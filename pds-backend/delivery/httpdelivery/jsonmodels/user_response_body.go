package jsonmodels

import "pds-backend/orm/gorm/model"

type UserProfileResponse struct {
	Fullname    string `json:"fullName"`
	Email       string `json:"email"`
	Picture     string `json:"picture"`
	Gender      string `json:"gender"`
	DivisionID  uint   `json:"divisionId"`
	Division    string `json:"divisionName"`
	RoleID      uint   `json:"roleId"`
	Role        string `json:"roleName"`
	Bio         string `json:"biography"`
	FirebaseUid string `json:"firebaseUid"`
}

type ListActivity struct {
	Row         []model.Activity `json:"row"`
	CurrentPage uint             `json:"currentPage"`
	TotalPage   uint             `json:"totalPage"`
}
