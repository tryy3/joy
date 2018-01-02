package utils

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/dom"
)

func New() *XPathEvaluator {
	macro.Rewrite("new XPathEvaluator()")
	return &XPathEvaluator{}
}

type XPathEvaluator struct {
}

func (*XPathEvaluator) CreateExpression(expression string, resolver *XPathNSResolver) (w *XPathExpression) {
	macro.Rewrite("$_.createExpression($1, $2)", expression, resolver)
	return w
}

func (*XPathEvaluator) CreateNSResolver(nodeResolver *dom.Node) (x *XPathNSResolver) {
	macro.Rewrite("$_.createNSResolver($1)", nodeResolver)
	return x
}

func (*XPathEvaluator) Evaluate(expression string, contextNode dom.Node, resolver *XPathNSResolver, kind uint8, result *XPathResult) (w *XPathResult) {
	macro.Rewrite("$_.evaluate($1, $2, $3, $4, $5)", expression, contextNode, resolver, kind, result)
	return w
}
