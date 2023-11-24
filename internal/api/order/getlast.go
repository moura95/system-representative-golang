package order

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"my-orders/internal/util"
)

func (o *Order) getLast(ctx *gin.Context) {
	representativeID := ctx.Keys["representativeID"].(int32)

	orderNumber, err := o.Db.GetLastOrderByRepresentativeID(ctx, representativeID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusOK, util.SuccessResponse(200, 1, ""))
			return
		}
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseRead.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, orderNumber+1, ""))
}
