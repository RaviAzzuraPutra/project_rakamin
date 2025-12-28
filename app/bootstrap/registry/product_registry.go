package registry

import (
	"last-project/app/controller/product_controller"
	"last-project/app/repository/category_repository"
	"last-project/app/repository/product_repository"
	"last-project/app/repository/toko_repository"
	"last-project/app/service/product_service"
)

type Product_Module struct {
	Product_Controller *product_controller.Product_Controller
}

func Product_Registry() *Product_Module {
	TokoRepository := toko_repository.NewTokoRpositoryRegistry()
	CategoryRepository := category_repository.NewCategoryRepositoryRegistry()
	ProductRepository := product_repository.NewProductRepositoryRegistry()

	ProductService := product_service.NewProductServiceRegistry(ProductRepository, TokoRepository, CategoryRepository)

	ProductController := product_controller.NewProductControllerRegistry(ProductService)

	return &Product_Module{
		Product_Controller: ProductController,
	}
}
