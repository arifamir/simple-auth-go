package controllers

import "github.com/gofiber/fiber"

func Welcome(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Welcome to simple auth application", "data": nil})
}
