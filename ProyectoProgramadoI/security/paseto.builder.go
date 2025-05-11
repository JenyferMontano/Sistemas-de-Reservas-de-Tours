package security

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

// PasetoBuilder implementa la interfaz Builder usando el esquema Paseto v2
type PasetoBuilder struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

// NewPasetoBuilder crea una nueva instancia de builder usando Paseto y una clave simétrica
func NewPasetoBuilder(symmetricKey string) (Builder, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("tamaño de la llave inválido: se requieren %d caracteres", chacha20poly1305.KeySize)
	}
	builder := &PasetoBuilder{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}
	return builder, nil
}

// CreateToken genera un token con username y email, válido por una duración específica
func (builder *PasetoBuilder) CreateToken(username string, email string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, email, duration)
	if err != nil {
		return "", err
	}
	return builder.paseto.Encrypt(builder.symmetricKey, payload, nil)
}

// VerifyToken verifica un token y retorna el payload si es válido
func (builder *PasetoBuilder) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}
	err := builder.paseto.Decrypt(token, builder.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}
	err = payload.Valid()
	if err != nil {
		return nil, err
	}
	return payload, nil
}
