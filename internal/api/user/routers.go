package user

import (
	"github.com/gin-gonic/gin"

	"my-orders/cfg"
	"my-orders/internal/repository"
	"my-orders/internal/token"
)

type IUser interface {
	SetupUserRoute(routes, authRoutes *gin.RouterGroup)
}

type User struct {
	Db         repository.Querier
	TokenMaker token.Maker
	Config     cfg.Config
}

func (u *User) SetupUserRoute(routes, authRoutes *gin.RouterGroup) {
	authRoutes.POST("/user", u.create)
	authRoutes.DELETE("/user/:id", u.delete)
	authRoutes.GET("/user/:id", u.get)
	authRoutes.GET("/user", u.list)
	routes.POST("/user/login", u.login)
	authRoutes.PUT("/user/password", u.updatePassword)
	authRoutes.DELETE("/user/:id/remove", u.remove)
	authRoutes.GET("/user/:id/restore", u.restore)
	authRoutes.PATCH("/user", u.update)
}
