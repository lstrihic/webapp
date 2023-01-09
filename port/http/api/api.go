package api

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/lstrihic/webapp/domain/user"
	"github.com/lstrihic/webapp/port/http/middleware/localization"
	"github.com/lstrihic/webapp/port/http/middleware/token"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type Route interface {
	Method() string
	Path() string
	IsSecure() bool
	Handler() fiber.Handler
}

// ErrorHandler http error handler.
var ErrorHandler = func(logger *zerolog.Logger) fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		logger = EnrichLogger(ctx, logger)

		var fiberError *fiber.Error
		var validationErrType validation.Errors

		if errors.As(err, &validationErrType) {
			logger.Warn().Interface("validation_errors", err).Msg("Validation failed")
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"code":   fiber.StatusUnprocessableEntity,
				"errors": err,
			})
		} else if errors.As(err, &fiberError) {
			logger.Warn().Interface("error", fiberError).Msg("Error occurred")
			return ctx.Status(fiberError.Code).JSON(fiberError)
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

// LocalizeMessage localize message.
func LocalizeMessage(ctx *fiber.Ctx, messageID, defaultMessage string) string {
	localizer, ok := ctx.Locals(localization.DefaultContextKey).(*i18n.Localizer)
	if !ok {
		return defaultMessage
	}
	message, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID: messageID,
	})
	if err != nil {
		return defaultMessage
	}
	return message
}
