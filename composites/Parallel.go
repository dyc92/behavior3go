package composites

import (
	b3 "github.com/dyc92/behavior3go"
	. "github.com/dyc92/behavior3go/config"
	. "github.com/dyc92/behavior3go/core"
)

type Parallel struct {
	Composite
}

func (n *Parallel) Initialize(setting *BTNodeCfg) {
	n.Composite.Initialize(setting)
}

func (n *Parallel) OnTick(tick *Tick) b3.Status {
	allSuccess := true
	for i := 0; i < n.GetChildCount(); i++ {
		status := n.GetChild(i).Execute(tick)
		if status == b3.ERROR {
			return b3.ERROR
		}
		if status != b3.SUCCESS {
			allSuccess = false
		}
	}
	if allSuccess {
		return b3.SUCCESS
	}
	return b3.RUNNING
}
