package fetch


import (
	"github.com/matthewmueller/joy/dom/file"
	"github.com/matthewmueller/joy/macro"
)

// ReferrerPolicy type
type ReferrerPolicy string

// RequestCache type
type RequestCache string

// RequestCredentials type
type RequestCredentials string

// RequestDestination type
type RequestDestination string

// RequestMode type
type RequestMode string

// RequestRedirect type
type RequestRedirect string

// RequestType type
type RequestType string

// RequestInit type
type RequestInit struct {
	body           *interface{}
	cache          *RequestCache
	credentials    *RequestCredentials
	headers        *interface{}
	integrity      *string
	keepalive      *bool
	method         *string
	mode           *RequestMode
	redirect       *RequestRedirect
	referrer       *string
	referrerPolicy *ReferrerPolicy
	window         *interface{}
}

// NewRequest fn
func NewRequest(input interface{}, init *RequestInit) *Request {
	macro.Rewrite("new Request($1, $2)", input, init)
	return &Request{}
}

// Request struct
// js:"Request,omit"
type Request struct {
}

// Clone fn
// js:"clone"
func (*Request) Clone() (r *Request) {
	macro.Rewrite("$_.clone()")
	return r
}

// ArrayBuffer fn
// js:"arrayBuffer"
func (*Request) ArrayBuffer() (b []byte) {
	macro.Rewrite("await $_.arrayBuffer()")
	return b
}

// Blob fn
// js:"blob"
func (*Request) Blob() (b file.Blob) {
	macro.Rewrite("await $_.blob()")
	return b
}

// JSON fn
// js:"json"
func (*Request) JSON() (i interface{}) {
	macro.Rewrite("await $_.json()")
	return i
}

// Text fn
// js:"text"
func (*Request) Text() (s string) {
	macro.Rewrite("await $_.text()")
	return s
}

// Cache prop
// js:"cache"
func (*Request) Cache() (cache *RequestCache) {
	macro.Rewrite("$_.cache")
	return cache
}

// Credentials prop
// js:"credentials"
func (*Request) Credentials() (credentials *RequestCredentials) {
	macro.Rewrite("$_.credentials")
	return credentials
}

// Destination prop
// js:"destination"
func (*Request) Destination() (destination *RequestDestination) {
	macro.Rewrite("$_.destination")
	return destination
}

// Headers prop
// js:"headers"
func (*Request) Headers() (headers *Headers) {
	macro.Rewrite("$_.headers")
	return headers
}

// Integrity prop
// js:"integrity"
func (*Request) Integrity() (integrity string) {
	macro.Rewrite("$_.integrity")
	return integrity
}

// Keepalive prop
// js:"keepalive"
func (*Request) Keepalive() (keepalive bool) {
	macro.Rewrite("$_.keepalive")
	return keepalive
}

// Method prop
// js:"method"
func (*Request) Method() (method string) {
	macro.Rewrite("$_.method")
	return method
}

// Mode prop
// js:"mode"
func (*Request) Mode() (mode *RequestMode) {
	macro.Rewrite("$_.mode")
	return mode
}

// Redirect prop
// js:"redirect"
func (*Request) Redirect() (redirect *RequestRedirect) {
	macro.Rewrite("$_.redirect")
	return redirect
}

// Referrer prop
// js:"referrer"
func (*Request) Referrer() (referrer string) {
	macro.Rewrite("$_.referrer")
	return referrer
}

// ReferrerPolicy prop
// js:"referrerPolicy"
func (*Request) ReferrerPolicy() (referrerPolicy *ReferrerPolicy) {
	macro.Rewrite("$_.referrerPolicy")
	return referrerPolicy
}

// Type prop
// js:"type"
func (*Request) Type() (kind *RequestType) {
	macro.Rewrite("$_.type")
	return kind
}

// URL prop
// js:"url"
func (*Request) URL() (url string) {
	macro.Rewrite("$_.url")
	return url
}

// BodyUsed prop
// js:"bodyUsed"
func (*Request) BodyUsed() (bodyUsed bool) {
	macro.Rewrite("$_.bodyUsed")
	return bodyUsed
}
