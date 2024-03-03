package main

import (
	"time"

	appInit "lolipad/boilerplate/init"

	log "go.uber.org/zap"

	_baseHttpHandler "lolipad/boilerplate/module/base/handler/http"
	_baseRepo "lolipad/boilerplate/module/base/store"
	_base "lolipad/boilerplate/module/base/usecase"

	// _ "lolipad/boilerplate/docs"

	// echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4" //nolint
	"github.com/labstack/echo/v4/middleware"
)

var config *appInit.Config
var logger *log.Logger

func init() {
	// Start pre-requisite app dependencies
	config, logger = appInit.StartAppInit()
}

// @title base
// @version 1.0

// @BasePath /base/v1
func main() {
	// Get PG Conn Instance
	pgDb, err := appInit.ConnectToPGServer()
	if err != nil {
		log.S().Fatal(err)
	}

	// init router
	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// Prometheus
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	router := e.Group("/bases")

	timeoutContext := time.Duration(config.Context.Timeout) * time.Second

	// DI: Repository & Usecase
	baseRepo := _baseRepo.NewStore(pgDb)

	baseUc := _base.NewAutoTextUsecase(baseRepo, timeoutContext, config)

	// End of DI Steps

	_baseHttpHandler.NewBaseHandler(router, baseUc)

	// router.GET("/swagger/*", echoSwagger.WrapHandler)

	// start serve
	e.Logger.Fatal(e.Start(config.API.Port))
}
