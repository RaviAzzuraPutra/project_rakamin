package trx_router

import (
	"last-project/app/controller/trx_controller"
	"last-project/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func TrxRouter(app *fiber.App, trxController *trx_controller.Trx_Controller) {
	trx := app.Group("/trx", middleware.JWTMiddleware())

	trx.Post("/add-trx", trxController.Create)
	trx.Get("/:id", middleware.AdminMiddleware(), trxController.GetVeryDetailedTrx)
}
