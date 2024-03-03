package jwt

import (
	"fmt"
	"otus_highload/internal/domain"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func New() *Jwt {
	return &Jwt{
		duration: time.Hour,
		secret:   "--SECRET--",
	}
}

type Jwt struct {
	duration time.Duration
	secret   string
}

func (j Jwt) NewUserToken(user domain.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(j.duration).Unix()
	tokenString, err := token.SignedString([]byte(j.secret))
	if err != nil {
		return "", fmt.Errorf("token signed; %w", err)
	}
	return tokenString, nil
}
