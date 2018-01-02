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

var _ SVGTextPositioningElement = (*SVGTSpanElement)(nil)
var _ SVGTextContentElement = (*SVGTSpanElement)(nil)
var _ SVGGraphicsElement = (*SVGTSpanElement)(nil)
var _ SVGTests = (*SVGTSpanElement)(nil)
var _ SVGElement = (*SVGTSpanElement)(nil)
var _ element.Element = (*SVGTSpanElement)(nil)
var _ utils.GlobalEventHandlers = (*SVGTSpanElement)(nil)
var _ element.ElementTraversal = (*SVGTSpanElement)(nil)
var _ NodeSelector = (*SVGTSpanElement)(nil)
var _ ChildNode = (*SVGTSpanElement)(nil)
var _ Node = (*SVGTSpanElement)(nil)
var _ event.EventTarget = (*SVGTSpanElement)(nil)

type SVGTSpanElement struct {
}

func (*SVGTSpanElement) GetCharNumAtPosition(point *SVGPoint) (i int) {
	macro.Rewrite("$_.getCharNumAtPosition($1)", point)
	return i
}

func (*SVGTSpanElement) GetComputedTextLength() (f float32) {
	macro.Rewrite("$_.getComputedTextLength()")
	return f
}

func (*SVGTSpanElement) GetEndPositionOfChar(charnum uint) (s *SVGPoint) {
	macro.Rewrite("$_.getEndPositionOfChar($1)", charnum)
	return s
}

func (*SVGTSpanElement) GetExtentOfChar(charnum uint) (s *SVGRect) {
	macro.Rewrite("$_.getExtentOfChar($1)", charnum)
	return s
}

func (*SVGTSpanElement) GetNumberOfChars() (i int) {
	macro.Rewrite("$_.getNumberOfChars()")
	return i
}

func (*SVGTSpanElement) GetRotationOfChar(charnum uint) (f float32) {
	macro.Rewrite("$_.getRotationOfChar($1)", charnum)
	return f
}

func (*SVGTSpanElement) GetStartPositionOfChar(charnum uint) (s *SVGPoint) {
	macro.Rewrite("$_.getStartPositionOfChar($1)", charnum)
	return s
}

func (*SVGTSpanElement) GetSubStringLength(charnum uint, nchars uint) (f float32) {
	macro.Rewrite("$_.getSubStringLength($1, $2)", charnum, nchars)
	return f
}

func (*SVGTSpanElement) SelectSubString(charnum uint, nchars uint) {
	macro.Rewrite("$_.selectSubString($1, $2)", charnum, nchars)
}

func (*SVGTSpanElement) GetBBox() (s *SVGRect) {
	macro.Rewrite("$_.getBBox()")
	return s
}

func (*SVGTSpanElement) GetCTM() (s *SVGMatrix) {
	macro.Rewrite("$_.getCTM()")
	return s
}

func (*SVGTSpanElement) GetScreenCTM() (s *SVGMatrix) {
	macro.Rewrite("$_.getScreenCTM()")
	return s
}

func (*SVGTSpanElement) GetTransformToElement(element SVGElement) (s *SVGMatrix) {
	macro.Rewrite("$_.getTransformToElement($1)", element)
	return s
}

func (*SVGTSpanElement) HasExtension(extension string) (b bool) {
	macro.Rewrite("$_.hasExtension($1)", extension)
	return b
}

func (*SVGTSpanElement) GetAttribute(qualifiedName string) (s string) {
	macro.Rewrite("$_.getAttribute($1)", qualifiedName)
	return s
}

func (*SVGTSpanElement) GetAttributeNode(name string) (w *Attr) {
	macro.Rewrite("$_.getAttributeNode($1)", name)
	return w
}

func (*SVGTSpanElement) GetAttributeNodeNS(namespaceURI string, localName string) (w *Attr) {
	macro.Rewrite("$_.getAttributeNodeNS($1, $2)", namespaceURI, localName)
	return w
}

func (*SVGTSpanElement) GetAttributeNS(namespaceURI string, localName string) (s string) {
	macro.Rewrite("$_.getAttributeNS($1, $2)", namespaceURI, localName)
	return s
}

func (*SVGTSpanElement) GetBoundingClientRect() (c *ClientRect) {
	macro.Rewrite("$_.getBoundingClientRect()")
	return c
}

func (*SVGTSpanElement) GetClientRects() (c *ClientRectList) {
	macro.Rewrite("$_.getClientRects()")
	return c
}

func (*SVGTSpanElement) GetElementsByTagName(name string) (w *NodeList) {
	macro.Rewrite("$_.getElementsByTagName($1)", name)
	return w
}

func (*SVGTSpanElement) GetElementsByTagNameNS(namespaceURI string, localName string) (w *NodeList) {
	macro.Rewrite("$_.getElementsByTagNameNS($1, $2)", namespaceURI, localName)
	return w
}

func (*SVGTSpanElement) HasAttribute(name string) (b bool) {
	macro.Rewrite("$_.hasAttribute($1)", name)
	return b
}

func (*SVGTSpanElement) HasAttributeNS(namespaceURI string, localName string) (b bool) {
	macro.Rewrite("$_.hasAttributeNS($1, $2)", namespaceURI, localName)
	return b
}

func (*SVGTSpanElement) MsGetRegionContent() (w *ms.MSRangeCollection) {
	macro.Rewrite("$_.msGetRegionContent()")
	return w
}

func (*SVGTSpanElement) MsGetUntransformedBounds() (c *ClientRect) {
	macro.Rewrite("$_.msGetUntransformedBounds()")
	return c
}

func (*SVGTSpanElement) MsMatchesSelector(selectors string) (b bool) {
	macro.Rewrite("$_.msMatchesSelector($1)", selectors)
	return b
}

func (*SVGTSpanElement) MsReleasePointerCapture(pointerId int) {
	macro.Rewrite("$_.msReleasePointerCapture($1)", pointerId)
}

func (*SVGTSpanElement) MsSetPointerCapture(pointerId int) {
	macro.Rewrite("$_.msSetPointerCapture($1)", pointerId)
}

func (*SVGTSpanElement) MsZoomTo(args *ms.MsZoomToOptions) {
	macro.Rewrite("$_.msZoomTo($1)", args)
}

func (*SVGTSpanElement) ReleasePointerCapture(pointerId int) {
	macro.Rewrite("$_.releasePointerCapture($1)", pointerId)
}

func (*SVGTSpanElement) RemoveAttribute(qualifiedName string) {
	macro.Rewrite("$_.removeAttribute($1)", qualifiedName)
}

func (*SVGTSpanElement) RemoveAttributeNode(oldAttr *Attr) (w *Attr) {
	macro.Rewrite("$_.removeAttributeNode($1)", oldAttr)
	return w
}

func (*SVGTSpanElement) RemoveAttributeNS(namespaceURI string, localName string) {
	macro.Rewrite("$_.removeAttributeNS($1, $2)", namespaceURI, localName)
}

func (*SVGTSpanElement) RequestFullscreen() {
	macro.Rewrite("$_.requestFullscreen()")
}

func (*SVGTSpanElement) RequestPointerLock() {
	macro.Rewrite("$_.requestPointerLock()")
}

func (*SVGTSpanElement) SetAttribute(qualifiedName string, value string) {
	macro.Rewrite("$_.setAttribute($1, $2)", qualifiedName, value)
}

func (*SVGTSpanElement) SetAttributeNode(newAttr *Attr) (w *Attr) {
	macro.Rewrite("$_.setAttributeNode($1)", newAttr)
	return w
}

func (*SVGTSpanElement) SetAttributeNodeNS(newAttr *Attr) (w *Attr) {
	macro.Rewrite("$_.setAttributeNodeNS($1)", newAttr)
	return w
}

func (*SVGTSpanElement) SetAttributeNS(namespaceURI string, qualifiedName string, value string) {
	macro.Rewrite("$_.setAttributeNS($1, $2, $3)", namespaceURI, qualifiedName, value)
}

func (*SVGTSpanElement) SetPointerCapture(pointerId int) {
	macro.Rewrite("$_.setPointerCapture($1)", pointerId)
}

func (*SVGTSpanElement) WebkitMatchesSelector(selectors string) (b bool) {
	macro.Rewrite("$_.webkitMatchesSelector($1)", selectors)
	return b
}

func (*SVGTSpanElement) WebkitRequestFullscreen() {
	macro.Rewrite("$_.webkitRequestFullscreen()")
}

func (*SVGTSpanElement) WebkitRequestFullScreen() {
	macro.Rewrite("$_.webkitRequestFullScreen()")
}

func (*SVGTSpanElement) QuerySelector(selectors string) (w element.Element) {
	macro.Rewrite("$_.querySelector($1)", selectors)
	return w
}

func (*SVGTSpanElement) QuerySelectorAll(selectors string) (w *NodeList) {
	macro.Rewrite("$_.querySelectorAll($1)", selectors)
	return w
}

func (*SVGTSpanElement) AppendChild(newChild Node) (w Node) {
	macro.Rewrite("$_.appendChild($1)", newChild)
	return w
}

func (*SVGTSpanElement) CloneNode(deep *bool) (w Node) {
	macro.Rewrite("$_.cloneNode($1)", deep)
	return w
}

func (*SVGTSpanElement) CompareDocumentPosition(other Node) (u uint8) {
	macro.Rewrite("$_.compareDocumentPosition($1)", other)
	return u
}

func (*SVGTSpanElement) Contains(child Node) (b bool) {
	macro.Rewrite("$_.contains($1)", child)
	return b
}

func (*SVGTSpanElement) HasAttributes() (b bool) {
	macro.Rewrite("$_.hasAttributes()")
	return b
}

func (*SVGTSpanElement) HasChildNodes() (b bool) {
	macro.Rewrite("$_.hasChildNodes()")
	return b
}

func (*SVGTSpanElement) InsertBefore(newChild Node, refChild *Node) (w Node) {
	macro.Rewrite("$_.insertBefore($1, $2)", newChild, refChild)
	return w
}

func (*SVGTSpanElement) IsDefaultNamespace(namespaceURI string) (b bool) {
	macro.Rewrite("$_.isDefaultNamespace($1)", namespaceURI)
	return b
}

func (*SVGTSpanElement) IsEqualNode(arg Node) (b bool) {
	macro.Rewrite("$_.isEqualNode($1)", arg)
	return b
}

func (*SVGTSpanElement) IsSameNode(other Node) (b bool) {
	macro.Rewrite("$_.isSameNode($1)", other)
	return b
}

func (*SVGTSpanElement) LookupNamespaceURI(prefix string) (s string) {
	macro.Rewrite("$_.lookupNamespaceURI($1)", prefix)
	return s
}

func (*SVGTSpanElement) LookupPrefix(namespaceURI string) (s string) {
	macro.Rewrite("$_.lookupPrefix($1)", namespaceURI)
	return s
}

func (*SVGTSpanElement) Normalize() {
	macro.Rewrite("$_.normalize()")
}

func (*SVGTSpanElement) RemoveChild(oldChild Node) (w Node) {
	macro.Rewrite("$_.removeChild($1)", oldChild)
	return w
}

func (*SVGTSpanElement) ReplaceChild(newChild Node, oldChild Node) (w Node) {
	macro.Rewrite("$_.replaceChild($1, $2)", newChild, oldChild)
	return w
}

func (*SVGTSpanElement) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*SVGTSpanElement) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*SVGTSpanElement) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*SVGTSpanElement) Dx() (dx *SVGAnimatedLengthList) {
	macro.Rewrite("$_.dx")
	return dx
}

func (*SVGTSpanElement) Dy() (dy *SVGAnimatedLengthList) {
	macro.Rewrite("$_.dy")
	return dy
}

func (*SVGTSpanElement) Rotate() (rotate *SVGAnimatedNumberList) {
	macro.Rewrite("$_.rotate")
	return rotate
}

func (*SVGTSpanElement) X() (x *SVGAnimatedLengthList) {
	macro.Rewrite("$_.x")
	return x
}

func (*SVGTSpanElement) Y() (y *SVGAnimatedLengthList) {
	macro.Rewrite("$_.y")
	return y
}

func (*SVGTSpanElement) LengthAdjust() (lengthAdjust *SVGAnimatedEnumeration) {
	macro.Rewrite("$_.lengthAdjust")
	return lengthAdjust
}

func (*SVGTSpanElement) TextLength() (textLength *SVGAnimatedLength) {
	macro.Rewrite("$_.textLength")
	return textLength
}

func (*SVGTSpanElement) FarthestViewportElement() (farthestViewportElement SVGElement) {
	macro.Rewrite("$_.farthestViewportElement")
	return farthestViewportElement
}

func (*SVGTSpanElement) NearestViewportElement() (nearestViewportElement SVGElement) {
	macro.Rewrite("$_.nearestViewportElement")
	return nearestViewportElement
}

func (*SVGTSpanElement) Transform() (transform *SVGAnimatedTransformList) {
	macro.Rewrite("$_.transform")
	return transform
}

func (*SVGTSpanElement) RequiredExtensions() (requiredExtensions *SVGStringList) {
	macro.Rewrite("$_.requiredExtensions")
	return requiredExtensions
}

func (*SVGTSpanElement) RequiredFeatures() (requiredFeatures *SVGStringList) {
	macro.Rewrite("$_.requiredFeatures")
	return requiredFeatures
}

func (*SVGTSpanElement) SystemLanguage() (systemLanguage *SVGStringList) {
	macro.Rewrite("$_.systemLanguage")
	return systemLanguage
}

func (*SVGTSpanElement) Dataset() (dataset *DOMStringMap) {
	macro.Rewrite("$_.dataset")
	return dataset
}

func (*SVGTSpanElement) OwnerSVGElement() (ownerSVGElement *SVGSVGElement) {
	macro.Rewrite("$_.ownerSVGElement")
	return ownerSVGElement
}

func (*SVGTSpanElement) ViewportElement() (viewportElement SVGElement) {
	macro.Rewrite("$_.viewportElement")
	return viewportElement
}

func (*SVGTSpanElement) Xmlbase() (xmlbase string) {
	macro.Rewrite("$_.xmlbase")
	return xmlbase
}

func (*SVGTSpanElement) SetXmlbase(xmlbase string) {
	macro.Rewrite("$_.xmlbase = $1", xmlbase)
}

func (*SVGTSpanElement) ClassList() (classList DOMTokenList) {
	macro.Rewrite("$_.classList")
	return classList
}

func (*SVGTSpanElement) ClassName() (className string) {
	macro.Rewrite("$_.className")
	return className
}

func (*SVGTSpanElement) SetClassName(className string) {
	macro.Rewrite("$_.className = $1", className)
}

func (*SVGTSpanElement) ClientHeight() (clientHeight int) {
	macro.Rewrite("$_.clientHeight")
	return clientHeight
}

func (*SVGTSpanElement) ClientLeft() (clientLeft int) {
	macro.Rewrite("$_.clientLeft")
	return clientLeft
}

func (*SVGTSpanElement) ClientTop() (clientTop int) {
	macro.Rewrite("$_.clientTop")
	return clientTop
}

func (*SVGTSpanElement) ClientWidth() (clientWidth int) {
	macro.Rewrite("$_.clientWidth")
	return clientWidth
}

func (*SVGTSpanElement) ID() (id string) {
	macro.Rewrite("$_.id")
	return id
}

func (*SVGTSpanElement) SetID(id string) {
	macro.Rewrite("$_.id = $1", id)
}

func (*SVGTSpanElement) InnerHTML() (innerHTML string) {
	macro.Rewrite("$_.innerHTML")
	return innerHTML
}

func (*SVGTSpanElement) SetInnerHTML(innerHTML string) {
	macro.Rewrite("$_.innerHTML = $1", innerHTML)
}

func (*SVGTSpanElement) MsContentZoomFactor() (msContentZoomFactor float32) {
	macro.Rewrite("$_.msContentZoomFactor")
	return msContentZoomFactor
}

func (*SVGTSpanElement) SetMsContentZoomFactor(msContentZoomFactor float32) {
	macro.Rewrite("$_.msContentZoomFactor = $1", msContentZoomFactor)
}

func (*SVGTSpanElement) MsRegionOverflow() (msRegionOverflow string) {
	macro.Rewrite("$_.msRegionOverflow")
	return msRegionOverflow
}

func (*SVGTSpanElement) Onariarequest() (onariarequest func(event.Event)) {
	macro.Rewrite("$_.onariarequest")
	return onariarequest
}

func (*SVGTSpanElement) SetOnariarequest(onariarequest func(event.Event)) {
	macro.Rewrite("$_.onariarequest = $1", onariarequest)
}

func (*SVGTSpanElement) Oncommand() (oncommand func(event.Event)) {
	macro.Rewrite("$_.oncommand")
	return oncommand
}

func (*SVGTSpanElement) SetOncommand(oncommand func(event.Event)) {
	macro.Rewrite("$_.oncommand = $1", oncommand)
}

func (*SVGTSpanElement) Ongotpointercapture() (ongotpointercapture func(*utils.PointerEvent)) {
	macro.Rewrite("$_.ongotpointercapture")
	return ongotpointercapture
}

func (*SVGTSpanElement) SetOngotpointercapture(ongotpointercapture func(*utils.PointerEvent)) {
	macro.Rewrite("$_.ongotpointercapture = $1", ongotpointercapture)
}

func (*SVGTSpanElement) Onlostpointercapture() (onlostpointercapture func(*utils.PointerEvent)) {
	macro.Rewrite("$_.onlostpointercapture")
	return onlostpointercapture
}

func (*SVGTSpanElement) SetOnlostpointercapture(onlostpointercapture func(*utils.PointerEvent)) {
	macro.Rewrite("$_.onlostpointercapture = $1", onlostpointercapture)
}

func (*SVGTSpanElement) Onmsgesturechange() (onmsgesturechange func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturechange")
	return onmsgesturechange
}

func (*SVGTSpanElement) SetOnmsgesturechange(onmsgesturechange func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturechange = $1", onmsgesturechange)
}

func (*SVGTSpanElement) Onmsgesturedoubletap() (onmsgesturedoubletap func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturedoubletap")
	return onmsgesturedoubletap
}

func (*SVGTSpanElement) SetOnmsgesturedoubletap(onmsgesturedoubletap func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturedoubletap = $1", onmsgesturedoubletap)
}

func (*SVGTSpanElement) Onmsgestureend() (onmsgestureend func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgestureend")
	return onmsgestureend
}

func (*SVGTSpanElement) SetOnmsgestureend(onmsgestureend func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgestureend = $1", onmsgestureend)
}

func (*SVGTSpanElement) Onmsgesturehold() (onmsgesturehold func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturehold")
	return onmsgesturehold
}

func (*SVGTSpanElement) SetOnmsgesturehold(onmsgesturehold func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturehold = $1", onmsgesturehold)
}

func (*SVGTSpanElement) Onmsgesturestart() (onmsgesturestart func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturestart")
	return onmsgesturestart
}

func (*SVGTSpanElement) SetOnmsgesturestart(onmsgesturestart func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturestart = $1", onmsgesturestart)
}

func (*SVGTSpanElement) Onmsgesturetap() (onmsgesturetap func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturetap")
	return onmsgesturetap
}

func (*SVGTSpanElement) SetOnmsgesturetap(onmsgesturetap func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturetap = $1", onmsgesturetap)
}

func (*SVGTSpanElement) Onmsgotpointercapture() (onmsgotpointercapture func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmsgotpointercapture")
	return onmsgotpointercapture
}

func (*SVGTSpanElement) SetOnmsgotpointercapture(onmsgotpointercapture func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmsgotpointercapture = $1", onmsgotpointercapture)
}

func (*SVGTSpanElement) Onmsinertiastart() (onmsinertiastart func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsinertiastart")
	return onmsinertiastart
}

func (*SVGTSpanElement) SetOnmsinertiastart(onmsinertiastart func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsinertiastart = $1", onmsinertiastart)
}

func (*SVGTSpanElement) Onmslostpointercapture() (onmslostpointercapture func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmslostpointercapture")
	return onmslostpointercapture
}

func (*SVGTSpanElement) SetOnmslostpointercapture(onmslostpointercapture func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmslostpointercapture = $1", onmslostpointercapture)
}

func (*SVGTSpanElement) Onmspointercancel() (onmspointercancel func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointercancel")
	return onmspointercancel
}

func (*SVGTSpanElement) SetOnmspointercancel(onmspointercancel func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointercancel = $1", onmspointercancel)
}

func (*SVGTSpanElement) Onmspointerdown() (onmspointerdown func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerdown")
	return onmspointerdown
}

func (*SVGTSpanElement) SetOnmspointerdown(onmspointerdown func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerdown = $1", onmspointerdown)
}

func (*SVGTSpanElement) Onmspointerenter() (onmspointerenter func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerenter")
	return onmspointerenter
}

func (*SVGTSpanElement) SetOnmspointerenter(onmspointerenter func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerenter = $1", onmspointerenter)
}

func (*SVGTSpanElement) Onmspointerleave() (onmspointerleave func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerleave")
	return onmspointerleave
}

func (*SVGTSpanElement) SetOnmspointerleave(onmspointerleave func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerleave = $1", onmspointerleave)
}

func (*SVGTSpanElement) Onmspointermove() (onmspointermove func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointermove")
	return onmspointermove
}

func (*SVGTSpanElement) SetOnmspointermove(onmspointermove func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointermove = $1", onmspointermove)
}

func (*SVGTSpanElement) Onmspointerout() (onmspointerout func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerout")
	return onmspointerout
}

func (*SVGTSpanElement) SetOnmspointerout(onmspointerout func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerout = $1", onmspointerout)
}

func (*SVGTSpanElement) Onmspointerover() (onmspointerover func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerover")
	return onmspointerover
}

func (*SVGTSpanElement) SetOnmspointerover(onmspointerover func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerover = $1", onmspointerover)
}

func (*SVGTSpanElement) Onmspointerup() (onmspointerup func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerup")
	return onmspointerup
}

func (*SVGTSpanElement) SetOnmspointerup(onmspointerup func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerup = $1", onmspointerup)
}

func (*SVGTSpanElement) Ontouchcancel() (ontouchcancel func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchcancel")
	return ontouchcancel
}

func (*SVGTSpanElement) SetOntouchcancel(ontouchcancel func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchcancel = $1", ontouchcancel)
}

func (*SVGTSpanElement) Ontouchend() (ontouchend func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchend")
	return ontouchend
}

func (*SVGTSpanElement) SetOntouchend(ontouchend func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchend = $1", ontouchend)
}

func (*SVGTSpanElement) Ontouchmove() (ontouchmove func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchmove")
	return ontouchmove
}

func (*SVGTSpanElement) SetOntouchmove(ontouchmove func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchmove = $1", ontouchmove)
}

func (*SVGTSpanElement) Ontouchstart() (ontouchstart func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchstart")
	return ontouchstart
}

func (*SVGTSpanElement) SetOntouchstart(ontouchstart func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchstart = $1", ontouchstart)
}

func (*SVGTSpanElement) Onwebkitfullscreenchange() (onwebkitfullscreenchange func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenchange")
	return onwebkitfullscreenchange
}

func (*SVGTSpanElement) SetOnwebkitfullscreenchange(onwebkitfullscreenchange func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenchange = $1", onwebkitfullscreenchange)
}

func (*SVGTSpanElement) Onwebkitfullscreenerror() (onwebkitfullscreenerror func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenerror")
	return onwebkitfullscreenerror
}

func (*SVGTSpanElement) SetOnwebkitfullscreenerror(onwebkitfullscreenerror func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenerror = $1", onwebkitfullscreenerror)
}

func (*SVGTSpanElement) OuterHTML() (outerHTML string) {
	macro.Rewrite("$_.outerHTML")
	return outerHTML
}

func (*SVGTSpanElement) SetOuterHTML(outerHTML string) {
	macro.Rewrite("$_.outerHTML = $1", outerHTML)
}

func (*SVGTSpanElement) Prefix() (prefix string) {
	macro.Rewrite("$_.prefix")
	return prefix
}

func (*SVGTSpanElement) ScrollHeight() (scrollHeight int) {
	macro.Rewrite("$_.scrollHeight")
	return scrollHeight
}

func (*SVGTSpanElement) ScrollLeft() (scrollLeft int) {
	macro.Rewrite("$_.scrollLeft")
	return scrollLeft
}

func (*SVGTSpanElement) SetScrollLeft(scrollLeft int) {
	macro.Rewrite("$_.scrollLeft = $1", scrollLeft)
}

func (*SVGTSpanElement) ScrollTop() (scrollTop int) {
	macro.Rewrite("$_.scrollTop")
	return scrollTop
}

func (*SVGTSpanElement) SetScrollTop(scrollTop int) {
	macro.Rewrite("$_.scrollTop = $1", scrollTop)
}

func (*SVGTSpanElement) ScrollWidth() (scrollWidth int) {
	macro.Rewrite("$_.scrollWidth")
	return scrollWidth
}

func (*SVGTSpanElement) TagName() (tagName string) {
	macro.Rewrite("$_.tagName")
	return tagName
}

func (*SVGTSpanElement) Onpointercancel() (onpointercancel func(event.Event)) {
	macro.Rewrite("$_.onpointercancel")
	return onpointercancel
}

func (*SVGTSpanElement) SetOnpointercancel(onpointercancel func(event.Event)) {
	macro.Rewrite("$_.onpointercancel = $1", onpointercancel)
}

func (*SVGTSpanElement) Onpointerdown() (onpointerdown func(event.Event)) {
	macro.Rewrite("$_.onpointerdown")
	return onpointerdown
}

func (*SVGTSpanElement) SetOnpointerdown(onpointerdown func(event.Event)) {
	macro.Rewrite("$_.onpointerdown = $1", onpointerdown)
}

func (*SVGTSpanElement) Onpointerenter() (onpointerenter func(event.Event)) {
	macro.Rewrite("$_.onpointerenter")
	return onpointerenter
}

func (*SVGTSpanElement) SetOnpointerenter(onpointerenter func(event.Event)) {
	macro.Rewrite("$_.onpointerenter = $1", onpointerenter)
}

func (*SVGTSpanElement) Onpointerleave() (onpointerleave func(event.Event)) {
	macro.Rewrite("$_.onpointerleave")
	return onpointerleave
}

func (*SVGTSpanElement) SetOnpointerleave(onpointerleave func(event.Event)) {
	macro.Rewrite("$_.onpointerleave = $1", onpointerleave)
}

func (*SVGTSpanElement) Onpointermove() (onpointermove func(event.Event)) {
	macro.Rewrite("$_.onpointermove")
	return onpointermove
}

func (*SVGTSpanElement) SetOnpointermove(onpointermove func(event.Event)) {
	macro.Rewrite("$_.onpointermove = $1", onpointermove)
}

func (*SVGTSpanElement) Onpointerout() (onpointerout func(event.Event)) {
	macro.Rewrite("$_.onpointerout")
	return onpointerout
}

func (*SVGTSpanElement) SetOnpointerout(onpointerout func(event.Event)) {
	macro.Rewrite("$_.onpointerout = $1", onpointerout)
}

func (*SVGTSpanElement) Onpointerover() (onpointerover func(event.Event)) {
	macro.Rewrite("$_.onpointerover")
	return onpointerover
}

func (*SVGTSpanElement) SetOnpointerover(onpointerover func(event.Event)) {
	macro.Rewrite("$_.onpointerover = $1", onpointerover)
}

func (*SVGTSpanElement) Onpointerup() (onpointerup func(event.Event)) {
	macro.Rewrite("$_.onpointerup")
	return onpointerup
}

func (*SVGTSpanElement) SetOnpointerup(onpointerup func(event.Event)) {
	macro.Rewrite("$_.onpointerup = $1", onpointerup)
}

func (*SVGTSpanElement) Onwheel() (onwheel func(event.Event)) {
	macro.Rewrite("$_.onwheel")
	return onwheel
}

func (*SVGTSpanElement) SetOnwheel(onwheel func(event.Event)) {
	macro.Rewrite("$_.onwheel = $1", onwheel)
}

func (*SVGTSpanElement) ChildElementCount() (childElementCount uint) {
	macro.Rewrite("$_.childElementCount")
	return childElementCount
}

func (*SVGTSpanElement) FirstElementChild() (firstElementChild element.Element) {
	macro.Rewrite("$_.firstElementChild")
	return firstElementChild
}

func (*SVGTSpanElement) LastElementChild() (lastElementChild element.Element) {
	macro.Rewrite("$_.lastElementChild")
	return lastElementChild
}

func (*SVGTSpanElement) NextElementSibling() (nextElementSibling element.Element) {
	macro.Rewrite("$_.nextElementSibling")
	return nextElementSibling
}

func (*SVGTSpanElement) PreviousElementSibling() (previousElementSibling element.Element) {
	macro.Rewrite("$_.previousElementSibling")
	return previousElementSibling
}

func (*SVGTSpanElement) Attributes() (attributes *NamedNodeMap) {
	macro.Rewrite("$_.attributes")
	return attributes
}

func (*SVGTSpanElement) BaseURI() (baseURI string) {
	macro.Rewrite("$_.baseURI")
	return baseURI
}

func (*SVGTSpanElement) ChildNodes() (childNodes *NodeList) {
	macro.Rewrite("$_.childNodes")
	return childNodes
}

func (*SVGTSpanElement) FirstChild() (firstChild Node) {
	macro.Rewrite("$_.firstChild")
	return firstChild
}

func (*SVGTSpanElement) LastChild() (lastChild Node) {
	macro.Rewrite("$_.lastChild")
	return lastChild
}

func (*SVGTSpanElement) LocalName() (localName string) {
	macro.Rewrite("$_.localName")
	return localName
}

func (*SVGTSpanElement) NamespaceURI() (namespaceURI string) {
	macro.Rewrite("$_.namespaceURI")
	return namespaceURI
}

func (*SVGTSpanElement) NextSibling() (nextSibling Node) {
	macro.Rewrite("$_.nextSibling")
	return nextSibling
}

func (*SVGTSpanElement) NodeName() (nodeName string) {
	macro.Rewrite("$_.nodeName")
	return nodeName
}

func (*SVGTSpanElement) NodeType() (nodeType uint8) {
	macro.Rewrite("$_.nodeType")
	return nodeType
}

func (*SVGTSpanElement) NodeValue() (nodeValue string) {
	macro.Rewrite("$_.nodeValue")
	return nodeValue
}

func (*SVGTSpanElement) SetNodeValue(nodeValue string) {
	macro.Rewrite("$_.nodeValue = $1", nodeValue)
}

func (*SVGTSpanElement) OwnerDocument() (ownerDocument document.Document) {
	macro.Rewrite("$_.ownerDocument")
	return ownerDocument
}

func (*SVGTSpanElement) ParentElement() (parentElement html.HTMLElement) {
	macro.Rewrite("$_.parentElement")
	return parentElement
}

func (*SVGTSpanElement) ParentNode() (parentNode Node) {
	macro.Rewrite("$_.parentNode")
	return parentNode
}

func (*SVGTSpanElement) PreviousSibling() (previousSibling Node) {
	macro.Rewrite("$_.previousSibling")
	return previousSibling
}

func (*SVGTSpanElement) TextContent() (textContent string) {
	macro.Rewrite("$_.textContent")
	return textContent
}

func (*SVGTSpanElement) SetTextContent(textContent string) {
	macro.Rewrite("$_.textContent = $1", textContent)
}
