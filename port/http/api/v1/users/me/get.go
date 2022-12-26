package me

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lstrihic/webapp/port/http/api"
)

type get struct {
}

func InitGetMe() api.Route {
	return &get{}
}

func (_ *get) Method() string {
	return fiber.MethodGet
}

func (_ *get) Path() string {
	return "/users/me"
}

func (_ *get) Handler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(api.GetUser(ctx))
	}
}
