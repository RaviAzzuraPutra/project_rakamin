package alamat_router

import (
	"last-project/app/controller/alamat_controller"
	"last-project/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func AlamatRouter(app *fiber.App, alamatController *alamat_controller.Alamat_Controller) {
	alamat := app.Group("/alamat", middleware.JWTMiddleware())

	alamat.Post("/add-alamat", alamatController.Create)
	alamat.Get("/", alamatController.GetByIDUser)
	alamat.Get("/:id", alamatController.GetByIDUserAndIDAlamat)
	alamat.Put("/update-alamat/:id", alamatController.Update)
	alamat.Delete("/delete-alamat/:id", alamatController.Delete)
}
