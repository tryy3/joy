package utils

import (
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/dom/channelmessage"
)

type MessageEventInit struct {
	*event.EventInit

	data	*interface{}
	origin	*string
	ports	*[]*channelmessage.MessagePort
	source	*window.Window
}
