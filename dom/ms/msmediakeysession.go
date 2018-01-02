package ms

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
)

var _ event.EventTarget = (*MSMediaKeySession)(nil)

type MSMediaKeySession struct {
}

func (*MSMediaKeySession) Close() {
	macro.Rewrite("$_.close()")
}

func (*MSMediaKeySession) Update(key []uint8) {
	macro.Rewrite("$_.update($1)", key)
}

func (*MSMediaKeySession) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*MSMediaKeySession) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*MSMediaKeySession) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*MSMediaKeySession) Error() (err *MSMediaKeyError) {
	macro.Rewrite("$_.error")
	return err
}

func (*MSMediaKeySession) KeySystem() (keySystem string) {
	macro.Rewrite("$_.keySystem")
	return keySystem
}

func (*MSMediaKeySession) SessionID() (sessionId string) {
	macro.Rewrite("$_.sessionId")
	return sessionId
}
