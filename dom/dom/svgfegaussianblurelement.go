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

var _ SVGElement = (*SVGFEGaussianBlurElement)(nil)
var _ element.Element = (*SVGFEGaussianBlurElement)(nil)
var _ utils.GlobalEventHandlers = (*SVGFEGaussianBlurElement)(nil)
var _ element.ElementTraversal = (*SVGFEGaussianBlurElement)(nil)
var _ NodeSelector = (*SVGFEGaussianBlurElement)(nil)
var _ ChildNode = (*SVGFEGaussianBlurElement)(nil)
var _ Node = (*SVGFEGaussianBlurElement)(nil)
var _ event.EventTarget = (*SVGFEGaussianBlurElement)(nil)

type SVGFEGaussianBlurElement struct {
}

func (*SVGFEGaussianBlurElement) SetStdDeviation(stdDeviationX float32, stdDeviationY float32) {
	macro.Rewrite("$_.setStdDeviation($1, $2)", stdDeviationX, stdDeviationY)
}

func (*SVGFEGaussianBlurElement) GetAttribute(qualifiedName string) (s string) {
	macro.Rewrite("$_.getAttribute($1)", qualifiedName)
	return s
}

func (*SVGFEGaussianBlurElement) GetAttributeNode(name string) (w *Attr) {
	macro.Rewrite("$_.getAttributeNode($1)", name)
	return w
}

func (*SVGFEGaussianBlurElement) GetAttributeNodeNS(namespaceURI string, localName string) (w *Attr) {
	macro.Rewrite("$_.getAttributeNodeNS($1, $2)", namespaceURI, localName)
	return w
}

func (*SVGFEGaussianBlurElement) GetAttributeNS(namespaceURI string, localName string) (s string) {
	macro.Rewrite("$_.getAttributeNS($1, $2)", namespaceURI, localName)
	return s
}

func (*SVGFEGaussianBlurElement) GetBoundingClientRect() (c *ClientRect) {
	macro.Rewrite("$_.getBoundingClientRect()")
	return c
}

func (*SVGFEGaussianBlurElement) GetClientRects() (c *ClientRectList) {
	macro.Rewrite("$_.getClientRects()")
	return c
}

func (*SVGFEGaussianBlurElement) GetElementsByTagName(name string) (w *NodeList) {
	macro.Rewrite("$_.getElementsByTagName($1)", name)
	return w
}

func (*SVGFEGaussianBlurElement) GetElementsByTagNameNS(namespaceURI string, localName string) (w *NodeList) {
	macro.Rewrite("$_.getElementsByTagNameNS($1, $2)", namespaceURI, localName)
	return w
}

func (*SVGFEGaussianBlurElement) HasAttribute(name string) (b bool) {
	macro.Rewrite("$_.hasAttribute($1)", name)
	return b
}

func (*SVGFEGaussianBlurElement) HasAttributeNS(namespaceURI string, localName string) (b bool) {
	macro.Rewrite("$_.hasAttributeNS($1, $2)", namespaceURI, localName)
	return b
}

func (*SVGFEGaussianBlurElement) MsGetRegionContent() (w *ms.MSRangeCollection) {
	macro.Rewrite("$_.msGetRegionContent()")
	return w
}

func (*SVGFEGaussianBlurElement) MsGetUntransformedBounds() (c *ClientRect) {
	macro.Rewrite("$_.msGetUntransformedBounds()")
	return c
}

func (*SVGFEGaussianBlurElement) MsMatchesSelector(selectors string) (b bool) {
	macro.Rewrite("$_.msMatchesSelector($1)", selectors)
	return b
}

func (*SVGFEGaussianBlurElement) MsReleasePointerCapture(pointerId int) {
	macro.Rewrite("$_.msReleasePointerCapture($1)", pointerId)
}

func (*SVGFEGaussianBlurElement) MsSetPointerCapture(pointerId int) {
	macro.Rewrite("$_.msSetPointerCapture($1)", pointerId)
}

func (*SVGFEGaussianBlurElement) MsZoomTo(args *ms.MsZoomToOptions) {
	macro.Rewrite("$_.msZoomTo($1)", args)
}

func (*SVGFEGaussianBlurElement) ReleasePointerCapture(pointerId int) {
	macro.Rewrite("$_.releasePointerCapture($1)", pointerId)
}

func (*SVGFEGaussianBlurElement) RemoveAttribute(qualifiedName string) {
	macro.Rewrite("$_.removeAttribute($1)", qualifiedName)
}

func (*SVGFEGaussianBlurElement) RemoveAttributeNode(oldAttr *Attr) (w *Attr) {
	macro.Rewrite("$_.removeAttributeNode($1)", oldAttr)
	return w
}

func (*SVGFEGaussianBlurElement) RemoveAttributeNS(namespaceURI string, localName string) {
	macro.Rewrite("$_.removeAttributeNS($1, $2)", namespaceURI, localName)
}

func (*SVGFEGaussianBlurElement) RequestFullscreen() {
	macro.Rewrite("$_.requestFullscreen()")
}

func (*SVGFEGaussianBlurElement) RequestPointerLock() {
	macro.Rewrite("$_.requestPointerLock()")
}

func (*SVGFEGaussianBlurElement) SetAttribute(qualifiedName string, value string) {
	macro.Rewrite("$_.setAttribute($1, $2)", qualifiedName, value)
}

func (*SVGFEGaussianBlurElement) SetAttributeNode(newAttr *Attr) (w *Attr) {
	macro.Rewrite("$_.setAttributeNode($1)", newAttr)
	return w
}

func (*SVGFEGaussianBlurElement) SetAttributeNodeNS(newAttr *Attr) (w *Attr) {
	macro.Rewrite("$_.setAttributeNodeNS($1)", newAttr)
	return w
}

func (*SVGFEGaussianBlurElement) SetAttributeNS(namespaceURI string, qualifiedName string, value string) {
	macro.Rewrite("$_.setAttributeNS($1, $2, $3)", namespaceURI, qualifiedName, value)
}

func (*SVGFEGaussianBlurElement) SetPointerCapture(pointerId int) {
	macro.Rewrite("$_.setPointerCapture($1)", pointerId)
}

func (*SVGFEGaussianBlurElement) WebkitMatchesSelector(selectors string) (b bool) {
	macro.Rewrite("$_.webkitMatchesSelector($1)", selectors)
	return b
}

func (*SVGFEGaussianBlurElement) WebkitRequestFullscreen() {
	macro.Rewrite("$_.webkitRequestFullscreen()")
}

func (*SVGFEGaussianBlurElement) WebkitRequestFullScreen() {
	macro.Rewrite("$_.webkitRequestFullScreen()")
}

func (*SVGFEGaussianBlurElement) QuerySelector(selectors string) (w element.Element) {
	macro.Rewrite("$_.querySelector($1)", selectors)
	return w
}

func (*SVGFEGaussianBlurElement) QuerySelectorAll(selectors string) (w *NodeList) {
	macro.Rewrite("$_.querySelectorAll($1)", selectors)
	return w
}

func (*SVGFEGaussianBlurElement) AppendChild(newChild Node) (w Node) {
	macro.Rewrite("$_.appendChild($1)", newChild)
	return w
}

func (*SVGFEGaussianBlurElement) CloneNode(deep *bool) (w Node) {
	macro.Rewrite("$_.cloneNode($1)", deep)
	return w
}

func (*SVGFEGaussianBlurElement) CompareDocumentPosition(other Node) (u uint8) {
	macro.Rewrite("$_.compareDocumentPosition($1)", other)
	return u
}

func (*SVGFEGaussianBlurElement) Contains(child Node) (b bool) {
	macro.Rewrite("$_.contains($1)", child)
	return b
}

func (*SVGFEGaussianBlurElement) HasAttributes() (b bool) {
	macro.Rewrite("$_.hasAttributes()")
	return b
}

func (*SVGFEGaussianBlurElement) HasChildNodes() (b bool) {
	macro.Rewrite("$_.hasChildNodes()")
	return b
}

func (*SVGFEGaussianBlurElement) InsertBefore(newChild Node, refChild *Node) (w Node) {
	macro.Rewrite("$_.insertBefore($1, $2)", newChild, refChild)
	return w
}

func (*SVGFEGaussianBlurElement) IsDefaultNamespace(namespaceURI string) (b bool) {
	macro.Rewrite("$_.isDefaultNamespace($1)", namespaceURI)
	return b
}

func (*SVGFEGaussianBlurElement) IsEqualNode(arg Node) (b bool) {
	macro.Rewrite("$_.isEqualNode($1)", arg)
	return b
}

func (*SVGFEGaussianBlurElement) IsSameNode(other Node) (b bool) {
	macro.Rewrite("$_.isSameNode($1)", other)
	return b
}

func (*SVGFEGaussianBlurElement) LookupNamespaceURI(prefix string) (s string) {
	macro.Rewrite("$_.lookupNamespaceURI($1)", prefix)
	return s
}

func (*SVGFEGaussianBlurElement) LookupPrefix(namespaceURI string) (s string) {
	macro.Rewrite("$_.lookupPrefix($1)", namespaceURI)
	return s
}

func (*SVGFEGaussianBlurElement) Normalize() {
	macro.Rewrite("$_.normalize()")
}

func (*SVGFEGaussianBlurElement) RemoveChild(oldChild Node) (w Node) {
	macro.Rewrite("$_.removeChild($1)", oldChild)
	return w
}

func (*SVGFEGaussianBlurElement) ReplaceChild(newChild Node, oldChild Node) (w Node) {
	macro.Rewrite("$_.replaceChild($1, $2)", newChild, oldChild)
	return w
}

func (*SVGFEGaussianBlurElement) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*SVGFEGaussianBlurElement) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*SVGFEGaussianBlurElement) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*SVGFEGaussianBlurElement) In1() (in1 *SVGAnimatedString) {
	macro.Rewrite("$_.in1")
	return in1
}

func (*SVGFEGaussianBlurElement) StdDeviationX() (stdDeviationX *SVGAnimatedNumber) {
	macro.Rewrite("$_.stdDeviationX")
	return stdDeviationX
}

func (*SVGFEGaussianBlurElement) StdDeviationY() (stdDeviationY *SVGAnimatedNumber) {
	macro.Rewrite("$_.stdDeviationY")
	return stdDeviationY
}

func (*SVGFEGaussianBlurElement) Height() (height *SVGAnimatedLength) {
	macro.Rewrite("$_.height")
	return height
}

func (*SVGFEGaussianBlurElement) Result() (result *SVGAnimatedString) {
	macro.Rewrite("$_.result")
	return result
}

func (*SVGFEGaussianBlurElement) Width() (width *SVGAnimatedLength) {
	macro.Rewrite("$_.width")
	return width
}

func (*SVGFEGaussianBlurElement) X() (x *SVGAnimatedLength) {
	macro.Rewrite("$_.x")
	return x
}

func (*SVGFEGaussianBlurElement) Y() (y *SVGAnimatedLength) {
	macro.Rewrite("$_.y")
	return y
}

func (*SVGFEGaussianBlurElement) Dataset() (dataset *DOMStringMap) {
	macro.Rewrite("$_.dataset")
	return dataset
}

func (*SVGFEGaussianBlurElement) OwnerSVGElement() (ownerSVGElement *SVGSVGElement) {
	macro.Rewrite("$_.ownerSVGElement")
	return ownerSVGElement
}

func (*SVGFEGaussianBlurElement) ViewportElement() (viewportElement SVGElement) {
	macro.Rewrite("$_.viewportElement")
	return viewportElement
}

func (*SVGFEGaussianBlurElement) Xmlbase() (xmlbase string) {
	macro.Rewrite("$_.xmlbase")
	return xmlbase
}

func (*SVGFEGaussianBlurElement) SetXmlbase(xmlbase string) {
	macro.Rewrite("$_.xmlbase = $1", xmlbase)
}

func (*SVGFEGaussianBlurElement) ClassList() (classList DOMTokenList) {
	macro.Rewrite("$_.classList")
	return classList
}

func (*SVGFEGaussianBlurElement) ClassName() (className string) {
	macro.Rewrite("$_.className")
	return className
}

func (*SVGFEGaussianBlurElement) SetClassName(className string) {
	macro.Rewrite("$_.className = $1", className)
}

func (*SVGFEGaussianBlurElement) ClientHeight() (clientHeight int) {
	macro.Rewrite("$_.clientHeight")
	return clientHeight
}

func (*SVGFEGaussianBlurElement) ClientLeft() (clientLeft int) {
	macro.Rewrite("$_.clientLeft")
	return clientLeft
}

func (*SVGFEGaussianBlurElement) ClientTop() (clientTop int) {
	macro.Rewrite("$_.clientTop")
	return clientTop
}

func (*SVGFEGaussianBlurElement) ClientWidth() (clientWidth int) {
	macro.Rewrite("$_.clientWidth")
	return clientWidth
}

func (*SVGFEGaussianBlurElement) ID() (id string) {
	macro.Rewrite("$_.id")
	return id
}

func (*SVGFEGaussianBlurElement) SetID(id string) {
	macro.Rewrite("$_.id = $1", id)
}

func (*SVGFEGaussianBlurElement) InnerHTML() (innerHTML string) {
	macro.Rewrite("$_.innerHTML")
	return innerHTML
}

func (*SVGFEGaussianBlurElement) SetInnerHTML(innerHTML string) {
	macro.Rewrite("$_.innerHTML = $1", innerHTML)
}

func (*SVGFEGaussianBlurElement) MsContentZoomFactor() (msContentZoomFactor float32) {
	macro.Rewrite("$_.msContentZoomFactor")
	return msContentZoomFactor
}

func (*SVGFEGaussianBlurElement) SetMsContentZoomFactor(msContentZoomFactor float32) {
	macro.Rewrite("$_.msContentZoomFactor = $1", msContentZoomFactor)
}

func (*SVGFEGaussianBlurElement) MsRegionOverflow() (msRegionOverflow string) {
	macro.Rewrite("$_.msRegionOverflow")
	return msRegionOverflow
}

func (*SVGFEGaussianBlurElement) Onariarequest() (onariarequest func(event.Event)) {
	macro.Rewrite("$_.onariarequest")
	return onariarequest
}

func (*SVGFEGaussianBlurElement) SetOnariarequest(onariarequest func(event.Event)) {
	macro.Rewrite("$_.onariarequest = $1", onariarequest)
}

func (*SVGFEGaussianBlurElement) Oncommand() (oncommand func(event.Event)) {
	macro.Rewrite("$_.oncommand")
	return oncommand
}

func (*SVGFEGaussianBlurElement) SetOncommand(oncommand func(event.Event)) {
	macro.Rewrite("$_.oncommand = $1", oncommand)
}

func (*SVGFEGaussianBlurElement) Ongotpointercapture() (ongotpointercapture func(*utils.PointerEvent)) {
	macro.Rewrite("$_.ongotpointercapture")
	return ongotpointercapture
}

func (*SVGFEGaussianBlurElement) SetOngotpointercapture(ongotpointercapture func(*utils.PointerEvent)) {
	macro.Rewrite("$_.ongotpointercapture = $1", ongotpointercapture)
}

func (*SVGFEGaussianBlurElement) Onlostpointercapture() (onlostpointercapture func(*utils.PointerEvent)) {
	macro.Rewrite("$_.onlostpointercapture")
	return onlostpointercapture
}

func (*SVGFEGaussianBlurElement) SetOnlostpointercapture(onlostpointercapture func(*utils.PointerEvent)) {
	macro.Rewrite("$_.onlostpointercapture = $1", onlostpointercapture)
}

func (*SVGFEGaussianBlurElement) Onmsgesturechange() (onmsgesturechange func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturechange")
	return onmsgesturechange
}

func (*SVGFEGaussianBlurElement) SetOnmsgesturechange(onmsgesturechange func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturechange = $1", onmsgesturechange)
}

func (*SVGFEGaussianBlurElement) Onmsgesturedoubletap() (onmsgesturedoubletap func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturedoubletap")
	return onmsgesturedoubletap
}

func (*SVGFEGaussianBlurElement) SetOnmsgesturedoubletap(onmsgesturedoubletap func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturedoubletap = $1", onmsgesturedoubletap)
}

func (*SVGFEGaussianBlurElement) Onmsgestureend() (onmsgestureend func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgestureend")
	return onmsgestureend
}

func (*SVGFEGaussianBlurElement) SetOnmsgestureend(onmsgestureend func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgestureend = $1", onmsgestureend)
}

func (*SVGFEGaussianBlurElement) Onmsgesturehold() (onmsgesturehold func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturehold")
	return onmsgesturehold
}

func (*SVGFEGaussianBlurElement) SetOnmsgesturehold(onmsgesturehold func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturehold = $1", onmsgesturehold)
}

func (*SVGFEGaussianBlurElement) Onmsgesturestart() (onmsgesturestart func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturestart")
	return onmsgesturestart
}

func (*SVGFEGaussianBlurElement) SetOnmsgesturestart(onmsgesturestart func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturestart = $1", onmsgesturestart)
}

func (*SVGFEGaussianBlurElement) Onmsgesturetap() (onmsgesturetap func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturetap")
	return onmsgesturetap
}

func (*SVGFEGaussianBlurElement) SetOnmsgesturetap(onmsgesturetap func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturetap = $1", onmsgesturetap)
}

func (*SVGFEGaussianBlurElement) Onmsgotpointercapture() (onmsgotpointercapture func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmsgotpointercapture")
	return onmsgotpointercapture
}

func (*SVGFEGaussianBlurElement) SetOnmsgotpointercapture(onmsgotpointercapture func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmsgotpointercapture = $1", onmsgotpointercapture)
}

func (*SVGFEGaussianBlurElement) Onmsinertiastart() (onmsinertiastart func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsinertiastart")
	return onmsinertiastart
}

func (*SVGFEGaussianBlurElement) SetOnmsinertiastart(onmsinertiastart func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsinertiastart = $1", onmsinertiastart)
}

func (*SVGFEGaussianBlurElement) Onmslostpointercapture() (onmslostpointercapture func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmslostpointercapture")
	return onmslostpointercapture
}

func (*SVGFEGaussianBlurElement) SetOnmslostpointercapture(onmslostpointercapture func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmslostpointercapture = $1", onmslostpointercapture)
}

func (*SVGFEGaussianBlurElement) Onmspointercancel() (onmspointercancel func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointercancel")
	return onmspointercancel
}

func (*SVGFEGaussianBlurElement) SetOnmspointercancel(onmspointercancel func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointercancel = $1", onmspointercancel)
}

func (*SVGFEGaussianBlurElement) Onmspointerdown() (onmspointerdown func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerdown")
	return onmspointerdown
}

func (*SVGFEGaussianBlurElement) SetOnmspointerdown(onmspointerdown func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerdown = $1", onmspointerdown)
}

func (*SVGFEGaussianBlurElement) Onmspointerenter() (onmspointerenter func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerenter")
	return onmspointerenter
}

func (*SVGFEGaussianBlurElement) SetOnmspointerenter(onmspointerenter func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerenter = $1", onmspointerenter)
}

func (*SVGFEGaussianBlurElement) Onmspointerleave() (onmspointerleave func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerleave")
	return onmspointerleave
}

func (*SVGFEGaussianBlurElement) SetOnmspointerleave(onmspointerleave func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerleave = $1", onmspointerleave)
}

func (*SVGFEGaussianBlurElement) Onmspointermove() (onmspointermove func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointermove")
	return onmspointermove
}

func (*SVGFEGaussianBlurElement) SetOnmspointermove(onmspointermove func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointermove = $1", onmspointermove)
}

func (*SVGFEGaussianBlurElement) Onmspointerout() (onmspointerout func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerout")
	return onmspointerout
}

func (*SVGFEGaussianBlurElement) SetOnmspointerout(onmspointerout func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerout = $1", onmspointerout)
}

func (*SVGFEGaussianBlurElement) Onmspointerover() (onmspointerover func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerover")
	return onmspointerover
}

func (*SVGFEGaussianBlurElement) SetOnmspointerover(onmspointerover func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerover = $1", onmspointerover)
}

func (*SVGFEGaussianBlurElement) Onmspointerup() (onmspointerup func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerup")
	return onmspointerup
}

func (*SVGFEGaussianBlurElement) SetOnmspointerup(onmspointerup func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerup = $1", onmspointerup)
}

func (*SVGFEGaussianBlurElement) Ontouchcancel() (ontouchcancel func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchcancel")
	return ontouchcancel
}

func (*SVGFEGaussianBlurElement) SetOntouchcancel(ontouchcancel func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchcancel = $1", ontouchcancel)
}

func (*SVGFEGaussianBlurElement) Ontouchend() (ontouchend func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchend")
	return ontouchend
}

func (*SVGFEGaussianBlurElement) SetOntouchend(ontouchend func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchend = $1", ontouchend)
}

func (*SVGFEGaussianBlurElement) Ontouchmove() (ontouchmove func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchmove")
	return ontouchmove
}

func (*SVGFEGaussianBlurElement) SetOntouchmove(ontouchmove func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchmove = $1", ontouchmove)
}

func (*SVGFEGaussianBlurElement) Ontouchstart() (ontouchstart func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchstart")
	return ontouchstart
}

func (*SVGFEGaussianBlurElement) SetOntouchstart(ontouchstart func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchstart = $1", ontouchstart)
}

func (*SVGFEGaussianBlurElement) Onwebkitfullscreenchange() (onwebkitfullscreenchange func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenchange")
	return onwebkitfullscreenchange
}

func (*SVGFEGaussianBlurElement) SetOnwebkitfullscreenchange(onwebkitfullscreenchange func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenchange = $1", onwebkitfullscreenchange)
}

func (*SVGFEGaussianBlurElement) Onwebkitfullscreenerror() (onwebkitfullscreenerror func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenerror")
	return onwebkitfullscreenerror
}

func (*SVGFEGaussianBlurElement) SetOnwebkitfullscreenerror(onwebkitfullscreenerror func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenerror = $1", onwebkitfullscreenerror)
}

func (*SVGFEGaussianBlurElement) OuterHTML() (outerHTML string) {
	macro.Rewrite("$_.outerHTML")
	return outerHTML
}

func (*SVGFEGaussianBlurElement) SetOuterHTML(outerHTML string) {
	macro.Rewrite("$_.outerHTML = $1", outerHTML)
}

func (*SVGFEGaussianBlurElement) Prefix() (prefix string) {
	macro.Rewrite("$_.prefix")
	return prefix
}

func (*SVGFEGaussianBlurElement) ScrollHeight() (scrollHeight int) {
	macro.Rewrite("$_.scrollHeight")
	return scrollHeight
}

func (*SVGFEGaussianBlurElement) ScrollLeft() (scrollLeft int) {
	macro.Rewrite("$_.scrollLeft")
	return scrollLeft
}

func (*SVGFEGaussianBlurElement) SetScrollLeft(scrollLeft int) {
	macro.Rewrite("$_.scrollLeft = $1", scrollLeft)
}

func (*SVGFEGaussianBlurElement) ScrollTop() (scrollTop int) {
	macro.Rewrite("$_.scrollTop")
	return scrollTop
}

func (*SVGFEGaussianBlurElement) SetScrollTop(scrollTop int) {
	macro.Rewrite("$_.scrollTop = $1", scrollTop)
}

func (*SVGFEGaussianBlurElement) ScrollWidth() (scrollWidth int) {
	macro.Rewrite("$_.scrollWidth")
	return scrollWidth
}

func (*SVGFEGaussianBlurElement) TagName() (tagName string) {
	macro.Rewrite("$_.tagName")
	return tagName
}

func (*SVGFEGaussianBlurElement) Onpointercancel() (onpointercancel func(event.Event)) {
	macro.Rewrite("$_.onpointercancel")
	return onpointercancel
}

func (*SVGFEGaussianBlurElement) SetOnpointercancel(onpointercancel func(event.Event)) {
	macro.Rewrite("$_.onpointercancel = $1", onpointercancel)
}

func (*SVGFEGaussianBlurElement) Onpointerdown() (onpointerdown func(event.Event)) {
	macro.Rewrite("$_.onpointerdown")
	return onpointerdown
}

func (*SVGFEGaussianBlurElement) SetOnpointerdown(onpointerdown func(event.Event)) {
	macro.Rewrite("$_.onpointerdown = $1", onpointerdown)
}

func (*SVGFEGaussianBlurElement) Onpointerenter() (onpointerenter func(event.Event)) {
	macro.Rewrite("$_.onpointerenter")
	return onpointerenter
}

func (*SVGFEGaussianBlurElement) SetOnpointerenter(onpointerenter func(event.Event)) {
	macro.Rewrite("$_.onpointerenter = $1", onpointerenter)
}

func (*SVGFEGaussianBlurElement) Onpointerleave() (onpointerleave func(event.Event)) {
	macro.Rewrite("$_.onpointerleave")
	return onpointerleave
}

func (*SVGFEGaussianBlurElement) SetOnpointerleave(onpointerleave func(event.Event)) {
	macro.Rewrite("$_.onpointerleave = $1", onpointerleave)
}

func (*SVGFEGaussianBlurElement) Onpointermove() (onpointermove func(event.Event)) {
	macro.Rewrite("$_.onpointermove")
	return onpointermove
}

func (*SVGFEGaussianBlurElement) SetOnpointermove(onpointermove func(event.Event)) {
	macro.Rewrite("$_.onpointermove = $1", onpointermove)
}

func (*SVGFEGaussianBlurElement) Onpointerout() (onpointerout func(event.Event)) {
	macro.Rewrite("$_.onpointerout")
	return onpointerout
}

func (*SVGFEGaussianBlurElement) SetOnpointerout(onpointerout func(event.Event)) {
	macro.Rewrite("$_.onpointerout = $1", onpointerout)
}

func (*SVGFEGaussianBlurElement) Onpointerover() (onpointerover func(event.Event)) {
	macro.Rewrite("$_.onpointerover")
	return onpointerover
}

func (*SVGFEGaussianBlurElement) SetOnpointerover(onpointerover func(event.Event)) {
	macro.Rewrite("$_.onpointerover = $1", onpointerover)
}

func (*SVGFEGaussianBlurElement) Onpointerup() (onpointerup func(event.Event)) {
	macro.Rewrite("$_.onpointerup")
	return onpointerup
}

func (*SVGFEGaussianBlurElement) SetOnpointerup(onpointerup func(event.Event)) {
	macro.Rewrite("$_.onpointerup = $1", onpointerup)
}

func (*SVGFEGaussianBlurElement) Onwheel() (onwheel func(event.Event)) {
	macro.Rewrite("$_.onwheel")
	return onwheel
}

func (*SVGFEGaussianBlurElement) SetOnwheel(onwheel func(event.Event)) {
	macro.Rewrite("$_.onwheel = $1", onwheel)
}

func (*SVGFEGaussianBlurElement) ChildElementCount() (childElementCount uint) {
	macro.Rewrite("$_.childElementCount")
	return childElementCount
}

func (*SVGFEGaussianBlurElement) FirstElementChild() (firstElementChild element.Element) {
	macro.Rewrite("$_.firstElementChild")
	return firstElementChild
}

func (*SVGFEGaussianBlurElement) LastElementChild() (lastElementChild element.Element) {
	macro.Rewrite("$_.lastElementChild")
	return lastElementChild
}

func (*SVGFEGaussianBlurElement) NextElementSibling() (nextElementSibling element.Element) {
	macro.Rewrite("$_.nextElementSibling")
	return nextElementSibling
}

func (*SVGFEGaussianBlurElement) PreviousElementSibling() (previousElementSibling element.Element) {
	macro.Rewrite("$_.previousElementSibling")
	return previousElementSibling
}

func (*SVGFEGaussianBlurElement) Attributes() (attributes *NamedNodeMap) {
	macro.Rewrite("$_.attributes")
	return attributes
}

func (*SVGFEGaussianBlurElement) BaseURI() (baseURI string) {
	macro.Rewrite("$_.baseURI")
	return baseURI
}

func (*SVGFEGaussianBlurElement) ChildNodes() (childNodes *NodeList) {
	macro.Rewrite("$_.childNodes")
	return childNodes
}

func (*SVGFEGaussianBlurElement) FirstChild() (firstChild Node) {
	macro.Rewrite("$_.firstChild")
	return firstChild
}

func (*SVGFEGaussianBlurElement) LastChild() (lastChild Node) {
	macro.Rewrite("$_.lastChild")
	return lastChild
}

func (*SVGFEGaussianBlurElement) LocalName() (localName string) {
	macro.Rewrite("$_.localName")
	return localName
}

func (*SVGFEGaussianBlurElement) NamespaceURI() (namespaceURI string) {
	macro.Rewrite("$_.namespaceURI")
	return namespaceURI
}

func (*SVGFEGaussianBlurElement) NextSibling() (nextSibling Node) {
	macro.Rewrite("$_.nextSibling")
	return nextSibling
}

func (*SVGFEGaussianBlurElement) NodeName() (nodeName string) {
	macro.Rewrite("$_.nodeName")
	return nodeName
}

func (*SVGFEGaussianBlurElement) NodeType() (nodeType uint8) {
	macro.Rewrite("$_.nodeType")
	return nodeType
}

func (*SVGFEGaussianBlurElement) NodeValue() (nodeValue string) {
	macro.Rewrite("$_.nodeValue")
	return nodeValue
}

func (*SVGFEGaussianBlurElement) SetNodeValue(nodeValue string) {
	macro.Rewrite("$_.nodeValue = $1", nodeValue)
}

func (*SVGFEGaussianBlurElement) OwnerDocument() (ownerDocument document.Document) {
	macro.Rewrite("$_.ownerDocument")
	return ownerDocument
}

func (*SVGFEGaussianBlurElement) ParentElement() (parentElement html.HTMLElement) {
	macro.Rewrite("$_.parentElement")
	return parentElement
}

func (*SVGFEGaussianBlurElement) ParentNode() (parentNode Node) {
	macro.Rewrite("$_.parentNode")
	return parentNode
}

func (*SVGFEGaussianBlurElement) PreviousSibling() (previousSibling Node) {
	macro.Rewrite("$_.previousSibling")
	return previousSibling
}

func (*SVGFEGaussianBlurElement) TextContent() (textContent string) {
	macro.Rewrite("$_.textContent")
	return textContent
}

func (*SVGFEGaussianBlurElement) SetTextContent(textContent string) {
	macro.Rewrite("$_.textContent = $1", textContent)
}
