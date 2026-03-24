package actions

import (
	b3 "github.com/magicsea/behavior3go"
	. "github.com/magicsea/behavior3go/config"
	. "github.com/magicsea/behavior3go/core"
)

type BattleSettlement struct {
	Action
}

func (this *BattleSettlement) Initialize(setting *BTNodeCfg) {
	this.Action.Initialize(setting)
}

func (this *BattleSettlement) OnTick(tick *Tick) b3.Status {

	return b3.SUCCESS
}
