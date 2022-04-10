package usecase

import (
	"pds-backend/orm/gorm/model"
	repo "pds-backend/repo/categoryrepo"
)

type CategoryUsecaseEntity interface {
	GetAllCategory() ([]model.Category, error)
	GetCategoryByID(id uint) (*model.Category, error)
}

type CategoryUsecase struct {
	categoryRepo repo.CategoryRepoEntity
}

func NewCategoryUsecase(categoryRepo repo.CategoryRepoEntity) CategoryUsecaseEntity {
	return &CategoryUsecase{
		categoryRepo: categoryRepo,
	}
}

func (c *CategoryUsecase) GetAllCategory() ([]model.Category, error) {
	return c.categoryRepo.GetAll()
}

func (c *CategoryUsecase) GetCategoryByID(id uint) (*model.Category, error) {
	return c.categoryRepo.GetById(id)
}
