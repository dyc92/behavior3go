package decorators

import (
	b3 "github.com/dyc92/behavior3go"
	. "github.com/dyc92/behavior3go/core"
)

/**
 * Repeater is a decorator that repeats the tick signal until the child node
 * return `RUNNING` or `ERROR`. Optionally, a maximum number of repetitions
 * can be defined.
 *
 * @module b3
 * @class Repeater
 * @extends Decorator
**/
type Once struct {
	Decorator
}

/**
 * Open method.
 * @method open
 * @param {Tick} tick A tick instance.
**/
func (this *Once) OnOpen(tick *Tick) {
	tick.Blackboard.Set("once", false, tick.GetTree().GetID(), this.GetID())
}

/**
 * Tick method.
 * @method tick
 * @param {b3.Tick} tick A tick instance.
 * @return {Constant} A state constant.
**/
func (this *Once) OnTick(tick *Tick) b3.Status {
	//fmt.Println("tick ", this.GetTitle())

	if this.GetChild(0) == nil {
		return b3.ERROR
	}
	var i = tick.Blackboard.GetBool("once", tick.GetTree().GetID(), this.GetID())
	if i {
		return b3.FAILURE
	}

	var status = b3.SUCCESS

	this.GetChild(0).Execute(tick)

	tick.Blackboard.Set("once", true, tick.GetTree().GetID(), this.GetID())
	return status
}
