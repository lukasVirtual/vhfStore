package handler

import (
	"github.com/gofiber/fiber/v2"
	application "www.github.com/vhfStore/internals/usecases/Application"
)

func (h *Handler) Register(app *fiber.App) {
	group := app.Group("/api/v1")

	application.Routes(group)
}
