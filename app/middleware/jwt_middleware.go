package middleware

import (
	"last-project/app/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func JWTMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authHeader := ctx.Get("Authorization")

		if authHeader == "" {
			return ctx.JSON(fiber.Map{
				"status":  401,
				"Message": "Unauthorized",
			})
		}

		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			return ctx.JSON(fiber.Map{
				"status":  401,
				"Message": "Invalid Format",
			})
		}

		tokenString := parts[1]

		claims, errClaim := utils.ParseJWT(tokenString)

		if errClaim != nil {
			return ctx.JSON(fiber.Map{
				"status":  401,
				"Message": "Token is invalid or has expired.",
			})
		}

		ctx.Locals("user_id", claims.UserID)
		ctx.Locals("is_admin", claims.IsAdmin)

		return ctx.Next()
	}
}
