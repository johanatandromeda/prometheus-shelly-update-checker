package http_proc

import (
	"fmt"
	"github.com/johanatandromeda/prometheus-shelly-update-checker/pkg/shelly"
	"log/slog"
	"net/http"
)

func ShellyUpdateHandler(w http.ResponseWriter, r *http.Request) {

	target := r.URL.Query()["target"]

	if len(target) != 1 {
		slog.Debug(fmt.Sprintf("A single target parameter not supplied"))
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("target parameter is required\n"))
		return
	}

	slog.Debug(fmt.Sprintf("Checking update status for '%s'", target[0]))

	update, status, err := shelly.UpdateNeeded(target[0])

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, "Internal error: %s", err)
		return
	}

	if status != 200 {
		w.WriteHeader(status)
	} else {
		resp := `# HELP shelly_needs_update Whether the Shelly device needs an update
# TYPE shelly_needs_update counter
shelly_needs_update `
		w.WriteHeader(http.StatusOK)
		if update {
			_, _ = fmt.Fprintf(w, "%s 1\n", resp)
		} else {
			_, _ = fmt.Fprintf(w, "%s 0\n", resp)
		}
	}

}
