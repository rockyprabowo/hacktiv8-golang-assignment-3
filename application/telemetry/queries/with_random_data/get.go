package telemetry_query_with_random

import (
	. "rocky.my.id/git/h8-assignment-3/application/telemetry/view_models"
)

func (t TelemetryRandomDataQuery) GetTelemetry() TelemetryVM {
	currentData := t.generator.Data()
	return TelemetryVM{
		Water: currentData.Water,
		Wind:  currentData.Wind,
	}
}
