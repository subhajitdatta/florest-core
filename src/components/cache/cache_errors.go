package cache

import (
	"errors"
	"github.com/jabong/florest-core/src/common/logger"
)

const (
	ErrNoPlatform         = "Platform not found"
	ErrInitialization     = "Initialization failed"
	ErrGetFailure         = "Failure in Get() method"
	ErrSetFailure         = "Failure in Set() method"
	ErrGetBatchFailure    = "Failure in GetBatch() method"
	ErrDeleteFailure      = "Failure in Delete() method"
	ErrDeleteBatchFailure = "Failure in DeleteBatch() method"
	ErrKeyPresent         = "Key is already present"
	ErrKeyNotPresent      = "Key is not present"
	ErrWrongType          = "Incorrect type sent"
)

// getErrObj returns error object with given details
func getErrObj(errCode string, developerMessage string) (ret error) {
	errorString := "ErrorCode: " + errCode + ", developerMessage : " + developerMessage
	logger.Error(errorString)
	ret = errors.New(errorString)
	return ret
}
