package auth_service

import (
	"errors"
	"last-project/app/models"
	"last-project/app/repository/interface/auth_interface_repository"
	"last-project/app/request"
	"last-project/app/utils"
)

type Auth_Service struct {
	repository auth_interface_repository.Auth_Interface_Repository
}

func NewAuthServiceRegistry(repository auth_interface_repository.Auth_Interface_Repository) *Auth_Service {
	return &Auth_Service{
		repository: repository,
	}
}

func (s *Auth_Service) Register(request *request.RegisterRequest) (*models.User, error) {

	if request.Password == nil || *request.Password == "" {
		return nil, errors.New("Password Cannot Be Empty")
	}

	hashedPassword, errHash := utils.HashPassword(*request.Password)

	if errHash != nil {
		return nil, errors.New("An error occurred during password hashing.")
	}

	if request.Nama == nil || *request.Nama == "" {
		return nil, errors.New("Name Cannot Be Empty")
	}

	if request.NoTelp == nil || *request.NoTelp == "" {
		return nil, errors.New("Phone numbers cannot be empty")
	}

	if request.Email == nil || *request.Email == " " {
		return nil, errors.New("Email numbers cannot be empty")
	}

	userID := utils.GenerateUUID()
	tokoID := utils.GenerateUUID()
	Nama_Toko := "TOKO " + *request.Nama

	user := models.User{
		ID:        &userID,
		Nama:      request.Nama,
		KataSandi: &hashedPassword,
		NoTelp:    request.NoTelp,
		Email:     request.Email,
		IsAdmin:   *request.IsAdmin,
		Toko: &models.Toko{
			ID:       &tokoID,
			IDUser:   &userID,
			NamaToko: &Nama_Toko,
		},
	}

	errRegister := s.repository.Register(&user)

	if errRegister != nil {
		return nil, errors.New("An error occurred during user registration. " + errRegister.Error())
	}

	return &user, errRegister

}

func (s *Auth_Service) Login(request *request.LoginRequest) (string, error) {
	login, errLogin := s.repository.Login(*request.Email)

	if errLogin != nil {
		return "", errors.New("An error occurred during login. " + errLogin.Error())
	}

	if request.Email == nil || *request.Email == "" {
		return "", errors.New("Email numbers cannot be empty")
	}

	if request.Password == nil || *request.Password == "" {
		return "", errors.New("Password Cannot Be Empty")
	}

	errPasswordChecked := utils.CheckPassword(*login.KataSandi, *request.Password)

	if errPasswordChecked != nil {
		return "", errors.New("Password Does Not Match " + errPasswordChecked.Error())
	}

	JWT, errJWT := utils.GenerateJWT(*login.ID, login.IsAdmin)

	if errJWT != nil {
		return "", errors.New("An error occurred while generating JWT. " + errJWT.Error())
	}

	return JWT, errJWT
}
