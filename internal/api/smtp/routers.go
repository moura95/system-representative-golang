package smtp

import (
	"github.com/gin-gonic/gin"

	"my-orders/internal/repository"
)

type ISmtp interface {
	SetupSmtpRoute(routes *gin.RouterGroup)
}

type Smtp struct {
	Db repository.Querier
}

func (s *Smtp) SetupSmtpRoute(routes *gin.RouterGroup) {
	routes.POST("/smtp", s.create)
	routes.DELETE("/smtp", s.delete)
	routes.GET("/smtp", s.get)
	routes.PATCH("/smtp", s.update)
}
