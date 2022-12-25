package http

import (
	"context"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/lstrihic/webapp/pkg/config"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

var Provider = fx.Module("api",
	fx.Provide(InitServer),
)

type Server interface {
	Start()
}

type server struct {
	fiber     *fiber.App
	lifecycle fx.Lifecycle
	logger    *zerolog.Logger
	cfg       *config.Config
}

// InitServer initialize http server.
func InitServer(lifecycle fx.Lifecycle, logger *zerolog.Logger, cfg *config.Config) Server {
	app := fiber.New(fiber.Config{
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
		DisableStartupMessage: true,
	})
	app.Use(requestid.New())
	app.Use(recover.New())

	// TODO: register routes

	return &server{
		fiber:     app,
		lifecycle: lifecycle,
		logger:    logger,
		cfg:       cfg,
	}
}

// Start the server.
func (s *server) Start() {
	s.lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) (err error) {
			go func() {
				err = s.fiber.Listen(fmt.Sprintf(":%d", s.cfg.Port))
				if err != nil {
					return
				}
			}()
			s.logger.Info().Interface("config", s.cfg).Msg("Service started")
			return
		},
		OnStop: func(ctx context.Context) error {
			s.logger.Info().Msg("Shutting down server")
			return s.fiber.Shutdown()
		},
	})
}
