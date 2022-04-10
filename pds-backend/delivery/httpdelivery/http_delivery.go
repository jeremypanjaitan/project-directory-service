package delivery

import (
	"pds-backend/config"
	"pds-backend/delivery/httpdelivery/httpconstant"
	"pds-backend/delivery/httpdelivery/middleware"
	"pds-backend/delivery/httpdelivery/routes"
	httpengine "pds-backend/httpengine/gin"
	"pds-backend/usecase"
)

type HttpDeliveryEntity interface {
	Run()
}

type HttpDelivery struct {
	httpEngine httpengine.GinHttpEngineEntity
}

func NewHttpDelivery(
	appConfig config.AppConfigEntity,
	authUsecase usecase.AuthUsecaseEntity,
	divisionUsecase usecase.DivisionUsecaseEntity,
	roleUsecase usecase.RoleUsecaseEntity,
	categoryUsecase usecase.CategoryUsecaseEntity,
	userUsecase usecase.UserUsecaseEntity,
	projectUsecase usecase.ProjectUsecaseEntity,
	likesUsecase usecase.LikesUsecaseEntity,
	viewsUsecase usecase.ViewsUsecaseEntity,
	commentUsecase usecase.CommentUsecaseEntity,
	tokenMiddleware middleware.AuthFirebaseMiddlewareEntity,
	cloudUsecase usecase.CloudUsecaseEntity,
) HttpDeliveryEntity {
	httpEngine := httpengine.NewGinHttpEngine(appConfig)
	httpEngine.GetGinEngine().Use(tokenMiddleware.RequireFirebaseToken())

	apiAuthRoute := httpEngine.GetGinEngine().Group(httpconstant.API_ROOT_ROUTE + httpconstant.ROUTE_AUTH)
	apiDivisionRoute := httpEngine.GetGinEngine().Group(httpconstant.API_ROOT_ROUTE + httpconstant.ROUTE_DIVISION)
	apiRoleRoute := httpEngine.GetGinEngine().Group(httpconstant.API_ROOT_ROUTE + httpconstant.ROUTE_ROLE)
	apiCategoryRoute := httpEngine.GetGinEngine().Group(httpconstant.API_ROOT_ROUTE + httpconstant.ROUTE_CATEGORY)
	apiUserRoute := httpEngine.GetGinEngine().Group(httpconstant.API_ROOT_ROUTE + httpconstant.ROUTE_USER_PROFILE)
	apiProjectRoute := httpEngine.GetGinEngine().Group(httpconstant.API_ROOT_ROUTE + httpconstant.ROUTE_PROJECT)
	apiLikesRoute := httpEngine.GetGinEngine().Group(httpconstant.API_ROOT_ROUTE + httpconstant.ROUTE_LIKES)
	apiViewsRoute := httpEngine.GetGinEngine().Group(httpconstant.API_ROOT_ROUTE + httpconstant.ROUTE_VIEWS)
	apiCommentRoute := httpEngine.GetGinEngine().Group(httpconstant.API_ROOT_ROUTE + httpconstant.ROUTE_COMMENT)

	routes.NewAuthRoute(apiAuthRoute, authUsecase, cloudUsecase, userUsecase, appConfig)
	routes.NewDivisionRoute(apiDivisionRoute, divisionUsecase)
	routes.NewRoleRoute(apiRoleRoute, roleUsecase)
	routes.NewCategoryRoute(apiCategoryRoute, categoryUsecase)
	routes.NewUserRoute(apiUserRoute, userUsecase, cloudUsecase)
	routes.NewProjectRoute(apiProjectRoute, projectUsecase)
	routes.NewLikesRoute(apiLikesRoute, likesUsecase, cloudUsecase, projectUsecase, userUsecase)
	routes.NewViewsRoute(apiViewsRoute, viewsUsecase)
	routes.NewCommentRoute(apiCommentRoute, commentUsecase, cloudUsecase, projectUsecase, userUsecase)

	return &HttpDelivery{
		httpEngine: httpEngine,
	}
}

func (h *HttpDelivery) Run() {
	h.httpEngine.Run()
}
