package routes

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"pds-backend/orm/gorm/model"
	"pds-backend/usecase"
	"time"

	"github.com/gin-gonic/gin"
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

var dummyCategoryEmpty = model.Category{}

type CategoryMockResponseWithArray struct {
	StatusCode int
	Message    string
	Data       []model.Category
}

type CategoryMockResponse struct {
	StatusCode int
	Message    string
	Data       model.Category
}

type CategoryMockErrorResponse struct {
	Message string
}

type categoryUseCaseMock struct {
	mock.Mock
}

func (c *categoryUseCaseMock) GetAllCategory() ([]model.Category, error) {
	args := c.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Category), args.Error(1)
}

func (c *categoryUseCaseMock) GetCategoryByID(id uint) (*model.Category, error) {
	args := c.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Category), args.Error(1)
}

type CategoryApiTestSuite struct {
	suite.Suite
	useCaseTest     usecase.CategoryUsecaseEntity
	routerTest      *gin.Engine
	routerGroupTest *gin.RouterGroup
}

func (suite *CategoryApiTestSuite) SetupTest() {
	suite.useCaseTest = new(categoryUseCaseMock)
	suite.routerTest = gin.Default()
	suite.routerGroupTest = suite.routerTest.Group("/apimock")
}

func (suite *CategoryApiTestSuite) TestCategoryApi_NewCategoryRoute() {
	NewCategoryRoute(suite.routerGroupTest, suite.useCaseTest)
}

func (suite *CategoryApiTestSuite) TestCategoryApi_CategoryGetAllHandler_Success() {
	suite.useCaseTest.(*categoryUseCaseMock).On("GetAllCategory").Return(dummyCategories, nil)
	delivery := CategoryRoute{categoryUsecase: suite.useCaseTest}
	handler := delivery.CategoryGetAllHandler

	getAll := "/apimock/category"
	categoryRoute := suite.routerGroupTest.Group("/category")
	categoryRoute.GET("", handler)

	rr := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, getAll, nil)
	request.Header.Set("Accept", "application/json")

	suite.routerTest.ServeHTTP(rr, request)
	assert.Equal(suite.T(), 200, rr.Code)

	a := rr.Body.String()
	actualCategories := new(CategoryMockResponseWithArray)
	json.Unmarshal([]byte(a), actualCategories)
	assert.Equal(suite.T(), dummyCategories[0].ID, actualCategories.Data[0].ID)
}

func (suite *CategoryApiTestSuite) TestCategoryApi_CategoryGetAllHandler_Failed() {
	suite.useCaseTest.(*categoryUseCaseMock).On("GetAllCategory").Return(nil, errors.New("failed"))
	delivery := CategoryRoute{categoryUsecase: suite.useCaseTest}
	handler := delivery.CategoryGetAllHandler

	getAll := "/apimock/category"
	categoryRoute := suite.routerGroupTest.Group("/category")
	categoryRoute.GET("", handler)

	rr := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, getAll, nil)
	request.Header.Set("Accept", "application/json")

	suite.routerTest.ServeHTTP(rr, request)
	assert.Equal(suite.T(), 401, rr.Code)

	a := rr.Body.String()
	actualCategories := new(CategoryMockResponseWithArray)
	json.Unmarshal([]byte(a), actualCategories)
	assert.Nil(suite.T(), actualCategories.Data)
}

func (suite *CategoryApiTestSuite) TestCategoryApi_CategoryGetByIDHandler_Success() {
	dummyCategory := dummyCategories[0]
	suite.useCaseTest.(*categoryUseCaseMock).On("GetCategoryByID", uint(1)).Return(&dummyCategory, nil)
	delivery := CategoryRoute{categoryUsecase: suite.useCaseTest}
	handler := delivery.CategoryGetByIDHandler

	getById := "/apimock/category/1"
	categoryRoute := suite.routerGroupTest.Group("/category")
	categoryRoute.GET("/:id", handler)

	rr := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, getById, nil)

	suite.routerTest.ServeHTTP(rr, request)
	assert.Equal(suite.T(), 200, rr.Code)

	a := rr.Body.String()
	actualCategory := new(CategoryMockResponse)
	json.Unmarshal([]byte(a), actualCategory)
	assert.Equal(suite.T(), dummyCategory.ID, actualCategory.Data.ID)
}

func (suite *CategoryApiTestSuite) TestDivisionApi_DivisionGetByIDHandler_Failed() {
	suite.useCaseTest.(*categoryUseCaseMock).On("GetCategoryByID", uint(1)).Return(nil, errors.New("failed"))
	delivery := CategoryRoute{categoryUsecase: suite.useCaseTest}
	handler := delivery.CategoryGetByIDHandler

	getById := "/apimock/category/1"
	categoryRoute := suite.routerGroupTest.Group("/category")
	categoryRoute.GET("/:id", handler)

	rr := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, getById, nil)

	suite.routerTest.ServeHTTP(rr, request)
	assert.Equal(suite.T(), 200, rr.Code)

	a := rr.Body.String()
	actualCategory := new(CategoryMockResponse)
	json.Unmarshal([]byte(a), actualCategory)
	assert.Equal(suite.T(), dummyCategoryEmpty, actualCategory.Data)
}

func TestCategoryApiSuite(t *testing.T) {
	suite.Run(t, new(CategoryApiTestSuite))
}
