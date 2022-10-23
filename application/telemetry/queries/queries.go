package telemetry_queries

import (
	"rocky.my.id/git/h8-assignment-3/application/telemetry/contracts"
)

type TelemetryQueries struct {
	Repository telemetry_contracts.TelemetryRepositoryContract
}

func NewTelemetryQueries(repository telemetry_contracts.TelemetryRepositoryContract) *TelemetryQueries {
	instance := &TelemetryQueries{
		Repository: repository,
	}
	return instance
}
