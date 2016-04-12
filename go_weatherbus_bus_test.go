package main_test

import (
	"net/http"
	"os/exec"

	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GoWeatherbusBus", func() {
	var sess *gexec.Session

	BeforeEach(func() {
		path, err := gexec.Build("github.com/seattle-beach/go-weatherbus-bus")
		Expect(err).NotTo(HaveOccurred())

		cmd := exec.Command(path)
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
		resp, err := http.Get("http://localhost:9092/api/v1/stops/1_619")
		Expect(err).NotTo(HaveOccurred())
		Expect(resp.StatusCode).To(Equal(200))
	})
})
