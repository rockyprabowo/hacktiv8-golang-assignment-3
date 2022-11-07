package telemetry_repository_random_data

import (
	"context"

	contracts "rocky.my.id/git/h8-assignment-3/application/telemetry/contracts"
	vm "rocky.my.id/git/h8-assignment-3/application/telemetry/view_models"
	"rocky.my.id/git/h8-assignment-3/domain/entities"
)

type TelemetryRepository struct {
	Generator contracts.TelemetryDataGeneratorContract
}

func NewTelemetryRepository(generator contracts.TelemetryDataGeneratorContract) *TelemetryRepository {
	instance := &TelemetryRepository{Generator: generator}
	instance.Generator.Start()
	return instance
}

func (t TelemetryRepository) Get(_ context.Context) *entities.Telemetry {
	return t.Generator.Data()
}

func (t TelemetryRepository) Config(_ context.Context) *vm.TelemetryConfigVM {
	interval := t.Generator.Interval()
	return &vm.TelemetryConfigVM{
		Interval:   interval.String(),
		IntervalMs: interval.Milliseconds(),
		WaterStates: map[string]string{
			"safe":    "<= 5",
			"warning": ">=<= 6 8",
			"danger":  "> 8",
		},
		WindStates: map[string]string{
			"safe":    "<= 6",
			"warning": ">=<= 7 15",
			"danger":  "> 15",
		},
	}
}
