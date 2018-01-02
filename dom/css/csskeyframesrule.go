package css

import (
	"github.com/matthewmueller/joy/macro"
)

var _ CSSRule = (*CSSKeyframesRule)(nil)

type CSSKeyframesRule struct {
}

func (*CSSKeyframesRule) AppendRule(rule string) {
	macro.Rewrite("$_.appendRule($1)", rule)
}

func (*CSSKeyframesRule) DeleteRule(rule string) {
	macro.Rewrite("$_.deleteRule($1)", rule)
}

func (*CSSKeyframesRule) FindRule(rule string) (c *CSSKeyframeRule) {
	macro.Rewrite("$_.findRule($1)", rule)
	return c
}

func (*CSSKeyframesRule) CSSRules() (cssRules *CSSRuleList) {
	macro.Rewrite("$_.cssRules")
	return cssRules
}

func (*CSSKeyframesRule) Name() (name string) {
	macro.Rewrite("$_.name")
	return name
}

func (*CSSKeyframesRule) SetName(name string) {
	macro.Rewrite("$_.name = $1", name)
}

func (*CSSKeyframesRule) CSSText() (cssText string) {
	macro.Rewrite("$_.cssText")
	return cssText
}

func (*CSSKeyframesRule) SetCSSText(cssText string) {
	macro.Rewrite("$_.cssText = $1", cssText)
}

func (*CSSKeyframesRule) ParentRule() (parentRule CSSRule) {
	macro.Rewrite("$_.parentRule")
	return parentRule
}

func (*CSSKeyframesRule) ParentStyleSheet() (parentStyleSheet *CSSStyleSheet) {
	macro.Rewrite("$_.parentStyleSheet")
	return parentStyleSheet
}

func (*CSSKeyframesRule) Type() (kind uint8) {
	macro.Rewrite("$_.type")
	return kind
}
