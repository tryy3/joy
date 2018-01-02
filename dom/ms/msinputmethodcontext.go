package ms

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/dom"
	"github.com/matthewmueller/joy/dom/html"
)

var _ event.EventTarget = (*MSInputMethodContext)(nil)

type MSInputMethodContext struct {
}

func (*MSInputMethodContext) GetCandidateWindowClientRect() (c *dom.ClientRect) {
	macro.Rewrite("$_.getCandidateWindowClientRect()")
	return c
}

func (*MSInputMethodContext) GetCompositionAlternatives() (s []string) {
	macro.Rewrite("$_.getCompositionAlternatives()")
	return s
}

func (*MSInputMethodContext) HasComposition() (b bool) {
	macro.Rewrite("$_.hasComposition()")
	return b
}

func (*MSInputMethodContext) IsCandidateWindowVisible() (b bool) {
	macro.Rewrite("$_.isCandidateWindowVisible()")
	return b
}

func (*MSInputMethodContext) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*MSInputMethodContext) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*MSInputMethodContext) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*MSInputMethodContext) CompositionEndOffset() (compositionEndOffset uint) {
	macro.Rewrite("$_.compositionEndOffset")
	return compositionEndOffset
}

func (*MSInputMethodContext) CompositionStartOffset() (compositionStartOffset uint) {
	macro.Rewrite("$_.compositionStartOffset")
	return compositionStartOffset
}

func (*MSInputMethodContext) Oncandidatewindowhide() (oncandidatewindowhide func(event.Event)) {
	macro.Rewrite("$_.oncandidatewindowhide")
	return oncandidatewindowhide
}

func (*MSInputMethodContext) SetOncandidatewindowhide(oncandidatewindowhide func(event.Event)) {
	macro.Rewrite("$_.oncandidatewindowhide = $1", oncandidatewindowhide)
}

func (*MSInputMethodContext) Oncandidatewindowshow() (oncandidatewindowshow func(event.Event)) {
	macro.Rewrite("$_.oncandidatewindowshow")
	return oncandidatewindowshow
}

func (*MSInputMethodContext) SetOncandidatewindowshow(oncandidatewindowshow func(event.Event)) {
	macro.Rewrite("$_.oncandidatewindowshow = $1", oncandidatewindowshow)
}

func (*MSInputMethodContext) Oncandidatewindowupdate() (oncandidatewindowupdate func(event.Event)) {
	macro.Rewrite("$_.oncandidatewindowupdate")
	return oncandidatewindowupdate
}

func (*MSInputMethodContext) SetOncandidatewindowupdate(oncandidatewindowupdate func(event.Event)) {
	macro.Rewrite("$_.oncandidatewindowupdate = $1", oncandidatewindowupdate)
}

func (*MSInputMethodContext) Target() (target html.HTMLElement) {
	macro.Rewrite("$_.target")
	return target
}
