package httpengine

import (
	"fmt"
	"pds-backend/config"

	"github.com/gin-gonic/gin"
)

type GinHttpEngineEntity interface {
	Run()
	GetGinEngine() *gin.Engine
}

type GinHttpEngine struct {
	ginEngine  *gin.Engine
	apiBaseUrl string
}

func NewGinHttpEngine(appConfig config.AppConfigEntity) GinHttpEngineEntity {
	apiBaseUrl := fmt.Sprintf("%s:%s", appConfig.GetApiHost(), appConfig.GetApiPort())
	gin.SetMode(appConfig.GetGinMode())
	ginEngine := gin.Default()
	return &GinHttpEngine{
		ginEngine:  ginEngine,
		apiBaseUrl: apiBaseUrl,
	}
}

func (g *GinHttpEngine) Run() {
	g.ginEngine.Run(g.apiBaseUrl)
}

func (g *GinHttpEngine) GetGinEngine() *gin.Engine {
	return g.ginEngine
}
