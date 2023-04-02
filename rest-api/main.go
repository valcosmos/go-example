package main

import (
	"github.com/joho/godotenv"
	"github.com/valcosmos/go-example/rest-api/database"
)

func main() {
	// init app
	err := initApp()
	if err != nil {
		panic(err)
	}

	// defer close database
	defer database.CloseMongoDB()

	app := generateApp()

	// app := fiber.New()

	// app.Post("/", func(c *fiber.Ctx) error {
	// 	sampleDoc := bson.M{"name": "Sample todo"}

	// 	collection := database.GetCollection("todos")

	// 	nDoc, err := collection.InsertOne(context.TODO(), sampleDoc)

	// 	if err != nil {
	// 		return c.Status(fiber.StatusInternalServerError).SendString("Error inserting todo")
	// 	}

	// 	return c.JSON(nDoc)

	// c.SendString("hello World")
	// })

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World!")
	// })

	app.Listen(":3000")
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
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}
