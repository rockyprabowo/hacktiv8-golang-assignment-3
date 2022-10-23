package telemetry_random_data_generator

import (
	"context"
	"log"
	"math/rand"
	"time"

	. "rocky.my.id/git/h8-assignment-3/domain/entities"
)

type TelemetryRandomDataGenerator struct {
	telemetry       *Telemetry
	context         context.Context
	refreshDuration time.Duration
}

func NewTelemetryRandomDataGenerator(ctx context.Context, refreshDuration time.Duration) *TelemetryRandomDataGenerator {
	instance := &TelemetryRandomDataGenerator{
		telemetry:       &Telemetry{},
		context:         ctx,
		refreshDuration: refreshDuration,
	}
	return instance
}

func (g TelemetryRandomDataGenerator) GeneratorFunc() {
	ticker := time.NewTicker(g.refreshDuration)
	log.Println("TelemetryRandomDataGenerator started.")
	for {
		*g.telemetry = Telemetry{
			Water: rand.Intn(100),
			Wind:  rand.Intn(100),
		}
		log.Printf("TelemetryRandomDataGenerator: new telemetry data recieved: %+v", *g.telemetry)

		select {
		case <-g.context.Done():
			ticker.Stop()
			log.Println("TelemetryRandomDataGenerator stopped.")
			return
		case <-ticker.C:
		}
	}
}

func (g TelemetryRandomDataGenerator) Data() *Telemetry {
	return g.telemetry
}

func (g TelemetryRandomDataGenerator) Start() {
	go g.GeneratorFunc()
}
