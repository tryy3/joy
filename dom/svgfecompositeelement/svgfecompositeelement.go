package svgfecompositeelement

import (
	"github.com/matthewmueller/joy/dom/dom"
	"github.com/matthewmueller/joy/dom/ms"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ window.SVGElement = (*SVGFECompositeElement)(nil)
var _ window.Element = (*SVGFECompositeElement)(nil)
var _ window.GlobalEventHandlers = (*SVGFECompositeElement)(nil)
var _ window.ElementTraversal = (*SVGFECompositeElement)(nil)
var _ window.NodeSelector = (*SVGFECompositeElement)(nil)
var _ dom.ChildNode = (*SVGFECompositeElement)(nil)
var _ window.Node = (*SVGFECompositeElement)(nil)
var _ window.EventTarget = (*SVGFECompositeElement)(nil)

// SVGFECompositeElement struct
// js:"SVGFECompositeElement,omit"
type SVGFECompositeElement struct {
}

// GetAttribute fn
// js:"getAttribute"
func (*SVGFECompositeElement) GetAttribute(qualifiedName string) (s string) {
	macro.Rewrite("$_.getAttribute($1)", qualifiedName)
	return s
}

// GetAttributeNode fn
// js:"getAttributeNode"
func (*SVGFECompositeElement) GetAttributeNode(name string) (w *window.Attr) {
	macro.Rewrite("$_.getAttributeNode($1)", name)
	return w
}

// GetAttributeNodeNS fn
// js:"getAttributeNodeNS"
func (*SVGFECompositeElement) GetAttributeNodeNS(namespaceURI string, localName string) (w *window.Attr) {
	macro.Rewrite("$_.getAttributeNodeNS($1, $2)", namespaceURI, localName)
	return w
}

// GetAttributeNS fn
// js:"getAttributeNS"
func (*SVGFECompositeElement) GetAttributeNS(namespaceURI string, localName string) (s string) {
	macro.Rewrite("$_.getAttributeNS($1, $2)", namespaceURI, localName)
	return s
}

// GetBoundingClientRect fn
// js:"getBoundingClientRect"
func (*SVGFECompositeElement) GetBoundingClientRect() (d *dom.ClientRect) {
	macro.Rewrite("$_.getBoundingClientRect()")
	return d
}

// GetClientRects fn
// js:"getClientRects"
func (*SVGFECompositeElement) GetClientRects() (d *dom.ClientRectList) {
	macro.Rewrite("$_.getClientRects()")
	return d
}

// GetElementsByTagName fn
// js:"getElementsByTagName"
func (*SVGFECompositeElement) GetElementsByTagName(name string) (w *window.NodeList) {
	macro.Rewrite("$_.getElementsByTagName($1)", name)
	return w
}

// GetElementsByTagNameNS fn
// js:"getElementsByTagNameNS"
func (*SVGFECompositeElement) GetElementsByTagNameNS(namespaceURI string, localName string) (w *window.NodeList) {
	macro.Rewrite("$_.getElementsByTagNameNS($1, $2)", namespaceURI, localName)
	return w
}

// HasAttribute fn
// js:"hasAttribute"
func (*SVGFECompositeElement) HasAttribute(name string) (b bool) {
	macro.Rewrite("$_.hasAttribute($1)", name)
	return b
}

// HasAttributeNS fn
// js:"hasAttributeNS"
func (*SVGFECompositeElement) HasAttributeNS(namespaceURI string, localName string) (b bool) {
	macro.Rewrite("$_.hasAttributeNS($1, $2)", namespaceURI, localName)
	return b
}

// MsGetRegionContent fn
// js:"msGetRegionContent"
func (*SVGFECompositeElement) MsGetRegionContent() (w *window.MSRangeCollection) {
	macro.Rewrite("$_.msGetRegionContent()")
	return w
}

// MsGetUntransformedBounds fn
// js:"msGetUntransformedBounds"
func (*SVGFECompositeElement) MsGetUntransformedBounds() (d *dom.ClientRect) {
	macro.Rewrite("$_.msGetUntransformedBounds()")
	return d
}

// MsMatchesSelector fn
// js:"msMatchesSelector"
func (*SVGFECompositeElement) MsMatchesSelector(selectors string) (b bool) {
	macro.Rewrite("$_.msMatchesSelector($1)", selectors)
	return b
}

// MsReleasePointerCapture fn
// js:"msReleasePointerCapture"
func (*SVGFECompositeElement) MsReleasePointerCapture(pointerId int) {
	macro.Rewrite("$_.msReleasePointerCapture($1)", pointerId)
}

// MsSetPointerCapture fn
// js:"msSetPointerCapture"
func (*SVGFECompositeElement) MsSetPointerCapture(pointerId int) {
	macro.Rewrite("$_.msSetPointerCapture($1)", pointerId)
}

// MsZoomTo fn
// js:"msZoomTo"
func (*SVGFECompositeElement) MsZoomTo(args *ms.MsZoomToOptions) {
	macro.Rewrite("$_.msZoomTo($1)", args)
}

// ReleasePointerCapture fn
// js:"releasePointerCapture"
func (*SVGFECompositeElement) ReleasePointerCapture(pointerId int) {
	macro.Rewrite("$_.releasePointerCapture($1)", pointerId)
}

// RemoveAttribute fn
// js:"removeAttribute"
func (*SVGFECompositeElement) RemoveAttribute(qualifiedName string) {
	macro.Rewrite("$_.removeAttribute($1)", qualifiedName)
}

// RemoveAttributeNode fn
// js:"removeAttributeNode"
func (*SVGFECompositeElement) RemoveAttributeNode(oldAttr *window.Attr) (w *window.Attr) {
	macro.Rewrite("$_.removeAttributeNode($1)", oldAttr)
	return w
}

// RemoveAttributeNS fn
// js:"removeAttributeNS"
func (*SVGFECompositeElement) RemoveAttributeNS(namespaceURI string, localName string) {
	macro.Rewrite("$_.removeAttributeNS($1, $2)", namespaceURI, localName)
}

// RequestFullscreen fn
// js:"requestFullscreen"
func (*SVGFECompositeElement) RequestFullscreen() {
	macro.Rewrite("$_.requestFullscreen()")
}

// RequestPointerLock fn
// js:"requestPointerLock"
func (*SVGFECompositeElement) RequestPointerLock() {
	macro.Rewrite("$_.requestPointerLock()")
}

// SetAttribute fn
// js:"setAttribute"
func (*SVGFECompositeElement) SetAttribute(qualifiedName string, value string) {
	macro.Rewrite("$_.setAttribute($1, $2)", qualifiedName, value)
}

// SetAttributeNode fn
// js:"setAttributeNode"
func (*SVGFECompositeElement) SetAttributeNode(newAttr *window.Attr) (w *window.Attr) {
	macro.Rewrite("$_.setAttributeNode($1)", newAttr)
	return w
}

// SetAttributeNodeNS fn
// js:"setAttributeNodeNS"
func (*SVGFECompositeElement) SetAttributeNodeNS(newAttr *window.Attr) (w *window.Attr) {
	macro.Rewrite("$_.setAttributeNodeNS($1)", newAttr)
	return w
}

// SetAttributeNS fn
// js:"setAttributeNS"
func (*SVGFECompositeElement) SetAttributeNS(namespaceURI string, qualifiedName string, value string) {
	macro.Rewrite("$_.setAttributeNS($1, $2, $3)", namespaceURI, qualifiedName, value)
}

// SetPointerCapture fn
// js:"setPointerCapture"
func (*SVGFECompositeElement) SetPointerCapture(pointerId int) {
	macro.Rewrite("$_.setPointerCapture($1)", pointerId)
}

// WebkitMatchesSelector fn
// js:"webkitMatchesSelector"
func (*SVGFECompositeElement) WebkitMatchesSelector(selectors string) (b bool) {
	macro.Rewrite("$_.webkitMatchesSelector($1)", selectors)
	return b
}

// WebkitRequestFullscreen fn
// js:"webkitRequestFullscreen"
func (*SVGFECompositeElement) WebkitRequestFullscreen() {
	macro.Rewrite("$_.webkitRequestFullscreen()")
}

// WebkitRequestFullScreen fn
// js:"webkitRequestFullScreen"
func (*SVGFECompositeElement) WebkitRequestFullScreen() {
	macro.Rewrite("$_.webkitRequestFullScreen()")
}

// QuerySelector fn
// js:"querySelector"
func (*SVGFECompositeElement) QuerySelector(selectors string) (w window.Element) {
	macro.Rewrite("$_.querySelector($1)", selectors)
	return w
}

// QuerySelectorAll fn
// js:"querySelectorAll"
func (*SVGFECompositeElement) QuerySelectorAll(selectors string) (w *window.NodeList) {
	macro.Rewrite("$_.querySelectorAll($1)", selectors)
	return w
}

// AppendChild fn
// js:"appendChild"
func (*SVGFECompositeElement) AppendChild(newChild window.Node) (w window.Node) {
	macro.Rewrite("$_.appendChild($1)", newChild)
	return w
}

// CloneNode fn
// js:"cloneNode"
func (*SVGFECompositeElement) CloneNode(deep *bool) (w window.Node) {
	macro.Rewrite("$_.cloneNode($1)", deep)
	return w
}

// CompareDocumentPosition fn
// js:"compareDocumentPosition"
func (*SVGFECompositeElement) CompareDocumentPosition(other window.Node) (u uint8) {
	macro.Rewrite("$_.compareDocumentPosition($1)", other)
	return u
}

// Contains fn
// js:"contains"
func (*SVGFECompositeElement) Contains(child window.Node) (b bool) {
	macro.Rewrite("$_.contains($1)", child)
	return b
}

// HasAttributes fn
// js:"hasAttributes"
func (*SVGFECompositeElement) HasAttributes() (b bool) {
	macro.Rewrite("$_.hasAttributes()")
	return b
}

// HasChildNodes fn
// js:"hasChildNodes"
func (*SVGFECompositeElement) HasChildNodes() (b bool) {
	macro.Rewrite("$_.hasChildNodes()")
	return b
}

// InsertBefore fn
// js:"insertBefore"
func (*SVGFECompositeElement) InsertBefore(newChild window.Node, refChild *window.Node) (w window.Node) {
	macro.Rewrite("$_.insertBefore($1, $2)", newChild, refChild)
	return w
}

// IsDefaultNamespace fn
// js:"isDefaultNamespace"
func (*SVGFECompositeElement) IsDefaultNamespace(namespaceURI string) (b bool) {
	macro.Rewrite("$_.isDefaultNamespace($1)", namespaceURI)
	return b
}

// IsEqualNode fn
// js:"isEqualNode"
func (*SVGFECompositeElement) IsEqualNode(arg window.Node) (b bool) {
	macro.Rewrite("$_.isEqualNode($1)", arg)
	return b
}

// IsSameNode fn
// js:"isSameNode"
func (*SVGFECompositeElement) IsSameNode(other window.Node) (b bool) {
	macro.Rewrite("$_.isSameNode($1)", other)
	return b
}

// LookupNamespaceURI fn
// js:"lookupNamespaceURI"
func (*SVGFECompositeElement) LookupNamespaceURI(prefix string) (s string) {
	macro.Rewrite("$_.lookupNamespaceURI($1)", prefix)
	return s
}

// LookupPrefix fn
// js:"lookupPrefix"
func (*SVGFECompositeElement) LookupPrefix(namespaceURI string) (s string) {
	macro.Rewrite("$_.lookupPrefix($1)", namespaceURI)
	return s
}

// Normalize fn
// js:"normalize"
func (*SVGFECompositeElement) Normalize() {
	macro.Rewrite("$_.normalize()")
}

// RemoveChild fn
// js:"removeChild"
func (*SVGFECompositeElement) RemoveChild(oldChild window.Node) (w window.Node) {
	macro.Rewrite("$_.removeChild($1)", oldChild)
	return w
}

// ReplaceChild fn
// js:"replaceChild"
func (*SVGFECompositeElement) ReplaceChild(newChild window.Node, oldChild window.Node) (w window.Node) {
	macro.Rewrite("$_.replaceChild($1, $2)", newChild, oldChild)
	return w
}

// AddEventListener fn
// js:"addEventListener"
func (*SVGFECompositeElement) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*SVGFECompositeElement) DispatchEvent(evt window.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*SVGFECompositeElement) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// In1 prop
// js:"in1"
func (*SVGFECompositeElement) In1() (in1 *dom.SVGAnimatedString) {
	macro.Rewrite("$_.in1")
	return in1
}

// In2 prop
// js:"in2"
func (*SVGFECompositeElement) In2() (in2 *dom.SVGAnimatedString) {
	macro.Rewrite("$_.in2")
	return in2
}

// K1 prop
// js:"k1"
func (*SVGFECompositeElement) K1() (k1 *dom.SVGAnimatedNumber) {
	macro.Rewrite("$_.k1")
	return k1
}

// K2 prop
// js:"k2"
func (*SVGFECompositeElement) K2() (k2 *dom.SVGAnimatedNumber) {
	macro.Rewrite("$_.k2")
	return k2
}

// K3 prop
// js:"k3"
func (*SVGFECompositeElement) K3() (k3 *dom.SVGAnimatedNumber) {
	macro.Rewrite("$_.k3")
	return k3
}

// K4 prop
// js:"k4"
func (*SVGFECompositeElement) K4() (k4 *dom.SVGAnimatedNumber) {
	macro.Rewrite("$_.k4")
	return k4
}

// Operator prop
// js:"operator"
func (*SVGFECompositeElement) Operator() (operator *dom.SVGAnimatedEnumeration) {
	macro.Rewrite("$_.operator")
	return operator
}

// Height prop
// js:"height"
func (*SVGFECompositeElement) Height() (height *dom.SVGAnimatedLength) {
	macro.Rewrite("$_.height")
	return height
}

// Result prop
// js:"result"
func (*SVGFECompositeElement) Result() (result *dom.SVGAnimatedString) {
	macro.Rewrite("$_.result")
	return result
}

// Width prop
// js:"width"
func (*SVGFECompositeElement) Width() (width *dom.SVGAnimatedLength) {
	macro.Rewrite("$_.width")
	return width
}

// X prop
// js:"x"
func (*SVGFECompositeElement) X() (x *dom.SVGAnimatedLength) {
	macro.Rewrite("$_.x")
	return x
}

// Y prop
// js:"y"
func (*SVGFECompositeElement) Y() (y *dom.SVGAnimatedLength) {
	macro.Rewrite("$_.y")
	return y
}

// Dataset prop
// js:"dataset"
func (*SVGFECompositeElement) Dataset() (dataset *dom.DOMStringMap) {
	macro.Rewrite("$_.dataset")
	return dataset
}

// OwnerSVGElement prop
// js:"ownerSVGElement"
func (*SVGFECompositeElement) OwnerSVGElement() (ownerSVGElement *window.SVGSVGElement) {
	macro.Rewrite("$_.ownerSVGElement")
	return ownerSVGElement
}

// ViewportElement prop
// js:"viewportElement"
func (*SVGFECompositeElement) ViewportElement() (viewportElement window.SVGElement) {
	macro.Rewrite("$_.viewportElement")
	return viewportElement
}

// Xmlbase prop
// js:"xmlbase"
func (*SVGFECompositeElement) Xmlbase() (xmlbase string) {
	macro.Rewrite("$_.xmlbase")
	return xmlbase
}

// SetXmlbase prop
// js:"xmlbase"
func (*SVGFECompositeElement) SetXmlbase(xmlbase string) {
	macro.Rewrite("$_.xmlbase = $1", xmlbase)
}

// ClassList prop
// js:"classList"
func (*SVGFECompositeElement) ClassList() (classList dom.DOMTokenList) {
	macro.Rewrite("$_.classList")
	return classList
}

// ClassName prop
// js:"className"
func (*SVGFECompositeElement) ClassName() (className string) {
	macro.Rewrite("$_.className")
	return className
}

// SetClassName prop
// js:"className"
func (*SVGFECompositeElement) SetClassName(className string) {
	macro.Rewrite("$_.className = $1", className)
}

// ClientHeight prop
// js:"clientHeight"
func (*SVGFECompositeElement) ClientHeight() (clientHeight int) {
	macro.Rewrite("$_.clientHeight")
	return clientHeight
}

// ClientLeft prop
// js:"clientLeft"
func (*SVGFECompositeElement) ClientLeft() (clientLeft int) {
	macro.Rewrite("$_.clientLeft")
	return clientLeft
}

// ClientTop prop
// js:"clientTop"
func (*SVGFECompositeElement) ClientTop() (clientTop int) {
	macro.Rewrite("$_.clientTop")
	return clientTop
}

// ClientWidth prop
// js:"clientWidth"
func (*SVGFECompositeElement) ClientWidth() (clientWidth int) {
	macro.Rewrite("$_.clientWidth")
	return clientWidth
}

// ID prop
// js:"id"
func (*SVGFECompositeElement) ID() (id string) {
	macro.Rewrite("$_.id")
	return id
}

// SetID prop
// js:"id"
func (*SVGFECompositeElement) SetID(id string) {
	macro.Rewrite("$_.id = $1", id)
}

// InnerHTML prop
// js:"innerHTML"
func (*SVGFECompositeElement) InnerHTML() (innerHTML string) {
	macro.Rewrite("$_.innerHTML")
	return innerHTML
}

// SetInnerHTML prop
// js:"innerHTML"
func (*SVGFECompositeElement) SetInnerHTML(innerHTML string) {
	macro.Rewrite("$_.innerHTML = $1", innerHTML)
}

// MsContentZoomFactor prop
// js:"msContentZoomFactor"
func (*SVGFECompositeElement) MsContentZoomFactor() (msContentZoomFactor float32) {
	macro.Rewrite("$_.msContentZoomFactor")
	return msContentZoomFactor
}

// SetMsContentZoomFactor prop
// js:"msContentZoomFactor"
func (*SVGFECompositeElement) SetMsContentZoomFactor(msContentZoomFactor float32) {
	macro.Rewrite("$_.msContentZoomFactor = $1", msContentZoomFactor)
}

// MsRegionOverflow prop
// js:"msRegionOverflow"
func (*SVGFECompositeElement) MsRegionOverflow() (msRegionOverflow string) {
	macro.Rewrite("$_.msRegionOverflow")
	return msRegionOverflow
}

// Onariarequest prop
// js:"onariarequest"
func (*SVGFECompositeElement) Onariarequest() (onariarequest func(window.Event)) {
	macro.Rewrite("$_.onariarequest")
	return onariarequest
}

// SetOnariarequest prop
// js:"onariarequest"
func (*SVGFECompositeElement) SetOnariarequest(onariarequest func(window.Event)) {
	macro.Rewrite("$_.onariarequest = $1", onariarequest)
}

// Oncommand prop
// js:"oncommand"
func (*SVGFECompositeElement) Oncommand() (oncommand func(window.Event)) {
	macro.Rewrite("$_.oncommand")
	return oncommand
}

// SetOncommand prop
// js:"oncommand"
func (*SVGFECompositeElement) SetOncommand(oncommand func(window.Event)) {
	macro.Rewrite("$_.oncommand = $1", oncommand)
}

// Ongotpointercapture prop
// js:"ongotpointercapture"
func (*SVGFECompositeElement) Ongotpointercapture() (ongotpointercapture func(*window.PointerEvent)) {
	macro.Rewrite("$_.ongotpointercapture")
	return ongotpointercapture
}

// SetOngotpointercapture prop
// js:"ongotpointercapture"
func (*SVGFECompositeElement) SetOngotpointercapture(ongotpointercapture func(*window.PointerEvent)) {
	macro.Rewrite("$_.ongotpointercapture = $1", ongotpointercapture)
}

// Onlostpointercapture prop
// js:"onlostpointercapture"
func (*SVGFECompositeElement) Onlostpointercapture() (onlostpointercapture func(*window.PointerEvent)) {
	macro.Rewrite("$_.onlostpointercapture")
	return onlostpointercapture
}

// SetOnlostpointercapture prop
// js:"onlostpointercapture"
func (*SVGFECompositeElement) SetOnlostpointercapture(onlostpointercapture func(*window.PointerEvent)) {
	macro.Rewrite("$_.onlostpointercapture = $1", onlostpointercapture)
}

// Onmsgesturechange prop
// js:"onmsgesturechange"
func (*SVGFECompositeElement) Onmsgesturechange() (onmsgesturechange func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturechange")
	return onmsgesturechange
}

// SetOnmsgesturechange prop
// js:"onmsgesturechange"
func (*SVGFECompositeElement) SetOnmsgesturechange(onmsgesturechange func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturechange = $1", onmsgesturechange)
}

// Onmsgesturedoubletap prop
// js:"onmsgesturedoubletap"
func (*SVGFECompositeElement) Onmsgesturedoubletap() (onmsgesturedoubletap func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturedoubletap")
	return onmsgesturedoubletap
}

// SetOnmsgesturedoubletap prop
// js:"onmsgesturedoubletap"
func (*SVGFECompositeElement) SetOnmsgesturedoubletap(onmsgesturedoubletap func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturedoubletap = $1", onmsgesturedoubletap)
}

// Onmsgestureend prop
// js:"onmsgestureend"
func (*SVGFECompositeElement) Onmsgestureend() (onmsgestureend func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgestureend")
	return onmsgestureend
}

// SetOnmsgestureend prop
// js:"onmsgestureend"
func (*SVGFECompositeElement) SetOnmsgestureend(onmsgestureend func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgestureend = $1", onmsgestureend)
}

// Onmsgesturehold prop
// js:"onmsgesturehold"
func (*SVGFECompositeElement) Onmsgesturehold() (onmsgesturehold func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturehold")
	return onmsgesturehold
}

// SetOnmsgesturehold prop
// js:"onmsgesturehold"
func (*SVGFECompositeElement) SetOnmsgesturehold(onmsgesturehold func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturehold = $1", onmsgesturehold)
}

// Onmsgesturestart prop
// js:"onmsgesturestart"
func (*SVGFECompositeElement) Onmsgesturestart() (onmsgesturestart func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturestart")
	return onmsgesturestart
}

// SetOnmsgesturestart prop
// js:"onmsgesturestart"
func (*SVGFECompositeElement) SetOnmsgesturestart(onmsgesturestart func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturestart = $1", onmsgesturestart)
}

// Onmsgesturetap prop
// js:"onmsgesturetap"
func (*SVGFECompositeElement) Onmsgesturetap() (onmsgesturetap func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturetap")
	return onmsgesturetap
}

// SetOnmsgesturetap prop
// js:"onmsgesturetap"
func (*SVGFECompositeElement) SetOnmsgesturetap(onmsgesturetap func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturetap = $1", onmsgesturetap)
}

// Onmsgotpointercapture prop
// js:"onmsgotpointercapture"
func (*SVGFECompositeElement) Onmsgotpointercapture() (onmsgotpointercapture func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmsgotpointercapture")
	return onmsgotpointercapture
}

// SetOnmsgotpointercapture prop
// js:"onmsgotpointercapture"
func (*SVGFECompositeElement) SetOnmsgotpointercapture(onmsgotpointercapture func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmsgotpointercapture = $1", onmsgotpointercapture)
}

// Onmsinertiastart prop
// js:"onmsinertiastart"
func (*SVGFECompositeElement) Onmsinertiastart() (onmsinertiastart func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsinertiastart")
	return onmsinertiastart
}

// SetOnmsinertiastart prop
// js:"onmsinertiastart"
func (*SVGFECompositeElement) SetOnmsinertiastart(onmsinertiastart func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsinertiastart = $1", onmsinertiastart)
}

// Onmslostpointercapture prop
// js:"onmslostpointercapture"
func (*SVGFECompositeElement) Onmslostpointercapture() (onmslostpointercapture func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmslostpointercapture")
	return onmslostpointercapture
}

// SetOnmslostpointercapture prop
// js:"onmslostpointercapture"
func (*SVGFECompositeElement) SetOnmslostpointercapture(onmslostpointercapture func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmslostpointercapture = $1", onmslostpointercapture)
}

// Onmspointercancel prop
// js:"onmspointercancel"
func (*SVGFECompositeElement) Onmspointercancel() (onmspointercancel func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointercancel")
	return onmspointercancel
}

// SetOnmspointercancel prop
// js:"onmspointercancel"
func (*SVGFECompositeElement) SetOnmspointercancel(onmspointercancel func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointercancel = $1", onmspointercancel)
}

// Onmspointerdown prop
// js:"onmspointerdown"
func (*SVGFECompositeElement) Onmspointerdown() (onmspointerdown func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerdown")
	return onmspointerdown
}

// SetOnmspointerdown prop
// js:"onmspointerdown"
func (*SVGFECompositeElement) SetOnmspointerdown(onmspointerdown func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerdown = $1", onmspointerdown)
}

// Onmspointerenter prop
// js:"onmspointerenter"
func (*SVGFECompositeElement) Onmspointerenter() (onmspointerenter func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerenter")
	return onmspointerenter
}

// SetOnmspointerenter prop
// js:"onmspointerenter"
func (*SVGFECompositeElement) SetOnmspointerenter(onmspointerenter func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerenter = $1", onmspointerenter)
}

// Onmspointerleave prop
// js:"onmspointerleave"
func (*SVGFECompositeElement) Onmspointerleave() (onmspointerleave func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerleave")
	return onmspointerleave
}

// SetOnmspointerleave prop
// js:"onmspointerleave"
func (*SVGFECompositeElement) SetOnmspointerleave(onmspointerleave func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerleave = $1", onmspointerleave)
}

// Onmspointermove prop
// js:"onmspointermove"
func (*SVGFECompositeElement) Onmspointermove() (onmspointermove func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointermove")
	return onmspointermove
}

// SetOnmspointermove prop
// js:"onmspointermove"
func (*SVGFECompositeElement) SetOnmspointermove(onmspointermove func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointermove = $1", onmspointermove)
}

// Onmspointerout prop
// js:"onmspointerout"
func (*SVGFECompositeElement) Onmspointerout() (onmspointerout func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerout")
	return onmspointerout
}

// SetOnmspointerout prop
// js:"onmspointerout"
func (*SVGFECompositeElement) SetOnmspointerout(onmspointerout func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerout = $1", onmspointerout)
}

// Onmspointerover prop
// js:"onmspointerover"
func (*SVGFECompositeElement) Onmspointerover() (onmspointerover func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerover")
	return onmspointerover
}

// SetOnmspointerover prop
// js:"onmspointerover"
func (*SVGFECompositeElement) SetOnmspointerover(onmspointerover func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerover = $1", onmspointerover)
}

// Onmspointerup prop
// js:"onmspointerup"
func (*SVGFECompositeElement) Onmspointerup() (onmspointerup func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerup")
	return onmspointerup
}

// SetOnmspointerup prop
// js:"onmspointerup"
func (*SVGFECompositeElement) SetOnmspointerup(onmspointerup func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerup = $1", onmspointerup)
}

// Ontouchcancel prop
// js:"ontouchcancel"
func (*SVGFECompositeElement) Ontouchcancel() (ontouchcancel func(*window.TouchEvent)) {
	macro.Rewrite("$_.ontouchcancel")
	return ontouchcancel
}

// SetOntouchcancel prop
// js:"ontouchcancel"
func (*SVGFECompositeElement) SetOntouchcancel(ontouchcancel func(*window.TouchEvent)) {
	macro.Rewrite("$_.ontouchcancel = $1", ontouchcancel)
}

// Ontouchend prop
// js:"ontouchend"
func (*SVGFECompositeElement) Ontouchend() (ontouchend func(*window.TouchEvent)) {
	macro.Rewrite("$_.ontouchend")
	return ontouchend
}

// SetOntouchend prop
// js:"ontouchend"
func (*SVGFECompositeElement) SetOntouchend(ontouchend func(*window.TouchEvent)) {
	macro.Rewrite("$_.ontouchend = $1", ontouchend)
}

// Ontouchmove prop
// js:"ontouchmove"
func (*SVGFECompositeElement) Ontouchmove() (ontouchmove func(*window.TouchEvent)) {
	macro.Rewrite("$_.ontouchmove")
	return ontouchmove
}

// SetOntouchmove prop
// js:"ontouchmove"
func (*SVGFECompositeElement) SetOntouchmove(ontouchmove func(*window.TouchEvent)) {
	macro.Rewrite("$_.ontouchmove = $1", ontouchmove)
}

// Ontouchstart prop
// js:"ontouchstart"
func (*SVGFECompositeElement) Ontouchstart() (ontouchstart func(*window.TouchEvent)) {
	macro.Rewrite("$_.ontouchstart")
	return ontouchstart
}

// SetOntouchstart prop
// js:"ontouchstart"
func (*SVGFECompositeElement) SetOntouchstart(ontouchstart func(*window.TouchEvent)) {
	macro.Rewrite("$_.ontouchstart = $1", ontouchstart)
}

// Onwebkitfullscreenchange prop
// js:"onwebkitfullscreenchange"
func (*SVGFECompositeElement) Onwebkitfullscreenchange() (onwebkitfullscreenchange func(window.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenchange")
	return onwebkitfullscreenchange
}

// SetOnwebkitfullscreenchange prop
// js:"onwebkitfullscreenchange"
func (*SVGFECompositeElement) SetOnwebkitfullscreenchange(onwebkitfullscreenchange func(window.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenchange = $1", onwebkitfullscreenchange)
}

// Onwebkitfullscreenerror prop
// js:"onwebkitfullscreenerror"
func (*SVGFECompositeElement) Onwebkitfullscreenerror() (onwebkitfullscreenerror func(window.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenerror")
	return onwebkitfullscreenerror
}

// SetOnwebkitfullscreenerror prop
// js:"onwebkitfullscreenerror"
func (*SVGFECompositeElement) SetOnwebkitfullscreenerror(onwebkitfullscreenerror func(window.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenerror = $1", onwebkitfullscreenerror)
}

// OuterHTML prop
// js:"outerHTML"
func (*SVGFECompositeElement) OuterHTML() (outerHTML string) {
	macro.Rewrite("$_.outerHTML")
	return outerHTML
}

// SetOuterHTML prop
// js:"outerHTML"
func (*SVGFECompositeElement) SetOuterHTML(outerHTML string) {
	macro.Rewrite("$_.outerHTML = $1", outerHTML)
}

// Prefix prop
// js:"prefix"
func (*SVGFECompositeElement) Prefix() (prefix string) {
	macro.Rewrite("$_.prefix")
	return prefix
}

// ScrollHeight prop
// js:"scrollHeight"
func (*SVGFECompositeElement) ScrollHeight() (scrollHeight int) {
	macro.Rewrite("$_.scrollHeight")
	return scrollHeight
}

// ScrollLeft prop
// js:"scrollLeft"
func (*SVGFECompositeElement) ScrollLeft() (scrollLeft int) {
	macro.Rewrite("$_.scrollLeft")
	return scrollLeft
}

// SetScrollLeft prop
// js:"scrollLeft"
func (*SVGFECompositeElement) SetScrollLeft(scrollLeft int) {
	macro.Rewrite("$_.scrollLeft = $1", scrollLeft)
}

// ScrollTop prop
// js:"scrollTop"
func (*SVGFECompositeElement) ScrollTop() (scrollTop int) {
	macro.Rewrite("$_.scrollTop")
	return scrollTop
}

// SetScrollTop prop
// js:"scrollTop"
func (*SVGFECompositeElement) SetScrollTop(scrollTop int) {
	macro.Rewrite("$_.scrollTop = $1", scrollTop)
}

// ScrollWidth prop
// js:"scrollWidth"
func (*SVGFECompositeElement) ScrollWidth() (scrollWidth int) {
	macro.Rewrite("$_.scrollWidth")
	return scrollWidth
}

// TagName prop
// js:"tagName"
func (*SVGFECompositeElement) TagName() (tagName string) {
	macro.Rewrite("$_.tagName")
	return tagName
}

// Onpointercancel prop
// js:"onpointercancel"
func (*SVGFECompositeElement) Onpointercancel() (onpointercancel func(window.Event)) {
	macro.Rewrite("$_.onpointercancel")
	return onpointercancel
}

// SetOnpointercancel prop
// js:"onpointercancel"
func (*SVGFECompositeElement) SetOnpointercancel(onpointercancel func(window.Event)) {
	macro.Rewrite("$_.onpointercancel = $1", onpointercancel)
}

// Onpointerdown prop
// js:"onpointerdown"
func (*SVGFECompositeElement) Onpointerdown() (onpointerdown func(window.Event)) {
	macro.Rewrite("$_.onpointerdown")
	return onpointerdown
}

// SetOnpointerdown prop
// js:"onpointerdown"
func (*SVGFECompositeElement) SetOnpointerdown(onpointerdown func(window.Event)) {
	macro.Rewrite("$_.onpointerdown = $1", onpointerdown)
}

// Onpointerenter prop
// js:"onpointerenter"
func (*SVGFECompositeElement) Onpointerenter() (onpointerenter func(window.Event)) {
	macro.Rewrite("$_.onpointerenter")
	return onpointerenter
}

// SetOnpointerenter prop
// js:"onpointerenter"
func (*SVGFECompositeElement) SetOnpointerenter(onpointerenter func(window.Event)) {
	macro.Rewrite("$_.onpointerenter = $1", onpointerenter)
}

// Onpointerleave prop
// js:"onpointerleave"
func (*SVGFECompositeElement) Onpointerleave() (onpointerleave func(window.Event)) {
	macro.Rewrite("$_.onpointerleave")
	return onpointerleave
}

// SetOnpointerleave prop
// js:"onpointerleave"
func (*SVGFECompositeElement) SetOnpointerleave(onpointerleave func(window.Event)) {
	macro.Rewrite("$_.onpointerleave = $1", onpointerleave)
}

// Onpointermove prop
// js:"onpointermove"
func (*SVGFECompositeElement) Onpointermove() (onpointermove func(window.Event)) {
	macro.Rewrite("$_.onpointermove")
	return onpointermove
}

// SetOnpointermove prop
// js:"onpointermove"
func (*SVGFECompositeElement) SetOnpointermove(onpointermove func(window.Event)) {
	macro.Rewrite("$_.onpointermove = $1", onpointermove)
}

// Onpointerout prop
// js:"onpointerout"
func (*SVGFECompositeElement) Onpointerout() (onpointerout func(window.Event)) {
	macro.Rewrite("$_.onpointerout")
	return onpointerout
}

// SetOnpointerout prop
// js:"onpointerout"
func (*SVGFECompositeElement) SetOnpointerout(onpointerout func(window.Event)) {
	macro.Rewrite("$_.onpointerout = $1", onpointerout)
}

// Onpointerover prop
// js:"onpointerover"
func (*SVGFECompositeElement) Onpointerover() (onpointerover func(window.Event)) {
	macro.Rewrite("$_.onpointerover")
	return onpointerover
}

// SetOnpointerover prop
// js:"onpointerover"
func (*SVGFECompositeElement) SetOnpointerover(onpointerover func(window.Event)) {
	macro.Rewrite("$_.onpointerover = $1", onpointerover)
}

// Onpointerup prop
// js:"onpointerup"
func (*SVGFECompositeElement) Onpointerup() (onpointerup func(window.Event)) {
	macro.Rewrite("$_.onpointerup")
	return onpointerup
}

// SetOnpointerup prop
// js:"onpointerup"
func (*SVGFECompositeElement) SetOnpointerup(onpointerup func(window.Event)) {
	macro.Rewrite("$_.onpointerup = $1", onpointerup)
}

// Onwheel prop
// js:"onwheel"
func (*SVGFECompositeElement) Onwheel() (onwheel func(window.Event)) {
	macro.Rewrite("$_.onwheel")
	return onwheel
}

// SetOnwheel prop
// js:"onwheel"
func (*SVGFECompositeElement) SetOnwheel(onwheel func(window.Event)) {
	macro.Rewrite("$_.onwheel = $1", onwheel)
}

// ChildElementCount prop
// js:"childElementCount"
func (*SVGFECompositeElement) ChildElementCount() (childElementCount uint) {
	macro.Rewrite("$_.childElementCount")
	return childElementCount
}

// FirstElementChild prop
// js:"firstElementChild"
func (*SVGFECompositeElement) FirstElementChild() (firstElementChild window.Element) {
	macro.Rewrite("$_.firstElementChild")
	return firstElementChild
}

// LastElementChild prop
// js:"lastElementChild"
func (*SVGFECompositeElement) LastElementChild() (lastElementChild window.Element) {
	macro.Rewrite("$_.lastElementChild")
	return lastElementChild
}

// NextElementSibling prop
// js:"nextElementSibling"
func (*SVGFECompositeElement) NextElementSibling() (nextElementSibling window.Element) {
	macro.Rewrite("$_.nextElementSibling")
	return nextElementSibling
}

// PreviousElementSibling prop
// js:"previousElementSibling"
func (*SVGFECompositeElement) PreviousElementSibling() (previousElementSibling window.Element) {
	macro.Rewrite("$_.previousElementSibling")
	return previousElementSibling
}

// Attributes prop
// js:"attributes"
func (*SVGFECompositeElement) Attributes() (attributes *window.NamedNodeMap) {
	macro.Rewrite("$_.attributes")
	return attributes
}

// BaseURI prop
// js:"baseURI"
func (*SVGFECompositeElement) BaseURI() (baseURI string) {
	macro.Rewrite("$_.baseURI")
	return baseURI
}

// ChildNodes prop
// js:"childNodes"
func (*SVGFECompositeElement) ChildNodes() (childNodes *window.NodeList) {
	macro.Rewrite("$_.childNodes")
	return childNodes
}

// FirstChild prop
// js:"firstChild"
func (*SVGFECompositeElement) FirstChild() (firstChild window.Node) {
	macro.Rewrite("$_.firstChild")
	return firstChild
}

// LastChild prop
// js:"lastChild"
func (*SVGFECompositeElement) LastChild() (lastChild window.Node) {
	macro.Rewrite("$_.lastChild")
	return lastChild
}

// LocalName prop
// js:"localName"
func (*SVGFECompositeElement) LocalName() (localName string) {
	macro.Rewrite("$_.localName")
	return localName
}

// NamespaceURI prop
// js:"namespaceURI"
func (*SVGFECompositeElement) NamespaceURI() (namespaceURI string) {
	macro.Rewrite("$_.namespaceURI")
	return namespaceURI
}

// NextSibling prop
// js:"nextSibling"
func (*SVGFECompositeElement) NextSibling() (nextSibling window.Node) {
	macro.Rewrite("$_.nextSibling")
	return nextSibling
}

// NodeName prop
// js:"nodeName"
func (*SVGFECompositeElement) NodeName() (nodeName string) {
	macro.Rewrite("$_.nodeName")
	return nodeName
}

// NodeType prop
// js:"nodeType"
func (*SVGFECompositeElement) NodeType() (nodeType uint8) {
	macro.Rewrite("$_.nodeType")
	return nodeType
}

// NodeValue prop
// js:"nodeValue"
func (*SVGFECompositeElement) NodeValue() (nodeValue string) {
	macro.Rewrite("$_.nodeValue")
	return nodeValue
}

// SetNodeValue prop
// js:"nodeValue"
func (*SVGFECompositeElement) SetNodeValue(nodeValue string) {
	macro.Rewrite("$_.nodeValue = $1", nodeValue)
}

// OwnerDocument prop
// js:"ownerDocument"
func (*SVGFECompositeElement) OwnerDocument() (ownerDocument window.Document) {
	macro.Rewrite("$_.ownerDocument")
	return ownerDocument
}

// ParentElement prop
// js:"parentElement"
func (*SVGFECompositeElement) ParentElement() (parentElement window.HTMLElement) {
	macro.Rewrite("$_.parentElement")
	return parentElement
}

// ParentNode prop
// js:"parentNode"
func (*SVGFECompositeElement) ParentNode() (parentNode window.Node) {
	macro.Rewrite("$_.parentNode")
	return parentNode
}

// PreviousSibling prop
// js:"previousSibling"
func (*SVGFECompositeElement) PreviousSibling() (previousSibling window.Node) {
	macro.Rewrite("$_.previousSibling")
	return previousSibling
}

// TextContent prop
// js:"textContent"
func (*SVGFECompositeElement) TextContent() (textContent string) {
	macro.Rewrite("$_.textContent")
	return textContent
}

// SetTextContent prop
// js:"textContent"
func (*SVGFECompositeElement) SetTextContent(textContent string) {
	macro.Rewrite("$_.textContent = $1", textContent)
}
