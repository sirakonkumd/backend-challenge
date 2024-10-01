package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func SetupRouter(app *fiber.App) *fiber.App {
	app.Get("/pie", TDownloadPaymentController)
	return app
}

func TDownloadPaymentController(c *fiber.Ctx) error {
	resp, err := PieFireDireService()
	if err != nil {
		log.Error("PieFireDireService err ", err)
		return c.Status(fiber.StatusInternalServerError).JSON("fail to get meat")
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}
