package dom

import (
	"github.com/matthewmueller/joy/macro"
)

var _ DOMTokenList = (*DOMSettableTokenList)(nil)

type DOMSettableTokenList struct {
}

func (*DOMSettableTokenList) Add(token string) {
	macro.Rewrite("$_.add($1)", token)
}

func (*DOMSettableTokenList) Contains(token string) (b bool) {
	macro.Rewrite("$_.contains($1)", token)
	return b
}

func (*DOMSettableTokenList) Item(index uint) (s string) {
	macro.Rewrite("$_.item($1)", index)
	return s
}

func (*DOMSettableTokenList) Remove(token string) {
	macro.Rewrite("$_.remove($1)", token)
}

func (*DOMSettableTokenList) Toggle(token string, force *bool) (b bool) {
	macro.Rewrite("$_.toggle($1, $2)", token, force)
	return b
}

func (*DOMSettableTokenList) ToString() (s string) {
	macro.Rewrite("$_.toString()")
	return s
}

func (*DOMSettableTokenList) Value() (value string) {
	macro.Rewrite("$_.value")
	return value
}

func (*DOMSettableTokenList) SetValue(value string) {
	macro.Rewrite("$_.value = $1", value)
}

func (*DOMSettableTokenList) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}
