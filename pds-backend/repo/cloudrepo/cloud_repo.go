package repo

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	cloudengine "pds-backend/cloudengine/firebase"
	"pds-backend/config"
	"pds-backend/constant"
	"pds-backend/delivery/httpdelivery/jsonmodels"
	"pds-backend/utils"

	"firebase.google.com/go/auth"
	"firebase.google.com/go/messaging"
)

type CloudRepoEntity interface {
	CreateUser(email string, password string, displayName string) (*auth.UserRecord, error)
	SendEmailVer(idToken string) error
	CreateToken(uid string) (string, error)
	CreateEmailLinkVer(email string) (string, error)
	GetUserByEmail(email string) (*auth.UserRecord, error)
	VerifyIDToken(idToken string) (*auth.Token, error)
	GetUserByUid(uid string) (*auth.UserRecord, error)
	CreateCustomToken(uid string) (string, error)
	VerifyCustomToken(token string) (jsonmodels.VerifyCustomTokenResBody, error)
	CreateChangePasswordLink(user *auth.UserRecord) (string, error)
	CreateNewToken(refreshToken string) (jsonmodels.RefreshTokenResBody, error)
	GetFcmToken(uid string) (string, error)
	SendPushNotification(fcmToken string, message string, data map[string]string) error
	DeleteFcmToken(uid string) error
}

type CloudRepo struct {
	firebaseEngine cloudengine.FirebaseCloudEngineEntity
	appConfig      config.AppConfigEntity
}

func NewCloudRepo(firebaseEngine cloudengine.FirebaseCloudEngineEntity, appConfig config.AppConfigEntity) CloudRepoEntity {
	return &CloudRepo{firebaseEngine: firebaseEngine, appConfig: appConfig}
}

func (c *CloudRepo) CreateUser(email string, password string, displayName string) (*auth.UserRecord, error) {
	params := (&auth.UserToCreate{}).
		Email(email).
		EmailVerified(false).
		Password(password).
		DisplayName(displayName)

	return c.firebaseEngine.GetFirebaseAuth().CreateUser(context.Background(), params)
}

func (c *CloudRepo) SendEmailVer(idToken string) error {
	log.Println(idToken)
	var emailVerBody jsonmodels.EmailVerBody
	emailVerBody.IdToken = idToken
	emailVerBody.RequestType = constant.VERIFY_EMAIL

	req, err := json.Marshal(emailVerBody)
	if err != nil {
		log.Println(err)
		return err
	}
	resp, err := utils.PostRequest(fmt.Sprintf(constant.FIREBASE_EMAIL_VER, c.appConfig.GetFirebaseWebApiKey()), req)
	log.Println(err)
	log.Println(resp)
	return err
}

func (c *CloudRepo) CreateToken(uid string) (string, error) {
	return c.firebaseEngine.GetFirebaseAuth().CustomToken(context.Background(), uid)
}

func (c *CloudRepo) CreateEmailLinkVer(email string) (string, error) {
	return c.firebaseEngine.GetFirebaseAuth().EmailVerificationLink(context.Background(), email)
}

func (c *CloudRepo) GetUserByEmail(email string) (*auth.UserRecord, error) {
	return c.firebaseEngine.GetFirebaseAuth().GetUserByEmail(context.Background(), email)
}

func (c *CloudRepo) VerifyIDToken(idToken string) (*auth.Token, error) {
	return c.firebaseEngine.GetFirebaseAuth().VerifyIDToken(context.Background(), idToken)
}

func (c *CloudRepo) GetUserByUid(uid string) (*auth.UserRecord, error) {
	return c.firebaseEngine.GetFirebaseAuth().GetUser(context.Background(), uid)
}
func (c *CloudRepo) CreateCustomToken(uid string) (string, error) {
	return c.firebaseEngine.GetFirebaseAuth().CustomToken(context.Background(), uid)
}
func (c *CloudRepo) VerifyCustomToken(token string) (jsonmodels.VerifyCustomTokenResBody, error) {
	var verifyCustomTokenReqBody jsonmodels.VerifyCustomTokenReqBody
	var verifyCustomTokenResBody jsonmodels.VerifyCustomTokenResBody
	verifyCustomTokenReqBody.Token = token
	verifyCustomTokenReqBody.ReturnSecureToken = true

	req, err := json.Marshal(verifyCustomTokenReqBody)
	if err != nil {
		return jsonmodels.VerifyCustomTokenResBody{}, err
	}
	resp, err := utils.PostRequest(fmt.Sprintf(constant.GOOGLE_VERIFY_CUSTOM_TOKEN, c.appConfig.GetFirebaseWebApiKey()), req)
	if err != nil {
		return jsonmodels.VerifyCustomTokenResBody{}, err
	}
	if err := json.Unmarshal(resp, &verifyCustomTokenResBody); err != nil {
		return jsonmodels.VerifyCustomTokenResBody{}, err
	}
	return verifyCustomTokenResBody, nil
}

func (c *CloudRepo) CreateChangePasswordLink(user *auth.UserRecord) (string, error) {
	return c.firebaseEngine.GetFirebaseAuth().PasswordResetLink(context.Background(), user.Email)
}
func (c *CloudRepo) CreateNewToken(refreshToken string) (jsonmodels.RefreshTokenResBody, error) {
	var refreshTokenResbody jsonmodels.RefreshTokenResBody
	refreshTokenReqBody := struct {
		GrantType    string `json:"grant_type"`
		RefreshToken string `json:"refresh_token"`
	}{
		GrantType:    "refresh_token",
		RefreshToken: refreshToken,
	}
	reqBody, err := json.Marshal(refreshTokenReqBody)
	if err != nil {
		return jsonmodels.RefreshTokenResBody{}, nil
	}
	resp, err := utils.PostRequest(fmt.Sprintf(constant.GOOGLE_REFRESH_TOKEN, c.appConfig.GetFirebaseWebApiKey()), reqBody)
	if err != nil {
		return jsonmodels.RefreshTokenResBody{}, nil
	}
	err = json.Unmarshal(resp, &refreshTokenResbody)
	if err != nil {
		return jsonmodels.RefreshTokenResBody{}, nil
	}
	newAccessToken, err := c.CreateCustomToken(refreshTokenResbody.Uid)
	if err != nil {
		return jsonmodels.RefreshTokenResBody{}, nil
	}
	refreshTokenResbody.AccessToken = newAccessToken
	return refreshTokenResbody, nil
}
func (c *CloudRepo) GetFcmToken(uid string) (string, error) {
	dsnap, err := c.firebaseEngine.GetFirebaseFirestore().Collection(constant.USERS_COLLECTION).Doc(uid).Get(context.Background())
	if err != nil {
		return "", err
	}

	m := dsnap.Data()

	aInterface := m["fcmTokens"].([]interface{})
	aString := make([]string, len(aInterface))
	for i, v := range aInterface {
		aString[i] = v.(string)
	}

	return aString[0], nil
}

func (c *CloudRepo) DeleteFcmToken(uid string) error {
	_, err := c.firebaseEngine.GetFirebaseFirestore().Collection(constant.USERS_COLLECTION).Doc(uid).Delete(context.Background())
	return err
}

func (c *CloudRepo) SendPushNotification(fcmToken string, message string, data map[string]string) error {
	data["largeIconUrl"] = constant.IMAGE_NOTIFICATION
	data["body"] = message
	data["title"] = constant.TITLE_NOTIFICATION

	messageToSend := &messaging.Message{
		Data:  data,
		Token: fcmToken,
	}

	_, err := c.firebaseEngine.GetFirebaseFcm().Send(context.Background(), messageToSend)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
