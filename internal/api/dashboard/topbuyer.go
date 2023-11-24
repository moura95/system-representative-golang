package dashboard

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"my-orders/internal/repository"
	"my-orders/internal/util"
)

type topBuyerRequest struct {
	StartDate time.Time `form:"start_date" binding:"required" time_format:"2006-01-02"`
	EndDate   time.Time `form:"end_date" binding:"required" time_format:"2006-01-02"`
	Top       int32     `form:"top" binding:"required"`
}

type topBuyerResponse struct {
	CustomerName  []string `json:"customer_name"`
	TotalCustomer []string `json:"total_customer"`
	Total         string   `json:"total"`
}

func (d *Dashboard) topBuyer(ctx *gin.Context) {
	representativeID := ctx.Keys["representativeID"].(int32)

	var req topBuyerRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	listTop, err := d.Db.TopBuyerByRepresentativeID(ctx, repository.TopBuyerByRepresentativeIDParams{
		RepresentativeID: representativeID,
		StartDate:        req.StartDate,
		EndDate:          req.EndDate,
		Top:              req.Top,
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

	var response topBuyerResponse
	for _, row := range listTop {
		response.CustomerName = append(response.CustomerName, row.CustomerName)
		response.TotalCustomer = append(response.TotalCustomer, row.TotalByCustomer)
		response.Total = row.TotalSales
	}
	ctx.JSON(http.StatusOK, util.SuccessResponse(200, response, ""))
	return
}
