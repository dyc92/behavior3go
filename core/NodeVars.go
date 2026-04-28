package core

import "github.com/dyc92/behavior3go/config"

type NodeVars struct {
	Input  []string
	Output []string
}

func NewNodeVars(setting *config.BTNodeCfg) NodeVars {
	var vars NodeVars
	vars.Initialize(setting)
	return vars
}

func (n *NodeVars) Initialize(setting *config.BTNodeCfg) {
	if setting == nil {
		return
	}
	n.Input = append([]string(nil), setting.Input...)
	n.Output = append([]string(nil), setting.Output...)
}

func (n *NodeVars) FirstInput() string {
	for _, key := range n.Input {
		if key != "" {
			return key
		}
	}
	return ""
}

func (n *NodeVars) FirstOutput() string {
	for _, key := range n.Output {
		if key != "" {
			return key
		}
	}
	return ""
}

func (n *NodeVars) ReadInput(tick *Tick) (interface{}, bool) {
	key := n.FirstInput()
	if key == "" || tick == nil || tick.Blackboard == nil {
		return nil, false
	}
	value := tick.Blackboard.GetMem(key)
	return value, value != nil
}

func (n *NodeVars) WriteOutput(tick *Tick, value interface{}) {
	key := n.FirstOutput()
	if key == "" || tick == nil || tick.Blackboard == nil {
		return
	}
	tick.Blackboard.SetMem(key, value)
}
