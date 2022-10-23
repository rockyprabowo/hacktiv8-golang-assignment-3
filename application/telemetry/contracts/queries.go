package telemetry_contracts

import (
	"context"
	"rocky.my.id/git/h8-assignment-3/application/telemetry/view_models"
)

type TelemetryQueriesContract interface {
	GetTelemetry(ctx context.Context) telemetry_view_models.TelemetryVM
}
