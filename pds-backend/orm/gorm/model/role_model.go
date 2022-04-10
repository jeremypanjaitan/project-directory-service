package model

import "gorm.io/gorm"

type Role struct {
	Name *string `json:"name"`
	User []User
	gorm.Model
}
