package representative

import (
	"github.com/gin-gonic/gin"

	"my-orders/internal/repository"
)

type IRepresentative interface {
	SetupRepresentativeRoute(routes, planRoutes *gin.RouterGroup)
}
type Representative struct {
	Db repository.Querier
}

func (r *Representative) SetupRepresentativeRoute(routes, authRoutes *gin.RouterGroup) {
	routes.POST("/representative", r.create)
	authRoutes.GET("/representative", r.get)
	authRoutes.GET("/representative/plan", r.getPlan)
	authRoutes.DELETE("/representative", r.delete)
	authRoutes.PATCH("/representative", r.update)
	authRoutes.DELETE("/representative/remove", r.remove)
	authRoutes.GET("/representative/restore", r.restore)
}
