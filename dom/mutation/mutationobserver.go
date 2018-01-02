package mutation

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/dom"
)

func New(callback func(mutations []*MutationRecord, observer *MutationObserver)) *MutationObserver {
	macro.Rewrite("new MutationObserver($1)", callback)
	return &MutationObserver{}
}

type MutationObserver struct {
}

func (*MutationObserver) Disconnect() {
	macro.Rewrite("$_.disconnect()")
}

func (*MutationObserver) Observe(target dom.Node, options *MutationObserverInit) {
	macro.Rewrite("$_.observe($1, $2)", target, options)
}

func (*MutationObserver) TakeRecords() (m []*MutationRecord) {
	macro.Rewrite("$_.takeRecords()")
	return m
}
