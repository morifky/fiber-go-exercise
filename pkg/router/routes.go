package router

import (
	"fiber-go-exercise/pkg/handler"
	"fiber-go-exercise/utils/token"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v3"
)

func SetupRoutes(h *handler.Handler, app *fiber.App) {
	//Base
	app.Get("/", h.Home)

	//Middleware
	api := app.Group("/api", logger.New())
	jwtMiddleware := jwtware.New(jwtware.Config{
		SigningKey: token.New(h.GetHandlerCfg().JWTSecret).Secret,
	})

	//User
	user := api.Group("/users", jwtMiddleware)
	user.Get("/", h.GetAllUsers)
	user.Get("/:id", h.GetUserByID)
	user.Delete("/:id", h.RemoveUser)
	user.Patch("/:id", h.UpdateUser)

	//Auth
	auth := api.Group("auth")
	auth.Post("/login", h.SignIn)
	auth.Post("/register", h.AddUser)

	//Post
	post := api.Group("/posts", jwtMiddleware)
	post.Get("/", h.GetAllPosts)
	post.Post("/", h.AddPost)
	post.Get("/:id", h.GetPostByID)
	post.Delete("/:id", h.RemovePost)
	post.Patch("/:id", h.UpdatePost)
}
