package model

import "gorm.io/gorm"

type Activity struct {
	UserID *uint `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Type   *string
	Header *string
	Body   *string
	gorm.Model
}
