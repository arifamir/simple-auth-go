package controllers

import "github.com/gofiber/fiber"

func GetBook(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Get book", "data": nil})
}

func CreateBook(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Create book", "data": nil})
}

func UpdateBook(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Update book", "data": nil})
}

func DeleteBook(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Delete book", "data": nil})
}
