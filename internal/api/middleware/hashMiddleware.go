package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"

	"my-orders/cfg"
	"my-orders/internal/util"
)

type HashRequest struct {
	Signature string `json:"signature" description:"Signature"`
	TimeStamp string `json:"timestamp" description:"TimeStamp"`
}

func HashMiddleware(cfg cfg.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req HashRequest
		if err := ctx.ShouldBindHeader(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
			return
		}

		h := hmac.New(sha256.New, []byte(cfg.SecretKey))
		h.Write([]byte("timestamp=" + req.TimeStamp))
		signature := hex.EncodeToString(h.Sum(nil))

		if signature != req.Signature {
			ctx.AbortWithStatusJSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
			return
		}

		ctx.Next()
		return
	}
}
