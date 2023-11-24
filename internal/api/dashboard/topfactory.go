package dashboard

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"my-orders/internal/repository"
	"my-orders/internal/util"
)

type topFactoryRequest struct {
	StartDate time.Time `form:"start_date" binding:"required" time_format:"2006-01-02"`
	EndDate   time.Time `form:"end_date" binding:"required" time_format:"2006-01-02"`
	Top       int32     `form:"top" binding:"required"`
}

type topFactoryResponse struct {
	FactoryName  []string `json:"factory_name"`
	TotalFactory []string `json:"total_factory"`
	Total        string   `json:"total"`
}

func (d *Dashboard) topFactory(ctx *gin.Context) {
	representativeID := ctx.Keys["representativeID"].(int32)

	var req topFactoryRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	listTop, err := d.Db.TopFactoryByRepresentativeID(ctx, repository.TopFactoryByRepresentativeIDParams{
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

	var response topFactoryResponse
	for _, row := range listTop {
		response.FactoryName = append(response.FactoryName, row.FactoryName)
		response.TotalFactory = append(response.TotalFactory, row.SubTotal)
		response.Total = row.Total
	}
	ctx.JSON(http.StatusOK, util.SuccessResponse(200, response, ""))
	return
}
