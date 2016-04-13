package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/seattle-beach/go-weatherbus-bus/onebusaway"
)

func NewHandler(oneBusAwayClient onebusaway.Client) http.Handler {
	router := mux.NewRouter()
	router.Handle("/api/v1/stops/{stop_id}", NewStopHandler(oneBusAwayClient))
	router.Handle("/api/v1/stops", NewStopsForLocationHandler(oneBusAwayClient))
	return router
}
