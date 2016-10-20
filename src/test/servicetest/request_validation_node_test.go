package servicetest

import (
	"encoding/json"
	"fmt"
	"github.com/jabong/florest-core/src/test/api"
	gk "github.com/onsi/ginkgo"
	gm "github.com/onsi/gomega"
	"io/ioutil"
	testUtil "test/utils"
)

func testRequestValidation() {
	var tstData = new(api.TestData)
	bt, err := ioutil.ReadFile("testdata/requestValidationNode.json")
	if err != nil {
		fmt.Sprintf("Error loading Request Validation Node Data file %s \n %s", err)
	}
	err = json.Unmarshal(bt, tstData)
	apiName := "florest"
	gk.Describe("GET"+"/"+apiName+"/V1/REQVD/", func() {
		request := testUtil.CreateTestRequest("GET", "/"+apiName+"/V1/REQVD/")
		response := GetResponse(request)
		gk.Context("then the response", func() {
			gk.It("should return validation failure message", func() {
				gm.Expect(response.HeaderMap.Get("Content-Type")).To(gm.Equal("application/json"))
				gm.Expect(response.HeaderMap.Get("Cache-Control")).To(gm.Equal(""))
				gm.Expect(response.Code).To(gm.Equal(400))
				validateRequestValidationResponse(response.Body.String(), tstData.RequestValidationMessage)
			})
		})
	})

}
