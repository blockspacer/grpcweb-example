// Code generated by reactGen. DO NOT EDIT.

package book

import "myitcv.io/react"

func (g *GetBookDef) ShouldComponentUpdateIntf(nextProps, prevState, nextState interface{}) bool {
	res := false

	v := prevState.(GetBookState)
	res = !v.EqualsIntf(nextState) || res
	return res
}

// SetState is an auto-generated proxy proxy to update the state for the
// GetBook component.  SetState does not immediately mutate g.State()
// but creates a pending state transition.
func (g *GetBookDef) SetState(state GetBookState) {
	g.ComponentDef.SetState(state)
}

// State is an auto-generated proxy to return the current state in use for the
// render of the GetBook component
func (g *GetBookDef) State() GetBookState {
	return g.ComponentDef.State().(GetBookState)
}

// IsState is an auto-generated definition so that GetBookState implements
// the myitcv.io/react.State interface.
func (g GetBookState) IsState() {}

var _ react.State = GetBookState{}

// GetInitialStateIntf is an auto-generated proxy to GetInitialState
func (g *GetBookDef) GetInitialStateIntf() react.State {
	return GetBookState{}
}

func (g GetBookState) EqualsIntf(val interface{}) bool {
	return g == val.(GetBookState)
}