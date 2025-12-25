package category_service

import (
	"errors"
	"last-project/app/models"
	"last-project/app/repository/interface/category_interface_repository"
	"last-project/app/request"
	"last-project/app/utils"
)

type Category_Service struct {
	repository category_interface_repository.Category_Interface_Repository
}

func NewCategoryServiceRegistry(repository category_interface_repository.Category_Interface_Repository) *Category_Service {
	return &Category_Service{
		repository: repository,
	}
}

func (s *Category_Service) Create(request *request.Category_Request) (*models.Category, error) {

	if request.NamaCategory == nil || *request.NamaCategory == "" {
		return nil, errors.New("Category Name Cannot Be Empty.")
	}

	uuid := utils.GenerateUUID()

	category := &models.Category{
		ID:           &uuid,
		NamaCategory: request.NamaCategory,
	}

	errCrate := s.repository.Create(category)

	if errCrate != nil {
		return nil, errors.New("An error occurred while creating category data." + errCrate.Error())
	}

	return category, errCrate
}

func (s *Category_Service) Get() ([]models.Category, error) {
	category, errGet := s.repository.Get()

	if category == nil {
		return nil, errors.New("Data Category Still Empty")
	}

	if errGet != nil {
		return nil, errors.New("An error occurred while retrieving category data." + errGet.Error())
	}

	return category, errGet
}

func (s *Category_Service) GetById(IDCategory string) (*models.Category, error) {
	category, errGet := s.repository.GetById(IDCategory)

	if category == nil {
		return nil, errors.New("Data Category Still Empty")
	}

	if errGet != nil {
		return nil, errors.New("An error occurred while retrieving category data." + errGet.Error())
	}

	return category, errGet
}

func (s *Category_Service) Update(IDCategory string, request *request.Category_Request) (*models.Category, error) {
	category, errGet := s.repository.GetById(IDCategory)

	if category == nil {
		return nil, errors.New("Data Category Still Empty")
	}

	if errGet != nil {
		return nil, errors.New("An error occurred while retrieving category data." + errGet.Error())
	}

	if request.NamaCategory == nil || *request.NamaCategory == "" {
		return nil, errors.New("Category Name Cannot Be Empty.")
	}

	category.NamaCategory = request.NamaCategory

	errUpdate := s.repository.Update(IDCategory, category)

	if errUpdate != nil {
		return nil, errors.New("An error occurred while update category data." + errUpdate.Error())
	}

	return category, errUpdate
}

func (s *Category_Service) Delete(IDCategory string) error {
	category, errGet := s.repository.GetById(IDCategory)

	if category == nil {
		return errors.New("Data Category Still Empty")
	}

	if errGet != nil {
		return errors.New("An error occurred while retrieving category data.")
	}

	errDelete := s.repository.Delete(IDCategory)

	if errDelete != nil {
		return errors.New("An error occurred while Delete category data." + errDelete.Error())
	}

	return errDelete
}
