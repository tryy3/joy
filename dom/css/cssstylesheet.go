package css

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/element"
	"github.com/matthewmueller/joy/dom/dom"
)

var _ StyleSheet = (*CSSStyleSheet)(nil)

type CSSStyleSheet struct {
}

func (*CSSStyleSheet) AddImport(bstrURL string, lIndex *int) (i int) {
	macro.Rewrite("$_.addImport($1, $2)", bstrURL, lIndex)
	return i
}

func (*CSSStyleSheet) AddPageRule(bstrSelector string, bstrStyle string, lIndex *int) (i int) {
	macro.Rewrite("$_.addPageRule($1, $2, $3)", bstrSelector, bstrStyle, lIndex)
	return i
}

func (*CSSStyleSheet) AddRule(bstrSelector string, bstrStyle *string, lIndex *int) (i int) {
	macro.Rewrite("$_.addRule($1, $2, $3)", bstrSelector, bstrStyle, lIndex)
	return i
}

func (*CSSStyleSheet) DeleteRule(index *uint) {
	macro.Rewrite("$_.deleteRule($1)", index)
}

func (*CSSStyleSheet) InsertRule(rule string, index *uint) (u uint) {
	macro.Rewrite("$_.insertRule($1, $2)", rule, index)
	return u
}

func (*CSSStyleSheet) RemoveImport(lIndex int) {
	macro.Rewrite("$_.removeImport($1)", lIndex)
}

func (*CSSStyleSheet) RemoveRule(lIndex int) {
	macro.Rewrite("$_.removeRule($1)", lIndex)
}

func (*CSSStyleSheet) CSSRules() (cssRules *CSSRuleList) {
	macro.Rewrite("$_.cssRules")
	return cssRules
}

func (*CSSStyleSheet) CSSText() (cssText string) {
	macro.Rewrite("$_.cssText")
	return cssText
}

func (*CSSStyleSheet) SetCSSText(cssText string) {
	macro.Rewrite("$_.cssText = $1", cssText)
}

func (*CSSStyleSheet) ID() (id string) {
	macro.Rewrite("$_.id")
	return id
}

func (*CSSStyleSheet) Imports() (imports *StyleSheetList) {
	macro.Rewrite("$_.imports")
	return imports
}

func (*CSSStyleSheet) IsAlternate() (isAlternate bool) {
	macro.Rewrite("$_.isAlternate")
	return isAlternate
}

func (*CSSStyleSheet) IsPrefAlternate() (isPrefAlternate bool) {
	macro.Rewrite("$_.isPrefAlternate")
	return isPrefAlternate
}

func (*CSSStyleSheet) OwnerRule() (ownerRule CSSRule) {
	macro.Rewrite("$_.ownerRule")
	return ownerRule
}

func (*CSSStyleSheet) OwningElement() (owningElement element.Element) {
	macro.Rewrite("$_.owningElement")
	return owningElement
}

func (*CSSStyleSheet) Pages() (pages *StyleSheetPageList) {
	macro.Rewrite("$_.pages")
	return pages
}

func (*CSSStyleSheet) ReadOnly() (readOnly bool) {
	macro.Rewrite("$_.readOnly")
	return readOnly
}

func (*CSSStyleSheet) Rules() (rules *CSSRuleList) {
	macro.Rewrite("$_.rules")
	return rules
}

func (*CSSStyleSheet) Disabled() (disabled bool) {
	macro.Rewrite("$_.disabled")
	return disabled
}

func (*CSSStyleSheet) SetDisabled(disabled bool) {
	macro.Rewrite("$_.disabled = $1", disabled)
}

func (*CSSStyleSheet) Href() (href string) {
	macro.Rewrite("$_.href")
	return href
}

func (*CSSStyleSheet) Media() (media *MediaList) {
	macro.Rewrite("$_.media")
	return media
}

func (*CSSStyleSheet) OwnerNode() (ownerNode dom.Node) {
	macro.Rewrite("$_.ownerNode")
	return ownerNode
}

func (*CSSStyleSheet) ParentStyleSheet() (parentStyleSheet StyleSheet) {
	macro.Rewrite("$_.parentStyleSheet")
	return parentStyleSheet
}

func (*CSSStyleSheet) Title() (title string) {
	macro.Rewrite("$_.title")
	return title
}

func (*CSSStyleSheet) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}
