package user

import (
	"context"
	"github.com/lstrihic/webapp/adapter/db"
	"github.com/lstrihic/webapp/pkg/security"
	"go.uber.org/fx"
)

var Provider = fx.Provide(InitService)

type Service interface {
	CreateUser(ctx context.Context, form *Form) (*User, error)
}

type service struct {
	db db.DB
}

// InitService initialize user service.
func InitService(db db.DB) Service {
	return &service{
		db: db,
	}
}

// CreateUser create new user.
func (s *service) CreateUser(ctx context.Context, form *Form) (*User, error) {
	// hash password
	password, err := security.HashPassword(form.Password)
	if err != nil {
		return nil, err
	}

	result, err := s.db.Client().
		User.
		Create().
		SetUsername(form.Username).
		SetEmail(form.Email).
		SetPassword(string(password)).
		Save(ctx)

	return MapUserEntityToDto(result), err
}
