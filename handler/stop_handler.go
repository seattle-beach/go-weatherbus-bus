package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/seattle-beach/go-weatherbus-bus/onebusaway"
)

type stopHandler struct {
	oneBusAwayClient onebusaway.Client
}

type stopData struct {
	Data stop `json:"data"`
}

type stop struct {
	StopID    string  `json:"stopId"`
	Name      string  `json:"name"`
	Lat       float64 `json:"latitude"`
	Long      float64 `json:"longitude"`
	Direction string  `json:"direction"`
}

func NewStopHandler(client onebusaway.Client) http.Handler {
	return &stopHandler{client}
}

func (sh *stopHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pathAry := strings.Split(r.URL.Path, "/")

	obaStop, err := sh.oneBusAwayClient.GetStop(pathAry[len(pathAry)-1])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	presentedStop := getPresentedStop(obaStop)

	w.Header().Add("content-type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(stopData{presentedStop})
}

func getPresentedStop(obaStop onebusaway.Stop) stop {
	return stop{
		StopID:    obaStop.StopID,
		Name:      obaStop.Name,
		Lat:       obaStop.Lat,
		Long:      obaStop.Long,
		Direction: obaStop.Direction,
	}
}
