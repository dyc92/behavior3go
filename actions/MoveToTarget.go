package actions

import (
	b3 "github.com/magicsea/behavior3go"
	. "github.com/magicsea/behavior3go/config"
	. "github.com/magicsea/behavior3go/core"
)

type MoveToTarget struct {
	Action
}

func (this *MoveToTarget) Initialize(setting *BTNodeCfg) {

	this.Action.Initialize(setting)
}

func (this *MoveToTarget) OnTick(tick *Tick) b3.Status {

	unit := tick.Blackboard.GetMem("Unit")
	if unit == nil {

		return b3.ERROR
	}
	return b3.SUCCESS
}
