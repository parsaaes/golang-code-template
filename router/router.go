package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rashadansari/golang-code-template/buildinfo"
	"github.com/rashadansari/golang-code-template/config"
	"github.com/sirupsen/logrus"
)

func New(server config.Server) *echo.Echo {
	e := echo.New()

	debug := logrus.IsLevelEnabled(logrus.DebugLevel)

	e.Debug = debug

	e.HideBanner = true

	if !debug {
		e.HidePort = true
	}

	e.Server.ReadTimeout = server.ReadTimeout
	e.Server.WriteTimeout = server.WriteTimeout

	recoverConfig := middleware.DefaultRecoverConfig
	recoverConfig.DisablePrintStack = !debug
	e.Use(middleware.RecoverWithConfig(recoverConfig))

	e.Use(middleware.CORS())
	e.Use(buildinfo.Middleware)
	e.Use(prometheusMiddleware())

	return e
}
