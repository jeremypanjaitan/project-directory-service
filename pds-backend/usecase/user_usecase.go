package usecase

import (
	"pds-backend/orm/gorm/model"
	repo "pds-backend/repo/userrepo"
)

type UserUsecaseEntity interface {
	GetProfileUser(email string) (*model.User, *string, *string, error)
	GetProfileUserById(id uint) (*model.User, *string, *string, error)
	ChangePassword(email string, oldPassword string, newPassword string) error
	ChangeProfile(email string, updatedProfile model.User) (*model.User, *string, *string, error)
	GetActivity(pageNumber uint, pageSize uint, email string) ([]model.Activity, *uint, error)
	GetProject(pageNumber uint, pageSize uint, email string) ([]model.ProjectWithLikeViewComment, *uint, error)
}

type UserUsecase struct {
	userRepo repo.UserRepoEntity
}

func NewUserUsecase(userRepo repo.UserRepoEntity) UserUsecaseEntity {
	return &UserUsecase{userRepo: userRepo}
}

func (u *UserUsecase) GetProfileUser(email string) (*model.User, *string, *string, error) {
	profile, err := u.userRepo.FindOne(email)
	if err != nil {
		return nil, nil, nil, err
	}
	divisionName, err := u.userRepo.FindDivisionByDivisionId(*profile.DivisionID)
	if err != nil {
		return nil, nil, nil, err
	}
	roleName, err := u.userRepo.FindRoleByRoleId(*profile.RoleID)
	if err != nil {
		return nil, nil, nil, err
	}
	return profile, divisionName, roleName, nil
}

func (u *UserUsecase) GetProfileUserById(id uint) (*model.User, *string, *string, error) {
	profile, err := u.userRepo.FindOneById(id)
	if err != nil {
		return nil, nil, nil, err
	}
	divisionName, err := u.userRepo.FindDivisionByDivisionId(*profile.DivisionID)
	if err != nil {
		return nil, nil, nil, err
	}
	roleName, err := u.userRepo.FindRoleByRoleId(*profile.RoleID)
	if err != nil {
		return nil, nil, nil, err
	}
	return profile, divisionName, roleName, nil
}

func (u *UserUsecase) ChangePassword(email string, oldPassword string, newPassword string) error {
	return u.userRepo.UpdatePassword(email, oldPassword, newPassword)
}

func (u *UserUsecase) ChangeProfile(email string, updatedProfile model.User) (*model.User, *string, *string, error) {
	profile, err := u.userRepo.UpdateProfile(email, updatedProfile)
	if err != nil {
		return nil, nil, nil, err
	}
	divisionName, err := u.userRepo.FindDivisionByDivisionId(*profile.DivisionID)
	if err != nil {
		return nil, nil, nil, err
	}
	roleName, err := u.userRepo.FindRoleByRoleId(*profile.RoleID)
	if err != nil {
		return nil, nil, nil, err
	}
	return profile, divisionName, roleName, nil
}

func (u *UserUsecase) GetActivity(pageNumber uint, pageSize uint, email string) ([]model.Activity, *uint, error) {
	userId, err := u.userRepo.FindIdByEmail(email)
	if err != nil {
		return nil, nil, err
	}
	return u.userRepo.GetActivity(pageNumber, pageSize, *userId)
}

func (u *UserUsecase) GetProject(pageNumber uint, pageSize uint, email string) ([]model.ProjectWithLikeViewComment, *uint, error) {
	userId, err := u.userRepo.FindIdByEmail(email)
	if err != nil {
		return nil, nil, err
	}
	return u.userRepo.GetUserProject(pageNumber, pageSize, *userId)
}