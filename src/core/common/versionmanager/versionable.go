package versionmanager

/*
Any executable that wants to be plugged into version manager should implement this interface
*/
type Versionable interface {
	GetInstance() interface{}
}
