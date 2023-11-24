package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"my-orders/internal/repository"
	"my-orders/internal/util"
)

func PlanMiddleware(db repository.Querier) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		representativeID := ctx.Keys["representativeID"].(int32)
		userID := ctx.Keys["userID"].(int32)

		// Activity Info
		_, _ = db.CreateActivity(ctx, repository.CreateActivityParams{
			UserID:           userID,
			RepresentativeID: representativeID,
			Action:           ctx.Request.Method,
			ReferenceUrl:     ctx.Request.URL.Path,
		})

		exp, err := db.GetRepresentativeDateExpByID(ctx, representativeID)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusOK, util.ErrorResponse(402, "", util.ErrorPlanExpired.Error()))
			return
		}

		if time.Now().After(exp) {
			ctx.AbortWithStatusJSON(http.StatusOK, util.ErrorResponse(402, "", util.ErrorPlanExpired.Error()))
			return
		}
		ctx.Next()
	}
}
