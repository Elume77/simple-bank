package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	mockdb "tutorial.sqlc.dev/app/db/mock"
	db "tutorial.sqlc.dev/app/db/sqlc"
)

func TestCreateTransferAPI(t *testing.T) {
	amount := int64(10)

	user1 := randomAccount()
	user2 := randomAccount()

	user1.Curency = "USD"
	user2.Curency = "USD"

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"from_account_id": user1.ID,
				"to_account_id":   user2.ID,
				"amount":          amount,
				"currency":        "USD",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(user1.ID)).Times(1).Return(user1, nil)
				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(user2.ID)).Times(1).Return(user2, nil)

				arg := db.TransferTxParams{
					FromAccountID: user1.ID,
					ToAccountID:   user2.ID,
					Amount:        amount,
				}
				store.EXPECT().TransferTx(gomock.Any(), gomock.Eq(arg)).Times(1).Return(db.TransferTxResult{}, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "Invalid Currency",
			body: gin.H{
				"from_account_id": user1.ID,
				"to_account_id":   user2.ID,
				"amount":          amount,
				"currency":        "XYZ",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().TransferTx(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Internal Error",
			body: gin.H{
				"from_account_id": user1.ID,
				"to_account_id":   user2.ID,
				"amount":          amount,
				"currency":        "USD",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetAccount(gomock.Any(), gomock.Any()).AnyTimes().Return(user1, nil)
				store.EXPECT().TransferTx(gomock.Any(), gomock.Any()).Times(1).Return(db.TransferTxResult{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			request, err := http.NewRequest(http.MethodPost, "/transfers", bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}
