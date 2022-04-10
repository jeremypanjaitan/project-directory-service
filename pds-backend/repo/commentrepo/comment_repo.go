package repo

import (
	"pds-backend/constant"
	orm "pds-backend/orm/gorm"
	"pds-backend/orm/gorm/model"

	"gorm.io/gorm"
)

type CommentRepoEntity interface {
	GetAll(projectId uint) ([]model.CommentJointUser, error)
	GetAllWithPagination(projectId uint, pageNumber uint, pageSize uint) ([]model.CommentJointUser, *uint, error)
	CreateOne(comment model.Comment, projectTitle string, userName string) (*model.Comment, error)
	Delete(commentId uint) error
}

type CommentRepo struct {
	orm *gorm.DB
}

func NewCommentRepo(orm orm.GormOrmEntity) CommentRepoEntity {
	return &CommentRepo{orm: orm.GetOrm()}
}

func (c *CommentRepo) GetAll(projectId uint) ([]model.CommentJointUser, error) {
	var list []model.CommentJointUser
	tx := c.orm.Begin()

	err := tx.Model(model.Comment{}).Select(`full_name, picture, "comment"."id", body, "comment"."created_at"`).Joins(`join "user" on "user"."id" = "comment"."user_id"`).Where(`"comment"."project_id" = ?`, projectId).Order(`"comment"."created_at"`).Scan(&list).Error

	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return list, nil
}

func (c *CommentRepo) GetAllWithPagination(projectId uint, pageNumber uint, pageSize uint) ([]model.CommentJointUser, *uint, error) {
	var list []model.CommentJointUser
	var count int64

	tx := c.orm.Begin()
	tx.Model(&model.Comment{}).Where(`project_id = ?`, projectId).Count(&count)

	totalOffset := (pageSize * pageNumber) - pageSize
	totalPagination := uint(count / int64(pageSize))
	if (count % int64(pageSize)) != 0 {
		totalPagination += 1
	}

	err := tx.Model(model.Comment{}).Select(`full_name, picture, "comment"."id", body, "comment"."created_at"`).
		Joins(`join "user" on "user"."id" = "comment"."user_id"`).Where(`"comment"."project_id" = ?`, projectId).
		Order(`"comment"."created_at" DESC`).Limit(int(pageSize)).Offset(int(totalOffset)).Scan(&list).Error

	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}
	tx.Commit()
	return list, &totalPagination, nil
}

func (c *CommentRepo) CreateOne(comment model.Comment, projectTitle string, userName string) (*model.Comment, error) {
	tx := c.orm.Begin()
	createCommentErr := tx.Create(&comment).Error
	if createCommentErr != nil {
		tx.Rollback()
		return nil, createCommentErr
	}

	activityType := constant.COMMENTED
	activityHeader := constant.ActivityHeaderConstant(activityType, projectTitle, userName)

	createActivity := model.Activity{
		UserID: comment.UserID,
		Type:   &activityType,
		Header: &activityHeader,
		Body:   comment.Body,
	}

	createActivityErr := tx.Create(&createActivity).Error
	if createActivityErr != nil {
		tx.Rollback()
		return nil, createActivityErr
	}

	tx.Commit()
	return &comment, nil
}

func (c *CommentRepo) Delete(commentId uint) error {
	tx := c.orm.Begin()
	err := tx.Delete(&model.Comment{}, commentId).Error

	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
