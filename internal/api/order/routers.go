package order

import (
	"github.com/gin-gonic/gin"

	"my-orders/cfg"
	"my-orders/internal/repository"
)

type IOrder interface {
	SetupOrderRoute(planRoutes *gin.RouterGroup)
}

type Order struct {
	Db     repository.Querier
	Config cfg.Config
}

func (o *Order) SetupOrderRoute(planRoutes *gin.RouterGroup) {
	planRoutes.POST("/order", o.Create)
	planRoutes.DELETE("/order/:order_id", o.delete)
	planRoutes.GET("/order/:order_id", o.get)
	planRoutes.GET("/order/last", o.getLast)
	planRoutes.GET("/order/:order_id/pdf", o.getPDF)
	planRoutes.POST("/order/:order_id/email", o.sendMailOrder)
	planRoutes.GET("/order", o.list)
	planRoutes.DELETE("/order/:order_id/remove", o.remove)
	planRoutes.GET("/order/:order_id/restore", o.restore)
	planRoutes.PATCH("/order", o.update)
}
