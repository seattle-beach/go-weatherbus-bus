package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/seattle-beach/go-weatherbus-bus/webserver"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	http.HandleFunc("/api/v1/stops/1_619", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world")
	})

	server := webserver.NewWebServer(9092)
	serverErrors, err := server.Start()

	if err != nil {
		log.Fatal(fmt.Sprintf("Error starting server: %s", err.Error()))
		os.Exit(1)
	}

	fmt.Println("Ready")

	select {
	case err := <-serverErrors:
		log.Fatal(err.Error())
		os.Exit(1)
	case <-sigs:
		break
	}

	server.Stop()
}
