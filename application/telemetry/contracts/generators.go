package telemetry_contracts

import (
	"time"

	"rocky.my.id/git/h8-assignment-3/domain/entities"
)

type TelemetryDataGeneratorContract interface {
	Start()
	GeneratorFunc()
	Data() *entities.Telemetry
	Interval() time.Duration
}
