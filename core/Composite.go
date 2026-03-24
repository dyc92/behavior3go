package core

import (
	"fmt"

	b3 "github.com/magicsea/behavior3go"
	. "github.com/magicsea/behavior3go/config"
)

type IComposite interface {
	IBaseNode
	GetChildCount() int
	GetChild(index int) IBaseNode
	AddChild(child IBaseNode)
}

type Composite struct {
	BaseNode
	BaseWorker
}

func (this *Composite) Ctor() {

	this.category = b3.COMPOSITE
}

/**
 * Initialization method.
 *
 * @method Initialize
 * @construCtor
**/
func (this *Composite) Initialize(params *BTNodeCfg) {
	params.Category = b3.COMPOSITE
	this.BaseNode.Initialize(params)
	//this.BaseNode.IBaseWorker = this
	this.children = make([]IBaseNode, 0)
	//fmt.Println("Composite Initialize")
}

func (this *Composite) tick(tick *Tick) b3.Status {
	fmt.Println("tick Composite1")
	return b3.ERROR
}
