package loader

import (
	_ "fmt"
	. "github.com/dyc92/behavior3go/actions"
	. "github.com/dyc92/behavior3go/composites"
	. "github.com/dyc92/behavior3go/config"
	. "github.com/dyc92/behavior3go/core"
	. "github.com/dyc92/behavior3go/decorators"
	_ "reflect"

	b3 "github.com/dyc92/behavior3go"
)

func createBaseStructMaps() *b3.RegisterStructMaps {
	st := b3.NewRegisterStructMaps()
	//actions
	st.Register("Error", &Error{})
	st.Register("JustFailer", &JustFailer{})
	st.Register("Runner", &Runner{})
	st.Register("Succeeder", &Succeeder{})
	st.Register("Wait", &Wait{})
	st.Register("Log", &Log{})
	st.Register("TriggerEvent", &TriggerEvent{})
	st.Register("WaitForEvent", &WaitForEvent{})
	//composites
	st.Register("MemPriority", &MemPriority{})
	st.Register("MemSequence", &MemSequence{})
	st.Register("Priority", &Priority{})
	st.Register("RetryUntilFailure", &RetryUntilFailure{})
	st.Register("Sequence", &Sequence{})
	st.Register("IfElse", &IfElse{})

	//decorators
	st.Register("Inverter", &Inverter{})
	st.Register("Repeater", &Repeater{})
	st.Register("RepeatUntilFailure", &RepeatUntilFailure{})
	st.Register("RepeatUntilSuccess", &RepeatUntilSuccess{})
	st.Register("Once", &Once{})

	return st
}

func CreateBevTreeFromConfig(config *BTTreeCfg, extMap *b3.RegisterStructMaps) *BehaviorTree {
	baseMaps := createBaseStructMaps()
	tree := NewBeTree()

	tree.Load(config, baseMaps, extMap)
	return tree
}
