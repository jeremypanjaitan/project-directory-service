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

type DivisionRouteEntity interface{}

type DivisionRoute struct {
	divisionUsecase usecase.DivisionUsecaseEntity
}

func NewDivisionRoute(apiRoute *gin.RouterGroup, divisionUsecase usecase.DivisionUsecaseEntity) {
	divisionRoute := DivisionRoute{divisionUsecase: divisionUsecase}
	apiRoute.GET("", divisionRoute.DivisionGetAllHandler)
	apiRoute.GET(httpconstant.ROUTE_PARAM_ID, divisionRoute.DivisionGetByIDHandler)
}

func (d *DivisionRoute) DivisionGetAllHandler(c *gin.Context) {
	response := httpresponse.NewHttpResponse(c)
	divisionList, err := d.divisionUsecase.GetAllDivision()
	if err != nil {
		response.SendError(httpresponse.NewUnauthorizedMessage(httperror.ErrUnauthorized.Error(), nil))
		return
	}

	responseData := make([]jsonmodels.DivisionResponseBody, len(divisionList))

	for i := 0; i < len(divisionList); i++ {
		responseData[i].Id = divisionList[i].ID
		responseData[i].Name = divisionList[i].Name
	}

	response.SendData(httpresponse.NewSuccessMessage(httpresponse.SUCCESS_GETALLDIVISION, responseData))
}

func (d *DivisionRoute) DivisionGetByIDHandler(c *gin.Context) {
	response := httpresponse.NewHttpResponse(c)
	id, _ := strconv.Atoi(c.Param("id"))
	division, err := d.divisionUsecase.GetDivisionByID(uint(id))
	if err != nil {
		var emptyData = make([]struct{}, 0)
		response.SendData(httpresponse.NewSuccessMessage(httpresponse.SUCCESS_GETDIVISIONBYID, emptyData))
		return
	}

	responseData := jsonmodels.RoleResponseBody{}
	responseData.Id = division.ID
	responseData.Name = division.Name

	response.SendData(httpresponse.NewSuccessMessage(httpresponse.SUCCESS_GETDIVISIONBYID, responseData))
}
