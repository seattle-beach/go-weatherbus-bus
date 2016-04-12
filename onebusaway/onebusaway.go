package onebusaway

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

//go:generate counterfeiter . Client

type Client interface {
	GetStop(stopID string) (Stop, error)
}

var StopNotFoundError = errors.New("stop-not-found")

type stopResponseRoot struct {
	Data stopData `json:"data"`
}

type stopData struct {
	Entry Stop `json:"entry"`
}

type Stop struct {
	StopID    string  `json:"id"`
	Name      string  `json:"name"`
	Lat       float64 `json:"lat"`
	Long      float64 `json:"lon"`
	Direction string  `json:"direction"`
}

type client struct {
	baseUrl string
}

func NewClient(baseUrl string) Client {
	return &client{baseUrl}
}

func (c *client) GetStop(stopID string) (Stop, error) {
	url := c.baseUrl + "/api/where/stop/" + stopID + ".json?key=test"
	response, err := http.Get(url)

	if err != nil {
		return Stop{}, err
	}

	defer response.Body.Close()
	decoder := json.NewDecoder(response.Body)
	var root stopResponseRoot
	err = decoder.Decode(&root)

	if err != nil {
		fmt.Printf("Decode ERROR!!!!!! %s\n", err.Error())
	}

	return root.Data.Entry, nil
}
