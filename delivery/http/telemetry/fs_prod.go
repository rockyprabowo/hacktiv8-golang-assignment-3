//go:build prod
// +build prod

package telemetry_delivery_http

import (
	"embed"
	"io/fs"
	"log"

	"github.com/labstack/echo/v4"
)

//go:embed all:ui/dist
var dist embed.FS

func getFrontendAssets() fs.FS {
	log.Println("frontendAssets: using embedded filesystem")
	return echo.MustSubFS(dist, "ui/dist")
}
