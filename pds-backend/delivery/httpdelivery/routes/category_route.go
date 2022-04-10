package routes

import (
	"pds-backend/delivery/httpdelivery/httpconstant"
	"pds-backend/delivery/httpdelivery/httperror"
	"pds-backend/delivery/httpdelivery/httpresponse"
	"pds-backend/delivery/httpdelivery/jsonmodels"
	"pds-backend/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryRouteEntity interface{}

type CategoryRoute struct {
	categoryUsecase usecase.CategoryUsecaseEntity
}

func NewCategoryRoute(apiRoute *gin.RouterGroup, categoryUsecase usecase.CategoryUsecaseEntity) {
	categoryRoute := CategoryRoute{categoryUsecase: categoryUsecase}
	apiRoute.GET("", categoryRoute.CategoryGetAllHandler)
	apiRoute.GET(httpconstant.ROUTE_PARAM_ID, categoryRoute.CategoryGetByIDHandler)
}

func (t *CategoryRoute) CategoryGetAllHandler(c *gin.Context) {
	response := httpresponse.NewHttpResponse(c)
	categoryList, err := t.categoryUsecase.GetAllCategory()
	if err != nil {
		response.SendError(httpresponse.NewUnauthorizedMessage(httperror.ErrUnauthorized.Error(), nil))
		return
	}

	responseData := make([]jsonmodels.CategoryResponseBody, len(categoryList))

	for i := 0; i < len(categoryList); i++ {
		responseData[i].Id = categoryList[i].ID
		responseData[i].Name = categoryList[i].Name
	}

	response.SendData(httpresponse.NewSuccessMessage(httpresponse.SUCCESS_GETALLCATEGORY, responseData))
}

func (t *CategoryRoute) CategoryGetByIDHandler(c *gin.Context) {
	response := httpresponse.NewHttpResponse(c)
	id, _ := strconv.Atoi(c.Param("id"))
	category, err := t.categoryUsecase.GetCategoryByID(uint(id))
	if err != nil {
		var emptyData = make([]struct{}, 0)
		response.SendData(httpresponse.NewSuccessMessage(httpresponse.SUCCESS_GETCATEGORYBYID, emptyData))
		return
	}

	responseData := jsonmodels.CategoryResponseBody{}
	responseData.Id = category.ID
	responseData.Name = category.Name

	response.SendData(httpresponse.NewSuccessMessage(httpresponse.SUCCESS_GETCATEGORYBYID, responseData))
}
