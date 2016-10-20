package servicetest

import (
	"testing"

	gk "github.com/onsi/ginkgo"
	gm "github.com/onsi/gomega"
)

func TestSearch(t *testing.T) {
	gm.RegisterFailHandler(gk.Fail)
	gk.RunSpecs(t, "Service Suite")
}

var _ = gk.Describe("Starting scenario: \n", func() {
	InitializeTestService()
	getHealthCheckTest()
	testRequestValidation()
	sqldbTest()
	mongodbTest()
	cacheTest()
	loggerTest()
})
