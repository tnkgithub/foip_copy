package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	// Logger
	"github.com/brpaz/echozap"
	"go.uber.org/zap"

	// pkg
	"foip/core/config"
	"foip/core/pkg/controller"
)

/* swagの仕様書
 * https://github.com/swaggo/swag#api-operation
 */

// @title       Simple Backend API
// @version     0.0.1
// @description This is a sample server.
// @host        localhost:8080
// @BasePath    /api/v1
func New(cfg *config.Config) (*echo.Echo, error) {
	ctrl := controller.New(cfg)
	server := echo.New()

	// Setup Logger
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	server.Use(echozap.ZapLogger(logger))

	// Routes
	server.GET("/metrics", ctrl.Metrics())
	server.GET("/swagger/*", ctrl.Swagger())

	apiv1 := server.Group("api/v1")
	api := apiv1
	api.GET("/token", ctrl.GetConnectionToken)

	// Setup CORS
	server.Use(middleware.CORSWithConfig(
		cfg.Server.CORS,
	))

	return server, nil
}
