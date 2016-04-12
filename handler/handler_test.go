package handler_test

import (
	"errors"
	"net/http"
	"net/http/httptest"

	"github.com/seattle-beach/go-weatherbus-bus/onebusaway"
	"github.com/seattle-beach/go-weatherbus-bus/onebusaway/onebusawayfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/seattle-beach/go-weatherbus-bus/handler"
)

var _ = Describe("Handler", func() {
	var (
		handler    http.Handler
		fakeClient *onebusawayfakes.FakeClient
		writer     *httptest.ResponseRecorder
		request    *http.Request
	)

	BeforeEach(func() {
		var err error
		fakeClient = new(onebusawayfakes.FakeClient)
		handler = NewHandler(fakeClient)

		writer = httptest.NewRecorder()
		request, err = http.NewRequest("GET", "/api/v1/stops/real_stoopid", nil)
		Expect(err).NotTo(HaveOccurred())
	})

	JustBeforeEach(func() {
		handler.ServeHTTP(writer, request)
	})

	Describe("/api/v1/stops/{stop_id}", func() {
		It("calls the onebusaway client with the provided stop_id", func() {
			Expect(fakeClient.GetStopCallCount()).To(Equal(1))
			stopID := fakeClient.GetStopArgsForCall(0)
			Expect(stopID).To(Equal("real_stoopid"))
		})

		Context("when the call succeeds", func() {
			BeforeEach(func() {
				stop := onebusaway.Stop{
					StopID:    "1_619",
					Name:      "4th Ave S & S Jackson S",
					Lat:       47.599827,
					Long:      -122.328972,
					Direction: "N",
				}

				fakeClient.GetStopReturns(stop, nil)
			})

			It("writes a Status OK", func() {
				Expect(writer.Code).To(Equal(http.StatusOK))
			})

			It("sets the correct content-type header", func() {
				Expect(writer.Header().Get("content-type")).To(Equal("application/json; charset=utf-8"))
			})

			It("returns the stop data", func() {
				Expect(writer.Body.String()).To(MatchJSON(`{
	        "data": {
	          "stopId" :"1_619",
	          "name": "4th Ave S \u0026 S Jackson S",
	          "latitude": 47.599827,
	          "longitude": -122.328972,
	          "direction": "N"
	      	}
				}`))
			})
		})

		Context("when the call returns a StopNotFoundError", func() {
			BeforeEach(func() {
				fakeClient.GetStopReturns(onebusaway.Stop{}, onebusaway.StopNotFoundError)
			})

			It("returns a 404 error", func() {
				Expect(writer.Code).To(Equal(http.StatusNotFound))
			})
		})

		Context("when the call returns an unknown error", func() {
			BeforeEach(func() {
				fakeClient.GetStopReturns(onebusaway.Stop{}, errors.New("wtf"))
			})

			It("returns a 500 error", func() {
				Expect(writer.Code).To(Equal(http.StatusInternalServerError))
			})
		})
	})
})
