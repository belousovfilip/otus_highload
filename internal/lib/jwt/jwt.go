package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"otus_highload/internal/models"
	"time"
)

func NewToken(user models.User, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.Id
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(duration).Unix()
	tokenString, err := token.SignedString([]byte("--SECRET--"))
	if err != nil {
		return "", fmt.Errorf("token signed; %s", err)
	}
	return tokenString, nil
}
