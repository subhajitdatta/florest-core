package servicetest

import (
	"encoding/json"
	utilhttp "github.com/jabong/florest-core/src/common/utils/http"
	gm "github.com/onsi/gomega"
)

func validateHealthCheckResponse(responseBody string) {
	var utilResponse utilhttp.Response
	err := json.Unmarshal([]byte(responseBody), &utilResponse)
	gm.Expect(err).To(gm.BeNil())

	utilResponseData := utilResponse.Data
	if v, ok := utilResponseData.(map[string]interface{}); ok {
		node, helloNodePresent := v["hello world"]
		gm.Expect(helloNodePresent).To(gm.Equal(true))
		if helloNodePresent {
			if body, ok := node.(map[string]interface{}); ok {
				status, statusPresent := body["status"]
				gm.Expect(statusPresent).To(gm.Equal(true))
				gm.Expect(status).To(gm.Equal("success"))
			}
		}
	}
}
