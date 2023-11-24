package order

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"my-orders/internal/repository"
	"my-orders/internal/util"
)

type createRequest struct {
	FactoryID     int32                   `json:"factory_id"`
	CustomerID    int32                   `json:"customer_id"`
	PortageID     int32                   `json:"portage_id"`
	SellerID      int32                   `json:"seller_id"`
	FormPaymentID int32                   `json:"form_payment_id"`
	OrderNumber   int32                   `json:"order_number"`
	UrlPdf        string                  `json:"url_pdf"`
	Buyer         string                  `json:"buyer"`
	Shipping      repository.ShippingEnum `json:"shipping"`
	Status        repository.StatusEnum   `json:"status"`
	CreatedAt     time.Time               `json:"created_at"`
}

func (o *Order) Create(ctx *gin.Context) {
	representativeID := ctx.Keys["representativeID"].(int32)

	var req createRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	if req.Shipping == "" {
		req.Shipping = repository.ShippingEnumOutros

	}

	order, err := o.Db.CreateOrder(ctx, repository.CreateOrderParams{
		RepresentativeID: representativeID,
		FactoryID:        req.FactoryID,
		CustomerID:       req.CustomerID,
		PortageID:        req.PortageID,
		SellerID:         req.SellerID,
		FormPaymentID:    sql.NullInt32{Int32: req.FormPaymentID, Valid: req.FormPaymentID != 0},
		OrderNumber:      req.OrderNumber,
		UrlPdf:           sql.NullString{String: req.UrlPdf, Valid: req.UrlPdf != ""},
		Buyer:            sql.NullString{String: req.Buyer, Valid: req.Buyer != ""},
		Shipping:         repository.ShippingEnum(req.Shipping),
		Status:           req.Status,
		CreatedAt:        req.CreatedAt,
	})
	if err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseCreate.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, order.ID, ""))
	return
}
