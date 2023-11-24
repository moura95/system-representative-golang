package order

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"my-orders/internal/repository"
	"my-orders/internal/util"
)

type getRequest struct {
	OrderID int32 `uri:"order_id" binding:"required,numeric"`
}

type getResponse struct {
	OrderID         int32                   `json:"order_id"`
	FactoryID       int32                   `json:"factory_id"`
	CustomerID      int32                   `json:"customer_id"`
	PortageID       int32                   `json:"portage_id"`
	SellerID        int32                   `json:"seller_id"`
	FormPaymentID   int32                   `json:"form_payment_id"`
	OrderNumber     int32                   `json:"order_number"`
	UrlPdf          string                  `json:"url_pdf"`
	Buyer           string                  `json:"buyer"`
	Shipping        repository.ShippingEnum `json:"shipping"`
	Status          repository.StatusEnum   `json:"status"`
	ExpireAt        time.Time               `json:"expire_order"`
	Total           string                  `json:"total,float64"`
	IsActive        bool                    `json:"is_active"`
	FactoryName     string                  `json:"factory_name"`
	CustomerName    string                  `json:"customer_name"`
	PortageName     string                  `json:"portage_name"`
	SellerName      string                  `json:"seller_name"`
	SellerEmail     string                  `json:"seller_email"`
	CustomerEmail   string                  `json:"customer_email"`
	FormPaymentName string                  `json:"form_payment_name"`
	CreatedAt       time.Time               `json:"created_at"`
	CustomerCnpj    string                  `json:"customer_cnpj"`
	FactoryCnpj     string                  `json:"factory_cnpj"`
}

func (o *Order) get(ctx *gin.Context) {
	var req getRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	order, err := o.Db.GetOrderByID(ctx, req.OrderID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseRead.Error()))
			return
		}
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseRead.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, getResponse{
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
		CustomerCnpj:    order.CustomerCnpj.String,
		FactoryCnpj:     order.FactoryCnpj.String,
	}, ""))
}
