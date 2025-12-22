package toko_repository

import (
	"last-project/app/database"
	"last-project/app/models"

	"gorm.io/gorm"
)

type Toko_Repository struct {
	DB *gorm.DB
}

func NewTokoRpositoryRegistry() *Toko_Repository {
	return &Toko_Repository{
		DB: database.DB,
	}
}

func (repo *Toko_Repository) GetTokoByUser(userID string) (*models.Toko, error) {
	var toko models.Toko

	err := repo.DB.Table("tokos").Where("id_user = ?", userID).First(&toko).Error

	return &toko, err
}

func (repo *Toko_Repository) UpdateToko(toko *models.Toko) error {
	errUpdate := repo.DB.Table("tokos").Save(toko).Error

	return errUpdate
}
