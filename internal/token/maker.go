package token

import (
	"time"
)

type Maker interface {
	CreateToken(userID, RepresentativeID int32, duration time.Duration) (string, *Payload, error)

	VerifyToken(token string) (*Payload, error)
}
