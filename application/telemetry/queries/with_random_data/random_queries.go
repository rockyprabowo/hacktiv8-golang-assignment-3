package telemetry_query_with_random

import (
	"rocky.my.id/git/h8-assignment-3/application/telemetry/contracts"
)

type TelemetryRandomDataQuery struct {
	generator telemetry_contracts.TelemetryDataGeneratorContract
}

func NewTelemetryRandomQueries(generator telemetry_contracts.TelemetryDataGeneratorContract) *TelemetryRandomDataQuery {
	instance := &TelemetryRandomDataQuery{
		generator: generator,
	}
	instance.generator.Start()
	return instance
}
