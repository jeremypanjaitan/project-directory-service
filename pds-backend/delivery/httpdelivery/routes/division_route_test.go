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

var nameDivision1 = "Name1"
var nameDivision2 = "Name2"

var dummyDivisions = []model.Division{
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

var dummyDivisionEmpty = model.Division{}

type DivisionMockResponseWithArray struct {
	StatusCode int
	Message    string
	Data       []model.Division
}

type DivisionMockResponse struct {
	StatusCode int
	Message    string
	Data       model.Division
}

type divisionUseCaseMock struct {
	mock.Mock
}

func (d *divisionUseCaseMock) GetAllDivision() ([]model.Division, error) {
	args := d.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Division), args.Error(1)
}

func (d *divisionUseCaseMock) GetDivisionByID(id uint) (*model.Division, error) {
	args := d.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Division), args.Error(1)
}

type DivisionApiTestSuite struct {
	suite.Suite
	useCaseTest     usecase.DivisionUsecaseEntity
	routerTest      *gin.Engine
	routerGroupTest *gin.RouterGroup
}

func (suite *DivisionApiTestSuite) SetupTest() {
	suite.useCaseTest = new(divisionUseCaseMock)
	suite.routerTest = gin.Default()
	suite.routerGroupTest = suite.routerTest.Group("/apimock")
}

func (suite *DivisionApiTestSuite) TestDivisionApi_NewDivisionRoute() {
	NewDivisionRoute(suite.routerGroupTest, suite.useCaseTest)
}

func (suite *DivisionApiTestSuite) TestDivisionApi_DivisionGetAllHandler_Success() {
	suite.useCaseTest.(*divisionUseCaseMock).On("GetAllDivision").Return(dummyDivisions, nil)
	delivery := DivisionRoute{divisionUsecase: suite.useCaseTest}
	handler := delivery.DivisionGetAllHandler

	getAll := "/apimock/division"
	divisionRoute := suite.routerGroupTest.Group("/division")
	divisionRoute.GET("", handler)

	rr := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, getAll, nil)
	request.Header.Set("Accept", "application/json")

	suite.routerTest.ServeHTTP(rr, request)
	assert.Equal(suite.T(), 200, rr.Code)

	a := rr.Body.String()
	actualDivisions := new(DivisionMockResponseWithArray)
	json.Unmarshal([]byte(a), actualDivisions)
	assert.Equal(suite.T(), dummyDivisions[0].ID, actualDivisions.Data[0].ID)
}

func (suite *DivisionApiTestSuite) TestDivisionApi_DivisionGetAllHandler_Failed() {
	suite.useCaseTest.(*divisionUseCaseMock).On("GetAllDivision").Return(nil, errors.New("failed"))
	delivery := DivisionRoute{divisionUsecase: suite.useCaseTest}
	handler := delivery.DivisionGetAllHandler

	getAll := "/apimock/division"
	divisionRoute := suite.routerGroupTest.Group("/division")
	divisionRoute.GET("", handler)

	rr := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, getAll, nil)
	request.Header.Set("Accept", "application/json")

	suite.routerTest.ServeHTTP(rr, request)
	assert.Equal(suite.T(), 401, rr.Code)

	a := rr.Body.String()
	actualDivisions := new(DivisionMockResponseWithArray)
	json.Unmarshal([]byte(a), actualDivisions)
	assert.Nil(suite.T(), actualDivisions.Data)
}

func (suite *DivisionApiTestSuite) TestDivisionApi_DivisionGetByIDHandler_Success() {
	dummyDivision := dummyDivisions[0]
	suite.useCaseTest.(*divisionUseCaseMock).On("GetDivisionByID", uint(1)).Return(&dummyDivision, nil)
	delivery := DivisionRoute{divisionUsecase: suite.useCaseTest}
	handler := delivery.DivisionGetByIDHandler

	getById := "/apimock/division/1"
	categoryRoute := suite.routerGroupTest.Group("/division")
	categoryRoute.GET("/:id", handler)

	rr := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, getById, nil)

	suite.routerTest.ServeHTTP(rr, request)
	assert.Equal(suite.T(), 200, rr.Code)

	a := rr.Body.String()
	actualDivision := new(DivisionMockResponse)
	json.Unmarshal([]byte(a), actualDivision)
	assert.Equal(suite.T(), dummyDivision.ID, actualDivision.Data.ID)
}

func (suite *DivisionApiTestSuite) TestDivisionApi_DivisionGetByIDHandler_Failed() {
	suite.useCaseTest.(*divisionUseCaseMock).On("GetDivisionByID", uint(1)).Return(nil, errors.New("failed"))
	delivery := DivisionRoute{divisionUsecase: suite.useCaseTest}
	handler := delivery.DivisionGetByIDHandler

	getById := "/apimock/division/1"
	categoryRoute := suite.routerGroupTest.Group("/division")
	categoryRoute.GET("/:id", handler)

	rr := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, getById, nil)

	suite.routerTest.ServeHTTP(rr, request)
	assert.Equal(suite.T(), 200, rr.Code)

	a := rr.Body.String()
	actualDivision := new(DivisionMockResponse)
	json.Unmarshal([]byte(a), actualDivision)
	assert.Equal(suite.T(), dummyDivisionEmpty, actualDivision.Data)
}

func TestDivisionApiSuite(t *testing.T) {
	suite.Run(t, new(DivisionApiTestSuite))
}
