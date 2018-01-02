package indexdb

import (
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/dom"
)

type IDBRequest interface {
	event.
		EventTarget

	Error() (err *dom.DOMError)

	Onerror() (onerror func(event.Event))

	SetOnerror(onerror func(event.Event))

	Onsuccess() (onsuccess func(event.Event))

	SetOnsuccess(onsuccess func(event.Event))

	ReadyState() (readyState *IDBRequestReadyState)

	Result() (result interface{})

	Source() (source interface{})

	Transaction() (transaction *IDBTransaction)
}
