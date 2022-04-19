package controller

import (
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	app.Get("/", index)
}
