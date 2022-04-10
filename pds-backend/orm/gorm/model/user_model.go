package model

import "gorm.io/gorm"

type User struct {
	FullName   *string `json:"fullName" binding:"required"`
	Email      *string `json:"email" gorm:"unique" binding:"required"`
	Password   *string `json:"password" binding:"required"`
	Picture    *string `json:"picture"`
	Gender     *string `json:"gender" binding:"required"`
	Biography  *string `json:"biography"`
	DivisionID *uint   `json:"divisionId" binding:"required" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	RoleID     *uint   `json:"roleId" binding:"required" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" `
	Project    []Project
	Activity   []Activity
	Likes      []*Project `gorm:"many2many:like;"`
	Views      []*Project `gorm:"many2many:view;"`
	Comments   []*Comment
	gorm.Model
}
