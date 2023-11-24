package lead

import (
	"github.com/gin-gonic/gin"

	"my-orders/cfg"
	"my-orders/internal/repository"
	"my-orders/internal/token"
)

type ILead interface {
	SetupLeadRoute(routes, authRoutes *gin.RouterGroup)
}

type Lead struct {
	Db         repository.Querier
	TokenMaker token.Maker
	Config     cfg.Config
}

func (l *Lead) SetupLeadRoute(routes, authRoutes *gin.RouterGroup) {
	routes.POST("/lead", l.create)
	authRoutes.GET("/lead", l.list)
	authRoutes.GET("/lead/:id", l.get)
	authRoutes.DELETE("/lead/:id", l.delete)

}
