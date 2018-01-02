package css

type CSSRule interface {
	CSSText() (cssText string)

	SetCSSText(cssText string)

	ParentRule() (parentRule CSSRule)

	ParentStyleSheet() (parentStyleSheet *CSSStyleSheet)

	Type() (kind uint8)
}
