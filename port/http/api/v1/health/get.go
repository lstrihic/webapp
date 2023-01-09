package health

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lstrihic/webapp/port/http/api"
)

type get struct{}

func InitGetHealth() api.Route {
	return &get{}
}

func (_ *get) Method() string {
	return fiber.MethodGet
}

func (_ *get) Path() string {
	return "/health"
}

func (_ *get) IsSecure() bool {
	return false
}

func (_ *get) Handler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "OK",
		})
	}
}
