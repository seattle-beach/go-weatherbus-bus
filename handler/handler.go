package handler

import (
	"net/http"

	"github.com/seattle-beach/go-weatherbus-bus/onebusaway"
)

func NewHandler(oneBusAwayClient onebusaway.Client) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/api/v1/stops/", NewStopHandler(oneBusAwayClient))
	mux.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	return mux
}
