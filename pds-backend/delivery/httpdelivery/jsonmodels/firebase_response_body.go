package jsonmodels

type VerifyCustomTokenResBody struct {
	Kind         string `json:"kind"`
	IdToken      string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
	IsNewUser    bool   `json:"isNewUser"`
}

type RefreshTokenResBody struct {
	NewRefreshToken string `json:"refresh_token"`
	AccessToken     string
	Uid             string `json:"user_id"`
}
