package telemetry_delivery_http

import (
	"github.com/labstack/echo/v4"
	telemetryUseCases "rocky.my.id/git/h8-assignment-3/application/telemetry"
)

type TelemetryHTTPDeliveryDependencies struct {
	UseCases *telemetryUseCases.TelemetryUseCases
	Engine   *echo.Echo
}

func Setup(dependencies TelemetryHTTPDeliveryDependencies) {
	var (
		HTTPHandler = NewTelemetryHTTPHandler(dependencies.UseCases)
		router      = NewTelemetryHTTPRouter(dependencies.Engine, HTTPHandler)
	)
	router.Setup()
}
