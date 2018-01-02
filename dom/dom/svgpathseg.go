package dom

type SVGPathSeg interface {
	PathSegType() (pathSegType uint8)

	PathSegTypeAsLetter() (pathSegTypeAsLetter string)
}
