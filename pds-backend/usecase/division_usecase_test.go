package usecase

import (
	"errors"
	"pds-backend/orm/gorm/model"
	repo "pds-backend/repo/divisionrepo"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

var nameDivision1 = "Name1"
var nameDivision2 = "Name2"

var dummyDivision = []model.Division{
	{
		Name: &nameDivision1,
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
	},
	{
		Name: &nameDivision2,
		Model: gorm.Model{
			ID:        2,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
	},
}

type divisionRepoMock struct {
	mock.Mock
}

func (r *divisionRepoMock) GetAll() ([]model.Division, error) {
	args := r.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Division), nil
}

func (r *divisionRepoMock) GetById(id uint) (*model.Division, error) {
	args := r.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Division), nil
}

type DivisionUsecaseTestSuite struct {
	suite.Suite
	repoTest repo.DivisionRepoEntity
}

func (suite *DivisionUsecaseTestSuite) SetupTest() {
	suite.repoTest = new(divisionRepoMock)
}

func (suite *DivisionUsecaseTestSuite) TestDivisionUsecase_GetAllDivision_Success() {
	suite.repoTest.(*divisionRepoMock).On("GetAll").Return(dummyDivision, nil)

	divisionUsecaseTest := NewDivisionUsecase(suite.repoTest)
	divisions, err := divisionUsecaseTest.GetAllDivision()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummyDivision, divisions)
}

func (suite *DivisionUsecaseTestSuite) TestDivisionUsecase_GetAllDivision_Failed() {
	suite.repoTest.(*divisionRepoMock).On("GetAll").Return(nil, errors.New("failed"))

	divisionUsecaseTest := NewDivisionUsecase(suite.repoTest)
	divisions, err := divisionUsecaseTest.GetAllDivision()
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "failed", err.Error())
	assert.Nil(suite.T(), divisions)
}

func (suite *DivisionUsecaseTestSuite) TestDivisionUsecase_GetDivisionByID_Success() {
	dummyDivision := dummyDivision[0]
	suite.repoTest.(*divisionRepoMock).On("GetById", dummyDivision.ID).Return(&dummyDivision, nil)

	divisionUsecaseTest := NewDivisionUsecase(suite.repoTest)
	division, err := divisionUsecaseTest.GetDivisionByID(dummyDivision.ID)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummyDivision.ID, division.ID)
}

func (suite *DivisionUsecaseTestSuite) TestDivisionUsecase_GetDivisionByID_Failed() {
	dummyDivision := dummyDivision[1]
	suite.repoTest.(*divisionRepoMock).On("GetById", dummyDivision.ID).Return(nil, errors.New("failed"))

	divisionUsecaseTest := NewDivisionUsecase(suite.repoTest)
	division, err := divisionUsecaseTest.GetDivisionByID(dummyDivision.ID)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "failed", err.Error())
	assert.Nil(suite.T(), division)
}

func TestDivisionUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(DivisionUsecaseTestSuite))
}
