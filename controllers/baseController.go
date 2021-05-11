package controllers

import "github.com/gofiber/fiber"

func Hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Api is working!", "data": nil})
}
