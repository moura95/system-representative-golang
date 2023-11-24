package paymentReceipt

import (
	"github.com/gin-gonic/gin"

	"my-orders/internal/repository"
)

type IPaymentReceipt interface {
	SetupPaymentReceiptRoute(planRoutes *gin.RouterGroup)
}

type PaymentReceipt struct {
	Db repository.Querier
}

func (p *PaymentReceipt) SetupPaymentReceiptRoute(planRoutes *gin.RouterGroup) {
	planRoutes.POST("/payment_receipt", p.create)
	planRoutes.DELETE("/payment_receipt/:id", p.delete)
	planRoutes.GET("/payment_receipt/:id", p.get)
	planRoutes.GET("/payment_receipt", p.list)
	planRoutes.PATCH("/payment_receipt", p.update)
	planRoutes.POST("/payment_receipt/file/:payment_receipt_id/", p.uploadFile)
	planRoutes.DELETE("/payment_receipt/file/:payment_receipt_id/", p.deleteFile)
}
