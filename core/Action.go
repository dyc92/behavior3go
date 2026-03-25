package core

import (
	b3 "github.com/dyc92/behavior3go"
	. "github.com/dyc92/behavior3go/config"
)

type IAction interface {
	IBaseNode
}

/**
 * Action is the base class for all action nodes. Thus, if you want to create
 * new custom action nodes, you need to inherit from this class. For example,
 * take a look at the Runner action:
 *
 *     var Runner = b3.Class(b3.Action, {
 *       name: 'Runner',
 *
 *       tick: function(tick) {
 *         return b3.RUNNING;
 *       }
 *     });
 *
 * @module b3
 * @class Action
 * @extends BaseNode
**/
type Action struct {
	BaseNode
	BaseWorker
}

func (this *Action) Ctor() {
	this.category = b3.ACTION
}
func (this *Action) Initialize(params *BTNodeCfg) {
	params.Category = b3.ACTION
	//this.id = b3.CreateUUID()
	this.BaseNode.Initialize(params)
	//this.BaseNode.IBaseWorker = this
	this.args = make(map[string]interface{})
}
