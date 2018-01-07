package window

import "github.com/matthewmueller/joy/dom/event"

type ServiceWorkerMessageEventInit struct {
	*event.EventInit

	data        *interface{}
	lastEventId *string
	origin      *string
	ports       *[]*MessagePort
	source      *interface{}
}
