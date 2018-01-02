package element

import (
	"github.com/matthewmueller/joy/dom/dom"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/touch"
	"github.com/matthewmueller/joy/dom/ms"
	"github.com/matthewmueller/joy/dom/utils"
)

type Element interface {
	dom.
		Node

	QuerySelector(selectors string) (e Element)

	QuerySelectorAll(selectors string) (n *dom.NodeList)

	GetAttribute(qualifiedName string) (s string)

	GetAttributeNode(name string) (a *dom.Attr)

	GetAttributeNodeNS(namespaceURI string, localName string) (a *dom.Attr)

	GetAttributeNS(namespaceURI string, localName string) (s string)

	GetBoundingClientRect() (c *dom.ClientRect)

	GetClientRects() (c *dom.ClientRectList)

	GetElementsByTagName(name string) (n *dom.NodeList)

	GetElementsByTagNameNS(namespaceURI string, localName string) (n *dom.NodeList)

	HasAttribute(name string) (b bool)

	HasAttributeNS(namespaceURI string, localName string) (b bool)

	MsGetRegionContent() (m *ms.MSRangeCollection)

	MsGetUntransformedBounds() (c *dom.ClientRect)

	MsMatchesSelector(selectors string) (b bool)

	MsReleasePointerCapture(pointerId int)

	MsSetPointerCapture(pointerId int)

	MsZoomTo(args *ms.MsZoomToOptions)

	ReleasePointerCapture(pointerId int)

	RemoveAttribute(qualifiedName string)

	RemoveAttributeNode(oldAttr *dom.Attr) (a *dom.Attr)

	RemoveAttributeNS(namespaceURI string, localName string)

	RequestFullscreen()

	RequestPointerLock()

	SetAttribute(qualifiedName string, value string)

	SetAttributeNode(newAttr *dom.Attr) (a *dom.Attr)

	SetAttributeNodeNS(newAttr *dom.Attr) (a *dom.Attr)

	SetAttributeNS(namespaceURI string, qualifiedName string, value string)

	SetPointerCapture(pointerId int)

	WebkitMatchesSelector(selectors string) (b bool)

	WebkitRequestFullscreen()

	WebkitRequestFullScreen()

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

	ChildElementCount() (childElementCount uint)

	FirstElementChild() (firstElementChild Element)

	LastElementChild() (lastElementChild Element)

	NextElementSibling() (nextElementSibling Element)

	PreviousElementSibling() (previousElementSibling Element)

	ClassList() (classList dom.DOMTokenList)

	ClassName() (className string)

	SetClassName(className string)

	ClientHeight() (clientHeight int)

	ClientLeft() (clientLeft int)

	ClientTop() (clientTop int)

	ClientWidth() (clientWidth int)

	ID() (id string)

	SetID(id string)

	InnerHTML() (innerHTML string)

	SetInnerHTML(innerHTML string)

	MsContentZoomFactor() (msContentZoomFactor float32)

	SetMsContentZoomFactor(msContentZoomFactor float32)

	MsRegionOverflow() (msRegionOverflow string)

	Onariarequest() (onariarequest func(event.Event))

	SetOnariarequest(onariarequest func(event.Event))

	Oncommand() (oncommand func(event.Event))

	SetOncommand(oncommand func(event.Event))

	Ongotpointercapture() (ongotpointercapture func(*utils.PointerEvent))

	SetOngotpointercapture(ongotpointercapture func(*utils.PointerEvent))

	Onlostpointercapture() (onlostpointercapture func(*utils.PointerEvent))

	SetOnlostpointercapture(onlostpointercapture func(*utils.PointerEvent))

	Onmsgesturechange() (onmsgesturechange func(*ms.MSGestureEvent))

	SetOnmsgesturechange(onmsgesturechange func(*ms.MSGestureEvent))

	Onmsgesturedoubletap() (onmsgesturedoubletap func(*ms.MSGestureEvent))

	SetOnmsgesturedoubletap(onmsgesturedoubletap func(*ms.MSGestureEvent))

	Onmsgestureend() (onmsgestureend func(*ms.MSGestureEvent))

	SetOnmsgestureend(onmsgestureend func(*ms.MSGestureEvent))

	Onmsgesturehold() (onmsgesturehold func(*ms.MSGestureEvent))

	SetOnmsgesturehold(onmsgesturehold func(*ms.MSGestureEvent))

	Onmsgesturestart() (onmsgesturestart func(*ms.MSGestureEvent))

	SetOnmsgesturestart(onmsgesturestart func(*ms.MSGestureEvent))

	Onmsgesturetap() (onmsgesturetap func(*ms.MSGestureEvent))

	SetOnmsgesturetap(onmsgesturetap func(*ms.MSGestureEvent))

	Onmsgotpointercapture() (onmsgotpointercapture func(*ms.MSPointerEvent))

	SetOnmsgotpointercapture(onmsgotpointercapture func(*ms.MSPointerEvent))

	Onmsinertiastart() (onmsinertiastart func(*ms.MSGestureEvent))

	SetOnmsinertiastart(onmsinertiastart func(*ms.MSGestureEvent))

	Onmslostpointercapture() (onmslostpointercapture func(*ms.MSPointerEvent))

	SetOnmslostpointercapture(onmslostpointercapture func(*ms.MSPointerEvent))

	Onmspointercancel() (onmspointercancel func(*ms.MSPointerEvent))

	SetOnmspointercancel(onmspointercancel func(*ms.MSPointerEvent))

	Onmspointerdown() (onmspointerdown func(*ms.MSPointerEvent))

	SetOnmspointerdown(onmspointerdown func(*ms.MSPointerEvent))

	Onmspointerenter() (onmspointerenter func(*ms.MSPointerEvent))

	SetOnmspointerenter(onmspointerenter func(*ms.MSPointerEvent))

	Onmspointerleave() (onmspointerleave func(*ms.MSPointerEvent))

	SetOnmspointerleave(onmspointerleave func(*ms.MSPointerEvent))

	Onmspointermove() (onmspointermove func(*ms.MSPointerEvent))

	SetOnmspointermove(onmspointermove func(*ms.MSPointerEvent))

	Onmspointerout() (onmspointerout func(*ms.MSPointerEvent))

	SetOnmspointerout(onmspointerout func(*ms.MSPointerEvent))

	Onmspointerover() (onmspointerover func(*ms.MSPointerEvent))

	SetOnmspointerover(onmspointerover func(*ms.MSPointerEvent))

	Onmspointerup() (onmspointerup func(*ms.MSPointerEvent))

	SetOnmspointerup(onmspointerup func(*ms.MSPointerEvent))

	Ontouchcancel() (ontouchcancel func(*touch.TouchEvent))

	SetOntouchcancel(ontouchcancel func(*touch.TouchEvent))

	Ontouchend() (ontouchend func(*touch.TouchEvent))

	SetOntouchend(ontouchend func(*touch.TouchEvent))

	Ontouchmove() (ontouchmove func(*touch.TouchEvent))

	SetOntouchmove(ontouchmove func(*touch.TouchEvent))

	Ontouchstart() (ontouchstart func(*touch.TouchEvent))

	SetOntouchstart(ontouchstart func(*touch.TouchEvent))

	Onwebkitfullscreenchange() (onwebkitfullscreenchange func(event.Event))

	SetOnwebkitfullscreenchange(onwebkitfullscreenchange func(event.Event))

	Onwebkitfullscreenerror() (onwebkitfullscreenerror func(event.Event))

	SetOnwebkitfullscreenerror(onwebkitfullscreenerror func(event.Event))

	OuterHTML() (outerHTML string)

	SetOuterHTML(outerHTML string)

	Prefix() (prefix string)

	ScrollHeight() (scrollHeight int)

	ScrollLeft() (scrollLeft int)

	SetScrollLeft(scrollLeft int)

	ScrollTop() (scrollTop int)

	SetScrollTop(scrollTop int)

	ScrollWidth() (scrollWidth int)

	TagName() (tagName string)
}
