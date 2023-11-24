package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"my-orders/internal/token"
	"my-orders/internal/util"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
)

func AuthMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)

		if len(authorizationHeader) == 0 {
			ctx.AbortWithStatusJSON(http.StatusOK, util.ErrorResponse(401, "", util.ErrorAuthorizationHeaderNotProvided.Error()))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			_ = errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusOK, util.ErrorResponse(401, "", util.ErrorAuthorizationHeaderNotProvided.Error()))
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			_ = fmt.Errorf("unsupported authorization type %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusOK, util.ErrorResponse(401, "", util.ErrorAuthorizationHeaderNotProvided.Error()))
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusOK, util.ErrorResponse(401, "", util.ErrorAuthorizationHeaderNotProvided.Error()))
			return
		}

		ctx.Set("userID", payload.UserID)
		ctx.Set("representativeID", payload.RepresentativeID)
		ctx.Next()
	}
}
