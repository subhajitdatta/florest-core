package servicetest

import (
	"encoding/json"
	"github.com/jabong/florest-core/src/common/constants"
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

func validateRequestValidationResponse(responseBody string, errmsg []string) {
	var utilResponse utilhttp.Response
	_ = json.Unmarshal([]byte(responseBody), &utilResponse)
	errleng := len(errmsg)
	m := make(map[string]string, errleng)
	for i := 0; i < errleng; i++ {
		temp := errmsg[i]
		m[temp] = temp
	}
	var leng int = len(utilResponse.Status.Errors)
	for i := 0; i < leng; i++ {
		var err constants.AppError = utilResponse.Status.Errors[i]
		_, found := m[err.Message]
		gm.Expect(true).To(gm.Equal(found))
	}
}
