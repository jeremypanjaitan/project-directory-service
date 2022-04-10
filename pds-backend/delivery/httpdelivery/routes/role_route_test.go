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

var dummyRoleEmpty = model.Role{}

type RoleMockResponseWithArray struct {
	StatusCode int
	Message    string
	Data       []model.Role
}

type RoleMockResponse struct {
	StatusCode int
	Message    string
	Data       model.Role
}

type RoleMockErrorResponse struct {
	Message string
}

type roleUseCaseMock struct {
	mock.Mock
}

func (r *roleUseCaseMock) GetAllRole() ([]model.Role, error) {
	args := r.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Role), args.Error(1)
}

func (r *roleUseCaseMock) GetRoleByID(id uint) (*model.Role, error) {
	args := r.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Role), args.Error(1)
}

type RoleApiTestSuite struct {
	suite.Suite
	useCaseTest     usecase.RoleUsecaseEntity
	routerTest      *gin.Engine
	routerGroupTest *gin.RouterGroup
}

func (suite *RoleApiTestSuite) SetupTest() {
	suite.useCaseTest = new(roleUseCaseMock)
	suite.routerTest = gin.Default()
	suite.routerGroupTest = suite.routerTest.Group("/apimock")
}

func (suite *RoleApiTestSuite) TestRoleApi_NewRoleRoute() {
	NewRoleRoute(suite.routerGroupTest, suite.useCaseTest)
}

func (suite *RoleApiTestSuite) TestRoleApi_RoleGetAllHandler_Success() {
	suite.useCaseTest.(*roleUseCaseMock).On("GetAllRole").Return(dummyRoles, nil)
	delivery := RoleRoute{roleUsecase: suite.useCaseTest}
	handler := delivery.RoleGetAllHandler

	getAll := "/apimock/role"
	roleRoute := suite.routerGroupTest.Group("/role")
	roleRoute.GET("", handler)

	rr := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, getAll, nil)
	request.Header.Set("Accept", "application/json")

	suite.routerTest.ServeHTTP(rr, request)
	assert.Equal(suite.T(), 200, rr.Code)

	a := rr.Body.String()
	actualRoles := new(RoleMockResponseWithArray)
	json.Unmarshal([]byte(a), actualRoles)
	assert.Equal(suite.T(), dummyRoles[0].ID, actualRoles.Data[0].ID)
}

func (suite *RoleApiTestSuite) TestRoleApi_RoleGetAllHandler_Failed() {
	suite.useCaseTest.(*roleUseCaseMock).On("GetAllRole").Return(nil, errors.New("failed"))
	delivery := RoleRoute{roleUsecase: suite.useCaseTest}
	handler := delivery.RoleGetAllHandler

	getAll := "/apimock/role"
	roleRoute := suite.routerGroupTest.Group("/role")
	roleRoute.GET("", handler)

	rr := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, getAll, nil)
	request.Header.Set("Accept", "application/json")

	suite.routerTest.ServeHTTP(rr, request)
	assert.Equal(suite.T(), 401, rr.Code)

	a := rr.Body.String()
	actualRoles := new(RoleMockResponseWithArray)
	json.Unmarshal([]byte(a), actualRoles)
	assert.Nil(suite.T(), actualRoles.Data)
}

func (suite *RoleApiTestSuite) TestRoleApi_RoleGetByIDHandler_Success() {
	dummyRole := dummyRoles[0]
	suite.useCaseTest.(*roleUseCaseMock).On("GetRoleByID", uint(1)).Return(&dummyRole, nil)
	delivery := RoleRoute{roleUsecase: suite.useCaseTest}
	handler := delivery.RoleGetByIDHandler

	getById := "/apimock/role/1"
	roleRoute := suite.routerGroupTest.Group("/role")
	roleRoute.GET("/:id", handler)

	rr := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, getById, nil)

	suite.routerTest.ServeHTTP(rr, request)
	assert.Equal(suite.T(), 200, rr.Code)

	a := rr.Body.String()
	actualRole := new(RoleMockResponse)
	json.Unmarshal([]byte(a), actualRole)
	assert.Equal(suite.T(), dummyRole.ID, actualRole.Data.ID)
}

func (suite *RoleApiTestSuite) TestRoleApi_RoleGetByIDHandler_Failed() {
	suite.useCaseTest.(*roleUseCaseMock).On("GetRoleByID", uint(1)).Return(nil, errors.New("failed"))
	delivery := RoleRoute{roleUsecase: suite.useCaseTest}
	handler := delivery.RoleGetByIDHandler

	getById := "/apimock/role/1"
	roleRoute := suite.routerGroupTest.Group("/role")
	roleRoute.GET("/:id", handler)

	rr := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, getById, nil)

	suite.routerTest.ServeHTTP(rr, request)
	assert.Equal(suite.T(), 200, rr.Code)

	a := rr.Body.String()
	actualRole := new(RoleMockResponse)
	json.Unmarshal([]byte(a), actualRole)
	assert.Equal(suite.T(), dummyRoleEmpty, actualRole.Data)
}

func TestRoleApiSuite(t *testing.T) {
	suite.Run(t, new(RoleApiTestSuite))
}
