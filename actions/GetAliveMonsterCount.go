package actions

import (
	"fmt"
	b3 "github.com/magicsea/behavior3go"
	. "github.com/magicsea/behavior3go/config"
	. "github.com/magicsea/behavior3go/core"
)

type GetAliveMonsterCount struct {
	Action
	monsterID int64
}

func (this *GetAliveMonsterCount) Initialize(setting *BTNodeCfg) {
	this.Action.Initialize(setting)
	this.monsterID = setting.GetPropertyAsInt64("monster_id")
}

func (this *GetAliveMonsterCount) OnTick(tick *Tick) b3.Status {
	fmt.Println("log:", this.monsterID)
	return b3.SUCCESS
}
