package auth

import (
	"context"
	"github.com/lstrihic/webapp/adapter/db"
	"github.com/lstrihic/webapp/adapter/db/entity/session"
	"github.com/lstrihic/webapp/domain/user"
	"github.com/pkg/errors"
	"go.uber.org/fx"
)

var Provider = fx.Provide(InitAuth)

type Service interface {
	Authenticate(ctx context.Context, token string) (*user.User, error)
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

// Authenticate user by the token.
func (s *service) Authenticate(ctx context.Context, token string) (*user.User, error) {
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
