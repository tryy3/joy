package speech

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
)

var _ event.EventTarget = (*SpeechSynthesisUtterance)(nil)

func NewSpeechSynthesisUtterance(text *string) *SpeechSynthesisUtterance {
	macro.Rewrite("new SpeechSynthesisUtterance($1)", text)
	return &SpeechSynthesisUtterance{}
}

type SpeechSynthesisUtterance struct {
}

func (*SpeechSynthesisUtterance) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*SpeechSynthesisUtterance) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*SpeechSynthesisUtterance) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*SpeechSynthesisUtterance) Lang() (lang string) {
	macro.Rewrite("$_.lang")
	return lang
}

func (*SpeechSynthesisUtterance) SetLang(lang string) {
	macro.Rewrite("$_.lang = $1", lang)
}

func (*SpeechSynthesisUtterance) Onboundary() (onboundary func(event.Event)) {
	macro.Rewrite("$_.onboundary")
	return onboundary
}

func (*SpeechSynthesisUtterance) SetOnboundary(onboundary func(event.Event)) {
	macro.Rewrite("$_.onboundary = $1", onboundary)
}

func (*SpeechSynthesisUtterance) Onend() (onend func(event.Event)) {
	macro.Rewrite("$_.onend")
	return onend
}

func (*SpeechSynthesisUtterance) SetOnend(onend func(event.Event)) {
	macro.Rewrite("$_.onend = $1", onend)
}

func (*SpeechSynthesisUtterance) Onerror() (onerror func(event.Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

func (*SpeechSynthesisUtterance) SetOnerror(onerror func(event.Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

func (*SpeechSynthesisUtterance) Onmark() (onmark func(event.Event)) {
	macro.Rewrite("$_.onmark")
	return onmark
}

func (*SpeechSynthesisUtterance) SetOnmark(onmark func(event.Event)) {
	macro.Rewrite("$_.onmark = $1", onmark)
}

func (*SpeechSynthesisUtterance) Onpause() (onpause func(event.Event)) {
	macro.Rewrite("$_.onpause")
	return onpause
}

func (*SpeechSynthesisUtterance) SetOnpause(onpause func(event.Event)) {
	macro.Rewrite("$_.onpause = $1", onpause)
}

func (*SpeechSynthesisUtterance) Onresume() (onresume func(event.Event)) {
	macro.Rewrite("$_.onresume")
	return onresume
}

func (*SpeechSynthesisUtterance) SetOnresume(onresume func(event.Event)) {
	macro.Rewrite("$_.onresume = $1", onresume)
}

func (*SpeechSynthesisUtterance) Onstart() (onstart func(event.Event)) {
	macro.Rewrite("$_.onstart")
	return onstart
}

func (*SpeechSynthesisUtterance) SetOnstart(onstart func(event.Event)) {
	macro.Rewrite("$_.onstart = $1", onstart)
}

func (*SpeechSynthesisUtterance) Pitch() (pitch float32) {
	macro.Rewrite("$_.pitch")
	return pitch
}

func (*SpeechSynthesisUtterance) SetPitch(pitch float32) {
	macro.Rewrite("$_.pitch = $1", pitch)
}

func (*SpeechSynthesisUtterance) Rate() (rate float32) {
	macro.Rewrite("$_.rate")
	return rate
}

func (*SpeechSynthesisUtterance) SetRate(rate float32) {
	macro.Rewrite("$_.rate = $1", rate)
}

func (*SpeechSynthesisUtterance) Text() (text string) {
	macro.Rewrite("$_.text")
	return text
}

func (*SpeechSynthesisUtterance) SetText(text string) {
	macro.Rewrite("$_.text = $1", text)
}

func (*SpeechSynthesisUtterance) Voice() (voice *SpeechSynthesisVoice) {
	macro.Rewrite("$_.voice")
	return voice
}

func (*SpeechSynthesisUtterance) SetVoice(voice *SpeechSynthesisVoice) {
	macro.Rewrite("$_.voice = $1", voice)
}

func (*SpeechSynthesisUtterance) Volume() (volume float32) {
	macro.Rewrite("$_.volume")
	return volume
}

func (*SpeechSynthesisUtterance) SetVolume(volume float32) {
	macro.Rewrite("$_.volume = $1", volume)
}
