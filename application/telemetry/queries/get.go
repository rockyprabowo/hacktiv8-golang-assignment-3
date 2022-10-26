package telemetry_queries

import (
	"context"

	vm "rocky.my.id/git/h8-assignment-3/application/telemetry/view_models"
)

func (t TelemetryQueries) GetTelemetry(ctx context.Context) vm.TelemetryVM {
	currentData := t.Repository.Get(ctx)
	return vm.TelemetryVM{
		Water: currentData.Water,
		Wind:  currentData.Wind,
	}
}

func (t TelemetryQueries) GetTelemetryConfig(ctx context.Context) vm.TelemetryConfigVM {
	return *t.Repository.Config(ctx)
}
