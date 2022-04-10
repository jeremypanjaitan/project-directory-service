package repo

import (
	orm "pds-backend/orm/gorm"
	"pds-backend/orm/gorm/model"

	"gorm.io/gorm"
)

type UserRepoEntity interface {
	CreateOne(user model.User) (*model.User, error)
	FindOne(email string) (*model.User, error)
	FindOneById(id uint) (*model.User, error)
	FindOneByEmail(email string) (*model.User, error)
	UpdatePassword(email string, oldPassword string, newPassword string) error
	UpdateProfile(email string, updatedProfile model.User) (*model.User, error)
	FindRoleByRoleId(roleId uint) (*string, error)
	FindDivisionByDivisionId(divisionId uint) (*string, error)
	FindIdByEmail(email string) (*uint, error)
	GetActivity(pageNumber uint, pageSize uint, userId uint) ([]model.Activity, *uint, error)
	GetUserProject(pageNumber uint, pageSize uint, userId uint) ([]model.ProjectWithLikeViewComment, *uint, error)
}

type UserRepo struct {
	orm *gorm.DB
}

func NewUserRepo(orm orm.GormOrmEntity) UserRepoEntity {
	return &UserRepo{orm: orm.GetOrm()}
}

func (u *UserRepo) CreateOne(user model.User) (*model.User, error) {
	tx := u.orm.Begin()
	tx = tx.Create(&user)
	if tx.Error != nil {
		tx.Rollback()
		return nil, tx.Error
	}
	tx.Commit()
	return &user, nil
}

func (u *UserRepo) FindOne(email string) (*model.User, error) {
	var user model.User
	tx := u.orm.Begin()
	err := tx.Model(model.User{}).Where("email = ?", email).First(&user).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &user, nil
}

func (u *UserRepo) FindOneById(id uint) (*model.User, error) {
	var user model.User
	tx := u.orm.Begin()
	err := tx.Model(model.User{}).Where("id = ?", id).First(&user).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &user, nil
}

func (u *UserRepo) FindOneByEmail(email string) (*model.User, error) {
	var user model.User
	tx := u.orm.Begin()
	err := tx.Model(model.User{}).Where("email = ?", email).First(&user).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &user, nil
}

func (u *UserRepo) UpdatePassword(email string, oldPassword string, newPassword string) error {
	var user model.User
	err := u.orm.Where("email = ? AND password = ?", email, oldPassword).First(&user).Error
	if err != nil {
		return err
	}
	tx := u.orm.Begin()
	err = tx.Model(&user).Where("email = ? AND password = ?", email, oldPassword).Update("password", newPassword).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (u *UserRepo) UpdateProfile(email string, updatedProfile model.User) (*model.User, error) {
	var user model.User
	tx := u.orm.Begin()
	err := tx.Model(&user).Where("email = ?", email).Updates(updatedProfile).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &user, nil
}

func (u *UserRepo) FindRoleByRoleId(roleId uint) (*string, error) {
	var role model.Role
	tx := u.orm.Begin()
	err := tx.Select("name").Where("id = ?", roleId).Find(&role).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return role.Name, nil
}

func (u *UserRepo) FindDivisionByDivisionId(divisionId uint) (*string, error) {
	var division model.Division
	tx := u.orm.Begin()
	err := tx.Select("name").Where("id = ?", divisionId).Find(&division).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return division.Name, nil
}

func (u *UserRepo) FindIdByEmail(email string) (*uint, error) {
	var user model.User
	var id *uint
	tx := u.orm.Begin()
	err := tx.Select("id").Where("email = ?", email).Find(&user).Error
	id = &user.ID
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return id, nil
}

func (u *UserRepo) GetActivity(pageNumber uint, pageSize uint, userId uint) ([]model.Activity, *uint, error) {
	var activites []model.Activity
	var count int64
	u.orm.Model(&model.Activity{}).Where("user_id = ?", userId).Count(&count)

	totalOffset := (pageSize * pageNumber) - pageSize
	totalPagination := uint(count / int64(pageSize))
	if (count % int64(pageSize)) != 0 {
		totalPagination += 1
	}
	tx := u.orm.Begin()
	err := tx.Order("id DESC").Model(model.Activity{}).Where("user_id = ?", userId).
		Limit(int(pageSize)).Offset(int(totalOffset)).Find(&activites).Error
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}
	tx.Commit()
	return activites, &totalPagination, nil
}

func (u *UserRepo) GetUserProject(pageNumber uint, pageSize uint, userId uint) ([]model.ProjectWithLikeViewComment, *uint, error) {
	var projects []model.ProjectWithLikeViewComment
	var count int64
	u.orm.Model(&model.Project{}).Where("user_id = ?", userId).Count(&count)

	totalOffset := (pageSize * pageNumber) - pageSize
	totalPagination := uint(count / int64(pageSize))
	if (count % int64(pageSize)) != 0 {
		totalPagination += 1
	}
	tx := u.orm.Begin()
	subQueryLike := tx.Table("likes").Select("COUNT(likes.project_id) as total_likes").Joins("JOIN project p ON p.id = likes.project_id").Where("p.id = project.id")
	subQueryView := tx.Table("view").Select("COUNT(view.project_id) as total_views").Joins("JOIN project p ON p.id = view.project_id").Where("p.id = project.id")
	subQueryComment := tx.Table("comment").Select("COUNT(comment.project_id) as total_comments").Joins("JOIN project p ON p.id = comment.project_id").Where("p.id = project.id")
	err := tx.Model(model.Project{}).Order("project.id DESC").Select("*, (?), (?), (?)", subQueryLike, subQueryView, subQueryComment).Where("project.user_id = ?", userId).
		Limit(int(pageSize)).Offset(int(totalOffset)).Find(&projects).Error
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}
	tx.Commit()
	return projects, &totalPagination, nil
}
