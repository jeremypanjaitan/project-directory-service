package usecase

import (
	"errors"
	"pds-backend/orm/gorm/model"
	repo "pds-backend/repo/categoryrepo"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

var nameCategory1 = "Name1"
var nameCategory2 = "Name2"

var dummyCategories = []model.Category{
	{
		Name: &nameCategory1,
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
	},
	{
		Name: &nameCategory2,
		Model: gorm.Model{
			ID:        2,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
	},
}

type categoryRepoMock struct {
	mock.Mock
}

func (r *categoryRepoMock) GetAll() ([]model.Category, error) {
	args := r.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Category), nil
}

func (r *categoryRepoMock) GetById(id uint) (*model.Category, error) {
	args := r.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Category), nil
}

type CategoryUsecaseTestSuite struct {
	suite.Suite
	repoTest repo.CategoryRepoEntity
}

func (suite *CategoryUsecaseTestSuite) SetupTest() {
	suite.repoTest = new(categoryRepoMock)
}

func (suite *CategoryUsecaseTestSuite) TestCategoryUsecase_GetAllCategory_Success() {
	suite.repoTest.(*categoryRepoMock).On("GetAll").Return(dummyCategories, nil)

	categoryUsecaseTest := NewCategoryUsecase(suite.repoTest)
	categories, err := categoryUsecaseTest.GetAllCategory()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummyCategories, categories)
}

func (suite *CategoryUsecaseTestSuite) TestCategoryUsecase_GetAllCategory_Failed() {
	suite.repoTest.(*categoryRepoMock).On("GetAll").Return(nil, errors.New("failed"))

	categoryUsecaseTest := NewCategoryUsecase(suite.repoTest)
	categories, err := categoryUsecaseTest.GetAllCategory()
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "failed", err.Error())
	assert.Nil(suite.T(), categories)
}

func (suite *CategoryUsecaseTestSuite) TestCategoryUsecase_GetCategoryByID_Success() {
	dummyCategory := dummyCategories[0]
	suite.repoTest.(*categoryRepoMock).On("GetById", dummyCategory.ID).Return(&dummyCategory, nil)

	categoryUsecaseTest := NewCategoryUsecase(suite.repoTest)
	category, err := categoryUsecaseTest.GetCategoryByID(dummyCategory.ID)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummyCategory.ID, category.ID)
}

func (suite *CategoryUsecaseTestSuite) TestCategoryUsecase_GetCategoryByID_Failed() {
	dummyCategory := dummyCategories[1]
	suite.repoTest.(*categoryRepoMock).On("GetById", dummyCategory.ID).Return(nil, errors.New("failed"))

	categoryUsecaseTest := NewCategoryUsecase(suite.repoTest)
	category, err := categoryUsecaseTest.GetCategoryByID(dummyCategory.ID)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "failed", err.Error())
	assert.Nil(suite.T(), category)
}

func TestCategoryUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(CategoryUsecaseTestSuite))
}
