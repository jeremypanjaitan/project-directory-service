package usecase

import (
	projectRepo "pds-backend/repo/projectrepo"
	userRepo "pds-backend/repo/userrepo"
	viewsRepo "pds-backend/repo/viewsrepo"
)

type ViewsUsecaseEntity interface {
	CountViewsByProjectId(projectId uint) (uint, error)
	CreateViews(projectId uint, userEmail string) error
}

type ViewsUsecase struct {
	viewsRepo   viewsRepo.ViewsRepoEntity
	projectRepo projectRepo.ProjectRepoEntity
	userRepo    userRepo.UserRepoEntity
}

func NewViewsUsecase(viewsRepo viewsRepo.ViewsRepoEntity, projectRepo projectRepo.ProjectRepoEntity, userRepo userRepo.UserRepoEntity) ViewsUsecaseEntity {
	return &ViewsUsecase{
		viewsRepo:   viewsRepo,
		projectRepo: projectRepo,
		userRepo:    userRepo,
	}
}

func (v *ViewsUsecase) CountViewsByProjectId(projectId uint) (uint, error) {
	count := uint(0)

	getProject, err := v.projectRepo.FindProjectById(projectId)
	if err != nil {
		return count, err
	}

	count = v.viewsRepo.Count(getProject)

	return count, nil
}

func (v *ViewsUsecase) CreateViews(projectId uint, userEmail string) error {
	getProject, err := v.projectRepo.FindProjectById(projectId)
	if err != nil {
		return err
	}

	// getUserId, err := v.projectRepo.FindIdByEmail(userEmail)
	// if err != nil {
	// 	return err
	// }

	getUser, err := v.userRepo.FindOneByEmail(userEmail)
	if err != nil {
		return err
	}

	err = v.viewsRepo.Create(getProject, getUser)
	if err != nil {
		return err
	}
	return nil
}
