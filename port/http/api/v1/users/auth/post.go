package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lstrihic/webapp/domain/auth"
	"github.com/lstrihic/webapp/port/http/api"
	"github.com/lstrihic/webapp/port/http/api/utils"
	"github.com/rs/zerolog"
)

type post struct {
	service auth.Service
	logger  *zerolog.Logger
}

func InitPostAuth(service auth.Service, logger *zerolog.Logger) api.Route {
	return &post{
		service: service,
		logger:  logger,
	}
}

func (_ *post) Method() string {
	return fiber.MethodPost
}

func (_ *post) Path() string {
	return "/users/auth"
}

func (_ *post) IsSecure() bool {
	return false
}

func (p *post) Handler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// logger
		logger := utils.EnrichLogger(ctx, p.logger)

		// unmarshal body
		var form auth.Form
		if err := ctx.BodyParser(&form); err != nil {
			return fiber.ErrUnprocessableEntity
		}

		// validate form
		if err := form.ValidateWithContext(ctx.Context()); err != nil {
			return err
		}

		// create session
		token, err := p.service.AuthorizeUser(ctx.Context(), &form)
		if err != nil {
			logger.Warn().Err(err).Msg("Invalid username or password")

			return &utils.Error{
				Code:    fiber.StatusUnauthorized,
				Message: utils.LocalizeMessage(ctx, "Auth.InvalidCredentials", "Invalid user credentials."),
			}
		}

		return ctx.Status(fiber.StatusCreated).JSON(token)
	}
}
