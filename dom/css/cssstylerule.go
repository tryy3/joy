package css

import (
	"github.com/matthewmueller/joy/macro"
)

var _ CSSRule = (*CSSStyleRule)(nil)

type CSSStyleRule struct {
}

func (*CSSStyleRule) ReadOnly() (readOnly bool) {
	macro.Rewrite("$_.readOnly")
	return readOnly
}

func (*CSSStyleRule) SelectorText() (selectorText string) {
	macro.Rewrite("$_.selectorText")
	return selectorText
}

func (*CSSStyleRule) SetSelectorText(selectorText string) {
	macro.Rewrite("$_.selectorText = $1", selectorText)
}

func (*CSSStyleRule) Style() (style *CSSStyleDeclaration) {
	macro.Rewrite("$_.style")
	return style
}

func (*CSSStyleRule) CSSText() (cssText string) {
	macro.Rewrite("$_.cssText")
	return cssText
}

func (*CSSStyleRule) SetCSSText(cssText string) {
	macro.Rewrite("$_.cssText = $1", cssText)
}

func (*CSSStyleRule) ParentRule() (parentRule CSSRule) {
	macro.Rewrite("$_.parentRule")
	return parentRule
}

func (*CSSStyleRule) ParentStyleSheet() (parentStyleSheet *CSSStyleSheet) {
	macro.Rewrite("$_.parentStyleSheet")
	return parentStyleSheet
}

func (*CSSStyleRule) Type() (kind uint8) {
	macro.Rewrite("$_.type")
	return kind
}
