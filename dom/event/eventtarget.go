package event

type EventTarget interface {
	AddEventListener(kind string, listener func(evt Event), useCapture bool)

	DispatchEvent(evt Event) (b bool)

	RemoveEventListener(kind string, listener func(evt Event), useCapture bool)
}
