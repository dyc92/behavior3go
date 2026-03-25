package actions

import (
	b3 "github.com/dyc92/behavior3go"
	. "github.com/dyc92/behavior3go/config"
	. "github.com/dyc92/behavior3go/core"
)

type Runner struct {
	Action
}

func (this *Runner) Initialize(setting *BTNodeCfg) {
	this.Action.Initialize(setting)
}

func (this *Runner) OnTick(tick *Tick) b3.Status {
	return b3.RUNNING
}
