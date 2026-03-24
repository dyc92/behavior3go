package actions

import (
	b3 "github.com/magicsea/behavior3go"
	. "github.com/magicsea/behavior3go/config"
	. "github.com/magicsea/behavior3go/core"
)

type JustFailer struct {
	Action
}

func (this *JustFailer) Initialize(setting *BTNodeCfg) {
	this.Action.Initialize(setting)
}

func (this *JustFailer) OnTick(tick *Tick) b3.Status {
	return b3.FAILURE
}
