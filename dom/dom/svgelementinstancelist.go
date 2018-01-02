package dom

import "github.com/matthewmueller/joy/macro"

type SVGElementInstanceList struct {
}

func (*SVGElementInstanceList) Item(index uint) (s *SVGElementInstance) {
	macro.Rewrite("$_.item($1)", index)
	return s
}

func (*SVGElementInstanceList) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}
