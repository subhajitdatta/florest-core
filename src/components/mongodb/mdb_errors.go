package mongodb

import "fmt"

type MDBError struct {
	ErrCode          string
	DeveloperMessage string
}

func (e MDBError) Error() string {
	return fmt.Sprintf("{ErrCode:%s,DeveloperMessage:%s}", e.ErrCode, e.DeveloperMessage)
}

const (
	ErrInitialization = "Initialization failed"
	ErrFindOneFailure = "Failure in FindOne() method"
	ErrFindAllFailure = "Failure in FindAll() method"
	ErrInsertFailure  = "Failure in Insert() method"
	ErrUpdateFailure  = "Failure in Update() method"
	ErrUpsertFailure  = "Failure in Upsert() method"
	ErrRemoveFailure  = "Failure in Remove() method"
	ErrKeyPresent     = "Key is already present"
	ErrKeyNotPresent  = "Key is not present"
	ErrWrongType      = "Incorrect type sent"
)

// getErrObj returns error object with given details
func getErrObj(errCode string, developerMessage string) *MDBError {
	ret := new(MDBError)
	ret.ErrCode = errCode
	ret.DeveloperMessage = developerMessage
	return ret
}
