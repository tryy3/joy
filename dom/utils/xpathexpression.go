package utils

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/dom"
)

type XPathExpression struct {
}

func (*XPathExpression) Evaluate(contextNode dom.Node, kind uint8, result *XPathResult) (x *XPathExpression) {
	macro.Rewrite("$_.evaluate($1, $2, $3)", contextNode, kind, result)
	return x
}
