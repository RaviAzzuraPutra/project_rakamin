package registry

import (
	"last-project/app/controller/foto_controller"
	"last-project/app/repository/foto_repository"
	"last-project/app/repository/product_repository"
	"last-project/app/repository/toko_repository"
	"last-project/app/service/foto_service"
)

type Foto_Module struct {
	FotoController *foto_controller.Foto_Controller
}

func Foto_Registry() *Foto_Module {
	TokoRpository := toko_repository.NewTokoRpositoryRegistry()
	ProductRepositoy := product_repository.NewProductRepositoryRegistry()
	FotoRepotiroty := foto_repository.NewFotoRepositoryRegistry()

	FotoService := foto_service.NewFotoServiceRegistry(FotoRepotiroty, TokoRpository, ProductRepositoy)

	FotoController := foto_controller.NewFotoControllerRegistry(FotoService)

	return &Foto_Module{
		FotoController: FotoController,
	}
}
