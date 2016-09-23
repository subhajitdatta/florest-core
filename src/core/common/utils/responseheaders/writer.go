package responseheaders

import (
	"fmt"
	"github.com/jabong/florest-core/src/common/config"
	"github.com/jabong/florest-core/src/common/constants"
	utilhttp "github.com/jabong/florest-core/src/common/utils/http"
	workflow "github.com/jabong/florest-core/src/core/common/orchestrator"
	"strings"
)

const (
	publicResponse  string = "public"
	privateResponse string = "private"
	noCache         string = "no-cache"
	noStore         string = "no-store"

	//http headers
	cacheControl string = "Cache-Control"
	contentType  string = "Content-Type"
	maxAge       string = "max-age"
)

//IdentifierExecutor joins the result retrieved from multiple nodes
type Writer struct {
	id string
}

func (r *Writer) SetID(id string) {
	r.id = id
}

func (r *Writer) GetID() (id string, err error) {
	return r.id, nil
}

func (r *Writer) Name() string {
	return "Response Header Writer"
}

func (r *Writer) Execute(io workflow.WorkFlowData) (workflow.WorkFlowData, error) {
	rh, _ := io.IOData.Get(constants.ResponseHeadersConfig)
	res := utilhttp.NewAPIResponse()
	if rh != nil {
		if resHeaderConf, ok := rh.(config.ResponseHeaderFields); ok {
			cHeaders := resHeaderConf.CacheControl
			res.Headers[cacheControl] = getCacheControlHeader(cHeaders)
		}
	}

	res.Headers[contentType] = "application/json"
	io.IOData.Set(constants.APIResponse, res)
	return io, nil
}

func getCacheControlHeader(cHeaders config.CacheControlHeaders) string {
	c := ""

	var params []string

	if cHeaders.ResponseType == publicResponse {
		params = append(params, publicResponse)
	} else if cHeaders.ResponseType == privateResponse {
		params = append(params, privateResponse)
	}

	if cHeaders.NoCache {
		params = append(params, noCache)
	}

	if cHeaders.NoStore {
		params = append(params, noStore)
	}

	if cHeaders.MaxAgeInSeconds > 0 {
		params = append(params, fmt.Sprintf("%s=%d", maxAge, cHeaders.MaxAgeInSeconds))
	}

	if len(params) > 0 {
		c = strings.Join(params, ", ")
	}
	return c
}
