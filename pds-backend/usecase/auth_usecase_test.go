package usecase

import (
	"errors"
	"pds-backend/apperror"
	"pds-backend/orm/gorm/model"
	repo "pds-backend/repo/userrepo"
	"pds-backend/service"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type tokenServiceMock struct {
	mock.Mock
}

func (t *tokenServiceMock) CreateAccessToken(credential *service.Credential) (*service.TokenDetails, error) {
	args := t.Called(credential)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*service.TokenDetails), nil
}

func (t *tokenServiceMock) VerifyAccessToken(tokenString string) (*service.UserCredential, error) {
	panic("Not Implemented")
}

func (t *tokenServiceMock) StoreAccessToken(email string, tokenDetails *service.TokenDetails) error {
	args := t.Called(email, tokenDetails)
	return args.Error(0)
}

func (t *tokenServiceMock) FetchAccessToken(userCredential *service.UserCredential) (string, error) {
	panic("Not Implemented")
}

func (t *tokenServiceMock) DeleteAccessToken(accessUuid string) error {
	args := t.Called(accessUuid)
	return args.Error(0)
}

type userRepoMock struct {
	mock.Mock
}

func (u *userRepoMock) CreateOne(user model.User) (*model.User, error) {
	args := u.Called(user)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), nil
}

func (u *userRepoMock) FindOne(email string) (*model.User, error) {
	args := u.Called(email)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), nil
}

func (u *userRepoMock) FindOneById(id uint) (*model.User, error) {
	panic("Not Implemented")
}

func (u *userRepoMock) UpdatePassword(email string, oldPassword string, newPassword string) error {
	panic("Not Implemented")
}

func (u *userRepoMock) UpdateProfile(email string, updatedProfile model.User) (*model.User, error) {
	panic("Not Implemented")
}

func (u *userRepoMock) FindRoleByRoleId(roleId uint) (*string, error) {
	args := u.Called(roleId)
	return args.Get(0).(*string), args.Error(1)
}

func (u *userRepoMock) FindDivisionByDivisionId(divisionId uint) (*string, error) {
	args := u.Called(divisionId)
	return args.Get(0).(*string), args.Error(1)
}

type AuthUsecaseTestSuite struct {
	suite.Suite
	userRepoMock     repo.UserRepoEntity
	tokenServiceMock service.TokenServiceEntity
}

func (suite *AuthUsecaseTestSuite) SetupTest() {
	suite.userRepoMock = new(userRepoMock)
	suite.tokenServiceMock = new(tokenServiceMock)
}

func (a *AuthUsecaseTestSuite) TestAuthUsecase_NewAuthUsecase() {
	NewAuthUsecase(a.tokenServiceMock, a.userRepoMock)
}

func (a *AuthUsecaseTestSuite) TestAuthUsecase_Logout() {
	mockAccessUuid := "4ff91736-9851-11ec-b909-0242ac120002"
	a.tokenServiceMock.(*tokenServiceMock).On("DeleteAccessToken", mockAccessUuid).Return(nil)
	authUsecaseTest := AuthUsecase{tokenService: a.tokenServiceMock, userRepo: a.userRepoMock}
	result := authUsecaseTest.Logout(mockAccessUuid)
	assert.Equal(a.T(), result, nil)
}

func (a *AuthUsecaseTestSuite) TestAuthUsecase_GetDivisionName() {
	divisionId, _ := strconv.ParseUint("20", 10, 32)
	divisionIdUint := uint(divisionId)
	mockDivisionName := "my division"
	a.userRepoMock.(*userRepoMock).On("FindDivisionByDivisionId", divisionIdUint).Return(&mockDivisionName, nil)
	authUsecaseTest := AuthUsecase{tokenService: a.tokenServiceMock, userRepo: a.userRepoMock}
	result, _ := authUsecaseTest.GetDivisionName(divisionIdUint)
	assert.Equal(a.T(), *result, mockDivisionName)
}

func (a *AuthUsecaseTestSuite) TestAuthUsecase_GetRoleName() {
	roleId, _ := strconv.ParseUint("20", 10, 32)
	roleIdUint := uint(roleId)
	mockRoleName := "my role"
	a.userRepoMock.(*userRepoMock).On("FindRoleByRoleId", roleIdUint).Return(&mockRoleName, nil)
	authUsecaseTest := AuthUsecase{tokenService: a.tokenServiceMock, userRepo: a.userRepoMock}
	result, _ := authUsecaseTest.GetRoleName(roleIdUint)
	assert.Equal(a.T(), *result, mockRoleName)
}

func (a *AuthUsecaseTestSuite) TestAuthUsecase_Register_Success() {
	var userDummy model.User
	roleId, _ := strconv.ParseUint("20", 10, 32)
	roleIdUint := uint(roleId)
	divisionId, _ := strconv.ParseUint("20", 10, 32)
	divisionIdUint := uint(divisionId)
	userFullName := "Jeremy Panjaitan"
	userEmail := "jeremypanjaitan@gmail.com"
	picture := ""
	gender := "M"
	DivisionID := divisionIdUint
	RoleID := roleIdUint
	Bio := "this is my bio"
	Password := "jeremy password"

	userDummy.FullName = &userFullName
	userDummy.Email = &userEmail
	userDummy.Picture = &picture
	userDummy.Gender = &gender
	userDummy.DivisionID = &DivisionID
	userDummy.RoleID = &RoleID
	userDummy.Biography = &Bio
	userDummy.Password = &Password

	a.userRepoMock.(*userRepoMock).On("CreateOne", userDummy).Return(&userDummy, nil)
	authUsecaseTest := AuthUsecase{tokenService: a.tokenServiceMock, userRepo: a.userRepoMock}
	result, _ := authUsecaseTest.Register(userDummy)
	assert.Equal(a.T(), result.Email, userDummy.Email)

}

func (a *AuthUsecaseTestSuite) TestAuthUsecase_Register_Failed() {
	var userDummy model.User
	roleId, _ := strconv.ParseUint("20", 10, 32)
	roleIdUint := uint(roleId)
	divisionId, _ := strconv.ParseUint("20", 10, 32)
	divisionIdUint := uint(divisionId)
	userFullName := "Jeremy Panjaitan"
	userEmail := "jeremypanjaitan@gmail.com"
	picture := ""
	gender := "M"
	DivisionID := divisionIdUint
	RoleID := roleIdUint
	Bio := "this is my bio"
	Password := "jeremy password"

	userDummy.FullName = &userFullName
	userDummy.Email = &userEmail
	userDummy.Picture = &picture
	userDummy.Gender = &gender
	userDummy.DivisionID = &DivisionID
	userDummy.RoleID = &RoleID
	userDummy.Biography = &Bio
	userDummy.Password = &Password

	mockRegisterError := errors.New("errors register")
	a.userRepoMock.(*userRepoMock).On("CreateOne", userDummy).Return(nil, mockRegisterError)
	authUsecaseTest := AuthUsecase{tokenService: a.tokenServiceMock, userRepo: a.userRepoMock}
	_, errActual := authUsecaseTest.Register(userDummy)
	assert.Equal(a.T(), errActual, mockRegisterError)

}

func (a *AuthUsecaseTestSuite) TestAuthUsecase_Login_Success() {
	var credentialDummy service.Credential
	var userDummy model.User
	var tokenDetailsDummy service.TokenDetails

	credentialDummy.Email = "email@dummy.com"
	credentialDummy.Password = "dummypassword"
	roleId, _ := strconv.ParseUint("20", 10, 32)
	roleIdUint := uint(roleId)
	divisionId, _ := strconv.ParseUint("20", 10, 32)
	divisionIdUint := uint(divisionId)
	userFullName := "Jeremy Panjaitan"
	userEmail := "email@dummy.com"
	picture := ""
	gender := "M"
	DivisionID := divisionIdUint
	RoleID := roleIdUint
	Bio := "this is my bio"
	Password := "dummypassword"

	userDummy.FullName = &userFullName
	userDummy.Email = &userEmail
	userDummy.Picture = &picture
	userDummy.Gender = &gender
	userDummy.DivisionID = &DivisionID
	userDummy.RoleID = &RoleID
	userDummy.Biography = &Bio
	userDummy.Password = &Password

	tokenDetailsDummy.AccessToken = "access token"
	tokenDetailsDummy.AccessUuid = "access uuid"
	tokenDetailsDummy.AtExpires = time.Now().Unix()

	a.userRepoMock.(*userRepoMock).On("FindOne", credentialDummy.Email).Return(&userDummy, nil)
	a.tokenServiceMock.(*tokenServiceMock).On("CreateAccessToken", &credentialDummy).Return(&tokenDetailsDummy, nil)
	a.tokenServiceMock.(*tokenServiceMock).On("StoreAccessToken", credentialDummy.Email, &tokenDetailsDummy).Return(nil)

	authUsecaseTest := AuthUsecase{tokenService: a.tokenServiceMock, userRepo: a.userRepoMock}
	userActual, tokenDetailsActual, err := authUsecaseTest.Login(credentialDummy)
	assert.Equal(a.T(), err, nil)
	assert.Equal(a.T(), *userActual, userDummy)
	assert.Equal(a.T(), *tokenDetailsActual, tokenDetailsDummy)

}
func (a *AuthUsecaseTestSuite) TestAuthUsecase_Login_Error_FindEmail() {
	var credentialDummy service.Credential
	var userDummy model.User
	var tokenDetailsDummy service.TokenDetails

	credentialDummy.Email = "email@dummy.com"
	credentialDummy.Password = "dummypassword"
	roleId, _ := strconv.ParseUint("20", 10, 32)
	roleIdUint := uint(roleId)
	divisionId, _ := strconv.ParseUint("20", 10, 32)
	divisionIdUint := uint(divisionId)
	userFullName := "Jeremy Panjaitan"
	userEmail := "email@dummy.com"
	picture := ""
	gender := "M"
	DivisionID := divisionIdUint
	RoleID := roleIdUint
	Bio := "this is my bio"
	Password := "dummypassword"

	userDummy.FullName = &userFullName
	userDummy.Email = &userEmail
	userDummy.Picture = &picture
	userDummy.Gender = &gender
	userDummy.DivisionID = &DivisionID
	userDummy.RoleID = &RoleID
	userDummy.Biography = &Bio
	userDummy.Password = &Password

	tokenDetailsDummy.AccessToken = "access token"
	tokenDetailsDummy.AccessUuid = "access uuid"
	tokenDetailsDummy.AtExpires = time.Now().Unix()

	mockFindOneError := errors.New("error find one")
	a.userRepoMock.(*userRepoMock).On("FindOne", credentialDummy.Email).Return(nil, mockFindOneError)
	a.tokenServiceMock.(*tokenServiceMock).On("CreateAccessToken", &credentialDummy).Return(&tokenDetailsDummy, nil)
	a.tokenServiceMock.(*tokenServiceMock).On("StoreAccessToken", credentialDummy.Email, &tokenDetailsDummy).Return(nil)

	authUsecaseTest := AuthUsecase{tokenService: a.tokenServiceMock, userRepo: a.userRepoMock}
	userActual, tokenDetailsActual, errActual := authUsecaseTest.Login(credentialDummy)

	assert.Equal(a.T(), mockFindOneError, errActual)
	assert.Nil(a.T(), userActual)
	assert.Nil(a.T(), tokenDetailsActual)

}

func (a *AuthUsecaseTestSuite) TestAuthUsecase_Login_Error_WrongCredential() {
	var credentialDummy service.Credential
	var userDummy model.User
	var tokenDetailsDummy service.TokenDetails

	credentialDummy.Email = "email@dummy.com"
	credentialDummy.Password = "dummypassword"
	roleId, _ := strconv.ParseUint("20", 10, 32)
	roleIdUint := uint(roleId)
	divisionId, _ := strconv.ParseUint("20", 10, 32)
	divisionIdUint := uint(divisionId)
	userFullName := "Jeremy Panjaitan"
	userEmail := "email@dummy"
	picture := ""
	gender := "M"
	DivisionID := divisionIdUint
	RoleID := roleIdUint
	Bio := "this is my bio"
	Password := "dummypassword"

	userDummy.FullName = &userFullName
	userDummy.Email = &userEmail
	userDummy.Picture = &picture
	userDummy.Gender = &gender
	userDummy.DivisionID = &DivisionID
	userDummy.RoleID = &RoleID
	userDummy.Biography = &Bio
	userDummy.Password = &Password

	tokenDetailsDummy.AccessToken = "access token"
	tokenDetailsDummy.AccessUuid = "access uuid"
	tokenDetailsDummy.AtExpires = time.Now().Unix()

	a.userRepoMock.(*userRepoMock).On("FindOne", credentialDummy.Email).Return(&userDummy, nil)
	a.tokenServiceMock.(*tokenServiceMock).On("CreateAccessToken", &credentialDummy).Return(&tokenDetailsDummy, nil)
	a.tokenServiceMock.(*tokenServiceMock).On("StoreAccessToken", credentialDummy.Email, &tokenDetailsDummy).Return(nil)

	authUsecaseTest := AuthUsecase{tokenService: a.tokenServiceMock, userRepo: a.userRepoMock}
	userActual, tokenDetailsActual, errActual := authUsecaseTest.Login(credentialDummy)

	assert.Equal(a.T(), apperror.ErrCredential, errActual)
	assert.Nil(a.T(), userActual)
	assert.Nil(a.T(), tokenDetailsActual)

}
func (a *AuthUsecaseTestSuite) TestAuthUsecase_Login_Error_CreateAccessToken() {
	var credentialDummy service.Credential
	var userDummy model.User
	var tokenDetailsDummy service.TokenDetails

	credentialDummy.Email = "email@dummy.com"
	credentialDummy.Password = "dummypassword"
	roleId, _ := strconv.ParseUint("20", 10, 32)
	roleIdUint := uint(roleId)
	divisionId, _ := strconv.ParseUint("20", 10, 32)
	divisionIdUint := uint(divisionId)
	userFullName := "Jeremy Panjaitan"
	userEmail := "email@dummy.com"
	picture := ""
	gender := "M"
	DivisionID := divisionIdUint
	RoleID := roleIdUint
	Bio := "this is my bio"
	Password := "dummypassword"

	userDummy.FullName = &userFullName
	userDummy.Email = &userEmail
	userDummy.Picture = &picture
	userDummy.Gender = &gender
	userDummy.DivisionID = &DivisionID
	userDummy.RoleID = &RoleID
	userDummy.Biography = &Bio
	userDummy.Password = &Password

	tokenDetailsDummy.AccessToken = "access token"
	tokenDetailsDummy.AccessUuid = "access uuid"
	tokenDetailsDummy.AtExpires = time.Now().Unix()

	mockErrorCreateAccessToken := errors.New("errors create access token")
	a.userRepoMock.(*userRepoMock).On("FindOne", credentialDummy.Email).Return(&userDummy, nil)
	a.tokenServiceMock.(*tokenServiceMock).On("CreateAccessToken", &credentialDummy).Return(nil, mockErrorCreateAccessToken)
	a.tokenServiceMock.(*tokenServiceMock).On("StoreAccessToken", credentialDummy.Email, &tokenDetailsDummy).Return(nil)

	authUsecaseTest := AuthUsecase{tokenService: a.tokenServiceMock, userRepo: a.userRepoMock}
	userActual, tokenDetailsActual, errActual := authUsecaseTest.Login(credentialDummy)

	assert.Equal(a.T(), mockErrorCreateAccessToken, errActual)
	assert.Nil(a.T(), userActual)
	assert.Nil(a.T(), tokenDetailsActual)

}
func (a *AuthUsecaseTestSuite) TestAuthUsecase_Login_Error_StoreAccessToken() {
	var credentialDummy service.Credential
	var userDummy model.User
	var tokenDetailsDummy service.TokenDetails

	credentialDummy.Email = "email@dummy.com"
	credentialDummy.Password = "dummypassword"
	roleId, _ := strconv.ParseUint("20", 10, 32)
	roleIdUint := uint(roleId)
	divisionId, _ := strconv.ParseUint("20", 10, 32)
	divisionIdUint := uint(divisionId)
	userFullName := "Jeremy Panjaitan"
	userEmail := "email@dummy.com"
	picture := ""
	gender := "M"
	DivisionID := divisionIdUint
	RoleID := roleIdUint
	Bio := "this is my bio"
	Password := "dummypassword"

	userDummy.FullName = &userFullName
	userDummy.Email = &userEmail
	userDummy.Picture = &picture
	userDummy.Gender = &gender
	userDummy.DivisionID = &DivisionID
	userDummy.RoleID = &RoleID
	userDummy.Biography = &Bio
	userDummy.Password = &Password

	tokenDetailsDummy.AccessToken = "access token"
	tokenDetailsDummy.AccessUuid = "access uuid"
	tokenDetailsDummy.AtExpires = time.Now().Unix()

	mockErrorStoreAccessToken := errors.New("errors store access token")
	a.userRepoMock.(*userRepoMock).On("FindOne", credentialDummy.Email).Return(&userDummy, nil)
	a.tokenServiceMock.(*tokenServiceMock).On("CreateAccessToken", &credentialDummy).Return(&tokenDetailsDummy, nil)
	a.tokenServiceMock.(*tokenServiceMock).On("StoreAccessToken", credentialDummy.Email, &tokenDetailsDummy).Return(mockErrorStoreAccessToken)

	authUsecaseTest := AuthUsecase{tokenService: a.tokenServiceMock, userRepo: a.userRepoMock}
	userActual, tokenDetailsActual, errActual := authUsecaseTest.Login(credentialDummy)

	assert.Equal(a.T(), mockErrorStoreAccessToken, errActual)
	assert.Nil(a.T(), userActual)
	assert.Nil(a.T(), tokenDetailsActual)

}

func TestAuthUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(AuthUsecaseTestSuite))
}
