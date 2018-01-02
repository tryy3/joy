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

var _ HTMLElement = (*HTMLMarqueeElement)(nil)
var _ element.Element = (*HTMLMarqueeElement)(nil)
var _ utils.GlobalEventHandlers = (*HTMLMarqueeElement)(nil)
var _ element.ElementTraversal = (*HTMLMarqueeElement)(nil)
var _ dom.NodeSelector = (*HTMLMarqueeElement)(nil)
var _ dom.ChildNode = (*HTMLMarqueeElement)(nil)
var _ dom.Node = (*HTMLMarqueeElement)(nil)
var _ event.EventTarget = (*HTMLMarqueeElement)(nil)

type HTMLMarqueeElement struct {
}

func (*HTMLMarqueeElement) Start() {
	macro.Rewrite("$_.start()")
}

func (*HTMLMarqueeElement) Stop() {
	macro.Rewrite("$_.stop()")
}

func (*HTMLMarqueeElement) Blur() {
	macro.Rewrite("$_.blur()")
}

func (*HTMLMarqueeElement) Click() {
	macro.Rewrite("$_.click()")
}

func (*HTMLMarqueeElement) DragDrop() (b bool) {
	macro.Rewrite("$_.dragDrop()")
	return b
}

func (*HTMLMarqueeElement) Focus() {
	macro.Rewrite("$_.focus()")
}

func (*HTMLMarqueeElement) GetElementsByClassName(classNames string) (w *dom.NodeList) {
	macro.Rewrite("$_.getElementsByClassName($1)", classNames)
	return w
}

func (*HTMLMarqueeElement) InsertAdjacentElement(position string, insertedElement element.Element) (w element.Element) {
	macro.Rewrite("$_.insertAdjacentElement($1, $2)", position, insertedElement)
	return w
}

func (*HTMLMarqueeElement) InsertAdjacentHTML(where string, html string) {
	macro.Rewrite("$_.insertAdjacentHTML($1, $2)", where, html)
}

func (*HTMLMarqueeElement) InsertAdjacentText(where string, text string) {
	macro.Rewrite("$_.insertAdjacentText($1, $2)", where, text)
}

func (*HTMLMarqueeElement) MsGetInputContext() (w *ms.MSInputMethodContext) {
	macro.Rewrite("$_.msGetInputContext()")
	return w
}

func (*HTMLMarqueeElement) ScrollIntoView(top *bool) {
	macro.Rewrite("$_.scrollIntoView($1)", top)
}

func (*HTMLMarqueeElement) GetAttribute(qualifiedName string) (s string) {
	macro.Rewrite("$_.getAttribute($1)", qualifiedName)
	return s
}

func (*HTMLMarqueeElement) GetAttributeNode(name string) (w *dom.Attr) {
	macro.Rewrite("$_.getAttributeNode($1)", name)
	return w
}

func (*HTMLMarqueeElement) GetAttributeNodeNS(namespaceURI string, localName string) (w *dom.Attr) {
	macro.Rewrite("$_.getAttributeNodeNS($1, $2)", namespaceURI, localName)
	return w
}

func (*HTMLMarqueeElement) GetAttributeNS(namespaceURI string, localName string) (s string) {
	macro.Rewrite("$_.getAttributeNS($1, $2)", namespaceURI, localName)
	return s
}

func (*HTMLMarqueeElement) GetBoundingClientRect() (c *dom.ClientRect) {
	macro.Rewrite("$_.getBoundingClientRect()")
	return c
}

func (*HTMLMarqueeElement) GetClientRects() (c *dom.ClientRectList) {
	macro.Rewrite("$_.getClientRects()")
	return c
}

func (*HTMLMarqueeElement) GetElementsByTagName(name string) (w *dom.NodeList) {
	macro.Rewrite("$_.getElementsByTagName($1)", name)
	return w
}

func (*HTMLMarqueeElement) GetElementsByTagNameNS(namespaceURI string, localName string) (w *dom.NodeList) {
	macro.Rewrite("$_.getElementsByTagNameNS($1, $2)", namespaceURI, localName)
	return w
}

func (*HTMLMarqueeElement) HasAttribute(name string) (b bool) {
	macro.Rewrite("$_.hasAttribute($1)", name)
	return b
}

func (*HTMLMarqueeElement) HasAttributeNS(namespaceURI string, localName string) (b bool) {
	macro.Rewrite("$_.hasAttributeNS($1, $2)", namespaceURI, localName)
	return b
}

func (*HTMLMarqueeElement) MsGetRegionContent() (w *ms.MSRangeCollection) {
	macro.Rewrite("$_.msGetRegionContent()")
	return w
}

func (*HTMLMarqueeElement) MsGetUntransformedBounds() (c *dom.ClientRect) {
	macro.Rewrite("$_.msGetUntransformedBounds()")
	return c
}

func (*HTMLMarqueeElement) MsMatchesSelector(selectors string) (b bool) {
	macro.Rewrite("$_.msMatchesSelector($1)", selectors)
	return b
}

func (*HTMLMarqueeElement) MsReleasePointerCapture(pointerId int) {
	macro.Rewrite("$_.msReleasePointerCapture($1)", pointerId)
}

func (*HTMLMarqueeElement) MsSetPointerCapture(pointerId int) {
	macro.Rewrite("$_.msSetPointerCapture($1)", pointerId)
}

func (*HTMLMarqueeElement) MsZoomTo(args *ms.MsZoomToOptions) {
	macro.Rewrite("$_.msZoomTo($1)", args)
}

func (*HTMLMarqueeElement) ReleasePointerCapture(pointerId int) {
	macro.Rewrite("$_.releasePointerCapture($1)", pointerId)
}

func (*HTMLMarqueeElement) RemoveAttribute(qualifiedName string) {
	macro.Rewrite("$_.removeAttribute($1)", qualifiedName)
}

func (*HTMLMarqueeElement) RemoveAttributeNode(oldAttr *dom.Attr) (w *dom.Attr) {
	macro.Rewrite("$_.removeAttributeNode($1)", oldAttr)
	return w
}

func (*HTMLMarqueeElement) RemoveAttributeNS(namespaceURI string, localName string) {
	macro.Rewrite("$_.removeAttributeNS($1, $2)", namespaceURI, localName)
}

func (*HTMLMarqueeElement) RequestFullscreen() {
	macro.Rewrite("$_.requestFullscreen()")
}

func (*HTMLMarqueeElement) RequestPointerLock() {
	macro.Rewrite("$_.requestPointerLock()")
}

func (*HTMLMarqueeElement) SetAttribute(qualifiedName string, value string) {
	macro.Rewrite("$_.setAttribute($1, $2)", qualifiedName, value)
}

func (*HTMLMarqueeElement) SetAttributeNode(newAttr *dom.Attr) (w *dom.Attr) {
	macro.Rewrite("$_.setAttributeNode($1)", newAttr)
	return w
}

func (*HTMLMarqueeElement) SetAttributeNodeNS(newAttr *dom.Attr) (w *dom.Attr) {
	macro.Rewrite("$_.setAttributeNodeNS($1)", newAttr)
	return w
}

func (*HTMLMarqueeElement) SetAttributeNS(namespaceURI string, qualifiedName string, value string) {
	macro.Rewrite("$_.setAttributeNS($1, $2, $3)", namespaceURI, qualifiedName, value)
}

func (*HTMLMarqueeElement) SetPointerCapture(pointerId int) {
	macro.Rewrite("$_.setPointerCapture($1)", pointerId)
}

func (*HTMLMarqueeElement) WebkitMatchesSelector(selectors string) (b bool) {
	macro.Rewrite("$_.webkitMatchesSelector($1)", selectors)
	return b
}

func (*HTMLMarqueeElement) WebkitRequestFullscreen() {
	macro.Rewrite("$_.webkitRequestFullscreen()")
}

func (*HTMLMarqueeElement) WebkitRequestFullScreen() {
	macro.Rewrite("$_.webkitRequestFullScreen()")
}

func (*HTMLMarqueeElement) QuerySelector(selectors string) (w element.Element) {
	macro.Rewrite("$_.querySelector($1)", selectors)
	return w
}

func (*HTMLMarqueeElement) QuerySelectorAll(selectors string) (w *dom.NodeList) {
	macro.Rewrite("$_.querySelectorAll($1)", selectors)
	return w
}

func (*HTMLMarqueeElement) AppendChild(newChild dom.Node) (w dom.Node) {
	macro.Rewrite("$_.appendChild($1)", newChild)
	return w
}

func (*HTMLMarqueeElement) CloneNode(deep *bool) (w dom.Node) {
	macro.Rewrite("$_.cloneNode($1)", deep)
	return w
}

func (*HTMLMarqueeElement) CompareDocumentPosition(other dom.Node) (u uint8) {
	macro.Rewrite("$_.compareDocumentPosition($1)", other)
	return u
}

func (*HTMLMarqueeElement) Contains(child dom.Node) (b bool) {
	macro.Rewrite("$_.contains($1)", child)
	return b
}

func (*HTMLMarqueeElement) HasAttributes() (b bool) {
	macro.Rewrite("$_.hasAttributes()")
	return b
}

func (*HTMLMarqueeElement) HasChildNodes() (b bool) {
	macro.Rewrite("$_.hasChildNodes()")
	return b
}

func (*HTMLMarqueeElement) InsertBefore(newChild dom.Node, refChild *dom.Node) (w dom.Node) {
	macro.Rewrite("$_.insertBefore($1, $2)", newChild, refChild)
	return w
}

func (*HTMLMarqueeElement) IsDefaultNamespace(namespaceURI string) (b bool) {
	macro.Rewrite("$_.isDefaultNamespace($1)", namespaceURI)
	return b
}

func (*HTMLMarqueeElement) IsEqualNode(arg dom.Node) (b bool) {
	macro.Rewrite("$_.isEqualNode($1)", arg)
	return b
}

func (*HTMLMarqueeElement) IsSameNode(other dom.Node) (b bool) {
	macro.Rewrite("$_.isSameNode($1)", other)
	return b
}

func (*HTMLMarqueeElement) LookupNamespaceURI(prefix string) (s string) {
	macro.Rewrite("$_.lookupNamespaceURI($1)", prefix)
	return s
}

func (*HTMLMarqueeElement) LookupPrefix(namespaceURI string) (s string) {
	macro.Rewrite("$_.lookupPrefix($1)", namespaceURI)
	return s
}

func (*HTMLMarqueeElement) Normalize() {
	macro.Rewrite("$_.normalize()")
}

func (*HTMLMarqueeElement) RemoveChild(oldChild dom.Node) (w dom.Node) {
	macro.Rewrite("$_.removeChild($1)", oldChild)
	return w
}

func (*HTMLMarqueeElement) ReplaceChild(newChild dom.Node, oldChild dom.Node) (w dom.Node) {
	macro.Rewrite("$_.replaceChild($1, $2)", newChild, oldChild)
	return w
}

func (*HTMLMarqueeElement) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*HTMLMarqueeElement) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*HTMLMarqueeElement) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*HTMLMarqueeElement) Behavior() (behavior string) {
	macro.Rewrite("$_.behavior")
	return behavior
}

func (*HTMLMarqueeElement) SetBehavior(behavior string) {
	macro.Rewrite("$_.behavior = $1", behavior)
}

func (*HTMLMarqueeElement) BgColor() (bgColor interface{}) {
	macro.Rewrite("$_.bgColor")
	return bgColor
}

func (*HTMLMarqueeElement) SetBgColor(bgColor interface{}) {
	macro.Rewrite("$_.bgColor = $1", bgColor)
}

func (*HTMLMarqueeElement) Direction() (direction string) {
	macro.Rewrite("$_.direction")
	return direction
}

func (*HTMLMarqueeElement) SetDirection(direction string) {
	macro.Rewrite("$_.direction = $1", direction)
}

func (*HTMLMarqueeElement) Height() (height string) {
	macro.Rewrite("$_.height")
	return height
}

func (*HTMLMarqueeElement) SetHeight(height string) {
	macro.Rewrite("$_.height = $1", height)
}

func (*HTMLMarqueeElement) Hspace() (hspace int) {
	macro.Rewrite("$_.hspace")
	return hspace
}

func (*HTMLMarqueeElement) SetHspace(hspace int) {
	macro.Rewrite("$_.hspace = $1", hspace)
}

func (*HTMLMarqueeElement) Loop() (loop int) {
	macro.Rewrite("$_.loop")
	return loop
}

func (*HTMLMarqueeElement) SetLoop(loop int) {
	macro.Rewrite("$_.loop = $1", loop)
}

func (*HTMLMarqueeElement) Onbounce() (onbounce func(event.Event)) {
	macro.Rewrite("$_.onbounce")
	return onbounce
}

func (*HTMLMarqueeElement) SetOnbounce(onbounce func(event.Event)) {
	macro.Rewrite("$_.onbounce = $1", onbounce)
}

func (*HTMLMarqueeElement) Onfinish() (onfinish func(event.Event)) {
	macro.Rewrite("$_.onfinish")
	return onfinish
}

func (*HTMLMarqueeElement) SetOnfinish(onfinish func(event.Event)) {
	macro.Rewrite("$_.onfinish = $1", onfinish)
}

func (*HTMLMarqueeElement) Onstart() (onstart func(event.Event)) {
	macro.Rewrite("$_.onstart")
	return onstart
}

func (*HTMLMarqueeElement) SetOnstart(onstart func(event.Event)) {
	macro.Rewrite("$_.onstart = $1", onstart)
}

func (*HTMLMarqueeElement) ScrollAmount() (scrollAmount uint) {
	macro.Rewrite("$_.scrollAmount")
	return scrollAmount
}

func (*HTMLMarqueeElement) SetScrollAmount(scrollAmount uint) {
	macro.Rewrite("$_.scrollAmount = $1", scrollAmount)
}

func (*HTMLMarqueeElement) ScrollDelay() (scrollDelay uint) {
	macro.Rewrite("$_.scrollDelay")
	return scrollDelay
}

func (*HTMLMarqueeElement) SetScrollDelay(scrollDelay uint) {
	macro.Rewrite("$_.scrollDelay = $1", scrollDelay)
}

func (*HTMLMarqueeElement) TrueSpeed() (trueSpeed bool) {
	macro.Rewrite("$_.trueSpeed")
	return trueSpeed
}

func (*HTMLMarqueeElement) SetTrueSpeed(trueSpeed bool) {
	macro.Rewrite("$_.trueSpeed = $1", trueSpeed)
}

func (*HTMLMarqueeElement) Vspace() (vspace int) {
	macro.Rewrite("$_.vspace")
	return vspace
}

func (*HTMLMarqueeElement) SetVspace(vspace int) {
	macro.Rewrite("$_.vspace = $1", vspace)
}

func (*HTMLMarqueeElement) Width() (width string) {
	macro.Rewrite("$_.width")
	return width
}

func (*HTMLMarqueeElement) SetWidth(width string) {
	macro.Rewrite("$_.width = $1", width)
}

func (*HTMLMarqueeElement) AccessKey() (accessKey string) {
	macro.Rewrite("$_.accessKey")
	return accessKey
}

func (*HTMLMarqueeElement) SetAccessKey(accessKey string) {
	macro.Rewrite("$_.accessKey = $1", accessKey)
}

func (*HTMLMarqueeElement) Children() (children HTMLCollection) {
	macro.Rewrite("$_.children")
	return children
}

func (*HTMLMarqueeElement) ContentEditable() (contentEditable string) {
	macro.Rewrite("$_.contentEditable")
	return contentEditable
}

func (*HTMLMarqueeElement) SetContentEditable(contentEditable string) {
	macro.Rewrite("$_.contentEditable = $1", contentEditable)
}

func (*HTMLMarqueeElement) Dataset() (dataset *dom.DOMStringMap) {
	macro.Rewrite("$_.dataset")
	return dataset
}

func (*HTMLMarqueeElement) Dir() (dir string) {
	macro.Rewrite("$_.dir")
	return dir
}

func (*HTMLMarqueeElement) SetDir(dir string) {
	macro.Rewrite("$_.dir = $1", dir)
}

func (*HTMLMarqueeElement) Draggable() (draggable bool) {
	macro.Rewrite("$_.draggable")
	return draggable
}

func (*HTMLMarqueeElement) SetDraggable(draggable bool) {
	macro.Rewrite("$_.draggable = $1", draggable)
}

func (*HTMLMarqueeElement) Hidden() (hidden bool) {
	macro.Rewrite("$_.hidden")
	return hidden
}

func (*HTMLMarqueeElement) SetHidden(hidden bool) {
	macro.Rewrite("$_.hidden = $1", hidden)
}

func (*HTMLMarqueeElement) HideFocus() (hideFocus bool) {
	macro.Rewrite("$_.hideFocus")
	return hideFocus
}

func (*HTMLMarqueeElement) SetHideFocus(hideFocus bool) {
	macro.Rewrite("$_.hideFocus = $1", hideFocus)
}

func (*HTMLMarqueeElement) InnerText() (innerText string) {
	macro.Rewrite("$_.innerText")
	return innerText
}

func (*HTMLMarqueeElement) SetInnerText(innerText string) {
	macro.Rewrite("$_.innerText = $1", innerText)
}

func (*HTMLMarqueeElement) IsContentEditable() (isContentEditable bool) {
	macro.Rewrite("$_.isContentEditable")
	return isContentEditable
}

func (*HTMLMarqueeElement) Lang() (lang string) {
	macro.Rewrite("$_.lang")
	return lang
}

func (*HTMLMarqueeElement) SetLang(lang string) {
	macro.Rewrite("$_.lang = $1", lang)
}

func (*HTMLMarqueeElement) OffsetHeight() (offsetHeight int) {
	macro.Rewrite("$_.offsetHeight")
	return offsetHeight
}

func (*HTMLMarqueeElement) OffsetLeft() (offsetLeft int) {
	macro.Rewrite("$_.offsetLeft")
	return offsetLeft
}

func (*HTMLMarqueeElement) OffsetParent() (offsetParent element.Element) {
	macro.Rewrite("$_.offsetParent")
	return offsetParent
}

func (*HTMLMarqueeElement) OffsetTop() (offsetTop int) {
	macro.Rewrite("$_.offsetTop")
	return offsetTop
}

func (*HTMLMarqueeElement) OffsetWidth() (offsetWidth int) {
	macro.Rewrite("$_.offsetWidth")
	return offsetWidth
}

func (*HTMLMarqueeElement) Onabort() (onabort func(event.Event)) {
	macro.Rewrite("$_.onabort")
	return onabort
}

func (*HTMLMarqueeElement) SetOnabort(onabort func(event.Event)) {
	macro.Rewrite("$_.onabort = $1", onabort)
}

func (*HTMLMarqueeElement) Onactivate() (onactivate func(ui.UIEvent)) {
	macro.Rewrite("$_.onactivate")
	return onactivate
}

func (*HTMLMarqueeElement) SetOnactivate(onactivate func(ui.UIEvent)) {
	macro.Rewrite("$_.onactivate = $1", onactivate)
}

func (*HTMLMarqueeElement) Onbeforeactivate() (onbeforeactivate func(ui.UIEvent)) {
	macro.Rewrite("$_.onbeforeactivate")
	return onbeforeactivate
}

func (*HTMLMarqueeElement) SetOnbeforeactivate(onbeforeactivate func(ui.UIEvent)) {
	macro.Rewrite("$_.onbeforeactivate = $1", onbeforeactivate)
}

func (*HTMLMarqueeElement) Onbeforecopy() (onbeforecopy func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.onbeforecopy")
	return onbeforecopy
}

func (*HTMLMarqueeElement) SetOnbeforecopy(onbeforecopy func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.onbeforecopy = $1", onbeforecopy)
}

func (*HTMLMarqueeElement) Onbeforecut() (onbeforecut func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.onbeforecut")
	return onbeforecut
}

func (*HTMLMarqueeElement) SetOnbeforecut(onbeforecut func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.onbeforecut = $1", onbeforecut)
}

func (*HTMLMarqueeElement) Onbeforedeactivate() (onbeforedeactivate func(ui.UIEvent)) {
	macro.Rewrite("$_.onbeforedeactivate")
	return onbeforedeactivate
}

func (*HTMLMarqueeElement) SetOnbeforedeactivate(onbeforedeactivate func(ui.UIEvent)) {
	macro.Rewrite("$_.onbeforedeactivate = $1", onbeforedeactivate)
}

func (*HTMLMarqueeElement) Onbeforepaste() (onbeforepaste func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.onbeforepaste")
	return onbeforepaste
}

func (*HTMLMarqueeElement) SetOnbeforepaste(onbeforepaste func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.onbeforepaste = $1", onbeforepaste)
}

func (*HTMLMarqueeElement) Onblur() (onblur func(*utils.FocusEvent)) {
	macro.Rewrite("$_.onblur")
	return onblur
}

func (*HTMLMarqueeElement) SetOnblur(onblur func(*utils.FocusEvent)) {
	macro.Rewrite("$_.onblur = $1", onblur)
}

func (*HTMLMarqueeElement) Oncanplay() (oncanplay func(event.Event)) {
	macro.Rewrite("$_.oncanplay")
	return oncanplay
}

func (*HTMLMarqueeElement) SetOncanplay(oncanplay func(event.Event)) {
	macro.Rewrite("$_.oncanplay = $1", oncanplay)
}

func (*HTMLMarqueeElement) Oncanplaythrough() (oncanplaythrough func(event.Event)) {
	macro.Rewrite("$_.oncanplaythrough")
	return oncanplaythrough
}

func (*HTMLMarqueeElement) SetOncanplaythrough(oncanplaythrough func(event.Event)) {
	macro.Rewrite("$_.oncanplaythrough = $1", oncanplaythrough)
}

func (*HTMLMarqueeElement) Onchange() (onchange func(event.Event)) {
	macro.Rewrite("$_.onchange")
	return onchange
}

func (*HTMLMarqueeElement) SetOnchange(onchange func(event.Event)) {
	macro.Rewrite("$_.onchange = $1", onchange)
}

func (*HTMLMarqueeElement) Onclick() (onclick func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onclick")
	return onclick
}

func (*HTMLMarqueeElement) SetOnclick(onclick func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onclick = $1", onclick)
}

func (*HTMLMarqueeElement) Oncontextmenu() (oncontextmenu func(*utils.PointerEvent)) {
	macro.Rewrite("$_.oncontextmenu")
	return oncontextmenu
}

func (*HTMLMarqueeElement) SetOncontextmenu(oncontextmenu func(*utils.PointerEvent)) {
	macro.Rewrite("$_.oncontextmenu = $1", oncontextmenu)
}

func (*HTMLMarqueeElement) Oncopy() (oncopy func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.oncopy")
	return oncopy
}

func (*HTMLMarqueeElement) SetOncopy(oncopy func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.oncopy = $1", oncopy)
}

func (*HTMLMarqueeElement) Oncuechange() (oncuechange func(event.Event)) {
	macro.Rewrite("$_.oncuechange")
	return oncuechange
}

func (*HTMLMarqueeElement) SetOncuechange(oncuechange func(event.Event)) {
	macro.Rewrite("$_.oncuechange = $1", oncuechange)
}

func (*HTMLMarqueeElement) Oncut() (oncut func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.oncut")
	return oncut
}

func (*HTMLMarqueeElement) SetOncut(oncut func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.oncut = $1", oncut)
}

func (*HTMLMarqueeElement) Ondblclick() (ondblclick func(mouse.MouseEvent)) {
	macro.Rewrite("$_.ondblclick")
	return ondblclick
}

func (*HTMLMarqueeElement) SetOndblclick(ondblclick func(mouse.MouseEvent)) {
	macro.Rewrite("$_.ondblclick = $1", ondblclick)
}

func (*HTMLMarqueeElement) Ondeactivate() (ondeactivate func(ui.UIEvent)) {
	macro.Rewrite("$_.ondeactivate")
	return ondeactivate
}

func (*HTMLMarqueeElement) SetOndeactivate(ondeactivate func(ui.UIEvent)) {
	macro.Rewrite("$_.ondeactivate = $1", ondeactivate)
}

func (*HTMLMarqueeElement) Ondrag() (ondrag func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondrag")
	return ondrag
}

func (*HTMLMarqueeElement) SetOndrag(ondrag func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondrag = $1", ondrag)
}

func (*HTMLMarqueeElement) Ondragend() (ondragend func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragend")
	return ondragend
}

func (*HTMLMarqueeElement) SetOndragend(ondragend func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragend = $1", ondragend)
}

func (*HTMLMarqueeElement) Ondragenter() (ondragenter func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragenter")
	return ondragenter
}

func (*HTMLMarqueeElement) SetOndragenter(ondragenter func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragenter = $1", ondragenter)
}

func (*HTMLMarqueeElement) Ondragleave() (ondragleave func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragleave")
	return ondragleave
}

func (*HTMLMarqueeElement) SetOndragleave(ondragleave func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragleave = $1", ondragleave)
}

func (*HTMLMarqueeElement) Ondragover() (ondragover func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragover")
	return ondragover
}

func (*HTMLMarqueeElement) SetOndragover(ondragover func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragover = $1", ondragover)
}

func (*HTMLMarqueeElement) Ondragstart() (ondragstart func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragstart")
	return ondragstart
}

func (*HTMLMarqueeElement) SetOndragstart(ondragstart func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragstart = $1", ondragstart)
}

func (*HTMLMarqueeElement) Ondrop() (ondrop func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondrop")
	return ondrop
}

func (*HTMLMarqueeElement) SetOndrop(ondrop func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondrop = $1", ondrop)
}

func (*HTMLMarqueeElement) Ondurationchange() (ondurationchange func(event.Event)) {
	macro.Rewrite("$_.ondurationchange")
	return ondurationchange
}

func (*HTMLMarqueeElement) SetOndurationchange(ondurationchange func(event.Event)) {
	macro.Rewrite("$_.ondurationchange = $1", ondurationchange)
}

func (*HTMLMarqueeElement) Onemptied() (onemptied func(event.Event)) {
	macro.Rewrite("$_.onemptied")
	return onemptied
}

func (*HTMLMarqueeElement) SetOnemptied(onemptied func(event.Event)) {
	macro.Rewrite("$_.onemptied = $1", onemptied)
}

func (*HTMLMarqueeElement) Onended() (onended func(event.Event)) {
	macro.Rewrite("$_.onended")
	return onended
}

func (*HTMLMarqueeElement) SetOnended(onended func(event.Event)) {
	macro.Rewrite("$_.onended = $1", onended)
}

func (*HTMLMarqueeElement) Onerror() (onerror func(event.Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

func (*HTMLMarqueeElement) SetOnerror(onerror func(event.Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

func (*HTMLMarqueeElement) Onfocus() (onfocus func(*utils.FocusEvent)) {
	macro.Rewrite("$_.onfocus")
	return onfocus
}

func (*HTMLMarqueeElement) SetOnfocus(onfocus func(*utils.FocusEvent)) {
	macro.Rewrite("$_.onfocus = $1", onfocus)
}

func (*HTMLMarqueeElement) Oninput() (oninput func(event.Event)) {
	macro.Rewrite("$_.oninput")
	return oninput
}

func (*HTMLMarqueeElement) SetOninput(oninput func(event.Event)) {
	macro.Rewrite("$_.oninput = $1", oninput)
}

func (*HTMLMarqueeElement) Oninvalid() (oninvalid func(event.Event)) {
	macro.Rewrite("$_.oninvalid")
	return oninvalid
}

func (*HTMLMarqueeElement) SetOninvalid(oninvalid func(event.Event)) {
	macro.Rewrite("$_.oninvalid = $1", oninvalid)
}

func (*HTMLMarqueeElement) Onkeydown() (onkeydown func(*utils.KeyboardEvent)) {
	macro.Rewrite("$_.onkeydown")
	return onkeydown
}

func (*HTMLMarqueeElement) SetOnkeydown(onkeydown func(*utils.KeyboardEvent)) {
	macro.Rewrite("$_.onkeydown = $1", onkeydown)
}

func (*HTMLMarqueeElement) Onkeypress() (onkeypress func(*utils.KeyboardEvent)) {
	macro.Rewrite("$_.onkeypress")
	return onkeypress
}

func (*HTMLMarqueeElement) SetOnkeypress(onkeypress func(*utils.KeyboardEvent)) {
	macro.Rewrite("$_.onkeypress = $1", onkeypress)
}

func (*HTMLMarqueeElement) Onkeyup() (onkeyup func(*utils.KeyboardEvent)) {
	macro.Rewrite("$_.onkeyup")
	return onkeyup
}

func (*HTMLMarqueeElement) SetOnkeyup(onkeyup func(*utils.KeyboardEvent)) {
	macro.Rewrite("$_.onkeyup = $1", onkeyup)
}

func (*HTMLMarqueeElement) Onload() (onload func(event.Event)) {
	macro.Rewrite("$_.onload")
	return onload
}

func (*HTMLMarqueeElement) SetOnload(onload func(event.Event)) {
	macro.Rewrite("$_.onload = $1", onload)
}

func (*HTMLMarqueeElement) Onloadeddata() (onloadeddata func(event.Event)) {
	macro.Rewrite("$_.onloadeddata")
	return onloadeddata
}

func (*HTMLMarqueeElement) SetOnloadeddata(onloadeddata func(event.Event)) {
	macro.Rewrite("$_.onloadeddata = $1", onloadeddata)
}

func (*HTMLMarqueeElement) Onloadedmetadata() (onloadedmetadata func(event.Event)) {
	macro.Rewrite("$_.onloadedmetadata")
	return onloadedmetadata
}

func (*HTMLMarqueeElement) SetOnloadedmetadata(onloadedmetadata func(event.Event)) {
	macro.Rewrite("$_.onloadedmetadata = $1", onloadedmetadata)
}

func (*HTMLMarqueeElement) Onloadstart() (onloadstart func(event.Event)) {
	macro.Rewrite("$_.onloadstart")
	return onloadstart
}

func (*HTMLMarqueeElement) SetOnloadstart(onloadstart func(event.Event)) {
	macro.Rewrite("$_.onloadstart = $1", onloadstart)
}

func (*HTMLMarqueeElement) Onmousedown() (onmousedown func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmousedown")
	return onmousedown
}

func (*HTMLMarqueeElement) SetOnmousedown(onmousedown func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmousedown = $1", onmousedown)
}

func (*HTMLMarqueeElement) Onmouseenter() (onmouseenter func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseenter")
	return onmouseenter
}

func (*HTMLMarqueeElement) SetOnmouseenter(onmouseenter func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseenter = $1", onmouseenter)
}

func (*HTMLMarqueeElement) Onmouseleave() (onmouseleave func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseleave")
	return onmouseleave
}

func (*HTMLMarqueeElement) SetOnmouseleave(onmouseleave func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseleave = $1", onmouseleave)
}

func (*HTMLMarqueeElement) Onmousemove() (onmousemove func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmousemove")
	return onmousemove
}

func (*HTMLMarqueeElement) SetOnmousemove(onmousemove func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmousemove = $1", onmousemove)
}

func (*HTMLMarqueeElement) Onmouseout() (onmouseout func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseout")
	return onmouseout
}

func (*HTMLMarqueeElement) SetOnmouseout(onmouseout func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseout = $1", onmouseout)
}

func (*HTMLMarqueeElement) Onmouseover() (onmouseover func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseover")
	return onmouseover
}

func (*HTMLMarqueeElement) SetOnmouseover(onmouseover func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseover = $1", onmouseover)
}

func (*HTMLMarqueeElement) Onmouseup() (onmouseup func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseup")
	return onmouseup
}

func (*HTMLMarqueeElement) SetOnmouseup(onmouseup func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseup = $1", onmouseup)
}

func (*HTMLMarqueeElement) Onmousewheel() (onmousewheel func(*utils.WheelEvent)) {
	macro.Rewrite("$_.onmousewheel")
	return onmousewheel
}

func (*HTMLMarqueeElement) SetOnmousewheel(onmousewheel func(*utils.WheelEvent)) {
	macro.Rewrite("$_.onmousewheel = $1", onmousewheel)
}

func (*HTMLMarqueeElement) Onmscontentzoom() (onmscontentzoom func(ui.UIEvent)) {
	macro.Rewrite("$_.onmscontentzoom")
	return onmscontentzoom
}

func (*HTMLMarqueeElement) SetOnmscontentzoom(onmscontentzoom func(ui.UIEvent)) {
	macro.Rewrite("$_.onmscontentzoom = $1", onmscontentzoom)
}

func (*HTMLMarqueeElement) Onmsmanipulationstatechanged() (onmsmanipulationstatechanged func(*ms.MSManipulationEvent)) {
	macro.Rewrite("$_.onmsmanipulationstatechanged")
	return onmsmanipulationstatechanged
}

func (*HTMLMarqueeElement) SetOnmsmanipulationstatechanged(onmsmanipulationstatechanged func(*ms.MSManipulationEvent)) {
	macro.Rewrite("$_.onmsmanipulationstatechanged = $1", onmsmanipulationstatechanged)
}

func (*HTMLMarqueeElement) Onpaste() (onpaste func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.onpaste")
	return onpaste
}

func (*HTMLMarqueeElement) SetOnpaste(onpaste func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.onpaste = $1", onpaste)
}

func (*HTMLMarqueeElement) Onpause() (onpause func(event.Event)) {
	macro.Rewrite("$_.onpause")
	return onpause
}

func (*HTMLMarqueeElement) SetOnpause(onpause func(event.Event)) {
	macro.Rewrite("$_.onpause = $1", onpause)
}

func (*HTMLMarqueeElement) Onplay() (onplay func(event.Event)) {
	macro.Rewrite("$_.onplay")
	return onplay
}

func (*HTMLMarqueeElement) SetOnplay(onplay func(event.Event)) {
	macro.Rewrite("$_.onplay = $1", onplay)
}

func (*HTMLMarqueeElement) Onplaying() (onplaying func(event.Event)) {
	macro.Rewrite("$_.onplaying")
	return onplaying
}

func (*HTMLMarqueeElement) SetOnplaying(onplaying func(event.Event)) {
	macro.Rewrite("$_.onplaying = $1", onplaying)
}

func (*HTMLMarqueeElement) Onprogress() (onprogress func(event.Event)) {
	macro.Rewrite("$_.onprogress")
	return onprogress
}

func (*HTMLMarqueeElement) SetOnprogress(onprogress func(event.Event)) {
	macro.Rewrite("$_.onprogress = $1", onprogress)
}

func (*HTMLMarqueeElement) Onratechange() (onratechange func(event.Event)) {
	macro.Rewrite("$_.onratechange")
	return onratechange
}

func (*HTMLMarqueeElement) SetOnratechange(onratechange func(event.Event)) {
	macro.Rewrite("$_.onratechange = $1", onratechange)
}

func (*HTMLMarqueeElement) Onreset() (onreset func(event.Event)) {
	macro.Rewrite("$_.onreset")
	return onreset
}

func (*HTMLMarqueeElement) SetOnreset(onreset func(event.Event)) {
	macro.Rewrite("$_.onreset = $1", onreset)
}

func (*HTMLMarqueeElement) Onscroll() (onscroll func(ui.UIEvent)) {
	macro.Rewrite("$_.onscroll")
	return onscroll
}

func (*HTMLMarqueeElement) SetOnscroll(onscroll func(ui.UIEvent)) {
	macro.Rewrite("$_.onscroll = $1", onscroll)
}

func (*HTMLMarqueeElement) Onseeked() (onseeked func(event.Event)) {
	macro.Rewrite("$_.onseeked")
	return onseeked
}

func (*HTMLMarqueeElement) SetOnseeked(onseeked func(event.Event)) {
	macro.Rewrite("$_.onseeked = $1", onseeked)
}

func (*HTMLMarqueeElement) Onseeking() (onseeking func(event.Event)) {
	macro.Rewrite("$_.onseeking")
	return onseeking
}

func (*HTMLMarqueeElement) SetOnseeking(onseeking func(event.Event)) {
	macro.Rewrite("$_.onseeking = $1", onseeking)
}

func (*HTMLMarqueeElement) Onselect() (onselect func(ui.UIEvent)) {
	macro.Rewrite("$_.onselect")
	return onselect
}

func (*HTMLMarqueeElement) SetOnselect(onselect func(ui.UIEvent)) {
	macro.Rewrite("$_.onselect = $1", onselect)
}

func (*HTMLMarqueeElement) Onselectstart() (onselectstart func(event.Event)) {
	macro.Rewrite("$_.onselectstart")
	return onselectstart
}

func (*HTMLMarqueeElement) SetOnselectstart(onselectstart func(event.Event)) {
	macro.Rewrite("$_.onselectstart = $1", onselectstart)
}

func (*HTMLMarqueeElement) Onstalled() (onstalled func(event.Event)) {
	macro.Rewrite("$_.onstalled")
	return onstalled
}

func (*HTMLMarqueeElement) SetOnstalled(onstalled func(event.Event)) {
	macro.Rewrite("$_.onstalled = $1", onstalled)
}

func (*HTMLMarqueeElement) Onsubmit() (onsubmit func(event.Event)) {
	macro.Rewrite("$_.onsubmit")
	return onsubmit
}

func (*HTMLMarqueeElement) SetOnsubmit(onsubmit func(event.Event)) {
	macro.Rewrite("$_.onsubmit = $1", onsubmit)
}

func (*HTMLMarqueeElement) Onsuspend() (onsuspend func(event.Event)) {
	macro.Rewrite("$_.onsuspend")
	return onsuspend
}

func (*HTMLMarqueeElement) SetOnsuspend(onsuspend func(event.Event)) {
	macro.Rewrite("$_.onsuspend = $1", onsuspend)
}

func (*HTMLMarqueeElement) Ontimeupdate() (ontimeupdate func(event.Event)) {
	macro.Rewrite("$_.ontimeupdate")
	return ontimeupdate
}

func (*HTMLMarqueeElement) SetOntimeupdate(ontimeupdate func(event.Event)) {
	macro.Rewrite("$_.ontimeupdate = $1", ontimeupdate)
}

func (*HTMLMarqueeElement) Onvolumechange() (onvolumechange func(event.Event)) {
	macro.Rewrite("$_.onvolumechange")
	return onvolumechange
}

func (*HTMLMarqueeElement) SetOnvolumechange(onvolumechange func(event.Event)) {
	macro.Rewrite("$_.onvolumechange = $1", onvolumechange)
}

func (*HTMLMarqueeElement) Onwaiting() (onwaiting func(event.Event)) {
	macro.Rewrite("$_.onwaiting")
	return onwaiting
}

func (*HTMLMarqueeElement) SetOnwaiting(onwaiting func(event.Event)) {
	macro.Rewrite("$_.onwaiting = $1", onwaiting)
}

func (*HTMLMarqueeElement) OuterText() (outerText string) {
	macro.Rewrite("$_.outerText")
	return outerText
}

func (*HTMLMarqueeElement) SetOuterText(outerText string) {
	macro.Rewrite("$_.outerText = $1", outerText)
}

func (*HTMLMarqueeElement) Spellcheck() (spellcheck bool) {
	macro.Rewrite("$_.spellcheck")
	return spellcheck
}

func (*HTMLMarqueeElement) SetSpellcheck(spellcheck bool) {
	macro.Rewrite("$_.spellcheck = $1", spellcheck)
}

func (*HTMLMarqueeElement) Style() (style *css.CSSStyleDeclaration) {
	macro.Rewrite("$_.style")
	return style
}

func (*HTMLMarqueeElement) TabIndex() (tabIndex int8) {
	macro.Rewrite("$_.tabIndex")
	return tabIndex
}

func (*HTMLMarqueeElement) SetTabIndex(tabIndex int8) {
	macro.Rewrite("$_.tabIndex = $1", tabIndex)
}

func (*HTMLMarqueeElement) Title() (title string) {
	macro.Rewrite("$_.title")
	return title
}

func (*HTMLMarqueeElement) SetTitle(title string) {
	macro.Rewrite("$_.title = $1", title)
}

func (*HTMLMarqueeElement) ClassList() (classList dom.DOMTokenList) {
	macro.Rewrite("$_.classList")
	return classList
}

func (*HTMLMarqueeElement) ClassName() (className string) {
	macro.Rewrite("$_.className")
	return className
}

func (*HTMLMarqueeElement) SetClassName(className string) {
	macro.Rewrite("$_.className = $1", className)
}

func (*HTMLMarqueeElement) ClientHeight() (clientHeight int) {
	macro.Rewrite("$_.clientHeight")
	return clientHeight
}

func (*HTMLMarqueeElement) ClientLeft() (clientLeft int) {
	macro.Rewrite("$_.clientLeft")
	return clientLeft
}

func (*HTMLMarqueeElement) ClientTop() (clientTop int) {
	macro.Rewrite("$_.clientTop")
	return clientTop
}

func (*HTMLMarqueeElement) ClientWidth() (clientWidth int) {
	macro.Rewrite("$_.clientWidth")
	return clientWidth
}

func (*HTMLMarqueeElement) ID() (id string) {
	macro.Rewrite("$_.id")
	return id
}

func (*HTMLMarqueeElement) SetID(id string) {
	macro.Rewrite("$_.id = $1", id)
}

func (*HTMLMarqueeElement) InnerHTML() (innerHTML string) {
	macro.Rewrite("$_.innerHTML")
	return innerHTML
}

func (*HTMLMarqueeElement) SetInnerHTML(innerHTML string) {
	macro.Rewrite("$_.innerHTML = $1", innerHTML)
}

func (*HTMLMarqueeElement) MsContentZoomFactor() (msContentZoomFactor float32) {
	macro.Rewrite("$_.msContentZoomFactor")
	return msContentZoomFactor
}

func (*HTMLMarqueeElement) SetMsContentZoomFactor(msContentZoomFactor float32) {
	macro.Rewrite("$_.msContentZoomFactor = $1", msContentZoomFactor)
}

func (*HTMLMarqueeElement) MsRegionOverflow() (msRegionOverflow string) {
	macro.Rewrite("$_.msRegionOverflow")
	return msRegionOverflow
}

func (*HTMLMarqueeElement) Onariarequest() (onariarequest func(event.Event)) {
	macro.Rewrite("$_.onariarequest")
	return onariarequest
}

func (*HTMLMarqueeElement) SetOnariarequest(onariarequest func(event.Event)) {
	macro.Rewrite("$_.onariarequest = $1", onariarequest)
}

func (*HTMLMarqueeElement) Oncommand() (oncommand func(event.Event)) {
	macro.Rewrite("$_.oncommand")
	return oncommand
}

func (*HTMLMarqueeElement) SetOncommand(oncommand func(event.Event)) {
	macro.Rewrite("$_.oncommand = $1", oncommand)
}

func (*HTMLMarqueeElement) Ongotpointercapture() (ongotpointercapture func(*utils.PointerEvent)) {
	macro.Rewrite("$_.ongotpointercapture")
	return ongotpointercapture
}

func (*HTMLMarqueeElement) SetOngotpointercapture(ongotpointercapture func(*utils.PointerEvent)) {
	macro.Rewrite("$_.ongotpointercapture = $1", ongotpointercapture)
}

func (*HTMLMarqueeElement) Onlostpointercapture() (onlostpointercapture func(*utils.PointerEvent)) {
	macro.Rewrite("$_.onlostpointercapture")
	return onlostpointercapture
}

func (*HTMLMarqueeElement) SetOnlostpointercapture(onlostpointercapture func(*utils.PointerEvent)) {
	macro.Rewrite("$_.onlostpointercapture = $1", onlostpointercapture)
}

func (*HTMLMarqueeElement) Onmsgesturechange() (onmsgesturechange func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturechange")
	return onmsgesturechange
}

func (*HTMLMarqueeElement) SetOnmsgesturechange(onmsgesturechange func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturechange = $1", onmsgesturechange)
}

func (*HTMLMarqueeElement) Onmsgesturedoubletap() (onmsgesturedoubletap func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturedoubletap")
	return onmsgesturedoubletap
}

func (*HTMLMarqueeElement) SetOnmsgesturedoubletap(onmsgesturedoubletap func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturedoubletap = $1", onmsgesturedoubletap)
}

func (*HTMLMarqueeElement) Onmsgestureend() (onmsgestureend func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgestureend")
	return onmsgestureend
}

func (*HTMLMarqueeElement) SetOnmsgestureend(onmsgestureend func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgestureend = $1", onmsgestureend)
}

func (*HTMLMarqueeElement) Onmsgesturehold() (onmsgesturehold func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturehold")
	return onmsgesturehold
}

func (*HTMLMarqueeElement) SetOnmsgesturehold(onmsgesturehold func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturehold = $1", onmsgesturehold)
}

func (*HTMLMarqueeElement) Onmsgesturestart() (onmsgesturestart func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturestart")
	return onmsgesturestart
}

func (*HTMLMarqueeElement) SetOnmsgesturestart(onmsgesturestart func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturestart = $1", onmsgesturestart)
}

func (*HTMLMarqueeElement) Onmsgesturetap() (onmsgesturetap func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturetap")
	return onmsgesturetap
}

func (*HTMLMarqueeElement) SetOnmsgesturetap(onmsgesturetap func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturetap = $1", onmsgesturetap)
}

func (*HTMLMarqueeElement) Onmsgotpointercapture() (onmsgotpointercapture func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmsgotpointercapture")
	return onmsgotpointercapture
}

func (*HTMLMarqueeElement) SetOnmsgotpointercapture(onmsgotpointercapture func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmsgotpointercapture = $1", onmsgotpointercapture)
}

func (*HTMLMarqueeElement) Onmsinertiastart() (onmsinertiastart func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsinertiastart")
	return onmsinertiastart
}

func (*HTMLMarqueeElement) SetOnmsinertiastart(onmsinertiastart func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsinertiastart = $1", onmsinertiastart)
}

func (*HTMLMarqueeElement) Onmslostpointercapture() (onmslostpointercapture func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmslostpointercapture")
	return onmslostpointercapture
}

func (*HTMLMarqueeElement) SetOnmslostpointercapture(onmslostpointercapture func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmslostpointercapture = $1", onmslostpointercapture)
}

func (*HTMLMarqueeElement) Onmspointercancel() (onmspointercancel func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointercancel")
	return onmspointercancel
}

func (*HTMLMarqueeElement) SetOnmspointercancel(onmspointercancel func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointercancel = $1", onmspointercancel)
}

func (*HTMLMarqueeElement) Onmspointerdown() (onmspointerdown func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerdown")
	return onmspointerdown
}

func (*HTMLMarqueeElement) SetOnmspointerdown(onmspointerdown func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerdown = $1", onmspointerdown)
}

func (*HTMLMarqueeElement) Onmspointerenter() (onmspointerenter func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerenter")
	return onmspointerenter
}

func (*HTMLMarqueeElement) SetOnmspointerenter(onmspointerenter func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerenter = $1", onmspointerenter)
}

func (*HTMLMarqueeElement) Onmspointerleave() (onmspointerleave func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerleave")
	return onmspointerleave
}

func (*HTMLMarqueeElement) SetOnmspointerleave(onmspointerleave func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerleave = $1", onmspointerleave)
}

func (*HTMLMarqueeElement) Onmspointermove() (onmspointermove func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointermove")
	return onmspointermove
}

func (*HTMLMarqueeElement) SetOnmspointermove(onmspointermove func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointermove = $1", onmspointermove)
}

func (*HTMLMarqueeElement) Onmspointerout() (onmspointerout func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerout")
	return onmspointerout
}

func (*HTMLMarqueeElement) SetOnmspointerout(onmspointerout func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerout = $1", onmspointerout)
}

func (*HTMLMarqueeElement) Onmspointerover() (onmspointerover func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerover")
	return onmspointerover
}

func (*HTMLMarqueeElement) SetOnmspointerover(onmspointerover func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerover = $1", onmspointerover)
}

func (*HTMLMarqueeElement) Onmspointerup() (onmspointerup func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerup")
	return onmspointerup
}

func (*HTMLMarqueeElement) SetOnmspointerup(onmspointerup func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerup = $1", onmspointerup)
}

func (*HTMLMarqueeElement) Ontouchcancel() (ontouchcancel func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchcancel")
	return ontouchcancel
}

func (*HTMLMarqueeElement) SetOntouchcancel(ontouchcancel func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchcancel = $1", ontouchcancel)
}

func (*HTMLMarqueeElement) Ontouchend() (ontouchend func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchend")
	return ontouchend
}

func (*HTMLMarqueeElement) SetOntouchend(ontouchend func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchend = $1", ontouchend)
}

func (*HTMLMarqueeElement) Ontouchmove() (ontouchmove func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchmove")
	return ontouchmove
}

func (*HTMLMarqueeElement) SetOntouchmove(ontouchmove func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchmove = $1", ontouchmove)
}

func (*HTMLMarqueeElement) Ontouchstart() (ontouchstart func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchstart")
	return ontouchstart
}

func (*HTMLMarqueeElement) SetOntouchstart(ontouchstart func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchstart = $1", ontouchstart)
}

func (*HTMLMarqueeElement) Onwebkitfullscreenchange() (onwebkitfullscreenchange func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenchange")
	return onwebkitfullscreenchange
}

func (*HTMLMarqueeElement) SetOnwebkitfullscreenchange(onwebkitfullscreenchange func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenchange = $1", onwebkitfullscreenchange)
}

func (*HTMLMarqueeElement) Onwebkitfullscreenerror() (onwebkitfullscreenerror func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenerror")
	return onwebkitfullscreenerror
}

func (*HTMLMarqueeElement) SetOnwebkitfullscreenerror(onwebkitfullscreenerror func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenerror = $1", onwebkitfullscreenerror)
}

func (*HTMLMarqueeElement) OuterHTML() (outerHTML string) {
	macro.Rewrite("$_.outerHTML")
	return outerHTML
}

func (*HTMLMarqueeElement) SetOuterHTML(outerHTML string) {
	macro.Rewrite("$_.outerHTML = $1", outerHTML)
}

func (*HTMLMarqueeElement) Prefix() (prefix string) {
	macro.Rewrite("$_.prefix")
	return prefix
}

func (*HTMLMarqueeElement) ScrollHeight() (scrollHeight int) {
	macro.Rewrite("$_.scrollHeight")
	return scrollHeight
}

func (*HTMLMarqueeElement) ScrollLeft() (scrollLeft int) {
	macro.Rewrite("$_.scrollLeft")
	return scrollLeft
}

func (*HTMLMarqueeElement) SetScrollLeft(scrollLeft int) {
	macro.Rewrite("$_.scrollLeft = $1", scrollLeft)
}

func (*HTMLMarqueeElement) ScrollTop() (scrollTop int) {
	macro.Rewrite("$_.scrollTop")
	return scrollTop
}

func (*HTMLMarqueeElement) SetScrollTop(scrollTop int) {
	macro.Rewrite("$_.scrollTop = $1", scrollTop)
}

func (*HTMLMarqueeElement) ScrollWidth() (scrollWidth int) {
	macro.Rewrite("$_.scrollWidth")
	return scrollWidth
}

func (*HTMLMarqueeElement) TagName() (tagName string) {
	macro.Rewrite("$_.tagName")
	return tagName
}

func (*HTMLMarqueeElement) Onpointercancel() (onpointercancel func(event.Event)) {
	macro.Rewrite("$_.onpointercancel")
	return onpointercancel
}

func (*HTMLMarqueeElement) SetOnpointercancel(onpointercancel func(event.Event)) {
	macro.Rewrite("$_.onpointercancel = $1", onpointercancel)
}

func (*HTMLMarqueeElement) Onpointerdown() (onpointerdown func(event.Event)) {
	macro.Rewrite("$_.onpointerdown")
	return onpointerdown
}

func (*HTMLMarqueeElement) SetOnpointerdown(onpointerdown func(event.Event)) {
	macro.Rewrite("$_.onpointerdown = $1", onpointerdown)
}

func (*HTMLMarqueeElement) Onpointerenter() (onpointerenter func(event.Event)) {
	macro.Rewrite("$_.onpointerenter")
	return onpointerenter
}

func (*HTMLMarqueeElement) SetOnpointerenter(onpointerenter func(event.Event)) {
	macro.Rewrite("$_.onpointerenter = $1", onpointerenter)
}

func (*HTMLMarqueeElement) Onpointerleave() (onpointerleave func(event.Event)) {
	macro.Rewrite("$_.onpointerleave")
	return onpointerleave
}

func (*HTMLMarqueeElement) SetOnpointerleave(onpointerleave func(event.Event)) {
	macro.Rewrite("$_.onpointerleave = $1", onpointerleave)
}

func (*HTMLMarqueeElement) Onpointermove() (onpointermove func(event.Event)) {
	macro.Rewrite("$_.onpointermove")
	return onpointermove
}

func (*HTMLMarqueeElement) SetOnpointermove(onpointermove func(event.Event)) {
	macro.Rewrite("$_.onpointermove = $1", onpointermove)
}

func (*HTMLMarqueeElement) Onpointerout() (onpointerout func(event.Event)) {
	macro.Rewrite("$_.onpointerout")
	return onpointerout
}

func (*HTMLMarqueeElement) SetOnpointerout(onpointerout func(event.Event)) {
	macro.Rewrite("$_.onpointerout = $1", onpointerout)
}

func (*HTMLMarqueeElement) Onpointerover() (onpointerover func(event.Event)) {
	macro.Rewrite("$_.onpointerover")
	return onpointerover
}

func (*HTMLMarqueeElement) SetOnpointerover(onpointerover func(event.Event)) {
	macro.Rewrite("$_.onpointerover = $1", onpointerover)
}

func (*HTMLMarqueeElement) Onpointerup() (onpointerup func(event.Event)) {
	macro.Rewrite("$_.onpointerup")
	return onpointerup
}

func (*HTMLMarqueeElement) SetOnpointerup(onpointerup func(event.Event)) {
	macro.Rewrite("$_.onpointerup = $1", onpointerup)
}

func (*HTMLMarqueeElement) Onwheel() (onwheel func(event.Event)) {
	macro.Rewrite("$_.onwheel")
	return onwheel
}

func (*HTMLMarqueeElement) SetOnwheel(onwheel func(event.Event)) {
	macro.Rewrite("$_.onwheel = $1", onwheel)
}

func (*HTMLMarqueeElement) ChildElementCount() (childElementCount uint) {
	macro.Rewrite("$_.childElementCount")
	return childElementCount
}

func (*HTMLMarqueeElement) FirstElementChild() (firstElementChild element.Element) {
	macro.Rewrite("$_.firstElementChild")
	return firstElementChild
}

func (*HTMLMarqueeElement) LastElementChild() (lastElementChild element.Element) {
	macro.Rewrite("$_.lastElementChild")
	return lastElementChild
}

func (*HTMLMarqueeElement) NextElementSibling() (nextElementSibling element.Element) {
	macro.Rewrite("$_.nextElementSibling")
	return nextElementSibling
}

func (*HTMLMarqueeElement) PreviousElementSibling() (previousElementSibling element.Element) {
	macro.Rewrite("$_.previousElementSibling")
	return previousElementSibling
}

func (*HTMLMarqueeElement) Attributes() (attributes *dom.NamedNodeMap) {
	macro.Rewrite("$_.attributes")
	return attributes
}

func (*HTMLMarqueeElement) BaseURI() (baseURI string) {
	macro.Rewrite("$_.baseURI")
	return baseURI
}

func (*HTMLMarqueeElement) ChildNodes() (childNodes *dom.NodeList) {
	macro.Rewrite("$_.childNodes")
	return childNodes
}

func (*HTMLMarqueeElement) FirstChild() (firstChild dom.Node) {
	macro.Rewrite("$_.firstChild")
	return firstChild
}

func (*HTMLMarqueeElement) LastChild() (lastChild dom.Node) {
	macro.Rewrite("$_.lastChild")
	return lastChild
}

func (*HTMLMarqueeElement) LocalName() (localName string) {
	macro.Rewrite("$_.localName")
	return localName
}

func (*HTMLMarqueeElement) NamespaceURI() (namespaceURI string) {
	macro.Rewrite("$_.namespaceURI")
	return namespaceURI
}

func (*HTMLMarqueeElement) NextSibling() (nextSibling dom.Node) {
	macro.Rewrite("$_.nextSibling")
	return nextSibling
}

func (*HTMLMarqueeElement) NodeName() (nodeName string) {
	macro.Rewrite("$_.nodeName")
	return nodeName
}

func (*HTMLMarqueeElement) NodeType() (nodeType uint8) {
	macro.Rewrite("$_.nodeType")
	return nodeType
}

func (*HTMLMarqueeElement) NodeValue() (nodeValue string) {
	macro.Rewrite("$_.nodeValue")
	return nodeValue
}

func (*HTMLMarqueeElement) SetNodeValue(nodeValue string) {
	macro.Rewrite("$_.nodeValue = $1", nodeValue)
}

func (*HTMLMarqueeElement) OwnerDocument() (ownerDocument document.Document) {
	macro.Rewrite("$_.ownerDocument")
	return ownerDocument
}

func (*HTMLMarqueeElement) ParentElement() (parentElement HTMLElement) {
	macro.Rewrite("$_.parentElement")
	return parentElement
}

func (*HTMLMarqueeElement) ParentNode() (parentNode dom.Node) {
	macro.Rewrite("$_.parentNode")
	return parentNode
}

func (*HTMLMarqueeElement) PreviousSibling() (previousSibling dom.Node) {
	macro.Rewrite("$_.previousSibling")
	return previousSibling
}

func (*HTMLMarqueeElement) TextContent() (textContent string) {
	macro.Rewrite("$_.textContent")
	return textContent
}

func (*HTMLMarqueeElement) SetTextContent(textContent string) {
	macro.Rewrite("$_.textContent = $1", textContent)
}
