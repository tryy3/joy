package cssstylerule

import (
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ window.CSSRule = (*CSSStyleRule)(nil)

// CSSStyleRule struct
// js:"CSSStyleRule,omit"
type CSSStyleRule struct {
}

// ReadOnly prop
// js:"readOnly"
func (*CSSStyleRule) ReadOnly() (readOnly bool) {
	macro.Rewrite("$_.readOnly")
	return readOnly
}

// SelectorText prop
// js:"selectorText"
func (*CSSStyleRule) SelectorText() (selectorText string) {
	macro.Rewrite("$_.selectorText")
	return selectorText
}

// SetSelectorText prop
// js:"selectorText"
func (*CSSStyleRule) SetSelectorText(selectorText string) {
	macro.Rewrite("$_.selectorText = $1", selectorText)
}

// Style prop
// js:"style"
func (*CSSStyleRule) Style() (style *window.CSSStyleDeclaration) {
	macro.Rewrite("$_.style")
	return style
}

// CSSText prop
// js:"cssText"
func (*CSSStyleRule) CSSText() (cssText string) {
	macro.Rewrite("$_.cssText")
	return cssText
}

// SetCSSText prop
// js:"cssText"
func (*CSSStyleRule) SetCSSText(cssText string) {
	macro.Rewrite("$_.cssText = $1", cssText)
}

// ParentRule prop
// js:"parentRule"
func (*CSSStyleRule) ParentRule() (parentRule window.CSSRule) {
	macro.Rewrite("$_.parentRule")
	return parentRule
}

// ParentStyleSheet prop
// js:"parentStyleSheet"
func (*CSSStyleRule) ParentStyleSheet() (parentStyleSheet *window.CSSStyleSheet) {
	macro.Rewrite("$_.parentStyleSheet")
	return parentStyleSheet
}

// Type prop
// js:"type"
func (*CSSStyleRule) Type() (kind uint8) {
	macro.Rewrite("$_.type")
	return kind
}
