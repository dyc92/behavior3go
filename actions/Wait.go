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
	isLoop    bool
}

func (n *Wait) Initialize(setting *BTNodeCfg) {
	n.Action.Initialize(setting)
	timeSec, err := ParseArgToNumber[int64](setting.Args, "time")
	if err != nil {
		n.timeSec = 0
	} else {
		n.timeSec = timeSec
	}

	randomSec, err := ParseArgToNumber[int64](setting.Args, "random")
	if err != nil {
		n.randomSec = 0
	} else {
		n.randomSec = randomSec
	}

	n.isLoop = ParseArgToBool(setting.Args, "isLoop")

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
	now := time.Now().UnixMilli()
	if now >= endTime {
		return b3.SUCCESS
	}
	return b3.RUNNING
}
