package token

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
	"tutorial.sqlc.dev/app/utils"
)

func TestJWTMaker(t *testing.T) {
	// 1. Setup Maker
	maker, err := NewJWTMaker(utils.RandomString(32)) // Added utils. prefix
	require.NoError(t, err)

	username := utils.RandomOwner()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	// 2. Create Token
	token, payload, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	// 3. Verify Token
	payload, err = maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	// 4. Check Payload Data
	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredJWTMoken(t *testing.T) {
	maker, err := NewJWTMaker(utils.RandomString(32))
	require.NoError(t, err)

	// Create a token that is already expired (-time.Minute)
	token, payload, err := maker.CreateToken(utils.RandomOwner(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	// Verify should return an error
	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	// In v5, the error will wrap jwt.ErrTokenExpired
	require.ErrorIs(t, err, jwt.ErrTokenExpired)
	require.Nil(t, payload)
}

func TestInvalidJWTTokenAlgNone(t *testing.T) {
	payload, err := NewPayload(utils.RandomOwner(), time.Minute)
	require.NoError(t, err)

	// Create a token with the 'none' algorithm
	// jwt.UnsafeAllowNoneSignatureType is required to even create this token for testing
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	maker, err := NewJWTMaker(utils.RandomString(32))
	require.NoError(t, err)

	// Verify should return an error because our keyFunc only allows HMAC (HS256)
	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.ErrorIs(t, err, jwt.ErrTokenUnverifiable) // v5 specific error for bad algs
	require.Nil(t, payload)
}
