package onebusaway_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestOnebusaway(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Onebusaway Suite")
}
