package routes

import (
	"github.com/arifamir/simple-auth-go/controllers"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware/logger"
)

func Setup(app *fiber.App) {

	/**
	* Group is used for Routes with common prefix to define a new sub-router with optional middleware.
	**/

	// Base Api end point
	api := app.Group("/api", logger.New())
	api.Get("/", controllers.Welcome)

	// Authentication end points
	auth := api.Group("/auth")
	auth.Post("/login", controllers.LoginUser)
	auth.Post("/register", controllers.CreateUser)

	// User related end points
	user := api.Group("/user")
	user.Get("/:id", controllers.GetUser)
	user.Patch("/:id", controllers.UpdateUser)
	user.Delete("/:id", controllers.DeleteUser)

	// Book related end points
	book := api.Group("/book")
	book.Get("/:id", controllers.GetBook)
	book.Post("/", controllers.CreateBook)
	book.Patch("/:id", controllers.UpdateBook)
	book.Delete("/:id", controllers.DeleteBook)

}
