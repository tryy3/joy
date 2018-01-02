package utils

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
)

var _ event.EventTarget = (*ApplicationCache)(nil)

type ApplicationCache struct {
}

func (*ApplicationCache) Abort() {
	macro.Rewrite("$_.abort()")
}

func (*ApplicationCache) SwapCache() {
	macro.Rewrite("$_.swapCache()")
}

func (*ApplicationCache) Update() {
	macro.Rewrite("$_.update()")
}

func (*ApplicationCache) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*ApplicationCache) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*ApplicationCache) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*ApplicationCache) Oncached() (oncached func(event.Event)) {
	macro.Rewrite("$_.oncached")
	return oncached
}

func (*ApplicationCache) SetOncached(oncached func(event.Event)) {
	macro.Rewrite("$_.oncached = $1", oncached)
}

func (*ApplicationCache) Onchecking() (onchecking func(event.Event)) {
	macro.Rewrite("$_.onchecking")
	return onchecking
}

func (*ApplicationCache) SetOnchecking(onchecking func(event.Event)) {
	macro.Rewrite("$_.onchecking = $1", onchecking)
}

func (*ApplicationCache) Ondownloading() (ondownloading func(event.Event)) {
	macro.Rewrite("$_.ondownloading")
	return ondownloading
}

func (*ApplicationCache) SetOndownloading(ondownloading func(event.Event)) {
	macro.Rewrite("$_.ondownloading = $1", ondownloading)
}

func (*ApplicationCache) Onerror() (onerror func(event.Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

func (*ApplicationCache) SetOnerror(onerror func(event.Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

func (*ApplicationCache) Onnoupdate() (onnoupdate func(event.Event)) {
	macro.Rewrite("$_.onnoupdate")
	return onnoupdate
}

func (*ApplicationCache) SetOnnoupdate(onnoupdate func(event.Event)) {
	macro.Rewrite("$_.onnoupdate = $1", onnoupdate)
}

func (*ApplicationCache) Onobsolete() (onobsolete func(event.Event)) {
	macro.Rewrite("$_.onobsolete")
	return onobsolete
}

func (*ApplicationCache) SetOnobsolete(onobsolete func(event.Event)) {
	macro.Rewrite("$_.onobsolete = $1", onobsolete)
}

func (*ApplicationCache) Onprogress() (onprogress func(*ProgressEvent)) {
	macro.Rewrite("$_.onprogress")
	return onprogress
}

func (*ApplicationCache) SetOnprogress(onprogress func(*ProgressEvent)) {
	macro.Rewrite("$_.onprogress = $1", onprogress)
}

func (*ApplicationCache) Onupdateready() (onupdateready func(event.Event)) {
	macro.Rewrite("$_.onupdateready")
	return onupdateready
}

func (*ApplicationCache) SetOnupdateready(onupdateready func(event.Event)) {
	macro.Rewrite("$_.onupdateready = $1", onupdateready)
}

func (*ApplicationCache) Status() (status uint8) {
	macro.Rewrite("$_.status")
	return status
}
