package order

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"my-orders/internal/repository"
	"my-orders/internal/util"
)

type updateRequest struct {
	ID            int32                   `json:"id" binding:"required"`
	FactoryID     int32                   `json:"factory_id" binding:"required,numeric"`
	CustomerID    int32                   `json:"customer_id" binding:"required,numeric"`
	PortageID     int32                   `json:"portage_id" binding:"required,numeric"`
	SellerID      int32                   `json:"seller_id" binding:"required,numeric"`
	FormPaymentID int32                   `json:"form_payment_id"`
	UrlPdf        string                  `json:"url_pdf"`
	Buyer         string                  `json:"buyer"`
	Shipping      repository.ShippingEnum `json:"shipping"`
	Status        repository.StatusEnum   `json:"status"`
	ExpireAt      time.Time               `json:"expire_order"`
	IsActive      bool                    `json:"is_active"`
	CreatedAt     time.Time               `json:"created_at"`
	UpdateAt      time.Time               `json:"update_at"`
}

func (o *Order) update(ctx *gin.Context) {
	var req updateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	_, err := o.Db.UpdateOrderByID(ctx, repository.UpdateOrderByIDParams{
		ID:            req.ID,
		FactoryID:     sql.NullInt32{Int32: req.FactoryID, Valid: req.FactoryID != 0},
		CustomerID:    sql.NullInt32{Int32: req.CustomerID, Valid: req.CustomerID != 0},
		PortageID:     sql.NullInt32{Int32: req.PortageID, Valid: req.PortageID != 0},
		SellerID:      sql.NullInt32{Int32: req.SellerID, Valid: req.SellerID != 0},
		FormPaymentID: sql.NullInt32{Int32: req.FormPaymentID, Valid: req.FormPaymentID != 0},
		UrlPdf:        sql.NullString{String: req.UrlPdf, Valid: req.UrlPdf != ""},
		Buyer:         sql.NullString{String: req.Buyer, Valid: req.Buyer != ""},
		Shipping:      repository.NullShippingEnum{ShippingEnum: req.Shipping, Valid: true},
		Status:        repository.NullStatusEnum{StatusEnum: req.Status, Valid: true},
		ExpiredAt:     sql.NullTime{Time: req.ExpireAt, Valid: !req.ExpireAt.IsZero()},
		CreatedAt:     sql.NullTime{Time: req.CreatedAt, Valid: !req.CreatedAt.IsZero()},
	})
	if err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseUpdate.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, "ok", ""))
	return
}
