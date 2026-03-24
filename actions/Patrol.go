package actions

import (
	b3 "github.com/magicsea/behavior3go"
	. "github.com/magicsea/behavior3go/config"
	. "github.com/magicsea/behavior3go/core"
)

type Patrol struct {
	Action
	pointA []int64
	pointB []int64
}

func (this *Patrol) Initialize(setting *BTNodeCfg) {
	this.Action.Initialize(setting)
	this.pointA = setting.GetPropertyAsSliceInt("point_a")
	this.pointB = setting.GetPropertyAsSliceInt("point_b")
}

func (this *Patrol) OnTick(tick *Tick) b3.Status {

	//TODO: 实现巡逻逻辑
	return b3.SUCCESS
}
