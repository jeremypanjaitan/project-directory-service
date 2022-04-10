package manager

import (
	cacheengine "pds-backend/cacheengine/redis"
	"pds-backend/config"
	"pds-backend/service"
)

type ServiceManagerEntity interface {
	GetTokenService() service.TokenServiceEntity
}

type ServiceManager struct {
	tokenService service.TokenServiceEntity
}

func NewServiceManager(appConfig config.AppConfigEntity, redisCacheEngine cacheengine.RedisCacheEngineEntity) ServiceManagerEntity {
	tokenService := service.NewTokenService(appConfig, redisCacheEngine)
	return &ServiceManager{tokenService: tokenService}
}

func (s *ServiceManager) GetTokenService() service.TokenServiceEntity {
	return s.tokenService
}
