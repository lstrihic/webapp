package domain

import (
	"github.com/lstrihic/webapp/domain/auth"
	"github.com/lstrihic/webapp/domain/user"
	"go.uber.org/fx"
)

var Provider = fx.Module("service",
	auth.Provider,
	user.Provider,
)
