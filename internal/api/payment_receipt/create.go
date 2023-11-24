package paymentReceipt

import (
	"database/sql"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"my-orders/internal/repository"
	"my-orders/internal/util"
)

type createRequest struct {
	TypePayment         repository.PaymentReceiptType     `json:"type_payment"`
	Status              repository.PaymentReceiptStatus   `json:"status"`
	Description         string                            `json:"description"`
	Amount              string                            `json:"amount"`
	ExpirationDate      time.Time                         `json:"expiration_date,omitempty"`
	PaymentDate         time.Time                         `json:"payment_date,omitempty"`
	DocNumber           string                            `json:"doc_number"`
	Recipient           string                            `json:"recipient"`
	PaymentForm         repository.PaymentReceiptFormType `json:"payment_form"`
	NumbersInstallments int                               `json:"numbers_installments"`
	Interval            int                               `json:"interval"`
	AdditionalInfo      string                            `json:"additional_info"`
	CompetentDate       string                            `json:"competent_date"`
}

func (p *PaymentReceipt) create(ctx *gin.Context) {
	representativeID := ctx.Keys["representativeID"].(int32)

	var req createRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}
	for i := 0; i < req.NumbersInstallments; i++ {
		expirationDate := req.ExpirationDate.AddDate(0, 0, req.Interval*i)
		if strings.ToTitle(req.CompetentDate) == "Default" {
			expirationDate = time.Now().AddDate(0, 0, req.Interval*i)
		}
		_, err = p.Db.CreatePaymentReceipt(ctx, repository.CreatePaymentReceiptParams{
			RepresentativeID: representativeID,
			TypePayment:      req.TypePayment,
			Status:           req.Status,
			Description:      req.Description,
			Amount:           req.Amount,
			ExpirationDate:   sql.NullTime{Time: expirationDate, Valid: expirationDate != time.Time{}},
			PaymentDate:      sql.NullTime{Time: req.PaymentDate, Valid: req.PaymentDate != time.Time{}},
			DocNumber:        sql.NullString{String: req.DocNumber, Valid: req.DocNumber != ""},
			Recipient:        sql.NullString{String: req.Recipient, Valid: req.Recipient != ""},
			PaymentForm:      req.PaymentForm,
			Installment:      int32(i) + 1,
			AdditionalInfo:   sql.NullString{String: req.AdditionalInfo, Valid: req.AdditionalInfo != ""},
		})
		if err != nil {
			ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseCreate.Error()))
			return
		}
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, "ok", ""))
	return
}
