package seller

import (
	"github.com/gin-gonic/gin"

	"my-orders/internal/repository"
)

type ISeller interface {
	SetupSellerRoute(planRoutes *gin.RouterGroup)
}

type Seller struct {
	Db repository.Querier
}

func (s *Seller) SetupSellerRoute(planRoutes *gin.RouterGroup) {
	planRoutes.POST("/seller", s.create)
	planRoutes.DELETE("/seller/:id", s.delete)
	planRoutes.GET("/seller/:id", s.get)
	planRoutes.GET("/seller", s.list)
	planRoutes.DELETE("/seller/:id/remove", s.remove)
	planRoutes.GET("/seller/:id/restore", s.restore)
	planRoutes.PATCH("/seller", s.update)
}
