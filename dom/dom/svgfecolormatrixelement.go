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

var _ SVGElement = (*SVGFEColorMatrixElement)(nil)
var _ element.Element = (*SVGFEColorMatrixElement)(nil)
var _ utils.GlobalEventHandlers = (*SVGFEColorMatrixElement)(nil)
var _ element.ElementTraversal = (*SVGFEColorMatrixElement)(nil)
var _ NodeSelector = (*SVGFEColorMatrixElement)(nil)
var _ ChildNode = (*SVGFEColorMatrixElement)(nil)
var _ Node = (*SVGFEColorMatrixElement)(nil)
var _ event.EventTarget = (*SVGFEColorMatrixElement)(nil)

type SVGFEColorMatrixElement struct {
}

func (*SVGFEColorMatrixElement) GetAttribute(qualifiedName string) (s string) {
	macro.Rewrite("$_.getAttribute($1)", qualifiedName)
	return s
}

func (*SVGFEColorMatrixElement) GetAttributeNode(name string) (w *Attr) {
	macro.Rewrite("$_.getAttributeNode($1)", name)
	return w
}

func (*SVGFEColorMatrixElement) GetAttributeNodeNS(namespaceURI string, localName string) (w *Attr) {
	macro.Rewrite("$_.getAttributeNodeNS($1, $2)", namespaceURI, localName)
	return w
}

func (*SVGFEColorMatrixElement) GetAttributeNS(namespaceURI string, localName string) (s string) {
	macro.Rewrite("$_.getAttributeNS($1, $2)", namespaceURI, localName)
	return s
}

func (*SVGFEColorMatrixElement) GetBoundingClientRect() (c *ClientRect) {
	macro.Rewrite("$_.getBoundingClientRect()")
	return c
}

func (*SVGFEColorMatrixElement) GetClientRects() (c *ClientRectList) {
	macro.Rewrite("$_.getClientRects()")
	return c
}

func (*SVGFEColorMatrixElement) GetElementsByTagName(name string) (w *NodeList) {
	macro.Rewrite("$_.getElementsByTagName($1)", name)
	return w
}

func (*SVGFEColorMatrixElement) GetElementsByTagNameNS(namespaceURI string, localName string) (w *NodeList) {
	macro.Rewrite("$_.getElementsByTagNameNS($1, $2)", namespaceURI, localName)
	return w
}

func (*SVGFEColorMatrixElement) HasAttribute(name string) (b bool) {
	macro.Rewrite("$_.hasAttribute($1)", name)
	return b
}

func (*SVGFEColorMatrixElement) HasAttributeNS(namespaceURI string, localName string) (b bool) {
	macro.Rewrite("$_.hasAttributeNS($1, $2)", namespaceURI, localName)
	return b
}

func (*SVGFEColorMatrixElement) MsGetRegionContent() (w *ms.MSRangeCollection) {
	macro.Rewrite("$_.msGetRegionContent()")
	return w
}

func (*SVGFEColorMatrixElement) MsGetUntransformedBounds() (c *ClientRect) {
	macro.Rewrite("$_.msGetUntransformedBounds()")
	return c
}

func (*SVGFEColorMatrixElement) MsMatchesSelector(selectors string) (b bool) {
	macro.Rewrite("$_.msMatchesSelector($1)", selectors)
	return b
}

func (*SVGFEColorMatrixElement) MsReleasePointerCapture(pointerId int) {
	macro.Rewrite("$_.msReleasePointerCapture($1)", pointerId)
}

func (*SVGFEColorMatrixElement) MsSetPointerCapture(pointerId int) {
	macro.Rewrite("$_.msSetPointerCapture($1)", pointerId)
}

func (*SVGFEColorMatrixElement) MsZoomTo(args *ms.MsZoomToOptions) {
	macro.Rewrite("$_.msZoomTo($1)", args)
}

func (*SVGFEColorMatrixElement) ReleasePointerCapture(pointerId int) {
	macro.Rewrite("$_.releasePointerCapture($1)", pointerId)
}

func (*SVGFEColorMatrixElement) RemoveAttribute(qualifiedName string) {
	macro.Rewrite("$_.removeAttribute($1)", qualifiedName)
}

func (*SVGFEColorMatrixElement) RemoveAttributeNode(oldAttr *Attr) (w *Attr) {
	macro.Rewrite("$_.removeAttributeNode($1)", oldAttr)
	return w
}

func (*SVGFEColorMatrixElement) RemoveAttributeNS(namespaceURI string, localName string) {
	macro.Rewrite("$_.removeAttributeNS($1, $2)", namespaceURI, localName)
}

func (*SVGFEColorMatrixElement) RequestFullscreen() {
	macro.Rewrite("$_.requestFullscreen()")
}

func (*SVGFEColorMatrixElement) RequestPointerLock() {
	macro.Rewrite("$_.requestPointerLock()")
}

func (*SVGFEColorMatrixElement) SetAttribute(qualifiedName string, value string) {
	macro.Rewrite("$_.setAttribute($1, $2)", qualifiedName, value)
}

func (*SVGFEColorMatrixElement) SetAttributeNode(newAttr *Attr) (w *Attr) {
	macro.Rewrite("$_.setAttributeNode($1)", newAttr)
	return w
}

func (*SVGFEColorMatrixElement) SetAttributeNodeNS(newAttr *Attr) (w *Attr) {
	macro.Rewrite("$_.setAttributeNodeNS($1)", newAttr)
	return w
}

func (*SVGFEColorMatrixElement) SetAttributeNS(namespaceURI string, qualifiedName string, value string) {
	macro.Rewrite("$_.setAttributeNS($1, $2, $3)", namespaceURI, qualifiedName, value)
}

func (*SVGFEColorMatrixElement) SetPointerCapture(pointerId int) {
	macro.Rewrite("$_.setPointerCapture($1)", pointerId)
}

func (*SVGFEColorMatrixElement) WebkitMatchesSelector(selectors string) (b bool) {
	macro.Rewrite("$_.webkitMatchesSelector($1)", selectors)
	return b
}

func (*SVGFEColorMatrixElement) WebkitRequestFullscreen() {
	macro.Rewrite("$_.webkitRequestFullscreen()")
}

func (*SVGFEColorMatrixElement) WebkitRequestFullScreen() {
	macro.Rewrite("$_.webkitRequestFullScreen()")
}

func (*SVGFEColorMatrixElement) QuerySelector(selectors string) (w element.Element) {
	macro.Rewrite("$_.querySelector($1)", selectors)
	return w
}

func (*SVGFEColorMatrixElement) QuerySelectorAll(selectors string) (w *NodeList) {
	macro.Rewrite("$_.querySelectorAll($1)", selectors)
	return w
}

func (*SVGFEColorMatrixElement) AppendChild(newChild Node) (w Node) {
	macro.Rewrite("$_.appendChild($1)", newChild)
	return w
}

func (*SVGFEColorMatrixElement) CloneNode(deep *bool) (w Node) {
	macro.Rewrite("$_.cloneNode($1)", deep)
	return w
}

func (*SVGFEColorMatrixElement) CompareDocumentPosition(other Node) (u uint8) {
	macro.Rewrite("$_.compareDocumentPosition($1)", other)
	return u
}

func (*SVGFEColorMatrixElement) Contains(child Node) (b bool) {
	macro.Rewrite("$_.contains($1)", child)
	return b
}

func (*SVGFEColorMatrixElement) HasAttributes() (b bool) {
	macro.Rewrite("$_.hasAttributes()")
	return b
}

func (*SVGFEColorMatrixElement) HasChildNodes() (b bool) {
	macro.Rewrite("$_.hasChildNodes()")
	return b
}

func (*SVGFEColorMatrixElement) InsertBefore(newChild Node, refChild *Node) (w Node) {
	macro.Rewrite("$_.insertBefore($1, $2)", newChild, refChild)
	return w
}

func (*SVGFEColorMatrixElement) IsDefaultNamespace(namespaceURI string) (b bool) {
	macro.Rewrite("$_.isDefaultNamespace($1)", namespaceURI)
	return b
}

func (*SVGFEColorMatrixElement) IsEqualNode(arg Node) (b bool) {
	macro.Rewrite("$_.isEqualNode($1)", arg)
	return b
}

func (*SVGFEColorMatrixElement) IsSameNode(other Node) (b bool) {
	macro.Rewrite("$_.isSameNode($1)", other)
	return b
}

func (*SVGFEColorMatrixElement) LookupNamespaceURI(prefix string) (s string) {
	macro.Rewrite("$_.lookupNamespaceURI($1)", prefix)
	return s
}

func (*SVGFEColorMatrixElement) LookupPrefix(namespaceURI string) (s string) {
	macro.Rewrite("$_.lookupPrefix($1)", namespaceURI)
	return s
}

func (*SVGFEColorMatrixElement) Normalize() {
	macro.Rewrite("$_.normalize()")
}

func (*SVGFEColorMatrixElement) RemoveChild(oldChild Node) (w Node) {
	macro.Rewrite("$_.removeChild($1)", oldChild)
	return w
}

func (*SVGFEColorMatrixElement) ReplaceChild(newChild Node, oldChild Node) (w Node) {
	macro.Rewrite("$_.replaceChild($1, $2)", newChild, oldChild)
	return w
}

func (*SVGFEColorMatrixElement) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*SVGFEColorMatrixElement) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*SVGFEColorMatrixElement) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*SVGFEColorMatrixElement) In1() (in1 *SVGAnimatedString) {
	macro.Rewrite("$_.in1")
	return in1
}

func (*SVGFEColorMatrixElement) Type() (kind *SVGAnimatedEnumeration) {
	macro.Rewrite("$_.type")
	return kind
}

func (*SVGFEColorMatrixElement) Values() (values *SVGAnimatedNumberList) {
	macro.Rewrite("$_.values")
	return values
}

func (*SVGFEColorMatrixElement) Height() (height *SVGAnimatedLength) {
	macro.Rewrite("$_.height")
	return height
}

func (*SVGFEColorMatrixElement) Result() (result *SVGAnimatedString) {
	macro.Rewrite("$_.result")
	return result
}

func (*SVGFEColorMatrixElement) Width() (width *SVGAnimatedLength) {
	macro.Rewrite("$_.width")
	return width
}

func (*SVGFEColorMatrixElement) X() (x *SVGAnimatedLength) {
	macro.Rewrite("$_.x")
	return x
}

func (*SVGFEColorMatrixElement) Y() (y *SVGAnimatedLength) {
	macro.Rewrite("$_.y")
	return y
}

func (*SVGFEColorMatrixElement) Dataset() (dataset *DOMStringMap) {
	macro.Rewrite("$_.dataset")
	return dataset
}

func (*SVGFEColorMatrixElement) OwnerSVGElement() (ownerSVGElement *SVGSVGElement) {
	macro.Rewrite("$_.ownerSVGElement")
	return ownerSVGElement
}

func (*SVGFEColorMatrixElement) ViewportElement() (viewportElement SVGElement) {
	macro.Rewrite("$_.viewportElement")
	return viewportElement
}

func (*SVGFEColorMatrixElement) Xmlbase() (xmlbase string) {
	macro.Rewrite("$_.xmlbase")
	return xmlbase
}

func (*SVGFEColorMatrixElement) SetXmlbase(xmlbase string) {
	macro.Rewrite("$_.xmlbase = $1", xmlbase)
}

func (*SVGFEColorMatrixElement) ClassList() (classList DOMTokenList) {
	macro.Rewrite("$_.classList")
	return classList
}

func (*SVGFEColorMatrixElement) ClassName() (className string) {
	macro.Rewrite("$_.className")
	return className
}

func (*SVGFEColorMatrixElement) SetClassName(className string) {
	macro.Rewrite("$_.className = $1", className)
}

func (*SVGFEColorMatrixElement) ClientHeight() (clientHeight int) {
	macro.Rewrite("$_.clientHeight")
	return clientHeight
}

func (*SVGFEColorMatrixElement) ClientLeft() (clientLeft int) {
	macro.Rewrite("$_.clientLeft")
	return clientLeft
}

func (*SVGFEColorMatrixElement) ClientTop() (clientTop int) {
	macro.Rewrite("$_.clientTop")
	return clientTop
}

func (*SVGFEColorMatrixElement) ClientWidth() (clientWidth int) {
	macro.Rewrite("$_.clientWidth")
	return clientWidth
}

func (*SVGFEColorMatrixElement) ID() (id string) {
	macro.Rewrite("$_.id")
	return id
}

func (*SVGFEColorMatrixElement) SetID(id string) {
	macro.Rewrite("$_.id = $1", id)
}

func (*SVGFEColorMatrixElement) InnerHTML() (innerHTML string) {
	macro.Rewrite("$_.innerHTML")
	return innerHTML
}

func (*SVGFEColorMatrixElement) SetInnerHTML(innerHTML string) {
	macro.Rewrite("$_.innerHTML = $1", innerHTML)
}

func (*SVGFEColorMatrixElement) MsContentZoomFactor() (msContentZoomFactor float32) {
	macro.Rewrite("$_.msContentZoomFactor")
	return msContentZoomFactor
}

func (*SVGFEColorMatrixElement) SetMsContentZoomFactor(msContentZoomFactor float32) {
	macro.Rewrite("$_.msContentZoomFactor = $1", msContentZoomFactor)
}

func (*SVGFEColorMatrixElement) MsRegionOverflow() (msRegionOverflow string) {
	macro.Rewrite("$_.msRegionOverflow")
	return msRegionOverflow
}

func (*SVGFEColorMatrixElement) Onariarequest() (onariarequest func(event.Event)) {
	macro.Rewrite("$_.onariarequest")
	return onariarequest
}

func (*SVGFEColorMatrixElement) SetOnariarequest(onariarequest func(event.Event)) {
	macro.Rewrite("$_.onariarequest = $1", onariarequest)
}

func (*SVGFEColorMatrixElement) Oncommand() (oncommand func(event.Event)) {
	macro.Rewrite("$_.oncommand")
	return oncommand
}

func (*SVGFEColorMatrixElement) SetOncommand(oncommand func(event.Event)) {
	macro.Rewrite("$_.oncommand = $1", oncommand)
}

func (*SVGFEColorMatrixElement) Ongotpointercapture() (ongotpointercapture func(*utils.PointerEvent)) {
	macro.Rewrite("$_.ongotpointercapture")
	return ongotpointercapture
}

func (*SVGFEColorMatrixElement) SetOngotpointercapture(ongotpointercapture func(*utils.PointerEvent)) {
	macro.Rewrite("$_.ongotpointercapture = $1", ongotpointercapture)
}

func (*SVGFEColorMatrixElement) Onlostpointercapture() (onlostpointercapture func(*utils.PointerEvent)) {
	macro.Rewrite("$_.onlostpointercapture")
	return onlostpointercapture
}

func (*SVGFEColorMatrixElement) SetOnlostpointercapture(onlostpointercapture func(*utils.PointerEvent)) {
	macro.Rewrite("$_.onlostpointercapture = $1", onlostpointercapture)
}

func (*SVGFEColorMatrixElement) Onmsgesturechange() (onmsgesturechange func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturechange")
	return onmsgesturechange
}

func (*SVGFEColorMatrixElement) SetOnmsgesturechange(onmsgesturechange func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturechange = $1", onmsgesturechange)
}

func (*SVGFEColorMatrixElement) Onmsgesturedoubletap() (onmsgesturedoubletap func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturedoubletap")
	return onmsgesturedoubletap
}

func (*SVGFEColorMatrixElement) SetOnmsgesturedoubletap(onmsgesturedoubletap func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturedoubletap = $1", onmsgesturedoubletap)
}

func (*SVGFEColorMatrixElement) Onmsgestureend() (onmsgestureend func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgestureend")
	return onmsgestureend
}

func (*SVGFEColorMatrixElement) SetOnmsgestureend(onmsgestureend func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgestureend = $1", onmsgestureend)
}

func (*SVGFEColorMatrixElement) Onmsgesturehold() (onmsgesturehold func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturehold")
	return onmsgesturehold
}

func (*SVGFEColorMatrixElement) SetOnmsgesturehold(onmsgesturehold func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturehold = $1", onmsgesturehold)
}

func (*SVGFEColorMatrixElement) Onmsgesturestart() (onmsgesturestart func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturestart")
	return onmsgesturestart
}

func (*SVGFEColorMatrixElement) SetOnmsgesturestart(onmsgesturestart func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturestart = $1", onmsgesturestart)
}

func (*SVGFEColorMatrixElement) Onmsgesturetap() (onmsgesturetap func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturetap")
	return onmsgesturetap
}

func (*SVGFEColorMatrixElement) SetOnmsgesturetap(onmsgesturetap func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturetap = $1", onmsgesturetap)
}

func (*SVGFEColorMatrixElement) Onmsgotpointercapture() (onmsgotpointercapture func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmsgotpointercapture")
	return onmsgotpointercapture
}

func (*SVGFEColorMatrixElement) SetOnmsgotpointercapture(onmsgotpointercapture func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmsgotpointercapture = $1", onmsgotpointercapture)
}

func (*SVGFEColorMatrixElement) Onmsinertiastart() (onmsinertiastart func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsinertiastart")
	return onmsinertiastart
}

func (*SVGFEColorMatrixElement) SetOnmsinertiastart(onmsinertiastart func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsinertiastart = $1", onmsinertiastart)
}

func (*SVGFEColorMatrixElement) Onmslostpointercapture() (onmslostpointercapture func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmslostpointercapture")
	return onmslostpointercapture
}

func (*SVGFEColorMatrixElement) SetOnmslostpointercapture(onmslostpointercapture func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmslostpointercapture = $1", onmslostpointercapture)
}

func (*SVGFEColorMatrixElement) Onmspointercancel() (onmspointercancel func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointercancel")
	return onmspointercancel
}

func (*SVGFEColorMatrixElement) SetOnmspointercancel(onmspointercancel func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointercancel = $1", onmspointercancel)
}

func (*SVGFEColorMatrixElement) Onmspointerdown() (onmspointerdown func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerdown")
	return onmspointerdown
}

func (*SVGFEColorMatrixElement) SetOnmspointerdown(onmspointerdown func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerdown = $1", onmspointerdown)
}

func (*SVGFEColorMatrixElement) Onmspointerenter() (onmspointerenter func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerenter")
	return onmspointerenter
}

func (*SVGFEColorMatrixElement) SetOnmspointerenter(onmspointerenter func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerenter = $1", onmspointerenter)
}

func (*SVGFEColorMatrixElement) Onmspointerleave() (onmspointerleave func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerleave")
	return onmspointerleave
}

func (*SVGFEColorMatrixElement) SetOnmspointerleave(onmspointerleave func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerleave = $1", onmspointerleave)
}

func (*SVGFEColorMatrixElement) Onmspointermove() (onmspointermove func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointermove")
	return onmspointermove
}

func (*SVGFEColorMatrixElement) SetOnmspointermove(onmspointermove func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointermove = $1", onmspointermove)
}

func (*SVGFEColorMatrixElement) Onmspointerout() (onmspointerout func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerout")
	return onmspointerout
}

func (*SVGFEColorMatrixElement) SetOnmspointerout(onmspointerout func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerout = $1", onmspointerout)
}

func (*SVGFEColorMatrixElement) Onmspointerover() (onmspointerover func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerover")
	return onmspointerover
}

func (*SVGFEColorMatrixElement) SetOnmspointerover(onmspointerover func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerover = $1", onmspointerover)
}

func (*SVGFEColorMatrixElement) Onmspointerup() (onmspointerup func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerup")
	return onmspointerup
}

func (*SVGFEColorMatrixElement) SetOnmspointerup(onmspointerup func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerup = $1", onmspointerup)
}

func (*SVGFEColorMatrixElement) Ontouchcancel() (ontouchcancel func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchcancel")
	return ontouchcancel
}

func (*SVGFEColorMatrixElement) SetOntouchcancel(ontouchcancel func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchcancel = $1", ontouchcancel)
}

func (*SVGFEColorMatrixElement) Ontouchend() (ontouchend func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchend")
	return ontouchend
}

func (*SVGFEColorMatrixElement) SetOntouchend(ontouchend func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchend = $1", ontouchend)
}

func (*SVGFEColorMatrixElement) Ontouchmove() (ontouchmove func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchmove")
	return ontouchmove
}

func (*SVGFEColorMatrixElement) SetOntouchmove(ontouchmove func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchmove = $1", ontouchmove)
}

func (*SVGFEColorMatrixElement) Ontouchstart() (ontouchstart func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchstart")
	return ontouchstart
}

func (*SVGFEColorMatrixElement) SetOntouchstart(ontouchstart func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchstart = $1", ontouchstart)
}

func (*SVGFEColorMatrixElement) Onwebkitfullscreenchange() (onwebkitfullscreenchange func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenchange")
	return onwebkitfullscreenchange
}

func (*SVGFEColorMatrixElement) SetOnwebkitfullscreenchange(onwebkitfullscreenchange func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenchange = $1", onwebkitfullscreenchange)
}

func (*SVGFEColorMatrixElement) Onwebkitfullscreenerror() (onwebkitfullscreenerror func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenerror")
	return onwebkitfullscreenerror
}

func (*SVGFEColorMatrixElement) SetOnwebkitfullscreenerror(onwebkitfullscreenerror func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenerror = $1", onwebkitfullscreenerror)
}

func (*SVGFEColorMatrixElement) OuterHTML() (outerHTML string) {
	macro.Rewrite("$_.outerHTML")
	return outerHTML
}

func (*SVGFEColorMatrixElement) SetOuterHTML(outerHTML string) {
	macro.Rewrite("$_.outerHTML = $1", outerHTML)
}

func (*SVGFEColorMatrixElement) Prefix() (prefix string) {
	macro.Rewrite("$_.prefix")
	return prefix
}

func (*SVGFEColorMatrixElement) ScrollHeight() (scrollHeight int) {
	macro.Rewrite("$_.scrollHeight")
	return scrollHeight
}

func (*SVGFEColorMatrixElement) ScrollLeft() (scrollLeft int) {
	macro.Rewrite("$_.scrollLeft")
	return scrollLeft
}

func (*SVGFEColorMatrixElement) SetScrollLeft(scrollLeft int) {
	macro.Rewrite("$_.scrollLeft = $1", scrollLeft)
}

func (*SVGFEColorMatrixElement) ScrollTop() (scrollTop int) {
	macro.Rewrite("$_.scrollTop")
	return scrollTop
}

func (*SVGFEColorMatrixElement) SetScrollTop(scrollTop int) {
	macro.Rewrite("$_.scrollTop = $1", scrollTop)
}

func (*SVGFEColorMatrixElement) ScrollWidth() (scrollWidth int) {
	macro.Rewrite("$_.scrollWidth")
	return scrollWidth
}

func (*SVGFEColorMatrixElement) TagName() (tagName string) {
	macro.Rewrite("$_.tagName")
	return tagName
}

func (*SVGFEColorMatrixElement) Onpointercancel() (onpointercancel func(event.Event)) {
	macro.Rewrite("$_.onpointercancel")
	return onpointercancel
}

func (*SVGFEColorMatrixElement) SetOnpointercancel(onpointercancel func(event.Event)) {
	macro.Rewrite("$_.onpointercancel = $1", onpointercancel)
}

func (*SVGFEColorMatrixElement) Onpointerdown() (onpointerdown func(event.Event)) {
	macro.Rewrite("$_.onpointerdown")
	return onpointerdown
}

func (*SVGFEColorMatrixElement) SetOnpointerdown(onpointerdown func(event.Event)) {
	macro.Rewrite("$_.onpointerdown = $1", onpointerdown)
}

func (*SVGFEColorMatrixElement) Onpointerenter() (onpointerenter func(event.Event)) {
	macro.Rewrite("$_.onpointerenter")
	return onpointerenter
}

func (*SVGFEColorMatrixElement) SetOnpointerenter(onpointerenter func(event.Event)) {
	macro.Rewrite("$_.onpointerenter = $1", onpointerenter)
}

func (*SVGFEColorMatrixElement) Onpointerleave() (onpointerleave func(event.Event)) {
	macro.Rewrite("$_.onpointerleave")
	return onpointerleave
}

func (*SVGFEColorMatrixElement) SetOnpointerleave(onpointerleave func(event.Event)) {
	macro.Rewrite("$_.onpointerleave = $1", onpointerleave)
}

func (*SVGFEColorMatrixElement) Onpointermove() (onpointermove func(event.Event)) {
	macro.Rewrite("$_.onpointermove")
	return onpointermove
}

func (*SVGFEColorMatrixElement) SetOnpointermove(onpointermove func(event.Event)) {
	macro.Rewrite("$_.onpointermove = $1", onpointermove)
}

func (*SVGFEColorMatrixElement) Onpointerout() (onpointerout func(event.Event)) {
	macro.Rewrite("$_.onpointerout")
	return onpointerout
}

func (*SVGFEColorMatrixElement) SetOnpointerout(onpointerout func(event.Event)) {
	macro.Rewrite("$_.onpointerout = $1", onpointerout)
}

func (*SVGFEColorMatrixElement) Onpointerover() (onpointerover func(event.Event)) {
	macro.Rewrite("$_.onpointerover")
	return onpointerover
}

func (*SVGFEColorMatrixElement) SetOnpointerover(onpointerover func(event.Event)) {
	macro.Rewrite("$_.onpointerover = $1", onpointerover)
}

func (*SVGFEColorMatrixElement) Onpointerup() (onpointerup func(event.Event)) {
	macro.Rewrite("$_.onpointerup")
	return onpointerup
}

func (*SVGFEColorMatrixElement) SetOnpointerup(onpointerup func(event.Event)) {
	macro.Rewrite("$_.onpointerup = $1", onpointerup)
}

func (*SVGFEColorMatrixElement) Onwheel() (onwheel func(event.Event)) {
	macro.Rewrite("$_.onwheel")
	return onwheel
}

func (*SVGFEColorMatrixElement) SetOnwheel(onwheel func(event.Event)) {
	macro.Rewrite("$_.onwheel = $1", onwheel)
}

func (*SVGFEColorMatrixElement) ChildElementCount() (childElementCount uint) {
	macro.Rewrite("$_.childElementCount")
	return childElementCount
}

func (*SVGFEColorMatrixElement) FirstElementChild() (firstElementChild element.Element) {
	macro.Rewrite("$_.firstElementChild")
	return firstElementChild
}

func (*SVGFEColorMatrixElement) LastElementChild() (lastElementChild element.Element) {
	macro.Rewrite("$_.lastElementChild")
	return lastElementChild
}

func (*SVGFEColorMatrixElement) NextElementSibling() (nextElementSibling element.Element) {
	macro.Rewrite("$_.nextElementSibling")
	return nextElementSibling
}

func (*SVGFEColorMatrixElement) PreviousElementSibling() (previousElementSibling element.Element) {
	macro.Rewrite("$_.previousElementSibling")
	return previousElementSibling
}

func (*SVGFEColorMatrixElement) Attributes() (attributes *NamedNodeMap) {
	macro.Rewrite("$_.attributes")
	return attributes
}

func (*SVGFEColorMatrixElement) BaseURI() (baseURI string) {
	macro.Rewrite("$_.baseURI")
	return baseURI
}

func (*SVGFEColorMatrixElement) ChildNodes() (childNodes *NodeList) {
	macro.Rewrite("$_.childNodes")
	return childNodes
}

func (*SVGFEColorMatrixElement) FirstChild() (firstChild Node) {
	macro.Rewrite("$_.firstChild")
	return firstChild
}

func (*SVGFEColorMatrixElement) LastChild() (lastChild Node) {
	macro.Rewrite("$_.lastChild")
	return lastChild
}

func (*SVGFEColorMatrixElement) LocalName() (localName string) {
	macro.Rewrite("$_.localName")
	return localName
}

func (*SVGFEColorMatrixElement) NamespaceURI() (namespaceURI string) {
	macro.Rewrite("$_.namespaceURI")
	return namespaceURI
}

func (*SVGFEColorMatrixElement) NextSibling() (nextSibling Node) {
	macro.Rewrite("$_.nextSibling")
	return nextSibling
}

func (*SVGFEColorMatrixElement) NodeName() (nodeName string) {
	macro.Rewrite("$_.nodeName")
	return nodeName
}

func (*SVGFEColorMatrixElement) NodeType() (nodeType uint8) {
	macro.Rewrite("$_.nodeType")
	return nodeType
}

func (*SVGFEColorMatrixElement) NodeValue() (nodeValue string) {
	macro.Rewrite("$_.nodeValue")
	return nodeValue
}

func (*SVGFEColorMatrixElement) SetNodeValue(nodeValue string) {
	macro.Rewrite("$_.nodeValue = $1", nodeValue)
}

func (*SVGFEColorMatrixElement) OwnerDocument() (ownerDocument document.Document) {
	macro.Rewrite("$_.ownerDocument")
	return ownerDocument
}

func (*SVGFEColorMatrixElement) ParentElement() (parentElement html.HTMLElement) {
	macro.Rewrite("$_.parentElement")
	return parentElement
}

func (*SVGFEColorMatrixElement) ParentNode() (parentNode Node) {
	macro.Rewrite("$_.parentNode")
	return parentNode
}

func (*SVGFEColorMatrixElement) PreviousSibling() (previousSibling Node) {
	macro.Rewrite("$_.previousSibling")
	return previousSibling
}

func (*SVGFEColorMatrixElement) TextContent() (textContent string) {
	macro.Rewrite("$_.textContent")
	return textContent
}

func (*SVGFEColorMatrixElement) SetTextContent(textContent string) {
	macro.Rewrite("$_.textContent = $1", textContent)
}
