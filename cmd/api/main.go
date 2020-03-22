package main

import (
	"github.com/cobbinma/example-go-api/cmd/api/handler"
	"github.com/cobbinma/example-go-api/pkg/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func main() {
	p := config.GetPort()
	e := getRouter()
	logrus.Infof("Listening for requests on http://localhost:%s/", p)
	e.Logger.Fatal(e.Start(":" + p))
}

func getRouter() *echo.Echo {
	e := echo.New()
	h := handler.NewHandler()

	e.Use(middleware.Logger())

	e.GET("/healthz", h.Health)
	e.GET("/oas", h.Oas)

	return e
}