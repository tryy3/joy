package css

type CSSConditionRule interface {
	CSSGroupingRule

	ConditionText() (conditionText string)

	SetConditionText(conditionText string)
}
