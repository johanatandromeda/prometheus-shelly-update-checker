package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
)

var version = ""

func main() {

	var programLevel = new(slog.LevelVar) // Info by default
	h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: programLevel})
	slog.SetDefault(slog.New(h))

	slog.Info(fmt.Sprintf("Starting prometheus-shelly-update-checker V%s", version))

	debug := flag.Bool("d", false, "Debug")
	flag.Parse()

	if *debug {
		programLevel.Set(slog.LevelDebug)
	}

}
