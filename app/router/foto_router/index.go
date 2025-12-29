package foto_router

import (
	"last-project/app/controller/foto_controller"
	"last-project/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func FotoRouter(app *fiber.App, fotoController *foto_controller.Foto_Controller) {
	foto := app.Group("/foto", middleware.JWTMiddleware())

	foto.Post("/add-foto/:id", middleware.AdminMiddleware(), fotoController.Create)
	foto.Get("/:id", fotoController.GetById)
	foto.Get("/product/:id", fotoController.GetByIdProduct)
	foto.Delete("/delete-foto/:id", middleware.AdminMiddleware(), fotoController.Delete)
}
