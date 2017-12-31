package fetch


import (
	"github.com/matthewmueller/joy/dom/file"
	"github.com/matthewmueller/joy/dom/streams"
	"github.com/matthewmueller/joy/macro"
)

// ResponseType type
type ResponseType string

// ResponseInit type
type ResponseInit struct {
	headers    *interface{}
	status     *uint8
	statusText *string
}

// NewResponse fn
func NewResponse(body *interface{}, init *ResponseInit) *Response {
	macro.Rewrite("new Response($1, $2)", body, init)
	return &Response{}
}

// Response struct
// js:"Response,omit"
type Response struct {
}

// Clone fn
// js:"clone"
func (*Response) Clone() (r *Response) {
	macro.Rewrite("$_.clone()")
	return r
}

// ArrayBuffer fn
// js:"arrayBuffer"
func (*Response) ArrayBuffer() (b []byte) {
	macro.Rewrite("await $_.arrayBuffer()")
	return b
}

// Blob fn
// js:"blob"
func (*Response) Blob() (b file.Blob) {
	macro.Rewrite("await $_.blob()")
	return b
}

// JSON fn
// js:"json"
func (*Response) JSON() (i interface{}) {
	macro.Rewrite("await $_.json()")
	return i
}

// Text fn
// js:"text"
func (*Response) Text() (s string) {
	macro.Rewrite("await $_.text()")
	return s
}

// Body prop
// js:"body"
func (*Response) Body() (body *streams.ReadableStream) {
	macro.Rewrite("$_.body")
	return body
}

// Headers prop
// js:"headers"
func (*Response) Headers() (headers *Headers) {
	macro.Rewrite("$_.headers")
	return headers
}

// Ok prop
// js:"ok"
func (*Response) Ok() (ok bool) {
	macro.Rewrite("$_.ok")
	return ok
}

// Status prop
// js:"status"
func (*Response) Status() (status uint8) {
	macro.Rewrite("$_.status")
	return status
}

// StatusText prop
// js:"statusText"
func (*Response) StatusText() (statusText string) {
	macro.Rewrite("$_.statusText")
	return statusText
}

// Type prop
// js:"type"
func (*Response) Type() (kind *ResponseType) {
	macro.Rewrite("$_.type")
	return kind
}

// URL prop
// js:"url"
func (*Response) URL() (url string) {
	macro.Rewrite("$_.url")
	return url
}

// BodyUsed prop
// js:"bodyUsed"
func (*Response) BodyUsed() (bodyUsed bool) {
	macro.Rewrite("$_.bodyUsed")
	return bodyUsed
}
