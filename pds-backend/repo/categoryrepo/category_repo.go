package repo

import (
	orm "pds-backend/orm/gorm"
	"pds-backend/orm/gorm/model"

	"gorm.io/gorm"
)

type CategoryRepoEntity interface {
	GetAll() ([]model.Category, error)
	GetById(id uint) (*model.Category, error)
}

type CategoryRepo struct {
	orm *gorm.DB
}

func NewCategoryRepo(orm orm.GormOrmEntity) CategoryRepoEntity {
	return &CategoryRepo{orm: orm.GetOrm()}
}

func (c *CategoryRepo) GetAll() ([]model.Category, error) {
	var list []model.Category
	tx := c.orm.Begin()
	err := tx.Model(model.Category{}).Find(&list).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return list, nil
}

func (c *CategoryRepo) GetById(id uint) (*model.Category, error) {
	var category model.Category
	tx := c.orm.Begin()
	err := tx.Model(model.Category{}).Where("id = ?", id).First(&category).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &category, nil
}
