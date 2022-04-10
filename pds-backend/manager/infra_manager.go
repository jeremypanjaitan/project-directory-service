package manager

import (
	"pds-backend/config"
	"pds-backend/infra"
)

type InfraManagerEntity interface {
	GetOrmInfra() infra.OrmInfraEntity
	GetCacheInfra() infra.CacheInfraEntity
	GetCloudInfra() infra.CloudInfraEntity
}

type InfraManager struct {
	ormInfra   infra.OrmInfraEntity
	cacheInfra infra.CacheInfraEntity
	cloudInfra infra.CloudInfraEntity
}

func NewInfraManager(appConfig config.AppConfigEntity) InfraManagerEntity {
	ormInfra := infra.NewOrmInfra(appConfig)
	cacheInfra := infra.NewCacheInfra(appConfig)
	cloudInfra := infra.NewCloudInfra()
	return &InfraManager{ormInfra: ormInfra, cacheInfra: cacheInfra, cloudInfra: cloudInfra}
}

func (i *InfraManager) GetOrmInfra() infra.OrmInfraEntity {
	return i.ormInfra
}

func (i *InfraManager) GetCacheInfra() infra.CacheInfraEntity {
	return i.cacheInfra
}
func (i *InfraManager) GetCloudInfra() infra.CloudInfraEntity {
	return i.cloudInfra
}
