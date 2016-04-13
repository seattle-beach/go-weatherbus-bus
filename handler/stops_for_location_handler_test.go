package handler_test

import (
	"net/http"

	. "github.com/seattle-beach/go-weatherbus-bus/handler"

	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
)

var _ = Describe("Handler/StopsForLocationHandler", func() {
	var subject http.Handler

	BeforeEach(func() {
		subject = NewStopsForLocationHandler(nil)
	})
})
