package paymentReceipt

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"my-orders/internal/repository"
	"my-orders/internal/util"
	"net/http"
)

type uploadFileRequest struct {
	PaymentReceiceptID int32  `uri:"payment_receipt_id" binding:"required,numeric"`
	FileUrl            string `json:"file_url"`
}

func (p *PaymentReceipt) uploadFile(ctx *gin.Context) {
	var req uploadFileRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}
	_, err := p.Db.UploadFilePaymentReceipt(ctx, repository.UploadFilePaymentReceiptParams{
		PaymentReceiptID: req.PaymentReceiceptID,
		UrlFile:          req.FileUrl,
	},
	)
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
