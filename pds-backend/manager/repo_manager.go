package manager

import (
	"pds-backend/config"
	categoryRepo "pds-backend/repo/categoryrepo"
	cloudrepo "pds-backend/repo/cloudrepo"
	commentRepo "pds-backend/repo/commentrepo"
	divisionRepo "pds-backend/repo/divisionrepo"
	likesRepo "pds-backend/repo/likesrepo"
	projectRepo "pds-backend/repo/projectrepo"
	roleRepo "pds-backend/repo/rolerepo"
	userRepo "pds-backend/repo/userrepo"
	viewsRepo "pds-backend/repo/viewsrepo"
)

type RepoManagerEntity interface {
	GetUserRepo() userRepo.UserRepoEntity
	GetDivisionRepo() divisionRepo.DivisionRepoEntity
	GetRoleRepo() roleRepo.RoleRepoEntity
	GetCategoryRepo() categoryRepo.CategoryRepoEntity
	GetProjectRepo() projectRepo.ProjectRepoEntity
	GetLikesRepo() likesRepo.LikesRepoEntity
	GetViewsRepo() viewsRepo.ViewsRepoEntity
	GetCommentRepo() commentRepo.CommentRepoEntity
	GetCloudRepo() cloudrepo.CloudRepoEntity
}

type RepoManager struct {
	userRepo     userRepo.UserRepoEntity
	divisionRepo divisionRepo.DivisionRepoEntity
	roleRepo     roleRepo.RoleRepoEntity
	categoryRepo categoryRepo.CategoryRepoEntity
	projectRepo  projectRepo.ProjectRepoEntity
	likesRepo    likesRepo.LikesRepoEntity
	viewsRepo    viewsRepo.ViewsRepoEntity
	commentRepo  commentRepo.CommentRepoEntity
	cloudRepo    cloudrepo.CloudRepoEntity
}

func NewRepoManager(infraManager InfraManagerEntity, appConfig config.AppConfigEntity) RepoManagerEntity {
	userRepo := userRepo.NewUserRepo(infraManager.GetOrmInfra().GetGormOrm())
	divisionRepo := divisionRepo.NewDivisionRepo(infraManager.GetOrmInfra().GetGormOrm())
	roleRepo := roleRepo.NewRoleRepo(infraManager.GetOrmInfra().GetGormOrm())
	categoryRepo := categoryRepo.NewCategoryRepo(infraManager.GetOrmInfra().GetGormOrm())
	projectRepo := projectRepo.NewProjectRepo(infraManager.GetOrmInfra().GetGormOrm())
	likesRepo := likesRepo.NewLikesRepo(infraManager.GetOrmInfra().GetGormOrm())
	viewsRepo := viewsRepo.NewViewsRepo(infraManager.GetOrmInfra().GetGormOrm())
	commentRepo := commentRepo.NewCommentRepo(infraManager.GetOrmInfra().GetGormOrm())
	cloudRepo := cloudrepo.NewCloudRepo(infraManager.GetCloudInfra().GetFirebaseCloudEngine(), appConfig)

	return &RepoManager{
		userRepo:     userRepo,
		divisionRepo: divisionRepo,
		roleRepo:     roleRepo,
		categoryRepo: categoryRepo,
		projectRepo:  projectRepo,
		likesRepo:    likesRepo,
		cloudRepo:    cloudRepo,
		viewsRepo:    viewsRepo,
		commentRepo:  commentRepo,
	}
}

func (r *RepoManager) GetUserRepo() userRepo.UserRepoEntity {
	return r.userRepo
}

func (r *RepoManager) GetDivisionRepo() divisionRepo.DivisionRepoEntity {
	return r.divisionRepo
}

func (r *RepoManager) GetRoleRepo() roleRepo.RoleRepoEntity {
	return r.roleRepo
}

func (r *RepoManager) GetCategoryRepo() categoryRepo.CategoryRepoEntity {
	return r.categoryRepo
}

func (r *RepoManager) GetLikesRepo() likesRepo.LikesRepoEntity {
	return r.likesRepo
}

func (r *RepoManager) GetViewsRepo() viewsRepo.ViewsRepoEntity {
	return r.viewsRepo
}

func (r *RepoManager) GetCommentRepo() commentRepo.CommentRepoEntity {
	return r.commentRepo
}

func (r *RepoManager) GetProjectRepo() projectRepo.ProjectRepoEntity {
	return r.projectRepo
}

func (r *RepoManager) GetCloudRepo() cloudrepo.CloudRepoEntity {
	return r.cloudRepo
}
