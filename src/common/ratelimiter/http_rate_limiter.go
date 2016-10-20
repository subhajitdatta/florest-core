package ratelimiter

import (
	"encoding/json"
	"fmt"
	"github.com/jabong/florest-core/src/common/constants"
	utilhttp "github.com/jabong/florest-core/src/common/utils/http"
	"net/http"
)

func getAPIResponseFromAppErr(appError *constants.AppError) utilhttp.APIResponse {
	appErrors := new(constants.AppErrors)
	appErrors.Errors = []constants.AppError{*appError}
	status := constants.GetAppHTTPError(*appErrors)
	appResponse := utilhttp.Response{Status: *status}
	jsonBody, _ := json.Marshal(appResponse)

	apiResponse := utilhttp.NewAPIResponse()
	apiResponse.HTTPStatus = appResponse.Status.HTTPStatusCode
	apiResponse.Body = jsonBody

	return apiResponse
}

func getFailedRateLimitResponse(err error) utilhttp.APIResponse {
	appError := &constants.AppError{
		Code:    constants.RateLimiterInternalError,
		Message: err.Error(),
	}
	return getAPIResponseFromAppErr(appError)
}

func getExceededRateLimitResponse(res *RateLimitResult) utilhttp.APIResponse {
	appError := &constants.AppError{
		Code:    constants.RateLimitExceeded,
		Message: fmt.Sprintf("Retry after: %v", res.RetryAfter),
	}
	return getAPIResponseFromAppErr(appError)
}

func MakeRateLimitedHTTPHandler(fn http.HandlerFunc, rl RateLimiter,
	key string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if rl == nil {
			//rate limiter is not initialized
			fn(w, r)
			return
		}

		var apiResponse utilhttp.APIResponse
		exceeded, res, err := rl.RateLimit(key)

		if err != nil {
			//error in rate limiting
			apiResponse = getFailedRateLimitResponse(err)
		} else if exceeded {
			//rate limit has exceeded
			apiResponse = getExceededRateLimitResponse(res)
		} else {
			//exceeded is false i.e rate limit not exceeded
			fn(w, r)
			return
		}

		w.WriteHeader(int(apiResponse.HTTPStatus))
		_, _ = w.Write(apiResponse.Body)
		w.Header().Set("Content-Type", "application/json")
		return
	}
}
