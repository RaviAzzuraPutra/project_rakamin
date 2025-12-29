package foto_repository

import (
	"last-project/app/database"
	"last-project/app/models"

	"gorm.io/gorm"
)

type Foto_Repository struct {
	DB *gorm.DB
}

func NewFotoRepositoryRegistry() *Foto_Repository {
	return &Foto_Repository{
		DB: database.DB,
	}
}

func (repo *Foto_Repository) AddPhoto(foto *models.FotoProduk) error {
	errCreate := repo.DB.Table("foto_produks").Create(foto).Error

	return errCreate
}

func (repo *Foto_Repository) DeleteFoto(IDFoto string) error {

	var foto *models.FotoProduk

	errDelete := repo.DB.Table("foto_produks").Unscoped().Where("id = ?", IDFoto).Delete(&foto).Error

	return errDelete
}

func (repo *Foto_Repository) GetById(IDFoto string) (*models.FotoProduk, error) {

	var foto *models.FotoProduk

	errGet := repo.DB.Table("foto_produks").Where("id = ?", IDFoto).First(&foto).Error

	return foto, errGet
}

func (repo *Foto_Repository) GetByIdProduct(IDProduct string) ([]models.FotoProduk, error) {
	var foto []models.FotoProduk

	errGet := repo.DB.Table("foto_produks").Where("id_produk = ?", IDProduct).Find(&foto).Error

	return foto, errGet
}
