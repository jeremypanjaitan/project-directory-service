package jsonmodels

type LoginRequestBody struct {
	Email             string `json:"email" binding:"required"`
	Password          string `json:"password" binding:"required"`
	ReturnSecureToken bool   `json:"returnSecureToken"`
}

type EmailVerBody struct {
	RequestType string `json:"requestType"`
	IdToken     string `json:"idToken"`
}
