package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"tutorial.sqlc.dev/app/utils"
)

func TestCreateUser(t *testing.T) {
	createRandomUser(t)

}

func createRandomUser(t *testing.T) User {

	hashedpassword, err := utils.HashedPassword(utils.RandomString(6))
	require.NoError(t, err)
	arg := CreateUserParams{
		Username:       utils.RandomOwner(), // randomly generated data
		Hashedpassword: hashedpassword,
		Fullname:       utils.RandomOwner(),
		Email:          utils.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, arg.Username)
	require.Equal(t, arg.Username, user.Username)

	require.Equal(t, arg.Hashedpassword, user.Hashedpassword)
	require.Equal(t, arg.Fullname, user.Fullname)
	require.Equal(t, arg.Email, user.Email)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)

	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Hashedpassword, user2.Hashedpassword)
	require.Equal(t, user1.Fullname, user2.Fullname)
	require.Equal(t, user1.Email, user2.Email)
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)

}
