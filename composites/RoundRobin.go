package composites

import (
	b3 "github.com/dyc92/behavior3go"
	. "github.com/dyc92/behavior3go/core"
)

const roundRobinChildKey = "currentChild"

type RoundRobin struct {
	Composite
}

func (this *RoundRobin) OnTick(tick *Tick) b3.Status {
	childCount := this.GetChildCount()
	if childCount == 0 {
		return b3.FAILURE
	}

	child := tick.Blackboard.GetInt(roundRobinChildKey, tick.GetTree().GetID(), this.GetID())
	if child < 0 || child >= childCount {
		child = 0
	}

	status := this.GetChild(child).Execute(tick)
	if status == b3.SUCCESS {
		tick.Blackboard.Set(roundRobinChildKey, (child+1)%childCount, tick.GetTree().GetID(), this.GetID())
	}
	return status
}
