package model

import "gorm.io/gorm"

type Division struct {
	Name *string `json:"name"`
	User []User
	gorm.Model
}
