package auth

import (
	"context"
	"github.com/lstrihic/webapp/adapter/db"
	"github.com/lstrihic/webapp/adapter/db/entity/session"
	userSchema "github.com/lstrihic/webapp/adapter/db/entity/user"
	"github.com/lstrihic/webapp/domain/user"
	"github.com/lstrihic/webapp/pkg/security"
	"github.com/pkg/errors"
	"go.uber.org/fx"
)

var Provider = fx.Provide(InitAuth)

type Service interface {
	Authorize(ctx context.Context, token string) (*user.User, error)
	AuthorizeUser(ctx context.Context, form *Form) (*Token, error)
}

type service struct {
	db db.DB
}

// InitAuth initialize service.
func InitAuth(db db.DB) Service {
	return &service{
		db: db,
	}
}

// Authorize user by the token.
func (s *service) Authorize(ctx context.Context, token string) (*user.User, error) {
	result, err := s.db.Client().
		Session.
		Query().
		WithUser().
		Where(session.Token(token)).
		First(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find user session")
	}
	if !result.IsValid {
		return nil, errors.New("session is not valid")
	}
	return user.MapUserEntityToDto(result.Edges.User), nil
}

// AuthorizeUser authorize user by username and password.
func (s *service) AuthorizeUser(ctx context.Context, form *Form) (*Token, error) {
	// find user based on user form
	foundUser, err := s.db.Client().User.
		Query().
		Where(userSchema.Username(form.Username)).
		First(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch user")
	}

	// validate password
	if !foundUser.ValidatePassword(form.Password) {
		return nil, errors.New("invalid password")
	}

	// create token
	token := security.RandomString(50)

	// create new session
	if err = s.db.Client().
		Session.
		Create().
		SetToken(token).
		SetUser(foundUser).
		Exec(ctx); err != nil {
		return nil, errors.Wrap(err, "failed to create token")
	}

	return MapSessionToTokenDto(token, foundUser.Username), nil
}
