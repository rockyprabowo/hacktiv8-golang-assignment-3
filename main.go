package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"rocky.my.id/git/h8-assignment-3/application/telemetry"
	"rocky.my.id/git/h8-assignment-3/application/telemetry/queries"
	"rocky.my.id/git/h8-assignment-3/delivery/http/telemetry"
	"rocky.my.id/git/h8-assignment-3/infrastructure/telemetry/generators/random_data_generator"
	"rocky.my.id/git/h8-assignment-3/infrastructure/telemetry/repositories/with_random_data"
	"syscall"
	"time"
)

const (
	defaultRefreshDuration = 15 * time.Second
	defaultServeAddress    = "localhost:8080"
)

var (
	serveAddress    string
	baseURL         string
	refreshDuration time.Duration
	debug           bool
	ctx, stop       = signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
)

func main() {
	var engine = echo.New()
	engine.Server = &http.Server{
		Addr:         serveAddress,
		Handler:      engine,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
	}

	var (
		telemetryGenerator  = telemetry_random_data_generator.NewTelemetryRandomDataGenerator(ctx, refreshDuration)
		telemetryRepository = telemetry_repository_random_data.NewTelemetryRepository(telemetryGenerator)
		telemetryUseCases   = &telemetry_use_cases.TelemetryUseCases{
			Queries: telemetry_queries.NewTelemetryQueries(
				telemetryRepository,
			),
		}
		telemetryDeliveryDeps = telemetry_delivery_http.TelemetryHTTPDeliveryDependencies{
			UseCases: telemetryUseCases,
			Engine:   engine,
		}
	)

	telemetry_delivery_http.Setup(telemetryDeliveryDeps)

	go func() {
		if err := engine.Start(serveAddress); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server forced to shutdown: ", err)
		}
	}()

	<-time.After(1 * time.Second)
	fmt.Printf("Base URL: %s\n", baseURL)
	fmt.Printf("Refreshing telemetry data every %s\n", refreshDuration.String())

	<-ctx.Done()
	log.Println("Shutting down gracefully, press Ctrl+C again to force")
	stop()

	if err := engine.Shutdown(context.Background()); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server has been shutdown.")
}
