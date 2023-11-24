package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	_ "github.com/newrelic/go-agent/v3/integrations/nrpq"
	"github.com/stripe/stripe-go/v72"

	"my-orders/cfg"
	"my-orders/internal/api"
	"my-orders/internal/repository"
)

func main() {
	loadConfig, _ := cfg.LoadConfig(".")
	stripe.Key = loadConfig.StripeKey

	var conn *sql.DB

	if loadConfig.GinMode == "release" {
		conn, _ = sql.Open("nrpostgres", loadConfig.DBSource)
	} else {
		conn, _ = sql.Open("postgres", loadConfig.DBSource)
	}

	store := repository.New(conn)
	log.Print("connection is database establish")

	runGinServer(loadConfig, store)
}

func runGinServer(cfg cfg.Config, store repository.Querier) {
	server := api.NewServer(cfg, store)

	_ = server.Start(cfg.HTTPServerAddress)
}
