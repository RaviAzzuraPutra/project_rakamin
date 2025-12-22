package middleware

import "github.com/gofiber/fiber/v2"

func AdminMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		isAdmin, ok := ctx.Locals("is_admin").(bool)

		if !ok || !isAdmin {
			return ctx.JSON(fiber.Map{
				"Status":  401,
				"Message": "You're not an admin mate.",
			})
		}

		return ctx.Next()
	}
}
