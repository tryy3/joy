package css

type CSSGroupingRule interface {
	CSSRule

	DeleteRule(index uint)

	InsertRule(rule string, index uint) (u uint)

	CSSRules() (cssRules *CSSRuleList)
}
