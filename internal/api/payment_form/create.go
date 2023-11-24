package paymentForm

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"my-orders/internal/repository"
	"my-orders/internal/util"
)

type createRequest struct {
	Name string `json:"name"`
}

func (p *PaymentForm) create(ctx *gin.Context) {
	representativeID := ctx.Keys["representativeID"].(int32)

	var req createRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	_, err := p.Db.CreatePaymentForm(ctx, repository.CreatePaymentFormParams{
		RepresentativeID: representativeID,
		Name:             req.Name,
	})
	if err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseCreate.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, "ok", ""))
	return
}
