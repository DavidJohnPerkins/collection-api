package main

import (
	"context"
	"dperkins/collection-api/api"
	"dperkins/collection-api/config"
	"dperkins/collection-api/store"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	store := store.NewSqlServerCollectionStore(cfg.DatabaseURL)
	server := api.NewServer(cfg.HTTPServer, store)
	server.Start(ctx)
}
