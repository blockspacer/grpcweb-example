// Code generated by reactGen. DO NOT EDIT.

package book

import "myitcv.io/react"

type BookChatElem struct {
	react.Element
}

func (b BookChatDef) ShouldComponentUpdateIntf(nextProps react.Props, prevState, nextState react.State) bool {
	res := false

	{
		res = b.Props() != nextProps.(BookChatProps) || res
	}
	v := prevState.(BookChatState)
	res = !v.EqualsIntf(nextState) || res
	return res
}

func buildBookChat(cd react.ComponentDef) react.Component {
	return BookChatDef{ComponentDef: cd}
}

func buildBookChatElem(props BookChatProps, children ...react.Element) *BookChatElem {
	return &BookChatElem{
		Element: react.CreateElement(buildBookChat, props, children...),
	}
}

// SetState is an auto-generated proxy proxy to update the state for the
// BookChat component.  SetState does not immediately mutate b.State()
// but creates a pending state transition.
func (b BookChatDef) SetState(state BookChatState) {
	b.ComponentDef.SetState(state)
}

// State is an auto-generated proxy to return the current state in use for the
// render of the BookChat component
func (b BookChatDef) State() BookChatState {
	return b.ComponentDef.State().(BookChatState)
}

// IsState is an auto-generated definition so that BookChatState implements
// the myitcv.io/react.State interface.
func (b BookChatState) IsState() {}

var _ react.State = BookChatState{}

// GetInitialStateIntf is an auto-generated proxy to GetInitialState
func (b BookChatDef) GetInitialStateIntf() react.State {
	return BookChatState{}
}

func (b BookChatState) EqualsIntf(val react.State) bool {
	return b == val.(BookChatState)
}

// IsProps is an auto-generated definition so that BookChatProps implements
// the myitcv.io/react.Props interface.
func (b BookChatProps) IsProps() {}

// Props is an auto-generated proxy to the current props of BookChat
func (b BookChatDef) Props() BookChatProps {
	uprops := b.ComponentDef.Props()
	return uprops.(BookChatProps)
}

func (b BookChatProps) EqualsIntf(val react.Props) bool {
	return b == val.(BookChatProps)
}

var _ react.Props = BookChatProps{}
