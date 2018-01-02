package webvvt

import "github.com/matthewmueller/joy/macro"

type TextTrackCueList struct {
}

func (*TextTrackCueList) GetCueByID(id string) (t TextTrackCue) {
	macro.Rewrite("$_.getCueById($1)", id)
	return t
}

func (*TextTrackCueList) Item(index uint) (t TextTrackCue) {
	macro.Rewrite("$_.item($1)", index)
	return t
}

func (*TextTrackCueList) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}
