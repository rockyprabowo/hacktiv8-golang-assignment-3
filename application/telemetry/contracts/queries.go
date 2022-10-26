package telemetry_contracts

import (
	"context"

	vm "rocky.my.id/git/h8-assignment-3/application/telemetry/view_models"
)

type TelemetryQueriesContract interface {
	GetTelemetry(ctx context.Context) vm.TelemetryVM
	GetTelemetryConfig(ctx context.Context) vm.TelemetryConfigVM
}
