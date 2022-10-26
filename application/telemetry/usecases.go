package telemetry_use_cases

import contracts "rocky.my.id/git/h8-assignment-3/application/telemetry/contracts"

type TelemetryUseCases struct {
	Queries  contracts.TelemetryQueriesContract
	Commands interface{}
}

// noinspection GoUnusedExportedFunction
func NewTelemetryUseCases(queries contracts.TelemetryQueriesContract) *TelemetryUseCases {
	return &TelemetryUseCases{Queries: queries}
}
