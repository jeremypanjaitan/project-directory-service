package usecase

import (
	"errors"
	"pds-backend/orm/gorm/model"
	repoProject "pds-backend/repo/projectrepo"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
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
var dummyProjects = []model.Project{
	{
		Title:       &title,
		Picture:     &pic,
		Description: &desc,
		Story:       &story,
		CategoryID:  &category,
		UserID:      &user,
		Model:       gorm.Model{ID: 1},
	},
	{
		Title:       &title2,
		Picture:     &pic2,
		Description: &desc2,
		Story:       &story2,
		CategoryID:  &category2,
		UserID:      &user2,
		Model:       gorm.Model{ID: 1},
	},
}

type repoProjectMock struct {
	mock.Mock
}

type ProjectUsecaseTestSuite struct {
	suite.Suite
	repoTest repoProject.ProjectRepoEntity
}

func (suite *ProjectUsecaseTestSuite) SetupTest() {
	suite.repoTest = new(repoProjectMock)
}

func TestProjectUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(ProjectUsecaseTestSuite))
}

func (r *repoProjectMock) CreateOne(project model.Project) (*model.Project, error) {
	args := r.Called(project)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Project), nil
}

func (r *repoProjectMock) FindIdByEmail(email string) (*uint, error) {
	args := r.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*uint), nil
}

func (r *repoProjectMock) FindProjectById(id uint) (*model.Project, error) {
	args := r.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Project), nil
}

func (r *repoProjectMock) ShowListProjectWithPagination(pageNumber uint, pageSize uint) ([]model.Project, *uint, error) {
	args := r.Called(pageNumber, pageSize)
	if args.Get(0) == nil {
		return nil, nil, args.Error(2)
	}
	return args.Get(0).([]model.Project), args.Get(1).(*uint), nil
}

func (r *repoProjectMock) FindProjectByTitle(pageNumber uint, pageSize uint, titleSearch string) ([]model.Project, *uint, error) {
	args := r.Called(pageNumber, pageSize, titleSearch)
	if args.Get(0) == nil {
		return nil, nil, args.Error(2)
	}
	return args.Get(0).([]model.Project), args.Get(1).(*uint), nil
}

func (suite *ProjectUsecaseTestSuite) TestProjectUsecase_CreateProject_Success() {
	dummyProject := dummyProjects[0]
	suite.repoTest.(*repoProjectMock).On("CreateOne", dummyProject).Return(&dummyProject, nil)
	projectUsecaseTest := NewProjectUsecase(suite.repoTest)
	project, err := projectUsecaseTest.CreateProject(dummyProject)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummyProject.Title, project.Title)
}

func (suite *ProjectUsecaseTestSuite) TestProjectUsecase_CreateProject_Failed() {
	dummyProject := dummyProjects[0]
	suite.repoTest.(*repoProjectMock).On("CreateOne", dummyProject).Return(nil, errors.New("failed"))
	projectUsecaseTest := NewProjectUsecase(suite.repoTest)
	project, err := projectUsecaseTest.CreateProject(dummyProject)
	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), project)
	assert.Equal(suite.T(), "failed", err.Error())
}

func (suite *ProjectUsecaseTestSuite) TestProjectUsecase_FindProjectById_Success() {
	dummyProject := dummyProjects[1]
	dummyId := dummyProject.ID
	suite.repoTest.(*repoProjectMock).On("FindProjectById", dummyId).Return(&dummyProject, nil)
	projectUsecaseTest := NewProjectUsecase(suite.repoTest)
	project, err := projectUsecaseTest.FindProjectById(dummyId)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummyProject.Title, project.Title)
}

func (suite *ProjectUsecaseTestSuite) TestProjectUsecase_FindProjectById_Failed() {
	dummyProject := dummyProjects[1]
	dummyId := dummyProject.ID
	suite.repoTest.(*repoProjectMock).On("FindProjectById", dummyId).Return(nil, errors.New("failed"))
	projectUsecaseTest := NewProjectUsecase(suite.repoTest)
	project, err := projectUsecaseTest.FindProjectById(dummyId)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "failed", err.Error())
	assert.Nil(suite.T(), project)
}

func (suite *ProjectUsecaseTestSuite) TestProjectUsecase_FindIdByEmail_Success() {
	dummyEmail := "coba@example.com"
	dummyResult := uint(1)
	suite.repoTest.(*repoProjectMock).On("FindIdByEmail", dummyEmail).Return(&dummyResult, nil)
	projectUsecaseTest := NewProjectUsecase(suite.repoTest)
	id, err := projectUsecaseTest.FindIdByEmail(dummyEmail)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), &dummyResult, id)
}

func (suite *ProjectUsecaseTestSuite) TestProjectUsecase_FindIdByEmail_Failed() {
	dummyEmail := "coba@example.com"
	suite.repoTest.(*repoProjectMock).On("FindIdByEmail", dummyEmail).Return(nil, errors.New("failed"))
	projectUsecaseTest := NewProjectUsecase(suite.repoTest)
	id, err := projectUsecaseTest.FindIdByEmail(dummyEmail)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "failed", err.Error())
	assert.Nil(suite.T(), id)
}

func (suite *ProjectUsecaseTestSuite) TestProjectUsecase_ShowListProject_Success() {
	dummyPageNumber := uint(1)
	dummyPageSize := uint(2)
	dummyTotalPagination := uint(5)
	dummyTitle := ""
	suite.repoTest.(*repoProjectMock).On("FindProjectByTitle", dummyPageNumber, dummyPageSize, dummyTitle).Return(dummyProjects, &dummyTotalPagination, nil)
	suite.repoTest.(*repoProjectMock).On("ShowListProjectWithPagination", dummyPageNumber, dummyPageSize).Return(dummyProjects, &dummyTotalPagination, nil)
	projectUsecaseTest := NewProjectUsecase(suite.repoTest)
	projects, totalPagination, err := projectUsecaseTest.ShowListProject(dummyPageNumber, dummyPageSize, dummyTitle)
	if dummyTitle != "" {
		assert.Nil(suite.T(), err)
		assert.Equal(suite.T(), dummyProjects, projects)
		assert.Equal(suite.T(), &dummyTotalPagination, totalPagination)
	} else {
		assert.Nil(suite.T(), err)
		assert.Equal(suite.T(), dummyProjects, projects)
		assert.Equal(suite.T(), &dummyTotalPagination, totalPagination)
	}
}

func (suite *ProjectUsecaseTestSuite) TestProjectUsecase_ShowListProject_Failed() {
	dummyPageNumber := uint(1)
	dummyPageSize := uint(2)
	dummyTitle := "Sangkuriang"
	suite.repoTest.(*repoProjectMock).On("FindProjectByTitle", dummyPageNumber, dummyPageSize, dummyTitle).Return(nil, nil, errors.New("failed"))
	suite.repoTest.(*repoProjectMock).On("ShowListProjectWithPagination", dummyPageNumber, dummyPageSize).Return(nil, nil, errors.New("failed"))
	projectUsecaseTest := NewProjectUsecase(suite.repoTest)
	projects, totalPagination, err := projectUsecaseTest.ShowListProject(dummyPageNumber, dummyPageSize, dummyTitle)
	if dummyTitle != "" {
		assert.NotNil(suite.T(), err)
		assert.Nil(suite.T(), projects)
		assert.Nil(suite.T(), totalPagination)
		assert.Equal(suite.T(), "failed", err.Error())
	} else {
		assert.NotNil(suite.T(), err)
		assert.Nil(suite.T(), projects)
		assert.Nil(suite.T(), totalPagination)
		assert.Equal(suite.T(), "failed", err.Error())
	}
}