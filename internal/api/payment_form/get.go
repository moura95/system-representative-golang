package paymentForm

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"my-orders/internal/util"
)

type getRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type getResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func (p *PaymentForm) get(ctx *gin.Context) {
	var req getRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	paymentForm, err := p.Db.GetPaymentFormByID(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseRead.Error()))
			return
		}
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseRead.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, getResponse{
		ID:   paymentForm.ID,
		Name: paymentForm.Name,
	}, ""))
}
