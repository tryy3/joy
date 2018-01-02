package serviceworker

import (
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/channelmessage"
)

type ServiceWorkerMessageEventInit struct {
	*event.EventInit

	data		*interface{}
	lastEventId	*string
	origin		*string
	ports		*[]*channelmessage.MessagePort
	source		*interface{}
}
