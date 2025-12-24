package user_interface_repository

import "last-project/app/models"

type User_Interface_Repository interface {
	Update(IDUser string, user *models.User) error
	Get(IDUser string) (*models.User, error)
}
