package alamat_service

import (
	"errors"
	"last-project/app/models"
	"last-project/app/repository/interface/alamat_interface_repository"
	"last-project/app/request"
	"last-project/app/utils"
)

type Alamat_Service struct {
	repository alamat_interface_repository.Alamat_Interface_Repository
}

func NewAlamatServiceRegistry(repository alamat_interface_repository.Alamat_Interface_Repository) *Alamat_Service {
	return &Alamat_Service{
		repository: repository,
	}
}

func (s *Alamat_Service) Create(request *request.Alamat, IDUserParameter string) (*models.Alamat, error) {

	if request.JudulAlamat == nil || *request.JudulAlamat == "" {
		return nil, errors.New("Title Address Cannot Be Empty.")
	}

	if request.NamaPenerima == nil || *request.NamaPenerima == "" {
		return nil, errors.New("Recipient's Name Cannot Be Empty.")
	}

	if request.NoTelp == nil || *request.NoTelp == "" {
		return nil, errors.New("Phone Number Cannot Be Empty.")
	}

	if request.DetailAlamat == nil || *request.DetailAlamat == "" {
		return nil, errors.New("Detail Of Address Cannot Be Empty.")
	}

	uuid := utils.GenerateUUID()

	alamat := models.Alamat{
		ID:           &uuid,
		IDUser:       &IDUserParameter,
		JudulAlamat:  request.JudulAlamat,
		NamaPenerima: request.NamaPenerima,
		NoTelp:       request.NoTelp,
		DetailAlamat: request.DetailAlamat,
	}

	errCreate := s.repository.Create(&alamat)

	if errCreate != nil {
		return nil, errors.New("An error occurred while create address data. " + errCreate.Error())
	}

	return &alamat, errCreate
}

func (s *Alamat_Service) GetByIDUserAndIDAlamat(IDUserParameter string, IDAlamat string) (*models.Alamat, error) {

	alamat, errGet := s.repository.GetIDUserAndIDAlamat(IDUserParameter, IDAlamat)

	if alamat == nil {
		return nil, errors.New("Address Is Empty.")
	}

	if errGet != nil {
		return nil, errors.New("An error occurred while retrieving address data.")
	}

	return alamat, errGet
}

func (s *Alamat_Service) GetByIDUser(IDUserParameter string) ([]models.Alamat, error) {

	alamat, errGet := s.repository.GetIDUser(IDUserParameter)

	if alamat == nil {
		return nil, errors.New("Address Is Empty.")
	}

	if errGet != nil {
		return nil, errors.New("An error occurred while retrieving address data. " + errGet.Error())
	}

	return alamat, errGet
}

func (s *Alamat_Service) Update(IDUserParameter string, IDAlamat string, request *request.Alamat) (*models.Alamat, error) {
	getAlamat, errGet := s.repository.GetIDUserAndIDAlamat(IDUserParameter, IDAlamat)

	if getAlamat == nil {
		return nil, errors.New("Address Not Found")
	}

	if errGet != nil {
		return nil, errors.New("An error occurred while retrieving address data. " + errGet.Error())
	}

	if request.JudulAlamat == nil || *request.JudulAlamat == "" {
		return nil, errors.New("Title Address Cannot Be Empty.")
	}

	if request.NamaPenerima == nil || *request.NamaPenerima == "" {
		return nil, errors.New("Recipient's Name Cannot Be Empty.")
	}

	if request.NoTelp == nil || *request.NoTelp == "" {
		return nil, errors.New("Phone Number Cannot Be Empty.")
	}

	if request.DetailAlamat == nil || *request.DetailAlamat == "" {
		return nil, errors.New("Detail Of Address Cannot Be Empty.")
	}

	getAlamat.JudulAlamat = request.JudulAlamat
	getAlamat.NamaPenerima = request.NamaPenerima
	getAlamat.NoTelp = request.NoTelp
	getAlamat.DetailAlamat = request.DetailAlamat

	errUpdate := s.repository.Update(IDAlamat, IDUserParameter, getAlamat)

	return getAlamat, errUpdate
}

func (s *Alamat_Service) Delete(IDUserParameter string, IDAlamat string) error {
	getAlamat, errGet := s.repository.GetIDUserAndIDAlamat(IDUserParameter, IDAlamat)

	if getAlamat == nil {
		return errors.New("Address Not Found")
	}

	if errGet != nil {
		return errors.New("An error occurred while retrieving address data." + errGet.Error())
	}

	errDelete := s.repository.Delete(IDAlamat, IDUserParameter)

	if errDelete != nil {
		return errors.New("An error occurred while delete address data." + errDelete.Error())
	}

	return nil

}
