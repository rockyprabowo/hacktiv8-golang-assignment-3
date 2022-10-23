package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())

	flag.StringVar(&serveAddress, "listen", defaultServeAddress, "Listen address")
	flag.DurationVar(&refreshDuration, "refresh-duration", defaultRefreshDuration, "refreshDuration duration in hms format, example: 15s for 15 seconds")
	flag.BoolVar(&debug, "debug", false, "Debug mode")
	flag.Parse()

	if port, ok := os.LookupEnv("PORT"); ok {
		serveAddress = fmt.Sprintf(":%s", port)
	}

	if railwayStaticURL, ok := os.LookupEnv("RAILWAY_STATIC_URL"); ok {
		baseURL = railwayStaticURL
	} else {
		baseURL = serveAddress
	}

	if railwayEnv, ok := os.LookupEnv("RAILWAY_ENVIRONMENT"); ok && railwayEnv == "production" {
		debug = false
	}

	if !debug {
		log.SetOutput(io.Discard)
		log.SetFlags(0)
	}
}
