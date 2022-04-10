package routes

import (
	"log"
	"pds-backend/apperror"
	"pds-backend/constant"
	"pds-backend/delivery/httpdelivery/httpconstant"
	"pds-backend/delivery/httpdelivery/httpresponse"
	"pds-backend/delivery/httpdelivery/jsonmodels"
	"pds-backend/orm/gorm/model"
	"pds-backend/usecase"
	"pds-backend/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ProjectRouteEntity interface {
}

type ProjectRoute struct {
	projectUsecase usecase.ProjectUsecaseEntity
}

func NewProjectRoute(apiRoute *gin.RouterGroup, projectUsecase usecase.ProjectUsecaseEntity) {
	projectRoute := ProjectRoute{projectUsecase: projectUsecase}
	apiRoute.POST("", projectRoute.createProjectHandler)
	apiRoute.GET(httpconstant.ROUTE_PROJECT_ID, projectRoute.findProjectByIdHandler)
	apiRoute.GET("", projectRoute.showListOfProjectHandler)
	apiRoute.DELETE(httpconstant.ROUTE_PROJECT_ID, projectRoute.deleteProjectHandler)
	apiRoute.PUT(httpconstant.ROUTE_PROJECT_ID, projectRoute.updateProjectHandler)
}

func (p *ProjectRoute) createProjectHandler(c *gin.Context) {
	var projectRequestBody jsonmodels.ProjectRequestBody
	response := httpresponse.NewHttpResponse(c)
	userEmail := c.GetString(constant.EMAIL)

	if err := c.ShouldBindJSON(&projectRequestBody); err != nil {
		response.SendError(httpresponse.NewBadRequestMessage(apperror.ErrRequiredField.Error(), err.Error()))
		return
	}

	userId, err := p.projectUsecase.FindIdByEmail(userEmail)
	if err != nil {
		log.Println(err)
		return
	}

	project := model.Project{
		Title:       &projectRequestBody.Title,
		Picture:     &projectRequestBody.Picture,
		Description: &projectRequestBody.Description,
		Story:       &projectRequestBody.Story,
		CategoryID:  &projectRequestBody.CategoryID,
		UserID:      userId,
	}

	if !utils.CheckProjectResponseStringLength(project) {
		response.SendError(httpresponse.NewBadRequestMessage(apperror.ErrConstraintStringLength.Error(), ""))
		return
	}

	createdProject, err := p.projectUsecase.CreateProject(project)
	if err != nil {
		if strings.Contains(err.Error(), apperror.ErrDuplicateValue.Error()) {
			response.SendError(httpresponse.NewDuplicateValueMessage(apperror.ErrDuplicateValue.Error(), err.Error()))
			return
		}
		response.SendError(httpresponse.NewInternalServerErrorMessage(httpconstant.INTERNAL_SERVER_ERROR, err.Error()))
		return
	}
	response.SendData(httpresponse.ResponseMessage{Status: 200, Description: "success to create project", Data: createdProject})
}

func (p *ProjectRoute) findProjectByIdHandler(c *gin.Context) {
	response := httpresponse.NewHttpResponse(c)
	var projectResponseBody jsonmodels.ProjectResponseBody
	projectId, _ := strconv.Atoi(c.Param("id"))
	userEmail := c.GetString(constant.EMAIL)

	userId, err := p.projectUsecase.FindIdByEmail(userEmail)
	if err != nil {
		response.SendError(httpresponse.NewBadRequestMessage(apperror.ErrRequiredField.Error(), err.Error()))
		return
	}

	getProjectById, err := p.projectUsecase.GetProjectDetails(uint(projectId))
	if err != nil {
		response.SendError(httpresponse.NewBadRequestMessage(apperror.ErrRequiredField.Error(), err.Error()))
		return
	}

	projectResponseBody = jsonmodels.ProjectResponseBody{
		Title:         *getProjectById.Title,
		Picture:       *getProjectById.Picture,
		Description:   *getProjectById.Description,
		Story:         *getProjectById.Story,
		CategoryID:    *getProjectById.CategoryID,
		UserID:        *getProjectById.UserID,
		CanDelete:     utils.CheckUserId(*userId, *getProjectById.UserID),
		CanEdit:       utils.CheckUserId(*userId, *getProjectById.UserID),
		TotalLikes:    *getProjectById.TotalLikes,
		TotalViews:    *getProjectById.TotalViews,
		TotalComments: *getProjectById.TotalComments,
		CreatedAt:     getProjectById.CreatedAt,
	}

	response.SendData(httpresponse.ResponseMessage{Status: 200, Description: "Project data", Data: projectResponseBody})
}

func (p *ProjectRoute) showListOfProjectHandler(c *gin.Context) {
	response := httpresponse.NewHttpResponse(c)
	pageNumber := c.Query("pageNumber")
	pageSize := c.Query("pageSize")
	titleSearch := c.Query("searchByTitle")
	category := c.Query("categoryId")
	sortByLikes := c.Query("sortByLikes")

	convPageNumber, convPageSize := utils.CheckPageNumberAndPageSize(pageNumber, pageSize, c)

	list, pagination, err := p.projectUsecase.ShowListProject(uint(convPageNumber), uint(convPageSize), titleSearch, category, sortByLikes)
	if err != nil {
		response.SendError(httpresponse.NewBadRequestMessage(apperror.ErrRequiredField.Error(), err.Error()))
		return
	}

	rowsResponse := utils.RowsHandler(list)
	rows := utils.PageNumberHandler(convPageNumber, *pagination, rowsResponse)

	responseBody := jsonmodels.ListProject{
		Row:         rows,
		CurrentPage: uint(convPageNumber),
		TotalPage:   *pagination,
	}

	response.SendData(httpresponse.ResponseMessage{Status: 200, Description: "ok", Data: responseBody})
}

func (p *ProjectRoute) deleteProjectHandler(c *gin.Context) {
	response := httpresponse.NewHttpResponse(c)
	projectId, _ := strconv.Atoi(c.Param("id"))

	err := p.projectUsecase.DeleteProject(uint(projectId))
	if err != nil {
		response.SendError(httpresponse.NewBadRequestMessage(apperror.ErrRequiredField.Error(), err.Error()))
		return
	}

	response.SendData(httpresponse.NewSuccessMessage(constant.SUCCESS, projectId))
}

func (p *ProjectRoute) updateProjectHandler(c *gin.Context) {
	var projectRequestBody jsonmodels.ProjectRequestBody
	response := httpresponse.NewHttpResponse(c)
	projectId, _ := strconv.Atoi(c.Param("id"))

	if err := c.ShouldBindJSON(&projectRequestBody); err != nil {
		response.SendError(httpresponse.NewBadRequestMessage(apperror.ErrRequiredField.Error(), err.Error()))
		log.Println(err)
		return
	}

	updatedProject := model.Project{
		Title:       &projectRequestBody.Title,
		Picture:     &projectRequestBody.Picture,
		Description: &projectRequestBody.Description,
		Story:       &projectRequestBody.Story,
		CategoryID:  &projectRequestBody.CategoryID,
	}

	if !utils.CheckProjectResponseStringLength(updatedProject) {
		response.SendError(httpresponse.NewBadRequestMessage(apperror.ErrConstraintStringLength.Error(), ""))
		return
	}

	updatedProjectData, err := p.projectUsecase.UpdateProject(updatedProject, uint(projectId))
	if err != nil {
		response.SendError(httpresponse.NewBadRequestMessage(apperror.ErrFailedUpdateProject.Error(), err.Error()))
		log.Println(err)
		return
	}

	project := jsonmodels.UpdatedProjectResponseBody{
		Title:       *updatedProjectData.Title,
		Picture:     *updatedProject.Picture,
		Description: *updatedProject.Description,
		Story:       *updatedProjectData.Story,
		CategoryID:  *updatedProjectData.CategoryID,
		UpdatedAt:   updatedProject.UpdatedAt,
	}

	response.SendData(httpresponse.NewSuccessMessage(constant.SUCCESS, project))
}
