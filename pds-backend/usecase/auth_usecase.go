package usecase

import (
	"fmt"
	"pds-backend/apperror"
	"pds-backend/delivery/httpdelivery/jsonmodels"
	"pds-backend/orm/gorm/model"
	repo "pds-backend/repo/userrepo"
	"pds-backend/service"
	"pds-backend/utils"
)

type AuthUsecaseEntity interface {
	Login(credential service.Credential) (*model.User, *service.TokenDetails, error)
	Logout(accessUuid string) error
	Register(user model.User) (*jsonmodels.RegisterReponseBody, error)
	GetDivisionName(divisionId uint) (*string, error)
	GetRoleName(roleId uint) (*string, error)
}

type AuthUsecase struct {
	tokenService service.TokenServiceEntity
	userRepo     repo.UserRepoEntity
}

func NewAuthUsecase(
	tokenService service.TokenServiceEntity,
	userRepo repo.UserRepoEntity,
) AuthUsecaseEntity {
	return &AuthUsecase{
		tokenService: tokenService,
		userRepo:     userRepo,
	}
}

func (a *AuthUsecase) Login(credential service.Credential) (*model.User, *service.TokenDetails, error) {
	userCredential, err := a.userRepo.FindOne(credential.Email)
	if err != nil {
		return nil, nil, err
	}
	if !(*userCredential.Email == credential.Email) {
		return nil, nil, apperror.ErrCredential
	}

	tokenDetails, err := a.tokenService.CreateAccessToken(&credential)
	if err != nil {
		return nil, nil, err
	}
	err = a.tokenService.StoreAccessToken(credential.Email, tokenDetails)
	if err != nil {
		return nil, nil, err
	}
	return userCredential, tokenDetails, err
}

func (a *AuthUsecase) Logout(accessUuid string) error {
	return a.tokenService.DeleteAccessToken(accessUuid)
}

func (a *AuthUsecase) Register(user model.User) (*jsonmodels.RegisterReponseBody, error) {
	var registerResponseBody jsonmodels.RegisterReponseBody
	stringHashPassword := fmt.Sprint(utils.Hash(*user.Password))
	user.Password = &stringHashPassword
	createdUser, err := a.userRepo.CreateOne(user)
	if err != nil {
		return nil, err
	}
	registerResponseBody.ID = &createdUser.ID
	registerResponseBody.FullName = createdUser.FullName
	registerResponseBody.Email = createdUser.Email
	registerResponseBody.Gender = createdUser.Gender
	registerResponseBody.DivisionID = createdUser.DivisionID
	registerResponseBody.RoleID = createdUser.RoleID
	return &registerResponseBody, nil
}

func (a *AuthUsecase) GetDivisionName(divisionId uint) (*string, error) {
	return a.userRepo.FindDivisionByDivisionId(divisionId)
}

func (a *AuthUsecase) GetRoleName(roleId uint) (*string, error) {
	return a.userRepo.FindRoleByRoleId(roleId)
}
