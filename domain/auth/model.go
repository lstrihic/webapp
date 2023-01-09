package auth

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Token struct {
	Token    string `json:"token"`
	Username string `json:"username"`
}

func MapSessionToTokenDto(token, username string) *Token {
	return &Token{
		Token:    token,
		Username: username,
	}
}

type Form struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ValidateWithContext validate form.
func (form *Form) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(ctx,
		form,
		validation.Field(&form.Username, validation.Required),
		validation.Field(&form.Password, validation.Required),
	)
}
