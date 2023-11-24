package company

import (
	"github.com/gin-gonic/gin"

	"my-orders/internal/repository"
)

type ICompany interface {
	SetupCompanyRoute(planRoutes *gin.RouterGroup)
}

type Company struct {
	Db repository.Querier
}

func (c *Company) SetupCompanyRoute(planRoutes *gin.RouterGroup) {
	planRoutes.POST("/company", c.create)
	planRoutes.DELETE("/company/:id", c.delete)
	planRoutes.GET("/company/:id", c.get)
	planRoutes.GET("/company", c.list)
	planRoutes.DELETE("/company/:id/remove", c.remove)
	planRoutes.GET("/company/:id/restore", c.restore)
	planRoutes.PATCH("/company", c.update)
}
