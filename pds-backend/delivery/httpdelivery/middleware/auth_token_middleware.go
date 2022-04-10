package middleware

import (
	"pds-backend/constant"
	"pds-backend/delivery/httpdelivery/httpconstant"
	"pds-backend/delivery/httpdelivery/httperror"
	"pds-backend/delivery/httpdelivery/httpresponse"
	"pds-backend/service"
	"strings"

	"github.com/gin-gonic/gin"
)

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

type AuthTokenMiddlewareEntity interface {
	RequireToken() gin.HandlerFunc
}

type AuthTokenMiddleware struct {
	tokenService service.TokenServiceEntity
}

func NewTokenValidator(tokenService service.TokenServiceEntity) AuthTokenMiddlewareEntity {
	return &AuthTokenMiddleware{
		tokenService: tokenService,
	}
}

func (a *AuthTokenMiddleware) RequireToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := httpresponse.NewHttpResponse(c)
		if c.Request.URL.Path == httpconstant.ROUTE_NO_AUTH_LOGIN ||
			c.Request.URL.Path == httpconstant.ROUTE_NO_AUTH_REGISTER ||
			c.Request.URL.Path == httpconstant.ROUTE_DIVISION ||
			c.Request.URL.Path == httpconstant.ROUTE_ROLE {
			c.Next()
		} else {
			h := authHeader{}
			if err := c.ShouldBindHeader(&h); err != nil {
				response.SendError(httpresponse.NewUnauthorizedMessage(httperror.ErrUnauthorized.Error(), nil))
				return
			}

			tokenString := strings.Replace(h.AuthorizationHeader, constant.BEARER, "", -1)
			if tokenString == "" {
				response.SendError(httpresponse.NewUnauthorizedMessage(httperror.ErrUnauthorized.Error(), nil))
				return
			}
			token, err := a.tokenService.VerifyAccessToken(tokenString)
			if err != nil {
				response.SendError(httpresponse.NewUnauthorizedMessage(httperror.ErrUnauthorized.Error(), nil))
				return
			}

			email, err := a.tokenService.FetchAccessToken(token)
			if email == "" || err != nil {
				response.SendError(httpresponse.NewUnauthorizedMessage(httperror.ErrUnauthorized.Error(), nil))
				return
			}

			if token != nil {
				c.Set(constant.EMAIL, email)
				c.Set(constant.UUID, token.AccessUuid)
				c.Next()
			} else {
				response.SendError(httpresponse.NewBadRequestMessage(httperror.ErrUnauthorized.Error(), nil))
				return
			}
		}
	}

}
