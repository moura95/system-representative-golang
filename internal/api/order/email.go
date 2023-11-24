package order

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"my-orders/internal/reports"
	"my-orders/internal/reports/factory"
	"my-orders/internal/util"
)

type pdfEmailRequest struct {
	OrderID       int32  `uri:"order_id" binding:"required,numeric"`
	CustomerEmail string `json:"customer_email"`
}

func (o *Order) sendMailOrder(ctx *gin.Context) {
	representativeID := ctx.Keys["representativeID"].(int32)

	var req pdfEmailRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}
	smtp, _ := o.Db.GetSmtpByRepresentativeID(ctx, representativeID)

	pdf := reports.PDFReport{}
	file, err := pdf.CreatePDFReport(&factory.OrderReportFactory{
		Db:           o.Db,
		PdfClientURL: o.Config.PdfTurtleClient,
		OrderID:      req.OrderID,
		Ctx:          ctx,
	})
	if err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorGeneratePDF.Error()))
		return
	}

	data := util.Data{
		Title: "Pedido de compra",
		Name:  "MidasGestor",
		Msg:   "Segue em anexo o pedido de compra",
	}
	err = util.NewSender("MidasGestor", smtp).SendEmail("Pedido de compra", data, "envio_pedido",
		[]string{req.CustomerEmail}, nil, nil, file,
	)

	if err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorSendEmail.Error()))
		return
	}
	ctx.JSON(http.StatusOK, util.SuccessResponse(200, "ok", ""))
	return
}
