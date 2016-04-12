package handler

import (
	"net/http"
	"strings"

	"github.com/seattle-beach/go-weatherbus-bus/onebusaway"
)

type stopHandler struct {
	oneBusAwayClient onebusaway.Client
}

func NewStopHandler(client onebusaway.Client) http.Handler {
	return &stopHandler{client}
}

func (sh *stopHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pathAry := strings.Split(r.URL.Path, "/")

	sh.oneBusAwayClient.GetStop(pathAry[len(pathAry)-1])
}
