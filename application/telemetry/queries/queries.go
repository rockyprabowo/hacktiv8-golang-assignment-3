package telemetry_queries

import (
	contracts "rocky.my.id/git/h8-assignment-3/application/telemetry/contracts"
)

type TelemetryQueries struct {
	Repository contracts.TelemetryRepositoryContract
}

func NewTelemetryQueries(repository contracts.TelemetryRepositoryContract) *TelemetryQueries {
	return &TelemetryQueries{
		Repository: repository,
	}
}
