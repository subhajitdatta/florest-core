package servicetest

import (
	gk "github.com/onsi/ginkgo"
	gm "github.com/onsi/gomega"
	testUtil "test/utils"
)

func getHealthCheckTest() {
	apiName := "florest"
	gk.Describe("GET /"+apiName+"/healthcheck", func() {
		request := testUtil.CreateTestRequest("GET", "/"+apiName+"/healthcheck")
		response := GetResponse(request)
		gk.Context("then the response", func() {
			gk.It("should return api health status", func() {
				gm.Expect(response.HeaderMap.Get("Content-Type")).To(gm.Equal("application/json"))
				gm.Expect(response.HeaderMap.Get("Cache-Control")).To(gm.Equal(""))
				gm.Expect(response.Code).To(gm.Equal(200))
				validateHealthCheckResponse(response.Body.String())
			})
		})
	})
}
