package css

import (
	"github.com/matthewmueller/joy/macro"
)

var _ CSSRule = (*CSSImportRule)(nil)

type CSSImportRule struct {
}

func (*CSSImportRule) Href() (href string) {
	macro.Rewrite("$_.href")
	return href
}

func (*CSSImportRule) Media() (media *MediaList) {
	macro.Rewrite("$_.media")
	return media
}

func (*CSSImportRule) StyleSheet() (styleSheet *CSSStyleSheet) {
	macro.Rewrite("$_.styleSheet")
	return styleSheet
}

func (*CSSImportRule) CSSText() (cssText string) {
	macro.Rewrite("$_.cssText")
	return cssText
}

func (*CSSImportRule) SetCSSText(cssText string) {
	macro.Rewrite("$_.cssText = $1", cssText)
}

func (*CSSImportRule) ParentRule() (parentRule CSSRule) {
	macro.Rewrite("$_.parentRule")
	return parentRule
}

func (*CSSImportRule) ParentStyleSheet() (parentStyleSheet *CSSStyleSheet) {
	macro.Rewrite("$_.parentStyleSheet")
	return parentStyleSheet
}

func (*CSSImportRule) Type() (kind uint8) {
	macro.Rewrite("$_.type")
	return kind
}
