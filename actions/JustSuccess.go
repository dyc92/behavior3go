package actions

import (
	b3 "github.com/magicsea/behavior3go"
	. "github.com/magicsea/behavior3go/core"
)

type JustSuccess struct {
	Action
}

func (this *JustSuccess) OnTick(tick *Tick) b3.Status {
	return b3.SUCCESS
}
