package telemetry_use_cases

import "rocky.my.id/git/h8-assignment-3/application/telemetry/contracts"

type TelemetryUseCases struct {
	Queries  telemetry_contracts.TelemetryQueriesContract
	Commands interface{}
}

// noinspection GoUnusedExportedFunction
func NewTelemetryUseCases(queries telemetry_contracts.TelemetryQueriesContract) *TelemetryUseCases {
	return &TelemetryUseCases{Queries: queries}
}
