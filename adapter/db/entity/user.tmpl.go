package entity

import (
	"github.com/lstrihic/webapp/pkg/security"
	"golang.org/x/crypto/bcrypt"
)

func (u *User) ValidatePassword(password string) bool {
	bytePassword := []byte(password)
	bytePasswordHash := []byte(u.Password)

	// comparing the password with the hash
	err := bcrypt.CompareHashAndPassword(bytePasswordHash, bytePassword)

	// nil means it is a match
	return err == nil
}

func (u *User) SetPassword(password string) error {
	hashedPassword, err2 := security.HashPassword(password)
	if err2 != nil {
		return err2
	}

	u.Password = string(hashedPassword)
	return nil
}

func (u *User) RefreshTokenKey() error {
	u.TokenKey = security.RandomString(50)
	return nil
}
