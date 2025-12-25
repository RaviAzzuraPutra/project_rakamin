package registry

import (
	"last-project/app/controller/category_controller"
	"last-project/app/repository/category_repository"
	"last-project/app/service/category_service"
)

type Category_Module struct {
	CategoryController *category_controller.Category_Controller
}

func Category_Registry() *Category_Module {
	CategoryRepository := category_repository.NewCategoryRepositoryRegistry()

	CategoryService := category_service.NewCategoryServiceRegistry(CategoryRepository)

	CategoryController := category_controller.NewCategoryControllerRegistry(CategoryService)

	return &Category_Module{
		CategoryController: CategoryController,
	}
}
