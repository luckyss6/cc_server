package middleware

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/luckyss6/cc_server/pkg/constant"
)

func AuthRequired() func(ctx *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			JWTAlg: "HS256",
			Key:    []byte(constant.JwtSecret),
		},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		},
	})
}
