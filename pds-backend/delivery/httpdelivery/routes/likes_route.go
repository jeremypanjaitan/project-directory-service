package routes

import (
	"log"
	"pds-backend/constant"
	"pds-backend/delivery/httpdelivery/httpconstant"
	"pds-backend/delivery/httpdelivery/httpresponse"
	"pds-backend/delivery/httpdelivery/jsonmodels"
	"pds-backend/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LikesRouteEntity interface{}

type LikesRoute struct {
	likesUsecase   usecase.LikesUsecaseEntity
	cloudUsecase   usecase.CloudUsecaseEntity
	projectUsecase usecase.ProjectUsecaseEntity
	userUsecase    usecase.UserUsecaseEntity
}

func NewLikesRoute(apiRoute *gin.RouterGroup,
	likesUsecase usecase.LikesUsecaseEntity,
	cloudUsecase usecase.CloudUsecaseEntity,
	projectUsecase usecase.ProjectUsecaseEntity,
	userUsecase usecase.UserUsecaseEntity,
) {
	likesRoute := LikesRoute{likesUsecase: likesUsecase, cloudUsecase: cloudUsecase, projectUsecase: projectUsecase, userUsecase: userUsecase}
	apiRoute.GET(httpconstant.ROUTE_LIKES_LIKE, likesRoute.LikesGetProjectLikeHandler)
	apiRoute.POST(httpconstant.ROUTE_LIKES_LIKE, likesRoute.LikesPostLikeHandler)
	apiRoute.DELETE(httpconstant.ROUTE_LIKES_DISLIKE, likesRoute.LikesDislikeHandler)
}

func (l *LikesRoute) LikesGetProjectLikeHandler(c *gin.Context) {
	response := httpresponse.NewHttpResponse(c)
	projectId, _ := strconv.Atoi(c.Param("id"))
	userEmail := c.GetString(constant.EMAIL)

	// count, err := l.likesUsecase.GetLikeCount(uint(projectId))
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Count: ", count)
	// }

	responseData := jsonmodels.LikesResponseBody{}

	totalLikes, isUserLike, err := l.likesUsecase.GetLikeByProjectId(uint(projectId), userEmail)
	if err != nil {
		responseData.TotalLikes = 0
		responseData.IsUserLike = false
		response.SendData(httpresponse.NewSuccessMessage(httpresponse.SUCCESS_GETPROJECTLIKE, responseData))
		return
	}

	responseData.TotalLikes = totalLikes
	responseData.IsUserLike = isUserLike
	response.SendData(httpresponse.NewSuccessMessage(httpresponse.SUCCESS_GETPROJECTLIKE, responseData))
}

func (l *LikesRoute) LikesPostLikeHandler(c *gin.Context) {
	response := httpresponse.NewHttpResponse(c)
	projectId, _ := strconv.Atoi(c.Param("id"))
	userEmail := c.GetString(constant.EMAIL)

	err := l.likesUsecase.CreateLike(uint(projectId), userEmail)
	if err != nil {
		response.SendError(httpresponse.NewBadRequestMessage(httpresponse.FAILED_CREATELIKE, nil))
		return
	}

	pOwner, err := l.projectUsecase.FindProjectOwner(uint(projectId))
	if err != nil {
		log.Println(err)
	}
	if *pOwner.Email != userEmail {
		project, err := l.projectUsecase.GetProjectDetails(uint(projectId))
		if err != nil {
			log.Println(err)
		}
		user, _, _, err := l.userUsecase.GetProfileUser(userEmail)
		if err != nil {
			log.Println(err)
		}
		message := constant.NotificationMessageConstant(constant.LIKE, *project.Title, *user.FullName)
		_ = l.cloudUsecase.SendNotification(*pOwner.Email,
			message,
			map[string]string{"projectId": c.Param("id"), "projectPicture": *project.Picture},
		)
	}

	response.SendData(httpresponse.NewSuccessMessage(httpresponse.SUCCESS_CREATELIKE, nil))
}

func (l *LikesRoute) LikesDislikeHandler(c *gin.Context) {
	response := httpresponse.NewHttpResponse(c)
	projectId, _ := strconv.Atoi(c.Param("id"))
	userEmail := c.GetString(constant.EMAIL)

	_ = l.likesUsecase.DeleteLike(uint(projectId), userEmail)
	response.SendData(httpresponse.NewSuccessMessage(httpresponse.SUCCESS_DISLIKE, nil))
}
