package user_service

import (
	"errors"
	"last-project/app/models"
	"last-project/app/repository/interface/user_interface_repository"
	"last-project/app/request"
	"last-project/app/utils"
)

type User_Service struct {
	repository user_interface_repository.User_Interface_Repository
}

func NewUserServiceRegistry(repository user_interface_repository.User_Interface_Repository) *User_Service {
	return &User_Service{
		repository: repository,
	}
}

func (s *User_Service) Get(IDUser string) (*models.User, error) {
	user, err := s.repository.Get(IDUser)

	if err != nil {
		return nil, errors.New("An error occurred while retrieving user data. " + err.Error())
	}

	if user == nil {
		return nil, errors.New("User Data Not Found.")
	}

	return user, err
}

func (s *User_Service) Update(IDUser string, request *request.User_Request) (*models.User, error) {
	user, err := s.repository.Get(IDUser)

	if err != nil {
		return nil, errors.New("An error occurred while retrieving user data. " + err.Error())
	}

	if user == nil {
		return nil, errors.New("User Data Not Found.")
	}

	if request.Nama == nil || *request.Nama == "" {
		return nil, errors.New("Name Cannot Be Empty.")
	}

	if request.NoTelp == nil || *request.NoTelp == "" {
		return nil, errors.New("Phone Cannot Be Empty.")
	}

	if request.Password == nil || *request.Password == "" {
		return nil, errors.New("Password Cannot Be Empty.")
	}

	if request.Pekerjaan == nil || *request.Pekerjaan == "" {
		return nil, errors.New("Work Cannot Be Empty.")
	}

	if request.TanggalLahir.IsZero() {
		return nil, errors.New("Date of Birth Cannot Be Empty.")
	}

	if request.Tentang == nil || *request.Tentang == "" {
		return nil, errors.New("About Cannot Be Empty.")
	}

	if request.JenisKelamin == nil || *request.JenisKelamin == "" {
		return nil, errors.New("Gender Cannot Be Empty.")
	}

	hashingPassword, errHash := utils.HashPassword(*request.Password)

	if errHash != nil {
		return nil, errors.New("An error occurred during password hashing.")
	}

	user.Nama = request.Nama
	user.KataSandi = &hashingPassword
	user.NoTelp = request.NoTelp
	user.TanggalLahir = &request.TanggalLahir
	user.Tentang = request.Tentang
	user.Pekerjaan = request.Pekerjaan

	errUpdate := s.repository.Update(IDUser, user)

	return user, errUpdate
}
