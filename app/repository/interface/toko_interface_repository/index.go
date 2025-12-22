package toko_interfcae_repository

import "last-project/app/models"

type Toko_Interface_Repository interface {
	GetTokoByUser(userID string) (*models.Toko, error)
	UpdateToko(toko *models.Toko) error
}
