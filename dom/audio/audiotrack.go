package audio

import "github.com/matthewmueller/joy/macro"

type AudioTrack struct {
}

func (*AudioTrack) Enabled() (enabled bool) {
	macro.Rewrite("$_.enabled")
	return enabled
}

func (*AudioTrack) SetEnabled(enabled bool) {
	macro.Rewrite("$_.enabled = $1", enabled)
}

func (*AudioTrack) ID() (id string) {
	macro.Rewrite("$_.id")
	return id
}

func (*AudioTrack) Kind() (kind string) {
	macro.Rewrite("$_.kind")
	return kind
}

func (*AudioTrack) SetKind(kind string) {
	macro.Rewrite("$_.kind = $1", kind)
}

func (*AudioTrack) Label() (label string) {
	macro.Rewrite("$_.label")
	return label
}

func (*AudioTrack) Language() (language string) {
	macro.Rewrite("$_.language")
	return language
}

func (*AudioTrack) SetLanguage(language string) {
	macro.Rewrite("$_.language = $1", language)
}

func (*AudioTrack) SourceBuffer() (sourceBuffer *SourceBuffer) {
	macro.Rewrite("$_.sourceBuffer")
	return sourceBuffer
}
