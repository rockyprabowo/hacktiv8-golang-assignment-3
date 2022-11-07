//go:build prod
// +build prod

package main

import (
	"log"
	"net/http"
	"regexp"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func allowOriginFunc(origin string) (bool, error) {
	return regexp.MatchString(`^https?://`+serveAddress+`$`, origin)
}

func setupCORS(engine *echo.Echo) {
	log.Println("cors: allow " + serveAddress + " only")
	engine.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOriginFunc: allowOriginFunc,
		AllowMethods:    []string{http.MethodGet},
	}))
}
