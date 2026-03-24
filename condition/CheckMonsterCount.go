package condition

import (
	"fmt"
	b3 "github.com/magicsea/behavior3go"
	. "github.com/magicsea/behavior3go/config"
	. "github.com/magicsea/behavior3go/core"
)

type CheckMonsterCount struct {
	Condition
	monsterID int64
	count     int64
	compChar  string
}

func (this *CheckMonsterCount) Initialize(setting *BTNodeCfg) {
	this.Condition.Initialize(setting)
	this.monsterID = setting.GetPropertyAsInt64("monster_config_id")
	this.count = setting.GetPropertyAsInt64("count")
	this.compChar = setting.GetPropertyAsString("char")
}

func (this *CheckMonsterCount) OnTick(tick *Tick) b3.Status {
	//TODO: 检查怪物数量是否符合指定范围
	fmt.Println("CheckMonsterCount:", this.GetID())

	for i := 0; i < this.GetChildCount(); i++ {
		var status = this.GetChild(i).Execute(tick)
		if status != b3.SUCCESS {
			return status
		}
	}
	return b3.SUCCESS
}
