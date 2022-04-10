package jsonmodels

type VerifyCustomTokenReqBody struct {
	Token             string `json:"token"`
	ReturnSecureToken bool   `json:"returnSecureToken"`
}
