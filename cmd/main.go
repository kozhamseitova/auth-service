package main

import (
	"context"
	"log"

	"github.com/kozhamseitova/auth-service/internal/app"
	"github.com/kozhamseitova/auth-service/internal/config"
)

func main() {
	ctx := context.Background()
	cfg, err := config.InitConfig("config.yaml")
	if err != nil {
		panic(err)
	}
	
	err = app.Run(ctx, cfg)
	if err != nil {
		log.Fatalf("failed to run app: %v", err)
	}
}