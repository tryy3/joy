package xhr

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/document"
)

var _ event.EventTarget = (*XMLHTTPRequest)(nil)

func New() *XMLHTTPRequest {
	macro.Rewrite("new XMLHttpRequest()")
	return &XMLHTTPRequest{}
}

type XMLHTTPRequest struct {
}

func (*XMLHTTPRequest) Abort() {
	macro.Rewrite("$_.abort()")
}

func (*XMLHTTPRequest) GetAllResponseHeaders() (s string) {
	macro.Rewrite("$_.getAllResponseHeaders()")
	return s
}

func (*XMLHTTPRequest) GetResponseHeader(header string) (s string) {
	macro.Rewrite("$_.getResponseHeader($1)", header)
	return s
}

func (*XMLHTTPRequest) MsCachingEnabled() (b bool) {
	macro.Rewrite("$_.msCachingEnabled()")
	return b
}

func (*XMLHTTPRequest) Open(method string, url string, async *bool, user *string, password *string) {
	macro.Rewrite("$_.open($1, $2, $3, $4, $5)", method, url, async, user, password)
}

func (*XMLHTTPRequest) OverrideMimeType(mime string) {
	macro.Rewrite("$_.overrideMimeType($1)", mime)
}

func (*XMLHTTPRequest) Send(data *interface{}) {
	macro.Rewrite("$_.send($1)", data)
}

func (*XMLHTTPRequest) SetRequestHeader(header string, value string) {
	macro.Rewrite("$_.setRequestHeader($1, $2)", header, value)
}

func (*XMLHTTPRequest) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*XMLHTTPRequest) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*XMLHTTPRequest) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*XMLHTTPRequest) MsCaching() (msCaching string) {
	macro.Rewrite("$_.msCaching")
	return msCaching
}

func (*XMLHTTPRequest) SetMsCaching(msCaching string) {
	macro.Rewrite("$_.msCaching = $1", msCaching)
}

func (*XMLHTTPRequest) Onreadystatechange() (onreadystatechange func(event.Event)) {
	macro.Rewrite("$_.onreadystatechange")
	return onreadystatechange
}

func (*XMLHTTPRequest) SetOnreadystatechange(onreadystatechange func(event.Event)) {
	macro.Rewrite("$_.onreadystatechange = $1", onreadystatechange)
}

func (*XMLHTTPRequest) ReadyState() (readyState uint8) {
	macro.Rewrite("$_.readyState")
	return readyState
}

func (*XMLHTTPRequest) Response() (response interface{}) {
	macro.Rewrite("$_.response")
	return response
}

func (*XMLHTTPRequest) ResponseText() (responseText string) {
	macro.Rewrite("$_.responseText")
	return responseText
}

func (*XMLHTTPRequest) ResponseType() (responseType *XMLHTTPRequestResponseType) {
	macro.Rewrite("$_.responseType")
	return responseType
}

func (*XMLHTTPRequest) SetResponseType(responseType *XMLHTTPRequestResponseType) {
	macro.Rewrite("$_.responseType = $1", responseType)
}

func (*XMLHTTPRequest) ResponseURL() (responseURL string) {
	macro.Rewrite("$_.responseURL")
	return responseURL
}

func (*XMLHTTPRequest) ResponseXML() (responseXML document.Document) {
	macro.Rewrite("$_.responseXML")
	return responseXML
}

func (*XMLHTTPRequest) Status() (status uint8) {
	macro.Rewrite("$_.status")
	return status
}

func (*XMLHTTPRequest) StatusText() (statusText string) {
	macro.Rewrite("$_.statusText")
	return statusText
}

func (*XMLHTTPRequest) Timeout() (timeout uint) {
	macro.Rewrite("$_.timeout")
	return timeout
}

func (*XMLHTTPRequest) SetTimeout(timeout uint) {
	macro.Rewrite("$_.timeout = $1", timeout)
}

func (*XMLHTTPRequest) Upload() (upload *XMLHTTPRequestUpload) {
	macro.Rewrite("$_.upload")
	return upload
}

func (*XMLHTTPRequest) WithCredentials() (withCredentials bool) {
	macro.Rewrite("$_.withCredentials")
	return withCredentials
}

func (*XMLHTTPRequest) SetWithCredentials(withCredentials bool) {
	macro.Rewrite("$_.withCredentials = $1", withCredentials)
}

func (*XMLHTTPRequest) Onabort() (onabort func(event.Event)) {
	macro.Rewrite("$_.onabort")
	return onabort
}

func (*XMLHTTPRequest) SetOnabort(onabort func(event.Event)) {
	macro.Rewrite("$_.onabort = $1", onabort)
}

func (*XMLHTTPRequest) Onerror() (onerror func(event.Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

func (*XMLHTTPRequest) SetOnerror(onerror func(event.Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

func (*XMLHTTPRequest) Onload() (onload func(event.Event)) {
	macro.Rewrite("$_.onload")
	return onload
}

func (*XMLHTTPRequest) SetOnload(onload func(event.Event)) {
	macro.Rewrite("$_.onload = $1", onload)
}

func (*XMLHTTPRequest) Onloadend() (onloadend func(event.Event)) {
	macro.Rewrite("$_.onloadend")
	return onloadend
}

func (*XMLHTTPRequest) SetOnloadend(onloadend func(event.Event)) {
	macro.Rewrite("$_.onloadend = $1", onloadend)
}

func (*XMLHTTPRequest) Onloadstart() (onloadstart func(event.Event)) {
	macro.Rewrite("$_.onloadstart")
	return onloadstart
}

func (*XMLHTTPRequest) SetOnloadstart(onloadstart func(event.Event)) {
	macro.Rewrite("$_.onloadstart = $1", onloadstart)
}

func (*XMLHTTPRequest) Onprogress() (onprogress func(event.Event)) {
	macro.Rewrite("$_.onprogress")
	return onprogress
}

func (*XMLHTTPRequest) SetOnprogress(onprogress func(event.Event)) {
	macro.Rewrite("$_.onprogress = $1", onprogress)
}

func (*XMLHTTPRequest) Ontimeout() (ontimeout func(event.Event)) {
	macro.Rewrite("$_.ontimeout")
	return ontimeout
}

func (*XMLHTTPRequest) SetOntimeout(ontimeout func(event.Event)) {
	macro.Rewrite("$_.ontimeout = $1", ontimeout)
}
