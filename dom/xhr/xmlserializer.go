package xhr

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/dom"
)

func New() *XMLSerializer {
	macro.Rewrite("new XMLSerializer()")
	return &XMLSerializer{}
}

type XMLSerializer struct {
}

func (*XMLSerializer) SerializeToString(target dom.Node) (s string) {
	macro.Rewrite("$_.serializeToString($1)", target)
	return s
}
