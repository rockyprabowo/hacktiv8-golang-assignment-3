package telemetry_delivery_http

import (
	"github.com/labstack/echo/v4"
	"rocky.my.id/git/h8-assignment-3/application/telemetry"
)

type TelemetryHTTPDeliveryDependencies struct {
	UseCases *telemetry_use_cases.TelemetryUseCases
	Engine   *echo.Echo
}

func Setup(dependencies TelemetryHTTPDeliveryDependencies) {
	var (
		HTTPHandler = NewTelemetryHTTPHandler(dependencies.UseCases)
		router      = NewTelemetryHTTPRouter(dependencies.Engine, HTTPHandler)
	)
	router.Setup()

}
