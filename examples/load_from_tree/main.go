/*
从导出的树文件加载
*/
package main

import (
	"fmt"
	b3 "github.com/dyc92/behavior3go"
	"github.com/dyc92/behavior3go/actions"
	"github.com/dyc92/behavior3go/config"
	"github.com/dyc92/behavior3go/loader"
)

func main() {
	treeConfig, ok := config.LoadTreeCfg("tree.json")
	if !ok {
		fmt.Println("LoadTreeCfg err")
		return
	}
	//自定义节点注册
	maps := b3.NewRegisterStructMaps()
	maps.Register("Log", new(actions.Log))

	//载入
	tree := loader.CreateBevTreeFromConfig(treeConfig, maps)
	tree.Print()

	//循环每一帧
	for i := 0; i < 5; i++ {
		tree.Tick(i)
	}

}
