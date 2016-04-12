package onebusaway

//go:generate counterfeiter . Client

type Client interface {
	GetStop(stopID string) error
}
