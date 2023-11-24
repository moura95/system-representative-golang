package orderItems

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"my-orders/internal/repository"
	"my-orders/internal/util"
)

type updateRequest struct {
	OrderID   int32  `json:"order_id"`
	ProductID int32  `json:"product_id"`
	Quantity  int32  `json:"quantity"`
	Price     string `json:"price"`
	Discount  string `json:"discount"`
}

func (o *OrderItem) update(ctx *gin.Context) {
	var req updateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	_, err := o.Db.UpdateOrderItemByID(ctx, repository.UpdateOrderItemByIDParams{
		OrderID:   req.OrderID,
		ProductID: req.ProductID,
		Quantity:  sql.NullInt32{Int32: req.Quantity, Valid: req.Quantity != 0},
		Price:     sql.NullString{String: req.Price, Valid: req.Price != ""},
		Discount:  sql.NullString{String: req.Discount, Valid: req.Discount != ""},
	})
	if err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseUpdate.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, "ok", ""))
	return
}
