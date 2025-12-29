package foto_service

import (
	"errors"
	"last-project/app/models"
	"last-project/app/repository/interface/foto_interface_repository"
	"last-project/app/repository/interface/product_interface_repository"
	"last-project/app/repository/interface/toko_interface_repository"
	"last-project/app/request"
	"last-project/app/utils"
)

type Foto_Service struct {
	repository foto_interface_repository.Foto_Interface_Repository
	toko       toko_interface_repository.Toko_Interface_Repository
	product    product_interface_repository.Product_Interface_Repository
}

func NewFotoServiceRegistry(repository foto_interface_repository.Foto_Interface_Repository,
	toko toko_interface_repository.Toko_Interface_Repository,
	product product_interface_repository.Product_Interface_Repository) *Foto_Service {
	return &Foto_Service{
		repository: repository,
		toko:       toko,
		product:    product,
	}
}

func (s *Foto_Service) AddPhoto(request *request.Foto_Product_Request, IDUser string, IDProduct string) (*models.FotoProduk, error) {
	if request.URL == nil || *request.URL == "" {
		return nil, errors.New("Url Cannot Be Empty.")
	}

	IDToko, errToko := s.toko.GetTokoByUser(IDUser)

	if errToko != nil {
		return nil, errors.New("An error occurred while retrieving the store ID. " + errToko.Error())
	}

	if IDToko == nil {
		return nil, errors.New("Store Not Found.")
	}

	folderName := IDToko.NamaToko

	imageURL, errPhoto := utils.UploadFotoToCloudinary(*request.URL, *folderName, *request.URL)

	if errPhoto != nil {
		return nil, errors.New(errPhoto.Error())
	}

	uuid := utils.GenerateUUID()

	foto := &models.FotoProduk{
		ID:       &uuid,
		IDProduk: &IDProduct,
		URL:      &imageURL,
	}

	errCreate := s.repository.AddPhoto(foto)

	return foto, errCreate
}

func (s *Foto_Service) Delete(IDFoto string) error {
	foto, err := s.repository.GetById(IDFoto)

	if err != nil {
		return errors.New("An error occurred while retrieving Photo. " + err.Error())
	}

	if foto == nil {
		return errors.New("Photo Not Found.")
	}

	errDelete := s.repository.DeleteFoto(IDFoto)

	if errDelete != nil {
		errors.New("An error occurred while delete Photo. " + errDelete.Error())
	}

	return errDelete

}

func (s *Foto_Service) GetById(IDFoto string) (*models.FotoProduk, error) {
	foto, errGet := s.repository.GetById(IDFoto)

	if errGet != nil {
		return nil, errors.New("An error occurred while retrieving Photo. " + errGet.Error())
	}

	if foto == nil {
		return nil, errors.New("Photo Not Found.")
	}

	return foto, errGet
}

func (s *Foto_Service) GetByIdProduct(IDProduct string) ([]models.FotoProduk, error) {

	getProduct, errProduct := s.product.GetById(IDProduct)

	if errProduct != nil {
		return nil, errors.New("An error occurred while retrieving Product ID. " + errProduct.Error())
	}

	getFoto, errFoto := s.repository.GetByIdProduct(*getProduct.ID)

	if errFoto != nil {
		return nil, errors.New("An error occurred while retrieving Product ID. " + errFoto.Error())
	}

	return getFoto, errFoto
}
