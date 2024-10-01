package main

import (
	"backend-challenge/pieFireDire/services"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{})

	services.SetupRouter(app)

	if err := app.Listen(":9001"); err != nil {
		panic(fmt.Errorf("unable to start fiber: %s", err.Error()))
	}
}
