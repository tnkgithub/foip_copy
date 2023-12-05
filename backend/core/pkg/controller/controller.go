package controller

import (
	"foip/core/config"
	_ "foip/core/docs"
	"foip/core/pkg/service/livekit"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swagger "github.com/swaggo/echo-swagger"
)

type Controller struct {
	Livekit config.LivekitConfig
}

func New(cfg *config.Config) *Controller {
	ctrl := Controller{
		Livekit: cfg.Livekit,
	}

	return &ctrl
}

// Metrics godoc
// @Summary     Metrics for Prometheus
// @Description show metrics for prometheus
// @Tags        admin,monitoring
// @Produce     text/plain
// @Router      /metrics [get]
func (c *Controller) Metrics() echo.HandlerFunc { return echo.WrapHandler(promhttp.Handler()) }

// Swagger godoc
// @Summary     API Documentation
// @Description show API documentation
// @Tags        admin
// @Produce     text/html
// @Router      /swagger/ [get]
func (c *Controller) Swagger() echo.HandlerFunc { return swagger.WrapHandler }

// Token Generator godoc
// @Summary     API Documentation
// @Description show API documentation
// @Tags        livekit, user
// @Produce     text/json
// @Param       room       query  string true "Room"
// @Param       user query string true   "User"
// @Router      /api/v1/token [get]
func (c *Controller) GetConnectionToken(ctx echo.Context) error {
	room, user := ctx.QueryParam("room"), ctx.QueryParam("user")

	token, err := livekit.GenerateAccessToken(room, user, c.Livekit)
	if err != nil {
		ctx.JSON(http.StatusOK, ErrorResponse{Code: 200, Message: err.Error()})
	}

	return ctx.JSON(http.StatusOK, GetConnectionTokenRespone{Token: token})
}
