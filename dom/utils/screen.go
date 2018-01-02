package utils

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
)

var _ event.EventTarget = (*Screen)(nil)

type Screen struct {
}

func (*Screen) MsLockOrientation(orientations interface{}) (b bool) {
	macro.Rewrite("$_.msLockOrientation($1)", orientations)
	return b
}

func (*Screen) MsUnlockOrientation() {
	macro.Rewrite("$_.msUnlockOrientation()")
}

func (*Screen) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*Screen) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*Screen) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*Screen) AvailHeight() (availHeight uint) {
	macro.Rewrite("$_.availHeight")
	return availHeight
}

func (*Screen) AvailWidth() (availWidth uint) {
	macro.Rewrite("$_.availWidth")
	return availWidth
}

func (*Screen) BufferDepth() (bufferDepth int) {
	macro.Rewrite("$_.bufferDepth")
	return bufferDepth
}

func (*Screen) SetBufferDepth(bufferDepth int) {
	macro.Rewrite("$_.bufferDepth = $1", bufferDepth)
}

func (*Screen) ColorDepth() (colorDepth uint) {
	macro.Rewrite("$_.colorDepth")
	return colorDepth
}

func (*Screen) DeviceXDPI() (deviceXDPI int) {
	macro.Rewrite("$_.deviceXDPI")
	return deviceXDPI
}

func (*Screen) DeviceYDPI() (deviceYDPI int) {
	macro.Rewrite("$_.deviceYDPI")
	return deviceYDPI
}

func (*Screen) FontSmoothingEnabled() (fontSmoothingEnabled bool) {
	macro.Rewrite("$_.fontSmoothingEnabled")
	return fontSmoothingEnabled
}

func (*Screen) Height() (height uint) {
	macro.Rewrite("$_.height")
	return height
}

func (*Screen) LogicalXDPI() (logicalXDPI int) {
	macro.Rewrite("$_.logicalXDPI")
	return logicalXDPI
}

func (*Screen) LogicalYDPI() (logicalYDPI int) {
	macro.Rewrite("$_.logicalYDPI")
	return logicalYDPI
}

func (*Screen) MsOrientation() (msOrientation string) {
	macro.Rewrite("$_.msOrientation")
	return msOrientation
}

func (*Screen) Onmsorientationchange() (onmsorientationchange func(event.Event)) {
	macro.Rewrite("$_.onmsorientationchange")
	return onmsorientationchange
}

func (*Screen) SetOnmsorientationchange(onmsorientationchange func(event.Event)) {
	macro.Rewrite("$_.onmsorientationchange = $1", onmsorientationchange)
}

func (*Screen) PixelDepth() (pixelDepth uint) {
	macro.Rewrite("$_.pixelDepth")
	return pixelDepth
}

func (*Screen) SystemXDPI() (systemXDPI int) {
	macro.Rewrite("$_.systemXDPI")
	return systemXDPI
}

func (*Screen) SystemYDPI() (systemYDPI int) {
	macro.Rewrite("$_.systemYDPI")
	return systemYDPI
}

func (*Screen) Width() (width uint) {
	macro.Rewrite("$_.width")
	return width
}
