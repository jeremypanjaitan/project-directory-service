package jsonmodels

import (
	"pds-backend/orm/gorm/model"
	"time"
)

type CommentResponseBody struct {
	UserPhoto *string   `json:"userPhoto"`
	Body      *string   `json:"body"`
	CreatedAt time.Time `json:"createdAt"`
}

type ListComment struct {
	Row         []model.CommentJointUser `json:"row"`
	CurrentPage uint                     `json:"currentPage"`
	TotalPage   uint                     `json:"totalPage"`
}
