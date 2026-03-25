/*
从原生工程文件加载
*/
package main

import (
	"fmt"
	b3 "github.com/dyc92/behavior3go"
	"github.com/dyc92/behavior3go/actions"
	"github.com/dyc92/behavior3go/config"
	. "github.com/dyc92/behavior3go/core"
	"github.com/dyc92/behavior3go/loader"
	"sync"
	"time"
)

// 所有的树管理
var mapTreesByID = sync.Map{}
var maps = b3.NewRegisterStructMaps()

func init() {
	//自定义节点注册
	maps.Register("Log", new(actions.Log))
	maps.Register("SetValue", new(actions.SetValue))
	maps.Register("IsValue", new(actions.IsValue))

	//获取子树的方法
	SetSubTreeLoadFunc(func(id string) *BehaviorTree {
		//println("==>load subtree:",id)
		t, ok := mapTreesByID.Load(id)
		if ok {
			return t.(*BehaviorTree)
		}
		return nil
	})
}

func main() {
	projectConfig, ok := config.LoadTreeCfg("memsubtree.b3")
	if !ok {
		fmt.Println("LoadRawProjectCfg err")
		return
	}

	var firstTree *BehaviorTree
	//载入
	for _, v := range projectConfig.Import {
		subTree, _ := config.LoadTreeCfg(v)
		tree := loader.CreateBevTreeFromConfig(subTree, maps)
		tree.Print()
		//保存到树管理
		println("==>store subtree:", tree.GetID())
		mapTreesByID.Store(tree.GetID(), tree)
		if firstTree == nil {
			firstTree = tree
		}
	}

	//循环每一帧
	for i := 0; i < 100; i++ {
		firstTree.Tick(i)
		time.Sleep(time.Millisecond * 100)
	}
}
