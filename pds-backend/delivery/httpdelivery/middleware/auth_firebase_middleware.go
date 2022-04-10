package middleware

import (
	"net/http"
	"pds-backend/constant"
	"pds-backend/delivery/httpdelivery/httpconstant"
	"pds-backend/delivery/httpdelivery/httperror"
	"pds-backend/delivery/httpdelivery/httpresponse"
	"pds-backend/usecase"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthFirebaseMiddlewareEntity interface {
	RequireFirebaseToken() gin.HandlerFunc
}

type AuthFirebaseMiddleware struct {
	cloudUsecase usecase.CloudUsecaseEntity
}

func NewAuthFirebaseMiddleware(cloudUsecase usecase.CloudUsecaseEntity) AuthFirebaseMiddlewareEntity {
	return &AuthFirebaseMiddleware{cloudUsecase: cloudUsecase}
}

func (a *AuthFirebaseMiddleware) RequireFirebaseToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == httpconstant.ROUTE_NO_AUTH_LOGIN ||
			c.Request.URL.Path == httpconstant.ROUTE_NO_AUTH_REGISTER ||
			c.Request.URL.Path == httpconstant.ROUTE_DIVISION ||
			c.Request.URL.Path == httpconstant.ROUTE_ROLE ||
			c.Request.URL.Path == httpconstant.ROUTE_NO_AUTH_REFRESH_TOKEN {
			c.Next()
		} else if c.Request.URL.Path == httpconstant.ROUTE_NO_AUTH_CHANGE_PASSWORD && c.Request.Method == http.MethodPost {
			c.Next()
		} else {
			response := httpresponse.NewHttpResponse(c)
			authorizationToken := authHeader{}
			if err := c.ShouldBindHeader(&authorizationToken); err != nil {
				response.SendError(httpresponse.NewUnauthorizedMessage(httperror.ErrUnauthorized.Error(), nil))
				return
			}
			token := strings.TrimSpace(strings.Replace(authorizationToken.AuthorizationHeader, constant.BEARER, "", 1))
			if token == "" {
				response.SendError(httpresponse.NewUnauthorizedMessage(httperror.ErrUnauthorized.Error(), nil))
				c.Abort()
				return
			}

			customerTokenVerified, err := a.cloudUsecase.VerifyCustomToken(token)
			if err != nil {
				response.SendError(httpresponse.NewUnauthorizedMessage(httperror.ErrUnauthorized.Error(), nil))
				c.Abort()
				return
			}
			idTokenVerified, err := a.cloudUsecase.VerifyIDToken(customerTokenVerified.IdToken)
			if err != nil {
				response.SendError(httpresponse.NewInternalServerErrorMessage(httperror.ErrFirebase.Error(), nil))
				c.Abort()
				return
			}
			authRecord, err := a.cloudUsecase.GetUserByUid(idTokenVerified.UID)
			if err != nil {
				response.SendError(httpresponse.NewInternalServerErrorMessage(httperror.ErrFirebase.Error(), nil))
				c.Abort()
				return
			}

			c.Set(constant.UUID, authRecord.UID)
			c.Set(constant.EMAIL, authRecord.Email)
			c.Next()
		}

	}

}
