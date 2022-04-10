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

type RoleRouteEntity interface{}

type RoleRoute struct {
	roleUsecase usecase.RoleUsecaseEntity
}

func NewRoleRoute(apiRoute *gin.RouterGroup, roleUsecase usecase.RoleUsecaseEntity) {
	roleRoute := RoleRoute{roleUsecase: roleUsecase}
	apiRoute.GET("", roleRoute.RoleGetAllHandler)
	apiRoute.GET(httpconstant.ROUTE_PARAM_ID, roleRoute.RoleGetByIDHandler)
}

func (r *RoleRoute) RoleGetAllHandler(c *gin.Context) {
	response := httpresponse.NewHttpResponse(c)
	roleList, err := r.roleUsecase.GetAllRole()
	if err != nil {
		response.SendError(httpresponse.NewUnauthorizedMessage(httperror.ErrUnauthorized.Error(), nil))
		return
	}

	responseData := make([]jsonmodels.RoleResponseBody, len(roleList))

	for i := 0; i < len(roleList); i++ {
		responseData[i].Id = roleList[i].ID
		responseData[i].Name = roleList[i].Name
	}

	response.SendData(httpresponse.NewSuccessMessage(httpresponse.SUCCESS_GETALLROLE, responseData))
}

func (r *RoleRoute) RoleGetByIDHandler(c *gin.Context) {
	response := httpresponse.NewHttpResponse(c)
	id, _ := strconv.Atoi(c.Param("id"))
	role, err := r.roleUsecase.GetRoleByID(uint(id))
	if err != nil {
		var emptyData = make([]struct{}, 0)
		response.SendData(httpresponse.NewSuccessMessage(httpresponse.SUCCESS_GETROLEBYID, emptyData))
		return
	}

	responseData := jsonmodels.RoleResponseBody{}
	responseData.Id = role.ID
	responseData.Name = role.Name

	response.SendData(httpresponse.NewSuccessMessage(httpresponse.SUCCESS_GETROLEBYID, responseData))
}
