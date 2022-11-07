//go:build !prod
// +build !prod

package telemetry_delivery_http

import (
	"io/fs"
	"log"
	"os"
)

func getFrontendAssets() fs.FS {
	log.Println("frontendAssets: using local filesystem")
	return os.DirFS("delivery/http/telemetry/ui/dist")
}
