package telemetry_contracts

import (
	"rocky.my.id/git/h8-assignment-3/application/telemetry/view_models"
)

type TelemetryQueriesContract interface {
	GetTelemetry() telemetry_view_models.TelemetryVM
}
