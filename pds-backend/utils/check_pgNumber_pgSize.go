package utils

import (
	"pds-backend/apperror"
	"pds-backend/delivery/httpdelivery/httpresponse"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CheckPageNumberAndPageSize (pageNumber string, pageSize string, c *gin.Context) (int, int) {
	response := httpresponse.NewHttpResponse(c)
	if pageNumber == "" || pageSize == "" {
		response.SendError(httpresponse.NewBadRequestMessage(apperror.ErrRequiredField.Error(), nil))
		return 0, 0
	} else {
		convPageSize, _ := strconv.Atoi(pageSize)
		convPageNumber, _ := strconv.Atoi(pageNumber)
		return convPageNumber, convPageSize
	}
}