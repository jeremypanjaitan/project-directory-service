package usecase

import (
	"errors"
	"pds-backend/orm/gorm/model"
	repo "pds-backend/repo/rolerepo"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

var nameRole1 = "Name1"
var nameRole2 = "Name2"

var dummyRoles = []model.Role{
	{
		Name: &nameRole1,
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
	},
	{
		Name: &nameRole2,
		Model: gorm.Model{
			ID:        2,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
	},
}

type roleRepoMock struct {
	mock.Mock
}

func (r *roleRepoMock) GetAll() ([]model.Role, error) {
	args := r.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Role), nil
}

func (r *roleRepoMock) GetById(id uint) (*model.Role, error) {
	args := r.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Role), nil
}

type RoleUsecaseTestSuite struct {
	suite.Suite
	repoTest repo.RoleRepoEntity
}

func (suite *RoleUsecaseTestSuite) SetupTest() {
	suite.repoTest = new(roleRepoMock)
}

func (suite *RoleUsecaseTestSuite) TestRoleUsecase_GetAllRole_Success() {
	suite.repoTest.(*roleRepoMock).On("GetAll").Return(dummyRoles, nil)

	roleUsecaseTest := NewRoleUsecase(suite.repoTest)
	roles, err := roleUsecaseTest.GetAllRole()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummyRoles, roles)
}

func (suite *RoleUsecaseTestSuite) TestRoleUsecase_GetAllRole_Failed() {
	suite.repoTest.(*roleRepoMock).On("GetAll").Return(nil, errors.New("failed"))

	rolesUsecaseTest := NewRoleUsecase(suite.repoTest)
	roles, err := rolesUsecaseTest.GetAllRole()
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "failed", err.Error())
	assert.Nil(suite.T(), roles)
}

func (suite *RoleUsecaseTestSuite) TestRoleUsecase_GetRoleByID_Success() {
	dummyRole := dummyRoles[0]
	suite.repoTest.(*roleRepoMock).On("GetById", dummyRole.ID).Return(&dummyRole, nil)

	roleUsecaseTest := NewRoleUsecase(suite.repoTest)
	role, err := roleUsecaseTest.GetRoleByID(dummyRole.ID)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummyRole.ID, role.ID)
}

func (suite *RoleUsecaseTestSuite) TestRoleUsecase_GetRoleByID_Failed() {
	dummyRole := dummyRoles[1]
	suite.repoTest.(*roleRepoMock).On("GetById", dummyRole.ID).Return(nil, errors.New("failed"))

	roleUsecaseTest := NewRoleUsecase(suite.repoTest)
	role, err := roleUsecaseTest.GetRoleByID(dummyRole.ID)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "failed", err.Error())
	assert.Nil(suite.T(), role)
}

func TestRoleUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(RoleUsecaseTestSuite))
}
