//go:build !prod
// +build !prod

package main

import (
	"log"
	"net/http"
	"regexp"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func allowOriginFunc(origin string) (bool, error) {
	// TODO: Only allow localhost with optional port for now
	return regexp.MatchString(`^https?://(?:localhost|127(?:\.\d{1,3}){3})(?::\d+)?$`, origin)
}

func setupCORS(engine *echo.Echo) {
	log.Println("cors: allow localhost only")
	engine.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOriginFunc: allowOriginFunc,
		AllowMethods:    []string{http.MethodGet},
	}))
}
