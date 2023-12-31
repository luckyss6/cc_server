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
	DockerRouter(app)
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

func DockerRouter(app *fiber.App) {
	dockerHandler := handler.NewDockerHandler()
	docker := app.Group("/docker")
	{
		docker.Get("/image/list", dockerHandler.GetImageList)
		docker.Get("/container/list", dockerHandler.GetContainerList)
		docker.Get("/container/:id", dockerHandler.GetContainerInfo)
	}
}

func MachineRouter(app *fiber.App) {
	machineHandler := handler.NewMachineHandler()
	machine := app.Group("/machine")
	{
		machine.Get("/info", machineHandler.GetMachineInfo)
	}
}
