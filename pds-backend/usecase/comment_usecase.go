package usecase

import (
	"pds-backend/orm/gorm/model"
	commentRepo "pds-backend/repo/commentrepo"
	projectRepo "pds-backend/repo/projectrepo"
	userRepo "pds-backend/repo/userrepo"
)

type CommentUsecaseEntity interface {
	GetAllComment(projectId int, pageNumber int, pageSize int) ([]model.CommentJointUser, *uint, error)
	CreateComment(projectId int, userEmail string, comment string) (*model.Comment, *string, error)
	DeleteComment(commentId int) error
}

type CommentUsecase struct {
	commentRepo commentRepo.CommentRepoEntity
	projectRepo projectRepo.ProjectRepoEntity
	userRepo    userRepo.UserRepoEntity
}

func NewCommentUsecase(commentRepo commentRepo.CommentRepoEntity, projectRepo projectRepo.ProjectRepoEntity, userRepo userRepo.UserRepoEntity) CommentUsecaseEntity {
	return &CommentUsecase{
		commentRepo: commentRepo,
		projectRepo: projectRepo,
		userRepo:    userRepo,
	}
}

func (c *CommentUsecase) GetAllComment(projectId int, pageNumber int, pageSize int) ([]model.CommentJointUser, *uint, error) {
	return c.commentRepo.GetAllWithPagination(uint(projectId), uint(pageNumber), uint(pageSize))
}

func (c *CommentUsecase) CreateComment(projectId int, userEmail string, commentBody string) (*model.Comment, *string, error) {
	getProject, err := c.projectRepo.FindProjectById(uint(projectId))
	if err != nil {
		return nil, nil, err
	}

	// getUserId, err := c.projectRepo.FindIdByEmail(userEmail)
	// if err != nil {
	// 	return nil, nil, err
	// }

	getUser, err := c.userRepo.FindOneByEmail(userEmail)
	if err != nil {
		return nil, nil, err
	}

	getUserPicture := getUser.Picture
	getProjectId := uint(projectId)
	newComment := model.Comment{
		UserID:    &getUser.ID,
		ProjectID: &getProjectId,
		Body:      &commentBody,
	}

	getProjectTitle := getProject.Title
	getUserName := getUser.FullName

	createdComment, err := c.commentRepo.CreateOne(newComment, *getProjectTitle, *getUserName)
	if err != nil {
		return nil, nil, err
	}

	return createdComment, getUserPicture, nil
}

func (c *CommentUsecase) DeleteComment(commentId int) error {
	return c.commentRepo.Delete(uint(commentId))
}
