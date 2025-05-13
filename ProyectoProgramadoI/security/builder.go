package security

import "time"

type Builder interface {
	// CreateToken genera un token para un usuario con su username y email, válido durante una duración determinada.
	CreateToken(username string, email string, rol string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
