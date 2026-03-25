package composites

import (
	_ "fmt"

	b3 "github.com/dyc92/behavior3go"
	_ "github.com/dyc92/behavior3go/config"
	. "github.com/dyc92/behavior3go/core"
)

type IfElse struct {
	Composite
}

/**
 * Tick method.
 * @method tick
 * @param {b3.Tick} tick A tick instance.
 * @return {Constant} A state constant.
**/
func (this *IfElse) OnTick(tick *Tick) b3.Status {

	nodeCount := this.GetChildCount()
	if nodeCount != 3 {
		return b3.FAILURE
	}

	if this.GetChild(0).Execute(tick) == b3.SUCCESS {
		return this.GetChild(1).Execute(tick)
	} else {
		return this.GetChild(2).Execute(tick)
	}
}
