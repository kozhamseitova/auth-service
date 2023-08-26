package app

import (
	"context"
	"os"
	"os/signal"

	"github.com/kozhamseitova/auth-service/internal/config"
	"github.com/kozhamseitova/auth-service/internal/handler"
	"github.com/kozhamseitova/auth-service/internal/repository"
	"github.com/kozhamseitova/auth-service/internal/service"
	"github.com/kozhamseitova/auth-service/pkg/client/mongodb"
	"github.com/kozhamseitova/auth-service/pkg/httpserver"
	"github.com/kozhamseitova/auth-service/pkg/jwttoken"
	"github.com/kozhamseitova/auth-service/pkg/logger"
)

func Run(ctx context.Context, cfg *config.Config) error {
	dbClient, err := mongodb.Connect(
		mongodb.WithHost(cfg.DB.Host),
		mongodb.WithPort(cfg.DB.Port),
		mongodb.WithDBName(cfg.DB.DBName),
	)

	if err != nil {
		return err
	}

	token := jwttoken.New(cfg.Token.SecretKey)

	logger, err := logger.NewLogger(cfg.App.Production)
	if err != nil {
		return err
	}

	repo := repository.NewRepository(dbClient, cfg.DB, logger)
	srvc := service.NewService(repo, token, cfg.Token, logger)
	hndlr := handler.NewHandler(srvc, cfg)

	server := httpserver.New(
		hndlr,
		cfg,
	)

	server.Start()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	select {
	case s := <-interrupt:
		logger.Infof(ctx, "signal received: %s", s.String())
	case err = <-server.Notify():
		logger.Errorf(ctx, "server notify: %s", err.Error())
	}

	err = server.Shutdown()
	if err != nil {
		logger.Errorf(ctx, "server shutdown err: %s", err)
	}


	return nil
}