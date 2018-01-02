package fetch

import (
	"github.com/matthewmueller/joy/macro"
)

func NewHeaders(init *interface{}) *Headers {
	macro.Rewrite("new Headers($1)", init)
	return &Headers{}
}

type Headers struct {
}

func (*Headers) Append(name string, value string) {
	macro.Rewrite("$_.append($1, $2)", name, value)
}

func (*Headers) Delete(name string) {
	macro.Rewrite("$_.delete($1)", name)
}

func (*Headers) ForEach(callback func(keyId []byte, status *string)) {
	macro.Rewrite("$_.forEach($1)", callback)
}

func (*Headers) Get(name string) (s string) {
	macro.Rewrite("$_.get($1)", name)
	return s
}

func (*Headers) Has(name string) (b bool) {
	macro.Rewrite("$_.has($1)", name)
	return b
}

func (*Headers) Set(name string, value string) {
	macro.Rewrite("$_.set($1, $2)", name, value)
}
