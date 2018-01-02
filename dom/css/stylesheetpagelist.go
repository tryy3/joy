package css

import "github.com/matthewmueller/joy/macro"

type StyleSheetPageList struct {
}

func (*StyleSheetPageList) Item(index uint) (c *CSSPageRule) {
	macro.Rewrite("$_.item($1)", index)
	return c
}

func (*StyleSheetPageList) Length() (length int) {
	macro.Rewrite("$_.length")
	return length
}
