package telemetry_delivery_http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	telemetryUseCases "rocky.my.id/git/h8-assignment-3/application/telemetry"
)

type TelemetryHTTPHandler struct {
	UseCases *telemetryUseCases.TelemetryUseCases
}

func NewTelemetryHTTPHandler(useCases *telemetryUseCases.TelemetryUseCases) *TelemetryHTTPHandler {
	return &TelemetryHTTPHandler{UseCases: useCases}
}

func (h TelemetryHTTPHandler) Get(c echo.Context) (err error) {
	data := h.UseCases.Queries.GetTelemetry(c.Request().Context())

	jsonResponse := NewTelemetryResponse(data)

	return c.JSON(http.StatusOK, jsonResponse)
}

func (h TelemetryHTTPHandler) GetConfig(c echo.Context) (err error) {
	data := h.UseCases.Queries.GetTelemetryConfig(c.Request().Context())

	return c.JSON(http.StatusOK, data)
}
