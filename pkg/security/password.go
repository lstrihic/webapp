package security

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {
	if password == "" {
		return nil, errors.New("The provided plain password is empty")
	}

	// hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 13)
	if err != nil {
		return nil, err
	}
	return hashedPassword, nil
}
