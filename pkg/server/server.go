package server

import "github.com/gofiber/fiber/v2"

func InitServer() *fiber.App {
	app := fiber.New()
	return app
}