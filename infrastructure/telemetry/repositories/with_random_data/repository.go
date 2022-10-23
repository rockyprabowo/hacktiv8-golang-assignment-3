package telemetry_repository_random_data

import (
	"context"
	"rocky.my.id/git/h8-assignment-3/application/telemetry/contracts"
	"rocky.my.id/git/h8-assignment-3/domain/entities"
)

type TelemetryRepository struct {
	Generator telemetry_contracts.TelemetryDataGeneratorContract
}

func NewTelemetryRepository(generator telemetry_contracts.TelemetryDataGeneratorContract) *TelemetryRepository {
	instance := &TelemetryRepository{Generator: generator}
	instance.Generator.Start()
	return instance
}

func (t TelemetryRepository) Get(_ context.Context) *entities.Telemetry {
	return t.Generator.Data()
}
