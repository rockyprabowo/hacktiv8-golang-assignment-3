package telemetry_queries

import (
	"context"
	. "rocky.my.id/git/h8-assignment-3/application/telemetry/view_models"
)

func (t TelemetryQueries) GetTelemetry(ctx context.Context) TelemetryVM {
	currentData := t.Repository.Get(ctx)
	return TelemetryVM{
		Water: currentData.Water,
		Wind:  currentData.Wind,
	}
}
