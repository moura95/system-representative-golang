package orderItems

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"my-orders/internal/util"
)

type listRequest struct {
	OrderID int32 `uri:"order_id" binding:"required,numeric"`
}

func (o *OrderItem) list(ctx *gin.Context) {
	var req listRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	listOrdersItems, err := o.Db.ListOrdersItemsByOrderID(ctx, req.OrderID)
	if err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseRead.Error()))
		return
	}

	if len(listOrdersItems) == 0 {
		ctx.JSON(http.StatusOK, util.SuccessResponse(200, listOrdersItems, ""))
		return
	}

	var listResponse []getResponse
	for _, order := range listOrdersItems {
		listResponse = append(listResponse, getResponse{
			OrderID:     order.OrderID,
			ProductID:   order.ProductID,
			Quantity:    order.Quantity,
			Price:       order.Price,
			Discount:    order.Discount,
			ProductName: order.ProductName,
			Ipi:         order.Ipi.String,
			Description: order.Description.String,
			Code:        order.Code,
			Total:       order.Total,
		})
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, listResponse, ""))
}
