package orchestrator

import (
	"errors"
)

/*
The state of the workflow nodes is logged in this data structure.
The state is mapped as a key value store
This data structure is internal to the orchestrator
*/
type workFlowState struct {
	value map[string]interface{}
}

/*
Set the value for a  key in the workflow state
Private to the orchestrator package
*/
func (state *workFlowState) set(key string, val interface{}) error {
	//Check if the key is already present
	_, ok := state.value[key]
	if ok {
		return errors.New("Value already set for " + key)
	}
	state.value[key] = val
	return nil
}

/*
Get the value for a key in the workflow state
*/
func (state *workFlowState) Get(key string) (val interface{}, err error) {
	//Check if the key is already present
	res, ok := state.value[key]
	if !ok {
		return "", errors.New("Value not found for " + key)
	}
	return res, nil
}

/*
Get all the values in the workflow state
*/
func (state *workFlowState) GetAll() (res map[string]interface{}) {
	return state.value
}

/*
Initialize the workflow state
*/
func (state *workFlowState) create() {
	state.value = make(map[string]interface{})
}
