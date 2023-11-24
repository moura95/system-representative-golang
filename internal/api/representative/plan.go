package representative

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"my-orders/internal/repository"
	"my-orders/internal/util"
)

type planResponse struct {
	CurrentPlan   repository.PlanTypes `json:"current_plan"`
	PlanExpiresAt string               `json:"plan_expires_at"`
	CreateAt      string               `json:"create_at"`
	TotalUsers    int64                `json:"total_users"`
}

func (r *Representative) getPlan(ctx *gin.Context) {
	representativeID := ctx.Keys["representativeID"].(int32)

	rep, err := r.Db.GetRepresentativesByID(ctx, representativeID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseRead.Error()))
			return
		}
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseRead.Error()))
		return
	}

	users, err := r.Db.GetTotalUsersByRepresentativesID(ctx, representativeID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseRead.Error()))
			return
		}
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseRead.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, planResponse{
		CurrentPlan:   rep.Plan,
		PlanExpiresAt: rep.DataExpire.Format("02/01/2006"),
		CreateAt:      rep.CreatedAt.Format("02/01/2006"),
		TotalUsers:    users,
	}, ""))
}
