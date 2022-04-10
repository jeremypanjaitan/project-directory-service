package usecase

import (
	"errors"
	"pds-backend/orm/gorm/model"
	repoUser "pds-backend/repo/userrepo"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

var name = "fahmi uhuy"
var name2 = "diva ihiy"
var email = "fahmi@gmail.com"
var email2 = "diva@gmail.com"
var password = "test123"
var password2 = "asd123f"
var gender = "M"
var gender2 = "F"
var bio = "ini bio loh"
var bio2 = "itu bio loh"
var div = uint(2)
var div2 = uint(3)
var role = uint(1)
var role2 = uint(2)
var dummyUsers = []model.User{
	{
		FullName:   &name,
		Email:      &email,
		Password:   &password,
		Picture:    nil,
		Gender:     &gender,
		Biography:  &bio,
		DivisionID: &div,
		RoleID:     &role,
	},
	{
		FullName:   &name2,
		Email:      &email2,
		Password:   &password2,
		Gender:     &gender2,
		Biography:  &bio2,
		DivisionID: &div2,
		RoleID:     &role2,
		Model:      gorm.Model{ID: 2},
	},
}

var upName = "ganti nama"
var upEmail = "ganti email"
var updatedDummy = model.User{
	FullName: &upName,
	Email:    &upEmail,
}

var passwordDummy = "passBaruw"

type repoUserMock struct {
	mock.Mock
}

type UserUsecaseTestSuite struct {
	suite.Suite
	repoTest repoUser.UserRepoEntity
}

func (suite *UserUsecaseTestSuite) SetupTest() {
	suite.repoTest = new(repoUserMock)
}

func TestUserUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))
}

func (r *repoUserMock) CreateOne(user model.User) (*model.User, error) {
	args := r.Called(user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), nil
}

func (r *repoUserMock) FindOne(email string) (*model.User, error) {
	args := r.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), nil
}

func (r *repoUserMock) FindOneById(id uint) (*model.User, error) {
	args := r.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), nil
}

func (r *repoUserMock) UpdatePassword(email string, oldPassword string, newPassword string) error {
	args := r.Called(email, oldPassword, newPassword)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (r *repoUserMock) UpdateProfile(email string, updatedProfile model.User) (*model.User, error) {
	args := r.Called(email, updatedProfile)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), nil
}

func (r *repoUserMock) FindRoleByRoleId(roleId uint) (*string, error) {
	args := r.Called(roleId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*string), nil
}

func (r *repoUserMock) FindDivisionByDivisionId(divisionId uint) (*string, error) {
	args := r.Called(divisionId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*string), nil
}

func (suite *UserUsecaseTestSuite) TestUserUsecase_GetProfileUser_Success() {
	dummyUser := &dummyUsers[0]
	dummyEmail := dummyUser.Email
	dummyDivisionId := dummyUser.DivisionID
	dummyRoleId := dummyUser.RoleID
	dummyResult := "test"
	suite.repoTest.(*repoUserMock).On("FindOne", *dummyEmail).Return(dummyUser, nil)
	suite.repoTest.(*repoUserMock).On("FindDivisionByDivisionId", *dummyDivisionId).Return(&dummyResult, nil)
	suite.repoTest.(*repoUserMock).On("FindRoleByRoleId", *dummyRoleId).Return(&dummyResult, nil)
	userUsecaseTest := NewUserUsecase(suite.repoTest)
	user, divisionName, roleName, err := userUsecaseTest.GetProfileUser(*dummyEmail)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), &dummyUser.FullName, &user.FullName)
	assert.Equal(suite.T(), "test", *divisionName)
	assert.Equal(suite.T(), "test", *roleName)
}

func (suite *UserUsecaseTestSuite) TestUserUsecase_GetProfileUser_Failed() {
	dummyUser := &dummyUsers[0]
	dummyEmail := dummyUser.Email
	dummyDivisionId := dummyUser.DivisionID
	dummyRoleId := dummyUser.RoleID
	suite.repoTest.(*repoUserMock).On("FindOne", *dummyEmail).Return(nil, errors.New("failed"))
	suite.repoTest.(*repoUserMock).On("FindDivisionByDivisionId", *dummyDivisionId).Return(nil, errors.New("failed"))
	suite.repoTest.(*repoUserMock).On("FindRoleByRoleId", *dummyRoleId).Return(nil, errors.New("failed"))
	userUsecaseTest := NewUserUsecase(suite.repoTest)
	user, divisionName, roleName, err := userUsecaseTest.GetProfileUser(*dummyEmail)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "failed", err.Error())
	assert.Nil(suite.T(), user)
	assert.Nil(suite.T(), divisionName)
	assert.Nil(suite.T(), roleName)
}

func (suite *UserUsecaseTestSuite) TestUserUsecase_ChangeProfile_Success() {
	dummyUser := &dummyUsers[1]
	dummyEmail := dummyUser.Email
	dummyDivisionId := dummyUser.DivisionID
	dummyRoleId := dummyUser.RoleID
	dummyResult := "test"
	suite.repoTest.(*repoUserMock).On("UpdateProfile", *dummyEmail, updatedDummy).Return(dummyUser, nil)
	suite.repoTest.(*repoUserMock).On("FindDivisionByDivisionId", *dummyDivisionId).Return(&dummyResult, nil)
	suite.repoTest.(*repoUserMock).On("FindRoleByRoleId", *dummyRoleId).Return(&dummyResult, nil)
	userUsecaseTest := NewUserUsecase(suite.repoTest)
	user, divisionName, roleName, err := userUsecaseTest.ChangeProfile(*dummyEmail, updatedDummy)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), &dummyUser.FullName, &user.FullName)
	assert.Equal(suite.T(), "test", *divisionName)
	assert.Equal(suite.T(), "test", *roleName)
}

func (suite *UserUsecaseTestSuite) TestUserUsecase_ChangeProfile_Failed() {
	dummyUser := &dummyUsers[1]
	dummyEmail := dummyUser.Email
	dummyDivisionId := dummyUser.DivisionID
	dummyRoleId := dummyUser.RoleID
	suite.repoTest.(*repoUserMock).On("UpdateProfile", *dummyEmail, updatedDummy).Return(nil, errors.New("failed"))
	suite.repoTest.(*repoUserMock).On("FindDivisionByDivisionId", *dummyDivisionId).Return(nil, errors.New("failed"))
	suite.repoTest.(*repoUserMock).On("FindRoleByRoleId", *dummyRoleId).Return(nil, errors.New("failed"))
	userUsecaseTest := NewUserUsecase(suite.repoTest)
	user, divisionName, roleName, err := userUsecaseTest.ChangeProfile(*dummyEmail, updatedDummy)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "failed", err.Error())
	assert.Nil(suite.T(), user)
	assert.Nil(suite.T(), divisionName)
	assert.Nil(suite.T(), roleName)
}

func (suite *UserUsecaseTestSuite) TestUserUsecase_ChangePassword_Success() {
	dummyUser := &dummyUsers[0]
	dummyEmail := dummyUser.Email
	dummyOldPass := dummyUser.Password
	suite.repoTest.(*repoUserMock).On("UpdatePassword", *dummyEmail, *dummyOldPass, passwordDummy).Return(nil)
	userUsecaseTest := NewUserUsecase(suite.repoTest)
	err := userUsecaseTest.ChangePassword(*dummyEmail, *dummyOldPass, passwordDummy)
	assert.Nil(suite.T(), err)
}

func (suite *UserUsecaseTestSuite) TestUserUsecase_ChangePassword_Failed() {
	dummyUser := &dummyUsers[0]
	dummyEmail := dummyUser.Email
	dummyOldPass := dummyUsers[1].Password
	suite.repoTest.(*repoUserMock).On("UpdatePassword", *dummyEmail, *dummyOldPass, passwordDummy).Return(errors.New("failed"))
	userUsecaseTest := NewUserUsecase(suite.repoTest)
	err := userUsecaseTest.ChangePassword(*dummyEmail, *dummyOldPass, passwordDummy)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "failed", err.Error())
}

func (suite *UserUsecaseTestSuite) TestUserUsecase_GetProfileUserById_Success() {
	dummyUser := &dummyUsers[0]
	dummyId := dummyUser.Model.ID
	dummyDivisionId := dummyUser.DivisionID
	dummyRoleId := dummyUser.RoleID
	dummyResult := "test"
	suite.repoTest.(*repoUserMock).On("FindOneById", dummyId).Return(dummyUser, nil)
	suite.repoTest.(*repoUserMock).On("FindDivisionByDivisionId", *dummyDivisionId).Return(&dummyResult, nil)
	suite.repoTest.(*repoUserMock).On("FindRoleByRoleId", *dummyRoleId).Return(&dummyResult, nil)
	userUsecaseTest := NewUserUsecase(suite.repoTest)
	user, divisionName, roleName, err := userUsecaseTest.GetProfileUserById(dummyId)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), &dummyUser.FullName, &user.FullName)
	assert.Equal(suite.T(), "test", *divisionName)
	assert.Equal(suite.T(), "test", *roleName)
}

func (suite *UserUsecaseTestSuite) TestUserUsecase_GetProfileUserById_Failed() {
	dummyUser := &dummyUsers[0]
	dummyId := dummyUser.Model.ID
	dummyDivisionId := dummyUser.DivisionID
	dummyRoleId := dummyUser.RoleID
	suite.repoTest.(*repoUserMock).On("FindOneById", dummyId).Return(nil, errors.New("failed"))
	suite.repoTest.(*repoUserMock).On("FindDivisionByDivisionId", *dummyDivisionId).Return(nil, errors.New("failed"))
	suite.repoTest.(*repoUserMock).On("FindRoleByRoleId", *dummyRoleId).Return(nil, errors.New("failed"))
	userUsecaseTest := NewUserUsecase(suite.repoTest)
	user, divisionName, roleName, err := userUsecaseTest.GetProfileUserById(dummyId)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "failed", err.Error())
	assert.Nil(suite.T(), user)
	assert.Nil(suite.T(), divisionName)
	assert.Nil(suite.T(), roleName)
}
