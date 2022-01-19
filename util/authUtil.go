package util

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthUtil interface {
	GetToken(Email string) (*string, error)
}

type authUtil struct {
}

func NewAuthUtil() AuthUtil {
	return authUtil{}
}

func (au authUtil) GetToken(Email string) (*string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = Email
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString([]byte(os.Getenv("SIGNINGKEY")))
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}
