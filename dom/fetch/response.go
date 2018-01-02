package fetch

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/file"
	"github.com/matthewmueller/joy/dom/streams"
)

type ResponseType string

type ResponseInit struct {
	headers		*interface{}
	status		*uint8
	statusText	*string
}

func NewResponse(body *interface{}, init *ResponseInit) *Response {
	macro.Rewrite("new Response($1, $2)", body, init)
	return &Response{}
}

type Response struct {
}

func (*Response) Clone() (r *Response) {
	macro.Rewrite("$_.clone()")
	return r
}

func (*Response) ArrayBuffer() (b []byte) {
	macro.Rewrite("await $_.arrayBuffer()")
	return b
}

func (*Response) Blob() (b file.Blob) {
	macro.Rewrite("await $_.blob()")
	return b
}

func (*Response) JSON() (i interface{}) {
	macro.Rewrite("await $_.json()")
	return i
}

func (*Response) Text() (s string) {
	macro.Rewrite("await $_.text()")
	return s
}

func (*Response) Body() (body *streams.ReadableStream) {
	macro.Rewrite("$_.body")
	return body
}

func (*Response) Headers() (headers *Headers) {
	macro.Rewrite("$_.headers")
	return headers
}

func (*Response) Ok() (ok bool) {
	macro.Rewrite("$_.ok")
	return ok
}

func (*Response) Status() (status uint8) {
	macro.Rewrite("$_.status")
	return status
}

func (*Response) StatusText() (statusText string) {
	macro.Rewrite("$_.statusText")
	return statusText
}

func (*Response) Type() (kind *ResponseType) {
	macro.Rewrite("$_.type")
	return kind
}

func (*Response) URL() (url string) {
	macro.Rewrite("$_.url")
	return url
}

func (*Response) BodyUsed() (bodyUsed bool) {
	macro.Rewrite("$_.bodyUsed")
	return bodyUsed
}
