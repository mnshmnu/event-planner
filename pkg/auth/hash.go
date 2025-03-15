package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func (a *auth) CompareHash(hPass []byte, pass []byte) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hPass), []byte(pass)); err != nil {
		return false, errors.New("invalid credentials")
	}
	return true, nil
}

func (a *auth) GenerateHash(pass string) ([]byte, error) {
	hPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hPass, nil
}
