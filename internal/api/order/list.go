package order

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"my-orders/internal/repository"
	"my-orders/internal/util"
)

type listRequest struct {
	IsActive bool `form:"is_active"`
}

func (o *Order) list(ctx *gin.Context) {
	representativeID := ctx.Keys["representativeID"].(int32)

	var req listRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	listOrders, err := o.Db.ListOrdersByRepresentativeID(ctx, repository.ListOrdersByRepresentativeIDParams{
		RepresentativeID: representativeID,
		IsActive:         req.IsActive,
	})
	if err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseRead.Error()))
		return
	}

	if len(listOrders) == 0 {
		ctx.JSON(http.StatusOK, util.SuccessResponse(200, listOrders, ""))
		return
	}

	var listResponse []getResponse
	for _, order := range listOrders {
		listResponse = append(listResponse, getResponse{
			OrderID:         order.ID,
			FactoryID:       order.FactoryID,
			CustomerID:      order.CustomerID,
			PortageID:       order.PortageID,
			SellerID:        order.SellerID,
			FormPaymentID:   order.FormPaymentID.Int32,
			OrderNumber:     order.OrderNumber,
			UrlPdf:          order.UrlPdf.String,
			Buyer:           order.Buyer.String,
			Shipping:        order.Shipping,
			Status:          order.Status,
			ExpireAt:        order.ExpiredAt,
			Total:           order.Total,
			IsActive:        order.IsActive,
			FactoryName:     order.FactoryName,
			CustomerName:    order.CustomerName,
			PortageName:     order.PortageName,
			SellerName:      order.SellerName,
			SellerEmail:     order.SellerEmail.String,
			CustomerEmail:   order.CustomerEmail.String,
			FormPaymentName: order.FormPaymentName.String,
			CreatedAt:       order.CreatedAt,
		})
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, listResponse, ""))
	return
}
