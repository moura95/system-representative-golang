package file

import (
	"github.com/gin-gonic/gin"

	"my-orders/cfg"
	"my-orders/internal/repository"
)

type IFile interface {
	SetupFileRoute(routerGroup, planRoutes *gin.RouterGroup)
}

type File struct {
	Db     repository.Querier
	Config cfg.Config
}

func (f *File) SetupFileRoute(routes, planRoutes *gin.RouterGroup) {
	planRoutes.POST("/file/upload", f.upload)
	planRoutes.DELETE("/file/delete/:filename", f.delete)
}
