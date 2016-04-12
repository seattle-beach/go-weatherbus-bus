package onebusaway_test

import (
	"net/http"

	"github.com/onsi/gomega/ghttp"
	"github.com/seattle-beach/go-weatherbus-bus/onebusaway"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Onebusaway", func() {
	var (
		subject onebusaway.Client
		ts      *ghttp.Server
	)

	BeforeEach(func() {
		ts = ghttp.NewServer()
		subject = onebusaway.NewClient(ts.URL())
	})

	AfterEach(func() {
		defer ts.Close()
	})

	Describe("GetStop", func() {
		var (
			stop       onebusaway.Stop
			err        error
			statusCode int
			json       string
		)

		JustBeforeEach(func() {
			ts.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/api/where/stop/1_619.json", "key=test"),
					ghttp.RespondWith(statusCode, json),
				),
			)
			stop, err = subject.GetStop("1_619")
		})

		Context("When the OBA service returns an non 200 status", func() {
			BeforeEach(func() {
				statusCode = http.StatusTeapot
				json = ""
			})

			It("should return an error", func() {
				Expect(err).To(HaveOccurred())
			})
		})

		Context("When the OBA service returns a wrapped error", func() {
			BeforeEach(func() {
				statusCode = http.StatusOK
				json = `{"code":404,"currentTime":1460497519536,"data":null,"text":"No such stop: 1_1","version":1}`
			})

			It("should return an error", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("One Bus Away failure"))
			})
		})

		Context("When the OBA service responds with a stop", func() {
			BeforeEach(func() {
				statusCode = http.StatusOK
				json = `{
	        "code": 200,
	        "currentTime": 1460401388377,
	        "data": {
	          "entry": {
	            "code": "619",
	            "direction": "N",
	            "id": "1_619",
	            "lat": 47.599827,
	            "locationType": 0,
	            "lon": -122.328972,
	            "name": "4th Ave S & S Jackson St",
	            "routeIds": [
	              "1_100229",
	              "1_100044"
	            ],
	            "wheelchairBoarding": "UNKNOWN"
	          }
	        },
	        "text": "OK",
	        "version": 2
	      }`
			})

			It("should return that stop", func() {
				Expect(err).NotTo(HaveOccurred())
				Expect(stop.StopID).To(Equal("1_619"))
				Expect(stop.Name).To(Equal("4th Ave S & S Jackson St"))
				Expect(stop.Lat).To(Equal(47.599827))
				Expect(stop.Long).To(Equal(-122.328972))
				Expect(stop.Direction).To(Equal("N"))
			})
		})
	})
})
