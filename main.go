package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/seattle-beach/go-weatherbus-bus/handler"
	"github.com/seattle-beach/go-weatherbus-bus/onebusaway"
	"github.com/seattle-beach/go-weatherbus-bus/webserver"
)

func main() {
	var baseUrl string

	if len(os.Args) == 2 {
		baseUrl = os.Args[1]
	} else {
		baseUrl = "http://api.pugetsound.onebusaway.org"
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	oba := onebusaway.NewClient(baseUrl)
	server := webserver.NewWebServer(9092, handler.NewHandler(oba))
	serverErrors, err := server.Start()

	if err != nil {
		log.Fatal(fmt.Sprintf("Error starting server: %s", err.Error()))
		os.Exit(1)
	}

	defer server.Stop()

	fmt.Println("Ready")

	select {
	case err := <-serverErrors:
		log.Fatal(err.Error())
		os.Exit(1)
	case <-sigs:
	}
}
