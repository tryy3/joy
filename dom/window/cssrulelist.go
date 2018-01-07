package window

import "github.com/matthewmueller/joy/macro"

// CSSRuleList struct
// js:"CSSRuleList,omit"
type CSSRuleList struct {
}

// Item fn
// js:"item"
func (*CSSRuleList) Item(index uint) (c CSSRule) {
	macro.Rewrite("$_.item($1)", index)
	return c
}

// Length prop
// js:"length"
func (*CSSRuleList) Length() (length int) {
	macro.Rewrite("$_.length")
	return length
}
