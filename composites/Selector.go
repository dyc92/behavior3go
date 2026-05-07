package composites

import (
	_ "fmt"

	b3 "github.com/dyc92/behavior3go"
	_ "github.com/dyc92/behavior3go/config"
	. "github.com/dyc92/behavior3go/core"
)

type Selector struct {
	Composite
}

/**
 * Tick method.
 * @method tick
 * @param {b3.Tick} tick A tick instance.
 * @return {Constant} A state constant.
**/
func (this *Selector) OnTick(tick *Tick) b3.Status {

	for i := 0; i < this.GetChildCount(); i++ {

		var status = this.GetChild(i).Execute(tick)
		if status == b3.SUCCESS {
			return status
		}
	}
	return b3.FAILURE
}
