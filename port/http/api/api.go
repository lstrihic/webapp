package api

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/lstrihic/webapp/port/http/api/utils"
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
		logger = utils.EnrichLogger(ctx, logger)

		var fiberError *fiber.Error
		var validationErrType validation.Errors
		var apiError *utils.Error

		if errors.As(err, &validationErrType) {
			logger.Warn().Interface("validation_errors", err).Msg("Bad request")
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(utils.Error{
				Code:    fiber.StatusUnprocessableEntity,
				Message: err,
			})
		} else if errors.As(err, &fiberError) {
			logger.Warn().Interface("error", fiberError).Msg("Error occurred")
			return ctx.Status(fiberError.Code).JSON(utils.Error{
				Code:    fiberError.Code,
				Message: fiberError.Message,
			})
		} else if errors.As(err, &apiError) {
			logger.Warn().Interface("error", apiError).Msg("Message occurred")
			return ctx.Status(apiError.Code).JSON(apiError)
		}
		logger.Error().Err(err).Msg("Internal sever error")
		return ctx.Status(fiber.StatusInternalServerError).JSON(utils.Error{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}
}
