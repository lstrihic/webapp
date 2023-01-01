package http

import (
	"github.com/lstrihic/webapp/port/http/api"
	"github.com/lstrihic/webapp/port/http/api/v1/health"
	"github.com/lstrihic/webapp/port/http/api/v1/users"
	"github.com/lstrihic/webapp/port/http/api/v1/users/me"
	"go.uber.org/fx"
)

var Provider = fx.Module("api",
	fx.Provide(
		AsRoute(health.InitGetHealth),
		AsRoute(me.InitGetMe),
		AsRoute(users.InitPostUser),
		fx.Annotate(
			InitServer,
			fx.ParamTags(`group:"routes"`),
		),
	),
)

func AsRoute(fn any) any {
	return fx.Annotate(
		fn,
		fx.As(new(api.Route)),
		fx.ResultTags(`group:"routes"`),
	)
}
