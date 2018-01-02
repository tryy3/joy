package html

import "github.com/matthewmueller/joy/macro"

type HTMLAllCollection struct {
}

func (*HTMLAllCollection) Item(nameOrIndex *string) (i interface{}) {
	macro.Rewrite("$_.item($1)", nameOrIndex)
	return i
}

func (*HTMLAllCollection) NamedItem(name string) (i interface{}) {
	macro.Rewrite("$_.namedItem($1)", name)
	return i
}

func (*HTMLAllCollection) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}
