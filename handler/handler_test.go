package handler_test

import (
	"net/http"
	"net/http/httptest"

	. "github.com/seattle-beach/go-weatherbus-bus/handler"
	"github.com/seattle-beach/go-weatherbus-bus/onebusaway/onebusawayfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Handler", func() {
	var (
		handler    http.Handler
		fakeClient *onebusawayfakes.FakeClient
	)

	BeforeEach(func() {
		fakeClient = new(onebusawayfakes.FakeClient)
		handler = NewHandler(fakeClient)
	})

	Describe("/api/v1/stops/{stop_id}", func() {
		It("calls the onebusaway client with the provided stop_id", func() {
			writer := httptest.NewRecorder()
			request, err := http.NewRequest("GET", "/api/v1/stops/real_stoopid", nil)
			Expect(err).NotTo(HaveOccurred())

			handler.ServeHTTP(writer, request)

			Expect(fakeClient.GetStopCallCount()).To(Equal(1))
			stopID := fakeClient.GetStopArgsForCall(0)
			Expect(stopID).To(Equal("real_stoopid"))
		})
	})
})
