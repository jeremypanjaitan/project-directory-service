package repo

import (
	orm "pds-backend/orm/gorm"
	"pds-backend/orm/gorm/model"

	"gorm.io/gorm"
)

type RoleRepoEntity interface {
	GetAll() ([]model.Role, error)
	GetById(id uint) (*model.Role, error)
}

type RoleRepo struct {
	orm *gorm.DB
}

func NewRoleRepo(orm orm.GormOrmEntity) RoleRepoEntity {
	return &RoleRepo{orm: orm.GetOrm()}
}

func (r *RoleRepo) GetAll() ([]model.Role, error) {
	var list []model.Role
	tx := r.orm.Begin()
	err := tx.Model(model.Role{}).Find(&list).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return list, nil
}

func (r *RoleRepo) GetById(id uint) (*model.Role, error) {
	var role model.Role
	tx := r.orm.Begin()
	err := tx.Model(model.Role{}).Where("id = ?", id).First(&role).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &role, nil
}
