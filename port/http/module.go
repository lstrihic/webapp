package http

import (
	"github.com/lstrihic/webapp/port/http/api"
	"github.com/lstrihic/webapp/port/http/api/v1/health"
	"go.uber.org/fx"
)

var Provider = fx.Module("api",
	fx.Provide(
		AsRoute(health.InitHealth),
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
