package usecase

import (
	"pds-backend/orm/gorm/model"
	repo "pds-backend/repo/divisionrepo"
)

type DivisionUsecaseEntity interface {
	GetAllDivision() ([]model.Division, error)
	GetDivisionByID(id uint) (*model.Division, error)
}

type DivisionUsecase struct {
	divisionRepo repo.DivisionRepoEntity
}

func NewDivisionUsecase(divisionRepo repo.DivisionRepoEntity) DivisionUsecaseEntity {
	return &DivisionUsecase{
		divisionRepo: divisionRepo,
	}
}

func (d *DivisionUsecase) GetAllDivision() ([]model.Division, error) {
	return d.divisionRepo.GetAll()
}

func (d *DivisionUsecase) GetDivisionByID(id uint) (*model.Division, error) {
	return d.divisionRepo.GetById(id)
}
