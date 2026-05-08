package composites

import (
	_ "fmt"

	b3 "github.com/dyc92/behavior3go"
	_ "github.com/dyc92/behavior3go/config"
	. "github.com/dyc92/behavior3go/core"
)

type Case struct {
	Composite
}

func (this *Case) OnTick(tick *Tick) b3.Status {
	for i := 1; i < this.GetChildCount(); i++ {

		if this.GetChild(i).Execute(tick) == b3.SUCCESS {
			return b3.SUCCESS
		}
	}
	return b3.FAILURE
}
