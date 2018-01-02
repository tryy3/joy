package xhr

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
)

var _ event.EventTarget = (*XMLHTTPRequestUpload)(nil)

type XMLHTTPRequestUpload struct {
}

func (*XMLHTTPRequestUpload) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*XMLHTTPRequestUpload) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*XMLHTTPRequestUpload) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*XMLHTTPRequestUpload) Onabort() (onabort func(event.Event)) {
	macro.Rewrite("$_.onabort")
	return onabort
}

func (*XMLHTTPRequestUpload) SetOnabort(onabort func(event.Event)) {
	macro.Rewrite("$_.onabort = $1", onabort)
}

func (*XMLHTTPRequestUpload) Onerror() (onerror func(event.Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

func (*XMLHTTPRequestUpload) SetOnerror(onerror func(event.Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

func (*XMLHTTPRequestUpload) Onload() (onload func(event.Event)) {
	macro.Rewrite("$_.onload")
	return onload
}

func (*XMLHTTPRequestUpload) SetOnload(onload func(event.Event)) {
	macro.Rewrite("$_.onload = $1", onload)
}

func (*XMLHTTPRequestUpload) Onloadend() (onloadend func(event.Event)) {
	macro.Rewrite("$_.onloadend")
	return onloadend
}

func (*XMLHTTPRequestUpload) SetOnloadend(onloadend func(event.Event)) {
	macro.Rewrite("$_.onloadend = $1", onloadend)
}

func (*XMLHTTPRequestUpload) Onloadstart() (onloadstart func(event.Event)) {
	macro.Rewrite("$_.onloadstart")
	return onloadstart
}

func (*XMLHTTPRequestUpload) SetOnloadstart(onloadstart func(event.Event)) {
	macro.Rewrite("$_.onloadstart = $1", onloadstart)
}

func (*XMLHTTPRequestUpload) Onprogress() (onprogress func(event.Event)) {
	macro.Rewrite("$_.onprogress")
	return onprogress
}

func (*XMLHTTPRequestUpload) SetOnprogress(onprogress func(event.Event)) {
	macro.Rewrite("$_.onprogress = $1", onprogress)
}

func (*XMLHTTPRequestUpload) Ontimeout() (ontimeout func(event.Event)) {
	macro.Rewrite("$_.ontimeout")
	return ontimeout
}

func (*XMLHTTPRequestUpload) SetOntimeout(ontimeout func(event.Event)) {
	macro.Rewrite("$_.ontimeout = $1", ontimeout)
}
