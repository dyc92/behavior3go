package config

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
)

type Number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

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

func ParseArgToString(g map[string]interface{}, str string) string {
	v, ok := g[str].(string)
	if !ok {
		return ""
	}
	return v
}

func ParseArgToBool(g map[string]interface{}, str string) bool {
	value, ok := g[str]
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

func convertToType[T Number | string](val interface{}) (T, error) {
	var zero T
	targetType := reflect.TypeOf(zero)
	v := reflect.ValueOf(val)

	// 仅支持数值类型转换
	if (v.Kind() < reflect.Int || v.Kind() > reflect.Float64) && v.Kind() != reflect.String {
		return zero, fmt.Errorf("值[%v]不是数值或字符串类型", val)
	}

	// 反射安全转换为目标泛型类型
	return v.Convert(targetType).Interface().(T), nil
}

func ParseArgToNumber[T Number](m map[string]interface{}, key string) (T, error) {
	val, ok := m[key]
	if !ok {
		var zero T
		return zero, fmt.Errorf("key: %s 不存在", key)
	}

	var result T
	targetType := reflect.TypeOf(result)

	result, err := convertToType[T](val)
	if err != nil {
		return result, fmt.Errorf("key: %s %w", key, err)
	}

	return result, fmt.Errorf("key: %s 的值 %v 不是数值或字符串类型，无法转换为 %v", key, val, targetType)
}

func ParseArgToSlice[T Number | string](p map[string]interface{}, key string) ([]T, error) {
	val, ok := p[key]
	if !ok {
		return nil, fmt.Errorf("key: %s 不存在", key)
	}

	slice, ok := val.([]interface{})
	if !ok {
		return nil, fmt.Errorf("key: %s 的值不是数组类型", key)
	}

	result := make([]T, 0, len(slice))
	for idx, item := range slice {
		num, err := convertToType[T](item)
		if err != nil {
			return nil, fmt.Errorf("数组索引[%d] %w", idx, err)
		}
		result = append(result, num)
	}

	return result, nil
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
