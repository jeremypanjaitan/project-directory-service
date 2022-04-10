package routes

import (
	// "bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"pds-backend/delivery/httpdelivery/jsonmodels"
	"pds-backend/orm/gorm/model"
	"pds-backend/usecase"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var title = "project ini"
var title2 = "project itu"
var pic = "photo ini"
var pic2 = "photo itu"
var desc = "desc ini"
var desc2 = "desc itu"
var story = "story ini"
var story2 = "story itu"
var category = uint(1)
var category2 = uint(2)
var user = uint(1)
var user2 = uint(2)

var dummyProject = model.Project{
	Title:       &title,
	Picture:     &pic,
	Description: &desc,
	Story:       &story,
	CategoryID:  &category,
	UserID:      &user,
}

var dummyProjects = []model.Project{
	{
		Title:       &title,
		Picture:     &pic,
		Description: &desc,
		Story:       &story,
		CategoryID:  &category,
		UserID:      &user,
	},
	{
		Title: &title2,
		Picture: &pic2,
		Description: &desc2,
		Story: &story2,
		CategoryID: &category2,
		UserID: &user2,
	},
}

type ProjectMockResponse struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
	Data       model.Project
}

type ProjectPaginationMockResponse struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
	Data       jsonmodels.ListProject
}

type ProjectMockErrorResponse struct {
	Message string
}

type projectUsecaseMock struct {
	mock.Mock
}

type ProjectApiTestSuite struct {
	suite.Suite
	usecaseTest     usecase.ProjectUsecaseEntity
	routerTest      *gin.Engine
	routerGroupTest *gin.RouterGroup
}

func (suite *ProjectApiTestSuite) SetupTest() {
	suite.usecaseTest = new(projectUsecaseMock)
	suite.routerTest = gin.Default()
	suite.routerGroupTest = suite.routerTest.Group("/apimock")
}

func TestProjectApiTestSuite(t *testing.T) {
	suite.Run(t, new(ProjectApiTestSuite))
}

func (p *projectUsecaseMock) CreateProject(project model.Project) (*model.Project, error) {
	args := p.Called(project)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Project), nil
}

func (p *projectUsecaseMock) FindProjectById(id uint) (*model.Project, error) {
	args := p.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Project), nil
}

func (p *projectUsecaseMock) FindIdByEmail(email string) (*uint, error) {
	args := p.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*uint), nil
}

func (p *projectUsecaseMock) ShowListProject(pageNumber uint, pageSize uint, titleSearch string) ([]model.Project, *uint, error) {
	args := p.Called(pageNumber, pageSize, titleSearch)
	if args.Get(0) == nil || args.Get(1) == nil {
		return nil, nil, args.Error(2)
	}
	return args.Get(0).([]model.Project), args.Get(1).(*uint), nil
}

// func (suite *ProjectApiTestSuite) TestProjectApi_CreateProject_Success() {
// 	dummyEmail := ""
// 	dummyUserID := uint(8)
// 	suite.usecaseTest.(*projectUsecaseMock).On("FindIdByEmail", dummyEmail).Return(&dummyUserID, nil)
// 	suite.usecaseTest.(*projectUsecaseMock).On("CreateProject", dummyProject).Return(&dummyProject, nil)
// 	delivery := ProjectRoute{projectUsecase: suite.usecaseTest}
// 	handler := delivery.createProjectHandler
// 	pathRoute := suite.routerGroupTest.Group("/project")
// 	pathRoute.POST("", handler)

// 	rr := httptest.NewRecorder()
// 	reqBody, _ := json.Marshal(dummyProject)
// 	req, _ := http.NewRequest(http.MethodPost, "/apimock/project", bytes.NewBuffer(reqBody))

// 	suite.routerTest.ServeHTTP(rr, req)
// 	assert.Equal(suite.T(), 200, rr.Code)
// 	// a := rr.Body.String()

// // 	actualProject := new(ProjectMockResponse)
// // 	json.Unmarshal([]byte(a), actualProject)
// // 	assert.Equal(suite.T(), dummyProject.Title, actualProject.Data)
// }

func (suite *ProjectApiTestSuite) TestProjectApi_FindProjectByIdHandler_Success() {
	suite.usecaseTest.(*projectUsecaseMock).On("FindProjectById", uint(1)).Return(&dummyProject, nil)
	delivery := ProjectRoute{projectUsecase: suite.usecaseTest}
	handler := delivery.findProjectByIdHandler

	getById := "/apimock/project/1"
	categoryRoute := suite.routerGroupTest.Group("/project")
	categoryRoute.GET("/:id", handler)

	rr := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, getById, nil)

	suite.routerTest.ServeHTTP(rr, req)
	assert.Equal(suite.T(), 200, rr.Code)

	a := rr.Body.String()
	actual := new(ProjectMockResponse)
	json.Unmarshal([]byte(a), actual)
	log.Println(actual.Message)
	assert.Equal(suite.T(), dummyProject.Title, actual.Data.Title)
}

func (suite *ProjectApiTestSuite) TestProjectApi_FindProjectByIdHandler_Failed() {
	suite.usecaseTest.(*projectUsecaseMock).On("FindProjectById", uint(1)).Return(nil, errors.New("failed"))
	delivery := ProjectRoute{projectUsecase: suite.usecaseTest}
	handler := delivery.findProjectByIdHandler

	getById := "/apimock/project/1"
	categoryRoute := suite.routerGroupTest.Group("/project")
	categoryRoute.GET("/:id", handler)

	rr := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, getById, nil)

	suite.routerTest.ServeHTTP(rr, req)
	assert.Equal(suite.T(), 400, rr.Code)

	a := rr.Body.String()
	actual := new(ProjectMockErrorResponse)
	json.Unmarshal([]byte(a), actual)
	// log.Println("COAAA",actual)
	// assert.Equal(suite.T(), dummyProject.Title, actual.Data.Title)
}

func (suite *ProjectApiTestSuite) TestProjectApi_showListOfProjectHandler_Success() {
	dummyPagination := uint(3)
	suite.usecaseTest.(*projectUsecaseMock).On("ShowListProject", uint(1), uint(3), "project").Return(dummyProjects, &dummyPagination, nil)
	delivery := ProjectRoute{projectUsecase: suite.usecaseTest}
	handler := delivery.showListOfProjectHandler

	pagination := "/apimock/project?pageNumber=1&pageSize=3&searchQuery=project"
	categoryRoute := suite.routerGroupTest.Group("/project")
	categoryRoute.GET("", handler)

	rr := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, pagination, nil)
	req.URL.RawQuery = req.URL.Query().Encode()
	log.Println(req)

	suite.routerTest.ServeHTTP(rr, req)
	assert.Equal(suite.T(), 200, rr.Code)

	a := rr.Body.String()
	actual := new(ProjectPaginationMockResponse)
	json.Unmarshal([]byte(a), actual)
	assert.Equal(suite.T(), uint(1), actual.Data.CurrentPage)
}
