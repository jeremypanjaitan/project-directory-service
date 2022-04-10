package httpresponse

import "github.com/gin-gonic/gin"

type HttpResponseEntity interface {
	SendData(message ResponseMessage)
	SendError(message ResponseMessage)
}

type HttpResponse struct {
	ginContext *gin.Context
}

func NewHttpResponse(c *gin.Context) HttpResponseEntity {
	return &HttpResponse{ginContext: c}
}

func (h *HttpResponse) SendData(message ResponseMessage) {
	h.ginContext.JSON(message.Status, message)
}
func (h *HttpResponse) SendError(message ResponseMessage) {
	h.ginContext.AbortWithStatusJSON(message.Status, message)
}
