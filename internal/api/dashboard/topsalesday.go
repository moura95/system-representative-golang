package dashboard

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"my-orders/internal/repository"
	"my-orders/internal/util"
)

type topSalesRequest struct {
	StartDate time.Time `form:"start_date" binding:"required" time_format:"2006-01-02"`
	EndDate   time.Time `form:"end_date" binding:"required" time_format:"2006-01-02"`
}

type topSalesResponse struct {
	Day      []string `json:"day"`
	TotalDay []string `json:"total_day"`
	Total    string   `json:"total"`
}

func (d *Dashboard) topSales(ctx *gin.Context) {
	representativeID := ctx.Keys["representativeID"].(int32)

	var req topSalesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	listTop, err := d.Db.TotalSalesPerDayByRepresentativeID(ctx, repository.TotalSalesPerDayByRepresentativeIDParams{
		RepresentativeID: representativeID,
		StartDate:        req.StartDate,
		EndDate:          req.EndDate,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusOK, util.SuccessResponse(200, listTop, ""))
			return
		}
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseRead.Error()))
		return
	}

	if len(listTop) == 0 {
		ctx.JSON(http.StatusOK, util.SuccessResponse(200, listTop, ""))
		return
	}

	var response topSalesResponse
	for _, row := range listTop {
		response.Day = append(response.Day, row.Day)
		response.TotalDay = append(response.TotalDay, row.TotalDay)
		response.Total = row.Total
	}
	ctx.JSON(http.StatusOK, util.SuccessResponse(200, response, ""))
	return
}
