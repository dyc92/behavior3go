package actions

import (
	b3 "github.com/dyc92/behavior3go"
	. "github.com/dyc92/behavior3go/config"
	. "github.com/dyc92/behavior3go/core"
)

type Error struct {
	Action
}

func (this *Error) Initialize(setting *BTNodeCfg) {
	this.Action.Initialize(setting)
}

func (this *Error) OnTick(tick *Tick) b3.Status {
	return b3.ERROR
}
