package decorators

import (
	b3 "github.com/dyc92/behavior3go"
	. "github.com/dyc92/behavior3go/config"
	. "github.com/dyc92/behavior3go/core"
)

type AlwaysRunning struct {
	Decorator
}

func (n *AlwaysRunning) Initialize(setting *BTNodeCfg) {
	n.Decorator.Initialize(setting)
}

func (n *AlwaysRunning) OnTick(tick *Tick) b3.Status {
	if n.GetChildCount() > 0 && n.GetChild(0) != nil {
		n.GetChild(0).Execute(tick)
	}
	return b3.RUNNING
}
