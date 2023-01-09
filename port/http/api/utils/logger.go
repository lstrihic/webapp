package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/rs/zerolog"
)

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
