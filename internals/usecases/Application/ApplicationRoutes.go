package application

import "github.com/gofiber/fiber/v2"

func Routes(app fiber.Router) {
	appGroup := app.Group("/app")

	appGroup.Get("/", func(c *fiber.Ctx) error { return nil })
}
