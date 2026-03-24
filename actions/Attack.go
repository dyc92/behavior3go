package actions

import (
	b3 "github.com/magicsea/behavior3go"
	. "github.com/magicsea/behavior3go/config"
	. "github.com/magicsea/behavior3go/core"
)

type Attack struct {
	Action
}

func (this *Attack) Initialize(setting *BTNodeCfg) {
	this.Action.Initialize(setting)
}

func (this *Attack) OnTick(tick *Tick) b3.Status {

	unit := tick.Blackboard.GetMem("Unit")
	if unit == nil {

		return b3.ERROR
	}
	return b3.SUCCESS
}
