package telemetry_delivery_http

import (
	. "rocky.my.id/git/h8-assignment-3/application/telemetry/view_models"
)

type TelemetryResponse struct {
	Status TelemetryVM `json:"status"`
}

func NewTelemetryResponse(payload TelemetryVM) *TelemetryResponse {
	return &TelemetryResponse{Status: payload}
}
