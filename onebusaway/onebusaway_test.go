package onebusaway_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/seattle-beach/go-weatherbus-bus/onebusaway"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Onebusaway", func() {
	var (
		subject     onebusaway.Client
		ts          *httptest.Server
		handlerFunc http.HandlerFunc
	)

	BeforeEach(func() {
		ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handlerFunc(w, r)
		}))
		subject = onebusaway.NewClient(ts.URL)
	})

	AfterEach(func() {
		defer ts.Close()
	})

	Describe("GetStop", func() {
		var (
			stop         onebusaway.Stop
			err          error
			requestedUrl string
		)

		Context("When the OBA service responds with a stop", func() {
			BeforeEach(func() {
				handlerFunc = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					requestedUrl = r.URL.String()
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(`
						{
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
						      }
						`))
				})
				stop, err = subject.GetStop("1_619")
			})

			It("should request the stop from the web service", func() {
				Expect(requestedUrl).To(Equal("/api/where/stop/1_619.json?key=test"))
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
