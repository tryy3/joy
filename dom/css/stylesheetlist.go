package css

import "github.com/matthewmueller/joy/macro"

type StyleSheetList struct {
}

func (*StyleSheetList) Item(index *uint) (s StyleSheet) {
	macro.Rewrite("$_.item($1)", index)
	return s
}

func (*StyleSheetList) Length() (length int) {
	macro.Rewrite("$_.length")
	return length
}
