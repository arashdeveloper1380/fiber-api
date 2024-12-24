package middlewares

import "github.com/gofiber/fiber/v2"

func XToken(ctx *fiber.Ctx) error {
	token := ctx.Get("x-token")
	if token != "secret" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}
	return ctx.Next()
}
