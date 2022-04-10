package app

import (
	"errors"
	"log"
	"os"
	"pds-backend/config"
	"pds-backend/manager"

	"github.com/joho/godotenv"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	//Check if .env file is exist or not
	if _, err := os.Stat(".env"); !errors.Is(err, os.ErrNotExist) {
		godotenv.Load()
	}
}

type AppEntity interface {
	Run()
}

type App struct {
	deliveryManager manager.DeliveryManagerEntity
}

func NewApp() AppEntity {
	appConfig := config.NewAppConfig()

	infraManager := manager.NewInfraManager(appConfig)
	repoManager := manager.NewRepoManager(infraManager, appConfig)
	serviceManager := manager.NewServiceManager(appConfig, infraManager.GetCacheInfra().GetRedisCacheEngine())
	usecaseManager := manager.NewUsecaseManager(repoManager, serviceManager, appConfig)
	middlewareManager := manager.NewMiddlewareManager(serviceManager.GetTokenService(), usecaseManager.GetCloudUsecase())
	deliveryManager := manager.NewDeliveryManager(appConfig, usecaseManager, middlewareManager)

	app := new(App)
	app.deliveryManager = deliveryManager

	return app
}

func (a *App) Run() {
	a.deliveryManager.GetHttpDelivery().Run()
}
