package telemetry_contracts

import (
	"rocky.my.id/git/h8-assignment-3/domain/entities"
)

type TelemetryDataGeneratorContract interface {
	Start()
	GeneratorFunc()
	Data() *entities.Telemetry
}
