package controller

import "github.com/gofiber/fiber/v2"

func index(c *fiber.Ctx) error {
	return c.SendString("halo")
}
