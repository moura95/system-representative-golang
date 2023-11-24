package imports

import (
	"github.com/gin-gonic/gin"

	"my-orders/internal/repository"
)

type IImports interface {
	SetupImportRoute(planRoutes *gin.RouterGroup)
}

type Import struct {
	Db repository.Querier
}

func (i *Import) SetupImportRoute(planRoutes *gin.RouterGroup) {
	g := planRoutes.Group("/import")
	g.POST("/product/:factoryId", i.importProducts)
	g.POST("/company", i.importCompanies)
}
