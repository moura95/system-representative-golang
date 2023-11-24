package product

import (
	"github.com/gin-gonic/gin"

	"my-orders/internal/repository"
)

type IProduct interface {
	SetupProductRoute(planRoutes *gin.RouterGroup)
}

type Product struct {
	Db repository.Querier
}

func (p *Product) SetupProductRoute(planRoutes *gin.RouterGroup) {
	planRoutes.POST("/product", p.create)
	planRoutes.DELETE("/product/:id", p.delete)
	planRoutes.GET("/product/:id", p.get)
	planRoutes.GET("/product", p.list)
	planRoutes.DELETE("/product/:id/remove", p.remove)
	planRoutes.GET("/product/:id/restore", p.restore)
	planRoutes.PATCH("/product", p.update)
}
