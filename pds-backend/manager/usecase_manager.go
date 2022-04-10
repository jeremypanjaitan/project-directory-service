package manager

import (
	"pds-backend/config"
	"pds-backend/usecase"
)

type UsecaseManagerEntity interface {
	GetAuthUseCase() usecase.AuthUsecaseEntity
	GetDivisionUseCase() usecase.DivisionUsecaseEntity
	GetRoleUseCase() usecase.RoleUsecaseEntity
	GetCategoryUseCase() usecase.CategoryUsecaseEntity
	GetUserUsecase() usecase.UserUsecaseEntity
	GetProjectUsecase() usecase.ProjectUsecaseEntity
	GetLikesUsecase() usecase.LikesUsecaseEntity
	GetViewsUsecase() usecase.ViewsUsecaseEntity
	GetCommentUsecase() usecase.CommentUsecaseEntity
	GetCloudUsecase() usecase.CloudUsecaseEntity
}

type UsecaseManager struct {
	authUsecase     usecase.AuthUsecaseEntity
	divisionUsecase usecase.DivisionUsecaseEntity
	roleUsecase     usecase.RoleUsecaseEntity
	categoryUsecase usecase.CategoryUsecaseEntity
	userUsecase     usecase.UserUsecaseEntity
	projectUsecase  usecase.ProjectUsecaseEntity
	likesUsecase    usecase.LikesUsecaseEntity
	viewsUsecase    usecase.ViewsUsecaseEntity
	commentUsecase  usecase.CommentUsecaseEntity
	cloudUsecase    usecase.CloudUsecaseEntity
}

func NewUsecaseManager(repoManager RepoManagerEntity, serviceManager ServiceManagerEntity, appConfig config.AppConfigEntity) UsecaseManagerEntity {
	authUsecase := usecase.NewAuthUsecase(serviceManager.GetTokenService(), repoManager.GetUserRepo())
	divisionUsecase := usecase.NewDivisionUsecase(repoManager.GetDivisionRepo())
	roleUsecase := usecase.NewRoleUsecase(repoManager.GetRoleRepo())
	categoryUsecase := usecase.NewCategoryUsecase(repoManager.GetCategoryRepo())
	userUsecase := usecase.NewUserUsecase(repoManager.GetUserRepo())
	projectUsecase := usecase.NewProjectUsecase(repoManager.GetProjectRepo())
	likesUsecase := usecase.NewLikesUsecase(repoManager.GetLikesRepo(), repoManager.GetProjectRepo(), repoManager.GetUserRepo())
	viewsUsecase := usecase.NewViewsUsecase(repoManager.GetViewsRepo(), repoManager.GetProjectRepo(), repoManager.GetUserRepo())
	commentUsecase := usecase.NewCommentUsecase(repoManager.GetCommentRepo(), repoManager.GetProjectRepo(), repoManager.GetUserRepo())
	cloudUsecase := usecase.NewCloudUsecase(repoManager.GetCloudRepo(), appConfig)
	return &UsecaseManager{
		authUsecase:     authUsecase,
		divisionUsecase: divisionUsecase,
		roleUsecase:     roleUsecase,
		categoryUsecase: categoryUsecase,
		userUsecase:     userUsecase,
		projectUsecase:  projectUsecase,
		likesUsecase:    likesUsecase,
		viewsUsecase:    viewsUsecase,
		commentUsecase:  commentUsecase,
		cloudUsecase:    cloudUsecase,
	}
}

func (n *UsecaseManager) GetAuthUseCase() usecase.AuthUsecaseEntity {
	return n.authUsecase
}

func (n *UsecaseManager) GetDivisionUseCase() usecase.DivisionUsecaseEntity {
	return n.divisionUsecase
}

func (n *UsecaseManager) GetRoleUseCase() usecase.RoleUsecaseEntity {
	return n.roleUsecase
}

func (n *UsecaseManager) GetCategoryUseCase() usecase.CategoryUsecaseEntity {
	return n.categoryUsecase
}

func (n *UsecaseManager) GetUserUsecase() usecase.UserUsecaseEntity {
	return n.userUsecase
}

func (n *UsecaseManager) GetProjectUsecase() usecase.ProjectUsecaseEntity {
	return n.projectUsecase
}

func (n *UsecaseManager) GetLikesUsecase() usecase.LikesUsecaseEntity {
	return n.likesUsecase
}

func (n *UsecaseManager) GetViewsUsecase() usecase.ViewsUsecaseEntity {
	return n.viewsUsecase
}

func (n *UsecaseManager) GetCommentUsecase() usecase.CommentUsecaseEntity {
	return n.commentUsecase
}

func (n *UsecaseManager) GetCloudUsecase() usecase.CloudUsecaseEntity {
	return n.cloudUsecase
}
