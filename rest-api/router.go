package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/valcosmos/go-example/rest-api/handlers"
)

func generateApp() *fiber.App {
	app := fiber.New()

	// create health check route
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})
	// create the library group and routes
	libGroup := app.Group("/library")

	libGroup.Get("/", handlers.TestHandler)

	return app
}
