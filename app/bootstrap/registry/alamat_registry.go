package registry

import (
	"last-project/app/controller/alamat_controller"
	"last-project/app/repository/alamat_repository"
	"last-project/app/service/alamat_service"
)

type Alamat_Module struct {
	AlamatController *alamat_controller.Alamat_Controller
}

func Alamat_Registry() *Alamat_Module {
	AlamatRepository := alamat_repository.NewAlamatRepositoryRegistry()

	AlamatService := alamat_service.NewAlamatServiceRegistry(AlamatRepository)

	AlamatController := alamat_controller.NewAlamatControllerRegistry(AlamatService)

	return &Alamat_Module{
		AlamatController: AlamatController,
	}
}
