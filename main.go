package main

import (
	"github.com/arifamir/simple-auth-go/db"
	"github.com/arifamir/simple-auth-go/routes"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware/cors"
	"log"
)

func main() {
	db.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	log.Fatal(app.Listen(":8000"))
}
