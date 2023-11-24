package orderItems

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"my-orders/internal/repository"
	"my-orders/internal/util"
)

type deleteRequest struct {
	OrderID   int32 `uri:"order_id" binding:"required,numeric"`
	ProductID int32 `uri:"product_id" binding:"required,numeric"`
}

func (o *OrderItem) delete(ctx *gin.Context) {
	var req deleteRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	err := o.Db.DeleteOrderItemsByID(ctx, repository.DeleteOrderItemsByIDParams{
		OrderID:   req.OrderID,
		ProductID: req.ProductID,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseCreate.Error()))
			return
		}
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseCreate.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, "ok", ""))
	return
}
