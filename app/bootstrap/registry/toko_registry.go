package registry

import (
	"last-project/app/controller/toko_controller"
	"last-project/app/repository/toko_repository"
	"last-project/app/service/toko_service"
)

type Toko_Modules struct {
	TokoController *toko_controller.Toko_Controller
}

func Toko_Registry() *Toko_Modules {
	TokoRepository := toko_repository.NewTokoRpositoryRegistry()

	TokoService := toko_service.NewTokoServiceRegistry(TokoRepository)

	TokoController := toko_controller.NewTokoControllerRegistry(TokoService)

	return &Toko_Modules{
		TokoController: TokoController,
	}
}
