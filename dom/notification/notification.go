package notification

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
)

var _ event.EventTarget = (*Notification)(nil)

func NewNotification(title string, options *NotificationOptions) *Notification {
	macro.Rewrite("new Notification($1, $2)", title, options)
	return &Notification{}
}

type Notification struct {
}

func (*Notification) Close() {
	macro.Rewrite("$_.close()")
}

func (*Notification) RequestPermission(callback *func(permission *NotificationPermission)) (n *NotificationPermission) {
	macro.Rewrite("await $_.requestPermission($1)", callback)
	return n
}

func (*Notification) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*Notification) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*Notification) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*Notification) Body() (body string) {
	macro.Rewrite("$_.body")
	return body
}

func (*Notification) Dir() (dir *NotificationDirection) {
	macro.Rewrite("$_.dir")
	return dir
}

func (*Notification) Icon() (icon string) {
	macro.Rewrite("$_.icon")
	return icon
}

func (*Notification) Lang() (lang string) {
	macro.Rewrite("$_.lang")
	return lang
}

func (*Notification) Onclick() (onclick func(event.Event)) {
	macro.Rewrite("$_.onclick")
	return onclick
}

func (*Notification) SetOnclick(onclick func(event.Event)) {
	macro.Rewrite("$_.onclick = $1", onclick)
}

func (*Notification) Onclose() (onclose func(event.Event)) {
	macro.Rewrite("$_.onclose")
	return onclose
}

func (*Notification) SetOnclose(onclose func(event.Event)) {
	macro.Rewrite("$_.onclose = $1", onclose)
}

func (*Notification) Onerror() (onerror func(event.Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

func (*Notification) SetOnerror(onerror func(event.Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

func (*Notification) Onshow() (onshow func(event.Event)) {
	macro.Rewrite("$_.onshow")
	return onshow
}

func (*Notification) SetOnshow(onshow func(event.Event)) {
	macro.Rewrite("$_.onshow = $1", onshow)
}

func (*Notification) Permission() (permission *NotificationPermission) {
	macro.Rewrite("$_.permission")
	return permission
}

func (*Notification) Tag() (tag string) {
	macro.Rewrite("$_.tag")
	return tag
}

func (*Notification) Title() (title string) {
	macro.Rewrite("$_.title")
	return title
}
