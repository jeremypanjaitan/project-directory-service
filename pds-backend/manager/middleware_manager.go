package manager

import (
	"pds-backend/delivery/httpdelivery/middleware"
	"pds-backend/service"
	"pds-backend/usecase"
)

type MiddlewareManagerEntity interface {
	GetTokenMiddleware() middleware.AuthTokenMiddlewareEntity
	GetTokenFirebaseMiddleware() middleware.AuthFirebaseMiddlewareEntity
}

type MiddlewareManager struct {
	tokenMiddleware         middleware.AuthTokenMiddlewareEntity
	tokenFirebaseMiddleware middleware.AuthFirebaseMiddlewareEntity
}

func NewMiddlewareManager(tokenService service.TokenServiceEntity, cloudUsecase usecase.CloudUsecaseEntity) MiddlewareManagerEntity {
	tokenMiddleware := middleware.NewTokenValidator(tokenService)
	tokenFirebaseMiddleware := middleware.NewAuthFirebaseMiddleware(cloudUsecase)
	return &MiddlewareManager{tokenMiddleware: tokenMiddleware, tokenFirebaseMiddleware: tokenFirebaseMiddleware}
}

func (m *MiddlewareManager) GetTokenMiddleware() middleware.AuthTokenMiddlewareEntity {
	return m.tokenMiddleware
}

func (m *MiddlewareManager) GetTokenFirebaseMiddleware() middleware.AuthFirebaseMiddlewareEntity {
	return m.tokenFirebaseMiddleware
}
