package health

// Most simple health check probe to see whether our server is still alive

import (
	"fmt"
	"net/http"

	"github.com/argoproj-labs/argocd-image-updater/pkg/log"
)

func StartHealthServer(port int) chan error {
	errCh := make(chan error)
	go func() {
		http.HandleFunc("/healthz", HealthProbe)
		errCh <- http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	}()
	return errCh
}

func HealthProbe(w http.ResponseWriter, r *http.Request) {
	log.Tracef("/healthz ping request received, replying with pong")
	fmt.Fprintf(w, "OK\n")
}
