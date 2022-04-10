package repo

import (
	orm "pds-backend/orm/gorm"
	"pds-backend/orm/gorm/model"

	"gorm.io/gorm"
)

type ViewsRepoEntity interface {
	Count(project *model.Project) uint
	Create(project *model.Project, user *model.User) error
}

type ViewsRepo struct {
	orm *gorm.DB
}

func NewViewsRepo(orm orm.GormOrmEntity) ViewsRepoEntity {
	return &ViewsRepo{orm: orm.GetOrm()}
}

func (v *ViewsRepo) Count(project *model.Project) uint {
	tx := v.orm.Begin()
	count := tx.Model(project).Association("Views").Count()
	tx.Commit()

	return uint(count)
}

func (v *ViewsRepo) Create(project *model.Project, user *model.User) error {
	tx := v.orm.Begin()
	err := tx.Model(project).Association("Views").Append(user)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
