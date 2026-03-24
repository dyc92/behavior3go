package actions

import (
	b3 "github.com/magicsea/behavior3go"
	. "github.com/magicsea/behavior3go/config"
	. "github.com/magicsea/behavior3go/core"
)

type CreateMonster struct {
	Action
	monsterConfigId int32
	count           int32
}

func (this *CreateMonster) Initialize(setting *BTNodeCfg) {
	this.Action.Initialize(setting)
	this.monsterConfigId = setting.GetPropertyAsInt32("monster_config_id")
	this.count = setting.GetPropertyAsInt32("count")
}

func (this *CreateMonster) OnTick(tick *Tick) b3.Status {

	//TODO: create monster
	return b3.SUCCESS
}
