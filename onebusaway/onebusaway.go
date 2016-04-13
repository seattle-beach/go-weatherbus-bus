package onebusaway

//go:generate counterfeiter . Client

type Client interface {
	GetStop(stopID string) (Stop, error)
}

type Stop struct {
	StopID    string  `json:"id"`
	Name      string  `json:"name"`
	Lat       float64 `json:"lat"`
	Long      float64 `json:"lon"`
	Direction string  `json:"direction"`
}
