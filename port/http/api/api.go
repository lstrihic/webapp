package api

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/lstrihic/webapp/domain/user"
	"github.com/lstrihic/webapp/port/http/middleware/token"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type Route interface {
	Method() string
	Path() string
	Handler() fiber.Handler
}

var validationErrType validation.Errors

// ErrorHandler http error handler.
var ErrorHandler = func(logger *zerolog.Logger) fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		logger = EnrichLogger(ctx, logger)
		if errors.As(err, &validationErrType) {
			logger.Warn().Interface("validation_errors", err).Msg("Validation failed")
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(err)
		}
		logger.Error().Err(err).Msg("Internal sever error")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}
}

// EnrichLogger enrich logger with fields.
func EnrichLogger(ctx *fiber.Ctx, logger *zerolog.Logger) *zerolog.Logger {
	enrichedLogger := logger.
		With().
		Str("path", ctx.Path()).
		Str("method", ctx.Method()).
		Interface("trace_id", ctx.Locals(requestid.ConfigDefault.ContextKey)).
		Logger()
	return &enrichedLogger
}

// GetUser get username from request context.
func GetUser(ctx *fiber.Ctx) *user.User {
	usr := ctx.Locals(token.DefaultContextKey)
	if usr != nil {
		return usr.(*user.User)
	}
	return nil
}
