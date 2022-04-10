package cacheengine

import (
	"context"
	"fmt"
	"log"
	"pds-backend/config"
	"pds-backend/constant"

	"github.com/go-redis/redis/v8"
)

type RedisCacheEngineEntity interface {
	GetRedisClient() *redis.Client
}

type RedisCacheEngine struct {
	redisClient *redis.Client
}

func NewRedisCacheEngine(appConfig config.AppConfigEntity) RedisCacheEngineEntity {
	redisHost := appConfig.GetRedisHost()
	redisPort := appConfig.GetRedisPort()
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: appConfig.GetRedisPassword(),
		DB:       0,
	})
	if appConfig.GetDebug() == constant.YES {
		pong, err := redisClient.Ping(context.Background()).Result()
		log.Println(pong)
		if err != nil {
			log.Fatal(pong)
		}
	}

	return &RedisCacheEngine{redisClient: redisClient}
}

func (r *RedisCacheEngine) GetRedisClient() *redis.Client {
	return r.redisClient
}
