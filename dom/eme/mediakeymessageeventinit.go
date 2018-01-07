package eme

import "github.com/matthewmueller/joy/dom/event"

type MediaKeyMessageEventInit struct {
	*event.EventInit

	message     *[]byte
	messageType *MediaKeyMessageType
}
