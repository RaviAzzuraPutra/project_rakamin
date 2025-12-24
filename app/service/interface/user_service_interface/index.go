package user_service_interface

import (
	"last-project/app/models"
	"last-project/app/request"
)

type User_Service_Interface interface {
	Get(IDUser string) (*models.User, error)
	Update(IDUser string, request *request.User_Request) (*models.User, error)
}
