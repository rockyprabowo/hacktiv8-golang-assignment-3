package telemetry_contracts

import (
	"context"
	"rocky.my.id/git/h8-assignment-3/domain/entities"
)

type TelemetryRepositoryContract interface {
	Get(context.Context) *entities.Telemetry
}
