package routes

import (
	"pds-backend/constant"
	"pds-backend/delivery/httpdelivery/httpresponse"
	"pds-backend/delivery/httpdelivery/jsonmodels"
	"pds-backend/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ViewsRouteEntity interface{}

type ViewsRoute struct {
	viewsUsecase usecase.ViewsUsecaseEntity
}

func NewViewsRoute(apiRoute *gin.RouterGroup, viewsUsecase usecase.ViewsUsecaseEntity) {
	viewsRoute := ViewsRoute{viewsUsecase: viewsUsecase}
	apiRoute.GET("", viewsRoute.ViewsGetProjectViewsHandler)
	apiRoute.POST("", viewsRoute.ViewsPostViewsHandler)
}

func (v *ViewsRoute) ViewsGetProjectViewsHandler(c *gin.Context) {
	response := httpresponse.NewHttpResponse(c)
	projectId, _ := strconv.Atoi(c.Param("id"))

	responseData := jsonmodels.ViewsResponseBody{}

	totalViews, err := v.viewsUsecase.CountViewsByProjectId(uint(projectId))

	if err != nil {
		responseData.TotalViews = 0
		response.SendData(httpresponse.NewSuccessMessage(httpresponse.SUCCESS_COUNTPROJECTVIEW, responseData))
		return
	}

	responseData.TotalViews = totalViews
	response.SendData(httpresponse.NewSuccessMessage(httpresponse.SUCCESS_COUNTPROJECTVIEW, responseData))
}

func (v *ViewsRoute) ViewsPostViewsHandler(c *gin.Context) {
	response := httpresponse.NewHttpResponse(c)
	projectId, _ := strconv.Atoi(c.Param("id"))
	userEmail := c.GetString(constant.EMAIL)

	_ = v.viewsUsecase.CreateViews(uint(projectId), userEmail)

	response.SendData(httpresponse.NewSuccessMessage(httpresponse.SUCCESS_CREATEVIEW, nil))
}
