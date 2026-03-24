package actions

import (
	b3 "github.com/magicsea/behavior3go"
	. "github.com/magicsea/behavior3go/config"
	. "github.com/magicsea/behavior3go/core"
)

type Succeeder struct {
	Action
}

func (this *Succeeder) Initialize(setting *BTNodeCfg) {
	this.Action.Initialize(setting)
}

func (this *Succeeder) OnTick(tick *Tick) b3.Status {
	return b3.SUCCESS
}
