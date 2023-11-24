package dashboard

import (
	"github.com/gin-gonic/gin"

	"my-orders/internal/repository"
)

type IDashboard interface {
	SetupDashboardRoute(planRoutes *gin.RouterGroup)
}

type Dashboard struct {
	Db repository.Querier
}

func (d *Dashboard) SetupDashboardRoute(planRoutes *gin.RouterGroup) {
	planRoutes.GET("/dashboard/topbuyer", d.topBuyer)
	planRoutes.GET("/dashboard/totalsales", d.topSales)
	planRoutes.GET("/dashboard/topproduct", d.topProduct)
	planRoutes.GET("/dashboard/topfactory", d.topFactory)
}
