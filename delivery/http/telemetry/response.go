package telemetry_delivery_http

import (
	vm "rocky.my.id/git/h8-assignment-3/application/telemetry/view_models"
)

type TelemetryResponse struct {
	Status vm.TelemetryVM `json:"status"`
}

func NewTelemetryResponse(payload vm.TelemetryVM) *TelemetryResponse {
	return &TelemetryResponse{Status: payload}
}
