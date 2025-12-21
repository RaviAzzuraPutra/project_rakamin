package auth_router

import (
	"last-project/app/controller/auth_controller"

	"github.com/gofiber/fiber/v2"
)

func AuthRouter(app *fiber.App, auth *auth_controller.Auth_Controller) {
	api := app.Group("/auth")

	api.Post("/register", auth.Register)
	api.Post("Login", auth.Login)
}
