package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Specific errors for token validation
var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

// Payload contains the payload data of the token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// NewPayload creates a new token payload with a specific username and duration
func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

// --- JWT v5 Interface Methods ---

// GetExpirationTime is required for the jwt.Claims interface in v5
func (payload *Payload) GetExpirationTime() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(payload.ExpiredAt), nil
}

// GetIssuedAt is required for the jwt.Claims interface in v5
func (payload *Payload) GetIssuedAt() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(payload.IssuedAt), nil
}

// GetNotBefore is required but we can return nil if not used
func (payload *Payload) GetNotBefore() (*jwt.NumericDate, error) {
	return nil, nil
}

// GetIssuer is required but we can return an empty string if not used
func (payload *Payload) GetIssuer() (string, error) {
	return "", nil
}

// GetSubject is required but we can return an empty string if not used
func (payload *Payload) GetSubject() (string, error) {
	return "", nil
}

// GetAudience is required but we can return nil if not used
func (payload *Payload) GetAudience() (jwt.ClaimStrings, error) {
	return nil, nil
}

// Valid checks if the token payload is valid or not (custom logic)
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
