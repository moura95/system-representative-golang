package paseto

import (
	"github.com/gin-gonic/gin"

	"my-orders/cfg"
	"my-orders/internal/repository"
	"my-orders/internal/token"
)

type IToken interface {
	SetupTokenRoute(routes *gin.RouterGroup)
}

type Token struct {
	Db         repository.Querier
	TokenMaker token.Maker
	Config     cfg.Config
}

func (t *Token) SetupTokenRoute(routes *gin.RouterGroup) {
	routes.POST("/token", t.renewAccessToken)
}
