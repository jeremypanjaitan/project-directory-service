package model

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	UserID    *uint
	ProjectID *uint
	Body      *string
	gorm.Model
}

type CommentJointUser struct {
	FullName  *string   `json:"from"`
	Picture   *string   `json:"picture"`
	ID        uint      `json:"commentId"`
	Body      *string   `json:"body"`
	CreatedAt time.Time `json:"createdAt"`
}
