package actions

import (
	"fmt"

	b3 "github.com/dyc92/behavior3go"
	. "github.com/dyc92/behavior3go/config"
	. "github.com/dyc92/behavior3go/core"
)

const (
	EventTargetKey = "__behavior3go_event_target"
	eventKeyPrefix = "__behavior3go_event:"
)

type EventBus interface {
	TriggerEvent(event string)
	ConsumeEvent(event string) bool
}

type TriggerEvent struct {
	Action
	event string
}

func (n *TriggerEvent) Initialize(setting *BTNodeCfg) {
	n.Action.Initialize(setting)
	n.event = setting.GetPropertyAsStringSafe("event")
}

func (n *TriggerEvent) OnTick(tick *Tick) b3.Status {
	if n.event == "" || tick == nil || tick.Blackboard == nil {
		return b3.FAILURE
	}
	if bus, ok := findEventBus(tick); ok {
		bus.TriggerEvent(n.event)
		return b3.SUCCESS
	}
	triggerBlackboardEvent(tick.Blackboard, n.event)
	return b3.SUCCESS
}

type WaitForEvent struct {
	Action
	event string
}

func (n *WaitForEvent) Initialize(setting *BTNodeCfg) {
	n.Action.Initialize(setting)
	n.event = setting.GetPropertyAsStringSafe("event")
}

func (n *WaitForEvent) OnTick(tick *Tick) b3.Status {
	if n.event == "" || tick == nil || tick.Blackboard == nil {
		return b3.FAILURE
	}
	if bus, ok := findEventBus(tick); ok {
		if bus.ConsumeEvent(n.event) {
			return b3.SUCCESS
		}
		return b3.RUNNING
	}
	if consumeBlackboardEvent(tick.Blackboard, n.event) {
		return b3.SUCCESS
	}
	return b3.RUNNING
}

func findEventBus(tick *Tick) (EventBus, bool) {
	if bus, ok := tick.GetTarget().(EventBus); ok {
		return bus, true
	}
	if bus, ok := tick.Blackboard.GetMem(EventTargetKey).(EventBus); ok {
		return bus, true
	}
	return nil, false
}

func triggerBlackboardEvent(blackboard *Blackboard, event string) {
	key := eventKeyPrefix + event
	count := readEventCount(blackboard.GetMem(key))
	blackboard.SetMem(key, count+1)
}

func consumeBlackboardEvent(blackboard *Blackboard, event string) bool {
	key := eventKeyPrefix + event
	count := readEventCount(blackboard.GetMem(key))
	if count <= 0 {
		return false
	}
	count--
	if count <= 0 {
		blackboard.Remove(key)
		return true
	}
	blackboard.SetMem(key, count)
	return true
}

func readEventCount(value interface{}) int {
	switch v := value.(type) {
	case int:
		return v
	case int32:
		return int(v)
	case int64:
		return int(v)
	case float64:
		return int(v)
	default:
		var count int
		_, _ = fmt.Sscan(fmt.Sprintf("%v", value), &count)
		return count
	}
}
