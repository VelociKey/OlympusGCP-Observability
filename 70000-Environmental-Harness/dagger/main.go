package main

import (
	"context"
	"olympus.fleet/00SDLC/OlympusForge/70000-Environmental-Harness/dagger/olympusgcp-observability/internal/dagger"
)

type OlympusGCPObservability struct{}

func (m *OlympusGCPObservability) HelloWorld(ctx context.Context) string {
	return "Hello from OlympusGCP-Observability!"
}

func main() {
	dagger.Serve()
}
