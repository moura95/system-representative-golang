package orderItems

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"my-orders/internal/repository"
	"my-orders/internal/util"
)

type getRequest struct {
	OrderID   int32 `uri:"order_id" binding:"required,numeric"`
	ProductID int32 `uri:"product_id" binding:"required,numeric"`
}

type getResponse struct {
	OrderID     int32  `json:"order_id"`
	ProductID   int32  `json:"product_id"`
	Quantity    int32  `json:"quantity"`
	Price       string `json:"price"`
	Discount    string `json:"discount"`
	ProductName string `json:"product_name"`
	Ipi         string `json:"ipi"`
	Description string `json:"description"`
	Code        string `json:"code"`
	Total       string `json:"total"`
}

func (o *OrderItem) get(ctx *gin.Context) {
	var req getRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	order, err := o.Db.GetOrderItemsByID(ctx, repository.GetOrderItemsByIDParams{
		OrderID:   req.OrderID,
		ProductID: req.ProductID,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusOK, util.ErrorResponse(200, []string{}, util.ErrorDatabaseRead.Error()))
			return
		}
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseRead.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, getResponse{
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
	}, ""))
}
