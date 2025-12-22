package toko_service

import (
	"errors"
	"last-project/app/models"
	toko_interfcae_repository "last-project/app/repository/interface/toko_interface_repository"
	"last-project/app/request"
	"last-project/app/utils"
)

type Toko_Service struct {
	repository toko_interfcae_repository.Toko_Interface_Repository
}

func NewTokoServiceRegistry(repository toko_interfcae_repository.Toko_Interface_Repository) *Toko_Service {
	return &Toko_Service{
		repository: repository,
	}
}

func (s *Toko_Service) GetToko(userID string) (*models.Toko, error) {
	toko, err := s.repository.GetTokoByUser(userID)

	if err != nil {
		return nil, errors.New("An error occurred while retrieving store data. " + err.Error())
	}

	if toko == nil {
		return nil, errors.New("Store Data Not Found")
	}

	return toko, err
}

func (s *Toko_Service) UpdateToko(userID string, request *request.Toko_Request) (*models.Toko, error) {
	toko, err := s.repository.GetTokoByUser(userID)

	if err != nil {
		return nil, errors.New("An error occurred while retrieving store data. " + err.Error())
	}

	if toko == nil {
		return nil, errors.New("Store Data Not Found")
	}

	if request.UrlFoto == nil || *request.UrlFoto == "" {
		return nil, errors.New("No photo provided")
	}

	folderName := toko.NamaToko

	imageURL, errPhoto := utils.UploadFotoToCloudinary(*request.UrlFoto, *folderName, *request.UrlFoto)

	if errPhoto != nil {
		return nil, errors.New(errPhoto.Error())
	}

	toko.URLFoto = &imageURL

	errUpdate := s.repository.UpdateToko(toko)

	if errUpdate != nil {
		return nil, errors.New("An error occurred while updating the store. " + errUpdate.Error())
	}

	return toko, errUpdate

}
