package stripe

import (
	"github.com/gin-gonic/gin"

	"my-orders/cfg"
	"my-orders/internal/repository"
)

type IStripe interface {
	SetupStripeRoute(routerGroup, authRoutes *gin.RouterGroup)
}

type Stripe struct {
	Db     repository.Querier
	Config cfg.Config
}

func (s *Stripe) SetupStripeRoute(routes, authRoutes *gin.RouterGroup) {
	routes.Any("/stripe/webhook", s.handlerEvent)
	authRoutes.POST("/stripe/checkout", s.checkoutCreator)
}
