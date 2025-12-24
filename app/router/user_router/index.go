package user_router

import (
	"last-project/app/controller/user_controller"
	"last-project/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app *fiber.App, UserController *user_controller.User_Controller) {
	user := app.Group("/user", middleware.JWTMiddleware())

	user.Get("/", UserController.Get)
	user.Put("/update-user", UserController.Update)
}
