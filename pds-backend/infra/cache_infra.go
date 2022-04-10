package infra

import (
	cacheengine "pds-backend/cacheengine/redis"
	"pds-backend/config"
)

type CacheInfraEntity interface {
	GetRedisCacheEngine() cacheengine.RedisCacheEngineEntity
}

type CacheInfra struct {
	redisCacheEngine cacheengine.RedisCacheEngineEntity
}

func NewCacheInfra(appConfig config.AppConfigEntity) CacheInfraEntity {
	redisCacheEngine := cacheengine.NewRedisCacheEngine(appConfig)
	return &CacheInfra{redisCacheEngine: redisCacheEngine}
}

func (c *CacheInfra) GetRedisCacheEngine() cacheengine.RedisCacheEngineEntity {
	return c.redisCacheEngine
}
