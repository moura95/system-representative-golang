package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type Payload struct {
	ID               uuid.UUID `json:"id"`
	UserID           int32     `json:"user_id"`
	RepresentativeID int32     `json:"representative_id"`
	IssuedAt         time.Time `json:"issued_at"`
	ExpiredAt        time.Time `json:"expired_at"`
}

func NewPayload(userID, representativeID int32, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:               tokenID,
		UserID:           userID,
		RepresentativeID: representativeID,
		IssuedAt:         time.Now(),
		ExpiredAt:        time.Now().Add(duration),
	}
	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
