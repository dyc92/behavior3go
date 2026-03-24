package composites

import (
	_ "fmt"

	b3 "github.com/magicsea/behavior3go"
	_ "github.com/magicsea/behavior3go/config"
	. "github.com/magicsea/behavior3go/core"
)

type RetryUntilSuccess struct {
	Composite
}

/**
 * Tick method.
 * @method tick
 * @param {b3.Tick} tick A tick instance.
 * @return {Constant} A state constant.
**/
func (this *RetryUntilSuccess) OnTick(tick *Tick) b3.Status {
	count, ok := this.BaseNode.GetArgs("count").(int)
	if !ok {
		for {
			var status = this.GetChild(0).Execute(tick)
			if status == b3.SUCCESS {
				return b3.SUCCESS
			}
		}
	}
	for i := 0; i < count; i++ {
		var status = this.GetChild(0).Execute(tick)
		if status == b3.SUCCESS {
			return b3.SUCCESS
		}
	}
	return b3.SUCCESS
}
