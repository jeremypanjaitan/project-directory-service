package usecase

import (
	"pds-backend/orm/gorm/model"
	repo "pds-backend/repo/rolerepo"
)

type RoleUsecaseEntity interface {
	GetAllRole() ([]model.Role, error)
	GetRoleByID(id uint) (*model.Role, error)
}

type RoleUsecase struct {
	roleRepo repo.RoleRepoEntity
}

func NewRoleUsecase(roleRepo repo.RoleRepoEntity) RoleUsecaseEntity {
	return &RoleUsecase{
		roleRepo: roleRepo,
	}
}

func (r *RoleUsecase) GetAllRole() ([]model.Role, error) {
	return r.roleRepo.GetAll()
}

func (r *RoleUsecase) GetRoleByID(id uint) (*model.Role, error) {
	return r.roleRepo.GetById(id)
}
