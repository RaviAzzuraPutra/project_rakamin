package registry

import (
	"last-project/app/controller/trx_controller"
	"last-project/app/repository/product_repository"
	"last-project/app/repository/trx_repository"
	"last-project/app/service/trx_service"
)

type Trx_Module struct {
	TrxController *trx_controller.Trx_Controller
}

func Trx_Registry() *Trx_Module {
	TrxRepository := trx_repository.NewTrxRepositoryRegistry()
	ProductRepository := product_repository.NewProductRepositoryRegistry()

	TrxService := trx_service.NewTrxServiceRegistry(TrxRepository, ProductRepository)

	TrxController := trx_controller.NewTrxControllerRegistry(TrxService)

	return &Trx_Module{
		TrxController: TrxController,
	}

}
