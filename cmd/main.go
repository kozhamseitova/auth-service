package main

import (
	"context"
	"log"

	"github.com/kozhamseitova/auth-service/internal/app"
	"github.com/kozhamseitova/auth-service/internal/config"

	"github.com/swaggo/gin-swagger" // gin-swagger middleware
    "github.com/swaggo/files" // swagger embed files
)

// @title           Auth Service
// @version         0.0.1
// @description     Api for simple auth with access and refresh tokens.

// @contact.name   Aisha
// @contact.email  kozhamseitova91@gmail.com

// @host      localhost:8080
// @BasePath  

// @securitydefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization


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