package telemetry_delivery_http

import (
	"github.com/labstack/echo/v4"
	"rocky.my.id/git/h8-assignment-3/application/telemetry"
	"rocky.my.id/git/h8-assignment-3/delivery/http/common/contracts"
)

type TelemetryHTTPDelivery struct {
	UseCases    *telemetry_use_cases.TelemetryUseCases
	HTTPHandler TelemetryHTTPHandlerContract
	Routes      common_delivery_http_contracts.RouterContract
}

func SetupDefault(engine *echo.Echo, useCases *telemetry_use_cases.TelemetryUseCases) {
	var (
		HTTPHandler = NewTelemetryHTTPHandler(useCases)
		routes      = NewTelemetryHTTPRouter(HTTPHandler, engine)
	)

	app := TelemetryHTTPDelivery{
		UseCases:    useCases,
		HTTPHandler: HTTPHandler,
		Routes:      routes,
	}

	app.Routes.Setup()
}
