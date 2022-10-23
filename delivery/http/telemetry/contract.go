package telemetry_delivery_http

import "github.com/labstack/echo/v4"

type TelemetryHTTPHandlerContract interface {
	Get(echo.Context) error
}
