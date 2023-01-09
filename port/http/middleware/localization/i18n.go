package localization

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

const DefaultContextKey = "i18n"

type Config struct {
	ContextKey string
}

func New(bundle *i18n.Bundle, config ...Config) fiber.Handler {
	cfg := makeCfg(config...)
	return func(ctx *fiber.Ctx) error {
		lang := ctx.Query("lang")
		accept := ctx.Get("Accept-Language")
		localizer := i18n.NewLocalizer(bundle, lang, accept)
		ctx.Locals(cfg.ContextKey, localizer)
		return ctx.Next()
	}
}

func makeCfg(config ...Config) (cfg Config) {
	if len(config) > 0 {
		cfg = config[0]
	}
	if cfg.ContextKey == "" {
		cfg.ContextKey = DefaultContextKey
	}
	return
}
