package usecase

import (
	"log"
	"pds-backend/config"
	"pds-backend/constant"
	"pds-backend/delivery/httpdelivery/jsonmodels"
	"pds-backend/orm/gorm/model"
	repo "pds-backend/repo/cloudrepo"
	"pds-backend/utils"

	"firebase.google.com/go/auth"
)

type CloudUsecaseEntity interface {
	RegisterUser(user model.User) error
	GetUserByEmail(email string) (*auth.UserRecord, error)
	VerifyIDToken(idToken string) (*auth.Token, error)
	GetUserByUid(uid string) (*auth.UserRecord, error)
	CreateCustomToken(uid string) (string, error)
	VerifyCustomToken(token string) (jsonmodels.VerifyCustomTokenResBody, error)
	SendChangePasswordLink(email string) error
	GetNewToken(refreshToken string) (jsonmodels.RefreshTokenResBody, error)
	SendNotification(toEmail string, message string, data map[string]string) error
	DeleteFcmToken(uid string) error
}

type CloudUsecase struct {
	cloudRepo repo.CloudRepoEntity
	appConfig config.AppConfigEntity
}

func NewCloudUsecase(cloudRepo repo.CloudRepoEntity, appConfig config.AppConfigEntity) CloudUsecaseEntity {
	return &CloudUsecase{cloudRepo: cloudRepo, appConfig: appConfig}
}

func (c *CloudUsecase) RegisterUser(user model.User) error {
	userRecord, err := c.cloudRepo.CreateUser(*user.Email, *user.Password, *user.FullName)
	if err != nil {
		log.Println(err)
		return err
	}

	emailLinkVer, err := c.cloudRepo.CreateEmailLinkVer(userRecord.Email)
	if err != nil {
		log.Println(err)
		return err
	}
	return utils.SendEmail(constant.EMAIL_VER_SUB, emailLinkVer, userRecord.Email, c.appConfig.GetClientEmailSmtp(), c.appConfig.GetClientPasswordSmtp())

}

func (c *CloudUsecase) GetUserByEmail(email string) (*auth.UserRecord, error) {
	return c.cloudRepo.GetUserByEmail(email)
}

func (c *CloudUsecase) VerifyIDToken(idToken string) (*auth.Token, error) {
	return c.cloudRepo.VerifyIDToken(idToken)
}

func (c *CloudUsecase) GetUserByUid(uid string) (*auth.UserRecord, error) {
	return c.cloudRepo.GetUserByUid(uid)
}
func (c *CloudUsecase) CreateCustomToken(uid string) (string, error) {
	return c.cloudRepo.CreateCustomToken(uid)
}
func (c *CloudUsecase) VerifyCustomToken(token string) (jsonmodels.VerifyCustomTokenResBody, error) {
	return c.cloudRepo.VerifyCustomToken(token)
}
func (c *CloudUsecase) SendChangePasswordLink(email string) error {
	user, err := c.cloudRepo.GetUserByEmail(email)
	if err != nil {
		return err
	}
	link, err := c.cloudRepo.CreateChangePasswordLink(user)
	if err != nil {
		return err
	}
	return utils.SendEmail(constant.PASSWORD_CHANGE_SUB, link, user.Email, c.appConfig.GetClientEmailSmtp(), c.appConfig.GetClientPasswordSmtp())
}
func (c *CloudUsecase) GetNewToken(refreshToken string) (jsonmodels.RefreshTokenResBody, error) {
	return c.cloudRepo.CreateNewToken(refreshToken)

}
func (c *CloudUsecase) SendNotification(toEmail string, message string, data map[string]string) error {
	user, err := c.cloudRepo.GetUserByEmail(toEmail)
	if err != nil {
		return err
	}
	fcmToken, err := c.cloudRepo.GetFcmToken(user.UID)
	if err != nil {
		return err
	}
	c.cloudRepo.SendPushNotification(fcmToken, message, data)
	if err != nil {
		return err
	}
	return nil
}

func (c *CloudUsecase) DeleteFcmToken(uid string) error {
	return c.cloudRepo.DeleteFcmToken(uid)
}
