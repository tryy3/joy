package file

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/dom"
)

var _ event.EventTarget = (*FileReader)(nil)

func New() *FileReader {
	macro.Rewrite("new FileReader()")
	return &FileReader{}
}

type FileReader struct {
}

func (*FileReader) ReadAsArrayBuffer(blob Blob) {
	macro.Rewrite("$_.readAsArrayBuffer($1)", blob)
}

func (*FileReader) ReadAsBinaryString(blob Blob) {
	macro.Rewrite("$_.readAsBinaryString($1)", blob)
}

func (*FileReader) ReadAsDataURL(blob Blob) {
	macro.Rewrite("$_.readAsDataURL($1)", blob)
}

func (*FileReader) ReadAsText(blob Blob, encoding *string) {
	macro.Rewrite("$_.readAsText($1, $2)", blob, encoding)
}

func (*FileReader) Abort() {
	macro.Rewrite("$_.abort()")
}

func (*FileReader) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*FileReader) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*FileReader) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*FileReader) Error() (err *dom.DOMError) {
	macro.Rewrite("$_.error")
	return err
}

func (*FileReader) Onabort() (onabort func(event.Event)) {
	macro.Rewrite("$_.onabort")
	return onabort
}

func (*FileReader) SetOnabort(onabort func(event.Event)) {
	macro.Rewrite("$_.onabort = $1", onabort)
}

func (*FileReader) Onerror() (onerror func(event.Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

func (*FileReader) SetOnerror(onerror func(event.Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

func (*FileReader) Onload() (onload func(event.Event)) {
	macro.Rewrite("$_.onload")
	return onload
}

func (*FileReader) SetOnload(onload func(event.Event)) {
	macro.Rewrite("$_.onload = $1", onload)
}

func (*FileReader) Onloadend() (onloadend func(event.Event)) {
	macro.Rewrite("$_.onloadend")
	return onloadend
}

func (*FileReader) SetOnloadend(onloadend func(event.Event)) {
	macro.Rewrite("$_.onloadend = $1", onloadend)
}

func (*FileReader) Onloadstart() (onloadstart func(event.Event)) {
	macro.Rewrite("$_.onloadstart")
	return onloadstart
}

func (*FileReader) SetOnloadstart(onloadstart func(event.Event)) {
	macro.Rewrite("$_.onloadstart = $1", onloadstart)
}

func (*FileReader) Onprogress() (onprogress func(event.Event)) {
	macro.Rewrite("$_.onprogress")
	return onprogress
}

func (*FileReader) SetOnprogress(onprogress func(event.Event)) {
	macro.Rewrite("$_.onprogress = $1", onprogress)
}

func (*FileReader) ReadyState() (readyState uint8) {
	macro.Rewrite("$_.readyState")
	return readyState
}

func (*FileReader) Result() (result interface{}) {
	macro.Rewrite("$_.result")
	return result
}
