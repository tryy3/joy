package css

import (
	"github.com/matthewmueller/joy/macro"
)

var _ CSSConditionRule = (*CSSMediaRule)(nil)
var _ CSSGroupingRule = (*CSSMediaRule)(nil)
var _ CSSRule = (*CSSMediaRule)(nil)

type CSSMediaRule struct {
}

func (*CSSMediaRule) DeleteRule(index uint) {
	macro.Rewrite("$_.deleteRule($1)", index)
}

func (*CSSMediaRule) InsertRule(rule string, index uint) (u uint) {
	macro.Rewrite("$_.insertRule($1, $2)", rule, index)
	return u
}

func (*CSSMediaRule) Media() (media *MediaList) {
	macro.Rewrite("$_.media")
	return media
}

func (*CSSMediaRule) ConditionText() (conditionText string) {
	macro.Rewrite("$_.conditionText")
	return conditionText
}

func (*CSSMediaRule) SetConditionText(conditionText string) {
	macro.Rewrite("$_.conditionText = $1", conditionText)
}

func (*CSSMediaRule) CSSRules() (cssRules *CSSRuleList) {
	macro.Rewrite("$_.cssRules")
	return cssRules
}

func (*CSSMediaRule) CSSText() (cssText string) {
	macro.Rewrite("$_.cssText")
	return cssText
}

func (*CSSMediaRule) SetCSSText(cssText string) {
	macro.Rewrite("$_.cssText = $1", cssText)
}

func (*CSSMediaRule) ParentRule() (parentRule CSSRule) {
	macro.Rewrite("$_.parentRule")
	return parentRule
}

func (*CSSMediaRule) ParentStyleSheet() (parentStyleSheet *CSSStyleSheet) {
	macro.Rewrite("$_.parentStyleSheet")
	return parentStyleSheet
}

func (*CSSMediaRule) Type() (kind uint8) {
	macro.Rewrite("$_.type")
	return kind
}
