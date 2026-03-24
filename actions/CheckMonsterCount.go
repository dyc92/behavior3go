package actions

import (
	b3 "github.com/magicsea/behavior3go"
	. "github.com/magicsea/behavior3go/config"
	. "github.com/magicsea/behavior3go/core"
)

type CheckMonsterCount struct {
	Action
	monsterID int64
	count     int64
	compChar  string
}

func (this *CheckMonsterCount) Initialize(setting *BTNodeCfg) {
	this.Action.Initialize(setting)
	this.monsterID = setting.GetPropertyAsInt64("monster_config_id")
	this.count = setting.GetPropertyAsInt64("count")
	this.compChar = setting.GetPropertyAsString("char")
}

func (this *CheckMonsterCount) OnTick(tick *Tick) b3.Status {
	//TODO: 检查怪物数量是否符合指定范围
	return b3.SUCCESS
}
