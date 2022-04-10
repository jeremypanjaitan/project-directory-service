package repo

import (
	"log"
	orm "pds-backend/orm/gorm"
	"pds-backend/orm/gorm/model"

	"gorm.io/gorm"
)

type ProjectRepoEntity interface {
	CreateOne(project model.Project) (*model.Project, error)
	FindIdByEmail(email string) (*uint, error)
	GetProjectDetails(projectId uint) (*model.ProjectWithLikeViewComment, error)
	FindProjectById(id uint) (*model.Project, error)
	FindOwnerByProjectId(id uint) (*model.User, error)
	ShowListProjectWithPagination(pageNumber uint, pageSize uint, titleSearch string, sort string) ([]model.ProjectWithLikeViewComment, *uint, error)
	ShowListProjectWithPaginationByCategory(pageNumber uint, pageSize uint, category uint, titleSearch string, sort string) ([]model.ProjectWithLikeViewComment, *uint, error)
	DeleteProjectById(id uint) error
	UpdateProjectById(Updatedproject model.Project, id uint) (*model.Project, error)
}

type ProjectRepo struct {
	orm *gorm.DB
}

func NewProjectRepo(orm orm.GormOrmEntity) ProjectRepoEntity {
	return &ProjectRepo{orm: orm.GetOrm()}
}

func (p *ProjectRepo) CreateOne(project model.Project) (*model.Project, error) {
	tx := p.orm.Begin()
	tx = tx.Create(&project)
	if tx.Error != nil {
		tx.Rollback()
		return nil, tx.Error
	}
	tx.Commit()
	return &project, nil
}

func (p *ProjectRepo) FindIdByEmail(email string) (*uint, error) {
	var user model.User
	var id *uint
	tx := p.orm.Begin()
	err := tx.Select("id").Where("email = ?", email).Find(&user).Error
	id = &user.ID
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return id, nil
}

func (p *ProjectRepo) GetProjectDetails(projectId uint) (*model.ProjectWithLikeViewComment, error) {
	var project model.ProjectWithLikeViewComment
	tx := p.orm.Begin()
	subQueryLike := tx.Table("likes").Select("COUNT(likes.project_id) as total_likes").Joins("JOIN project p ON p.id = likes.project_id").Where("p.id = project.id")
	subQueryView := tx.Table("view").Select("COUNT(view.project_id) as total_views").Joins("JOIN project p ON p.id = view.project_id").Where("p.id = project.id")
	subQueryComment := tx.Table("comment").Select("COUNT(comment.project_id) as total_comments").Joins("JOIN project p ON p.id = comment.project_id").Where("p.id = project.id")
	err := tx.Model(model.Project{}).Select("*, (?), (?), (?)", subQueryLike, subQueryView, subQueryComment).
		Where("project.id = ?", projectId).Find(&project).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &project, nil
}

func (p *ProjectRepo) FindProjectById(id uint) (*model.Project, error) {
	var project model.Project
	tx := p.orm.Begin()
	err := tx.Model(model.Project{}).Where("id = ?", id).First(&project).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &project, nil
}

func (p *ProjectRepo) ShowListProjectWithPagination(pageNumber uint, pageSize uint, titleSearch string, sort string) ([]model.ProjectWithLikeViewComment, *uint, error) {
	var projects []model.ProjectWithLikeViewComment
	var count int64
	p.orm.Find(&model.Project{}).Where("project.title ILIKE ?", "%"+titleSearch+"%").Count(&count)

	totalOffset := (pageSize * pageNumber) - pageSize
	totalPagination := uint(count / int64(pageSize))
	if (count % int64(pageSize)) != 0 {
		totalPagination += 1
	}
	tx := p.orm.Begin()
	subQueryLike := tx.Table("likes").Select("COUNT(likes.project_id) as total_likes").Joins("JOIN project p ON p.id = likes.project_id").Where("p.id = project.id")
	subQueryView := tx.Table("view").Select("COUNT(view.project_id) as total_views").Joins("JOIN project p ON p.id = view.project_id").Where("p.id = project.id")
	subQueryComment := tx.Table("comment").Select("COUNT(comment.project_id) as total_comments").Joins("JOIN project p ON p.id = comment.project_id").Where("p.id = project.id")
	err := tx.Model(model.Project{}).Order("total_likes "+sort+", project.created_at DESC").Select("*, (?), (?), (?)", subQueryLike, subQueryView, subQueryComment).Where("project.title ILIKE ?", "%"+titleSearch+"%").
		Limit(int(pageSize)).Offset(int(totalOffset)).Find(&projects).Error
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}
	tx.Commit()
	return projects, &totalPagination, nil
}

func (p *ProjectRepo) ShowListProjectWithPaginationByCategory(pageNumber uint, pageSize uint, category uint, titleSearch string, sort string) ([]model.ProjectWithLikeViewComment, *uint, error) {
	var projects []model.ProjectWithLikeViewComment
	var count int64
	p.orm.Model(&model.Project{}).Where("category_id = ?", category).Count(&count)

	totalOffset := (pageSize * pageNumber) - pageSize
	totalPagination := uint(count / int64(pageSize))
	if (count % int64(pageSize)) != 0 {
		totalPagination += 1
	}
	tx := p.orm.Begin()
	subQueryLike := tx.Table("likes").Select("COUNT(likes.project_id) as total_likes").Joins("JOIN project p ON p.id = likes.project_id").Where("p.id = project.id")
	subQueryView := tx.Table("view").Select("COUNT(view.project_id) as total_views").Joins("JOIN project p ON p.id = view.project_id").Where("p.id = project.id")
	subQueryComment := tx.Table("comment").Select("COUNT(comment.project_id) as total_comments").Joins("JOIN project p ON p.id = comment.project_id").Where("p.id = project.id")
	err := tx.Model(model.Project{}).Order("total_likes "+sort+", project.created_at DESC").
		Select("*, (?), (?), (?)", subQueryLike, subQueryView, subQueryComment).
		Where("project.category_id = ? AND project.title ILIKE ?", category, "%"+titleSearch+"%").
		Limit(int(pageSize)).Offset(int(totalOffset)).Find(&projects).Error
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}
	tx.Commit()
	return projects, &totalPagination, nil
}

func (p *ProjectRepo) DeleteProjectById(id uint) error {
	var project model.Project
	tx := p.orm.Begin()
	err := tx.Delete(&project, id).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (p *ProjectRepo) UpdateProjectById(Updatedproject model.Project, id uint) (*model.Project, error) {
	var project model.Project
	tx := p.orm.Begin()
	tx = tx.Model(&project).Where("id = ?", id).Updates(Updatedproject)
	if tx.Error != nil {
		tx.Rollback()
		return nil, tx.Error
	}
	tx.Commit()
	return &project, nil
}

func (p *ProjectRepo) FindOwnerByProjectId(id uint) (*model.User, error) {
	var user model.User
	var project model.Project
	tx := p.orm.Begin()
	tx.First(&project, id)
	if tx.Error != nil {
		tx.Rollback()
		log.Println(tx.Error)
		return nil, tx.Error
	}
	userId := *project.UserID
	tx.Model(model.User{}).Where("id = ?", userId).First(&user)
	if tx.Error != nil {
		tx.Rollback()
		return nil, tx.Error
	}
	return &user, nil
}
