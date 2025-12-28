package product_router

import (
	"last-project/app/controller/product_controller"
	"last-project/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func ProductRouter(app *fiber.App, productController *product_controller.Product_Controller) {
	product := app.Group("/product", middleware.JWTMiddleware())

	product.Post("/add-product", middleware.AdminMiddleware(), productController.Create)
	product.Get("/", productController.GetAll)
	product.Get("/product-toko/:id", middleware.AdminMiddleware(), productController.GetByIdAndIdToko)
	product.Get("/product-toko", middleware.AdminMiddleware(), productController.GetByIdToko)
	product.Get("/search", productController.GetByCategory)
	product.Get("/:id", productController.GetById)
	product.Put("/update-product/:id", middleware.AdminMiddleware(), productController.Update)
	product.Delete("/delete-product/:id", middleware.AdminMiddleware(), productController.Delete)
}
