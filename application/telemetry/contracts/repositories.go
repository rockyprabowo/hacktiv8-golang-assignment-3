package telemetry_contracts

import (
	"context"

	vm "rocky.my.id/git/h8-assignment-3/application/telemetry/view_models"
	"rocky.my.id/git/h8-assignment-3/domain/entities"
)

type TelemetryRepositoryContract interface {
	Get(context.Context) *entities.Telemetry
	Config(context.Context) *vm.TelemetryConfigVM
}
