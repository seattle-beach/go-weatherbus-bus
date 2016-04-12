package handler

import (
	"net/http"

	"github.com/seattle-beach/go-weatherbus-bus/onebusaway"
)

func NewHandler(oneBusAwayClient onebusaway.Client) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/api/v1/stops/", NewStopHandler(oneBusAwayClient))

	return mux
}
