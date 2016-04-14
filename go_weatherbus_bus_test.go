package main_test

import (
	"io/ioutil"
	"net/http"
	"os/exec"

	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
	"github.com/onsi/gomega/ghttp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GoWeatherbusBus", func() {
	var (
		sess *gexec.Session
		ts   *ghttp.Server
	)

	BeforeEach(func() {
		path, err := gexec.Build("github.com/seattle-beach/go-weatherbus-bus")
		Expect(err).NotTo(HaveOccurred())

		ts = ghttp.NewServer()

		cmd := exec.Command(path, ts.URL())
		sess, err = gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		Eventually(sess.Out).Should(gbytes.Say("Ready"))
		Expect(sess).ShouldNot(gexec.Exit())
	})

	AfterEach(func() {
		sess.Terminate()
		Eventually(sess).Should(gexec.Exit(0))
	})

	It("should respond to /api/where/stop/[stopid]", func() {
		json := `{
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
		ts.AppendHandlers(
			ghttp.CombineHandlers(
				ghttp.RespondWith(http.StatusOK, json),
			),
		)

		resp, err := http.Get("http://localhost:9090/api/v1/stops/1_619")
		Expect(err).NotTo(HaveOccurred())
		Expect(resp.StatusCode).To(Equal(http.StatusOK))
		bytes, err := ioutil.ReadAll(resp.Body)
		Expect(err).NotTo(HaveOccurred())
		Expect(string(bytes)).To(MatchJSON(`{"data":{"stopId":"1_619",
			"name":"4th Ave S \u0026 S Jackson St","latitude":47.599827,
			"longitude":-122.328972,"direction":"N"}}`))
	})
})
