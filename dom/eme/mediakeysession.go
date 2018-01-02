package eme

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
)

var _ event.EventTarget = (*MediaKeySession)(nil)

type MediaKeySession struct {
}

func (*MediaKeySession) Close() {
	macro.Rewrite("await $_.close()")
}

func (*MediaKeySession) GenerateRequest(initDataType string, initData []byte) {
	macro.Rewrite("await $_.generateRequest($1, $2)", initDataType, initData)
}

func (*MediaKeySession) Load(sessionId string) (b bool) {
	macro.Rewrite("await $_.load($1)", sessionId)
	return b
}

func (*MediaKeySession) Remove() {
	macro.Rewrite("await $_.remove()")
}

func (*MediaKeySession) Update(response []byte) {
	macro.Rewrite("await $_.update($1)", response)
}

func (*MediaKeySession) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*MediaKeySession) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*MediaKeySession) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*MediaKeySession) Closed() {
	macro.Rewrite("await $_.closed")
}

func (*MediaKeySession) Expiration() (expiration float32) {
	macro.Rewrite("$_.expiration")
	return expiration
}

func (*MediaKeySession) KeyStatuses() (keyStatuses *MediaKeyStatusMap) {
	macro.Rewrite("$_.keyStatuses")
	return keyStatuses
}

func (*MediaKeySession) SessionID() (sessionId string) {
	macro.Rewrite("$_.sessionId")
	return sessionId
}
