package http

import (
	"context"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/lstrihic/webapp/domain/auth"
	"github.com/lstrihic/webapp/pkg/config"
	"github.com/lstrihic/webapp/pkg/localization"
	"github.com/lstrihic/webapp/port/http/api"
	localizationMiddleware "github.com/lstrihic/webapp/port/http/middleware/localization"
	"github.com/lstrihic/webapp/port/http/middleware/token"
	"github.com/lstrihic/webapp/ui"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

type Server interface {
	Start() error
}

type server struct {
	fiber     *fiber.App
	lifecycle fx.Lifecycle
	logger    *zerolog.Logger
	cfg       *config.Config
	auth      auth.Service
}

// InitServer initialize http server.
func InitServer(
	routes []api.Route,
	lifecycle fx.Lifecycle,
	logger *zerolog.Logger,
	cfg *config.Config,
	auth auth.Service,
	localization localization.Service,
) Server {
	app := fiber.New(fiber.Config{
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
		DisableStartupMessage: true,
		ErrorHandler:          api.ErrorHandler(logger),
	})
	app.Use(favicon.New())
	app.Use(requestid.New())
	app.Use(etag.New())
	app.Use(localizationMiddleware.New(localization.GetBundle()))
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	// register routes
	authMiddleware := token.New(auth.Authorize)
	v1Group := app.Group("/api/v1")
	for _, route := range routes {
		var handlers []fiber.Handler
		if route.IsSecure() {
			handlers = append(handlers, authMiddleware)
		}
		handlers = append(handlers, route.Handler())

		v1Group.Add(route.Method(), route.Path(), handlers...)
	}

	// server UI
	app.Use("/", ui.Middleware)

	return &server{
		fiber:     app,
		lifecycle: lifecycle,
		logger:    logger,
		cfg:       cfg,
	}
}

// Start the server.
func (s *server) Start() error {
	s.lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go s.run()
			s.logger.Info().Interface("config", s.cfg).Msg("Server started")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			s.logger.Info().Msg("Shutting down server")
			return s.fiber.Shutdown()
		},
	})
	return nil
}

// run the server.
func (s *server) run() {
	_ = s.fiber.Listen(fmt.Sprintf(":%d", s.cfg.Port))
}
