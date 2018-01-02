package fetch

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/file"
)

type ReferrerPolicy string

type RequestCache string

type RequestCredentials string

type RequestDestination string

type RequestMode string

type RequestRedirect string

type RequestType string

type RequestInit struct {
	body		*interface{}
	cache		*RequestCache
	credentials	*RequestCredentials
	headers		*interface{}
	integrity	*string
	keepalive	*bool
	method		*string
	mode		*RequestMode
	redirect	*RequestRedirect
	referrer	*string
	referrerPolicy	*ReferrerPolicy
	window		*interface{}
}

func NewRequest(input interface{}, init *RequestInit) *Request {
	macro.Rewrite("new Request($1, $2)", input, init)
	return &Request{}
}

type Request struct {
}

func (*Request) Clone() (r *Request) {
	macro.Rewrite("$_.clone()")
	return r
}

func (*Request) ArrayBuffer() (b []byte) {
	macro.Rewrite("await $_.arrayBuffer()")
	return b
}

func (*Request) Blob() (b file.Blob) {
	macro.Rewrite("await $_.blob()")
	return b
}

func (*Request) JSON() (i interface{}) {
	macro.Rewrite("await $_.json()")
	return i
}

func (*Request) Text() (s string) {
	macro.Rewrite("await $_.text()")
	return s
}

func (*Request) Cache() (cache *RequestCache) {
	macro.Rewrite("$_.cache")
	return cache
}

func (*Request) Credentials() (credentials *RequestCredentials) {
	macro.Rewrite("$_.credentials")
	return credentials
}

func (*Request) Destination() (destination *RequestDestination) {
	macro.Rewrite("$_.destination")
	return destination
}

func (*Request) Headers() (headers *Headers) {
	macro.Rewrite("$_.headers")
	return headers
}

func (*Request) Integrity() (integrity string) {
	macro.Rewrite("$_.integrity")
	return integrity
}

func (*Request) Keepalive() (keepalive bool) {
	macro.Rewrite("$_.keepalive")
	return keepalive
}

func (*Request) Method() (method string) {
	macro.Rewrite("$_.method")
	return method
}

func (*Request) Mode() (mode *RequestMode) {
	macro.Rewrite("$_.mode")
	return mode
}

func (*Request) Redirect() (redirect *RequestRedirect) {
	macro.Rewrite("$_.redirect")
	return redirect
}

func (*Request) Referrer() (referrer string) {
	macro.Rewrite("$_.referrer")
	return referrer
}

func (*Request) ReferrerPolicy() (referrerPolicy *ReferrerPolicy) {
	macro.Rewrite("$_.referrerPolicy")
	return referrerPolicy
}

func (*Request) Type() (kind *RequestType) {
	macro.Rewrite("$_.type")
	return kind
}

func (*Request) URL() (url string) {
	macro.Rewrite("$_.url")
	return url
}

func (*Request) BodyUsed() (bodyUsed bool) {
	macro.Rewrite("$_.bodyUsed")
	return bodyUsed
}
