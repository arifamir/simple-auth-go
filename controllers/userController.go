package controllers

import "github.com/gofiber/fiber"

func GetUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Get user", "data": nil})
}

func CreateUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Create user", "data": nil})
}

func UpdateUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Update user", "data": nil})
}

func DeleteUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Delete user", "data": nil})
}

func LoginUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Login user", "data": nil})
}
