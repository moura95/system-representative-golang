package paymentReceipt

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"my-orders/internal/util"
)

func (p *PaymentReceipt) list(ctx *gin.Context) {
	representativeID := ctx.Keys["representativeID"].(int32)

	listPaymentsReceipt, err := p.Db.ListPaymentReceiptByRepresentativeID(ctx, representativeID)
	if err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}
	if len(listPaymentsReceipt) == 0 {
		ctx.JSON(http.StatusOK, util.SuccessResponse(200, listPaymentsReceipt, ""))
		return
	}

	var listResponse []getResponse
	for _, paymentReceipt := range listPaymentsReceipt {
		files := []fileResponse{}
		err = json.Unmarshal(paymentReceipt.Files, &files)
		if err != nil {
			ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
			return
		}
		if files[0].UrlFile == "" {
			files = []fileResponse{}
		}
		listResponse = append(listResponse, getResponse{
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
		})
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, listResponse, ""))
}
