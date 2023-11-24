package calendar

import (
	"github.com/gin-gonic/gin"

	"my-orders/internal/repository"
)

type ICalendar interface {
	SetupCalendarRoute(planRoutes *gin.RouterGroup)
}

type Calendar struct {
	Db repository.Querier
}

func (c *Calendar) SetupCalendarRoute(planRoutes *gin.RouterGroup) {
	planRoutes.POST("/calendar", c.create)
	planRoutes.DELETE("/calendar/:id", c.delete)
	planRoutes.GET("/calendar/:id", c.get)
	planRoutes.GET("/calendar", c.list)
	planRoutes.PATCH("/calendar", c.update)
}
