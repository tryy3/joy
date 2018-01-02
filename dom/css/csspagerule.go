package css

import "github.com/matthewmueller/joy/macro"

var _ CSSRule = (*CSSPageRule)(nil)

type CSSPageRule struct {
}

func (*CSSPageRule) PseudoClass() (pseudoClass string) {
	macro.Rewrite("$_.pseudoClass")
	return pseudoClass
}

func (*CSSPageRule) Selector() (selector string) {
	macro.Rewrite("$_.selector")
	return selector
}

func (*CSSPageRule) SelectorText() (selectorText string) {
	macro.Rewrite("$_.selectorText")
	return selectorText
}

func (*CSSPageRule) SetSelectorText(selectorText string) {
	macro.Rewrite("$_.selectorText = $1", selectorText)
}

func (*CSSPageRule) Style() (style *CSSStyleDeclaration) {
	macro.Rewrite("$_.style")
	return style
}

func (*CSSPageRule) CSSText() (cssText string) {
	macro.Rewrite("$_.cssText")
	return cssText
}

func (*CSSPageRule) SetCSSText(cssText string) {
	macro.Rewrite("$_.cssText = $1", cssText)
}

func (*CSSPageRule) ParentRule() (parentRule CSSRule) {
	macro.Rewrite("$_.parentRule")
	return parentRule
}

func (*CSSPageRule) ParentStyleSheet() (parentStyleSheet *CSSStyleSheet) {
	macro.Rewrite("$_.parentStyleSheet")
	return parentStyleSheet
}

func (*CSSPageRule) Type() (kind uint8) {
	macro.Rewrite("$_.type")
	return kind
}
