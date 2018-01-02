package dom

type SVGGraphicsElement interface {
	SVGElement

	HasExtension(extension string) (b bool)

	GetBBox() (s *SVGRect)

	GetCTM() (s *SVGMatrix)

	GetScreenCTM() (s *SVGMatrix)

	GetTransformToElement(element SVGElement) (s *SVGMatrix)

	RequiredExtensions() (requiredExtensions *SVGStringList)

	RequiredFeatures() (requiredFeatures *SVGStringList)

	SystemLanguage() (systemLanguage *SVGStringList)

	FarthestViewportElement() (farthestViewportElement SVGElement)

	NearestViewportElement() (nearestViewportElement SVGElement)

	Transform() (transform *SVGAnimatedTransformList)
}
