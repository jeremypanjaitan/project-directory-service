package routes

import (
	"log"
	"pds-backend/apperror"
	"pds-backend/constant"
	"pds-backend/delivery/httpdelivery/httpconstant"
	"pds-backend/delivery/httpdelivery/httpresponse"
	"pds-backend/delivery/httpdelivery/jsonmodels"
	"pds-backend/orm/gorm/model"
	"pds-backend/usecase"
	"pds-backend/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserRouteEntity interface {
}

type UserRoute struct {
	userUsecase  usecase.UserUsecaseEntity
	cloudUsecase usecase.CloudUsecaseEntity
}

func NewUserRoute(apiRoute *gin.RouterGroup,
	userUsecase usecase.UserUsecaseEntity,
	cloudUsecase usecase.CloudUsecaseEntity) {
	userRoute := UserRoute{userUsecase: userUsecase, cloudUsecase: cloudUsecase}
	apiRoute.GET("", userRoute.getProfileHandler)
	apiRoute.GET(httpconstant.ROUTE_USER_PROFILE_ID, userRoute.getProfileByIdHandler)
	apiRoute.PUT(httpconstant.ROUTE_CHANGE_PASSWORD, userRoute.changePasswordHandler)
	apiRoute.POST(httpconstant.ROUTE_CHANGE_PASSWORD, userRoute.changePasswordFirebaseHandler)
	apiRoute.PUT("", userRoute.changeProfileHandler)
	apiRoute.GET(httpconstant.ROUTE_USER_PROFILE_ACTIVITY, userRoute.getActivityHandler)
	apiRoute.GET(httpconstant.ROUTE_PROJECT, userRoute.getProjectHandler)
}

func (u *UserRoute) getProfileHandler(c *gin.Context) {
	response := httpresponse.NewHttpResponse(c)
	userEmail := c.GetString(constant.EMAIL)
	firebaseUid := c.GetString(constant.UUID)
	getProfile, getDivision, getRole, err := u.userUsecase.GetProfileUser(userEmail)
	if err != nil {
		response.SendError(httpresponse.NewBadRequestMessage(apperror.ErrRequiredField.Error(), err.Error()))
		return
	}
	picture, bio := utils.HandlerPictureAndBiography(getProfile.Picture, getProfile.Biography)

	profile := jsonmodels.UserProfileResponse{
		Fullname:    *getProfile.FullName,
		Email:       *getProfile.Email,
		Picture:     picture,
		Gender:      *getProfile.Gender,
		DivisionID:  *getProfile.DivisionID,
		Division:    *getDivision,
		RoleID:      *getProfile.RoleID,
		Role:        *getRole,
		Bio:         bio,
		FirebaseUid: firebaseUid,
	}

	response.SendData(httpresponse.NewSuccessMessage(httpconstant.SUCCESS, profile))
}

func (u *UserRoute) getProfileByIdHandler(c *gin.Context) {
	response := httpresponse.NewHttpResponse(c)
	userId, _ := strconv.Atoi(c.Param("id"))

	getProfileById, getDivision, getRole, err := u.userUsecase.GetProfileUserById(uint(userId))
	if err != nil {
		response.SendError(httpresponse.NewBadRequestMessage(apperror.ErrRequiredField.Error(), err.Error()))
		return
	}

	picture, bio := utils.HandlerPictureAndBiography(getProfileById.Picture, getProfileById.Biography)

	profile := jsonmodels.UserProfileResponse{
		Fullname:   *getProfileById.FullName,
		Email:      *getProfileById.Email,
		Picture:    picture,
		Gender:     *getProfileById.Gender,
		DivisionID: *getProfileById.DivisionID,
		Division:   *getDivision,
		RoleID:     *getProfileById.RoleID,
		Role:       *getRole,
		Bio:        bio,
	}

	response.SendData(httpresponse.NewSuccessMessage(httpconstant.SUCCESS, profile))
}

func (u *UserRoute) changePasswordHandler(c *gin.Context) {
	var userRequestBody jsonmodels.UserChangePassword
	response := httpresponse.NewHttpResponse(c)
	userEmail := c.GetString(constant.EMAIL)

	if err := c.ShouldBindJSON(&userRequestBody); err != nil {
		response.SendError(httpresponse.NewBadRequestMessage(apperror.ErrRequiredField.Error(), err.Error()))
		log.Println(err)
		return
	}

	err := u.userUsecase.ChangePassword(userEmail, userRequestBody.OldPassword, userRequestBody.NewPassword)
	if err != nil {
		response.SendError(httpresponse.NewBadRequestMessage(apperror.ErrRequiredField.Error(), err.Error()))
		log.Println(err)
		return
	}

	response.SendData(httpresponse.ResponseMessage{Status: 200, Description: "successfully update password", Data: ""})
}

func (u *UserRoute) changePasswordFirebaseHandler(c *gin.Context) {
	response := httpresponse.NewHttpResponse(c)
	emailRequest := struct {
		Email string `json:"email"`
	}{}
	if err := c.BindJSON(&emailRequest); err != nil {
		response.SendError(httpresponse.NewBadRequestMessage(apperror.ErrRequiredField.Error(), err.Error()))
		log.Println(err)
		return
	}
	err := u.cloudUsecase.SendChangePasswordLink(emailRequest.Email)
	if err != nil {
		response.SendData(httpresponse.NewBadRequestMessage(apperror.ErrEmailAddress.Error(), gin.H{"message": err.Error()}))
		return
	}
	response.SendData(httpresponse.NewSuccessMessage("ok", gin.H{"message": "password change link has been sent to your email"}))
}

func (u *UserRoute) changeProfileHandler(c *gin.Context) {
	var userRequestBody jsonmodels.UserChangeProfileData
	response := httpresponse.NewHttpResponse(c)
	userEmail := c.GetString(constant.EMAIL)

	if err := c.ShouldBindJSON(&userRequestBody); err != nil {
		response.SendError(httpresponse.NewBadRequestMessage(apperror.ErrRequiredField.Error(), err.Error()))
		log.Println(err)
		return
	}

	updatedProfile := model.User{
		FullName:   &userRequestBody.Fullname,
		Picture:    &userRequestBody.Picture,
		Gender:     &userRequestBody.Gender,
		DivisionID: &userRequestBody.DivisionID,
		RoleID:     &userRequestBody.RoleID,
		Biography:  &userRequestBody.Bio,
	}

	if !utils.CheckProfileResponseStringLength(updatedProfile) {
		response.SendError(httpresponse.NewBadRequestMessage(apperror.ErrConstraintStringLength.Error(), ""))
		return
	}

	updatedProfileData, getDivision, getRole, err := u.userUsecase.ChangeProfile(userEmail, updatedProfile)
	if err != nil {
		response.SendError(httpresponse.NewBadRequestMessage(apperror.ErrRequiredField.Error(), err.Error()))
		log.Println(err)
		return
	}

	picture, bio := utils.HandlerPictureAndBiography(updatedProfileData.Picture, updatedProfile.Biography)
	profile := jsonmodels.UserProfileResponse{
		Fullname:   *updatedProfileData.FullName,
		Email:      userEmail,
		Picture:    picture,
		Gender:     *updatedProfileData.Gender,
		DivisionID: *updatedProfileData.DivisionID,
		Division:   *getDivision,
		RoleID:     *updatedProfileData.RoleID,
		Role:       *getRole,
		Bio:        bio,
	}
	response.SendData(httpresponse.NewSuccessMessage(constant.SUCCESS, profile))
}

func (u *UserRoute) getActivityHandler(c *gin.Context) {
	response := httpresponse.NewHttpResponse(c)
	pageNumber := c.Query("pageNumber")
	pageSize := c.Query("pageSize")
	userEmail := c.GetString(constant.EMAIL)

	convPageNumber, convPageSize := utils.CheckPageNumberAndPageSize(pageNumber, pageSize, c)

	activities, pagination, err := u.userUsecase.GetActivity(uint(convPageNumber), uint(convPageSize), userEmail)
	if err != nil {
		response.SendError(httpresponse.NewBadRequestMessage(apperror.ErrRequiredField.Error(), err.Error()))
		log.Println(err)
		return
	}

	responseBody := jsonmodels.ListActivity{
		Row:         activities,
		CurrentPage: uint(convPageNumber),
		TotalPage:   *pagination,
	}
	response.SendData(httpresponse.ResponseMessage{Status: 200, Description: "ok", Data: responseBody})
}

func (u *UserRoute) getProjectHandler(c *gin.Context) {
	response := httpresponse.NewHttpResponse(c)
	pageNumber := c.Query("pageNumber")
	pageSize := c.Query("pageSize")
	userEmail := c.GetString(constant.EMAIL)

	convPageNumber, convPageSize := utils.CheckPageNumberAndPageSize(pageNumber, pageSize, c)

	projects, pagination, err := u.userUsecase.GetProject(uint(convPageNumber), uint(convPageSize), userEmail)
	if err != nil {
		response.SendError(httpresponse.NewBadRequestMessage(apperror.ErrRequiredField.Error(), err.Error()))
		log.Println(err)
		return
	}

	rowsResponse := utils.RowsHandler(projects)
	rows := utils.PageNumberHandler(convPageNumber, *pagination, rowsResponse)

	responseBody := jsonmodels.ListProject{
		Row:         rows,
		CurrentPage: uint(convPageNumber),
		TotalPage:   *pagination,
	}

	response.SendData(httpresponse.ResponseMessage{Status: 200, Description: "ok", Data: responseBody})
}
