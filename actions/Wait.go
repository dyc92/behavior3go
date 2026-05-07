package actions

import (
	"math/rand"
	"time"

	b3 "github.com/dyc92/behavior3go"
	. "github.com/dyc92/behavior3go/config"
	. "github.com/dyc92/behavior3go/core"
)

type Wait struct {
	Action
	timeSec   int64
	randomSec int64
}

func (n *Wait) Initialize(setting *BTNodeCfg) {
	n.Action.Initialize(setting)
	n.timeSec = setting.GetPropertyAsInt64("time")
	n.randomSec = setting.GetPropertyAsInt64("random")
}

func (n *Wait) OnOpen(tick *Tick) {
	now := time.Now().UnixMilli()
	wait := n.timeSec * 1000
	if n.randomSec > 0 {
		wait += rand.Int63n(n.randomSec * 1000)
	}
	tick.Blackboard.Set("wait_end_time", now+wait, tick.GetTree().GetID(), n.GetID())
}

func (n *Wait) OnTick(tick *Tick) b3.Status {
	endTime := tick.Blackboard.GetInt64("wait_end_time", tick.GetTree().GetID(), n.GetID())
	if time.Now().UnixMilli() >= endTime {
		return b3.SUCCESS
	}
	return b3.RUNNING
}
