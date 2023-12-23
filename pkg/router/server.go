package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luckyss6/cc_server/internal/handler"
	"github.com/luckyss6/cc_server/pkg/middleware"
)

func InitServer() *fiber.App {
	app := fiber.New()
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})
	UserRouter(app)
	return app
}

func UserRouter(app *fiber.App) {
	userHandler := handler.NewUserHandler()
	user := app.Group("/user")
	{
		user.Post("/login", userHandler.Login)
		user.Post("/register", userHandler.Register)
		user.Get("/ok", middleware.AuthRequired(), func(c *fiber.Ctx) error {
			return c.SendString("ok")
		})
	}
}
