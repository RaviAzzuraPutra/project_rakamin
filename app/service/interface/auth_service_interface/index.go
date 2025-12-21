package auth_service_interface

import (
	"last-project/app/models"
	"last-project/app/request"
)

type Auth_Interface_Service interface {
	Register(request *request.RegisterRequest) (*models.User, error)
	Login(request *request.LoginRequest) (string, error)
}
