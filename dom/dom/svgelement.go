package dom

import "github.com/matthewmueller/joy/dom/element"

type SVGElement interface {
	element.
		Element

	Dataset() (dataset *DOMStringMap)

	OwnerSVGElement() (ownerSVGElement *SVGSVGElement)

	ViewportElement() (viewportElement SVGElement)

	Xmlbase() (xmlbase string)

	SetXmlbase(xmlbase string)
}
