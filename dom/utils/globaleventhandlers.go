package utils

import "github.com/matthewmueller/joy/dom/event"

type GlobalEventHandlers interface {
	Onpointercancel() (onpointercancel func(event.Event))

	SetOnpointercancel(onpointercancel func(event.Event))

	Onpointerdown() (onpointerdown func(event.Event))

	SetOnpointerdown(onpointerdown func(event.Event))

	Onpointerenter() (onpointerenter func(event.Event))

	SetOnpointerenter(onpointerenter func(event.Event))

	Onpointerleave() (onpointerleave func(event.Event))

	SetOnpointerleave(onpointerleave func(event.Event))

	Onpointermove() (onpointermove func(event.Event))

	SetOnpointermove(onpointermove func(event.Event))

	Onpointerout() (onpointerout func(event.Event))

	SetOnpointerout(onpointerout func(event.Event))

	Onpointerover() (onpointerover func(event.Event))

	SetOnpointerover(onpointerover func(event.Event))

	Onpointerup() (onpointerup func(event.Event))

	SetOnpointerup(onpointerup func(event.Event))

	Onwheel() (onwheel func(event.Event))

	SetOnwheel(onwheel func(event.Event))
}
