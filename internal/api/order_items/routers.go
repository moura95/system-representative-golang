package orderItems

import (
	"github.com/gin-gonic/gin"

	"my-orders/internal/repository"
)

type IOrderItem interface {
	SetupOrderItemRoute(planRoutes *gin.RouterGroup)
}

type OrderItem struct {
	Db repository.Querier
}

func (o *OrderItem) SetupOrderItemRoute(planRoutes *gin.RouterGroup) {
	planRoutes.POST("/order/item", o.Create)
	planRoutes.DELETE("/order/:order_id/item/:product_id", o.delete)
	planRoutes.GET("/order/:order_id/item/:product_id", o.get)
	planRoutes.GET("/order/:order_id/item", o.list)
	planRoutes.PATCH("/order/item", o.update)
}
