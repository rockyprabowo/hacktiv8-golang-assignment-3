package telemetry_delivery_http

import (
	"github.com/labstack/echo/v4"
)

type TelemetryHTTPRouter struct {
	Handler TelemetryHTTPHandlerContract
	Router  *echo.Echo
}

func NewTelemetryHTTPRouter(handler TelemetryHTTPHandlerContract, router *echo.Echo) *TelemetryHTTPRouter {
	return &TelemetryHTTPRouter{Handler: handler, Router: router}
}

func (routes *TelemetryHTTPRouter) Setup() {
	routeGroup := routes.Router.Group("/api/v1/telemetry")
	{
		routeGroup.GET("", routes.Handler.Get)
	}
}
