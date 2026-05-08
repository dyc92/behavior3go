package composites

import (
	_ "fmt"

	b3 "github.com/dyc92/behavior3go"
	_ "github.com/dyc92/behavior3go/config"
	. "github.com/dyc92/behavior3go/core"
)

type Switch struct {
	Composite
}

func (this *Switch) OnTick(tick *Tick) b3.Status {

	for i := 0; i < this.GetChildCount(); i++ {
		caseNode := this.GetChild(i)
		if caseNode.GetChild(0).Execute(tick) == b3.SUCCESS {
			return caseNode.Execute(tick)
		}
	}
	return b3.FAILURE
}
