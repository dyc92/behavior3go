package actions

import (
	"fmt"
	b3 "github.com/dyc92/behavior3go"
	. "github.com/dyc92/behavior3go/config"
	. "github.com/dyc92/behavior3go/core"
)

type Attack struct {
	Action
}

func (this *Attack) Initialize(setting *BTNodeCfg) {
	this.Action.Initialize(setting)
}

func (this *Attack) OnTick(tick *Tick) b3.Status {

	fmt.Println("Attack")
	unit := tick.Blackboard.GetMem("Unit")
	if unit == nil {

		return b3.ERROR
	}
	return b3.SUCCESS
}
