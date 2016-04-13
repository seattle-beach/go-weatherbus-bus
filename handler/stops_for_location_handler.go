package handler

import (
	"net/http"

	"github.com/seattle-beach/go-weatherbus-bus/onebusaway"
)

type stopsForLocationHandler struct {
}

func NewStopsForLocationHandler(client onebusaway.Client) http.Handler {
	return nil
}

//
// func (sh *stopsForLocationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// }
