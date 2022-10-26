package telemetry_delivery_http

import (
	"net/http"

	"github.com/labstack/echo/v4"
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

	err = c.JSON(http.StatusOK, jsonResponse)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return
}

func (h TelemetryHTTPHandler) GetConfig(c echo.Context) (err error) {
	data := h.UseCases.Queries.GetTelemetryConfig(c.Request().Context())

	err = c.JSON(http.StatusOK, data)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return
}
