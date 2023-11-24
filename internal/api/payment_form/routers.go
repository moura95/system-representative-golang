package paymentForm

import (
	"github.com/gin-gonic/gin"

	"my-orders/internal/repository"
)

type IPaymentForm interface {
	SetupPaymentFormRoute(planRoutes *gin.RouterGroup)
}

type PaymentForm struct {
	Db repository.Querier
}

func (p *PaymentForm) SetupPaymentFormRoute(planRoutes *gin.RouterGroup) {
	planRoutes.POST("/payment_form", p.create)
	planRoutes.DELETE("/payment_form/:id", p.delete)
	planRoutes.GET("/payment_form/:id", p.get)
	planRoutes.GET("/payment_form", p.list)
	planRoutes.PATCH("/payment_form", p.update)
}
