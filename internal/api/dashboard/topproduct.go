package dashboard

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"my-orders/internal/repository"
	"my-orders/internal/util"
)

type topProductRequest struct {
	StartDate time.Time `form:"start_date" binding:"required" time_format:"2006-01-02"`
	EndDate   time.Time `form:"end_date" binding:"required" time_format:"2006-01-02"`
	Top       int32     `form:"top" binding:"required"`
}

type topProductResponse struct {
	ProductName []string `json:"product_name"`
	SubTotal    []string `json:"sub_total"`
	Total       string   `json:"total"`
}

func (d *Dashboard) topProduct(ctx *gin.Context) {
	representativeID := ctx.Keys["representativeID"].(int32)

	var req topProductRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	listTop, err := d.Db.TopSalesPerProductByRepresentativeID(ctx, repository.TopSalesPerProductByRepresentativeIDParams{
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

	var response topProductResponse
	for _, row := range listTop {
		response.ProductName = append(response.ProductName, row.ProductName)
		response.SubTotal = append(response.SubTotal, row.SubTotal)
		response.Total = row.Total
	}
	ctx.JSON(http.StatusOK, util.SuccessResponse(200, response, ""))
	return
}
