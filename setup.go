package main

import (
	"flag"
	"io"
	"log"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())

	flag.StringVar(&serveAddress, "listen", defaultServeAddress, "Listen address")
	flag.DurationVar(&refreshDuration, "refresh-duration", defaultRefreshDuration, "refreshDuration duration in hms format, example: 15s for 15 seconds")
	flag.BoolVar(&debug, "debug", false, "Debug mode")
	flag.Parse()

	if !debug {
		log.SetOutput(io.Discard)
		log.SetFlags(0)
	}
}
