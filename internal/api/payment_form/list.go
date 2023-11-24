package paymentForm

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"my-orders/internal/util"
)

func (p *PaymentForm) list(ctx *gin.Context) {
	representativeID := ctx.Keys["representativeID"].(int32)

	listPaymentForms, err := p.Db.ListPaymentFormsByRepresentativeID(ctx, representativeID)
	if err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}
	if len(listPaymentForms) == 0 {
		ctx.JSON(http.StatusOK, util.SuccessResponse(200, listPaymentForms, ""))
		return
	}

	var listResponse []getResponse
	for _, seller := range listPaymentForms {
		listResponse = append(listResponse, getResponse{
			ID:   seller.ID,
			Name: seller.Name,
		})
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, listResponse, ""))
}
