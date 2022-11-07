package telemetry_random_data_generator

import (
	"context"
	"log"
	"math/rand"
	"time"

	"rocky.my.id/git/h8-assignment-3/domain/entities"
)

type TelemetryRandomDataGenerator struct {
	telemetry       *entities.Telemetry
	context         context.Context
	refreshDuration time.Duration
	maxRandomInt    int
}

func NewTelemetryRandomDataGenerator(ctx context.Context, refreshDuration time.Duration, maxRandomInt int) *TelemetryRandomDataGenerator {
	instance := &TelemetryRandomDataGenerator{
		telemetry:       &entities.Telemetry{},
		context:         ctx,
		refreshDuration: refreshDuration,
		maxRandomInt:    maxRandomInt,
	}
	return instance
}

func (g TelemetryRandomDataGenerator) GeneratorFunc() {
	ticker := time.NewTicker(g.refreshDuration)
	log.Println("TelemetryRandomDataGenerator started.")
	for {
		*g.telemetry = entities.Telemetry{
			Water: rand.Intn(g.maxRandomInt),
			Wind:  rand.Intn(g.maxRandomInt),
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

func (g TelemetryRandomDataGenerator) Data() *entities.Telemetry {
	return g.telemetry
}

func (g TelemetryRandomDataGenerator) Start() {
	go g.GeneratorFunc()
}

func (g TelemetryRandomDataGenerator) Interval() time.Duration {
	return g.refreshDuration
}
