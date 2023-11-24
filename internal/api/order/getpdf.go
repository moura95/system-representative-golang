package order

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"my-orders/internal/reports"
	"my-orders/internal/reports/factory"
	"my-orders/internal/util"
)

type pdfRequest struct {
	OrderID int32 `uri:"order_id" binding:"required,numeric"`
}

func (o *Order) getPDF(ctx *gin.Context) {
	var req pdfRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	report := reports.PDFReport{}
	pdf, err := report.CreatePDFReport(&factory.OrderReportFactory{
		Db:           o.Db,
		PdfClientURL: o.Config.PdfTurtleClient,
		OrderID:      req.OrderID,
		Ctx:          ctx,
	})
	if err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorGeneratePDF.Error()))
		return
	}

	ctx.Data(http.StatusOK, "application/pdf", pdf)
	return
}
