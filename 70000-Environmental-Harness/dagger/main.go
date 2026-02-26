package main

import (
	"context"
	"dagger/olympusgcp-observability/internal/dagger"
)

type OlympusGCPObservability struct{}

func (m *OlympusGCPObservability) HelloWorld(ctx context.Context) string {
	return "Hello from OlympusGCP-Observability!"
}

func main() {
	dagger.Serve()
}
