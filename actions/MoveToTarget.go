package actions

import (
	"fmt"
	b3 "github.com/dyc92/behavior3go"
	. "github.com/dyc92/behavior3go/config"
	. "github.com/dyc92/behavior3go/core"
)

type MoveToTarget struct {
	Action
}

func (this *MoveToTarget) Initialize(setting *BTNodeCfg) {

	this.Action.Initialize(setting)
}

func (this *MoveToTarget) OnTick(tick *Tick) b3.Status {
	fmt.Println("MoveToTarget")
	unit := tick.Blackboard.GetMem("Unit")
	if unit == nil {

		return b3.ERROR
	}
	return b3.SUCCESS
}
