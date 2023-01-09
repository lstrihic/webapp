package token

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/lstrihic/webapp/domain/user"
	"github.com/lstrihic/webapp/port/http/api/utils"
	"github.com/pkg/errors"
	"strings"
)

const DefaultContextKey = "user"

type UserFn func(ctx context.Context, token string) (*user.User, error)

type Config struct {
	AuthSchema     string
	ContextKey     string
	SuccessHandler fiber.Handler
	ErrorHandler   fiber.ErrorHandler
}

// New auth middleware.
func New(userFn UserFn, config ...Config) fiber.Handler {
	cfg := makeCfg(config)
	return func(ctx *fiber.Ctx) error {
		token, err := tokenFromHeader(ctx, fiber.HeaderAuthorization, cfg.AuthSchema)
		if err != nil {
			return cfg.ErrorHandler(ctx, err)
		}
		usr, err := userFn(ctx.Context(), token)
		if err != nil {
			return cfg.ErrorHandler(ctx, err)
		}
		ctx.Locals(cfg.ContextKey, usr)
		return cfg.SuccessHandler(ctx)
	}
}

// tokenFromHeader extract token from header.
func tokenFromHeader(ctx *fiber.Ctx, header, authScheme string) (string, error) {
	headerValue := ctx.Get(header)
	l := len(authScheme)
	if len(headerValue) > l+1 && strings.EqualFold(headerValue[:l], authScheme) {
		return strings.TrimSpace(headerValue[l:]), nil
	}
	return "", errors.New("Missing or malformed JWT")
}

// makeCfg create default config.
func makeCfg(config []Config) (cfg Config) {
	if len(config) > 0 {
		cfg = config[0]
	}
	if cfg.SuccessHandler == nil {
		cfg.SuccessHandler = func(ctx *fiber.Ctx) error {
			return ctx.Next()
		}
	}
	if cfg.AuthSchema == "" {
		cfg.AuthSchema = "Bearer"
	}
	if cfg.ContextKey == "" {
		cfg.ContextKey = DefaultContextKey
	}

	if cfg.ErrorHandler == nil {
		cfg.ErrorHandler = func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(utils.Error{
				Code:    fiber.StatusUnauthorized,
				Message: utils.LocalizeMessage(ctx, "Auth.FailedToAuthorize", "Failed to authorize user."),
			})
		}
	}
	return cfg
}

// GetUser get username from request context.
func GetUser(ctx *fiber.Ctx) *user.User {
	usr := ctx.Locals(DefaultContextKey)
	if usr != nil {
		return usr.(*user.User)
	}
	return nil
}
