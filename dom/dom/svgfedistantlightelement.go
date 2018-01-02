package dom

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/element"
	"github.com/matthewmueller/joy/dom/utils"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/ms"
	"github.com/matthewmueller/joy/dom/touch"
	"github.com/matthewmueller/joy/dom/document"
	"github.com/matthewmueller/joy/dom/html"
)

var _ SVGElement = (*SVGFEDistantLightElement)(nil)
var _ element.Element = (*SVGFEDistantLightElement)(nil)
var _ utils.GlobalEventHandlers = (*SVGFEDistantLightElement)(nil)
var _ element.ElementTraversal = (*SVGFEDistantLightElement)(nil)
var _ NodeSelector = (*SVGFEDistantLightElement)(nil)
var _ ChildNode = (*SVGFEDistantLightElement)(nil)
var _ Node = (*SVGFEDistantLightElement)(nil)
var _ event.EventTarget = (*SVGFEDistantLightElement)(nil)

type SVGFEDistantLightElement struct {
}

func (*SVGFEDistantLightElement) GetAttribute(qualifiedName string) (s string) {
	macro.Rewrite("$_.getAttribute($1)", qualifiedName)
	return s
}

func (*SVGFEDistantLightElement) GetAttributeNode(name string) (w *Attr) {
	macro.Rewrite("$_.getAttributeNode($1)", name)
	return w
}

func (*SVGFEDistantLightElement) GetAttributeNodeNS(namespaceURI string, localName string) (w *Attr) {
	macro.Rewrite("$_.getAttributeNodeNS($1, $2)", namespaceURI, localName)
	return w
}

func (*SVGFEDistantLightElement) GetAttributeNS(namespaceURI string, localName string) (s string) {
	macro.Rewrite("$_.getAttributeNS($1, $2)", namespaceURI, localName)
	return s
}

func (*SVGFEDistantLightElement) GetBoundingClientRect() (c *ClientRect) {
	macro.Rewrite("$_.getBoundingClientRect()")
	return c
}

func (*SVGFEDistantLightElement) GetClientRects() (c *ClientRectList) {
	macro.Rewrite("$_.getClientRects()")
	return c
}

func (*SVGFEDistantLightElement) GetElementsByTagName(name string) (w *NodeList) {
	macro.Rewrite("$_.getElementsByTagName($1)", name)
	return w
}

func (*SVGFEDistantLightElement) GetElementsByTagNameNS(namespaceURI string, localName string) (w *NodeList) {
	macro.Rewrite("$_.getElementsByTagNameNS($1, $2)", namespaceURI, localName)
	return w
}

func (*SVGFEDistantLightElement) HasAttribute(name string) (b bool) {
	macro.Rewrite("$_.hasAttribute($1)", name)
	return b
}

func (*SVGFEDistantLightElement) HasAttributeNS(namespaceURI string, localName string) (b bool) {
	macro.Rewrite("$_.hasAttributeNS($1, $2)", namespaceURI, localName)
	return b
}

func (*SVGFEDistantLightElement) MsGetRegionContent() (w *ms.MSRangeCollection) {
	macro.Rewrite("$_.msGetRegionContent()")
	return w
}

func (*SVGFEDistantLightElement) MsGetUntransformedBounds() (c *ClientRect) {
	macro.Rewrite("$_.msGetUntransformedBounds()")
	return c
}

func (*SVGFEDistantLightElement) MsMatchesSelector(selectors string) (b bool) {
	macro.Rewrite("$_.msMatchesSelector($1)", selectors)
	return b
}

func (*SVGFEDistantLightElement) MsReleasePointerCapture(pointerId int) {
	macro.Rewrite("$_.msReleasePointerCapture($1)", pointerId)
}

func (*SVGFEDistantLightElement) MsSetPointerCapture(pointerId int) {
	macro.Rewrite("$_.msSetPointerCapture($1)", pointerId)
}

func (*SVGFEDistantLightElement) MsZoomTo(args *ms.MsZoomToOptions) {
	macro.Rewrite("$_.msZoomTo($1)", args)
}

func (*SVGFEDistantLightElement) ReleasePointerCapture(pointerId int) {
	macro.Rewrite("$_.releasePointerCapture($1)", pointerId)
}

func (*SVGFEDistantLightElement) RemoveAttribute(qualifiedName string) {
	macro.Rewrite("$_.removeAttribute($1)", qualifiedName)
}

func (*SVGFEDistantLightElement) RemoveAttributeNode(oldAttr *Attr) (w *Attr) {
	macro.Rewrite("$_.removeAttributeNode($1)", oldAttr)
	return w
}

func (*SVGFEDistantLightElement) RemoveAttributeNS(namespaceURI string, localName string) {
	macro.Rewrite("$_.removeAttributeNS($1, $2)", namespaceURI, localName)
}

func (*SVGFEDistantLightElement) RequestFullscreen() {
	macro.Rewrite("$_.requestFullscreen()")
}

func (*SVGFEDistantLightElement) RequestPointerLock() {
	macro.Rewrite("$_.requestPointerLock()")
}

func (*SVGFEDistantLightElement) SetAttribute(qualifiedName string, value string) {
	macro.Rewrite("$_.setAttribute($1, $2)", qualifiedName, value)
}

func (*SVGFEDistantLightElement) SetAttributeNode(newAttr *Attr) (w *Attr) {
	macro.Rewrite("$_.setAttributeNode($1)", newAttr)
	return w
}

func (*SVGFEDistantLightElement) SetAttributeNodeNS(newAttr *Attr) (w *Attr) {
	macro.Rewrite("$_.setAttributeNodeNS($1)", newAttr)
	return w
}

func (*SVGFEDistantLightElement) SetAttributeNS(namespaceURI string, qualifiedName string, value string) {
	macro.Rewrite("$_.setAttributeNS($1, $2, $3)", namespaceURI, qualifiedName, value)
}

func (*SVGFEDistantLightElement) SetPointerCapture(pointerId int) {
	macro.Rewrite("$_.setPointerCapture($1)", pointerId)
}

func (*SVGFEDistantLightElement) WebkitMatchesSelector(selectors string) (b bool) {
	macro.Rewrite("$_.webkitMatchesSelector($1)", selectors)
	return b
}

func (*SVGFEDistantLightElement) WebkitRequestFullscreen() {
	macro.Rewrite("$_.webkitRequestFullscreen()")
}

func (*SVGFEDistantLightElement) WebkitRequestFullScreen() {
	macro.Rewrite("$_.webkitRequestFullScreen()")
}

func (*SVGFEDistantLightElement) QuerySelector(selectors string) (w element.Element) {
	macro.Rewrite("$_.querySelector($1)", selectors)
	return w
}

func (*SVGFEDistantLightElement) QuerySelectorAll(selectors string) (w *NodeList) {
	macro.Rewrite("$_.querySelectorAll($1)", selectors)
	return w
}

func (*SVGFEDistantLightElement) AppendChild(newChild Node) (w Node) {
	macro.Rewrite("$_.appendChild($1)", newChild)
	return w
}

func (*SVGFEDistantLightElement) CloneNode(deep *bool) (w Node) {
	macro.Rewrite("$_.cloneNode($1)", deep)
	return w
}

func (*SVGFEDistantLightElement) CompareDocumentPosition(other Node) (u uint8) {
	macro.Rewrite("$_.compareDocumentPosition($1)", other)
	return u
}

func (*SVGFEDistantLightElement) Contains(child Node) (b bool) {
	macro.Rewrite("$_.contains($1)", child)
	return b
}

func (*SVGFEDistantLightElement) HasAttributes() (b bool) {
	macro.Rewrite("$_.hasAttributes()")
	return b
}

func (*SVGFEDistantLightElement) HasChildNodes() (b bool) {
	macro.Rewrite("$_.hasChildNodes()")
	return b
}

func (*SVGFEDistantLightElement) InsertBefore(newChild Node, refChild *Node) (w Node) {
	macro.Rewrite("$_.insertBefore($1, $2)", newChild, refChild)
	return w
}

func (*SVGFEDistantLightElement) IsDefaultNamespace(namespaceURI string) (b bool) {
	macro.Rewrite("$_.isDefaultNamespace($1)", namespaceURI)
	return b
}

func (*SVGFEDistantLightElement) IsEqualNode(arg Node) (b bool) {
	macro.Rewrite("$_.isEqualNode($1)", arg)
	return b
}

func (*SVGFEDistantLightElement) IsSameNode(other Node) (b bool) {
	macro.Rewrite("$_.isSameNode($1)", other)
	return b
}

func (*SVGFEDistantLightElement) LookupNamespaceURI(prefix string) (s string) {
	macro.Rewrite("$_.lookupNamespaceURI($1)", prefix)
	return s
}

func (*SVGFEDistantLightElement) LookupPrefix(namespaceURI string) (s string) {
	macro.Rewrite("$_.lookupPrefix($1)", namespaceURI)
	return s
}

func (*SVGFEDistantLightElement) Normalize() {
	macro.Rewrite("$_.normalize()")
}

func (*SVGFEDistantLightElement) RemoveChild(oldChild Node) (w Node) {
	macro.Rewrite("$_.removeChild($1)", oldChild)
	return w
}

func (*SVGFEDistantLightElement) ReplaceChild(newChild Node, oldChild Node) (w Node) {
	macro.Rewrite("$_.replaceChild($1, $2)", newChild, oldChild)
	return w
}

func (*SVGFEDistantLightElement) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*SVGFEDistantLightElement) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*SVGFEDistantLightElement) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*SVGFEDistantLightElement) Azimuth() (azimuth *SVGAnimatedNumber) {
	macro.Rewrite("$_.azimuth")
	return azimuth
}

func (*SVGFEDistantLightElement) Elevation() (elevation *SVGAnimatedNumber) {
	macro.Rewrite("$_.elevation")
	return elevation
}

func (*SVGFEDistantLightElement) Dataset() (dataset *DOMStringMap) {
	macro.Rewrite("$_.dataset")
	return dataset
}

func (*SVGFEDistantLightElement) OwnerSVGElement() (ownerSVGElement *SVGSVGElement) {
	macro.Rewrite("$_.ownerSVGElement")
	return ownerSVGElement
}

func (*SVGFEDistantLightElement) ViewportElement() (viewportElement SVGElement) {
	macro.Rewrite("$_.viewportElement")
	return viewportElement
}

func (*SVGFEDistantLightElement) Xmlbase() (xmlbase string) {
	macro.Rewrite("$_.xmlbase")
	return xmlbase
}

func (*SVGFEDistantLightElement) SetXmlbase(xmlbase string) {
	macro.Rewrite("$_.xmlbase = $1", xmlbase)
}

func (*SVGFEDistantLightElement) ClassList() (classList DOMTokenList) {
	macro.Rewrite("$_.classList")
	return classList
}

func (*SVGFEDistantLightElement) ClassName() (className string) {
	macro.Rewrite("$_.className")
	return className
}

func (*SVGFEDistantLightElement) SetClassName(className string) {
	macro.Rewrite("$_.className = $1", className)
}

func (*SVGFEDistantLightElement) ClientHeight() (clientHeight int) {
	macro.Rewrite("$_.clientHeight")
	return clientHeight
}

func (*SVGFEDistantLightElement) ClientLeft() (clientLeft int) {
	macro.Rewrite("$_.clientLeft")
	return clientLeft
}

func (*SVGFEDistantLightElement) ClientTop() (clientTop int) {
	macro.Rewrite("$_.clientTop")
	return clientTop
}

func (*SVGFEDistantLightElement) ClientWidth() (clientWidth int) {
	macro.Rewrite("$_.clientWidth")
	return clientWidth
}

func (*SVGFEDistantLightElement) ID() (id string) {
	macro.Rewrite("$_.id")
	return id
}

func (*SVGFEDistantLightElement) SetID(id string) {
	macro.Rewrite("$_.id = $1", id)
}

func (*SVGFEDistantLightElement) InnerHTML() (innerHTML string) {
	macro.Rewrite("$_.innerHTML")
	return innerHTML
}

func (*SVGFEDistantLightElement) SetInnerHTML(innerHTML string) {
	macro.Rewrite("$_.innerHTML = $1", innerHTML)
}

func (*SVGFEDistantLightElement) MsContentZoomFactor() (msContentZoomFactor float32) {
	macro.Rewrite("$_.msContentZoomFactor")
	return msContentZoomFactor
}

func (*SVGFEDistantLightElement) SetMsContentZoomFactor(msContentZoomFactor float32) {
	macro.Rewrite("$_.msContentZoomFactor = $1", msContentZoomFactor)
}

func (*SVGFEDistantLightElement) MsRegionOverflow() (msRegionOverflow string) {
	macro.Rewrite("$_.msRegionOverflow")
	return msRegionOverflow
}

func (*SVGFEDistantLightElement) Onariarequest() (onariarequest func(event.Event)) {
	macro.Rewrite("$_.onariarequest")
	return onariarequest
}

func (*SVGFEDistantLightElement) SetOnariarequest(onariarequest func(event.Event)) {
	macro.Rewrite("$_.onariarequest = $1", onariarequest)
}

func (*SVGFEDistantLightElement) Oncommand() (oncommand func(event.Event)) {
	macro.Rewrite("$_.oncommand")
	return oncommand
}

func (*SVGFEDistantLightElement) SetOncommand(oncommand func(event.Event)) {
	macro.Rewrite("$_.oncommand = $1", oncommand)
}

func (*SVGFEDistantLightElement) Ongotpointercapture() (ongotpointercapture func(*utils.PointerEvent)) {
	macro.Rewrite("$_.ongotpointercapture")
	return ongotpointercapture
}

func (*SVGFEDistantLightElement) SetOngotpointercapture(ongotpointercapture func(*utils.PointerEvent)) {
	macro.Rewrite("$_.ongotpointercapture = $1", ongotpointercapture)
}

func (*SVGFEDistantLightElement) Onlostpointercapture() (onlostpointercapture func(*utils.PointerEvent)) {
	macro.Rewrite("$_.onlostpointercapture")
	return onlostpointercapture
}

func (*SVGFEDistantLightElement) SetOnlostpointercapture(onlostpointercapture func(*utils.PointerEvent)) {
	macro.Rewrite("$_.onlostpointercapture = $1", onlostpointercapture)
}

func (*SVGFEDistantLightElement) Onmsgesturechange() (onmsgesturechange func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturechange")
	return onmsgesturechange
}

func (*SVGFEDistantLightElement) SetOnmsgesturechange(onmsgesturechange func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturechange = $1", onmsgesturechange)
}

func (*SVGFEDistantLightElement) Onmsgesturedoubletap() (onmsgesturedoubletap func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturedoubletap")
	return onmsgesturedoubletap
}

func (*SVGFEDistantLightElement) SetOnmsgesturedoubletap(onmsgesturedoubletap func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturedoubletap = $1", onmsgesturedoubletap)
}

func (*SVGFEDistantLightElement) Onmsgestureend() (onmsgestureend func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgestureend")
	return onmsgestureend
}

func (*SVGFEDistantLightElement) SetOnmsgestureend(onmsgestureend func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgestureend = $1", onmsgestureend)
}

func (*SVGFEDistantLightElement) Onmsgesturehold() (onmsgesturehold func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturehold")
	return onmsgesturehold
}

func (*SVGFEDistantLightElement) SetOnmsgesturehold(onmsgesturehold func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturehold = $1", onmsgesturehold)
}

func (*SVGFEDistantLightElement) Onmsgesturestart() (onmsgesturestart func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturestart")
	return onmsgesturestart
}

func (*SVGFEDistantLightElement) SetOnmsgesturestart(onmsgesturestart func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturestart = $1", onmsgesturestart)
}

func (*SVGFEDistantLightElement) Onmsgesturetap() (onmsgesturetap func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturetap")
	return onmsgesturetap
}

func (*SVGFEDistantLightElement) SetOnmsgesturetap(onmsgesturetap func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturetap = $1", onmsgesturetap)
}

func (*SVGFEDistantLightElement) Onmsgotpointercapture() (onmsgotpointercapture func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmsgotpointercapture")
	return onmsgotpointercapture
}

func (*SVGFEDistantLightElement) SetOnmsgotpointercapture(onmsgotpointercapture func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmsgotpointercapture = $1", onmsgotpointercapture)
}

func (*SVGFEDistantLightElement) Onmsinertiastart() (onmsinertiastart func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsinertiastart")
	return onmsinertiastart
}

func (*SVGFEDistantLightElement) SetOnmsinertiastart(onmsinertiastart func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsinertiastart = $1", onmsinertiastart)
}

func (*SVGFEDistantLightElement) Onmslostpointercapture() (onmslostpointercapture func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmslostpointercapture")
	return onmslostpointercapture
}

func (*SVGFEDistantLightElement) SetOnmslostpointercapture(onmslostpointercapture func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmslostpointercapture = $1", onmslostpointercapture)
}

func (*SVGFEDistantLightElement) Onmspointercancel() (onmspointercancel func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointercancel")
	return onmspointercancel
}

func (*SVGFEDistantLightElement) SetOnmspointercancel(onmspointercancel func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointercancel = $1", onmspointercancel)
}

func (*SVGFEDistantLightElement) Onmspointerdown() (onmspointerdown func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerdown")
	return onmspointerdown
}

func (*SVGFEDistantLightElement) SetOnmspointerdown(onmspointerdown func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerdown = $1", onmspointerdown)
}

func (*SVGFEDistantLightElement) Onmspointerenter() (onmspointerenter func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerenter")
	return onmspointerenter
}

func (*SVGFEDistantLightElement) SetOnmspointerenter(onmspointerenter func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerenter = $1", onmspointerenter)
}

func (*SVGFEDistantLightElement) Onmspointerleave() (onmspointerleave func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerleave")
	return onmspointerleave
}

func (*SVGFEDistantLightElement) SetOnmspointerleave(onmspointerleave func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerleave = $1", onmspointerleave)
}

func (*SVGFEDistantLightElement) Onmspointermove() (onmspointermove func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointermove")
	return onmspointermove
}

func (*SVGFEDistantLightElement) SetOnmspointermove(onmspointermove func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointermove = $1", onmspointermove)
}

func (*SVGFEDistantLightElement) Onmspointerout() (onmspointerout func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerout")
	return onmspointerout
}

func (*SVGFEDistantLightElement) SetOnmspointerout(onmspointerout func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerout = $1", onmspointerout)
}

func (*SVGFEDistantLightElement) Onmspointerover() (onmspointerover func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerover")
	return onmspointerover
}

func (*SVGFEDistantLightElement) SetOnmspointerover(onmspointerover func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerover = $1", onmspointerover)
}

func (*SVGFEDistantLightElement) Onmspointerup() (onmspointerup func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerup")
	return onmspointerup
}

func (*SVGFEDistantLightElement) SetOnmspointerup(onmspointerup func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerup = $1", onmspointerup)
}

func (*SVGFEDistantLightElement) Ontouchcancel() (ontouchcancel func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchcancel")
	return ontouchcancel
}

func (*SVGFEDistantLightElement) SetOntouchcancel(ontouchcancel func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchcancel = $1", ontouchcancel)
}

func (*SVGFEDistantLightElement) Ontouchend() (ontouchend func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchend")
	return ontouchend
}

func (*SVGFEDistantLightElement) SetOntouchend(ontouchend func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchend = $1", ontouchend)
}

func (*SVGFEDistantLightElement) Ontouchmove() (ontouchmove func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchmove")
	return ontouchmove
}

func (*SVGFEDistantLightElement) SetOntouchmove(ontouchmove func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchmove = $1", ontouchmove)
}

func (*SVGFEDistantLightElement) Ontouchstart() (ontouchstart func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchstart")
	return ontouchstart
}

func (*SVGFEDistantLightElement) SetOntouchstart(ontouchstart func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchstart = $1", ontouchstart)
}

func (*SVGFEDistantLightElement) Onwebkitfullscreenchange() (onwebkitfullscreenchange func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenchange")
	return onwebkitfullscreenchange
}

func (*SVGFEDistantLightElement) SetOnwebkitfullscreenchange(onwebkitfullscreenchange func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenchange = $1", onwebkitfullscreenchange)
}

func (*SVGFEDistantLightElement) Onwebkitfullscreenerror() (onwebkitfullscreenerror func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenerror")
	return onwebkitfullscreenerror
}

func (*SVGFEDistantLightElement) SetOnwebkitfullscreenerror(onwebkitfullscreenerror func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenerror = $1", onwebkitfullscreenerror)
}

func (*SVGFEDistantLightElement) OuterHTML() (outerHTML string) {
	macro.Rewrite("$_.outerHTML")
	return outerHTML
}

func (*SVGFEDistantLightElement) SetOuterHTML(outerHTML string) {
	macro.Rewrite("$_.outerHTML = $1", outerHTML)
}

func (*SVGFEDistantLightElement) Prefix() (prefix string) {
	macro.Rewrite("$_.prefix")
	return prefix
}

func (*SVGFEDistantLightElement) ScrollHeight() (scrollHeight int) {
	macro.Rewrite("$_.scrollHeight")
	return scrollHeight
}

func (*SVGFEDistantLightElement) ScrollLeft() (scrollLeft int) {
	macro.Rewrite("$_.scrollLeft")
	return scrollLeft
}

func (*SVGFEDistantLightElement) SetScrollLeft(scrollLeft int) {
	macro.Rewrite("$_.scrollLeft = $1", scrollLeft)
}

func (*SVGFEDistantLightElement) ScrollTop() (scrollTop int) {
	macro.Rewrite("$_.scrollTop")
	return scrollTop
}

func (*SVGFEDistantLightElement) SetScrollTop(scrollTop int) {
	macro.Rewrite("$_.scrollTop = $1", scrollTop)
}

func (*SVGFEDistantLightElement) ScrollWidth() (scrollWidth int) {
	macro.Rewrite("$_.scrollWidth")
	return scrollWidth
}

func (*SVGFEDistantLightElement) TagName() (tagName string) {
	macro.Rewrite("$_.tagName")
	return tagName
}

func (*SVGFEDistantLightElement) Onpointercancel() (onpointercancel func(event.Event)) {
	macro.Rewrite("$_.onpointercancel")
	return onpointercancel
}

func (*SVGFEDistantLightElement) SetOnpointercancel(onpointercancel func(event.Event)) {
	macro.Rewrite("$_.onpointercancel = $1", onpointercancel)
}

func (*SVGFEDistantLightElement) Onpointerdown() (onpointerdown func(event.Event)) {
	macro.Rewrite("$_.onpointerdown")
	return onpointerdown
}

func (*SVGFEDistantLightElement) SetOnpointerdown(onpointerdown func(event.Event)) {
	macro.Rewrite("$_.onpointerdown = $1", onpointerdown)
}

func (*SVGFEDistantLightElement) Onpointerenter() (onpointerenter func(event.Event)) {
	macro.Rewrite("$_.onpointerenter")
	return onpointerenter
}

func (*SVGFEDistantLightElement) SetOnpointerenter(onpointerenter func(event.Event)) {
	macro.Rewrite("$_.onpointerenter = $1", onpointerenter)
}

func (*SVGFEDistantLightElement) Onpointerleave() (onpointerleave func(event.Event)) {
	macro.Rewrite("$_.onpointerleave")
	return onpointerleave
}

func (*SVGFEDistantLightElement) SetOnpointerleave(onpointerleave func(event.Event)) {
	macro.Rewrite("$_.onpointerleave = $1", onpointerleave)
}

func (*SVGFEDistantLightElement) Onpointermove() (onpointermove func(event.Event)) {
	macro.Rewrite("$_.onpointermove")
	return onpointermove
}

func (*SVGFEDistantLightElement) SetOnpointermove(onpointermove func(event.Event)) {
	macro.Rewrite("$_.onpointermove = $1", onpointermove)
}

func (*SVGFEDistantLightElement) Onpointerout() (onpointerout func(event.Event)) {
	macro.Rewrite("$_.onpointerout")
	return onpointerout
}

func (*SVGFEDistantLightElement) SetOnpointerout(onpointerout func(event.Event)) {
	macro.Rewrite("$_.onpointerout = $1", onpointerout)
}

func (*SVGFEDistantLightElement) Onpointerover() (onpointerover func(event.Event)) {
	macro.Rewrite("$_.onpointerover")
	return onpointerover
}

func (*SVGFEDistantLightElement) SetOnpointerover(onpointerover func(event.Event)) {
	macro.Rewrite("$_.onpointerover = $1", onpointerover)
}

func (*SVGFEDistantLightElement) Onpointerup() (onpointerup func(event.Event)) {
	macro.Rewrite("$_.onpointerup")
	return onpointerup
}

func (*SVGFEDistantLightElement) SetOnpointerup(onpointerup func(event.Event)) {
	macro.Rewrite("$_.onpointerup = $1", onpointerup)
}

func (*SVGFEDistantLightElement) Onwheel() (onwheel func(event.Event)) {
	macro.Rewrite("$_.onwheel")
	return onwheel
}

func (*SVGFEDistantLightElement) SetOnwheel(onwheel func(event.Event)) {
	macro.Rewrite("$_.onwheel = $1", onwheel)
}

func (*SVGFEDistantLightElement) ChildElementCount() (childElementCount uint) {
	macro.Rewrite("$_.childElementCount")
	return childElementCount
}

func (*SVGFEDistantLightElement) FirstElementChild() (firstElementChild element.Element) {
	macro.Rewrite("$_.firstElementChild")
	return firstElementChild
}

func (*SVGFEDistantLightElement) LastElementChild() (lastElementChild element.Element) {
	macro.Rewrite("$_.lastElementChild")
	return lastElementChild
}

func (*SVGFEDistantLightElement) NextElementSibling() (nextElementSibling element.Element) {
	macro.Rewrite("$_.nextElementSibling")
	return nextElementSibling
}

func (*SVGFEDistantLightElement) PreviousElementSibling() (previousElementSibling element.Element) {
	macro.Rewrite("$_.previousElementSibling")
	return previousElementSibling
}

func (*SVGFEDistantLightElement) Attributes() (attributes *NamedNodeMap) {
	macro.Rewrite("$_.attributes")
	return attributes
}

func (*SVGFEDistantLightElement) BaseURI() (baseURI string) {
	macro.Rewrite("$_.baseURI")
	return baseURI
}

func (*SVGFEDistantLightElement) ChildNodes() (childNodes *NodeList) {
	macro.Rewrite("$_.childNodes")
	return childNodes
}

func (*SVGFEDistantLightElement) FirstChild() (firstChild Node) {
	macro.Rewrite("$_.firstChild")
	return firstChild
}

func (*SVGFEDistantLightElement) LastChild() (lastChild Node) {
	macro.Rewrite("$_.lastChild")
	return lastChild
}

func (*SVGFEDistantLightElement) LocalName() (localName string) {
	macro.Rewrite("$_.localName")
	return localName
}

func (*SVGFEDistantLightElement) NamespaceURI() (namespaceURI string) {
	macro.Rewrite("$_.namespaceURI")
	return namespaceURI
}

func (*SVGFEDistantLightElement) NextSibling() (nextSibling Node) {
	macro.Rewrite("$_.nextSibling")
	return nextSibling
}

func (*SVGFEDistantLightElement) NodeName() (nodeName string) {
	macro.Rewrite("$_.nodeName")
	return nodeName
}

func (*SVGFEDistantLightElement) NodeType() (nodeType uint8) {
	macro.Rewrite("$_.nodeType")
	return nodeType
}

func (*SVGFEDistantLightElement) NodeValue() (nodeValue string) {
	macro.Rewrite("$_.nodeValue")
	return nodeValue
}

func (*SVGFEDistantLightElement) SetNodeValue(nodeValue string) {
	macro.Rewrite("$_.nodeValue = $1", nodeValue)
}

func (*SVGFEDistantLightElement) OwnerDocument() (ownerDocument document.Document) {
	macro.Rewrite("$_.ownerDocument")
	return ownerDocument
}

func (*SVGFEDistantLightElement) ParentElement() (parentElement html.HTMLElement) {
	macro.Rewrite("$_.parentElement")
	return parentElement
}

func (*SVGFEDistantLightElement) ParentNode() (parentNode Node) {
	macro.Rewrite("$_.parentNode")
	return parentNode
}

func (*SVGFEDistantLightElement) PreviousSibling() (previousSibling Node) {
	macro.Rewrite("$_.previousSibling")
	return previousSibling
}

func (*SVGFEDistantLightElement) TextContent() (textContent string) {
	macro.Rewrite("$_.textContent")
	return textContent
}

func (*SVGFEDistantLightElement) SetTextContent(textContent string) {
	macro.Rewrite("$_.textContent = $1", textContent)
}
