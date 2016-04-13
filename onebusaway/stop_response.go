package onebusaway

type stopResponseRoot struct {
	Data stopData `json:"data"`
}

type stopData struct {
	Entry Stop `json:"entry"`
}
