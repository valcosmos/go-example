package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type libraryDTO struct {
	Name    string `json:"name" bson:"name"`
	Address string `json:"address" bson:"address"`
}

func CreateLibrary(c *fiber.Ctx) error {
	nLibrary := new(libraryDTO)

	if err := c.BodyParser(nLibrary); err != nil {
		return err
	}

	fmt.Println(nLibrary)

	return c.SendString("Create library")
}
