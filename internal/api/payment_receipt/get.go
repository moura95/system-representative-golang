package paymentReceipt

import (
	"database/sql"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"my-orders/internal/repository"
	"net/http"

	"my-orders/internal/util"
)

type getRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type getResponse struct {
	ID             int32          `json:"id"`
	TypePayment    string         `json:"type_payment"`
	Status         string         `json:"status"`
	Amount         string         `json:"amount"`
	Description    string         `json:"description"`
	ExpirationDate sql.NullTime   `json:"expiration_date"`
	PaymentDate    sql.NullTime   `json:"payment_date"`
	DocNumber      string         `json:"doc_number"`
	Recipient      string         `json:"recipient"`
	PaymentForm    string         `json:"payment_form"`
	Installment    int32          `json:"installment"`
	AdditionalInfo string         `json:"additional_info"`
	Files          []fileResponse `json:"files"`
}
type fileResponse struct {
	ID               int32  `json:"file_id"`
	PaymentReceiptID int32  `json:"payment_receipt_id"`
	UrlFile          string `json:"url_file"`
	CreateAt         string `json:"create_at"`
}

func (p *PaymentReceipt) get(ctx *gin.Context) {
	representativeID := ctx.Keys["representativeID"].(int32)
	var req getRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}
	params := repository.GetPaymentReceiptByIDParams{
		ID:               req.ID,
		RepresentativeID: representativeID,
	}
	paymentReceipt, err := p.Db.GetPaymentReceiptByID(ctx, params)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseRead.Error()))
			return
		}
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseRead.Error()))
		return
	}
	files := []fileResponse{}
	err = json.Unmarshal(paymentReceipt.Files, &files)
	if err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, getResponse{
		ID:             paymentReceipt.ID,
		TypePayment:    string(paymentReceipt.TypePayment),
		Status:         string(paymentReceipt.Status),
		Amount:         paymentReceipt.Amount,
		Description:    paymentReceipt.Description,
		ExpirationDate: paymentReceipt.ExpirationDate,
		PaymentDate:    paymentReceipt.PaymentDate,
		DocNumber:      paymentReceipt.DocNumber.String,
		Recipient:      paymentReceipt.Recipient.String,
		PaymentForm:    string(paymentReceipt.PaymentForm),
		Installment:    paymentReceipt.Installment,
		AdditionalInfo: paymentReceipt.AdditionalInfo.String,
		Files:          files,
	}, ""))
}
