package domain

import (
	"github.com/lstrihic/webapp/domain/auth"
	"go.uber.org/fx"
)

var Provider = fx.Module("service",
	auth.Provider,
)
