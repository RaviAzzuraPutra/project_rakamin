package product_repository

import (
	"last-project/app/database"
	"last-project/app/models"

	"gorm.io/gorm"
)

type Product_Repository struct {
	DB *gorm.DB
}

func NewProductRepositoryRegistry() *Product_Repository {
	return &Product_Repository{
		DB: database.DB,
	}
}

func (repo *Product_Repository) Create(product *models.Produk) error {
	errProduct := repo.DB.Table("produks").Create(product).Error

	return errProduct
}

func (repo *Product_Repository) GetAll() ([]models.Produk, error) {

	var Product []models.Produk

	errGet := repo.DB.Table("produks").Find(&Product).Error

	return Product, errGet
}

func (repo *Product_Repository) GetByIdAndIdToko(IDProduct string, IDToko string) (*models.Produk, error) {
	var Product *models.Produk

	errGet := repo.DB.Table("produks").Where("id = ? and id_toko = ?", IDProduct, IDToko).First(&Product).Error

	return Product, errGet
}

func (repo *Product_Repository) GetByIdToko(IDToko string) ([]models.Produk, error) {
	var Product []models.Produk

	errGet := repo.DB.Table("produks").Where("id_toko = ?", IDToko).Find(&Product).Error

	return Product, errGet
}

func (repo *Product_Repository) GetById(IDProduct string) (*models.Produk, error) {
	var Product *models.Produk

	errGet := repo.DB.Table("produks").Where("id = ?", IDProduct).First(&Product).Error

	return Product, errGet
}

func (repo *Product_Repository) GetByCategory(search string) ([]models.Produk, error) {
	var Product []models.Produk

	errGet := repo.DB.Joins("JOIN categories ON categories.id = produks.id_category").Where("categories.nama_category LIKE ?", "%"+search+"%").Find(&Product).Error

	return Product, errGet
}

func (repo *Product_Repository) Update(IDProduct string, IDToko string, Product *models.Produk) error {

	errUpdate := repo.DB.Table("produks").Where("id = ? and id_toko = ?", IDProduct, IDToko).Updates(Product).Error

	return errUpdate
}

func (repo *Product_Repository) Delete(IDProduct string, IDToko string) error {

	var Product *models.Produk

	errDelete := repo.DB.Table("produks").Unscoped().Where("id = ? and id_toko = ?", IDProduct, IDToko).Delete(&Product).Error

	return errDelete
}

func (repo *Product_Repository) CreateFoto(foto *models.FotoProduk) error {
	errCreate := repo.DB.Table("foto_produks").Create(foto).Error

	return errCreate
}
