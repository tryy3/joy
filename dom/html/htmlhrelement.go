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

var _ HTMLElement = (*HTMLHRElement)(nil)
var _ element.Element = (*HTMLHRElement)(nil)
var _ utils.GlobalEventHandlers = (*HTMLHRElement)(nil)
var _ element.ElementTraversal = (*HTMLHRElement)(nil)
var _ dom.NodeSelector = (*HTMLHRElement)(nil)
var _ dom.ChildNode = (*HTMLHRElement)(nil)
var _ dom.Node = (*HTMLHRElement)(nil)
var _ event.EventTarget = (*HTMLHRElement)(nil)

type HTMLHRElement struct {
}

func (*HTMLHRElement) Blur() {
	macro.Rewrite("$_.blur()")
}

func (*HTMLHRElement) Click() {
	macro.Rewrite("$_.click()")
}

func (*HTMLHRElement) DragDrop() (b bool) {
	macro.Rewrite("$_.dragDrop()")
	return b
}

func (*HTMLHRElement) Focus() {
	macro.Rewrite("$_.focus()")
}

func (*HTMLHRElement) GetElementsByClassName(classNames string) (w *dom.NodeList) {
	macro.Rewrite("$_.getElementsByClassName($1)", classNames)
	return w
}

func (*HTMLHRElement) InsertAdjacentElement(position string, insertedElement element.Element) (w element.Element) {
	macro.Rewrite("$_.insertAdjacentElement($1, $2)", position, insertedElement)
	return w
}

func (*HTMLHRElement) InsertAdjacentHTML(where string, html string) {
	macro.Rewrite("$_.insertAdjacentHTML($1, $2)", where, html)
}

func (*HTMLHRElement) InsertAdjacentText(where string, text string) {
	macro.Rewrite("$_.insertAdjacentText($1, $2)", where, text)
}

func (*HTMLHRElement) MsGetInputContext() (w *ms.MSInputMethodContext) {
	macro.Rewrite("$_.msGetInputContext()")
	return w
}

func (*HTMLHRElement) ScrollIntoView(top *bool) {
	macro.Rewrite("$_.scrollIntoView($1)", top)
}

func (*HTMLHRElement) GetAttribute(qualifiedName string) (s string) {
	macro.Rewrite("$_.getAttribute($1)", qualifiedName)
	return s
}

func (*HTMLHRElement) GetAttributeNode(name string) (w *dom.Attr) {
	macro.Rewrite("$_.getAttributeNode($1)", name)
	return w
}

func (*HTMLHRElement) GetAttributeNodeNS(namespaceURI string, localName string) (w *dom.Attr) {
	macro.Rewrite("$_.getAttributeNodeNS($1, $2)", namespaceURI, localName)
	return w
}

func (*HTMLHRElement) GetAttributeNS(namespaceURI string, localName string) (s string) {
	macro.Rewrite("$_.getAttributeNS($1, $2)", namespaceURI, localName)
	return s
}

func (*HTMLHRElement) GetBoundingClientRect() (c *dom.ClientRect) {
	macro.Rewrite("$_.getBoundingClientRect()")
	return c
}

func (*HTMLHRElement) GetClientRects() (c *dom.ClientRectList) {
	macro.Rewrite("$_.getClientRects()")
	return c
}

func (*HTMLHRElement) GetElementsByTagName(name string) (w *dom.NodeList) {
	macro.Rewrite("$_.getElementsByTagName($1)", name)
	return w
}

func (*HTMLHRElement) GetElementsByTagNameNS(namespaceURI string, localName string) (w *dom.NodeList) {
	macro.Rewrite("$_.getElementsByTagNameNS($1, $2)", namespaceURI, localName)
	return w
}

func (*HTMLHRElement) HasAttribute(name string) (b bool) {
	macro.Rewrite("$_.hasAttribute($1)", name)
	return b
}

func (*HTMLHRElement) HasAttributeNS(namespaceURI string, localName string) (b bool) {
	macro.Rewrite("$_.hasAttributeNS($1, $2)", namespaceURI, localName)
	return b
}

func (*HTMLHRElement) MsGetRegionContent() (w *ms.MSRangeCollection) {
	macro.Rewrite("$_.msGetRegionContent()")
	return w
}

func (*HTMLHRElement) MsGetUntransformedBounds() (c *dom.ClientRect) {
	macro.Rewrite("$_.msGetUntransformedBounds()")
	return c
}

func (*HTMLHRElement) MsMatchesSelector(selectors string) (b bool) {
	macro.Rewrite("$_.msMatchesSelector($1)", selectors)
	return b
}

func (*HTMLHRElement) MsReleasePointerCapture(pointerId int) {
	macro.Rewrite("$_.msReleasePointerCapture($1)", pointerId)
}

func (*HTMLHRElement) MsSetPointerCapture(pointerId int) {
	macro.Rewrite("$_.msSetPointerCapture($1)", pointerId)
}

func (*HTMLHRElement) MsZoomTo(args *ms.MsZoomToOptions) {
	macro.Rewrite("$_.msZoomTo($1)", args)
}

func (*HTMLHRElement) ReleasePointerCapture(pointerId int) {
	macro.Rewrite("$_.releasePointerCapture($1)", pointerId)
}

func (*HTMLHRElement) RemoveAttribute(qualifiedName string) {
	macro.Rewrite("$_.removeAttribute($1)", qualifiedName)
}

func (*HTMLHRElement) RemoveAttributeNode(oldAttr *dom.Attr) (w *dom.Attr) {
	macro.Rewrite("$_.removeAttributeNode($1)", oldAttr)
	return w
}

func (*HTMLHRElement) RemoveAttributeNS(namespaceURI string, localName string) {
	macro.Rewrite("$_.removeAttributeNS($1, $2)", namespaceURI, localName)
}

func (*HTMLHRElement) RequestFullscreen() {
	macro.Rewrite("$_.requestFullscreen()")
}

func (*HTMLHRElement) RequestPointerLock() {
	macro.Rewrite("$_.requestPointerLock()")
}

func (*HTMLHRElement) SetAttribute(qualifiedName string, value string) {
	macro.Rewrite("$_.setAttribute($1, $2)", qualifiedName, value)
}

func (*HTMLHRElement) SetAttributeNode(newAttr *dom.Attr) (w *dom.Attr) {
	macro.Rewrite("$_.setAttributeNode($1)", newAttr)
	return w
}

func (*HTMLHRElement) SetAttributeNodeNS(newAttr *dom.Attr) (w *dom.Attr) {
	macro.Rewrite("$_.setAttributeNodeNS($1)", newAttr)
	return w
}

func (*HTMLHRElement) SetAttributeNS(namespaceURI string, qualifiedName string, value string) {
	macro.Rewrite("$_.setAttributeNS($1, $2, $3)", namespaceURI, qualifiedName, value)
}

func (*HTMLHRElement) SetPointerCapture(pointerId int) {
	macro.Rewrite("$_.setPointerCapture($1)", pointerId)
}

func (*HTMLHRElement) WebkitMatchesSelector(selectors string) (b bool) {
	macro.Rewrite("$_.webkitMatchesSelector($1)", selectors)
	return b
}

func (*HTMLHRElement) WebkitRequestFullscreen() {
	macro.Rewrite("$_.webkitRequestFullscreen()")
}

func (*HTMLHRElement) WebkitRequestFullScreen() {
	macro.Rewrite("$_.webkitRequestFullScreen()")
}

func (*HTMLHRElement) QuerySelector(selectors string) (w element.Element) {
	macro.Rewrite("$_.querySelector($1)", selectors)
	return w
}

func (*HTMLHRElement) QuerySelectorAll(selectors string) (w *dom.NodeList) {
	macro.Rewrite("$_.querySelectorAll($1)", selectors)
	return w
}

func (*HTMLHRElement) AppendChild(newChild dom.Node) (w dom.Node) {
	macro.Rewrite("$_.appendChild($1)", newChild)
	return w
}

func (*HTMLHRElement) CloneNode(deep *bool) (w dom.Node) {
	macro.Rewrite("$_.cloneNode($1)", deep)
	return w
}

func (*HTMLHRElement) CompareDocumentPosition(other dom.Node) (u uint8) {
	macro.Rewrite("$_.compareDocumentPosition($1)", other)
	return u
}

func (*HTMLHRElement) Contains(child dom.Node) (b bool) {
	macro.Rewrite("$_.contains($1)", child)
	return b
}

func (*HTMLHRElement) HasAttributes() (b bool) {
	macro.Rewrite("$_.hasAttributes()")
	return b
}

func (*HTMLHRElement) HasChildNodes() (b bool) {
	macro.Rewrite("$_.hasChildNodes()")
	return b
}

func (*HTMLHRElement) InsertBefore(newChild dom.Node, refChild *dom.Node) (w dom.Node) {
	macro.Rewrite("$_.insertBefore($1, $2)", newChild, refChild)
	return w
}

func (*HTMLHRElement) IsDefaultNamespace(namespaceURI string) (b bool) {
	macro.Rewrite("$_.isDefaultNamespace($1)", namespaceURI)
	return b
}

func (*HTMLHRElement) IsEqualNode(arg dom.Node) (b bool) {
	macro.Rewrite("$_.isEqualNode($1)", arg)
	return b
}

func (*HTMLHRElement) IsSameNode(other dom.Node) (b bool) {
	macro.Rewrite("$_.isSameNode($1)", other)
	return b
}

func (*HTMLHRElement) LookupNamespaceURI(prefix string) (s string) {
	macro.Rewrite("$_.lookupNamespaceURI($1)", prefix)
	return s
}

func (*HTMLHRElement) LookupPrefix(namespaceURI string) (s string) {
	macro.Rewrite("$_.lookupPrefix($1)", namespaceURI)
	return s
}

func (*HTMLHRElement) Normalize() {
	macro.Rewrite("$_.normalize()")
}

func (*HTMLHRElement) RemoveChild(oldChild dom.Node) (w dom.Node) {
	macro.Rewrite("$_.removeChild($1)", oldChild)
	return w
}

func (*HTMLHRElement) ReplaceChild(newChild dom.Node, oldChild dom.Node) (w dom.Node) {
	macro.Rewrite("$_.replaceChild($1, $2)", newChild, oldChild)
	return w
}

func (*HTMLHRElement) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*HTMLHRElement) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*HTMLHRElement) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*HTMLHRElement) Align() (align string) {
	macro.Rewrite("$_.align")
	return align
}

func (*HTMLHRElement) SetAlign(align string) {
	macro.Rewrite("$_.align = $1", align)
}

func (*HTMLHRElement) NoShade() (noShade bool) {
	macro.Rewrite("$_.noShade")
	return noShade
}

func (*HTMLHRElement) SetNoShade(noShade bool) {
	macro.Rewrite("$_.noShade = $1", noShade)
}

func (*HTMLHRElement) Width() (width int) {
	macro.Rewrite("$_.width")
	return width
}

func (*HTMLHRElement) SetWidth(width int) {
	macro.Rewrite("$_.width = $1", width)
}

func (*HTMLHRElement) Color() (color string) {
	macro.Rewrite("$_.color")
	return color
}

func (*HTMLHRElement) SetColor(color string) {
	macro.Rewrite("$_.color = $1", color)
}

func (*HTMLHRElement) Size() (size int) {
	macro.Rewrite("$_.size")
	return size
}

func (*HTMLHRElement) SetSize(size int) {
	macro.Rewrite("$_.size = $1", size)
}

func (*HTMLHRElement) AccessKey() (accessKey string) {
	macro.Rewrite("$_.accessKey")
	return accessKey
}

func (*HTMLHRElement) SetAccessKey(accessKey string) {
	macro.Rewrite("$_.accessKey = $1", accessKey)
}

func (*HTMLHRElement) Children() (children HTMLCollection) {
	macro.Rewrite("$_.children")
	return children
}

func (*HTMLHRElement) ContentEditable() (contentEditable string) {
	macro.Rewrite("$_.contentEditable")
	return contentEditable
}

func (*HTMLHRElement) SetContentEditable(contentEditable string) {
	macro.Rewrite("$_.contentEditable = $1", contentEditable)
}

func (*HTMLHRElement) Dataset() (dataset *dom.DOMStringMap) {
	macro.Rewrite("$_.dataset")
	return dataset
}

func (*HTMLHRElement) Dir() (dir string) {
	macro.Rewrite("$_.dir")
	return dir
}

func (*HTMLHRElement) SetDir(dir string) {
	macro.Rewrite("$_.dir = $1", dir)
}

func (*HTMLHRElement) Draggable() (draggable bool) {
	macro.Rewrite("$_.draggable")
	return draggable
}

func (*HTMLHRElement) SetDraggable(draggable bool) {
	macro.Rewrite("$_.draggable = $1", draggable)
}

func (*HTMLHRElement) Hidden() (hidden bool) {
	macro.Rewrite("$_.hidden")
	return hidden
}

func (*HTMLHRElement) SetHidden(hidden bool) {
	macro.Rewrite("$_.hidden = $1", hidden)
}

func (*HTMLHRElement) HideFocus() (hideFocus bool) {
	macro.Rewrite("$_.hideFocus")
	return hideFocus
}

func (*HTMLHRElement) SetHideFocus(hideFocus bool) {
	macro.Rewrite("$_.hideFocus = $1", hideFocus)
}

func (*HTMLHRElement) InnerText() (innerText string) {
	macro.Rewrite("$_.innerText")
	return innerText
}

func (*HTMLHRElement) SetInnerText(innerText string) {
	macro.Rewrite("$_.innerText = $1", innerText)
}

func (*HTMLHRElement) IsContentEditable() (isContentEditable bool) {
	macro.Rewrite("$_.isContentEditable")
	return isContentEditable
}

func (*HTMLHRElement) Lang() (lang string) {
	macro.Rewrite("$_.lang")
	return lang
}

func (*HTMLHRElement) SetLang(lang string) {
	macro.Rewrite("$_.lang = $1", lang)
}

func (*HTMLHRElement) OffsetHeight() (offsetHeight int) {
	macro.Rewrite("$_.offsetHeight")
	return offsetHeight
}

func (*HTMLHRElement) OffsetLeft() (offsetLeft int) {
	macro.Rewrite("$_.offsetLeft")
	return offsetLeft
}

func (*HTMLHRElement) OffsetParent() (offsetParent element.Element) {
	macro.Rewrite("$_.offsetParent")
	return offsetParent
}

func (*HTMLHRElement) OffsetTop() (offsetTop int) {
	macro.Rewrite("$_.offsetTop")
	return offsetTop
}

func (*HTMLHRElement) OffsetWidth() (offsetWidth int) {
	macro.Rewrite("$_.offsetWidth")
	return offsetWidth
}

func (*HTMLHRElement) Onabort() (onabort func(event.Event)) {
	macro.Rewrite("$_.onabort")
	return onabort
}

func (*HTMLHRElement) SetOnabort(onabort func(event.Event)) {
	macro.Rewrite("$_.onabort = $1", onabort)
}

func (*HTMLHRElement) Onactivate() (onactivate func(ui.UIEvent)) {
	macro.Rewrite("$_.onactivate")
	return onactivate
}

func (*HTMLHRElement) SetOnactivate(onactivate func(ui.UIEvent)) {
	macro.Rewrite("$_.onactivate = $1", onactivate)
}

func (*HTMLHRElement) Onbeforeactivate() (onbeforeactivate func(ui.UIEvent)) {
	macro.Rewrite("$_.onbeforeactivate")
	return onbeforeactivate
}

func (*HTMLHRElement) SetOnbeforeactivate(onbeforeactivate func(ui.UIEvent)) {
	macro.Rewrite("$_.onbeforeactivate = $1", onbeforeactivate)
}

func (*HTMLHRElement) Onbeforecopy() (onbeforecopy func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.onbeforecopy")
	return onbeforecopy
}

func (*HTMLHRElement) SetOnbeforecopy(onbeforecopy func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.onbeforecopy = $1", onbeforecopy)
}

func (*HTMLHRElement) Onbeforecut() (onbeforecut func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.onbeforecut")
	return onbeforecut
}

func (*HTMLHRElement) SetOnbeforecut(onbeforecut func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.onbeforecut = $1", onbeforecut)
}

func (*HTMLHRElement) Onbeforedeactivate() (onbeforedeactivate func(ui.UIEvent)) {
	macro.Rewrite("$_.onbeforedeactivate")
	return onbeforedeactivate
}

func (*HTMLHRElement) SetOnbeforedeactivate(onbeforedeactivate func(ui.UIEvent)) {
	macro.Rewrite("$_.onbeforedeactivate = $1", onbeforedeactivate)
}

func (*HTMLHRElement) Onbeforepaste() (onbeforepaste func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.onbeforepaste")
	return onbeforepaste
}

func (*HTMLHRElement) SetOnbeforepaste(onbeforepaste func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.onbeforepaste = $1", onbeforepaste)
}

func (*HTMLHRElement) Onblur() (onblur func(*utils.FocusEvent)) {
	macro.Rewrite("$_.onblur")
	return onblur
}

func (*HTMLHRElement) SetOnblur(onblur func(*utils.FocusEvent)) {
	macro.Rewrite("$_.onblur = $1", onblur)
}

func (*HTMLHRElement) Oncanplay() (oncanplay func(event.Event)) {
	macro.Rewrite("$_.oncanplay")
	return oncanplay
}

func (*HTMLHRElement) SetOncanplay(oncanplay func(event.Event)) {
	macro.Rewrite("$_.oncanplay = $1", oncanplay)
}

func (*HTMLHRElement) Oncanplaythrough() (oncanplaythrough func(event.Event)) {
	macro.Rewrite("$_.oncanplaythrough")
	return oncanplaythrough
}

func (*HTMLHRElement) SetOncanplaythrough(oncanplaythrough func(event.Event)) {
	macro.Rewrite("$_.oncanplaythrough = $1", oncanplaythrough)
}

func (*HTMLHRElement) Onchange() (onchange func(event.Event)) {
	macro.Rewrite("$_.onchange")
	return onchange
}

func (*HTMLHRElement) SetOnchange(onchange func(event.Event)) {
	macro.Rewrite("$_.onchange = $1", onchange)
}

func (*HTMLHRElement) Onclick() (onclick func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onclick")
	return onclick
}

func (*HTMLHRElement) SetOnclick(onclick func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onclick = $1", onclick)
}

func (*HTMLHRElement) Oncontextmenu() (oncontextmenu func(*utils.PointerEvent)) {
	macro.Rewrite("$_.oncontextmenu")
	return oncontextmenu
}

func (*HTMLHRElement) SetOncontextmenu(oncontextmenu func(*utils.PointerEvent)) {
	macro.Rewrite("$_.oncontextmenu = $1", oncontextmenu)
}

func (*HTMLHRElement) Oncopy() (oncopy func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.oncopy")
	return oncopy
}

func (*HTMLHRElement) SetOncopy(oncopy func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.oncopy = $1", oncopy)
}

func (*HTMLHRElement) Oncuechange() (oncuechange func(event.Event)) {
	macro.Rewrite("$_.oncuechange")
	return oncuechange
}

func (*HTMLHRElement) SetOncuechange(oncuechange func(event.Event)) {
	macro.Rewrite("$_.oncuechange = $1", oncuechange)
}

func (*HTMLHRElement) Oncut() (oncut func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.oncut")
	return oncut
}

func (*HTMLHRElement) SetOncut(oncut func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.oncut = $1", oncut)
}

func (*HTMLHRElement) Ondblclick() (ondblclick func(mouse.MouseEvent)) {
	macro.Rewrite("$_.ondblclick")
	return ondblclick
}

func (*HTMLHRElement) SetOndblclick(ondblclick func(mouse.MouseEvent)) {
	macro.Rewrite("$_.ondblclick = $1", ondblclick)
}

func (*HTMLHRElement) Ondeactivate() (ondeactivate func(ui.UIEvent)) {
	macro.Rewrite("$_.ondeactivate")
	return ondeactivate
}

func (*HTMLHRElement) SetOndeactivate(ondeactivate func(ui.UIEvent)) {
	macro.Rewrite("$_.ondeactivate = $1", ondeactivate)
}

func (*HTMLHRElement) Ondrag() (ondrag func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondrag")
	return ondrag
}

func (*HTMLHRElement) SetOndrag(ondrag func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondrag = $1", ondrag)
}

func (*HTMLHRElement) Ondragend() (ondragend func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragend")
	return ondragend
}

func (*HTMLHRElement) SetOndragend(ondragend func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragend = $1", ondragend)
}

func (*HTMLHRElement) Ondragenter() (ondragenter func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragenter")
	return ondragenter
}

func (*HTMLHRElement) SetOndragenter(ondragenter func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragenter = $1", ondragenter)
}

func (*HTMLHRElement) Ondragleave() (ondragleave func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragleave")
	return ondragleave
}

func (*HTMLHRElement) SetOndragleave(ondragleave func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragleave = $1", ondragleave)
}

func (*HTMLHRElement) Ondragover() (ondragover func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragover")
	return ondragover
}

func (*HTMLHRElement) SetOndragover(ondragover func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragover = $1", ondragover)
}

func (*HTMLHRElement) Ondragstart() (ondragstart func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragstart")
	return ondragstart
}

func (*HTMLHRElement) SetOndragstart(ondragstart func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragstart = $1", ondragstart)
}

func (*HTMLHRElement) Ondrop() (ondrop func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondrop")
	return ondrop
}

func (*HTMLHRElement) SetOndrop(ondrop func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondrop = $1", ondrop)
}

func (*HTMLHRElement) Ondurationchange() (ondurationchange func(event.Event)) {
	macro.Rewrite("$_.ondurationchange")
	return ondurationchange
}

func (*HTMLHRElement) SetOndurationchange(ondurationchange func(event.Event)) {
	macro.Rewrite("$_.ondurationchange = $1", ondurationchange)
}

func (*HTMLHRElement) Onemptied() (onemptied func(event.Event)) {
	macro.Rewrite("$_.onemptied")
	return onemptied
}

func (*HTMLHRElement) SetOnemptied(onemptied func(event.Event)) {
	macro.Rewrite("$_.onemptied = $1", onemptied)
}

func (*HTMLHRElement) Onended() (onended func(event.Event)) {
	macro.Rewrite("$_.onended")
	return onended
}

func (*HTMLHRElement) SetOnended(onended func(event.Event)) {
	macro.Rewrite("$_.onended = $1", onended)
}

func (*HTMLHRElement) Onerror() (onerror func(event.Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

func (*HTMLHRElement) SetOnerror(onerror func(event.Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

func (*HTMLHRElement) Onfocus() (onfocus func(*utils.FocusEvent)) {
	macro.Rewrite("$_.onfocus")
	return onfocus
}

func (*HTMLHRElement) SetOnfocus(onfocus func(*utils.FocusEvent)) {
	macro.Rewrite("$_.onfocus = $1", onfocus)
}

func (*HTMLHRElement) Oninput() (oninput func(event.Event)) {
	macro.Rewrite("$_.oninput")
	return oninput
}

func (*HTMLHRElement) SetOninput(oninput func(event.Event)) {
	macro.Rewrite("$_.oninput = $1", oninput)
}

func (*HTMLHRElement) Oninvalid() (oninvalid func(event.Event)) {
	macro.Rewrite("$_.oninvalid")
	return oninvalid
}

func (*HTMLHRElement) SetOninvalid(oninvalid func(event.Event)) {
	macro.Rewrite("$_.oninvalid = $1", oninvalid)
}

func (*HTMLHRElement) Onkeydown() (onkeydown func(*utils.KeyboardEvent)) {
	macro.Rewrite("$_.onkeydown")
	return onkeydown
}

func (*HTMLHRElement) SetOnkeydown(onkeydown func(*utils.KeyboardEvent)) {
	macro.Rewrite("$_.onkeydown = $1", onkeydown)
}

func (*HTMLHRElement) Onkeypress() (onkeypress func(*utils.KeyboardEvent)) {
	macro.Rewrite("$_.onkeypress")
	return onkeypress
}

func (*HTMLHRElement) SetOnkeypress(onkeypress func(*utils.KeyboardEvent)) {
	macro.Rewrite("$_.onkeypress = $1", onkeypress)
}

func (*HTMLHRElement) Onkeyup() (onkeyup func(*utils.KeyboardEvent)) {
	macro.Rewrite("$_.onkeyup")
	return onkeyup
}

func (*HTMLHRElement) SetOnkeyup(onkeyup func(*utils.KeyboardEvent)) {
	macro.Rewrite("$_.onkeyup = $1", onkeyup)
}

func (*HTMLHRElement) Onload() (onload func(event.Event)) {
	macro.Rewrite("$_.onload")
	return onload
}

func (*HTMLHRElement) SetOnload(onload func(event.Event)) {
	macro.Rewrite("$_.onload = $1", onload)
}

func (*HTMLHRElement) Onloadeddata() (onloadeddata func(event.Event)) {
	macro.Rewrite("$_.onloadeddata")
	return onloadeddata
}

func (*HTMLHRElement) SetOnloadeddata(onloadeddata func(event.Event)) {
	macro.Rewrite("$_.onloadeddata = $1", onloadeddata)
}

func (*HTMLHRElement) Onloadedmetadata() (onloadedmetadata func(event.Event)) {
	macro.Rewrite("$_.onloadedmetadata")
	return onloadedmetadata
}

func (*HTMLHRElement) SetOnloadedmetadata(onloadedmetadata func(event.Event)) {
	macro.Rewrite("$_.onloadedmetadata = $1", onloadedmetadata)
}

func (*HTMLHRElement) Onloadstart() (onloadstart func(event.Event)) {
	macro.Rewrite("$_.onloadstart")
	return onloadstart
}

func (*HTMLHRElement) SetOnloadstart(onloadstart func(event.Event)) {
	macro.Rewrite("$_.onloadstart = $1", onloadstart)
}

func (*HTMLHRElement) Onmousedown() (onmousedown func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmousedown")
	return onmousedown
}

func (*HTMLHRElement) SetOnmousedown(onmousedown func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmousedown = $1", onmousedown)
}

func (*HTMLHRElement) Onmouseenter() (onmouseenter func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseenter")
	return onmouseenter
}

func (*HTMLHRElement) SetOnmouseenter(onmouseenter func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseenter = $1", onmouseenter)
}

func (*HTMLHRElement) Onmouseleave() (onmouseleave func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseleave")
	return onmouseleave
}

func (*HTMLHRElement) SetOnmouseleave(onmouseleave func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseleave = $1", onmouseleave)
}

func (*HTMLHRElement) Onmousemove() (onmousemove func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmousemove")
	return onmousemove
}

func (*HTMLHRElement) SetOnmousemove(onmousemove func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmousemove = $1", onmousemove)
}

func (*HTMLHRElement) Onmouseout() (onmouseout func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseout")
	return onmouseout
}

func (*HTMLHRElement) SetOnmouseout(onmouseout func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseout = $1", onmouseout)
}

func (*HTMLHRElement) Onmouseover() (onmouseover func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseover")
	return onmouseover
}

func (*HTMLHRElement) SetOnmouseover(onmouseover func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseover = $1", onmouseover)
}

func (*HTMLHRElement) Onmouseup() (onmouseup func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseup")
	return onmouseup
}

func (*HTMLHRElement) SetOnmouseup(onmouseup func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseup = $1", onmouseup)
}

func (*HTMLHRElement) Onmousewheel() (onmousewheel func(*utils.WheelEvent)) {
	macro.Rewrite("$_.onmousewheel")
	return onmousewheel
}

func (*HTMLHRElement) SetOnmousewheel(onmousewheel func(*utils.WheelEvent)) {
	macro.Rewrite("$_.onmousewheel = $1", onmousewheel)
}

func (*HTMLHRElement) Onmscontentzoom() (onmscontentzoom func(ui.UIEvent)) {
	macro.Rewrite("$_.onmscontentzoom")
	return onmscontentzoom
}

func (*HTMLHRElement) SetOnmscontentzoom(onmscontentzoom func(ui.UIEvent)) {
	macro.Rewrite("$_.onmscontentzoom = $1", onmscontentzoom)
}

func (*HTMLHRElement) Onmsmanipulationstatechanged() (onmsmanipulationstatechanged func(*ms.MSManipulationEvent)) {
	macro.Rewrite("$_.onmsmanipulationstatechanged")
	return onmsmanipulationstatechanged
}

func (*HTMLHRElement) SetOnmsmanipulationstatechanged(onmsmanipulationstatechanged func(*ms.MSManipulationEvent)) {
	macro.Rewrite("$_.onmsmanipulationstatechanged = $1", onmsmanipulationstatechanged)
}

func (*HTMLHRElement) Onpaste() (onpaste func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.onpaste")
	return onpaste
}

func (*HTMLHRElement) SetOnpaste(onpaste func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.onpaste = $1", onpaste)
}

func (*HTMLHRElement) Onpause() (onpause func(event.Event)) {
	macro.Rewrite("$_.onpause")
	return onpause
}

func (*HTMLHRElement) SetOnpause(onpause func(event.Event)) {
	macro.Rewrite("$_.onpause = $1", onpause)
}

func (*HTMLHRElement) Onplay() (onplay func(event.Event)) {
	macro.Rewrite("$_.onplay")
	return onplay
}

func (*HTMLHRElement) SetOnplay(onplay func(event.Event)) {
	macro.Rewrite("$_.onplay = $1", onplay)
}

func (*HTMLHRElement) Onplaying() (onplaying func(event.Event)) {
	macro.Rewrite("$_.onplaying")
	return onplaying
}

func (*HTMLHRElement) SetOnplaying(onplaying func(event.Event)) {
	macro.Rewrite("$_.onplaying = $1", onplaying)
}

func (*HTMLHRElement) Onprogress() (onprogress func(event.Event)) {
	macro.Rewrite("$_.onprogress")
	return onprogress
}

func (*HTMLHRElement) SetOnprogress(onprogress func(event.Event)) {
	macro.Rewrite("$_.onprogress = $1", onprogress)
}

func (*HTMLHRElement) Onratechange() (onratechange func(event.Event)) {
	macro.Rewrite("$_.onratechange")
	return onratechange
}

func (*HTMLHRElement) SetOnratechange(onratechange func(event.Event)) {
	macro.Rewrite("$_.onratechange = $1", onratechange)
}

func (*HTMLHRElement) Onreset() (onreset func(event.Event)) {
	macro.Rewrite("$_.onreset")
	return onreset
}

func (*HTMLHRElement) SetOnreset(onreset func(event.Event)) {
	macro.Rewrite("$_.onreset = $1", onreset)
}

func (*HTMLHRElement) Onscroll() (onscroll func(ui.UIEvent)) {
	macro.Rewrite("$_.onscroll")
	return onscroll
}

func (*HTMLHRElement) SetOnscroll(onscroll func(ui.UIEvent)) {
	macro.Rewrite("$_.onscroll = $1", onscroll)
}

func (*HTMLHRElement) Onseeked() (onseeked func(event.Event)) {
	macro.Rewrite("$_.onseeked")
	return onseeked
}

func (*HTMLHRElement) SetOnseeked(onseeked func(event.Event)) {
	macro.Rewrite("$_.onseeked = $1", onseeked)
}

func (*HTMLHRElement) Onseeking() (onseeking func(event.Event)) {
	macro.Rewrite("$_.onseeking")
	return onseeking
}

func (*HTMLHRElement) SetOnseeking(onseeking func(event.Event)) {
	macro.Rewrite("$_.onseeking = $1", onseeking)
}

func (*HTMLHRElement) Onselect() (onselect func(ui.UIEvent)) {
	macro.Rewrite("$_.onselect")
	return onselect
}

func (*HTMLHRElement) SetOnselect(onselect func(ui.UIEvent)) {
	macro.Rewrite("$_.onselect = $1", onselect)
}

func (*HTMLHRElement) Onselectstart() (onselectstart func(event.Event)) {
	macro.Rewrite("$_.onselectstart")
	return onselectstart
}

func (*HTMLHRElement) SetOnselectstart(onselectstart func(event.Event)) {
	macro.Rewrite("$_.onselectstart = $1", onselectstart)
}

func (*HTMLHRElement) Onstalled() (onstalled func(event.Event)) {
	macro.Rewrite("$_.onstalled")
	return onstalled
}

func (*HTMLHRElement) SetOnstalled(onstalled func(event.Event)) {
	macro.Rewrite("$_.onstalled = $1", onstalled)
}

func (*HTMLHRElement) Onsubmit() (onsubmit func(event.Event)) {
	macro.Rewrite("$_.onsubmit")
	return onsubmit
}

func (*HTMLHRElement) SetOnsubmit(onsubmit func(event.Event)) {
	macro.Rewrite("$_.onsubmit = $1", onsubmit)
}

func (*HTMLHRElement) Onsuspend() (onsuspend func(event.Event)) {
	macro.Rewrite("$_.onsuspend")
	return onsuspend
}

func (*HTMLHRElement) SetOnsuspend(onsuspend func(event.Event)) {
	macro.Rewrite("$_.onsuspend = $1", onsuspend)
}

func (*HTMLHRElement) Ontimeupdate() (ontimeupdate func(event.Event)) {
	macro.Rewrite("$_.ontimeupdate")
	return ontimeupdate
}

func (*HTMLHRElement) SetOntimeupdate(ontimeupdate func(event.Event)) {
	macro.Rewrite("$_.ontimeupdate = $1", ontimeupdate)
}

func (*HTMLHRElement) Onvolumechange() (onvolumechange func(event.Event)) {
	macro.Rewrite("$_.onvolumechange")
	return onvolumechange
}

func (*HTMLHRElement) SetOnvolumechange(onvolumechange func(event.Event)) {
	macro.Rewrite("$_.onvolumechange = $1", onvolumechange)
}

func (*HTMLHRElement) Onwaiting() (onwaiting func(event.Event)) {
	macro.Rewrite("$_.onwaiting")
	return onwaiting
}

func (*HTMLHRElement) SetOnwaiting(onwaiting func(event.Event)) {
	macro.Rewrite("$_.onwaiting = $1", onwaiting)
}

func (*HTMLHRElement) OuterText() (outerText string) {
	macro.Rewrite("$_.outerText")
	return outerText
}

func (*HTMLHRElement) SetOuterText(outerText string) {
	macro.Rewrite("$_.outerText = $1", outerText)
}

func (*HTMLHRElement) Spellcheck() (spellcheck bool) {
	macro.Rewrite("$_.spellcheck")
	return spellcheck
}

func (*HTMLHRElement) SetSpellcheck(spellcheck bool) {
	macro.Rewrite("$_.spellcheck = $1", spellcheck)
}

func (*HTMLHRElement) Style() (style *css.CSSStyleDeclaration) {
	macro.Rewrite("$_.style")
	return style
}

func (*HTMLHRElement) TabIndex() (tabIndex int8) {
	macro.Rewrite("$_.tabIndex")
	return tabIndex
}

func (*HTMLHRElement) SetTabIndex(tabIndex int8) {
	macro.Rewrite("$_.tabIndex = $1", tabIndex)
}

func (*HTMLHRElement) Title() (title string) {
	macro.Rewrite("$_.title")
	return title
}

func (*HTMLHRElement) SetTitle(title string) {
	macro.Rewrite("$_.title = $1", title)
}

func (*HTMLHRElement) ClassList() (classList dom.DOMTokenList) {
	macro.Rewrite("$_.classList")
	return classList
}

func (*HTMLHRElement) ClassName() (className string) {
	macro.Rewrite("$_.className")
	return className
}

func (*HTMLHRElement) SetClassName(className string) {
	macro.Rewrite("$_.className = $1", className)
}

func (*HTMLHRElement) ClientHeight() (clientHeight int) {
	macro.Rewrite("$_.clientHeight")
	return clientHeight
}

func (*HTMLHRElement) ClientLeft() (clientLeft int) {
	macro.Rewrite("$_.clientLeft")
	return clientLeft
}

func (*HTMLHRElement) ClientTop() (clientTop int) {
	macro.Rewrite("$_.clientTop")
	return clientTop
}

func (*HTMLHRElement) ClientWidth() (clientWidth int) {
	macro.Rewrite("$_.clientWidth")
	return clientWidth
}

func (*HTMLHRElement) ID() (id string) {
	macro.Rewrite("$_.id")
	return id
}

func (*HTMLHRElement) SetID(id string) {
	macro.Rewrite("$_.id = $1", id)
}

func (*HTMLHRElement) InnerHTML() (innerHTML string) {
	macro.Rewrite("$_.innerHTML")
	return innerHTML
}

func (*HTMLHRElement) SetInnerHTML(innerHTML string) {
	macro.Rewrite("$_.innerHTML = $1", innerHTML)
}

func (*HTMLHRElement) MsContentZoomFactor() (msContentZoomFactor float32) {
	macro.Rewrite("$_.msContentZoomFactor")
	return msContentZoomFactor
}

func (*HTMLHRElement) SetMsContentZoomFactor(msContentZoomFactor float32) {
	macro.Rewrite("$_.msContentZoomFactor = $1", msContentZoomFactor)
}

func (*HTMLHRElement) MsRegionOverflow() (msRegionOverflow string) {
	macro.Rewrite("$_.msRegionOverflow")
	return msRegionOverflow
}

func (*HTMLHRElement) Onariarequest() (onariarequest func(event.Event)) {
	macro.Rewrite("$_.onariarequest")
	return onariarequest
}

func (*HTMLHRElement) SetOnariarequest(onariarequest func(event.Event)) {
	macro.Rewrite("$_.onariarequest = $1", onariarequest)
}

func (*HTMLHRElement) Oncommand() (oncommand func(event.Event)) {
	macro.Rewrite("$_.oncommand")
	return oncommand
}

func (*HTMLHRElement) SetOncommand(oncommand func(event.Event)) {
	macro.Rewrite("$_.oncommand = $1", oncommand)
}

func (*HTMLHRElement) Ongotpointercapture() (ongotpointercapture func(*utils.PointerEvent)) {
	macro.Rewrite("$_.ongotpointercapture")
	return ongotpointercapture
}

func (*HTMLHRElement) SetOngotpointercapture(ongotpointercapture func(*utils.PointerEvent)) {
	macro.Rewrite("$_.ongotpointercapture = $1", ongotpointercapture)
}

func (*HTMLHRElement) Onlostpointercapture() (onlostpointercapture func(*utils.PointerEvent)) {
	macro.Rewrite("$_.onlostpointercapture")
	return onlostpointercapture
}

func (*HTMLHRElement) SetOnlostpointercapture(onlostpointercapture func(*utils.PointerEvent)) {
	macro.Rewrite("$_.onlostpointercapture = $1", onlostpointercapture)
}

func (*HTMLHRElement) Onmsgesturechange() (onmsgesturechange func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturechange")
	return onmsgesturechange
}

func (*HTMLHRElement) SetOnmsgesturechange(onmsgesturechange func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturechange = $1", onmsgesturechange)
}

func (*HTMLHRElement) Onmsgesturedoubletap() (onmsgesturedoubletap func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturedoubletap")
	return onmsgesturedoubletap
}

func (*HTMLHRElement) SetOnmsgesturedoubletap(onmsgesturedoubletap func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturedoubletap = $1", onmsgesturedoubletap)
}

func (*HTMLHRElement) Onmsgestureend() (onmsgestureend func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgestureend")
	return onmsgestureend
}

func (*HTMLHRElement) SetOnmsgestureend(onmsgestureend func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgestureend = $1", onmsgestureend)
}

func (*HTMLHRElement) Onmsgesturehold() (onmsgesturehold func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturehold")
	return onmsgesturehold
}

func (*HTMLHRElement) SetOnmsgesturehold(onmsgesturehold func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturehold = $1", onmsgesturehold)
}

func (*HTMLHRElement) Onmsgesturestart() (onmsgesturestart func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturestart")
	return onmsgesturestart
}

func (*HTMLHRElement) SetOnmsgesturestart(onmsgesturestart func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturestart = $1", onmsgesturestart)
}

func (*HTMLHRElement) Onmsgesturetap() (onmsgesturetap func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturetap")
	return onmsgesturetap
}

func (*HTMLHRElement) SetOnmsgesturetap(onmsgesturetap func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturetap = $1", onmsgesturetap)
}

func (*HTMLHRElement) Onmsgotpointercapture() (onmsgotpointercapture func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmsgotpointercapture")
	return onmsgotpointercapture
}

func (*HTMLHRElement) SetOnmsgotpointercapture(onmsgotpointercapture func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmsgotpointercapture = $1", onmsgotpointercapture)
}

func (*HTMLHRElement) Onmsinertiastart() (onmsinertiastart func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsinertiastart")
	return onmsinertiastart
}

func (*HTMLHRElement) SetOnmsinertiastart(onmsinertiastart func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsinertiastart = $1", onmsinertiastart)
}

func (*HTMLHRElement) Onmslostpointercapture() (onmslostpointercapture func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmslostpointercapture")
	return onmslostpointercapture
}

func (*HTMLHRElement) SetOnmslostpointercapture(onmslostpointercapture func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmslostpointercapture = $1", onmslostpointercapture)
}

func (*HTMLHRElement) Onmspointercancel() (onmspointercancel func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointercancel")
	return onmspointercancel
}

func (*HTMLHRElement) SetOnmspointercancel(onmspointercancel func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointercancel = $1", onmspointercancel)
}

func (*HTMLHRElement) Onmspointerdown() (onmspointerdown func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerdown")
	return onmspointerdown
}

func (*HTMLHRElement) SetOnmspointerdown(onmspointerdown func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerdown = $1", onmspointerdown)
}

func (*HTMLHRElement) Onmspointerenter() (onmspointerenter func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerenter")
	return onmspointerenter
}

func (*HTMLHRElement) SetOnmspointerenter(onmspointerenter func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerenter = $1", onmspointerenter)
}

func (*HTMLHRElement) Onmspointerleave() (onmspointerleave func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerleave")
	return onmspointerleave
}

func (*HTMLHRElement) SetOnmspointerleave(onmspointerleave func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerleave = $1", onmspointerleave)
}

func (*HTMLHRElement) Onmspointermove() (onmspointermove func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointermove")
	return onmspointermove
}

func (*HTMLHRElement) SetOnmspointermove(onmspointermove func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointermove = $1", onmspointermove)
}

func (*HTMLHRElement) Onmspointerout() (onmspointerout func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerout")
	return onmspointerout
}

func (*HTMLHRElement) SetOnmspointerout(onmspointerout func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerout = $1", onmspointerout)
}

func (*HTMLHRElement) Onmspointerover() (onmspointerover func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerover")
	return onmspointerover
}

func (*HTMLHRElement) SetOnmspointerover(onmspointerover func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerover = $1", onmspointerover)
}

func (*HTMLHRElement) Onmspointerup() (onmspointerup func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerup")
	return onmspointerup
}

func (*HTMLHRElement) SetOnmspointerup(onmspointerup func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerup = $1", onmspointerup)
}

func (*HTMLHRElement) Ontouchcancel() (ontouchcancel func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchcancel")
	return ontouchcancel
}

func (*HTMLHRElement) SetOntouchcancel(ontouchcancel func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchcancel = $1", ontouchcancel)
}

func (*HTMLHRElement) Ontouchend() (ontouchend func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchend")
	return ontouchend
}

func (*HTMLHRElement) SetOntouchend(ontouchend func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchend = $1", ontouchend)
}

func (*HTMLHRElement) Ontouchmove() (ontouchmove func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchmove")
	return ontouchmove
}

func (*HTMLHRElement) SetOntouchmove(ontouchmove func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchmove = $1", ontouchmove)
}

func (*HTMLHRElement) Ontouchstart() (ontouchstart func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchstart")
	return ontouchstart
}

func (*HTMLHRElement) SetOntouchstart(ontouchstart func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchstart = $1", ontouchstart)
}

func (*HTMLHRElement) Onwebkitfullscreenchange() (onwebkitfullscreenchange func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenchange")
	return onwebkitfullscreenchange
}

func (*HTMLHRElement) SetOnwebkitfullscreenchange(onwebkitfullscreenchange func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenchange = $1", onwebkitfullscreenchange)
}

func (*HTMLHRElement) Onwebkitfullscreenerror() (onwebkitfullscreenerror func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenerror")
	return onwebkitfullscreenerror
}

func (*HTMLHRElement) SetOnwebkitfullscreenerror(onwebkitfullscreenerror func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenerror = $1", onwebkitfullscreenerror)
}

func (*HTMLHRElement) OuterHTML() (outerHTML string) {
	macro.Rewrite("$_.outerHTML")
	return outerHTML
}

func (*HTMLHRElement) SetOuterHTML(outerHTML string) {
	macro.Rewrite("$_.outerHTML = $1", outerHTML)
}

func (*HTMLHRElement) Prefix() (prefix string) {
	macro.Rewrite("$_.prefix")
	return prefix
}

func (*HTMLHRElement) ScrollHeight() (scrollHeight int) {
	macro.Rewrite("$_.scrollHeight")
	return scrollHeight
}

func (*HTMLHRElement) ScrollLeft() (scrollLeft int) {
	macro.Rewrite("$_.scrollLeft")
	return scrollLeft
}

func (*HTMLHRElement) SetScrollLeft(scrollLeft int) {
	macro.Rewrite("$_.scrollLeft = $1", scrollLeft)
}

func (*HTMLHRElement) ScrollTop() (scrollTop int) {
	macro.Rewrite("$_.scrollTop")
	return scrollTop
}

func (*HTMLHRElement) SetScrollTop(scrollTop int) {
	macro.Rewrite("$_.scrollTop = $1", scrollTop)
}

func (*HTMLHRElement) ScrollWidth() (scrollWidth int) {
	macro.Rewrite("$_.scrollWidth")
	return scrollWidth
}

func (*HTMLHRElement) TagName() (tagName string) {
	macro.Rewrite("$_.tagName")
	return tagName
}

func (*HTMLHRElement) Onpointercancel() (onpointercancel func(event.Event)) {
	macro.Rewrite("$_.onpointercancel")
	return onpointercancel
}

func (*HTMLHRElement) SetOnpointercancel(onpointercancel func(event.Event)) {
	macro.Rewrite("$_.onpointercancel = $1", onpointercancel)
}

func (*HTMLHRElement) Onpointerdown() (onpointerdown func(event.Event)) {
	macro.Rewrite("$_.onpointerdown")
	return onpointerdown
}

func (*HTMLHRElement) SetOnpointerdown(onpointerdown func(event.Event)) {
	macro.Rewrite("$_.onpointerdown = $1", onpointerdown)
}

func (*HTMLHRElement) Onpointerenter() (onpointerenter func(event.Event)) {
	macro.Rewrite("$_.onpointerenter")
	return onpointerenter
}

func (*HTMLHRElement) SetOnpointerenter(onpointerenter func(event.Event)) {
	macro.Rewrite("$_.onpointerenter = $1", onpointerenter)
}

func (*HTMLHRElement) Onpointerleave() (onpointerleave func(event.Event)) {
	macro.Rewrite("$_.onpointerleave")
	return onpointerleave
}

func (*HTMLHRElement) SetOnpointerleave(onpointerleave func(event.Event)) {
	macro.Rewrite("$_.onpointerleave = $1", onpointerleave)
}

func (*HTMLHRElement) Onpointermove() (onpointermove func(event.Event)) {
	macro.Rewrite("$_.onpointermove")
	return onpointermove
}

func (*HTMLHRElement) SetOnpointermove(onpointermove func(event.Event)) {
	macro.Rewrite("$_.onpointermove = $1", onpointermove)
}

func (*HTMLHRElement) Onpointerout() (onpointerout func(event.Event)) {
	macro.Rewrite("$_.onpointerout")
	return onpointerout
}

func (*HTMLHRElement) SetOnpointerout(onpointerout func(event.Event)) {
	macro.Rewrite("$_.onpointerout = $1", onpointerout)
}

func (*HTMLHRElement) Onpointerover() (onpointerover func(event.Event)) {
	macro.Rewrite("$_.onpointerover")
	return onpointerover
}

func (*HTMLHRElement) SetOnpointerover(onpointerover func(event.Event)) {
	macro.Rewrite("$_.onpointerover = $1", onpointerover)
}

func (*HTMLHRElement) Onpointerup() (onpointerup func(event.Event)) {
	macro.Rewrite("$_.onpointerup")
	return onpointerup
}

func (*HTMLHRElement) SetOnpointerup(onpointerup func(event.Event)) {
	macro.Rewrite("$_.onpointerup = $1", onpointerup)
}

func (*HTMLHRElement) Onwheel() (onwheel func(event.Event)) {
	macro.Rewrite("$_.onwheel")
	return onwheel
}

func (*HTMLHRElement) SetOnwheel(onwheel func(event.Event)) {
	macro.Rewrite("$_.onwheel = $1", onwheel)
}

func (*HTMLHRElement) ChildElementCount() (childElementCount uint) {
	macro.Rewrite("$_.childElementCount")
	return childElementCount
}

func (*HTMLHRElement) FirstElementChild() (firstElementChild element.Element) {
	macro.Rewrite("$_.firstElementChild")
	return firstElementChild
}

func (*HTMLHRElement) LastElementChild() (lastElementChild element.Element) {
	macro.Rewrite("$_.lastElementChild")
	return lastElementChild
}

func (*HTMLHRElement) NextElementSibling() (nextElementSibling element.Element) {
	macro.Rewrite("$_.nextElementSibling")
	return nextElementSibling
}

func (*HTMLHRElement) PreviousElementSibling() (previousElementSibling element.Element) {
	macro.Rewrite("$_.previousElementSibling")
	return previousElementSibling
}

func (*HTMLHRElement) Attributes() (attributes *dom.NamedNodeMap) {
	macro.Rewrite("$_.attributes")
	return attributes
}

func (*HTMLHRElement) BaseURI() (baseURI string) {
	macro.Rewrite("$_.baseURI")
	return baseURI
}

func (*HTMLHRElement) ChildNodes() (childNodes *dom.NodeList) {
	macro.Rewrite("$_.childNodes")
	return childNodes
}

func (*HTMLHRElement) FirstChild() (firstChild dom.Node) {
	macro.Rewrite("$_.firstChild")
	return firstChild
}

func (*HTMLHRElement) LastChild() (lastChild dom.Node) {
	macro.Rewrite("$_.lastChild")
	return lastChild
}

func (*HTMLHRElement) LocalName() (localName string) {
	macro.Rewrite("$_.localName")
	return localName
}

func (*HTMLHRElement) NamespaceURI() (namespaceURI string) {
	macro.Rewrite("$_.namespaceURI")
	return namespaceURI
}

func (*HTMLHRElement) NextSibling() (nextSibling dom.Node) {
	macro.Rewrite("$_.nextSibling")
	return nextSibling
}

func (*HTMLHRElement) NodeName() (nodeName string) {
	macro.Rewrite("$_.nodeName")
	return nodeName
}

func (*HTMLHRElement) NodeType() (nodeType uint8) {
	macro.Rewrite("$_.nodeType")
	return nodeType
}

func (*HTMLHRElement) NodeValue() (nodeValue string) {
	macro.Rewrite("$_.nodeValue")
	return nodeValue
}

func (*HTMLHRElement) SetNodeValue(nodeValue string) {
	macro.Rewrite("$_.nodeValue = $1", nodeValue)
}

func (*HTMLHRElement) OwnerDocument() (ownerDocument document.Document) {
	macro.Rewrite("$_.ownerDocument")
	return ownerDocument
}

func (*HTMLHRElement) ParentElement() (parentElement HTMLElement) {
	macro.Rewrite("$_.parentElement")
	return parentElement
}

func (*HTMLHRElement) ParentNode() (parentNode dom.Node) {
	macro.Rewrite("$_.parentNode")
	return parentNode
}

func (*HTMLHRElement) PreviousSibling() (previousSibling dom.Node) {
	macro.Rewrite("$_.previousSibling")
	return previousSibling
}

func (*HTMLHRElement) TextContent() (textContent string) {
	macro.Rewrite("$_.textContent")
	return textContent
}

func (*HTMLHRElement) SetTextContent(textContent string) {
	macro.Rewrite("$_.textContent = $1", textContent)
}
