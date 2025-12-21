package auth_interface_repository

import "last-project/app/models"

type Auth_Interface_Repository interface {
	Register(register *models.User) error
	Login(email string) (*models.User, error)
}
