package jwt

import (
	"numtostr/gotodo/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// TokenPayload defines the payload for the token
type TokenPayload struct {
	ID string
}

// Generate generates the jwt token based on payload
func Generate(payload TokenPayload) string {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		"ID":  payload.ID,
	})

	token, err := t.SignedString([]byte(config.TOKENKEY))

	if err != nil {
		panic(err)
	}

	return token
}
