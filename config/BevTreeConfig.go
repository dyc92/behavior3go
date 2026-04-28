package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	Category string                 `json:"category,omitempty"`
}

func (node *BTNodeCfg) GetPropertyAsString(str string) string {
	return node.Args[str].(string)
}

func (node *BTNodeCfg) GetPropertyAsStringSafe(str string) string {
	if node == nil || node.Args == nil {
		return ""
	}
	value, ok := node.Args[str]
	if !ok || value == nil {
		return ""
	}
	return strings.TrimSpace(fmt.Sprintf("%v", value))
}

func (node *BTNodeCfg) GetPropertyAsBool(str string) bool {
	if node == nil || node.Args == nil {
		return false
	}
	value, ok := node.Args[str]
	if !ok || value == nil {
		return false
	}
	switch v := value.(type) {
	case bool:
		return v
	case string:
		return strings.EqualFold(strings.TrimSpace(v), "true") || strings.TrimSpace(v) == "1"
	default:
		return fmt.Sprintf("%v", value) == "1"
	}
}

func (node *BTNodeCfg) GetPropertyAsInt32Slice(str string) []int32 {
	if node == nil || node.Args == nil {
		return nil
	}
	value, ok := node.Args[str]
	if !ok || value == nil {
		return nil
	}
	items, ok := value.([]interface{})
	if !ok {
		return nil
	}
	ret := make([]int32, 0, len(items))
	for _, item := range items {
		i, err := strconv.ParseInt(fmt.Sprintf("%v", item), 10, 32)
		if err == nil && i != 0 {
			ret = append(ret, int32(i))
		}
	}
	return ret
}

func (node *BTNodeCfg) FirstInput() string {
	if node == nil {
		return ""
	}
	for _, key := range node.Input {
		if key != "" {
			return key
		}
	}
	return ""
}

func (node *BTNodeCfg) FirstOutput() string {
	if node == nil {
		return ""
	}
	for _, key := range node.Output {
		if key != "" {
			return key
		}
	}
	return ""
}

func (node *BTNodeCfg) GetPropertyAsInt32(str string) int32 {
	v, ok := node.Args[str]
	if !ok {
		fmt.Println("fail, not found:", str)
		return 0
	}
	if v == "" {
		fmt.Println("fail, empty:", str)
		return 0
	}

	i, err := strconv.ParseInt(fmt.Sprintf("%v", v), 10, 32)
	if err != nil {
		fmt.Println("fail, parse int42:", err, v)
		return 0
	}
	return int32(i)
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
func (node *BTNodeCfg) GetPropertyAsFloat64(str string) float64 {
	v, ok := node.Args[str]
	if !ok {
		fmt.Println("fail, not found:", str)
		return 0
	}
	if v == "" {
		fmt.Println("fail, empty:", str)
		return 0
	}
	i, err := strconv.ParseFloat(fmt.Sprintf("%v", v), 10)
	if err != nil {
		fmt.Println("fail, parse int64:", err, v)
		return 0
	}
	return i
}

func (node *BTNodeCfg) GetPropertyAsSlice(str string) []interface{} {
	return node.Args[str].([]interface{})
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
