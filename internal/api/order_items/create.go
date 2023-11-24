package orderItems

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"

	"my-orders/internal/repository"
	"my-orders/internal/util"
)

type createRequest struct {
	OrderID   int32  `json:"order_id"`
	ProductID int32  `json:"product_id"`
	Quantity  int32  `json:"quantity"`
	Price     string `json:"price"`
	Discount  string `json:"discount"`
}

func (o *OrderItem) Create(ctx *gin.Context) {
	var req createRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}
	if req.Discount == "" {
		req.Discount = "0"
	}

	_, err := o.Db.CreateOrderItems(ctx, repository.CreateOrderItemsParams{
		OrderID:   req.OrderID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
		Price:     req.Price,
		Discount:  req.Discount,
	})
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusOK, util.ErrorResponse(409, "", util.ErrorItemDuplicate.Error()))
				return
			}
		}
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseCreate.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, "ok", ""))
	return
}
