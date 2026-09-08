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

	dbx, err := store.InitSharedDB(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to init db: %v", err)
	}

	collectionStore := store.NewSqlServerCollectionStore(dbx)
	server := api.NewServer(cfg.HTTPServer, collectionStore)
	server.Start(ctx)
}
