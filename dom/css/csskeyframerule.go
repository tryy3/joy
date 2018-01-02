package css

import (
	"github.com/matthewmueller/joy/macro"
)

var _ CSSRule = (*CSSKeyframeRule)(nil)

type CSSKeyframeRule struct {
}

func (*CSSKeyframeRule) KeyText() (keyText string) {
	macro.Rewrite("$_.keyText")
	return keyText
}

func (*CSSKeyframeRule) SetKeyText(keyText string) {
	macro.Rewrite("$_.keyText = $1", keyText)
}

func (*CSSKeyframeRule) Style() (style *CSSStyleDeclaration) {
	macro.Rewrite("$_.style")
	return style
}

func (*CSSKeyframeRule) CSSText() (cssText string) {
	macro.Rewrite("$_.cssText")
	return cssText
}

func (*CSSKeyframeRule) SetCSSText(cssText string) {
	macro.Rewrite("$_.cssText = $1", cssText)
}

func (*CSSKeyframeRule) ParentRule() (parentRule CSSRule) {
	macro.Rewrite("$_.parentRule")
	return parentRule
}

func (*CSSKeyframeRule) ParentStyleSheet() (parentStyleSheet *CSSStyleSheet) {
	macro.Rewrite("$_.parentStyleSheet")
	return parentStyleSheet
}

func (*CSSKeyframeRule) Type() (kind uint8) {
	macro.Rewrite("$_.type")
	return kind
}
