package util

import "golang.org/x/crypto/bcrypt"

type PasswordUtil interface {
	PasswordVerify(hash, pw string) error
	PasswordGenerate(password string) ([]byte, error)
}

type passwordUtil struct {
}

func NewPasswordUtil() PasswordUtil {
	return passwordUtil{}
}

func (pu passwordUtil) PasswordVerify(hash, pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
}

func (pu passwordUtil) PasswordGenerate(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
