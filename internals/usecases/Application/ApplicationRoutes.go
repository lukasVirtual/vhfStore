package application

import "github.com/gofiber/fiber/v2"

func Routes(app fiber.Router) {
	appV1 := app.Group("/app")

	appV1.Get("/", func(c *fiber.Ctx) error { return nil })
}
