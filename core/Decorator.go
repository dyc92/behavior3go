package core

import (
	b3 "github.com/dyc92/behavior3go"
	. "github.com/dyc92/behavior3go/config"
)

type IDecorator interface {
	IBaseNode
}

type Decorator struct {
	BaseNode
	BaseWorker
}

func (this *Decorator) Ctor() {

	this.category = b3.DECORATOR
}

/**
 * Initialization method.
 *
 * @method Initialize
 * @construCtor
**/
func (this *Decorator) Initialize(params *BTNodeCfg) {
	params.Category = b3.DECORATOR
	this.BaseNode.Initialize(params)
	//this.BaseNode.IBaseWorker = this
}
