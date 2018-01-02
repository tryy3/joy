package css

import (
	"github.com/matthewmueller/joy/macro"
)

var _ CSSRule = (*CSSFontFaceRule)(nil)

type CSSFontFaceRule struct {
}

func (*CSSFontFaceRule) Style() (style *CSSStyleDeclaration) {
	macro.Rewrite("$_.style")
	return style
}

func (*CSSFontFaceRule) CSSText() (cssText string) {
	macro.Rewrite("$_.cssText")
	return cssText
}

func (*CSSFontFaceRule) SetCSSText(cssText string) {
	macro.Rewrite("$_.cssText = $1", cssText)
}

func (*CSSFontFaceRule) ParentRule() (parentRule CSSRule) {
	macro.Rewrite("$_.parentRule")
	return parentRule
}

func (*CSSFontFaceRule) ParentStyleSheet() (parentStyleSheet *CSSStyleSheet) {
	macro.Rewrite("$_.parentStyleSheet")
	return parentStyleSheet
}

func (*CSSFontFaceRule) Type() (kind uint8) {
	macro.Rewrite("$_.type")
	return kind
}
