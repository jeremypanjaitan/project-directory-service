package config

import (
	"fmt"
	"os"
	"strconv"
)

type AppConfigEntity interface {
	GetDataSourceName() string
	GetRunMigration() string
	GetDebug() string
	GetApiHost() string
	GetApiPort() string
	GetAppName() string
	GetJwtSignature() string
	GetRedisHost() string
	GetRedisPort() string
	GetTokenLifeTime() int
	GetDbSeed() string
	GetDbSchema() string
	GetGinMode() string
	GetRedisPassword() string
	GetFirebaseWebApiKey() string
	GetClientEmailSmtp() string
	GetClientPasswordSmtp() string
}
type AppConfig struct {
	dataSourceName     string
	runMigration       string
	debug              string
	apiHost            string
	apiPort            string
	appName            string
	jwtSignature       string
	redisHost          string
	redisPort          string
	tokenLifeTime      int
	dbSeed             string
	dbSchema           string
	ginMode            string
	redisPasword       string
	firebaseWebApiKey  string
	clientEmailSmtp    string
	clientPasswordSmtp string
}

func NewAppConfig() AppConfigEntity {
	appConfig := AppConfig{}

	runMigration := os.Getenv("DB_MIGRATION")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbSchema := os.Getenv("DB_SCHEMA")
	debug := os.Getenv("DEBUG")
	apiHost := os.Getenv("API_HOST")
	apiPort := os.Getenv("API_PORT")
	appName := os.Getenv("APP_NAME")
	jwtSignature := os.Getenv("JWT_SIGNATURE")
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	tokenLifeTime, _ := strconv.Atoi(os.Getenv("TOKEN_LIFE_TIME"))
	dbSeed := os.Getenv("DB_SEED")
	ginMode := os.Getenv("GIN_MODE")
	firebaseWebApiKey := os.Getenv("FIREBASE_WEB_API_KEY")
	clientEmailSmtp := os.Getenv("CLIENT_EMAIL_SMTP")
	clientPasswordSmtp := os.Getenv("CLIENT_PASSWORD_SMTP")

	dataSourceName := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPassword, dbName, dbPort)

	appConfig.dataSourceName = dataSourceName
	appConfig.runMigration = runMigration
	appConfig.debug = debug
	appConfig.apiHost = apiHost
	appConfig.apiPort = apiPort
	appConfig.appName = appName
	appConfig.jwtSignature = jwtSignature
	appConfig.redisHost = redisHost
	appConfig.redisPort = redisPort
	appConfig.tokenLifeTime = tokenLifeTime
	appConfig.dbSeed = dbSeed
	appConfig.dbSchema = dbSchema
	appConfig.ginMode = ginMode
	appConfig.redisPasword = redisPassword
	appConfig.firebaseWebApiKey = firebaseWebApiKey
	appConfig.clientEmailSmtp = clientEmailSmtp
	appConfig.clientPasswordSmtp = clientPasswordSmtp
	return &appConfig
}

func (a *AppConfig) GetDataSourceName() string {
	return a.dataSourceName
}

func (a *AppConfig) GetRunMigration() string {
	return a.runMigration
}

func (a *AppConfig) GetDebug() string {
	return a.debug
}

func (a *AppConfig) GetApiHost() string {
	return a.apiHost
}

func (a *AppConfig) GetApiPort() string {
	return a.apiPort
}

func (a *AppConfig) GetAppName() string {
	return a.appName
}

func (a *AppConfig) GetJwtSignature() string {
	return a.jwtSignature
}

func (a *AppConfig) GetRedisHost() string {
	return a.redisHost
}
func (a *AppConfig) GetRedisPort() string {
	return a.redisPort
}

func (a *AppConfig) GetTokenLifeTime() int {
	return a.tokenLifeTime
}

func (a *AppConfig) GetDbSeed() string {
	return a.dbSeed
}

func (a *AppConfig) GetDbSchema() string {
	return a.dbSchema
}
func (a *AppConfig) GetGinMode() string {
	return a.ginMode
}
func (a *AppConfig) GetRedisPassword() string {
	return a.redisPasword
}
func (a *AppConfig) GetFirebaseWebApiKey() string {
	return a.firebaseWebApiKey
}
func (a *AppConfig) GetClientEmailSmtp() string {
	return a.clientEmailSmtp
}
func (a *AppConfig) GetClientPasswordSmtp() string {
	return a.clientPasswordSmtp
}
