package paymentReceipt

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"my-orders/internal/util"
	"net/http"
)

type deleteFileRequest struct {
	PaymentReceiceptID int32 `uri:"payment_receipt_id" binding:"required,numeric"`
}

func (p *PaymentReceipt) deleteFile(ctx *gin.Context) {
	var req deleteFileRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}
	_, err := p.Db.DeleteFilePaymentReceiptByID(ctx, req.PaymentReceiceptID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseUpdate.Error()))
			return
		}
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseUpdate.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, "ok", ""))
	return
}
