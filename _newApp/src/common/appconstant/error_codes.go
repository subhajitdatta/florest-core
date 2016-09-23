package appconstant

import (
	florest_Constant "github.com/jabong/florest-core/src/common/constants"
)

const (
	InconsistantDataStateErrorCode       florest_Constant.APPErrorCode = 1407
	FunctionalityNotImplementedErrorCode florest_Constant.APPErrorCode = 1408
)

const (
	HttpStatusNotImplementedErrorCode florest_Constant.HTTPCode = 501
)

var APPErrorCodeToHTTPCodeMap = map[florest_Constant.APPErrorCode]florest_Constant.HTTPCode{
	InconsistantDataStateErrorCode:       florest_Constant.HTTPStatusInternalServerErrorCode,
	FunctionalityNotImplementedErrorCode: HttpStatusNotImplementedErrorCode,
}
