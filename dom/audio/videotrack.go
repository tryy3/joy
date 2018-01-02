package audio

import "github.com/matthewmueller/joy/macro"

type VideoTrack struct {
}

func (*VideoTrack) ID() (id string) {
	macro.Rewrite("$_.id")
	return id
}

func (*VideoTrack) Kind() (kind string) {
	macro.Rewrite("$_.kind")
	return kind
}

func (*VideoTrack) SetKind(kind string) {
	macro.Rewrite("$_.kind = $1", kind)
}

func (*VideoTrack) Label() (label string) {
	macro.Rewrite("$_.label")
	return label
}

func (*VideoTrack) Language() (language string) {
	macro.Rewrite("$_.language")
	return language
}

func (*VideoTrack) SetLanguage(language string) {
	macro.Rewrite("$_.language = $1", language)
}

func (*VideoTrack) Selected() (selected bool) {
	macro.Rewrite("$_.selected")
	return selected
}

func (*VideoTrack) SetSelected(selected bool) {
	macro.Rewrite("$_.selected = $1", selected)
}

func (*VideoTrack) SourceBuffer() (sourceBuffer *SourceBuffer) {
	macro.Rewrite("$_.sourceBuffer")
	return sourceBuffer
}
