package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"pds-backend/apperror"
	"pds-backend/config"
	"pds-backend/constant"
	"pds-backend/delivery/httpdelivery/httpconstant"
	"pds-backend/delivery/httpdelivery/httperror"
	"pds-backend/delivery/httpdelivery/httpresponse"
	"pds-backend/delivery/httpdelivery/jsonmodels"
	"pds-backend/orm/gorm/model"
	"pds-backend/service"
	"pds-backend/usecase"
	"pds-backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthRouteEntity interface {
}

type AuthRoute struct {
	authUsecase  usecase.AuthUsecaseEntity
	cloudUsecase usecase.CloudUsecaseEntity
	userUsecase  usecase.UserUsecaseEntity
	appConfig    config.AppConfigEntity
}

func NewAuthRoute(
	apiRoute *gin.RouterGroup,
	authUsecase usecase.AuthUsecaseEntity,
	cloudUsecase usecase.CloudUsecaseEntity,
	userUsecase usecase.UserUsecaseEntity,
	appConfig config.AppConfigEntity,
) {
	authRoute := AuthRoute{authUsecase: authUsecase, cloudUsecase: cloudUsecase, userUsecase: userUsecase, appConfig: appConfig}
	apiRoute.POST(httpconstant.ROUTE_LOGIN, authRoute.authLoginHandlerFirebase)
	apiRoute.POST(httpconstant.ROUTE_LOGOUT, authRoute.authLogoutHandler)
	apiRoute.POST(httpconstant.ROUTE_REGISTER, authRoute.authRegisterHandlerFirebase)
	apiRoute.POST(httpconstant.ROUTE_REFRESH_TOKEN, authRoute.authRefreshToken)
	apiRoute.DELETE(httpconstant.ROUTE_FCM_TOKEN, authRoute.authDeleteFcmToken)
}

func (a *AuthRoute) authLoginHandler(c *gin.Context) {
	var loginRequestBody jsonmodels.LoginRequestBody
	var credential service.Credential
	var loginResponseBody jsonmodels.LoginResponseBody
	var picture string
	var bio string
	response := httpresponse.NewHttpResponse(c)
	if err := c.ShouldBindJSON(&loginRequestBody); err != nil {
		response.SendError(httpresponse.NewBadRequestMessage(apperror.ErrRequiredField.Error(), err.Error()))
		return
	}

	credential.Email = loginRequestBody.Email
	credential.Password = loginRequestBody.Password

	userData, tokenDetails, err := a.authUsecase.Login(credential)
	if err != nil {
		response.SendError(httpresponse.NewUnauthorizedMessage(httpresponse.FAILED_LOGIN, nil))
		return
	}

	if userData.Picture != nil || userData.Biography != nil {
		picture = *userData.Picture
		bio = *userData.Biography
	} else {
		picture = ""
		bio = ""
	}
	getRole, err := a.authUsecase.GetRoleName(*userData.RoleID)
	if err != nil {
		response.SendError(httpresponse.NewBadRequestMessage(apperror.ErrRequiredField.Error(), err.Error()))
		return
	}
	getDivision, err := a.authUsecase.GetDivisionName(*userData.DivisionID)
	if err != nil {
		response.SendError(httpresponse.NewBadRequestMessage(apperror.ErrRequiredField.Error(), err.Error()))
		return
	}

	loginResponseBody.TokenData = *tokenDetails
	loginResponseBody.UserData = jsonmodels.LoginProfileResponse{
		Fullname:   *userData.FullName,
		Email:      *userData.Email,
		Password:   *userData.Password,
		Picture:    picture,
		Gender:     *userData.Gender,
		DivisionID: *userData.DivisionID,
		Division:   *getDivision,
		RoleID:     *userData.RoleID,
		Role:       *getRole,
		Bio:        bio,
	}
	response.SendData(httpresponse.NewSuccessMessage(httpresponse.SUCCESS_LOGIN, loginResponseBody))
}
func (a *AuthRoute) authLoginHandlerFirebase(c *gin.Context) {
	var loginRequestBody jsonmodels.LoginRequestBody
	var loginResponseBody jsonmodels.LoginResponseBodyFirebase
	var tokenDetailsFirebase service.TokenDetailsFirebase

	response := httpresponse.NewHttpResponse(c)
	if err := c.ShouldBindJSON(&loginRequestBody); err != nil {
		response.SendError(httpresponse.NewBadRequestMessage(apperror.ErrRequiredField.Error(), err.Error()))
		return
	}

	userAuth, err := a.cloudUsecase.GetUserByEmail(loginRequestBody.Email)
	if err != nil {
		response.SendError(httpresponse.NewUnauthorizedMessage(httpresponse.WRONG_EMAIL_OR_PASSWORD, nil))
		return
	}
	if !userAuth.EmailVerified {
		response.SendError(httpresponse.NewUnauthorizedMessage(httpresponse.FAILED_EMAIL_NOT_VERIFIED, nil))
		return
	}
	loginRequestBody.ReturnSecureToken = true
	req, err := json.Marshal(loginRequestBody)
	if err != nil {
		response.SendError(httpresponse.NewInternalServerErrorMessage(http.StatusText(http.StatusInternalServerError), nil))
		return
	}
	resp, err := utils.PostRequest(fmt.Sprintf(constant.GOOGLE_SIGN_IN_PASSWORD, a.appConfig.GetFirebaseWebApiKey()), req)
	if err != nil {
		response.SendError(httpresponse.NewUnauthorizedMessage(httpresponse.WRONG_EMAIL_OR_PASSWORD, nil))
		return
	}
	respBody := struct {
		RefreshToken string `json:"refreshToken"`
	}{}
	err = json.Unmarshal(resp, &respBody)
	if err != nil {
		response.SendError(httpresponse.NewInternalServerErrorMessage(httperror.ErrFirebase.Error(), err.Error()))
		return
	}
	accessToken, err := a.cloudUsecase.CreateCustomToken(userAuth.UID)
	if err != nil {
		response.SendError(httpresponse.NewInternalServerErrorMessage(httperror.ErrFirebase.Error(), err.Error()))
		return
	}
	tokenDetailsFirebase.AccessToken = accessToken
	tokenDetailsFirebase.RefreshToken = respBody.RefreshToken

	getProfile, getDivision, getRole, err := a.userUsecase.GetProfileUser(loginRequestBody.Email)
	if err != nil {
		response.SendError(httpresponse.NewBadRequestMessage(apperror.ErrRequiredField.Error(), err.Error()))
		return
	}

	picture, bio := utils.HandlerPictureAndBiography(getProfile.Picture, getProfile.Biography)

	loginResponseBody.TokenData = tokenDetailsFirebase
	loginResponseBody.UserData = jsonmodels.LoginProfileResponse{
		Fullname:    *getProfile.FullName,
		Email:       *getProfile.Email,
		Password:    *getProfile.Password,
		Picture:     picture,
		Gender:      *getProfile.Gender,
		DivisionID:  *getProfile.DivisionID,
		Division:    *getDivision,
		RoleID:      *getProfile.RoleID,
		Role:        *getRole,
		Bio:         bio,
		FirebaseUID: userAuth.UID,
	}
	response.SendData(httpresponse.NewSuccessMessage(httpresponse.SUCCESS_LOGIN, loginResponseBody))
}
func (a *AuthRoute) authLogoutHandler(c *gin.Context) {
	accessUuid := c.GetString(constant.UUID)
	response := httpresponse.NewHttpResponse(c)
	err := a.authUsecase.Logout(accessUuid)
	if err != nil {
		response.SendError(httpresponse.NewBadRequestMessage(httperror.ErrUnauthorized.Error(), nil))
		return
	}
	response.SendData(httpresponse.NewSuccessMessage(httpresponse.SUCCESS_LOGOUT, nil))
}

func (a *AuthRoute) authRegisterHandlerFirebase(c *gin.Context) {
	var user model.User
	response := httpresponse.NewHttpResponse(c)
	if err := c.ShouldBindJSON(&user); err != nil {

		log.Println(err)
		response.SendError(httpresponse.NewBadRequestMessage(apperror.ErrRequiredField.Error(), err.Error()))
		return
	}

	err := a.cloudUsecase.RegisterUser(user)
	if err != nil {
		response.SendError(httpresponse.NewDuplicateValueMessage(apperror.ErrDuplicateEmail.Error(), err.Error()))
		return
	}
	registeredUser, err := a.authUsecase.Register(user)
	if err != nil {
		if strings.Contains(err.Error(), apperror.ErrDuplicateEmail.Error()) {
			response.SendError(httpresponse.NewDuplicateValueMessage(apperror.ErrDuplicateEmail.Error(), err.Error()))
			return
		}
		response.SendError(httpresponse.NewInternalServerErrorMessage(httpconstant.INTERNAL_SERVER_ERROR, err.Error()))
		return
	}
	response.SendData(httpresponse.NewSuccessMessage(httpconstant.SUCCESS, registeredUser))
}

func (a *AuthRoute) authRegisterHandler(c *gin.Context) {
	var user model.User
	response := httpresponse.NewHttpResponse(c)
	if err := c.ShouldBindJSON(&user); err != nil {
		response.SendError(httpresponse.NewBadRequestMessage(apperror.ErrRequiredField.Error(), err.Error()))
		return
	}

	registeredUser, err := a.authUsecase.Register(user)
	if err != nil {
		if strings.Contains(err.Error(), apperror.ErrDuplicateValue.Error()) {
			response.SendError(httpresponse.NewDuplicateValueMessage(apperror.ErrDuplicateValue.Error(), err.Error()))
			return
		}
		response.SendError(httpresponse.NewInternalServerErrorMessage(httpconstant.INTERNAL_SERVER_ERROR, err.Error()))
		return
	}
	response.SendData(httpresponse.NewSuccessMessage(httpconstant.SUCCESS, registeredUser))
}

func (a *AuthRoute) authRefreshToken(c *gin.Context) {
	response := httpresponse.NewHttpResponse(c)
	refreshTokenBody := struct {
		RefreshToken string `json:"refreshToken" binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&refreshTokenBody); err != nil {
		response.SendError(httpresponse.NewUnauthorizedMessage(apperror.ErrRefreshToken.Error(), nil))
		return
	}
	newTokenData, err := a.cloudUsecase.GetNewToken(refreshTokenBody.RefreshToken)
	if err != nil {
		response.SendError(httpresponse.NewInternalServerErrorMessage(http.StatusText(http.StatusInternalServerError), nil))
		return
	}
	log.Println(newTokenData)
	if newTokenData.AccessToken == "" && newTokenData.NewRefreshToken == "" {

		response.SendError(httpresponse.NewUnauthorizedMessage(apperror.ErrRefreshToken.Error(), nil))
		return
	}
	response.SendData(httpresponse.NewSuccessMessage("ok", gin.H{
		"newRefreshToken": newTokenData.NewRefreshToken,
		"accessToken":     newTokenData.AccessToken,
	}))
}
func (a *AuthRoute) authDeleteFcmToken(c *gin.Context) {
	response := httpresponse.NewHttpResponse(c)
	accessUuid := c.GetString(constant.UUID)
	err := a.cloudUsecase.DeleteFcmToken(accessUuid)
	if err != nil {
		response.SendError(httpresponse.NewBadRequestMessage("fcm token not found", nil))
		return
	}
	response.SendError(httpresponse.NewSuccessMessage("success delete fcm token", nil))
}
