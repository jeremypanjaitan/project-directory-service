package manager

import (
	"pds-backend/config"
	delivery "pds-backend/delivery/httpdelivery"
)

type DeliveryManagerEntity interface {
	GetHttpDelivery() delivery.HttpDeliveryEntity
}
type DeliveryManager struct {
	httpDelivery delivery.HttpDeliveryEntity
}

func NewDeliveryManager(
	appConfig config.AppConfigEntity,
	usecaseManager UsecaseManagerEntity,
	middlewareManager MiddlewareManagerEntity,
) DeliveryManagerEntity {
	httpDelivery := delivery.NewHttpDelivery(appConfig,
		usecaseManager.GetAuthUseCase(),
		usecaseManager.GetDivisionUseCase(),
		usecaseManager.GetRoleUseCase(),
		usecaseManager.GetCategoryUseCase(),
		usecaseManager.GetUserUsecase(),
		usecaseManager.GetProjectUsecase(),
		usecaseManager.GetLikesUsecase(),
		usecaseManager.GetViewsUsecase(),
		usecaseManager.GetCommentUsecase(),
		middlewareManager.GetTokenFirebaseMiddleware(),
		usecaseManager.GetCloudUsecase(),
	)
	return &DeliveryManager{httpDelivery: httpDelivery}
}

func (d *DeliveryManager) GetHttpDelivery() delivery.HttpDeliveryEntity {
	return d.httpDelivery
}
