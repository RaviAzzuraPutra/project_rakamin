package category_router

import (
	"last-project/app/controller/category_controller"
	"last-project/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func CategoryRouter(app *fiber.App, CategoryController *category_controller.Category_Controller) {
	category := app.Group("/category", middleware.JWTMiddleware())

	category.Get("/", CategoryController.Get)
	category.Post("/add-category", middleware.AdminMiddleware(), CategoryController.Create)
	category.Put("/update-category/:id", middleware.AdminMiddleware(), CategoryController.Update)
	category.Delete("/delete-category/:id", middleware.AdminMiddleware(), CategoryController.Delete)
	category.Get("/:id", CategoryController.GetById)
}
