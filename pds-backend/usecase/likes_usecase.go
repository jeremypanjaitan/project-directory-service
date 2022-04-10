package usecase

import (
	likeRepo "pds-backend/repo/likesrepo"
	projectRepo "pds-backend/repo/projectrepo"
	userRepo "pds-backend/repo/userrepo"
)

type LikesUsecaseEntity interface {
	GetLikeByProjectId(projectId uint, userEmail string) (uint, bool, error)
	CreateLike(projectId uint, userEmail string) error
	DeleteLike(projectId uint, userEmail string) error
}

type LikesUsecase struct {
	likesRepo   likeRepo.LikesRepoEntity
	projectRepo projectRepo.ProjectRepoEntity
	userRepo    userRepo.UserRepoEntity
}

func NewLikesUsecase(likesRepo likeRepo.LikesRepoEntity, projectRepo projectRepo.ProjectRepoEntity, userRepo userRepo.UserRepoEntity) LikesUsecaseEntity {
	return &LikesUsecase{
		likesRepo:   likesRepo,
		projectRepo: projectRepo,
		userRepo:    userRepo,
	}
}

func (l *LikesUsecase) GetLikeByProjectId(projectId uint, userEmail string) (uint, bool, error) {
	getProject, err := l.projectRepo.FindProjectById(projectId)
	if err != nil {
		return 0, false, err
	}

	count := uint(0)

	// getUserId, err := l.projectRepo.FindIdByEmail(userEmail)
	// if err != nil {
	// 	return count, false, nil
	// }

	getUser, err := l.userRepo.FindOneByEmail(userEmail)
	if err != nil {
		return count, false, err
	}

	count = l.likesRepo.Count(getProject)

	isUserLike := l.likesRepo.CheckIfUserLike(getProject, getUser)

	return count, isUserLike, nil
}

func (l *LikesUsecase) CreateLike(projectId uint, userEmail string) error {
	getProject, err := l.projectRepo.FindProjectById(projectId)
	if err != nil {
		return err
	}

	// getUserId, err := l.projectRepo.FindIdByEmail(userEmail)
	// if err != nil {
	// 	return err
	// }

	getUser, err := l.userRepo.FindOneByEmail(userEmail)
	if err != nil {
		return err
	}

	err = l.likesRepo.Create(getProject, getUser)
	if err != nil {
		return err
	}
	return nil
}

func (l *LikesUsecase) DeleteLike(projectId uint, userEmail string) error {
	getProject, err := l.projectRepo.FindProjectById(projectId)
	if err != nil {
		return err
	}

	// getUserId, err := l.projectRepo.FindIdByEmail(userEmail)
	// if err != nil {
	// 	return err
	// }

	getUser, err := l.userRepo.FindOneByEmail(userEmail)
	if err != nil {
		return err
	}

	err = l.likesRepo.Delete(getProject, getUser)
	if err != nil {
		return err
	}
	return nil
}
