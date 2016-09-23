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

	if _, ok := utilResponse.Data.(map[string]interface{}); ok {
		gm.Expect(ok).To(gm.Equal(true))
	}
}
