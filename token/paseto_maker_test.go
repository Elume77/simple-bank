package token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"tutorial.sqlc.dev/app/utils"
)

func TestPasetoMaker(t *testing.T) {
	// Generate a random 32-character key for the test
	symmetricKey := utils.RandomString(32)
	maker, err := NewPasetoMaker(symmetricKey)
	require.NoError(t, err)

	username := utils.RandomOwner()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	// Act
	token, payload, err := maker.CreateToken(username, duration)

	// Assert - Token Creation
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	// Assert - Payload Data Integrity
	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)

	// Assert - Decryption Verification
	// This ensures the token string actually contains what the payload says it does
	decryptedPayload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, decryptedPayload)

	require.Equal(t, payload.ID, decryptedPayload.ID)
	require.Equal(t, payload.Username, decryptedPayload.Username)
	require.WithinDuration(t, payload.IssuedAt, decryptedPayload.IssuedAt, time.Second)
	require.WithinDuration(t, payload.ExpiredAt, decryptedPayload.ExpiredAt, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	maker, _ := NewPasetoMaker(utils.RandomString(32))

	// Create a token that expired 1 minute ago
	token, payload, err := maker.CreateToken(utils.RandomOwner(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	// Verify should fail
	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error()) // Assuming you defined ErrExpiredToken
	require.Nil(t, payload)
}
