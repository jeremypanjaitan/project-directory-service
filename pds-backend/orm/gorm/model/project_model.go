package model

import "gorm.io/gorm"

type Project struct {
	Title       *string `json:"title" gorm:"index:idx_title"`
	Picture     *string `json:"picture"`
	Description *string `json:"description"`
	Story       *string `json:"story"`
	CategoryID  *uint   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID      *uint   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Likes       []*User `gorm:"many2many:likes;"`
	Views       []*User `gorm:"many2many:view;"`
	Comments    []*Comment
	gorm.Model
}

type ProjectWithLikeViewComment struct {
	Title         *string
	Picture       *string
	Description   *string
	Story         *string
	CategoryID    *uint
	UserID        *uint
	TotalLikes    *uint
	TotalViews    *uint
	TotalComments *uint
	gorm.Model
}
