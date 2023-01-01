package ui

import (
	"embed"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"net/http"
	"path/filepath"
)

//go:embed all:dist
var distDir embed.FS

var Middleware = filesystem.New(filesystem.Config{
	Next: func(ctx *fiber.Ctx) bool {
		if _, err := distDir.Open(filepath.Join("dist", ctx.Path())); err != nil {
			ctx.Path("/")
		}
		return false
	},
	Root:       http.FS(distDir),
	PathPrefix: "dist",
	Browse:     true,
})
