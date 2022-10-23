package telemetry_delivery_http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"rocky.my.id/git/h8-assignment-3/application/telemetry"
)

type TelemetryHTTPHandler struct {
	UseCases *telemetry_use_cases.TelemetryUseCases
}

func NewTelemetryHTTPHandler(useCases *telemetry_use_cases.TelemetryUseCases) *TelemetryHTTPHandler {
	return &TelemetryHTTPHandler{UseCases: useCases}
}

func (h TelemetryHTTPHandler) Get(c echo.Context) (err error) {
	data := h.UseCases.Queries.GetTelemetry()

	jsonResponse := NewTelemetryResponse(data)

	err = c.JSON(http.StatusOK, jsonResponse)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return
}
