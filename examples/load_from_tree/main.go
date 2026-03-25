/*
从导出的树文件加载
*/
package main

import (
	"fmt"
	b3 "github.com/dyc92/behavior3go"
)

func main() {
	treeConfig, ok := LoadTreeCfg("tree.json")
	if !ok {
		fmt.Println("LoadTreeCfg err")
		return
	}
	//自定义节点注册
	maps := b3.NewRegisterStructMaps()
	maps.Register("Log", new(LogTest))

	//载入
	tree := CreateBevTreeFromConfig(treeConfig, maps)
	tree.Print()

	//循环每一帧
	for i := 0; i < 5; i++ {
		tree.Tick(i)
	}

}
