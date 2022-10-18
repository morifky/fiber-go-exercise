package router

import (
	"fiber-go-exercise/pkg/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(h *handler.Handler, app *fiber.App) {
	//Base
	app.Get("/", h.Home)
	//Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", h.Home)

	//User
	user := api.Group("/user")
	user.Get("/", h.GetAllUsers)
}
