package repo

import (
	"pds-backend/constant"
	orm "pds-backend/orm/gorm"
	"pds-backend/orm/gorm/model"

	"gorm.io/gorm"
)

type LikesRepoEntity interface {
	// StartAssociation(project *model.Project) error
	CheckIfUserLike(project *model.Project, user *model.User) bool
	Count(project *model.Project) uint
	Create(project *model.Project, user *model.User) error
	Delete(project *model.Project, user *model.User) error
}

type LikesRepo struct {
	orm *gorm.DB
}

func NewLikesRepo(orm orm.GormOrmEntity) LikesRepoEntity {
	return &LikesRepo{orm: orm.GetOrm()}
}

// func (l *LikesRepo) StartAssociation(project *model.Project) error {
// 	tx := l.orm.Begin()
// 	err := l.orm.Model(project).Association("Likes").Error
// 	if err != nil {
// 		fmt.Println(err)
// 		tx.Rollback()
// 		return err
// 	}
// 	tx.Commit()
//
// 	return nil
// }

func (l *LikesRepo) CheckIfUserLike(project *model.Project, user *model.User) bool {
	tx := l.orm.Begin()
	// err := l.orm.Model(project).Association("Likes").Find(user)
	// if err != nil {
	// 	fmt.Println(err)
	// 	tx.Rollback()
	// 	return err
	// }

	result := tx.Joins(`JOIN "likes" ON "likes"."user_id" = "user"."id" AND "likes"."project_id" = ?`, project.ID).Find(user)
	isUserLike := false
	if result.RowsAffected > 0 {
		isUserLike = true
	}
	tx.Commit()

	return isUserLike
}

func (l *LikesRepo) Count(project *model.Project) uint {
	tx := l.orm.Begin()
	count := tx.Model(project).Association("Likes").Count()
	tx.Commit()

	return uint(count)
}

func (l *LikesRepo) Create(project *model.Project, user *model.User) error {
	tx := l.orm.Begin()

	createLikeErr := tx.Model(project).Association("Likes").Append(user)
	if createLikeErr != nil {
		tx.Rollback()
		return createLikeErr
	}

	activityType := constant.LIKE
	activityHeader := constant.ActivityHeaderConstant(activityType, *project.Title, *user.FullName)
	createActivity := model.Activity{
		UserID: &user.ID,
		Type:   &activityType,
		Header: &activityHeader,
	}

	createActivityErr := tx.Create(&createActivity).Error
	if createActivityErr != nil {
		tx.Rollback()
		return createActivityErr
	}

	tx.Commit()
	return nil
}

func (l *LikesRepo) Delete(project *model.Project, user *model.User) error {
	tx := l.orm.Begin()

	deleteLikeErr := tx.Model(project).Association("Likes").Delete(user)
	if deleteLikeErr != nil {
		tx.Rollback()
		return deleteLikeErr
	}

	activityType := constant.DISLIKE
	activityHeader := constant.ActivityHeaderConstant(activityType, *project.Title, *user.FullName)

	createActivity := model.Activity{
		UserID: &user.ID,
		Type:   &activityType,
		Header: &activityHeader,
	}

	createActivityErr := tx.Create(&createActivity).Error
	if createActivityErr != nil {
		tx.Rollback()
		return createActivityErr
	}

	tx.Commit()
	return nil
}
