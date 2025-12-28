package product_service

import (
	"errors"
	"last-project/app/models"
	"last-project/app/repository/interface/category_interface_repository"
	"last-project/app/repository/interface/product_interface_repository"
	"last-project/app/repository/interface/toko_interface_repository"
	"last-project/app/request"
	"last-project/app/utils"
)

type Product_Service struct {
	repository product_interface_repository.Product_Interface_Repository
	toko       toko_interface_repository.Toko_Interface_Repository
	category   category_interface_repository.Category_Interface_Repository
}

func NewProductServiceRegistry(repository product_interface_repository.Product_Interface_Repository,
	toko toko_interface_repository.Toko_Interface_Repository,
	category category_interface_repository.Category_Interface_Repository,
) *Product_Service {
	return &Product_Service{
		repository: repository,
		toko:       toko,
		category:   category,
	}
}

func (s *Product_Service) Create(request *request.Product_Request, IDUser string) (*models.Produk, error) {

	if request.NamaProduk == nil || *request.NamaProduk == "" {
		return nil, errors.New("Product Name Cannot Be Empty.")
	}

	if request.HargaReseller == nil || *request.HargaReseller == "" {
		return nil, errors.New("Reseller Price Cannot Be Empty.")
	}

	if request.HargaKonsumen == nil || *request.HargaKonsumen == "" {
		return nil, errors.New("Consumer Price Cannot Be Empty.")
	}

	if request.Stok == nil {
		return nil, errors.New("Stock Cannot Be Empty.")
	}

	if request.IDCategory == nil || *request.IDCategory == "" {
		return nil, errors.New("ID Category Cannot Be Empty.")
	}

	if request.Deskripsi == nil || *request.Deskripsi == "" {
		return nil, errors.New("Description Cannot Be Empty.")
	}

	uuid := utils.GenerateUUID()

	slug := utils.GenerateSlug(*request.NamaProduk)

	IDToko, errToko := s.toko.GetTokoByUser(IDUser)

	if errToko != nil {
		return nil, errors.New("An error occurred while retrieving the store ID. " + errToko.Error())
	}

	if IDToko == nil {
		return nil, errors.New("Store Not Found.")
	}

	IDCategory, errCategory := s.category.GetById(*request.IDCategory)

	if errCategory != nil {
		return nil, errors.New("An error occurred while retrieving the category ID. " + errCategory.Error())
	}

	if IDCategory == nil {
		return nil, errors.New("Category Not Found.")
	}

	product := &models.Produk{
		ID:            &uuid,
		Slug:          &slug,
		NamaProduk:    request.NamaProduk,
		HargaReseller: request.HargaReseller,
		HargaKonsumen: request.HargaKonsumen,
		Stok:          request.Stok,
		Deskripsi:     request.Deskripsi,
		IDToko:        IDToko.ID,
		IDCategory:    IDCategory.ID,
	}

	errCreate := s.repository.Create(product)

	if errCreate != nil {
		return nil, errors.New("An error occurred while create product " + errCreate.Error())
	}

	return product, errCreate
}

func (s *Product_Service) GetAll() ([]models.Produk, error) {

	product, err := s.repository.GetAll()

	if err != nil {
		return nil, errors.New("An error occurred while retrieving All Product. " + err.Error())
	}

	if product == nil {
		return nil, errors.New("Product Still Empty")
	}

	return product, err
}

func (s *Product_Service) GetByIdAndIdToko(IDProduct string, IDUser string) (*models.Produk, error) {

	IDToko, errToko := s.toko.GetTokoByUser(IDUser)

	if errToko != nil {
		return nil, errors.New("An error occurred while retrieving the store ID. " + errToko.Error())
	}

	if IDToko == nil {
		return nil, errors.New("Store Not Found.")
	}

	product, err := s.repository.GetByIdAndIdToko(IDProduct, *IDToko.ID)

	if err != nil {
		return nil, errors.New("An error occurred while retrieving the spesific Product. " + err.Error())
	}

	if product == nil {
		return nil, errors.New("Product Not Found")
	}

	return product, err
}

func (s *Product_Service) GetByIdToko(IDUser string) ([]models.Produk, error) {

	IDToko, errToko := s.toko.GetTokoByUser(IDUser)

	if errToko != nil {
		return nil, errors.New("An error occurred while retrieving the store ID. " + errToko.Error())
	}

	if IDToko == nil {
		return nil, errors.New("Store Not Found.")
	}

	product, err := s.repository.GetByIdToko(*IDToko.ID)

	if err != nil {
		return nil, errors.New("An error occurred while retrieving the spesific Product. " + err.Error())
	}

	if product == nil {
		return nil, errors.New("Product Not Found")
	}

	return product, err
}

func (s *Product_Service) GetById(IDProduct string) (*models.Produk, error) {

	product, err := s.repository.GetById(IDProduct)

	if err != nil {
		return nil, errors.New("An error occurred while retrieving the spesific Product. " + err.Error())
	}

	if product == nil {
		return nil, errors.New("Product Not Found")
	}

	return product, err
}

func (s *Product_Service) GetByCategory(search string) ([]models.Produk, error) {

	product, err := s.repository.GetByCategory(search)

	if err != nil {
		return nil, errors.New("An error occurred while retrieving Product By Categories. " + err.Error())
	}

	if product == nil {
		return nil, errors.New("Product Is Not Found")
	}

	return product, err
}

func (s *Product_Service) Update(IDProduct string, IDUser string, request *request.Product_Request) (*models.Produk, error) {

	IDToko, errToko := s.toko.GetTokoByUser(IDUser)

	if errToko != nil {
		return nil, errors.New("An error occurred while retrieving the store ID. " + errToko.Error())
	}

	if IDToko == nil {
		return nil, errors.New("Store Not Found.")
	}

	product, err := s.repository.GetByIdAndIdToko(IDProduct, *IDToko.ID)

	if err != nil {
		return nil, errors.New("An error occurred while retrieving the spesific Product. " + err.Error())
	}

	if product == nil {
		return nil, errors.New("Product Not Found")
	}

	if request.NamaProduk == nil || *request.NamaProduk == "" {
		return nil, errors.New("Product Name Cannot Be Empty.")
	}

	if request.HargaReseller == nil || *request.HargaReseller == "" {
		return nil, errors.New("Reseller Price Cannot Be Empty.")
	}

	if request.HargaKonsumen == nil || *request.HargaKonsumen == "" {
		return nil, errors.New("Consumer Price Cannot Be Empty.")
	}

	if request.Stok == nil {
		return nil, errors.New("Stock Cannot Be Empty.")
	}

	if request.IDCategory == nil || *request.IDCategory == "" {
		return nil, errors.New("ID Category Cannot Be Empty.")
	}

	if request.Deskripsi == nil || *request.Deskripsi == "" {
		return nil, errors.New("Description Cannot Be Empty.")
	}

	slug := utils.GenerateSlug(*request.NamaProduk)

	IDCategory, errCategory := s.category.GetById(*request.IDCategory)

	if errCategory != nil {
		return nil, errors.New("An error occurred while retrieving the category ID. " + errCategory.Error())
	}

	if IDCategory == nil {
		return nil, errors.New("Category Not Found.")
	}

	product.NamaProduk = request.NamaProduk
	product.Slug = &slug
	product.HargaReseller = request.HargaReseller
	product.HargaKonsumen = request.HargaKonsumen
	product.Stok = request.Stok
	product.Deskripsi = request.Deskripsi
	product.IDCategory = IDCategory.ID

	errUpdate := s.repository.Update(IDProduct, *IDToko.ID, product)

	return product, errUpdate

}

func (s *Product_Service) Delete(IDProduct string, IDUser string) error {

	IDToko, errToko := s.toko.GetTokoByUser(IDUser)

	if errToko != nil {
		return errors.New("An error occurred while retrieving the store ID. " + errToko.Error())
	}

	if IDToko == nil {
		return errors.New("Store Not Found.")
	}

	product, err := s.repository.GetByIdAndIdToko(IDProduct, *IDToko.ID)

	if err != nil {
		return errors.New("An error occurred while retrieving the spesific Product. " + err.Error())
	}

	if product == nil {
		return errors.New("Product Not Found")
	}

	errDelete := s.repository.Delete(IDProduct, *IDToko.ID)

	return errDelete
}
