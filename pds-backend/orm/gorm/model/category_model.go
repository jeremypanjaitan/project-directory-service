package model

import "gorm.io/gorm"

type Category struct {
	Name    *string `json:"name"`
	Project []Project
	gorm.Model
}
