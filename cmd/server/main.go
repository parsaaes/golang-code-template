package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/rashadansari/golang-code-template/config"
	"github.com/rashadansari/golang-code-template/handler"
	"github.com/rashadansari/golang-code-template/metric"
	"github.com/rashadansari/golang-code-template/model"
	"github.com/rashadansari/golang-code-template/postgres"
	"github.com/rashadansari/golang-code-template/redis"
	"github.com/rashadansari/golang-code-template/router"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

//nolint:funlen
func main(cfg config.Config) {
	e := router.New(cfg.Server)

	postgresDb := postgres.WithRetry(postgres.Create, cfg.Postgres)

	defer func() {
		if err := postgresDb.Close(); err != nil {
			logrus.Errorf("postgres connection close error: %s", err.Error())
		}
	}()

	_, redisClose := redis.Create(cfg.Redis.Master)

	defer func() {
		if err := redisClose(); err != nil {
			logrus.Errorf("redis master connection close error: %s", err.Error())
		}
	}()

	templateRepo := model.SQLTemplateRepo{DB: postgresDb}

	templateHandler := handler.TemplateHandler{TemplateRepo: templateRepo}

	e.GET("/healthz", func(c echo.Context) error { return c.NoContent(http.StatusNoContent) })

	v1 := e.Group("/v1")

	v1.POST("/template", templateHandler.Create)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := e.Start(cfg.Server.Address); err != nil {
			logrus.Fatalf("failed to start server: %s", err.Error())
		}
	}()

	go metric.StartPrometheusServer(cfg.Monitoring.Prometheus)

	logrus.Info("start server!")

	s := <-sig

	logrus.Infof("signal %s received", s)

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Server.GracefulTimeout)
	defer cancel()

	e.Server.SetKeepAlivesEnabled(false)

	if err := e.Shutdown(ctx); err != nil {
		logrus.Errorf("failed to shutdown server: %s", err.Error())
	}
}

// Register server command
func Register(root *cobra.Command, cfg config.Config) {
	root.AddCommand(
		&cobra.Command{
			Use:   "server",
			Short: "Template Server Component",
			Run: func(cmd *cobra.Command, args []string) {
				main(cfg)
			},
		},
	)
}
