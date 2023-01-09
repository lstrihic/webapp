package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lstrihic/webapp/domain/user"
	"github.com/lstrihic/webapp/port/http/api"
)

type post struct {
	service user.Service
}

func InitPostUser(service user.Service) api.Route {
	return &post{
		service: service,
	}
}

func (_ *post) Method() string {
	return fiber.MethodPost
}

func (_ *post) Path() string {
	return "/users"
}

func (_ *post) IsSecure() bool {
	return true
}

func (p *post) Handler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// unmarshal body
		var form user.Form
		if err := ctx.BodyParser(&form); err != nil {
			return err
		}

		// validate job
		if err := form.ValidateWithContext(ctx.Context()); err != nil {
			return err
		}

		// create user
		result, err := p.service.CreateUser(ctx.Context(), &form)
		if err != nil {
			return err
		}

		return ctx.Status(fiber.StatusOK).JSON(result)
	}
}
