package repo

import (
	orm "pds-backend/orm/gorm"
	"pds-backend/orm/gorm/model"

	"gorm.io/gorm"
)

type DivisionRepoEntity interface {
	GetAll() ([]model.Division, error)
	GetById(id uint) (*model.Division, error)
}

type DivisionRepo struct {
	orm *gorm.DB
}

func NewDivisionRepo(orm orm.GormOrmEntity) DivisionRepoEntity {
	return &DivisionRepo{orm: orm.GetOrm()}
}

func (d *DivisionRepo) GetAll() ([]model.Division, error) {
	var list []model.Division
	tx := d.orm.Begin()
	err := tx.Model(model.Division{}).Find(&list).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return list, nil
}

func (d *DivisionRepo) GetById(id uint) (*model.Division, error) {
	var division model.Division
	tx := d.orm.Begin()
	err := tx.Model(model.Division{}).Where("id = ?", id).First(&division).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &division, nil
}
