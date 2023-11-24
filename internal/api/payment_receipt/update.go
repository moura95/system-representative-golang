package paymentReceipt

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"

	"my-orders/internal/repository"
	"my-orders/internal/util"
)

type updateRequest struct {
	ID             int32                             `json:"id"`
	PaymentForm    repository.PaymentReceiptFormType `json:"payment_form"`
	Status         repository.PaymentReceiptStatus   `json:"status"`
	Description    string                            `json:"description"`
	Amount         string                            `json:"amount"`
	ExpirationDate time.Time                         `json:"expiration_date,omitempty"`
	PaymentDate    time.Time                         `json:"payment_date,omitempty"`
	DocNumber      string                            `json:"doc_number"`
	Recipient      string                            `json:"recipient"`
	AdditionalInfo string                            `json:"additional_info"`
}

func (p *PaymentReceipt) update(ctx *gin.Context) {
	var req updateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	_, err := p.Db.UpdatePaymentPaymentReceiptByID(ctx, repository.UpdatePaymentPaymentReceiptByIDParams{
		ID:             req.ID,
		PaymentForm:    repository.NullPaymentReceiptFormType{PaymentReceiptFormType: req.PaymentForm, Valid: req.PaymentForm != ""},
		Description:    sql.NullString{String: req.Description, Valid: req.Description != ""},
		Amount:         sql.NullString{String: req.Amount, Valid: req.Amount != ""},
		ExpirationDate: sql.NullTime{Time: req.ExpirationDate, Valid: !req.ExpirationDate.IsZero()},
		PaymentDate:    sql.NullTime{Time: req.PaymentDate, Valid: !req.PaymentDate.IsZero()},
		DocNumber:      sql.NullString{String: req.DocNumber, Valid: req.DocNumber != ""},
		Recipient:      sql.NullString{String: req.Recipient, Valid: req.Recipient != ""},
		AdditionalInfo: sql.NullString{String: req.AdditionalInfo, Valid: req.AdditionalInfo != ""},
	})

	if err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseUpdate.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, "ok", ""))
	return
}
