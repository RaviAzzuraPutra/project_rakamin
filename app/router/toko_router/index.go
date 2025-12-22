package toko_router

import (
	"last-project/app/controller/toko_controller"
	"last-project/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func TokoRouter(app *fiber.App, tokoController *toko_controller.Toko_Controller) {
	toko := app.Group("/toko", middleware.JWTMiddleware())

	toko.Get("/", tokoController.GetToko)
	toko.Put("/photo", tokoController.UpdateToko)
}
