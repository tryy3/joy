package css

import (
	"github.com/matthewmueller/joy/macro"
)

var _ CSSConditionRule = (*CSSSupportsRule)(nil)
var _ CSSGroupingRule = (*CSSSupportsRule)(nil)
var _ CSSRule = (*CSSSupportsRule)(nil)

type CSSSupportsRule struct {
}

func (*CSSSupportsRule) DeleteRule(index uint) {
	macro.Rewrite("$_.deleteRule($1)", index)
}

func (*CSSSupportsRule) InsertRule(rule string, index uint) (u uint) {
	macro.Rewrite("$_.insertRule($1, $2)", rule, index)
	return u
}

func (*CSSSupportsRule) ConditionText() (conditionText string) {
	macro.Rewrite("$_.conditionText")
	return conditionText
}

func (*CSSSupportsRule) SetConditionText(conditionText string) {
	macro.Rewrite("$_.conditionText = $1", conditionText)
}

func (*CSSSupportsRule) CSSRules() (cssRules *CSSRuleList) {
	macro.Rewrite("$_.cssRules")
	return cssRules
}

func (*CSSSupportsRule) CSSText() (cssText string) {
	macro.Rewrite("$_.cssText")
	return cssText
}

func (*CSSSupportsRule) SetCSSText(cssText string) {
	macro.Rewrite("$_.cssText = $1", cssText)
}

func (*CSSSupportsRule) ParentRule() (parentRule CSSRule) {
	macro.Rewrite("$_.parentRule")
	return parentRule
}

func (*CSSSupportsRule) ParentStyleSheet() (parentStyleSheet *CSSStyleSheet) {
	macro.Rewrite("$_.parentStyleSheet")
	return parentStyleSheet
}

func (*CSSSupportsRule) Type() (kind uint8) {
	macro.Rewrite("$_.type")
	return kind
}
