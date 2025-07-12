package main

import (
	"flag"
	"fmt"
	"github.com/johanatandromeda/prometheus-shelly-update-checker/pkg/http_proc"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log/slog"
	"net/http"
	"os"
)

var version = ""

func main() {

	var programLevel = new(slog.LevelVar) // Info by default
	h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: programLevel})
	slog.SetDefault(slog.New(h))

	slog.Info(fmt.Sprintf("Starting prometheus-shelly-update-checker V%s", version))

	debug := flag.Bool("d", false, "Debug")
	port := flag.Uint("port", 9111, "Port to expose the metrics on")
	host := flag.String("host", "", "Host to bind to")
	flag.Parse()

	if *debug {
		programLevel.Set(slog.LevelDebug)
	}

	slog.Debug(fmt.Sprintf("Starting server on %s:%d", *host, *port))
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/probe", http_proc.ShellyUpdateHandler)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", *host, *port), nil)
	if err != nil {
		panic(err)
	}

}
