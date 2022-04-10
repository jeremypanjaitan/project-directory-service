package jsonmodels

type UserChangePassword struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type UserChangeProfileData struct {
	Fullname   string `json:"fullName"`
	Picture    string `json:"picture"`
	Gender     string `json:"gender"`
	DivisionID uint   `json:"division"`
	RoleID     uint   `json:"role"`
	Bio        string `json:"biography"`
}
