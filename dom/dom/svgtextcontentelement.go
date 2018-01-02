package dom

type SVGTextContentElement interface {
	SVGGraphicsElement

	GetCharNumAtPosition(point *SVGPoint) (i int)

	GetComputedTextLength() (f float32)

	GetEndPositionOfChar(charnum uint) (s *SVGPoint)

	GetExtentOfChar(charnum uint) (s *SVGRect)

	GetNumberOfChars() (i int)

	GetRotationOfChar(charnum uint) (f float32)

	GetStartPositionOfChar(charnum uint) (s *SVGPoint)

	GetSubStringLength(charnum uint, nchars uint) (f float32)

	SelectSubString(charnum uint, nchars uint)

	LengthAdjust() (lengthAdjust *SVGAnimatedEnumeration)

	TextLength() (textLength *SVGAnimatedLength)
}
