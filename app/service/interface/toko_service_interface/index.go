package toko_service_interface

import (
	"last-project/app/models"
	"last-project/app/request"
)

type Toko_Service_Interface interface {
	GetToko(userID string) (*models.Toko, error)
	UpdateToko(userID string, request *request.Toko_Request) (*models.Toko, error)
}
