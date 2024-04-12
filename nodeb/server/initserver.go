package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/keshu12345/guardianlink/nodeb/config"

	logger "github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

func Initialize(router *gin.Engine, cfg *config.Configuration, lifecycle fx.Lifecycle) {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: router,
	}
	server.ReadTimeout = time.Duration(cfg.Server.ReadTimeout) * time.Second
	server.WriteTimeout = time.Duration(cfg.Server.WriteTimeout) * time.Second
	server.IdleTimeout = time.Duration(cfg.Server.IdleTimeout) * time.Second

	lifecycle.Append(fx.Hook{
		OnStart: func(context context.Context) error {
			logger.Info(fmt.Sprintf("Starting the REST application with %s environment and with port is %v", cfg.EnvironmentName, cfg.Server.Port))
			go func() {
				if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					logger.Fatalf("listen: %s\n", err)
				}
			}()
			return nil
		},
		OnStop: func(i context.Context) error {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			err := server.Shutdown(ctx)
			if err != nil {
				logger.Fatal("Server Shutdown: ", err)
			}
			logger.Info("Server exiting")
			return err
		},
	})
}
