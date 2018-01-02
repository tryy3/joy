package ms

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/dom"
)

var _ event.EventTarget = (*MSStreamReader)(nil)

func New() *MSStreamReader {
	macro.Rewrite("new MSStreamReader()")
	return &MSStreamReader{}
}

type MSStreamReader struct {
}

func (*MSStreamReader) ReadAsArrayBuffer(stream *MSStream, size *int64) {
	macro.Rewrite("$_.readAsArrayBuffer($1, $2)", stream, size)
}

func (*MSStreamReader) ReadAsBinaryString(stream *MSStream, size *int64) {
	macro.Rewrite("$_.readAsBinaryString($1, $2)", stream, size)
}

func (*MSStreamReader) ReadAsBlob(stream *MSStream, size *int64) {
	macro.Rewrite("$_.readAsBlob($1, $2)", stream, size)
}

func (*MSStreamReader) ReadAsDataURL(stream *MSStream, size *int64) {
	macro.Rewrite("$_.readAsDataURL($1, $2)", stream, size)
}

func (*MSStreamReader) ReadAsText(stream *MSStream, encoding *string, size *int64) {
	macro.Rewrite("$_.readAsText($1, $2, $3)", stream, encoding, size)
}

func (*MSStreamReader) Abort() {
	macro.Rewrite("$_.abort()")
}

func (*MSStreamReader) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*MSStreamReader) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*MSStreamReader) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*MSStreamReader) Error() (err *dom.DOMError) {
	macro.Rewrite("$_.error")
	return err
}

func (*MSStreamReader) Onabort() (onabort func(event.Event)) {
	macro.Rewrite("$_.onabort")
	return onabort
}

func (*MSStreamReader) SetOnabort(onabort func(event.Event)) {
	macro.Rewrite("$_.onabort = $1", onabort)
}

func (*MSStreamReader) Onerror() (onerror func(event.Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

func (*MSStreamReader) SetOnerror(onerror func(event.Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

func (*MSStreamReader) Onload() (onload func(event.Event)) {
	macro.Rewrite("$_.onload")
	return onload
}

func (*MSStreamReader) SetOnload(onload func(event.Event)) {
	macro.Rewrite("$_.onload = $1", onload)
}

func (*MSStreamReader) Onloadend() (onloadend func(event.Event)) {
	macro.Rewrite("$_.onloadend")
	return onloadend
}

func (*MSStreamReader) SetOnloadend(onloadend func(event.Event)) {
	macro.Rewrite("$_.onloadend = $1", onloadend)
}

func (*MSStreamReader) Onloadstart() (onloadstart func(event.Event)) {
	macro.Rewrite("$_.onloadstart")
	return onloadstart
}

func (*MSStreamReader) SetOnloadstart(onloadstart func(event.Event)) {
	macro.Rewrite("$_.onloadstart = $1", onloadstart)
}

func (*MSStreamReader) Onprogress() (onprogress func(event.Event)) {
	macro.Rewrite("$_.onprogress")
	return onprogress
}

func (*MSStreamReader) SetOnprogress(onprogress func(event.Event)) {
	macro.Rewrite("$_.onprogress = $1", onprogress)
}

func (*MSStreamReader) ReadyState() (readyState uint8) {
	macro.Rewrite("$_.readyState")
	return readyState
}

func (*MSStreamReader) Result() (result interface{}) {
	macro.Rewrite("$_.result")
	return result
}
