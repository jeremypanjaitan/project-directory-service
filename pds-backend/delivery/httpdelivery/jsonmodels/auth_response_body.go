package jsonmodels

import (
	"pds-backend/service"
)

type LoginResponseBody struct {
	TokenData service.TokenDetails `json:"tokenData"`
	UserData  LoginProfileResponse `json:"userData"`
}

type LoginResponseBodyFirebase struct {
	TokenData service.TokenDetailsFirebase `json:"tokenData"`
	UserData  LoginProfileResponse         `json:"userData"`
}

type RegisterReponseBody struct {
	ID         *uint   `json:"id"`
	FullName   *string `json:"fullName"`
	Email      *string `json:"email" gorm:"unique"`
	Gender     *string `json:"gender"`
	DivisionID *uint   `json:"divisionId"`
	RoleID     *uint   `json:"roleId"`
}

type LoginProfileResponse struct {
	Fullname    string `json:"fullName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Picture     string `json:"picture"`
	Gender      string `json:"gender"`
	DivisionID  uint   `json:"divisionId"`
	Division    string `json:"divisionName"`
	RoleID      uint   `json:"roleId"`
	Role        string `json:"roleName"`
	Bio         string `json:"biography"`
	FirebaseUID string `json:"firebaseUid"`
}
