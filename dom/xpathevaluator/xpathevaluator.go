package xpathevaluator

import (
	"github.com/matthewmueller/joy/dom/utils"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

// New fn
func New() *XPathEvaluator {
	macro.Rewrite("new XPathEvaluator()")
	return &XPathEvaluator{}
}

// XPathEvaluator struct
// js:"XPathEvaluator,omit"
type XPathEvaluator struct {
}

// CreateExpression fn
// js:"createExpression"
func (*XPathEvaluator) CreateExpression(expression string, resolver *utils.XPathNSResolver) (w *window.XPathExpression) {
	macro.Rewrite("$_.createExpression($1, $2)", expression, resolver)
	return w
}

// CreateNSResolver fn
// js:"createNSResolver"
func (*XPathEvaluator) CreateNSResolver(nodeResolver *window.Node) (u *utils.XPathNSResolver) {
	macro.Rewrite("$_.createNSResolver($1)", nodeResolver)
	return u
}

// Evaluate fn
// js:"evaluate"
func (*XPathEvaluator) Evaluate(expression string, contextNode window.Node, resolver *utils.XPathNSResolver, kind uint8, result *window.XPathResult) (w *window.XPathResult) {
	macro.Rewrite("$_.evaluate($1, $2, $3, $4, $5)", expression, contextNode, resolver, kind, result)
	return w
}
