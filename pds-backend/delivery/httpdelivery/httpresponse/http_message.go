package httpresponse

import (
	"net/http"
)

type ResponseMessage struct {
	Status      int         `json:"status"`
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
}

func NewSuccessMessage(description string, data interface{}) ResponseMessage {
	return ResponseMessage{
		Status:      http.StatusOK,
		Description: description,
		Data:        data,
	}
}

func NewBadRequestMessage(description string, data interface{}) ResponseMessage {
	return ResponseMessage{
		Status:      http.StatusBadRequest,
		Description: description,
		Data:        data,
	}
}

func NewUnauthorizedMessage(description string, data interface{}) ResponseMessage {
	return ResponseMessage{
		Status:      http.StatusUnauthorized,
		Description: description,
		Data:        data,
	}
}

func NewDuplicateValueMessage(description string, data interface{}) ResponseMessage {
	return ResponseMessage{
		Status:      http.StatusConflict,
		Description: description,
		Data:        data,
	}
}

func NewInternalServerErrorMessage(description string, data interface{}) ResponseMessage {
	return ResponseMessage{
		Status:      http.StatusInternalServerError,
		Description: description,
		Data:        data,
	}
}
