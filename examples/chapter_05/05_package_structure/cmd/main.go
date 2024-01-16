package main

import (
	"demo/internal/works"
	"demo/pkg/helpers"
	"log/slog"
	"time"
)

func main() {
	start := time.Now()
	slog.Info("main() started", "at", start.UTC())
	helpers.Example()
	works.Example()
	slog.Warn("package works", "state", works.GetState())
	slog.Info("main() finished", "duration", time.Since(start))
}
