package css

import (
	"github.com/matthewmueller/joy/macro"
)

var _ CSSRule = (*CSSNamespaceRule)(nil)

type CSSNamespaceRule struct {
}

func (*CSSNamespaceRule) NamespaceURI() (namespaceURI string) {
	macro.Rewrite("$_.namespaceURI")
	return namespaceURI
}

func (*CSSNamespaceRule) Prefix() (prefix string) {
	macro.Rewrite("$_.prefix")
	return prefix
}

func (*CSSNamespaceRule) CSSText() (cssText string) {
	macro.Rewrite("$_.cssText")
	return cssText
}

func (*CSSNamespaceRule) SetCSSText(cssText string) {
	macro.Rewrite("$_.cssText = $1", cssText)
}

func (*CSSNamespaceRule) ParentRule() (parentRule CSSRule) {
	macro.Rewrite("$_.parentRule")
	return parentRule
}

func (*CSSNamespaceRule) ParentStyleSheet() (parentStyleSheet *CSSStyleSheet) {
	macro.Rewrite("$_.parentStyleSheet")
	return parentStyleSheet
}

func (*CSSNamespaceRule) Type() (kind uint8) {
	macro.Rewrite("$_.type")
	return kind
}
