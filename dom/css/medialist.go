package css

import "github.com/matthewmueller/joy/macro"

type MediaList struct {
}

func (*MediaList) AppendMedium(newMedium string) {
	macro.Rewrite("$_.appendMedium($1)", newMedium)
}

func (*MediaList) DeleteMedium(oldMedium string) {
	macro.Rewrite("$_.deleteMedium($1)", oldMedium)
}

func (*MediaList) Item(index uint) (s string) {
	macro.Rewrite("$_.item($1)", index)
	return s
}

func (*MediaList) ToString() (s string) {
	macro.Rewrite("$_.toString()")
	return s
}

func (*MediaList) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}

func (*MediaList) MediaText() (mediaText string) {
	macro.Rewrite("$_.mediaText")
	return mediaText
}

func (*MediaList) SetMediaText(mediaText string) {
	macro.Rewrite("$_.mediaText = $1", mediaText)
}
