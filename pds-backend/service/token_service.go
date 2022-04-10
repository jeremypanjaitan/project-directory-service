package service

import (
	"context"
	cacheengine "pds-backend/cacheengine/redis"
	"pds-backend/config"
	"pds-backend/constant"
	"pds-backend/service/serviceerror"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type TokenDetails struct {
	AccessToken string `json:"accessToken"`
	AccessUuid  string `json:"accessUuid"`
	AtExpires   int64  `json:"atExpires"`
}

type TokenDetailsFirebase struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type TokenConfig struct {
	ApplicationName     string
	JwtSignatureKey     string
	JwtSigningMethod    *jwt.SigningMethodHMAC
	AccessTokenLifeTime time.Duration
	Client              *redis.Client
}

type JwtClaims struct {
	jwt.StandardClaims
	Email      string `json:"Email"`
	AccessUUID string `json:"AccessUUID"`
}

type UserCredential struct {
	AccessUuid string
	Email      string
}

type Credential struct {
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type TokenServiceEntity interface {
	CreateAccessToken(credential *Credential) (*TokenDetails, error)
	VerifyAccessToken(tokenString string) (*UserCredential, error)
	StoreAccessToken(userName string, tokenDetails *TokenDetails) error
	FetchAccessToken(userCredential *UserCredential) (string, error)
	DeleteAccessToken(accessUuid string) error
}

type TokenService struct {
	tokenConfig TokenConfig
}

func NewTokenService(appConfig config.AppConfigEntity, redisCacheEngine cacheengine.RedisCacheEngineEntity) TokenServiceEntity {
	tokenConfig := TokenConfig{
		ApplicationName:     appConfig.GetAppName(),
		JwtSigningMethod:    jwt.SigningMethodHS256,
		JwtSignatureKey:     appConfig.GetJwtSignature(),
		AccessTokenLifeTime: time.Duration(appConfig.GetTokenLifeTime()) * time.Second,
		Client:              redisCacheEngine.GetRedisClient(),
	}
	return &TokenService{
		tokenConfig: tokenConfig,
	}
}
func (t *TokenService) CreateAccessToken(credential *Credential) (*TokenDetails, error) {
	td := &TokenDetails{}
	now := time.Now().UTC()
	end := now.Add(t.tokenConfig.AccessTokenLifeTime)

	td.AtExpires = end.Unix()
	td.AccessUuid = uuid.New().String()

	claims := JwtClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:   t.tokenConfig.ApplicationName,
			IssuedAt: time.Now().Unix(),
		},
		Email:      credential.Email,
		AccessUUID: td.AccessUuid,
	}
	claims.IssuedAt = now.Unix()
	claims.ExpiresAt = end.Unix()

	token := jwt.NewWithClaims(t.tokenConfig.JwtSigningMethod, claims)
	newToken, err := token.SignedString([]byte(t.tokenConfig.JwtSignatureKey))
	td.AccessToken = newToken
	return td, err
}

func (t *TokenService) VerifyAccessToken(tokenString string) (*UserCredential, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, serviceerror.ErrSigningMethod
		} else if method != t.tokenConfig.JwtSigningMethod {
			return nil, serviceerror.ErrSigningMethod
		}
		return []byte(t.tokenConfig.JwtSignatureKey), nil
	})
	if err != nil {
		return nil, serviceerror.ErrInvalidNumberOfSegments
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}
	accessUUID := claims[constant.ACCESS_UUID].(string)
	email := claims[constant.EMAIL].(string)
	return &UserCredential{
		AccessUuid: accessUUID,
		Email:      email,
	}, nil

}

func (t *TokenService) StoreAccessToken(email string, tokenDetails *TokenDetails) error {
	at := time.Unix(tokenDetails.AtExpires, 0)
	now := time.Now()
	err := t.tokenConfig.Client.Set(context.Background(), tokenDetails.AccessUuid, email, at.Sub(now)).Err()
	if err != nil {
		return err
	}
	return nil
}

func (t *TokenService) FetchAccessToken(userCredential *UserCredential) (string, error) {
	if userCredential != nil {
		email, err := t.tokenConfig.Client.Get(context.Background(), userCredential.AccessUuid).Result()
		if err != nil {
			return "", err
		}
		return email, nil
	} else {
		return "", serviceerror.ErrInvalidAccess
	}
}

func (t *TokenService) DeleteAccessToken(accessUuid string) error {
	if accessUuid != "" {
		rd := t.tokenConfig.Client.Del(context.Background(), accessUuid)
		return rd.Err()
	} else {
		return serviceerror.ErrTokenNotFound
	}
}
