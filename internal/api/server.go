package api

import (
	"github.com/axiaoxin-com/ratelimiter"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
	"log"
	"my-orders/cfg"
	"my-orders/internal/repository"
	"my-orders/internal/token"
	"time"
)

type Server struct {
	store      *repository.Querier
	router     *gin.Engine
	tokenMaker *token.Maker
	config     *cfg.Config
}

func NewServer(cfg cfg.Config, store repository.Querier) *Server {
	tokenMaker, _ := token.NewPasetoMaker(cfg.TokenSymmetricKey)

	server := &Server{
		config:     &cfg,
		store:      &store,
		tokenMaker: &tokenMaker,
	}

	var router *gin.Engine

	if server.config.GinMode == "release" {
		relic, err := newrelic.NewApplication(
			newrelic.ConfigAppName(cfg.NewRelicAppName),
			newrelic.ConfigLicense(cfg.NewRelicKey),
			newrelic.ConfigAppLogForwardingEnabled(cfg.NewRelicLogForward),
		)
		if err != nil {
			log.Fatal("failed to instantiate newrelic")
		}

		router = gin.Default()
		router.Use(nrgin.Middleware(relic))
		router.Use(gzip.Gzip(gzip.DefaultCompression))
		router.Use(ratelimiter.GinMemRatelimiter(ratelimiter.GinRatelimiterConfig{
			// config: how to generate a limit key
			LimitKey: func(c *gin.Context) string {
				return c.ClientIP()
			},
			// config: how to respond when limiting
			LimitedHandler: func(c *gin.Context) {
				c.JSON(200, "too many requests!!!")
				c.Abort()
				return
			},
			// config: return ratelimiter token fill interval and bucket size
			TokenBucketConfig: func(*gin.Context) (time.Duration, int) {
				return time.Second * 60, 4000
			},
		}))

	} else {
		router = gin.Default()
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders("plan")
	corsConfig.AddAllowHeaders("redirect")
	corsConfig.AddAllowHeaders("Authorization")
	router.Use(cors.New(corsConfig))
	server.createRoutesV1(router)

	server.router = router
	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
