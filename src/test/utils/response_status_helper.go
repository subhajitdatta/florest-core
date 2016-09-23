package utils

import (
	"github.com/jabong/florest-core/src/common/constants"
	utilhttp "github.com/jabong/florest-core/src/common/utils/http"
	gm "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

func MatchHeaderStatus(responseRecorder *httptest.ResponseRecorder, httpCode int) {
	gm.Expect(responseRecorder.HeaderMap.Get("Content-Type")).To(gm.Equal("application/json"))
	gm.Expect(responseRecorder.HeaderMap.Get("Cache-Control")).To(gm.Equal("public, max-age=300"))
	gm.Expect(responseRecorder.Code).To(gm.Equal(httpCode))
}

//MatchSuccessResponseStatus verifies status for successful response
func MatchSuccessResponseStatus(responseBody *utilhttp.Response) {
	gm.Expect(responseBody.Status.HTTPStatusCode).To(gm.Equal(constants.HTTPCode(200)))
	gm.Expect(responseBody.Status.Success).To(gm.Equal(true))
}

//MatchVersionableNotFound
func MatchVersionableNotFound(responseBody *utilhttp.Response) {
	gm.Expect(responseBody.Status.HTTPStatusCode).To(gm.Equal(constants.HTTPCode(http.StatusNotFound)))
	gm.Expect(responseBody.Status.Errors[0].Code).To(gm.Equal(constants.APPErrorCode(1601)))
	gm.Expect(responseBody.Status.Errors[0].Message).To(gm.Equal("Versionable not found in version manager"))
	//TODO
	//gm.Expect(responeBody.DebugData).To(gm.Equal(""))
}

//MatchNotFoundResponseStatus verifies status for Not Found response
func MatchNotFoundResponseStatus(responseBody *utilhttp.Response) {
	gm.Expect(responseBody.Status.HTTPStatusCode).To(gm.Equal(constants.HTTPCode(404)))
	gm.Expect(responseBody.Status.Success).To(gm.Equal(false))
}
