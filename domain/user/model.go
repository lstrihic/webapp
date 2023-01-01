package user

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/lstrihic/webapp/adapter/db/entity"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

// MapUserEntityToDto map user entity to dto.
func MapUserEntityToDto(user *entity.User) *User {
	return &User{
		ID:       user.ID,
		Username: user.Username,
	}
}

type Form struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ValidateWithContext validate form.
func (form *Form) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(ctx,
		form,
		validation.Field(&form.Username, validation.Required, validation.Length(1, 50), is.Alphanumeric),
		validation.Field(&form.Email, validation.Required, validation.Length(1, 255), is.Email),
		validation.Field(&form.Password, validation.Required, is.PrintableASCII),
	)
}
