package onebusaway

import (
	"encoding/json"
	"errors"
	"net/http"
)

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

	if response.StatusCode != http.StatusOK {
		return Stop{}, errors.New("Service returned " + response.Status)
	}

	defer response.Body.Close()
	var root stopResponseRoot
	json.NewDecoder(response.Body).Decode(&root)

	if (root.Data == stopData{}) {
		return Stop{}, errors.New("One Bus Away failure")
	}

	return root.Data.Entry, nil
}
