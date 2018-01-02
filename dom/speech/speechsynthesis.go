package speech

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
)

var _ event.EventTarget = (*SpeechSynthesis)(nil)

type SpeechSynthesis struct {
}

func (*SpeechSynthesis) Cancel() {
	macro.Rewrite("$_.cancel()")
}

func (*SpeechSynthesis) GetVoices() (s []*SpeechSynthesisVoice) {
	macro.Rewrite("$_.getVoices()")
	return s
}

func (*SpeechSynthesis) Pause() {
	macro.Rewrite("$_.pause()")
}

func (*SpeechSynthesis) Resume() {
	macro.Rewrite("$_.resume()")
}

func (*SpeechSynthesis) Speak(utterance *SpeechSynthesisUtterance) {
	macro.Rewrite("$_.speak($1)", utterance)
}

func (*SpeechSynthesis) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*SpeechSynthesis) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*SpeechSynthesis) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*SpeechSynthesis) Onvoiceschanged() (onvoiceschanged func(event.Event)) {
	macro.Rewrite("$_.onvoiceschanged")
	return onvoiceschanged
}

func (*SpeechSynthesis) SetOnvoiceschanged(onvoiceschanged func(event.Event)) {
	macro.Rewrite("$_.onvoiceschanged = $1", onvoiceschanged)
}

func (*SpeechSynthesis) Paused() (paused bool) {
	macro.Rewrite("$_.paused")
	return paused
}

func (*SpeechSynthesis) Pending() (pending bool) {
	macro.Rewrite("$_.pending")
	return pending
}

func (*SpeechSynthesis) Speaking() (speaking bool) {
	macro.Rewrite("$_.speaking")
	return speaking
}
