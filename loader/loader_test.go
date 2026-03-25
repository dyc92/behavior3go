package loader

import (
	"fmt"
	"github.com/dyc92/behavior3go/actions"
	"github.com/dyc92/behavior3go/condition"
	"github.com/dyc92/behavior3go/config"
	"github.com/dyc92/behavior3go/decorators"
	"reflect"
	"testing"

	b3 "github.com/dyc92/behavior3go"

	. "github.com/dyc92/behavior3go/core"
	//. "github.com/dyc92/behavior3go/decorators"
)

type Test struct {
	value string
}

func (test *Test) Print() {
	fmt.Println(test.value)
}

func TestExample(t *testing.T) {
	maps := createBaseStructMaps()
	if data, err := maps.New("Runner"); err != nil {
		t.Error("Error:", err, data)
	} else {
		t.Log(reflect.TypeOf(data))
	}

}

// /////////////////////加载事例///////////////////////////
// 自定义action节点
type LogTest struct {
	Action
	info string
}

func (this *LogTest) OnTick(tick *Tick) b3.Status {
	fmt.Println("logtest:", this.info)
	return b3.SUCCESS
}

func TestLoadTree(t *testing.T) {
	treeConfig, ok := config.LoadTreeCfg("base_dungeon_tree.json")
	if ok {
		//自定义节点注册
		maps := b3.NewRegisterStructMaps()
		maps.Register("Log", new(LogTest))
		maps.Register("CheckMonsterCount", new(condition.CheckMonsterCount))
		maps.Register("Attack", new(actions.Attack))
		maps.Register("BattleSettlement", new(actions.BattleSettlement))
		maps.Register("MoveToTarget", new(actions.MoveToTarget))
		maps.Register("Patrol", new(actions.Patrol))
		maps.Register("Once", new(decorators.Once))
		maps.Register("CreateMonster", new(actions.CreateMonster))

		//载入
		tree := CreateBevTreeFromConfig(treeConfig, maps)
		tree.Print()

		//循环每一帧
		for i := 0; i < 1; i++ {
			tree.Tick(i)
		}
	} else {
		t.Error("LoadTreeCfg err")
	}

}
