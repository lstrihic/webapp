package api

import "github.com/gofiber/fiber/v2"

type Route interface {
	Method() string
	Path() string
	Handler() fiber.Handler
}
