package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// 编辑器地址@http://editor.behavior3.com/#/editor
// 节点json类型
type BTNodeCfg struct {
	Id       string                 `json:"id"`                 // 对应原 id: number
	Name     string                 `json:"name"`               // 对应原 name: string
	Desc     *string                `json:"desc,omitempty"`     // 可选字符串，用指针标识 omitempty
	Args     map[string]interface{} `json:"args,omitempty"`     // 对应原 args?: { [key: string]: any }
	Input    []string               `json:"input,omitempty"`    // 可选字符串数组
	Output   []string               `json:"output,omitempty"`   // 可选字符串数组
	Children []*BTNodeCfg           `json:"children,omitempty"` // 嵌套自身结构的切片，指针避免循环拷贝
	Debug    *bool                  `json:"debug,omitempty"`    // 可选布尔值，指针标识 omitempty
	Disabled *bool                  `json:"disabled,omitempty"` // 可选布尔值，指针标识 omitempty
	Path     *string                `json:"path,omitempty"`     // 可选字符串
}

func (node *BTNodeCfg) GetPropertyAsString(str string) string {
	return node.Args[str].(string)
}

func (node *BTNodeCfg) GetPropertyAsInt64(str string) int64 {
	v, ok := node.Args[str]
	if !ok {
		fmt.Println("fail, not found:", str)
		return 0
	}
	if v == "" {
		fmt.Println("fail, empty:", str)
		return 0
	}
	i, err := strconv.ParseInt(fmt.Sprintf("%v", v), 10, 64)
	if err != nil {
		fmt.Println("fail, parse int64:", err, v)
		return 0
	}
	return i
}

func (node *BTNodeCfg) GetPropertyAsSliceInt(str string) []int64 {
	return node.Args[str].([]int64)
}

type TreeVar struct {
	Name string `json:"name"` // 变量名
	Desc string `json:"desc"` // 变量描述
}

// BTTreeCfg 树json类型
type BTTreeCfg struct {
	Version string    `json:"version"`          // 对应 version: string
	Name    string    `json:"name"`             // 对应 name: string
	Desc    *string   `json:"desc,omitempty"`   // 可选字符串，指针+omitempty实现可选
	Export  *bool     `json:"export,omitempty"` // 可选布尔值，同上
	FirstID int       `json:"firstid"`          // 对应 firstid: number（Go 字段名遵循驼峰规范，json标签保持原字段名）
	Group   []string  `json:"group"`            // 对应 group: string[]
	Import  []string  `json:"import"`           // 对应 import: string[]
	Vars    []TreeVar `json:"vars"`             // 对应 vars: { name: string; desc: string }[]
	Root    BTNodeCfg `json:"root"`             // 对应 root: NodeModel（非可选，直接值类型）
}

// 加载
func LoadTreeCfg(path string) (*BTTreeCfg, bool) {

	var tree BTTreeCfg
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("fail:", err)
		return nil, false
	}
	err = json.Unmarshal(file, &tree)
	if err != nil {
		fmt.Println("fail, ummarshal:", err, len(file))
		return nil, false
	}

	//fmt.Println("load tree:", tree.Title, " nodes:", len(tree.Nodes))
	return &tree, true
}
