package handler

import (
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	application "www.github.com/vhfStore/internals/usecases/Application"
)

func (h *Handler) Register(app *fiber.App) {
	log.Debug("Handler::Register Routes")
	group := app.Group("/api/v1")

	application.Routes(group)
}
