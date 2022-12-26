package health

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lstrihic/webapp/port/http/api"
)

type health struct {
}

func InitHealth() api.Route {
	return &health{}
}

func (p *health) Method() string {
	return fiber.MethodGet
}

func (p *health) Path() string {
	return "/health"
}

func (p *health) Handler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "OK",
		})
	}
}
