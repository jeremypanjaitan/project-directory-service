package routes

import (
	"fmt"
	"log"
	"pds-backend/apperror"
	"pds-backend/constant"
	"pds-backend/delivery/httpdelivery/httpresponse"
	"pds-backend/delivery/httpdelivery/jsonmodels"
	"pds-backend/orm/gorm/model"
	"pds-backend/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentRouteEntity interface {
}

type CommentRoute struct {
	commentUsecase usecase.CommentUsecaseEntity
	cloudUsecase   usecase.CloudUsecaseEntity
	projectUsecase usecase.ProjectUsecaseEntity
	userUsecase    usecase.UserUsecaseEntity
}

func NewCommentRoute(apiRoute *gin.RouterGroup,
	commentUsecase usecase.CommentUsecaseEntity,
	cloudUsecase usecase.CloudUsecaseEntity,
	projectUsecase usecase.ProjectUsecaseEntity,
	userUsecase usecase.UserUsecaseEntity,
) {
	commentRoute := CommentRoute{commentUsecase: commentUsecase, cloudUsecase: cloudUsecase, projectUsecase: projectUsecase, userUsecase: userUsecase}
	apiRoute.GET("", commentRoute.CommentGetAllHandler)
	apiRoute.POST("", commentRoute.CommentCreateHandler)
	apiRoute.DELETE("", commentRoute.CommentDeleteHandler)
}

func (m *CommentRoute) CommentGetAllHandler(c *gin.Context) {
	response := httpresponse.NewHttpResponse(c)
	projectId, _ := strconv.Atoi(c.Param("id"))
	pageNumber := c.Query("pageNumber")
	pageSize := c.Query("pageSize")
	var convPageNumber int
	var convPageSize int

	if pageNumber == "" || pageSize == "" {
		response.SendError(httpresponse.NewBadRequestMessage(apperror.ErrRequiredField.Error(), nil))
		return
	} else {
		convPageNumber, _ = strconv.Atoi(pageNumber)
		convPageSize, _ = strconv.Atoi(pageSize)
	}

	getComments, getPagination, err := m.commentUsecase.GetAllComment(projectId, convPageNumber, convPageSize)

	if err != nil {
		emptyResponse := jsonmodels.ListComment{
			Row:         make([]model.CommentJointUser, 0),
			CurrentPage: 0,
			TotalPage:   0,
		}
		response.SendData(httpresponse.NewSuccessMessage(httpresponse.SUCCESS_GETCOMMENTLIST, emptyResponse))
		return
	}

	responseData := jsonmodels.ListComment{
		CurrentPage: uint(convPageNumber),
		TotalPage:   *getPagination,
	}

	if len(getComments) < 1 {
		responseData.Row = make([]model.CommentJointUser, 0)
	} else {
		responseData.Row = getComments
	}

	response.SendData(httpresponse.NewSuccessMessage(httpresponse.SUCCESS_GETCOMMENTLIST, responseData))
}

func (m *CommentRoute) CommentCreateHandler(c *gin.Context) {
	var commentRequestBody jsonmodels.CommentRequestBody
	response := httpresponse.NewHttpResponse(c)
	projectId, _ := strconv.Atoi(c.Param("id"))
	userEmail := c.GetString(constant.EMAIL)

	if err := c.ShouldBindJSON(&commentRequestBody); err != nil {
		response.SendError(httpresponse.NewBadRequestMessage(httpresponse.FAILED_CREATECOMMENT, nil))
		return
	}

	if len(commentRequestBody.Body) < 2 || len(commentRequestBody.Body) > 300 {
		response.SendError(httpresponse.NewBadRequestMessage(httpresponse.FAILED_CREATECOMMENT, nil))
		return
	}

	getComment, getUserPhoto, err := m.commentUsecase.CreateComment(projectId, userEmail, commentRequestBody.Body)
	if err != nil {
		fmt.Println("error creating comment")
		response.SendError(httpresponse.NewBadRequestMessage(httpresponse.FAILED_CREATECOMMENT, nil))
		return
	}
	pOwner, err := m.projectUsecase.FindProjectOwner(uint(projectId))
	if err != nil {
		log.Println(err)
	}

	if *pOwner.Email != userEmail {
		project, err := m.projectUsecase.GetProjectDetails(uint(projectId))
		if err != nil {
			log.Println(err)
		}
		user, _, _, err := m.userUsecase.GetProfileUser(userEmail)
		if err != nil {
			log.Println(err)
		}
		message := constant.NotificationMessageConstant(constant.COMMENTED, *project.Title, *user.FullName)
		_ = m.cloudUsecase.SendNotification(*pOwner.Email, message, map[string]string{"projectId": c.Param("id"), "projectPicture": *project.Picture})
	}

	responseData := jsonmodels.CommentResponseBody{
		UserPhoto: getUserPhoto,
		Body:      getComment.Body,
		CreatedAt: getComment.CreatedAt,
	}

	response.SendData(httpresponse.NewSuccessMessage(httpresponse.SUCCESS_CREATECOMMENT, responseData))
}

func (m *CommentRoute) CommentDeleteHandler(c *gin.Context) {
	response := httpresponse.NewHttpResponse(c)
	commentId, _ := strconv.Atoi(c.Param("id"))

	_ = m.commentUsecase.DeleteComment(commentId)

	response.SendData(httpresponse.NewSuccessMessage(httpresponse.SUCCESS_DELETECOMMENT, nil))
}
