package core

import (
	b3 "github.com/dyc92/behavior3go"
	. "github.com/dyc92/behavior3go/config"
)

type ICondition interface {
	IBaseNode
}

type Condition struct {
	BaseNode
	BaseWorker
}

func (this *Condition) Ctor() {

	this.category = b3.CONDITION
}

/**
 * Initialization method.
 *
 * @method Initialize
 * @construCtor
**/
func (this *Condition) Initialize(params *BTNodeCfg) {
	params.Category = b3.CONDITION
	this.BaseNode.Initialize(params)
	//this.BaseNode.IBaseWorker = this
}
