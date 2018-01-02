package html

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/element"
	"github.com/matthewmueller/joy/dom/utils"
	"github.com/matthewmueller/joy/dom/dom"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/ms"
	"github.com/matthewmueller/joy/dom/ui"
	"github.com/matthewmueller/joy/dom/mouse"
	"github.com/matthewmueller/joy/dom/draganddrop"
	"github.com/matthewmueller/joy/dom/css"
	"github.com/matthewmueller/joy/dom/touch"
	"github.com/matthewmueller/joy/dom/document"
)

var _ HTMLElement = (*HTMLTableCaptionElement)(nil)
var _ element.Element = (*HTMLTableCaptionElement)(nil)
var _ utils.GlobalEventHandlers = (*HTMLTableCaptionElement)(nil)
var _ element.ElementTraversal = (*HTMLTableCaptionElement)(nil)
var _ dom.NodeSelector = (*HTMLTableCaptionElement)(nil)
var _ dom.ChildNode = (*HTMLTableCaptionElement)(nil)
var _ dom.Node = (*HTMLTableCaptionElement)(nil)
var _ event.EventTarget = (*HTMLTableCaptionElement)(nil)

type HTMLTableCaptionElement struct {
}

func (*HTMLTableCaptionElement) Blur() {
	macro.Rewrite("$_.blur()")
}

func (*HTMLTableCaptionElement) Click() {
	macro.Rewrite("$_.click()")
}

func (*HTMLTableCaptionElement) DragDrop() (b bool) {
	macro.Rewrite("$_.dragDrop()")
	return b
}

func (*HTMLTableCaptionElement) Focus() {
	macro.Rewrite("$_.focus()")
}

func (*HTMLTableCaptionElement) GetElementsByClassName(classNames string) (w *dom.NodeList) {
	macro.Rewrite("$_.getElementsByClassName($1)", classNames)
	return w
}

func (*HTMLTableCaptionElement) InsertAdjacentElement(position string, insertedElement element.Element) (w element.Element) {
	macro.Rewrite("$_.insertAdjacentElement($1, $2)", position, insertedElement)
	return w
}

func (*HTMLTableCaptionElement) InsertAdjacentHTML(where string, html string) {
	macro.Rewrite("$_.insertAdjacentHTML($1, $2)", where, html)
}

func (*HTMLTableCaptionElement) InsertAdjacentText(where string, text string) {
	macro.Rewrite("$_.insertAdjacentText($1, $2)", where, text)
}

func (*HTMLTableCaptionElement) MsGetInputContext() (w *ms.MSInputMethodContext) {
	macro.Rewrite("$_.msGetInputContext()")
	return w
}

func (*HTMLTableCaptionElement) ScrollIntoView(top *bool) {
	macro.Rewrite("$_.scrollIntoView($1)", top)
}

func (*HTMLTableCaptionElement) GetAttribute(qualifiedName string) (s string) {
	macro.Rewrite("$_.getAttribute($1)", qualifiedName)
	return s
}

func (*HTMLTableCaptionElement) GetAttributeNode(name string) (w *dom.Attr) {
	macro.Rewrite("$_.getAttributeNode($1)", name)
	return w
}

func (*HTMLTableCaptionElement) GetAttributeNodeNS(namespaceURI string, localName string) (w *dom.Attr) {
	macro.Rewrite("$_.getAttributeNodeNS($1, $2)", namespaceURI, localName)
	return w
}

func (*HTMLTableCaptionElement) GetAttributeNS(namespaceURI string, localName string) (s string) {
	macro.Rewrite("$_.getAttributeNS($1, $2)", namespaceURI, localName)
	return s
}

func (*HTMLTableCaptionElement) GetBoundingClientRect() (c *dom.ClientRect) {
	macro.Rewrite("$_.getBoundingClientRect()")
	return c
}

func (*HTMLTableCaptionElement) GetClientRects() (c *dom.ClientRectList) {
	macro.Rewrite("$_.getClientRects()")
	return c
}

func (*HTMLTableCaptionElement) GetElementsByTagName(name string) (w *dom.NodeList) {
	macro.Rewrite("$_.getElementsByTagName($1)", name)
	return w
}

func (*HTMLTableCaptionElement) GetElementsByTagNameNS(namespaceURI string, localName string) (w *dom.NodeList) {
	macro.Rewrite("$_.getElementsByTagNameNS($1, $2)", namespaceURI, localName)
	return w
}

func (*HTMLTableCaptionElement) HasAttribute(name string) (b bool) {
	macro.Rewrite("$_.hasAttribute($1)", name)
	return b
}

func (*HTMLTableCaptionElement) HasAttributeNS(namespaceURI string, localName string) (b bool) {
	macro.Rewrite("$_.hasAttributeNS($1, $2)", namespaceURI, localName)
	return b
}

func (*HTMLTableCaptionElement) MsGetRegionContent() (w *ms.MSRangeCollection) {
	macro.Rewrite("$_.msGetRegionContent()")
	return w
}

func (*HTMLTableCaptionElement) MsGetUntransformedBounds() (c *dom.ClientRect) {
	macro.Rewrite("$_.msGetUntransformedBounds()")
	return c
}

func (*HTMLTableCaptionElement) MsMatchesSelector(selectors string) (b bool) {
	macro.Rewrite("$_.msMatchesSelector($1)", selectors)
	return b
}

func (*HTMLTableCaptionElement) MsReleasePointerCapture(pointerId int) {
	macro.Rewrite("$_.msReleasePointerCapture($1)", pointerId)
}

func (*HTMLTableCaptionElement) MsSetPointerCapture(pointerId int) {
	macro.Rewrite("$_.msSetPointerCapture($1)", pointerId)
}

func (*HTMLTableCaptionElement) MsZoomTo(args *ms.MsZoomToOptions) {
	macro.Rewrite("$_.msZoomTo($1)", args)
}

func (*HTMLTableCaptionElement) ReleasePointerCapture(pointerId int) {
	macro.Rewrite("$_.releasePointerCapture($1)", pointerId)
}

func (*HTMLTableCaptionElement) RemoveAttribute(qualifiedName string) {
	macro.Rewrite("$_.removeAttribute($1)", qualifiedName)
}

func (*HTMLTableCaptionElement) RemoveAttributeNode(oldAttr *dom.Attr) (w *dom.Attr) {
	macro.Rewrite("$_.removeAttributeNode($1)", oldAttr)
	return w
}

func (*HTMLTableCaptionElement) RemoveAttributeNS(namespaceURI string, localName string) {
	macro.Rewrite("$_.removeAttributeNS($1, $2)", namespaceURI, localName)
}

func (*HTMLTableCaptionElement) RequestFullscreen() {
	macro.Rewrite("$_.requestFullscreen()")
}

func (*HTMLTableCaptionElement) RequestPointerLock() {
	macro.Rewrite("$_.requestPointerLock()")
}

func (*HTMLTableCaptionElement) SetAttribute(qualifiedName string, value string) {
	macro.Rewrite("$_.setAttribute($1, $2)", qualifiedName, value)
}

func (*HTMLTableCaptionElement) SetAttributeNode(newAttr *dom.Attr) (w *dom.Attr) {
	macro.Rewrite("$_.setAttributeNode($1)", newAttr)
	return w
}

func (*HTMLTableCaptionElement) SetAttributeNodeNS(newAttr *dom.Attr) (w *dom.Attr) {
	macro.Rewrite("$_.setAttributeNodeNS($1)", newAttr)
	return w
}

func (*HTMLTableCaptionElement) SetAttributeNS(namespaceURI string, qualifiedName string, value string) {
	macro.Rewrite("$_.setAttributeNS($1, $2, $3)", namespaceURI, qualifiedName, value)
}

func (*HTMLTableCaptionElement) SetPointerCapture(pointerId int) {
	macro.Rewrite("$_.setPointerCapture($1)", pointerId)
}

func (*HTMLTableCaptionElement) WebkitMatchesSelector(selectors string) (b bool) {
	macro.Rewrite("$_.webkitMatchesSelector($1)", selectors)
	return b
}

func (*HTMLTableCaptionElement) WebkitRequestFullscreen() {
	macro.Rewrite("$_.webkitRequestFullscreen()")
}

func (*HTMLTableCaptionElement) WebkitRequestFullScreen() {
	macro.Rewrite("$_.webkitRequestFullScreen()")
}

func (*HTMLTableCaptionElement) QuerySelector(selectors string) (w element.Element) {
	macro.Rewrite("$_.querySelector($1)", selectors)
	return w
}

func (*HTMLTableCaptionElement) QuerySelectorAll(selectors string) (w *dom.NodeList) {
	macro.Rewrite("$_.querySelectorAll($1)", selectors)
	return w
}

func (*HTMLTableCaptionElement) AppendChild(newChild dom.Node) (w dom.Node) {
	macro.Rewrite("$_.appendChild($1)", newChild)
	return w
}

func (*HTMLTableCaptionElement) CloneNode(deep *bool) (w dom.Node) {
	macro.Rewrite("$_.cloneNode($1)", deep)
	return w
}

func (*HTMLTableCaptionElement) CompareDocumentPosition(other dom.Node) (u uint8) {
	macro.Rewrite("$_.compareDocumentPosition($1)", other)
	return u
}

func (*HTMLTableCaptionElement) Contains(child dom.Node) (b bool) {
	macro.Rewrite("$_.contains($1)", child)
	return b
}

func (*HTMLTableCaptionElement) HasAttributes() (b bool) {
	macro.Rewrite("$_.hasAttributes()")
	return b
}

func (*HTMLTableCaptionElement) HasChildNodes() (b bool) {
	macro.Rewrite("$_.hasChildNodes()")
	return b
}

func (*HTMLTableCaptionElement) InsertBefore(newChild dom.Node, refChild *dom.Node) (w dom.Node) {
	macro.Rewrite("$_.insertBefore($1, $2)", newChild, refChild)
	return w
}

func (*HTMLTableCaptionElement) IsDefaultNamespace(namespaceURI string) (b bool) {
	macro.Rewrite("$_.isDefaultNamespace($1)", namespaceURI)
	return b
}

func (*HTMLTableCaptionElement) IsEqualNode(arg dom.Node) (b bool) {
	macro.Rewrite("$_.isEqualNode($1)", arg)
	return b
}

func (*HTMLTableCaptionElement) IsSameNode(other dom.Node) (b bool) {
	macro.Rewrite("$_.isSameNode($1)", other)
	return b
}

func (*HTMLTableCaptionElement) LookupNamespaceURI(prefix string) (s string) {
	macro.Rewrite("$_.lookupNamespaceURI($1)", prefix)
	return s
}

func (*HTMLTableCaptionElement) LookupPrefix(namespaceURI string) (s string) {
	macro.Rewrite("$_.lookupPrefix($1)", namespaceURI)
	return s
}

func (*HTMLTableCaptionElement) Normalize() {
	macro.Rewrite("$_.normalize()")
}

func (*HTMLTableCaptionElement) RemoveChild(oldChild dom.Node) (w dom.Node) {
	macro.Rewrite("$_.removeChild($1)", oldChild)
	return w
}

func (*HTMLTableCaptionElement) ReplaceChild(newChild dom.Node, oldChild dom.Node) (w dom.Node) {
	macro.Rewrite("$_.replaceChild($1, $2)", newChild, oldChild)
	return w
}

func (*HTMLTableCaptionElement) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*HTMLTableCaptionElement) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*HTMLTableCaptionElement) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*HTMLTableCaptionElement) Align() (align string) {
	macro.Rewrite("$_.align")
	return align
}

func (*HTMLTableCaptionElement) SetAlign(align string) {
	macro.Rewrite("$_.align = $1", align)
}

func (*HTMLTableCaptionElement) VAlign() (vAlign string) {
	macro.Rewrite("$_.vAlign")
	return vAlign
}

func (*HTMLTableCaptionElement) SetVAlign(vAlign string) {
	macro.Rewrite("$_.vAlign = $1", vAlign)
}

func (*HTMLTableCaptionElement) AccessKey() (accessKey string) {
	macro.Rewrite("$_.accessKey")
	return accessKey
}

func (*HTMLTableCaptionElement) SetAccessKey(accessKey string) {
	macro.Rewrite("$_.accessKey = $1", accessKey)
}

func (*HTMLTableCaptionElement) Children() (children HTMLCollection) {
	macro.Rewrite("$_.children")
	return children
}

func (*HTMLTableCaptionElement) ContentEditable() (contentEditable string) {
	macro.Rewrite("$_.contentEditable")
	return contentEditable
}

func (*HTMLTableCaptionElement) SetContentEditable(contentEditable string) {
	macro.Rewrite("$_.contentEditable = $1", contentEditable)
}

func (*HTMLTableCaptionElement) Dataset() (dataset *dom.DOMStringMap) {
	macro.Rewrite("$_.dataset")
	return dataset
}

func (*HTMLTableCaptionElement) Dir() (dir string) {
	macro.Rewrite("$_.dir")
	return dir
}

func (*HTMLTableCaptionElement) SetDir(dir string) {
	macro.Rewrite("$_.dir = $1", dir)
}

func (*HTMLTableCaptionElement) Draggable() (draggable bool) {
	macro.Rewrite("$_.draggable")
	return draggable
}

func (*HTMLTableCaptionElement) SetDraggable(draggable bool) {
	macro.Rewrite("$_.draggable = $1", draggable)
}

func (*HTMLTableCaptionElement) Hidden() (hidden bool) {
	macro.Rewrite("$_.hidden")
	return hidden
}

func (*HTMLTableCaptionElement) SetHidden(hidden bool) {
	macro.Rewrite("$_.hidden = $1", hidden)
}

func (*HTMLTableCaptionElement) HideFocus() (hideFocus bool) {
	macro.Rewrite("$_.hideFocus")
	return hideFocus
}

func (*HTMLTableCaptionElement) SetHideFocus(hideFocus bool) {
	macro.Rewrite("$_.hideFocus = $1", hideFocus)
}

func (*HTMLTableCaptionElement) InnerText() (innerText string) {
	macro.Rewrite("$_.innerText")
	return innerText
}

func (*HTMLTableCaptionElement) SetInnerText(innerText string) {
	macro.Rewrite("$_.innerText = $1", innerText)
}

func (*HTMLTableCaptionElement) IsContentEditable() (isContentEditable bool) {
	macro.Rewrite("$_.isContentEditable")
	return isContentEditable
}

func (*HTMLTableCaptionElement) Lang() (lang string) {
	macro.Rewrite("$_.lang")
	return lang
}

func (*HTMLTableCaptionElement) SetLang(lang string) {
	macro.Rewrite("$_.lang = $1", lang)
}

func (*HTMLTableCaptionElement) OffsetHeight() (offsetHeight int) {
	macro.Rewrite("$_.offsetHeight")
	return offsetHeight
}

func (*HTMLTableCaptionElement) OffsetLeft() (offsetLeft int) {
	macro.Rewrite("$_.offsetLeft")
	return offsetLeft
}

func (*HTMLTableCaptionElement) OffsetParent() (offsetParent element.Element) {
	macro.Rewrite("$_.offsetParent")
	return offsetParent
}

func (*HTMLTableCaptionElement) OffsetTop() (offsetTop int) {
	macro.Rewrite("$_.offsetTop")
	return offsetTop
}

func (*HTMLTableCaptionElement) OffsetWidth() (offsetWidth int) {
	macro.Rewrite("$_.offsetWidth")
	return offsetWidth
}

func (*HTMLTableCaptionElement) Onabort() (onabort func(event.Event)) {
	macro.Rewrite("$_.onabort")
	return onabort
}

func (*HTMLTableCaptionElement) SetOnabort(onabort func(event.Event)) {
	macro.Rewrite("$_.onabort = $1", onabort)
}

func (*HTMLTableCaptionElement) Onactivate() (onactivate func(ui.UIEvent)) {
	macro.Rewrite("$_.onactivate")
	return onactivate
}

func (*HTMLTableCaptionElement) SetOnactivate(onactivate func(ui.UIEvent)) {
	macro.Rewrite("$_.onactivate = $1", onactivate)
}

func (*HTMLTableCaptionElement) Onbeforeactivate() (onbeforeactivate func(ui.UIEvent)) {
	macro.Rewrite("$_.onbeforeactivate")
	return onbeforeactivate
}

func (*HTMLTableCaptionElement) SetOnbeforeactivate(onbeforeactivate func(ui.UIEvent)) {
	macro.Rewrite("$_.onbeforeactivate = $1", onbeforeactivate)
}

func (*HTMLTableCaptionElement) Onbeforecopy() (onbeforecopy func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.onbeforecopy")
	return onbeforecopy
}

func (*HTMLTableCaptionElement) SetOnbeforecopy(onbeforecopy func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.onbeforecopy = $1", onbeforecopy)
}

func (*HTMLTableCaptionElement) Onbeforecut() (onbeforecut func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.onbeforecut")
	return onbeforecut
}

func (*HTMLTableCaptionElement) SetOnbeforecut(onbeforecut func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.onbeforecut = $1", onbeforecut)
}

func (*HTMLTableCaptionElement) Onbeforedeactivate() (onbeforedeactivate func(ui.UIEvent)) {
	macro.Rewrite("$_.onbeforedeactivate")
	return onbeforedeactivate
}

func (*HTMLTableCaptionElement) SetOnbeforedeactivate(onbeforedeactivate func(ui.UIEvent)) {
	macro.Rewrite("$_.onbeforedeactivate = $1", onbeforedeactivate)
}

func (*HTMLTableCaptionElement) Onbeforepaste() (onbeforepaste func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.onbeforepaste")
	return onbeforepaste
}

func (*HTMLTableCaptionElement) SetOnbeforepaste(onbeforepaste func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.onbeforepaste = $1", onbeforepaste)
}

func (*HTMLTableCaptionElement) Onblur() (onblur func(*utils.FocusEvent)) {
	macro.Rewrite("$_.onblur")
	return onblur
}

func (*HTMLTableCaptionElement) SetOnblur(onblur func(*utils.FocusEvent)) {
	macro.Rewrite("$_.onblur = $1", onblur)
}

func (*HTMLTableCaptionElement) Oncanplay() (oncanplay func(event.Event)) {
	macro.Rewrite("$_.oncanplay")
	return oncanplay
}

func (*HTMLTableCaptionElement) SetOncanplay(oncanplay func(event.Event)) {
	macro.Rewrite("$_.oncanplay = $1", oncanplay)
}

func (*HTMLTableCaptionElement) Oncanplaythrough() (oncanplaythrough func(event.Event)) {
	macro.Rewrite("$_.oncanplaythrough")
	return oncanplaythrough
}

func (*HTMLTableCaptionElement) SetOncanplaythrough(oncanplaythrough func(event.Event)) {
	macro.Rewrite("$_.oncanplaythrough = $1", oncanplaythrough)
}

func (*HTMLTableCaptionElement) Onchange() (onchange func(event.Event)) {
	macro.Rewrite("$_.onchange")
	return onchange
}

func (*HTMLTableCaptionElement) SetOnchange(onchange func(event.Event)) {
	macro.Rewrite("$_.onchange = $1", onchange)
}

func (*HTMLTableCaptionElement) Onclick() (onclick func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onclick")
	return onclick
}

func (*HTMLTableCaptionElement) SetOnclick(onclick func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onclick = $1", onclick)
}

func (*HTMLTableCaptionElement) Oncontextmenu() (oncontextmenu func(*utils.PointerEvent)) {
	macro.Rewrite("$_.oncontextmenu")
	return oncontextmenu
}

func (*HTMLTableCaptionElement) SetOncontextmenu(oncontextmenu func(*utils.PointerEvent)) {
	macro.Rewrite("$_.oncontextmenu = $1", oncontextmenu)
}

func (*HTMLTableCaptionElement) Oncopy() (oncopy func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.oncopy")
	return oncopy
}

func (*HTMLTableCaptionElement) SetOncopy(oncopy func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.oncopy = $1", oncopy)
}

func (*HTMLTableCaptionElement) Oncuechange() (oncuechange func(event.Event)) {
	macro.Rewrite("$_.oncuechange")
	return oncuechange
}

func (*HTMLTableCaptionElement) SetOncuechange(oncuechange func(event.Event)) {
	macro.Rewrite("$_.oncuechange = $1", oncuechange)
}

func (*HTMLTableCaptionElement) Oncut() (oncut func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.oncut")
	return oncut
}

func (*HTMLTableCaptionElement) SetOncut(oncut func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.oncut = $1", oncut)
}

func (*HTMLTableCaptionElement) Ondblclick() (ondblclick func(mouse.MouseEvent)) {
	macro.Rewrite("$_.ondblclick")
	return ondblclick
}

func (*HTMLTableCaptionElement) SetOndblclick(ondblclick func(mouse.MouseEvent)) {
	macro.Rewrite("$_.ondblclick = $1", ondblclick)
}

func (*HTMLTableCaptionElement) Ondeactivate() (ondeactivate func(ui.UIEvent)) {
	macro.Rewrite("$_.ondeactivate")
	return ondeactivate
}

func (*HTMLTableCaptionElement) SetOndeactivate(ondeactivate func(ui.UIEvent)) {
	macro.Rewrite("$_.ondeactivate = $1", ondeactivate)
}

func (*HTMLTableCaptionElement) Ondrag() (ondrag func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondrag")
	return ondrag
}

func (*HTMLTableCaptionElement) SetOndrag(ondrag func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondrag = $1", ondrag)
}

func (*HTMLTableCaptionElement) Ondragend() (ondragend func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragend")
	return ondragend
}

func (*HTMLTableCaptionElement) SetOndragend(ondragend func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragend = $1", ondragend)
}

func (*HTMLTableCaptionElement) Ondragenter() (ondragenter func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragenter")
	return ondragenter
}

func (*HTMLTableCaptionElement) SetOndragenter(ondragenter func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragenter = $1", ondragenter)
}

func (*HTMLTableCaptionElement) Ondragleave() (ondragleave func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragleave")
	return ondragleave
}

func (*HTMLTableCaptionElement) SetOndragleave(ondragleave func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragleave = $1", ondragleave)
}

func (*HTMLTableCaptionElement) Ondragover() (ondragover func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragover")
	return ondragover
}

func (*HTMLTableCaptionElement) SetOndragover(ondragover func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragover = $1", ondragover)
}

func (*HTMLTableCaptionElement) Ondragstart() (ondragstart func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragstart")
	return ondragstart
}

func (*HTMLTableCaptionElement) SetOndragstart(ondragstart func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragstart = $1", ondragstart)
}

func (*HTMLTableCaptionElement) Ondrop() (ondrop func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondrop")
	return ondrop
}

func (*HTMLTableCaptionElement) SetOndrop(ondrop func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondrop = $1", ondrop)
}

func (*HTMLTableCaptionElement) Ondurationchange() (ondurationchange func(event.Event)) {
	macro.Rewrite("$_.ondurationchange")
	return ondurationchange
}

func (*HTMLTableCaptionElement) SetOndurationchange(ondurationchange func(event.Event)) {
	macro.Rewrite("$_.ondurationchange = $1", ondurationchange)
}

func (*HTMLTableCaptionElement) Onemptied() (onemptied func(event.Event)) {
	macro.Rewrite("$_.onemptied")
	return onemptied
}

func (*HTMLTableCaptionElement) SetOnemptied(onemptied func(event.Event)) {
	macro.Rewrite("$_.onemptied = $1", onemptied)
}

func (*HTMLTableCaptionElement) Onended() (onended func(event.Event)) {
	macro.Rewrite("$_.onended")
	return onended
}

func (*HTMLTableCaptionElement) SetOnended(onended func(event.Event)) {
	macro.Rewrite("$_.onended = $1", onended)
}

func (*HTMLTableCaptionElement) Onerror() (onerror func(event.Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

func (*HTMLTableCaptionElement) SetOnerror(onerror func(event.Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

func (*HTMLTableCaptionElement) Onfocus() (onfocus func(*utils.FocusEvent)) {
	macro.Rewrite("$_.onfocus")
	return onfocus
}

func (*HTMLTableCaptionElement) SetOnfocus(onfocus func(*utils.FocusEvent)) {
	macro.Rewrite("$_.onfocus = $1", onfocus)
}

func (*HTMLTableCaptionElement) Oninput() (oninput func(event.Event)) {
	macro.Rewrite("$_.oninput")
	return oninput
}

func (*HTMLTableCaptionElement) SetOninput(oninput func(event.Event)) {
	macro.Rewrite("$_.oninput = $1", oninput)
}

func (*HTMLTableCaptionElement) Oninvalid() (oninvalid func(event.Event)) {
	macro.Rewrite("$_.oninvalid")
	return oninvalid
}

func (*HTMLTableCaptionElement) SetOninvalid(oninvalid func(event.Event)) {
	macro.Rewrite("$_.oninvalid = $1", oninvalid)
}

func (*HTMLTableCaptionElement) Onkeydown() (onkeydown func(*utils.KeyboardEvent)) {
	macro.Rewrite("$_.onkeydown")
	return onkeydown
}

func (*HTMLTableCaptionElement) SetOnkeydown(onkeydown func(*utils.KeyboardEvent)) {
	macro.Rewrite("$_.onkeydown = $1", onkeydown)
}

func (*HTMLTableCaptionElement) Onkeypress() (onkeypress func(*utils.KeyboardEvent)) {
	macro.Rewrite("$_.onkeypress")
	return onkeypress
}

func (*HTMLTableCaptionElement) SetOnkeypress(onkeypress func(*utils.KeyboardEvent)) {
	macro.Rewrite("$_.onkeypress = $1", onkeypress)
}

func (*HTMLTableCaptionElement) Onkeyup() (onkeyup func(*utils.KeyboardEvent)) {
	macro.Rewrite("$_.onkeyup")
	return onkeyup
}

func (*HTMLTableCaptionElement) SetOnkeyup(onkeyup func(*utils.KeyboardEvent)) {
	macro.Rewrite("$_.onkeyup = $1", onkeyup)
}

func (*HTMLTableCaptionElement) Onload() (onload func(event.Event)) {
	macro.Rewrite("$_.onload")
	return onload
}

func (*HTMLTableCaptionElement) SetOnload(onload func(event.Event)) {
	macro.Rewrite("$_.onload = $1", onload)
}

func (*HTMLTableCaptionElement) Onloadeddata() (onloadeddata func(event.Event)) {
	macro.Rewrite("$_.onloadeddata")
	return onloadeddata
}

func (*HTMLTableCaptionElement) SetOnloadeddata(onloadeddata func(event.Event)) {
	macro.Rewrite("$_.onloadeddata = $1", onloadeddata)
}

func (*HTMLTableCaptionElement) Onloadedmetadata() (onloadedmetadata func(event.Event)) {
	macro.Rewrite("$_.onloadedmetadata")
	return onloadedmetadata
}

func (*HTMLTableCaptionElement) SetOnloadedmetadata(onloadedmetadata func(event.Event)) {
	macro.Rewrite("$_.onloadedmetadata = $1", onloadedmetadata)
}

func (*HTMLTableCaptionElement) Onloadstart() (onloadstart func(event.Event)) {
	macro.Rewrite("$_.onloadstart")
	return onloadstart
}

func (*HTMLTableCaptionElement) SetOnloadstart(onloadstart func(event.Event)) {
	macro.Rewrite("$_.onloadstart = $1", onloadstart)
}

func (*HTMLTableCaptionElement) Onmousedown() (onmousedown func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmousedown")
	return onmousedown
}

func (*HTMLTableCaptionElement) SetOnmousedown(onmousedown func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmousedown = $1", onmousedown)
}

func (*HTMLTableCaptionElement) Onmouseenter() (onmouseenter func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseenter")
	return onmouseenter
}

func (*HTMLTableCaptionElement) SetOnmouseenter(onmouseenter func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseenter = $1", onmouseenter)
}

func (*HTMLTableCaptionElement) Onmouseleave() (onmouseleave func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseleave")
	return onmouseleave
}

func (*HTMLTableCaptionElement) SetOnmouseleave(onmouseleave func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseleave = $1", onmouseleave)
}

func (*HTMLTableCaptionElement) Onmousemove() (onmousemove func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmousemove")
	return onmousemove
}

func (*HTMLTableCaptionElement) SetOnmousemove(onmousemove func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmousemove = $1", onmousemove)
}

func (*HTMLTableCaptionElement) Onmouseout() (onmouseout func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseout")
	return onmouseout
}

func (*HTMLTableCaptionElement) SetOnmouseout(onmouseout func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseout = $1", onmouseout)
}

func (*HTMLTableCaptionElement) Onmouseover() (onmouseover func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseover")
	return onmouseover
}

func (*HTMLTableCaptionElement) SetOnmouseover(onmouseover func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseover = $1", onmouseover)
}

func (*HTMLTableCaptionElement) Onmouseup() (onmouseup func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseup")
	return onmouseup
}

func (*HTMLTableCaptionElement) SetOnmouseup(onmouseup func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseup = $1", onmouseup)
}

func (*HTMLTableCaptionElement) Onmousewheel() (onmousewheel func(*utils.WheelEvent)) {
	macro.Rewrite("$_.onmousewheel")
	return onmousewheel
}

func (*HTMLTableCaptionElement) SetOnmousewheel(onmousewheel func(*utils.WheelEvent)) {
	macro.Rewrite("$_.onmousewheel = $1", onmousewheel)
}

func (*HTMLTableCaptionElement) Onmscontentzoom() (onmscontentzoom func(ui.UIEvent)) {
	macro.Rewrite("$_.onmscontentzoom")
	return onmscontentzoom
}

func (*HTMLTableCaptionElement) SetOnmscontentzoom(onmscontentzoom func(ui.UIEvent)) {
	macro.Rewrite("$_.onmscontentzoom = $1", onmscontentzoom)
}

func (*HTMLTableCaptionElement) Onmsmanipulationstatechanged() (onmsmanipulationstatechanged func(*ms.MSManipulationEvent)) {
	macro.Rewrite("$_.onmsmanipulationstatechanged")
	return onmsmanipulationstatechanged
}

func (*HTMLTableCaptionElement) SetOnmsmanipulationstatechanged(onmsmanipulationstatechanged func(*ms.MSManipulationEvent)) {
	macro.Rewrite("$_.onmsmanipulationstatechanged = $1", onmsmanipulationstatechanged)
}

func (*HTMLTableCaptionElement) Onpaste() (onpaste func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.onpaste")
	return onpaste
}

func (*HTMLTableCaptionElement) SetOnpaste(onpaste func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.onpaste = $1", onpaste)
}

func (*HTMLTableCaptionElement) Onpause() (onpause func(event.Event)) {
	macro.Rewrite("$_.onpause")
	return onpause
}

func (*HTMLTableCaptionElement) SetOnpause(onpause func(event.Event)) {
	macro.Rewrite("$_.onpause = $1", onpause)
}

func (*HTMLTableCaptionElement) Onplay() (onplay func(event.Event)) {
	macro.Rewrite("$_.onplay")
	return onplay
}

func (*HTMLTableCaptionElement) SetOnplay(onplay func(event.Event)) {
	macro.Rewrite("$_.onplay = $1", onplay)
}

func (*HTMLTableCaptionElement) Onplaying() (onplaying func(event.Event)) {
	macro.Rewrite("$_.onplaying")
	return onplaying
}

func (*HTMLTableCaptionElement) SetOnplaying(onplaying func(event.Event)) {
	macro.Rewrite("$_.onplaying = $1", onplaying)
}

func (*HTMLTableCaptionElement) Onprogress() (onprogress func(event.Event)) {
	macro.Rewrite("$_.onprogress")
	return onprogress
}

func (*HTMLTableCaptionElement) SetOnprogress(onprogress func(event.Event)) {
	macro.Rewrite("$_.onprogress = $1", onprogress)
}

func (*HTMLTableCaptionElement) Onratechange() (onratechange func(event.Event)) {
	macro.Rewrite("$_.onratechange")
	return onratechange
}

func (*HTMLTableCaptionElement) SetOnratechange(onratechange func(event.Event)) {
	macro.Rewrite("$_.onratechange = $1", onratechange)
}

func (*HTMLTableCaptionElement) Onreset() (onreset func(event.Event)) {
	macro.Rewrite("$_.onreset")
	return onreset
}

func (*HTMLTableCaptionElement) SetOnreset(onreset func(event.Event)) {
	macro.Rewrite("$_.onreset = $1", onreset)
}

func (*HTMLTableCaptionElement) Onscroll() (onscroll func(ui.UIEvent)) {
	macro.Rewrite("$_.onscroll")
	return onscroll
}

func (*HTMLTableCaptionElement) SetOnscroll(onscroll func(ui.UIEvent)) {
	macro.Rewrite("$_.onscroll = $1", onscroll)
}

func (*HTMLTableCaptionElement) Onseeked() (onseeked func(event.Event)) {
	macro.Rewrite("$_.onseeked")
	return onseeked
}

func (*HTMLTableCaptionElement) SetOnseeked(onseeked func(event.Event)) {
	macro.Rewrite("$_.onseeked = $1", onseeked)
}

func (*HTMLTableCaptionElement) Onseeking() (onseeking func(event.Event)) {
	macro.Rewrite("$_.onseeking")
	return onseeking
}

func (*HTMLTableCaptionElement) SetOnseeking(onseeking func(event.Event)) {
	macro.Rewrite("$_.onseeking = $1", onseeking)
}

func (*HTMLTableCaptionElement) Onselect() (onselect func(ui.UIEvent)) {
	macro.Rewrite("$_.onselect")
	return onselect
}

func (*HTMLTableCaptionElement) SetOnselect(onselect func(ui.UIEvent)) {
	macro.Rewrite("$_.onselect = $1", onselect)
}

func (*HTMLTableCaptionElement) Onselectstart() (onselectstart func(event.Event)) {
	macro.Rewrite("$_.onselectstart")
	return onselectstart
}

func (*HTMLTableCaptionElement) SetOnselectstart(onselectstart func(event.Event)) {
	macro.Rewrite("$_.onselectstart = $1", onselectstart)
}

func (*HTMLTableCaptionElement) Onstalled() (onstalled func(event.Event)) {
	macro.Rewrite("$_.onstalled")
	return onstalled
}

func (*HTMLTableCaptionElement) SetOnstalled(onstalled func(event.Event)) {
	macro.Rewrite("$_.onstalled = $1", onstalled)
}

func (*HTMLTableCaptionElement) Onsubmit() (onsubmit func(event.Event)) {
	macro.Rewrite("$_.onsubmit")
	return onsubmit
}

func (*HTMLTableCaptionElement) SetOnsubmit(onsubmit func(event.Event)) {
	macro.Rewrite("$_.onsubmit = $1", onsubmit)
}

func (*HTMLTableCaptionElement) Onsuspend() (onsuspend func(event.Event)) {
	macro.Rewrite("$_.onsuspend")
	return onsuspend
}

func (*HTMLTableCaptionElement) SetOnsuspend(onsuspend func(event.Event)) {
	macro.Rewrite("$_.onsuspend = $1", onsuspend)
}

func (*HTMLTableCaptionElement) Ontimeupdate() (ontimeupdate func(event.Event)) {
	macro.Rewrite("$_.ontimeupdate")
	return ontimeupdate
}

func (*HTMLTableCaptionElement) SetOntimeupdate(ontimeupdate func(event.Event)) {
	macro.Rewrite("$_.ontimeupdate = $1", ontimeupdate)
}

func (*HTMLTableCaptionElement) Onvolumechange() (onvolumechange func(event.Event)) {
	macro.Rewrite("$_.onvolumechange")
	return onvolumechange
}

func (*HTMLTableCaptionElement) SetOnvolumechange(onvolumechange func(event.Event)) {
	macro.Rewrite("$_.onvolumechange = $1", onvolumechange)
}

func (*HTMLTableCaptionElement) Onwaiting() (onwaiting func(event.Event)) {
	macro.Rewrite("$_.onwaiting")
	return onwaiting
}

func (*HTMLTableCaptionElement) SetOnwaiting(onwaiting func(event.Event)) {
	macro.Rewrite("$_.onwaiting = $1", onwaiting)
}

func (*HTMLTableCaptionElement) OuterText() (outerText string) {
	macro.Rewrite("$_.outerText")
	return outerText
}

func (*HTMLTableCaptionElement) SetOuterText(outerText string) {
	macro.Rewrite("$_.outerText = $1", outerText)
}

func (*HTMLTableCaptionElement) Spellcheck() (spellcheck bool) {
	macro.Rewrite("$_.spellcheck")
	return spellcheck
}

func (*HTMLTableCaptionElement) SetSpellcheck(spellcheck bool) {
	macro.Rewrite("$_.spellcheck = $1", spellcheck)
}

func (*HTMLTableCaptionElement) Style() (style *css.CSSStyleDeclaration) {
	macro.Rewrite("$_.style")
	return style
}

func (*HTMLTableCaptionElement) TabIndex() (tabIndex int8) {
	macro.Rewrite("$_.tabIndex")
	return tabIndex
}

func (*HTMLTableCaptionElement) SetTabIndex(tabIndex int8) {
	macro.Rewrite("$_.tabIndex = $1", tabIndex)
}

func (*HTMLTableCaptionElement) Title() (title string) {
	macro.Rewrite("$_.title")
	return title
}

func (*HTMLTableCaptionElement) SetTitle(title string) {
	macro.Rewrite("$_.title = $1", title)
}

func (*HTMLTableCaptionElement) ClassList() (classList dom.DOMTokenList) {
	macro.Rewrite("$_.classList")
	return classList
}

func (*HTMLTableCaptionElement) ClassName() (className string) {
	macro.Rewrite("$_.className")
	return className
}

func (*HTMLTableCaptionElement) SetClassName(className string) {
	macro.Rewrite("$_.className = $1", className)
}

func (*HTMLTableCaptionElement) ClientHeight() (clientHeight int) {
	macro.Rewrite("$_.clientHeight")
	return clientHeight
}

func (*HTMLTableCaptionElement) ClientLeft() (clientLeft int) {
	macro.Rewrite("$_.clientLeft")
	return clientLeft
}

func (*HTMLTableCaptionElement) ClientTop() (clientTop int) {
	macro.Rewrite("$_.clientTop")
	return clientTop
}

func (*HTMLTableCaptionElement) ClientWidth() (clientWidth int) {
	macro.Rewrite("$_.clientWidth")
	return clientWidth
}

func (*HTMLTableCaptionElement) ID() (id string) {
	macro.Rewrite("$_.id")
	return id
}

func (*HTMLTableCaptionElement) SetID(id string) {
	macro.Rewrite("$_.id = $1", id)
}

func (*HTMLTableCaptionElement) InnerHTML() (innerHTML string) {
	macro.Rewrite("$_.innerHTML")
	return innerHTML
}

func (*HTMLTableCaptionElement) SetInnerHTML(innerHTML string) {
	macro.Rewrite("$_.innerHTML = $1", innerHTML)
}

func (*HTMLTableCaptionElement) MsContentZoomFactor() (msContentZoomFactor float32) {
	macro.Rewrite("$_.msContentZoomFactor")
	return msContentZoomFactor
}

func (*HTMLTableCaptionElement) SetMsContentZoomFactor(msContentZoomFactor float32) {
	macro.Rewrite("$_.msContentZoomFactor = $1", msContentZoomFactor)
}

func (*HTMLTableCaptionElement) MsRegionOverflow() (msRegionOverflow string) {
	macro.Rewrite("$_.msRegionOverflow")
	return msRegionOverflow
}

func (*HTMLTableCaptionElement) Onariarequest() (onariarequest func(event.Event)) {
	macro.Rewrite("$_.onariarequest")
	return onariarequest
}

func (*HTMLTableCaptionElement) SetOnariarequest(onariarequest func(event.Event)) {
	macro.Rewrite("$_.onariarequest = $1", onariarequest)
}

func (*HTMLTableCaptionElement) Oncommand() (oncommand func(event.Event)) {
	macro.Rewrite("$_.oncommand")
	return oncommand
}

func (*HTMLTableCaptionElement) SetOncommand(oncommand func(event.Event)) {
	macro.Rewrite("$_.oncommand = $1", oncommand)
}

func (*HTMLTableCaptionElement) Ongotpointercapture() (ongotpointercapture func(*utils.PointerEvent)) {
	macro.Rewrite("$_.ongotpointercapture")
	return ongotpointercapture
}

func (*HTMLTableCaptionElement) SetOngotpointercapture(ongotpointercapture func(*utils.PointerEvent)) {
	macro.Rewrite("$_.ongotpointercapture = $1", ongotpointercapture)
}

func (*HTMLTableCaptionElement) Onlostpointercapture() (onlostpointercapture func(*utils.PointerEvent)) {
	macro.Rewrite("$_.onlostpointercapture")
	return onlostpointercapture
}

func (*HTMLTableCaptionElement) SetOnlostpointercapture(onlostpointercapture func(*utils.PointerEvent)) {
	macro.Rewrite("$_.onlostpointercapture = $1", onlostpointercapture)
}

func (*HTMLTableCaptionElement) Onmsgesturechange() (onmsgesturechange func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturechange")
	return onmsgesturechange
}

func (*HTMLTableCaptionElement) SetOnmsgesturechange(onmsgesturechange func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturechange = $1", onmsgesturechange)
}

func (*HTMLTableCaptionElement) Onmsgesturedoubletap() (onmsgesturedoubletap func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturedoubletap")
	return onmsgesturedoubletap
}

func (*HTMLTableCaptionElement) SetOnmsgesturedoubletap(onmsgesturedoubletap func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturedoubletap = $1", onmsgesturedoubletap)
}

func (*HTMLTableCaptionElement) Onmsgestureend() (onmsgestureend func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgestureend")
	return onmsgestureend
}

func (*HTMLTableCaptionElement) SetOnmsgestureend(onmsgestureend func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgestureend = $1", onmsgestureend)
}

func (*HTMLTableCaptionElement) Onmsgesturehold() (onmsgesturehold func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturehold")
	return onmsgesturehold
}

func (*HTMLTableCaptionElement) SetOnmsgesturehold(onmsgesturehold func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturehold = $1", onmsgesturehold)
}

func (*HTMLTableCaptionElement) Onmsgesturestart() (onmsgesturestart func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturestart")
	return onmsgesturestart
}

func (*HTMLTableCaptionElement) SetOnmsgesturestart(onmsgesturestart func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturestart = $1", onmsgesturestart)
}

func (*HTMLTableCaptionElement) Onmsgesturetap() (onmsgesturetap func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturetap")
	return onmsgesturetap
}

func (*HTMLTableCaptionElement) SetOnmsgesturetap(onmsgesturetap func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturetap = $1", onmsgesturetap)
}

func (*HTMLTableCaptionElement) Onmsgotpointercapture() (onmsgotpointercapture func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmsgotpointercapture")
	return onmsgotpointercapture
}

func (*HTMLTableCaptionElement) SetOnmsgotpointercapture(onmsgotpointercapture func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmsgotpointercapture = $1", onmsgotpointercapture)
}

func (*HTMLTableCaptionElement) Onmsinertiastart() (onmsinertiastart func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsinertiastart")
	return onmsinertiastart
}

func (*HTMLTableCaptionElement) SetOnmsinertiastart(onmsinertiastart func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsinertiastart = $1", onmsinertiastart)
}

func (*HTMLTableCaptionElement) Onmslostpointercapture() (onmslostpointercapture func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmslostpointercapture")
	return onmslostpointercapture
}

func (*HTMLTableCaptionElement) SetOnmslostpointercapture(onmslostpointercapture func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmslostpointercapture = $1", onmslostpointercapture)
}

func (*HTMLTableCaptionElement) Onmspointercancel() (onmspointercancel func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointercancel")
	return onmspointercancel
}

func (*HTMLTableCaptionElement) SetOnmspointercancel(onmspointercancel func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointercancel = $1", onmspointercancel)
}

func (*HTMLTableCaptionElement) Onmspointerdown() (onmspointerdown func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerdown")
	return onmspointerdown
}

func (*HTMLTableCaptionElement) SetOnmspointerdown(onmspointerdown func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerdown = $1", onmspointerdown)
}

func (*HTMLTableCaptionElement) Onmspointerenter() (onmspointerenter func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerenter")
	return onmspointerenter
}

func (*HTMLTableCaptionElement) SetOnmspointerenter(onmspointerenter func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerenter = $1", onmspointerenter)
}

func (*HTMLTableCaptionElement) Onmspointerleave() (onmspointerleave func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerleave")
	return onmspointerleave
}

func (*HTMLTableCaptionElement) SetOnmspointerleave(onmspointerleave func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerleave = $1", onmspointerleave)
}

func (*HTMLTableCaptionElement) Onmspointermove() (onmspointermove func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointermove")
	return onmspointermove
}

func (*HTMLTableCaptionElement) SetOnmspointermove(onmspointermove func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointermove = $1", onmspointermove)
}

func (*HTMLTableCaptionElement) Onmspointerout() (onmspointerout func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerout")
	return onmspointerout
}

func (*HTMLTableCaptionElement) SetOnmspointerout(onmspointerout func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerout = $1", onmspointerout)
}

func (*HTMLTableCaptionElement) Onmspointerover() (onmspointerover func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerover")
	return onmspointerover
}

func (*HTMLTableCaptionElement) SetOnmspointerover(onmspointerover func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerover = $1", onmspointerover)
}

func (*HTMLTableCaptionElement) Onmspointerup() (onmspointerup func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerup")
	return onmspointerup
}

func (*HTMLTableCaptionElement) SetOnmspointerup(onmspointerup func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerup = $1", onmspointerup)
}

func (*HTMLTableCaptionElement) Ontouchcancel() (ontouchcancel func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchcancel")
	return ontouchcancel
}

func (*HTMLTableCaptionElement) SetOntouchcancel(ontouchcancel func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchcancel = $1", ontouchcancel)
}

func (*HTMLTableCaptionElement) Ontouchend() (ontouchend func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchend")
	return ontouchend
}

func (*HTMLTableCaptionElement) SetOntouchend(ontouchend func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchend = $1", ontouchend)
}

func (*HTMLTableCaptionElement) Ontouchmove() (ontouchmove func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchmove")
	return ontouchmove
}

func (*HTMLTableCaptionElement) SetOntouchmove(ontouchmove func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchmove = $1", ontouchmove)
}

func (*HTMLTableCaptionElement) Ontouchstart() (ontouchstart func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchstart")
	return ontouchstart
}

func (*HTMLTableCaptionElement) SetOntouchstart(ontouchstart func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchstart = $1", ontouchstart)
}

func (*HTMLTableCaptionElement) Onwebkitfullscreenchange() (onwebkitfullscreenchange func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenchange")
	return onwebkitfullscreenchange
}

func (*HTMLTableCaptionElement) SetOnwebkitfullscreenchange(onwebkitfullscreenchange func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenchange = $1", onwebkitfullscreenchange)
}

func (*HTMLTableCaptionElement) Onwebkitfullscreenerror() (onwebkitfullscreenerror func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenerror")
	return onwebkitfullscreenerror
}

func (*HTMLTableCaptionElement) SetOnwebkitfullscreenerror(onwebkitfullscreenerror func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenerror = $1", onwebkitfullscreenerror)
}

func (*HTMLTableCaptionElement) OuterHTML() (outerHTML string) {
	macro.Rewrite("$_.outerHTML")
	return outerHTML
}

func (*HTMLTableCaptionElement) SetOuterHTML(outerHTML string) {
	macro.Rewrite("$_.outerHTML = $1", outerHTML)
}

func (*HTMLTableCaptionElement) Prefix() (prefix string) {
	macro.Rewrite("$_.prefix")
	return prefix
}

func (*HTMLTableCaptionElement) ScrollHeight() (scrollHeight int) {
	macro.Rewrite("$_.scrollHeight")
	return scrollHeight
}

func (*HTMLTableCaptionElement) ScrollLeft() (scrollLeft int) {
	macro.Rewrite("$_.scrollLeft")
	return scrollLeft
}

func (*HTMLTableCaptionElement) SetScrollLeft(scrollLeft int) {
	macro.Rewrite("$_.scrollLeft = $1", scrollLeft)
}

func (*HTMLTableCaptionElement) ScrollTop() (scrollTop int) {
	macro.Rewrite("$_.scrollTop")
	return scrollTop
}

func (*HTMLTableCaptionElement) SetScrollTop(scrollTop int) {
	macro.Rewrite("$_.scrollTop = $1", scrollTop)
}

func (*HTMLTableCaptionElement) ScrollWidth() (scrollWidth int) {
	macro.Rewrite("$_.scrollWidth")
	return scrollWidth
}

func (*HTMLTableCaptionElement) TagName() (tagName string) {
	macro.Rewrite("$_.tagName")
	return tagName
}

func (*HTMLTableCaptionElement) Onpointercancel() (onpointercancel func(event.Event)) {
	macro.Rewrite("$_.onpointercancel")
	return onpointercancel
}

func (*HTMLTableCaptionElement) SetOnpointercancel(onpointercancel func(event.Event)) {
	macro.Rewrite("$_.onpointercancel = $1", onpointercancel)
}

func (*HTMLTableCaptionElement) Onpointerdown() (onpointerdown func(event.Event)) {
	macro.Rewrite("$_.onpointerdown")
	return onpointerdown
}

func (*HTMLTableCaptionElement) SetOnpointerdown(onpointerdown func(event.Event)) {
	macro.Rewrite("$_.onpointerdown = $1", onpointerdown)
}

func (*HTMLTableCaptionElement) Onpointerenter() (onpointerenter func(event.Event)) {
	macro.Rewrite("$_.onpointerenter")
	return onpointerenter
}

func (*HTMLTableCaptionElement) SetOnpointerenter(onpointerenter func(event.Event)) {
	macro.Rewrite("$_.onpointerenter = $1", onpointerenter)
}

func (*HTMLTableCaptionElement) Onpointerleave() (onpointerleave func(event.Event)) {
	macro.Rewrite("$_.onpointerleave")
	return onpointerleave
}

func (*HTMLTableCaptionElement) SetOnpointerleave(onpointerleave func(event.Event)) {
	macro.Rewrite("$_.onpointerleave = $1", onpointerleave)
}

func (*HTMLTableCaptionElement) Onpointermove() (onpointermove func(event.Event)) {
	macro.Rewrite("$_.onpointermove")
	return onpointermove
}

func (*HTMLTableCaptionElement) SetOnpointermove(onpointermove func(event.Event)) {
	macro.Rewrite("$_.onpointermove = $1", onpointermove)
}

func (*HTMLTableCaptionElement) Onpointerout() (onpointerout func(event.Event)) {
	macro.Rewrite("$_.onpointerout")
	return onpointerout
}

func (*HTMLTableCaptionElement) SetOnpointerout(onpointerout func(event.Event)) {
	macro.Rewrite("$_.onpointerout = $1", onpointerout)
}

func (*HTMLTableCaptionElement) Onpointerover() (onpointerover func(event.Event)) {
	macro.Rewrite("$_.onpointerover")
	return onpointerover
}

func (*HTMLTableCaptionElement) SetOnpointerover(onpointerover func(event.Event)) {
	macro.Rewrite("$_.onpointerover = $1", onpointerover)
}

func (*HTMLTableCaptionElement) Onpointerup() (onpointerup func(event.Event)) {
	macro.Rewrite("$_.onpointerup")
	return onpointerup
}

func (*HTMLTableCaptionElement) SetOnpointerup(onpointerup func(event.Event)) {
	macro.Rewrite("$_.onpointerup = $1", onpointerup)
}

func (*HTMLTableCaptionElement) Onwheel() (onwheel func(event.Event)) {
	macro.Rewrite("$_.onwheel")
	return onwheel
}

func (*HTMLTableCaptionElement) SetOnwheel(onwheel func(event.Event)) {
	macro.Rewrite("$_.onwheel = $1", onwheel)
}

func (*HTMLTableCaptionElement) ChildElementCount() (childElementCount uint) {
	macro.Rewrite("$_.childElementCount")
	return childElementCount
}

func (*HTMLTableCaptionElement) FirstElementChild() (firstElementChild element.Element) {
	macro.Rewrite("$_.firstElementChild")
	return firstElementChild
}

func (*HTMLTableCaptionElement) LastElementChild() (lastElementChild element.Element) {
	macro.Rewrite("$_.lastElementChild")
	return lastElementChild
}

func (*HTMLTableCaptionElement) NextElementSibling() (nextElementSibling element.Element) {
	macro.Rewrite("$_.nextElementSibling")
	return nextElementSibling
}

func (*HTMLTableCaptionElement) PreviousElementSibling() (previousElementSibling element.Element) {
	macro.Rewrite("$_.previousElementSibling")
	return previousElementSibling
}

func (*HTMLTableCaptionElement) Attributes() (attributes *dom.NamedNodeMap) {
	macro.Rewrite("$_.attributes")
	return attributes
}

func (*HTMLTableCaptionElement) BaseURI() (baseURI string) {
	macro.Rewrite("$_.baseURI")
	return baseURI
}

func (*HTMLTableCaptionElement) ChildNodes() (childNodes *dom.NodeList) {
	macro.Rewrite("$_.childNodes")
	return childNodes
}

func (*HTMLTableCaptionElement) FirstChild() (firstChild dom.Node) {
	macro.Rewrite("$_.firstChild")
	return firstChild
}

func (*HTMLTableCaptionElement) LastChild() (lastChild dom.Node) {
	macro.Rewrite("$_.lastChild")
	return lastChild
}

func (*HTMLTableCaptionElement) LocalName() (localName string) {
	macro.Rewrite("$_.localName")
	return localName
}

func (*HTMLTableCaptionElement) NamespaceURI() (namespaceURI string) {
	macro.Rewrite("$_.namespaceURI")
	return namespaceURI
}

func (*HTMLTableCaptionElement) NextSibling() (nextSibling dom.Node) {
	macro.Rewrite("$_.nextSibling")
	return nextSibling
}

func (*HTMLTableCaptionElement) NodeName() (nodeName string) {
	macro.Rewrite("$_.nodeName")
	return nodeName
}

func (*HTMLTableCaptionElement) NodeType() (nodeType uint8) {
	macro.Rewrite("$_.nodeType")
	return nodeType
}

func (*HTMLTableCaptionElement) NodeValue() (nodeValue string) {
	macro.Rewrite("$_.nodeValue")
	return nodeValue
}

func (*HTMLTableCaptionElement) SetNodeValue(nodeValue string) {
	macro.Rewrite("$_.nodeValue = $1", nodeValue)
}

func (*HTMLTableCaptionElement) OwnerDocument() (ownerDocument document.Document) {
	macro.Rewrite("$_.ownerDocument")
	return ownerDocument
}

func (*HTMLTableCaptionElement) ParentElement() (parentElement HTMLElement) {
	macro.Rewrite("$_.parentElement")
	return parentElement
}

func (*HTMLTableCaptionElement) ParentNode() (parentNode dom.Node) {
	macro.Rewrite("$_.parentNode")
	return parentNode
}

func (*HTMLTableCaptionElement) PreviousSibling() (previousSibling dom.Node) {
	macro.Rewrite("$_.previousSibling")
	return previousSibling
}

func (*HTMLTableCaptionElement) TextContent() (textContent string) {
	macro.Rewrite("$_.textContent")
	return textContent
}

func (*HTMLTableCaptionElement) SetTextContent(textContent string) {
	macro.Rewrite("$_.textContent = $1", textContent)
}
