package webserver

import (
	"fmt"
	"net"
	"net/http"
)

type WebServer interface {
	Start() (chan (error), error)
	Stop() error
}

type webServer struct {
	port     int
	listener net.Listener
}

func NewWebServer(port int) WebServer {
	return &webServer{
		port: port,
	}
}

func (ws *webServer) Start() (chan (error), error) {
	var err error

	ws.listener, err = net.Listen("tcp", fmt.Sprintf(":%d", ws.port))
	if err != nil {
		return nil, err
	}

	serverErr := make(chan error)

	go func() {
		serverErr <- http.Serve(ws.listener, nil)
	}()

	return serverErr, nil
}

func (ws *webServer) Stop() error {
	return ws.listener.Close()
}
