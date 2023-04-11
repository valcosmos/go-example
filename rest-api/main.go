package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/valcosmos/go-example/rest-api/database"
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

func main() {
	// init app
	err := initApp()
	if err != nil {
		panic(err)
	}

	// defer close database
	defer database.CloseMongoDB()

	app := generateApp()

	port := os.Getenv("PORT")

	app.Listen(":" + port)
}

func initApp() error {
	err := loadENV()
	if err != nil {
		return err
	}
	err = database.StartMongoDB()
	if err != nil {
		return err
	}
	return nil
}

func loadENV() error {
	goEnv := os.Getenv("GO_ENV")
	if goEnv == "" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}
	return nil
}
