package usecase

import (
	"pds-backend/orm/gorm/model"
	repo "pds-backend/repo/projectrepo"
	"strconv"
)

type ProjectUsecaseEntity interface {
	CreateProject(project model.Project) (*model.Project, error)
	FindProjectById(id uint) (*model.Project, error)
	FindIdByEmail(email string) (*uint, error)
	GetProjectDetails(projectId uint) (*model.ProjectWithLikeViewComment, error)
	ShowListProject(pageNumber uint, pageSize uint, titleSearch string, category string, sortLike string) ([]model.ProjectWithLikeViewComment, *uint, error)
	DeleteProject(id uint) error
	UpdateProject(Updatedproject model.Project, id uint) (*model.Project, error)
	FindProjectOwner(id uint) (*model.User, error)
}

type ProjectUsecase struct {
	projectRepo repo.ProjectRepoEntity
}

func NewProjectUsecase(projectRepo repo.ProjectRepoEntity) ProjectUsecaseEntity {
	return &ProjectUsecase{projectRepo: projectRepo}
}

func (p *ProjectUsecase) CreateProject(project model.Project) (*model.Project, error) {
	return p.projectRepo.CreateOne(project)
}

func (p *ProjectUsecase) FindProjectById(id uint) (*model.Project, error) {
	return p.projectRepo.FindProjectById(id)
}

func (p *ProjectUsecase) GetProjectDetails(projectId uint) (*model.ProjectWithLikeViewComment, error) {
	return p.projectRepo.GetProjectDetails(projectId)
}

func (p *ProjectUsecase) FindIdByEmail(email string) (*uint, error) {
	return p.projectRepo.FindIdByEmail(email)
}

func (p *ProjectUsecase) ShowListProject(pageNumber uint, pageSize uint, titleSearch string, category string, sortLike string) ([]model.ProjectWithLikeViewComment, *uint, error) {
	if category != "" {
		categoryNum, _ := strconv.Atoi(category)
		return p.projectRepo.ShowListProjectWithPaginationByCategory(pageNumber, pageSize, uint(categoryNum), titleSearch, sortLike)
	}
	return p.projectRepo.ShowListProjectWithPagination(pageNumber, pageSize, titleSearch, sortLike)
}

func (p *ProjectUsecase) DeleteProject(id uint) error {
	return p.projectRepo.DeleteProjectById(id)
}

func (p *ProjectUsecase) UpdateProject(Updatedproject model.Project, id uint) (*model.Project, error) {
	return p.projectRepo.UpdateProjectById(Updatedproject, id)
}

func (p *ProjectUsecase) FindProjectOwner(id uint) (*model.User, error) {
	return p.projectRepo.FindOwnerByProjectId(id)
}
