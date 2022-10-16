package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Telemetry struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

type TelemetryResponse struct {
	Status Telemetry `json:"status"`
}

func main() {
	var (
		telemetry    Telemetry
		serveAddress = "localhost:8080"
		timeout      = 15 * time.Second
		mux          = http.NewServeMux()
	)

	rand.Seed(time.Now().UnixNano())

	go TelemetryDataFetcher(&telemetry, timeout)

	mux.HandleFunc("/api/telemetry", TelemetryDataHandler(&telemetry))

	log.Println("Listening and serving on " + serveAddress)
	err := http.ListenAndServe(serveAddress, mux)
	if err != nil {
		log.Fatal(err)
	}
}

func TelemetryDataGenerator(c chan Telemetry, timeout time.Duration) {
	for {
		c <- Telemetry{
			Water: rand.Intn(100),
			Wind:  rand.Intn(100),
		}
		time.Sleep(timeout)
	}
}

func TelemetryDataFetcher(currentData *Telemetry, timeout time.Duration) {
	c := make(chan Telemetry)

	go TelemetryDataGenerator(c, timeout)
	for {
		*currentData = <-c
		log.Printf("new telemetry data recieved: %+v", currentData)
	}
}

func TelemetryDataHandler(telemetry *Telemetry) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		response := TelemetryResponse{
			Status: *telemetry,
		}

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			writeErrorResponse(writer, err)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		_, err = writer.Write(jsonResponse)
		if err != nil {
			writeErrorResponse(writer, err)
			return
		}
	}
}

func writeErrorResponse(writer http.ResponseWriter, err error) {
	code := http.StatusInternalServerError
	errorResponse, _ := json.Marshal(map[string]any{
		"code":  code,
		"error": err.Error(),
	})
	writer.WriteHeader(code)
	_, _ = writer.Write(errorResponse)
}
