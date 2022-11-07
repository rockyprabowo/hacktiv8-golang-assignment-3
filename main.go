package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	telemetryUseCases "rocky.my.id/git/h8-assignment-3/application/telemetry"
	telemetryQueries "rocky.my.id/git/h8-assignment-3/application/telemetry/queries"
	telemetryDeliveryHttp "rocky.my.id/git/h8-assignment-3/delivery/http/telemetry"
	telemetryRandomDataGenerator "rocky.my.id/git/h8-assignment-3/infrastructure/telemetry/generators/random_data_generator"
	telemetryRepositoryRandomData "rocky.my.id/git/h8-assignment-3/infrastructure/telemetry/repositories/with_random_data"
)

const (
	defaultMaxRandomInt    = 100
	defaultRefreshDuration = 15 * time.Second
	defaultServeAddress    = "localhost:8080"
)

var (
	serveAddress    string
	baseURL         string
	refreshDuration time.Duration
	maxRandomInt    int
	debug           bool
)

func main() {
	var (
		engine    = echo.New()
		ctx, stop = signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	)

	engine.Server = &http.Server{
		Addr:              serveAddress,
		Handler:           engine,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}

	var (
		generator  = telemetryRandomDataGenerator.NewTelemetryRandomDataGenerator(ctx, refreshDuration, maxRandomInt)
		repository = telemetryRepositoryRandomData.NewTelemetryRepository(generator)
		useCases   = &telemetryUseCases.TelemetryUseCases{
			Queries: telemetryQueries.NewTelemetryQueries(
				repository,
			),
		}
		deliveryDeps = telemetryDeliveryHttp.TelemetryHTTPDeliveryDependencies{
			UseCases: useCases,
			Engine:   engine,
		}
	)

	engine.Use(middleware.Recover())
	setupCORS(engine)

	telemetryDeliveryHttp.Setup(deliveryDeps)

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

	{
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := engine.Shutdown(shutdownCtx); err != nil {
			log.Fatal("Server forced to shutdown: ", err)
		}
	}

	log.Println("Server has been shutdown.")
}
