package trx_service_interface

import (
	"last-project/app/models"
	"last-project/app/request"
)

type Trx_Service_Interface interface {
	Create(request *request.Create_Trx_Request, IDUser string) (*models.Trx, error)
	GetVeryDetailedTrx(IDTrx string) (*models.Trx, []models.DetailTrx, error)
}
