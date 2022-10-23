package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
	"os/signal"
	"rocky.my.id/git/h8-assignment-3/application/telemetry"
	"rocky.my.id/git/h8-assignment-3/application/telemetry/generators/random_data_generator"
	"rocky.my.id/git/h8-assignment-3/application/telemetry/queries/with_random_data"
	"rocky.my.id/git/h8-assignment-3/delivery/http/telemetry"
	"syscall"
	"time"
)

const (
	defaultRefreshDuration = 15 * time.Second
	defaultServeAddress    = "localhost:8080"
)

var (
	serveAddress    string
	refreshDuration time.Duration
	debug           bool
	ctx, stop       = signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
)

func main() {
	var e = echo.New()
	e.Server = &http.Server{
		Addr:         serveAddress,
		Handler:      e,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	var (
		telemetryUseCases = &telemetry_use_cases.TelemetryUseCases{
			Queries: telemetry_query_with_random.NewTelemetryRandomQueries(
				telemetry_random_data_generator.NewTelemetryRandomDataGenerator(ctx, refreshDuration),
			),
		}
	)

	telemetry_delivery_http.SetupDefault(e, telemetryUseCases)

	go func() {
		if err := e.Start(serveAddress); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server forced to shutdown: ", err)
		}
	}()

	<-time.After(1 * time.Second)
	fmt.Printf("Refresing telemetry data every %s\n", refreshDuration.String())

	<-ctx.Done()
	stop()
	log.Println("Shutting down gracefully, press Ctrl+C again to force")

	if err := e.Shutdown(context.Background()); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server has been shutdown.")
}
