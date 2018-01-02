package css

import "github.com/matthewmueller/joy/macro"

type CSSRuleList struct {
}

func (*CSSRuleList) Item(index uint) (c CSSRule) {
	macro.Rewrite("$_.item($1)", index)
	return c
}

func (*CSSRuleList) Length() (length int) {
	macro.Rewrite("$_.length")
	return length
}
