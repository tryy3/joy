package mouse

import (
	"github.com/matthewmueller/joy/dom/ui"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/dom/element"
)

type MouseEvent interface {
	ui.
		UIEvent

	GetModifierState(keyArg string) (b bool)

	InitMouseEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, detailArg int, screenXArg int, screenYArg int, clientXArg int, clientYArg int, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg uint8, relatedTargetArg event.EventTarget)

	AltKey() (altKey bool)

	Button() (button uint8)

	Buttons() (buttons uint8)

	ClientX() (clientX int)

	ClientY() (clientY int)

	CtrlKey() (ctrlKey bool)

	FromElement() (fromElement element.Element)

	LayerX() (layerX int)

	LayerY() (layerY int)

	MetaKey() (metaKey bool)

	MovementX() (movementX int)

	MovementY() (movementY int)

	OffsetX() (offsetX int)

	OffsetY() (offsetY int)

	PageX() (pageX int)

	PageY() (pageY int)

	RelatedTarget() (relatedTarget event.EventTarget)

	ScreenX() (screenX int)

	ScreenY() (screenY int)

	ShiftKey() (shiftKey bool)

	ToElement() (toElement element.Element)

	Which() (which uint8)

	X() (x int)

	Y() (y int)
}
