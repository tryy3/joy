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

var _ HTMLElement = (*HTMLSelectElement)(nil)
var _ element.Element = (*HTMLSelectElement)(nil)
var _ utils.GlobalEventHandlers = (*HTMLSelectElement)(nil)
var _ element.ElementTraversal = (*HTMLSelectElement)(nil)
var _ dom.NodeSelector = (*HTMLSelectElement)(nil)
var _ dom.ChildNode = (*HTMLSelectElement)(nil)
var _ dom.Node = (*HTMLSelectElement)(nil)
var _ event.EventTarget = (*HTMLSelectElement)(nil)

type HTMLSelectElement struct {
}

func (*HTMLSelectElement) Add(element HTMLElement, before *interface{}) {
	macro.Rewrite("$_.add($1, $2)", element, before)
}

func (*HTMLSelectElement) CheckValidity() (b bool) {
	macro.Rewrite("$_.checkValidity()")
	return b
}

func (*HTMLSelectElement) Item(name *interface{}, index *interface{}) (i interface{}) {
	macro.Rewrite("$_.item($1, $2)", name, index)
	return i
}

func (*HTMLSelectElement) NamedItem(name string) (i interface{}) {
	macro.Rewrite("$_.namedItem($1)", name)
	return i
}

func (*HTMLSelectElement) Remove(index *int) {
	macro.Rewrite("$_.remove($1)", index)
}

func (*HTMLSelectElement) SetCustomValidity(err string) {
	macro.Rewrite("$_.setCustomValidity($1)", err)
}

func (*HTMLSelectElement) Blur() {
	macro.Rewrite("$_.blur()")
}

func (*HTMLSelectElement) Click() {
	macro.Rewrite("$_.click()")
}

func (*HTMLSelectElement) DragDrop() (b bool) {
	macro.Rewrite("$_.dragDrop()")
	return b
}

func (*HTMLSelectElement) Focus() {
	macro.Rewrite("$_.focus()")
}

func (*HTMLSelectElement) GetElementsByClassName(classNames string) (w *dom.NodeList) {
	macro.Rewrite("$_.getElementsByClassName($1)", classNames)
	return w
}

func (*HTMLSelectElement) InsertAdjacentElement(position string, insertedElement element.Element) (w element.Element) {
	macro.Rewrite("$_.insertAdjacentElement($1, $2)", position, insertedElement)
	return w
}

func (*HTMLSelectElement) InsertAdjacentHTML(where string, html string) {
	macro.Rewrite("$_.insertAdjacentHTML($1, $2)", where, html)
}

func (*HTMLSelectElement) InsertAdjacentText(where string, text string) {
	macro.Rewrite("$_.insertAdjacentText($1, $2)", where, text)
}

func (*HTMLSelectElement) MsGetInputContext() (w *ms.MSInputMethodContext) {
	macro.Rewrite("$_.msGetInputContext()")
	return w
}

func (*HTMLSelectElement) ScrollIntoView(top *bool) {
	macro.Rewrite("$_.scrollIntoView($1)", top)
}

func (*HTMLSelectElement) GetAttribute(qualifiedName string) (s string) {
	macro.Rewrite("$_.getAttribute($1)", qualifiedName)
	return s
}

func (*HTMLSelectElement) GetAttributeNode(name string) (w *dom.Attr) {
	macro.Rewrite("$_.getAttributeNode($1)", name)
	return w
}

func (*HTMLSelectElement) GetAttributeNodeNS(namespaceURI string, localName string) (w *dom.Attr) {
	macro.Rewrite("$_.getAttributeNodeNS($1, $2)", namespaceURI, localName)
	return w
}

func (*HTMLSelectElement) GetAttributeNS(namespaceURI string, localName string) (s string) {
	macro.Rewrite("$_.getAttributeNS($1, $2)", namespaceURI, localName)
	return s
}

func (*HTMLSelectElement) GetBoundingClientRect() (c *dom.ClientRect) {
	macro.Rewrite("$_.getBoundingClientRect()")
	return c
}

func (*HTMLSelectElement) GetClientRects() (c *dom.ClientRectList) {
	macro.Rewrite("$_.getClientRects()")
	return c
}

func (*HTMLSelectElement) GetElementsByTagName(name string) (w *dom.NodeList) {
	macro.Rewrite("$_.getElementsByTagName($1)", name)
	return w
}

func (*HTMLSelectElement) GetElementsByTagNameNS(namespaceURI string, localName string) (w *dom.NodeList) {
	macro.Rewrite("$_.getElementsByTagNameNS($1, $2)", namespaceURI, localName)
	return w
}

func (*HTMLSelectElement) HasAttribute(name string) (b bool) {
	macro.Rewrite("$_.hasAttribute($1)", name)
	return b
}

func (*HTMLSelectElement) HasAttributeNS(namespaceURI string, localName string) (b bool) {
	macro.Rewrite("$_.hasAttributeNS($1, $2)", namespaceURI, localName)
	return b
}

func (*HTMLSelectElement) MsGetRegionContent() (w *ms.MSRangeCollection) {
	macro.Rewrite("$_.msGetRegionContent()")
	return w
}

func (*HTMLSelectElement) MsGetUntransformedBounds() (c *dom.ClientRect) {
	macro.Rewrite("$_.msGetUntransformedBounds()")
	return c
}

func (*HTMLSelectElement) MsMatchesSelector(selectors string) (b bool) {
	macro.Rewrite("$_.msMatchesSelector($1)", selectors)
	return b
}

func (*HTMLSelectElement) MsReleasePointerCapture(pointerId int) {
	macro.Rewrite("$_.msReleasePointerCapture($1)", pointerId)
}

func (*HTMLSelectElement) MsSetPointerCapture(pointerId int) {
	macro.Rewrite("$_.msSetPointerCapture($1)", pointerId)
}

func (*HTMLSelectElement) MsZoomTo(args *ms.MsZoomToOptions) {
	macro.Rewrite("$_.msZoomTo($1)", args)
}

func (*HTMLSelectElement) ReleasePointerCapture(pointerId int) {
	macro.Rewrite("$_.releasePointerCapture($1)", pointerId)
}

func (*HTMLSelectElement) RemoveAttribute(qualifiedName string) {
	macro.Rewrite("$_.removeAttribute($1)", qualifiedName)
}

func (*HTMLSelectElement) RemoveAttributeNode(oldAttr *dom.Attr) (w *dom.Attr) {
	macro.Rewrite("$_.removeAttributeNode($1)", oldAttr)
	return w
}

func (*HTMLSelectElement) RemoveAttributeNS(namespaceURI string, localName string) {
	macro.Rewrite("$_.removeAttributeNS($1, $2)", namespaceURI, localName)
}

func (*HTMLSelectElement) RequestFullscreen() {
	macro.Rewrite("$_.requestFullscreen()")
}

func (*HTMLSelectElement) RequestPointerLock() {
	macro.Rewrite("$_.requestPointerLock()")
}

func (*HTMLSelectElement) SetAttribute(qualifiedName string, value string) {
	macro.Rewrite("$_.setAttribute($1, $2)", qualifiedName, value)
}

func (*HTMLSelectElement) SetAttributeNode(newAttr *dom.Attr) (w *dom.Attr) {
	macro.Rewrite("$_.setAttributeNode($1)", newAttr)
	return w
}

func (*HTMLSelectElement) SetAttributeNodeNS(newAttr *dom.Attr) (w *dom.Attr) {
	macro.Rewrite("$_.setAttributeNodeNS($1)", newAttr)
	return w
}

func (*HTMLSelectElement) SetAttributeNS(namespaceURI string, qualifiedName string, value string) {
	macro.Rewrite("$_.setAttributeNS($1, $2, $3)", namespaceURI, qualifiedName, value)
}

func (*HTMLSelectElement) SetPointerCapture(pointerId int) {
	macro.Rewrite("$_.setPointerCapture($1)", pointerId)
}

func (*HTMLSelectElement) WebkitMatchesSelector(selectors string) (b bool) {
	macro.Rewrite("$_.webkitMatchesSelector($1)", selectors)
	return b
}

func (*HTMLSelectElement) WebkitRequestFullscreen() {
	macro.Rewrite("$_.webkitRequestFullscreen()")
}

func (*HTMLSelectElement) WebkitRequestFullScreen() {
	macro.Rewrite("$_.webkitRequestFullScreen()")
}

func (*HTMLSelectElement) QuerySelector(selectors string) (w element.Element) {
	macro.Rewrite("$_.querySelector($1)", selectors)
	return w
}

func (*HTMLSelectElement) QuerySelectorAll(selectors string) (w *dom.NodeList) {
	macro.Rewrite("$_.querySelectorAll($1)", selectors)
	return w
}

func (*HTMLSelectElement) AppendChild(newChild dom.Node) (w dom.Node) {
	macro.Rewrite("$_.appendChild($1)", newChild)
	return w
}

func (*HTMLSelectElement) CloneNode(deep *bool) (w dom.Node) {
	macro.Rewrite("$_.cloneNode($1)", deep)
	return w
}

func (*HTMLSelectElement) CompareDocumentPosition(other dom.Node) (u uint8) {
	macro.Rewrite("$_.compareDocumentPosition($1)", other)
	return u
}

func (*HTMLSelectElement) Contains(child dom.Node) (b bool) {
	macro.Rewrite("$_.contains($1)", child)
	return b
}

func (*HTMLSelectElement) HasAttributes() (b bool) {
	macro.Rewrite("$_.hasAttributes()")
	return b
}

func (*HTMLSelectElement) HasChildNodes() (b bool) {
	macro.Rewrite("$_.hasChildNodes()")
	return b
}

func (*HTMLSelectElement) InsertBefore(newChild dom.Node, refChild *dom.Node) (w dom.Node) {
	macro.Rewrite("$_.insertBefore($1, $2)", newChild, refChild)
	return w
}

func (*HTMLSelectElement) IsDefaultNamespace(namespaceURI string) (b bool) {
	macro.Rewrite("$_.isDefaultNamespace($1)", namespaceURI)
	return b
}

func (*HTMLSelectElement) IsEqualNode(arg dom.Node) (b bool) {
	macro.Rewrite("$_.isEqualNode($1)", arg)
	return b
}

func (*HTMLSelectElement) IsSameNode(other dom.Node) (b bool) {
	macro.Rewrite("$_.isSameNode($1)", other)
	return b
}

func (*HTMLSelectElement) LookupNamespaceURI(prefix string) (s string) {
	macro.Rewrite("$_.lookupNamespaceURI($1)", prefix)
	return s
}

func (*HTMLSelectElement) LookupPrefix(namespaceURI string) (s string) {
	macro.Rewrite("$_.lookupPrefix($1)", namespaceURI)
	return s
}

func (*HTMLSelectElement) Normalize() {
	macro.Rewrite("$_.normalize()")
}

func (*HTMLSelectElement) RemoveChild(oldChild dom.Node) (w dom.Node) {
	macro.Rewrite("$_.removeChild($1)", oldChild)
	return w
}

func (*HTMLSelectElement) ReplaceChild(newChild dom.Node, oldChild dom.Node) (w dom.Node) {
	macro.Rewrite("$_.replaceChild($1, $2)", newChild, oldChild)
	return w
}

func (*HTMLSelectElement) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*HTMLSelectElement) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*HTMLSelectElement) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*HTMLSelectElement) Autofocus() (autofocus bool) {
	macro.Rewrite("$_.autofocus")
	return autofocus
}

func (*HTMLSelectElement) SetAutofocus(autofocus bool) {
	macro.Rewrite("$_.autofocus = $1", autofocus)
}

func (*HTMLSelectElement) Disabled() (disabled bool) {
	macro.Rewrite("$_.disabled")
	return disabled
}

func (*HTMLSelectElement) SetDisabled(disabled bool) {
	macro.Rewrite("$_.disabled = $1", disabled)
}

func (*HTMLSelectElement) Form() (form *HTMLFormElement) {
	macro.Rewrite("$_.form")
	return form
}

func (*HTMLSelectElement) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}

func (*HTMLSelectElement) SetLength(length uint) {
	macro.Rewrite("$_.length = $1", length)
}

func (*HTMLSelectElement) Multiple() (multiple bool) {
	macro.Rewrite("$_.multiple")
	return multiple
}

func (*HTMLSelectElement) SetMultiple(multiple bool) {
	macro.Rewrite("$_.multiple = $1", multiple)
}

func (*HTMLSelectElement) Name() (name string) {
	macro.Rewrite("$_.name")
	return name
}

func (*HTMLSelectElement) SetName(name string) {
	macro.Rewrite("$_.name = $1", name)
}

func (*HTMLSelectElement) Options() (options *HTMLOptionsCollection) {
	macro.Rewrite("$_.options")
	return options
}

func (*HTMLSelectElement) Required() (required bool) {
	macro.Rewrite("$_.required")
	return required
}

func (*HTMLSelectElement) SetRequired(required bool) {
	macro.Rewrite("$_.required = $1", required)
}

func (*HTMLSelectElement) SelectedIndex() (selectedIndex int) {
	macro.Rewrite("$_.selectedIndex")
	return selectedIndex
}

func (*HTMLSelectElement) SetSelectedIndex(selectedIndex int) {
	macro.Rewrite("$_.selectedIndex = $1", selectedIndex)
}

func (*HTMLSelectElement) SelectedOptions() (selectedOptions HTMLCollection) {
	macro.Rewrite("$_.selectedOptions")
	return selectedOptions
}

func (*HTMLSelectElement) Size() (size uint) {
	macro.Rewrite("$_.size")
	return size
}

func (*HTMLSelectElement) SetSize(size uint) {
	macro.Rewrite("$_.size = $1", size)
}

func (*HTMLSelectElement) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

func (*HTMLSelectElement) ValidationMessage() (validationMessage string) {
	macro.Rewrite("$_.validationMessage")
	return validationMessage
}

func (*HTMLSelectElement) Validity() (validity *utils.ValidityState) {
	macro.Rewrite("$_.validity")
	return validity
}

func (*HTMLSelectElement) Value() (value string) {
	macro.Rewrite("$_.value")
	return value
}

func (*HTMLSelectElement) SetValue(value string) {
	macro.Rewrite("$_.value = $1", value)
}

func (*HTMLSelectElement) WillValidate() (willValidate bool) {
	macro.Rewrite("$_.willValidate")
	return willValidate
}

func (*HTMLSelectElement) AccessKey() (accessKey string) {
	macro.Rewrite("$_.accessKey")
	return accessKey
}

func (*HTMLSelectElement) SetAccessKey(accessKey string) {
	macro.Rewrite("$_.accessKey = $1", accessKey)
}

func (*HTMLSelectElement) Children() (children HTMLCollection) {
	macro.Rewrite("$_.children")
	return children
}

func (*HTMLSelectElement) ContentEditable() (contentEditable string) {
	macro.Rewrite("$_.contentEditable")
	return contentEditable
}

func (*HTMLSelectElement) SetContentEditable(contentEditable string) {
	macro.Rewrite("$_.contentEditable = $1", contentEditable)
}

func (*HTMLSelectElement) Dataset() (dataset *dom.DOMStringMap) {
	macro.Rewrite("$_.dataset")
	return dataset
}

func (*HTMLSelectElement) Dir() (dir string) {
	macro.Rewrite("$_.dir")
	return dir
}

func (*HTMLSelectElement) SetDir(dir string) {
	macro.Rewrite("$_.dir = $1", dir)
}

func (*HTMLSelectElement) Draggable() (draggable bool) {
	macro.Rewrite("$_.draggable")
	return draggable
}

func (*HTMLSelectElement) SetDraggable(draggable bool) {
	macro.Rewrite("$_.draggable = $1", draggable)
}

func (*HTMLSelectElement) Hidden() (hidden bool) {
	macro.Rewrite("$_.hidden")
	return hidden
}

func (*HTMLSelectElement) SetHidden(hidden bool) {
	macro.Rewrite("$_.hidden = $1", hidden)
}

func (*HTMLSelectElement) HideFocus() (hideFocus bool) {
	macro.Rewrite("$_.hideFocus")
	return hideFocus
}

func (*HTMLSelectElement) SetHideFocus(hideFocus bool) {
	macro.Rewrite("$_.hideFocus = $1", hideFocus)
}

func (*HTMLSelectElement) InnerText() (innerText string) {
	macro.Rewrite("$_.innerText")
	return innerText
}

func (*HTMLSelectElement) SetInnerText(innerText string) {
	macro.Rewrite("$_.innerText = $1", innerText)
}

func (*HTMLSelectElement) IsContentEditable() (isContentEditable bool) {
	macro.Rewrite("$_.isContentEditable")
	return isContentEditable
}

func (*HTMLSelectElement) Lang() (lang string) {
	macro.Rewrite("$_.lang")
	return lang
}

func (*HTMLSelectElement) SetLang(lang string) {
	macro.Rewrite("$_.lang = $1", lang)
}

func (*HTMLSelectElement) OffsetHeight() (offsetHeight int) {
	macro.Rewrite("$_.offsetHeight")
	return offsetHeight
}

func (*HTMLSelectElement) OffsetLeft() (offsetLeft int) {
	macro.Rewrite("$_.offsetLeft")
	return offsetLeft
}

func (*HTMLSelectElement) OffsetParent() (offsetParent element.Element) {
	macro.Rewrite("$_.offsetParent")
	return offsetParent
}

func (*HTMLSelectElement) OffsetTop() (offsetTop int) {
	macro.Rewrite("$_.offsetTop")
	return offsetTop
}

func (*HTMLSelectElement) OffsetWidth() (offsetWidth int) {
	macro.Rewrite("$_.offsetWidth")
	return offsetWidth
}

func (*HTMLSelectElement) Onabort() (onabort func(event.Event)) {
	macro.Rewrite("$_.onabort")
	return onabort
}

func (*HTMLSelectElement) SetOnabort(onabort func(event.Event)) {
	macro.Rewrite("$_.onabort = $1", onabort)
}

func (*HTMLSelectElement) Onactivate() (onactivate func(ui.UIEvent)) {
	macro.Rewrite("$_.onactivate")
	return onactivate
}

func (*HTMLSelectElement) SetOnactivate(onactivate func(ui.UIEvent)) {
	macro.Rewrite("$_.onactivate = $1", onactivate)
}

func (*HTMLSelectElement) Onbeforeactivate() (onbeforeactivate func(ui.UIEvent)) {
	macro.Rewrite("$_.onbeforeactivate")
	return onbeforeactivate
}

func (*HTMLSelectElement) SetOnbeforeactivate(onbeforeactivate func(ui.UIEvent)) {
	macro.Rewrite("$_.onbeforeactivate = $1", onbeforeactivate)
}

func (*HTMLSelectElement) Onbeforecopy() (onbeforecopy func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.onbeforecopy")
	return onbeforecopy
}

func (*HTMLSelectElement) SetOnbeforecopy(onbeforecopy func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.onbeforecopy = $1", onbeforecopy)
}

func (*HTMLSelectElement) Onbeforecut() (onbeforecut func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.onbeforecut")
	return onbeforecut
}

func (*HTMLSelectElement) SetOnbeforecut(onbeforecut func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.onbeforecut = $1", onbeforecut)
}

func (*HTMLSelectElement) Onbeforedeactivate() (onbeforedeactivate func(ui.UIEvent)) {
	macro.Rewrite("$_.onbeforedeactivate")
	return onbeforedeactivate
}

func (*HTMLSelectElement) SetOnbeforedeactivate(onbeforedeactivate func(ui.UIEvent)) {
	macro.Rewrite("$_.onbeforedeactivate = $1", onbeforedeactivate)
}

func (*HTMLSelectElement) Onbeforepaste() (onbeforepaste func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.onbeforepaste")
	return onbeforepaste
}

func (*HTMLSelectElement) SetOnbeforepaste(onbeforepaste func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.onbeforepaste = $1", onbeforepaste)
}

func (*HTMLSelectElement) Onblur() (onblur func(*utils.FocusEvent)) {
	macro.Rewrite("$_.onblur")
	return onblur
}

func (*HTMLSelectElement) SetOnblur(onblur func(*utils.FocusEvent)) {
	macro.Rewrite("$_.onblur = $1", onblur)
}

func (*HTMLSelectElement) Oncanplay() (oncanplay func(event.Event)) {
	macro.Rewrite("$_.oncanplay")
	return oncanplay
}

func (*HTMLSelectElement) SetOncanplay(oncanplay func(event.Event)) {
	macro.Rewrite("$_.oncanplay = $1", oncanplay)
}

func (*HTMLSelectElement) Oncanplaythrough() (oncanplaythrough func(event.Event)) {
	macro.Rewrite("$_.oncanplaythrough")
	return oncanplaythrough
}

func (*HTMLSelectElement) SetOncanplaythrough(oncanplaythrough func(event.Event)) {
	macro.Rewrite("$_.oncanplaythrough = $1", oncanplaythrough)
}

func (*HTMLSelectElement) Onchange() (onchange func(event.Event)) {
	macro.Rewrite("$_.onchange")
	return onchange
}

func (*HTMLSelectElement) SetOnchange(onchange func(event.Event)) {
	macro.Rewrite("$_.onchange = $1", onchange)
}

func (*HTMLSelectElement) Onclick() (onclick func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onclick")
	return onclick
}

func (*HTMLSelectElement) SetOnclick(onclick func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onclick = $1", onclick)
}

func (*HTMLSelectElement) Oncontextmenu() (oncontextmenu func(*utils.PointerEvent)) {
	macro.Rewrite("$_.oncontextmenu")
	return oncontextmenu
}

func (*HTMLSelectElement) SetOncontextmenu(oncontextmenu func(*utils.PointerEvent)) {
	macro.Rewrite("$_.oncontextmenu = $1", oncontextmenu)
}

func (*HTMLSelectElement) Oncopy() (oncopy func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.oncopy")
	return oncopy
}

func (*HTMLSelectElement) SetOncopy(oncopy func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.oncopy = $1", oncopy)
}

func (*HTMLSelectElement) Oncuechange() (oncuechange func(event.Event)) {
	macro.Rewrite("$_.oncuechange")
	return oncuechange
}

func (*HTMLSelectElement) SetOncuechange(oncuechange func(event.Event)) {
	macro.Rewrite("$_.oncuechange = $1", oncuechange)
}

func (*HTMLSelectElement) Oncut() (oncut func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.oncut")
	return oncut
}

func (*HTMLSelectElement) SetOncut(oncut func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.oncut = $1", oncut)
}

func (*HTMLSelectElement) Ondblclick() (ondblclick func(mouse.MouseEvent)) {
	macro.Rewrite("$_.ondblclick")
	return ondblclick
}

func (*HTMLSelectElement) SetOndblclick(ondblclick func(mouse.MouseEvent)) {
	macro.Rewrite("$_.ondblclick = $1", ondblclick)
}

func (*HTMLSelectElement) Ondeactivate() (ondeactivate func(ui.UIEvent)) {
	macro.Rewrite("$_.ondeactivate")
	return ondeactivate
}

func (*HTMLSelectElement) SetOndeactivate(ondeactivate func(ui.UIEvent)) {
	macro.Rewrite("$_.ondeactivate = $1", ondeactivate)
}

func (*HTMLSelectElement) Ondrag() (ondrag func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondrag")
	return ondrag
}

func (*HTMLSelectElement) SetOndrag(ondrag func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondrag = $1", ondrag)
}

func (*HTMLSelectElement) Ondragend() (ondragend func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragend")
	return ondragend
}

func (*HTMLSelectElement) SetOndragend(ondragend func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragend = $1", ondragend)
}

func (*HTMLSelectElement) Ondragenter() (ondragenter func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragenter")
	return ondragenter
}

func (*HTMLSelectElement) SetOndragenter(ondragenter func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragenter = $1", ondragenter)
}

func (*HTMLSelectElement) Ondragleave() (ondragleave func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragleave")
	return ondragleave
}

func (*HTMLSelectElement) SetOndragleave(ondragleave func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragleave = $1", ondragleave)
}

func (*HTMLSelectElement) Ondragover() (ondragover func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragover")
	return ondragover
}

func (*HTMLSelectElement) SetOndragover(ondragover func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragover = $1", ondragover)
}

func (*HTMLSelectElement) Ondragstart() (ondragstart func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragstart")
	return ondragstart
}

func (*HTMLSelectElement) SetOndragstart(ondragstart func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragstart = $1", ondragstart)
}

func (*HTMLSelectElement) Ondrop() (ondrop func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondrop")
	return ondrop
}

func (*HTMLSelectElement) SetOndrop(ondrop func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondrop = $1", ondrop)
}

func (*HTMLSelectElement) Ondurationchange() (ondurationchange func(event.Event)) {
	macro.Rewrite("$_.ondurationchange")
	return ondurationchange
}

func (*HTMLSelectElement) SetOndurationchange(ondurationchange func(event.Event)) {
	macro.Rewrite("$_.ondurationchange = $1", ondurationchange)
}

func (*HTMLSelectElement) Onemptied() (onemptied func(event.Event)) {
	macro.Rewrite("$_.onemptied")
	return onemptied
}

func (*HTMLSelectElement) SetOnemptied(onemptied func(event.Event)) {
	macro.Rewrite("$_.onemptied = $1", onemptied)
}

func (*HTMLSelectElement) Onended() (onended func(event.Event)) {
	macro.Rewrite("$_.onended")
	return onended
}

func (*HTMLSelectElement) SetOnended(onended func(event.Event)) {
	macro.Rewrite("$_.onended = $1", onended)
}

func (*HTMLSelectElement) Onerror() (onerror func(event.Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

func (*HTMLSelectElement) SetOnerror(onerror func(event.Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

func (*HTMLSelectElement) Onfocus() (onfocus func(*utils.FocusEvent)) {
	macro.Rewrite("$_.onfocus")
	return onfocus
}

func (*HTMLSelectElement) SetOnfocus(onfocus func(*utils.FocusEvent)) {
	macro.Rewrite("$_.onfocus = $1", onfocus)
}

func (*HTMLSelectElement) Oninput() (oninput func(event.Event)) {
	macro.Rewrite("$_.oninput")
	return oninput
}

func (*HTMLSelectElement) SetOninput(oninput func(event.Event)) {
	macro.Rewrite("$_.oninput = $1", oninput)
}

func (*HTMLSelectElement) Oninvalid() (oninvalid func(event.Event)) {
	macro.Rewrite("$_.oninvalid")
	return oninvalid
}

func (*HTMLSelectElement) SetOninvalid(oninvalid func(event.Event)) {
	macro.Rewrite("$_.oninvalid = $1", oninvalid)
}

func (*HTMLSelectElement) Onkeydown() (onkeydown func(*utils.KeyboardEvent)) {
	macro.Rewrite("$_.onkeydown")
	return onkeydown
}

func (*HTMLSelectElement) SetOnkeydown(onkeydown func(*utils.KeyboardEvent)) {
	macro.Rewrite("$_.onkeydown = $1", onkeydown)
}

func (*HTMLSelectElement) Onkeypress() (onkeypress func(*utils.KeyboardEvent)) {
	macro.Rewrite("$_.onkeypress")
	return onkeypress
}

func (*HTMLSelectElement) SetOnkeypress(onkeypress func(*utils.KeyboardEvent)) {
	macro.Rewrite("$_.onkeypress = $1", onkeypress)
}

func (*HTMLSelectElement) Onkeyup() (onkeyup func(*utils.KeyboardEvent)) {
	macro.Rewrite("$_.onkeyup")
	return onkeyup
}

func (*HTMLSelectElement) SetOnkeyup(onkeyup func(*utils.KeyboardEvent)) {
	macro.Rewrite("$_.onkeyup = $1", onkeyup)
}

func (*HTMLSelectElement) Onload() (onload func(event.Event)) {
	macro.Rewrite("$_.onload")
	return onload
}

func (*HTMLSelectElement) SetOnload(onload func(event.Event)) {
	macro.Rewrite("$_.onload = $1", onload)
}

func (*HTMLSelectElement) Onloadeddata() (onloadeddata func(event.Event)) {
	macro.Rewrite("$_.onloadeddata")
	return onloadeddata
}

func (*HTMLSelectElement) SetOnloadeddata(onloadeddata func(event.Event)) {
	macro.Rewrite("$_.onloadeddata = $1", onloadeddata)
}

func (*HTMLSelectElement) Onloadedmetadata() (onloadedmetadata func(event.Event)) {
	macro.Rewrite("$_.onloadedmetadata")
	return onloadedmetadata
}

func (*HTMLSelectElement) SetOnloadedmetadata(onloadedmetadata func(event.Event)) {
	macro.Rewrite("$_.onloadedmetadata = $1", onloadedmetadata)
}

func (*HTMLSelectElement) Onloadstart() (onloadstart func(event.Event)) {
	macro.Rewrite("$_.onloadstart")
	return onloadstart
}

func (*HTMLSelectElement) SetOnloadstart(onloadstart func(event.Event)) {
	macro.Rewrite("$_.onloadstart = $1", onloadstart)
}

func (*HTMLSelectElement) Onmousedown() (onmousedown func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmousedown")
	return onmousedown
}

func (*HTMLSelectElement) SetOnmousedown(onmousedown func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmousedown = $1", onmousedown)
}

func (*HTMLSelectElement) Onmouseenter() (onmouseenter func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseenter")
	return onmouseenter
}

func (*HTMLSelectElement) SetOnmouseenter(onmouseenter func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseenter = $1", onmouseenter)
}

func (*HTMLSelectElement) Onmouseleave() (onmouseleave func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseleave")
	return onmouseleave
}

func (*HTMLSelectElement) SetOnmouseleave(onmouseleave func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseleave = $1", onmouseleave)
}

func (*HTMLSelectElement) Onmousemove() (onmousemove func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmousemove")
	return onmousemove
}

func (*HTMLSelectElement) SetOnmousemove(onmousemove func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmousemove = $1", onmousemove)
}

func (*HTMLSelectElement) Onmouseout() (onmouseout func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseout")
	return onmouseout
}

func (*HTMLSelectElement) SetOnmouseout(onmouseout func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseout = $1", onmouseout)
}

func (*HTMLSelectElement) Onmouseover() (onmouseover func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseover")
	return onmouseover
}

func (*HTMLSelectElement) SetOnmouseover(onmouseover func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseover = $1", onmouseover)
}

func (*HTMLSelectElement) Onmouseup() (onmouseup func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseup")
	return onmouseup
}

func (*HTMLSelectElement) SetOnmouseup(onmouseup func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseup = $1", onmouseup)
}

func (*HTMLSelectElement) Onmousewheel() (onmousewheel func(*utils.WheelEvent)) {
	macro.Rewrite("$_.onmousewheel")
	return onmousewheel
}

func (*HTMLSelectElement) SetOnmousewheel(onmousewheel func(*utils.WheelEvent)) {
	macro.Rewrite("$_.onmousewheel = $1", onmousewheel)
}

func (*HTMLSelectElement) Onmscontentzoom() (onmscontentzoom func(ui.UIEvent)) {
	macro.Rewrite("$_.onmscontentzoom")
	return onmscontentzoom
}

func (*HTMLSelectElement) SetOnmscontentzoom(onmscontentzoom func(ui.UIEvent)) {
	macro.Rewrite("$_.onmscontentzoom = $1", onmscontentzoom)
}

func (*HTMLSelectElement) Onmsmanipulationstatechanged() (onmsmanipulationstatechanged func(*ms.MSManipulationEvent)) {
	macro.Rewrite("$_.onmsmanipulationstatechanged")
	return onmsmanipulationstatechanged
}

func (*HTMLSelectElement) SetOnmsmanipulationstatechanged(onmsmanipulationstatechanged func(*ms.MSManipulationEvent)) {
	macro.Rewrite("$_.onmsmanipulationstatechanged = $1", onmsmanipulationstatechanged)
}

func (*HTMLSelectElement) Onpaste() (onpaste func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.onpaste")
	return onpaste
}

func (*HTMLSelectElement) SetOnpaste(onpaste func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.onpaste = $1", onpaste)
}

func (*HTMLSelectElement) Onpause() (onpause func(event.Event)) {
	macro.Rewrite("$_.onpause")
	return onpause
}

func (*HTMLSelectElement) SetOnpause(onpause func(event.Event)) {
	macro.Rewrite("$_.onpause = $1", onpause)
}

func (*HTMLSelectElement) Onplay() (onplay func(event.Event)) {
	macro.Rewrite("$_.onplay")
	return onplay
}

func (*HTMLSelectElement) SetOnplay(onplay func(event.Event)) {
	macro.Rewrite("$_.onplay = $1", onplay)
}

func (*HTMLSelectElement) Onplaying() (onplaying func(event.Event)) {
	macro.Rewrite("$_.onplaying")
	return onplaying
}

func (*HTMLSelectElement) SetOnplaying(onplaying func(event.Event)) {
	macro.Rewrite("$_.onplaying = $1", onplaying)
}

func (*HTMLSelectElement) Onprogress() (onprogress func(event.Event)) {
	macro.Rewrite("$_.onprogress")
	return onprogress
}

func (*HTMLSelectElement) SetOnprogress(onprogress func(event.Event)) {
	macro.Rewrite("$_.onprogress = $1", onprogress)
}

func (*HTMLSelectElement) Onratechange() (onratechange func(event.Event)) {
	macro.Rewrite("$_.onratechange")
	return onratechange
}

func (*HTMLSelectElement) SetOnratechange(onratechange func(event.Event)) {
	macro.Rewrite("$_.onratechange = $1", onratechange)
}

func (*HTMLSelectElement) Onreset() (onreset func(event.Event)) {
	macro.Rewrite("$_.onreset")
	return onreset
}

func (*HTMLSelectElement) SetOnreset(onreset func(event.Event)) {
	macro.Rewrite("$_.onreset = $1", onreset)
}

func (*HTMLSelectElement) Onscroll() (onscroll func(ui.UIEvent)) {
	macro.Rewrite("$_.onscroll")
	return onscroll
}

func (*HTMLSelectElement) SetOnscroll(onscroll func(ui.UIEvent)) {
	macro.Rewrite("$_.onscroll = $1", onscroll)
}

func (*HTMLSelectElement) Onseeked() (onseeked func(event.Event)) {
	macro.Rewrite("$_.onseeked")
	return onseeked
}

func (*HTMLSelectElement) SetOnseeked(onseeked func(event.Event)) {
	macro.Rewrite("$_.onseeked = $1", onseeked)
}

func (*HTMLSelectElement) Onseeking() (onseeking func(event.Event)) {
	macro.Rewrite("$_.onseeking")
	return onseeking
}

func (*HTMLSelectElement) SetOnseeking(onseeking func(event.Event)) {
	macro.Rewrite("$_.onseeking = $1", onseeking)
}

func (*HTMLSelectElement) Onselect() (onselect func(ui.UIEvent)) {
	macro.Rewrite("$_.onselect")
	return onselect
}

func (*HTMLSelectElement) SetOnselect(onselect func(ui.UIEvent)) {
	macro.Rewrite("$_.onselect = $1", onselect)
}

func (*HTMLSelectElement) Onselectstart() (onselectstart func(event.Event)) {
	macro.Rewrite("$_.onselectstart")
	return onselectstart
}

func (*HTMLSelectElement) SetOnselectstart(onselectstart func(event.Event)) {
	macro.Rewrite("$_.onselectstart = $1", onselectstart)
}

func (*HTMLSelectElement) Onstalled() (onstalled func(event.Event)) {
	macro.Rewrite("$_.onstalled")
	return onstalled
}

func (*HTMLSelectElement) SetOnstalled(onstalled func(event.Event)) {
	macro.Rewrite("$_.onstalled = $1", onstalled)
}

func (*HTMLSelectElement) Onsubmit() (onsubmit func(event.Event)) {
	macro.Rewrite("$_.onsubmit")
	return onsubmit
}

func (*HTMLSelectElement) SetOnsubmit(onsubmit func(event.Event)) {
	macro.Rewrite("$_.onsubmit = $1", onsubmit)
}

func (*HTMLSelectElement) Onsuspend() (onsuspend func(event.Event)) {
	macro.Rewrite("$_.onsuspend")
	return onsuspend
}

func (*HTMLSelectElement) SetOnsuspend(onsuspend func(event.Event)) {
	macro.Rewrite("$_.onsuspend = $1", onsuspend)
}

func (*HTMLSelectElement) Ontimeupdate() (ontimeupdate func(event.Event)) {
	macro.Rewrite("$_.ontimeupdate")
	return ontimeupdate
}

func (*HTMLSelectElement) SetOntimeupdate(ontimeupdate func(event.Event)) {
	macro.Rewrite("$_.ontimeupdate = $1", ontimeupdate)
}

func (*HTMLSelectElement) Onvolumechange() (onvolumechange func(event.Event)) {
	macro.Rewrite("$_.onvolumechange")
	return onvolumechange
}

func (*HTMLSelectElement) SetOnvolumechange(onvolumechange func(event.Event)) {
	macro.Rewrite("$_.onvolumechange = $1", onvolumechange)
}

func (*HTMLSelectElement) Onwaiting() (onwaiting func(event.Event)) {
	macro.Rewrite("$_.onwaiting")
	return onwaiting
}

func (*HTMLSelectElement) SetOnwaiting(onwaiting func(event.Event)) {
	macro.Rewrite("$_.onwaiting = $1", onwaiting)
}

func (*HTMLSelectElement) OuterText() (outerText string) {
	macro.Rewrite("$_.outerText")
	return outerText
}

func (*HTMLSelectElement) SetOuterText(outerText string) {
	macro.Rewrite("$_.outerText = $1", outerText)
}

func (*HTMLSelectElement) Spellcheck() (spellcheck bool) {
	macro.Rewrite("$_.spellcheck")
	return spellcheck
}

func (*HTMLSelectElement) SetSpellcheck(spellcheck bool) {
	macro.Rewrite("$_.spellcheck = $1", spellcheck)
}

func (*HTMLSelectElement) Style() (style *css.CSSStyleDeclaration) {
	macro.Rewrite("$_.style")
	return style
}

func (*HTMLSelectElement) TabIndex() (tabIndex int8) {
	macro.Rewrite("$_.tabIndex")
	return tabIndex
}

func (*HTMLSelectElement) SetTabIndex(tabIndex int8) {
	macro.Rewrite("$_.tabIndex = $1", tabIndex)
}

func (*HTMLSelectElement) Title() (title string) {
	macro.Rewrite("$_.title")
	return title
}

func (*HTMLSelectElement) SetTitle(title string) {
	macro.Rewrite("$_.title = $1", title)
}

func (*HTMLSelectElement) ClassList() (classList dom.DOMTokenList) {
	macro.Rewrite("$_.classList")
	return classList
}

func (*HTMLSelectElement) ClassName() (className string) {
	macro.Rewrite("$_.className")
	return className
}

func (*HTMLSelectElement) SetClassName(className string) {
	macro.Rewrite("$_.className = $1", className)
}

func (*HTMLSelectElement) ClientHeight() (clientHeight int) {
	macro.Rewrite("$_.clientHeight")
	return clientHeight
}

func (*HTMLSelectElement) ClientLeft() (clientLeft int) {
	macro.Rewrite("$_.clientLeft")
	return clientLeft
}

func (*HTMLSelectElement) ClientTop() (clientTop int) {
	macro.Rewrite("$_.clientTop")
	return clientTop
}

func (*HTMLSelectElement) ClientWidth() (clientWidth int) {
	macro.Rewrite("$_.clientWidth")
	return clientWidth
}

func (*HTMLSelectElement) ID() (id string) {
	macro.Rewrite("$_.id")
	return id
}

func (*HTMLSelectElement) SetID(id string) {
	macro.Rewrite("$_.id = $1", id)
}

func (*HTMLSelectElement) InnerHTML() (innerHTML string) {
	macro.Rewrite("$_.innerHTML")
	return innerHTML
}

func (*HTMLSelectElement) SetInnerHTML(innerHTML string) {
	macro.Rewrite("$_.innerHTML = $1", innerHTML)
}

func (*HTMLSelectElement) MsContentZoomFactor() (msContentZoomFactor float32) {
	macro.Rewrite("$_.msContentZoomFactor")
	return msContentZoomFactor
}

func (*HTMLSelectElement) SetMsContentZoomFactor(msContentZoomFactor float32) {
	macro.Rewrite("$_.msContentZoomFactor = $1", msContentZoomFactor)
}

func (*HTMLSelectElement) MsRegionOverflow() (msRegionOverflow string) {
	macro.Rewrite("$_.msRegionOverflow")
	return msRegionOverflow
}

func (*HTMLSelectElement) Onariarequest() (onariarequest func(event.Event)) {
	macro.Rewrite("$_.onariarequest")
	return onariarequest
}

func (*HTMLSelectElement) SetOnariarequest(onariarequest func(event.Event)) {
	macro.Rewrite("$_.onariarequest = $1", onariarequest)
}

func (*HTMLSelectElement) Oncommand() (oncommand func(event.Event)) {
	macro.Rewrite("$_.oncommand")
	return oncommand
}

func (*HTMLSelectElement) SetOncommand(oncommand func(event.Event)) {
	macro.Rewrite("$_.oncommand = $1", oncommand)
}

func (*HTMLSelectElement) Ongotpointercapture() (ongotpointercapture func(*utils.PointerEvent)) {
	macro.Rewrite("$_.ongotpointercapture")
	return ongotpointercapture
}

func (*HTMLSelectElement) SetOngotpointercapture(ongotpointercapture func(*utils.PointerEvent)) {
	macro.Rewrite("$_.ongotpointercapture = $1", ongotpointercapture)
}

func (*HTMLSelectElement) Onlostpointercapture() (onlostpointercapture func(*utils.PointerEvent)) {
	macro.Rewrite("$_.onlostpointercapture")
	return onlostpointercapture
}

func (*HTMLSelectElement) SetOnlostpointercapture(onlostpointercapture func(*utils.PointerEvent)) {
	macro.Rewrite("$_.onlostpointercapture = $1", onlostpointercapture)
}

func (*HTMLSelectElement) Onmsgesturechange() (onmsgesturechange func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturechange")
	return onmsgesturechange
}

func (*HTMLSelectElement) SetOnmsgesturechange(onmsgesturechange func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturechange = $1", onmsgesturechange)
}

func (*HTMLSelectElement) Onmsgesturedoubletap() (onmsgesturedoubletap func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturedoubletap")
	return onmsgesturedoubletap
}

func (*HTMLSelectElement) SetOnmsgesturedoubletap(onmsgesturedoubletap func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturedoubletap = $1", onmsgesturedoubletap)
}

func (*HTMLSelectElement) Onmsgestureend() (onmsgestureend func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgestureend")
	return onmsgestureend
}

func (*HTMLSelectElement) SetOnmsgestureend(onmsgestureend func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgestureend = $1", onmsgestureend)
}

func (*HTMLSelectElement) Onmsgesturehold() (onmsgesturehold func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturehold")
	return onmsgesturehold
}

func (*HTMLSelectElement) SetOnmsgesturehold(onmsgesturehold func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturehold = $1", onmsgesturehold)
}

func (*HTMLSelectElement) Onmsgesturestart() (onmsgesturestart func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturestart")
	return onmsgesturestart
}

func (*HTMLSelectElement) SetOnmsgesturestart(onmsgesturestart func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturestart = $1", onmsgesturestart)
}

func (*HTMLSelectElement) Onmsgesturetap() (onmsgesturetap func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturetap")
	return onmsgesturetap
}

func (*HTMLSelectElement) SetOnmsgesturetap(onmsgesturetap func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturetap = $1", onmsgesturetap)
}

func (*HTMLSelectElement) Onmsgotpointercapture() (onmsgotpointercapture func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmsgotpointercapture")
	return onmsgotpointercapture
}

func (*HTMLSelectElement) SetOnmsgotpointercapture(onmsgotpointercapture func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmsgotpointercapture = $1", onmsgotpointercapture)
}

func (*HTMLSelectElement) Onmsinertiastart() (onmsinertiastart func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsinertiastart")
	return onmsinertiastart
}

func (*HTMLSelectElement) SetOnmsinertiastart(onmsinertiastart func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsinertiastart = $1", onmsinertiastart)
}

func (*HTMLSelectElement) Onmslostpointercapture() (onmslostpointercapture func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmslostpointercapture")
	return onmslostpointercapture
}

func (*HTMLSelectElement) SetOnmslostpointercapture(onmslostpointercapture func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmslostpointercapture = $1", onmslostpointercapture)
}

func (*HTMLSelectElement) Onmspointercancel() (onmspointercancel func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointercancel")
	return onmspointercancel
}

func (*HTMLSelectElement) SetOnmspointercancel(onmspointercancel func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointercancel = $1", onmspointercancel)
}

func (*HTMLSelectElement) Onmspointerdown() (onmspointerdown func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerdown")
	return onmspointerdown
}

func (*HTMLSelectElement) SetOnmspointerdown(onmspointerdown func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerdown = $1", onmspointerdown)
}

func (*HTMLSelectElement) Onmspointerenter() (onmspointerenter func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerenter")
	return onmspointerenter
}

func (*HTMLSelectElement) SetOnmspointerenter(onmspointerenter func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerenter = $1", onmspointerenter)
}

func (*HTMLSelectElement) Onmspointerleave() (onmspointerleave func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerleave")
	return onmspointerleave
}

func (*HTMLSelectElement) SetOnmspointerleave(onmspointerleave func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerleave = $1", onmspointerleave)
}

func (*HTMLSelectElement) Onmspointermove() (onmspointermove func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointermove")
	return onmspointermove
}

func (*HTMLSelectElement) SetOnmspointermove(onmspointermove func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointermove = $1", onmspointermove)
}

func (*HTMLSelectElement) Onmspointerout() (onmspointerout func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerout")
	return onmspointerout
}

func (*HTMLSelectElement) SetOnmspointerout(onmspointerout func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerout = $1", onmspointerout)
}

func (*HTMLSelectElement) Onmspointerover() (onmspointerover func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerover")
	return onmspointerover
}

func (*HTMLSelectElement) SetOnmspointerover(onmspointerover func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerover = $1", onmspointerover)
}

func (*HTMLSelectElement) Onmspointerup() (onmspointerup func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerup")
	return onmspointerup
}

func (*HTMLSelectElement) SetOnmspointerup(onmspointerup func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerup = $1", onmspointerup)
}

func (*HTMLSelectElement) Ontouchcancel() (ontouchcancel func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchcancel")
	return ontouchcancel
}

func (*HTMLSelectElement) SetOntouchcancel(ontouchcancel func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchcancel = $1", ontouchcancel)
}

func (*HTMLSelectElement) Ontouchend() (ontouchend func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchend")
	return ontouchend
}

func (*HTMLSelectElement) SetOntouchend(ontouchend func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchend = $1", ontouchend)
}

func (*HTMLSelectElement) Ontouchmove() (ontouchmove func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchmove")
	return ontouchmove
}

func (*HTMLSelectElement) SetOntouchmove(ontouchmove func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchmove = $1", ontouchmove)
}

func (*HTMLSelectElement) Ontouchstart() (ontouchstart func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchstart")
	return ontouchstart
}

func (*HTMLSelectElement) SetOntouchstart(ontouchstart func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchstart = $1", ontouchstart)
}

func (*HTMLSelectElement) Onwebkitfullscreenchange() (onwebkitfullscreenchange func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenchange")
	return onwebkitfullscreenchange
}

func (*HTMLSelectElement) SetOnwebkitfullscreenchange(onwebkitfullscreenchange func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenchange = $1", onwebkitfullscreenchange)
}

func (*HTMLSelectElement) Onwebkitfullscreenerror() (onwebkitfullscreenerror func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenerror")
	return onwebkitfullscreenerror
}

func (*HTMLSelectElement) SetOnwebkitfullscreenerror(onwebkitfullscreenerror func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenerror = $1", onwebkitfullscreenerror)
}

func (*HTMLSelectElement) OuterHTML() (outerHTML string) {
	macro.Rewrite("$_.outerHTML")
	return outerHTML
}

func (*HTMLSelectElement) SetOuterHTML(outerHTML string) {
	macro.Rewrite("$_.outerHTML = $1", outerHTML)
}

func (*HTMLSelectElement) Prefix() (prefix string) {
	macro.Rewrite("$_.prefix")
	return prefix
}

func (*HTMLSelectElement) ScrollHeight() (scrollHeight int) {
	macro.Rewrite("$_.scrollHeight")
	return scrollHeight
}

func (*HTMLSelectElement) ScrollLeft() (scrollLeft int) {
	macro.Rewrite("$_.scrollLeft")
	return scrollLeft
}

func (*HTMLSelectElement) SetScrollLeft(scrollLeft int) {
	macro.Rewrite("$_.scrollLeft = $1", scrollLeft)
}

func (*HTMLSelectElement) ScrollTop() (scrollTop int) {
	macro.Rewrite("$_.scrollTop")
	return scrollTop
}

func (*HTMLSelectElement) SetScrollTop(scrollTop int) {
	macro.Rewrite("$_.scrollTop = $1", scrollTop)
}

func (*HTMLSelectElement) ScrollWidth() (scrollWidth int) {
	macro.Rewrite("$_.scrollWidth")
	return scrollWidth
}

func (*HTMLSelectElement) TagName() (tagName string) {
	macro.Rewrite("$_.tagName")
	return tagName
}

func (*HTMLSelectElement) Onpointercancel() (onpointercancel func(event.Event)) {
	macro.Rewrite("$_.onpointercancel")
	return onpointercancel
}

func (*HTMLSelectElement) SetOnpointercancel(onpointercancel func(event.Event)) {
	macro.Rewrite("$_.onpointercancel = $1", onpointercancel)
}

func (*HTMLSelectElement) Onpointerdown() (onpointerdown func(event.Event)) {
	macro.Rewrite("$_.onpointerdown")
	return onpointerdown
}

func (*HTMLSelectElement) SetOnpointerdown(onpointerdown func(event.Event)) {
	macro.Rewrite("$_.onpointerdown = $1", onpointerdown)
}

func (*HTMLSelectElement) Onpointerenter() (onpointerenter func(event.Event)) {
	macro.Rewrite("$_.onpointerenter")
	return onpointerenter
}

func (*HTMLSelectElement) SetOnpointerenter(onpointerenter func(event.Event)) {
	macro.Rewrite("$_.onpointerenter = $1", onpointerenter)
}

func (*HTMLSelectElement) Onpointerleave() (onpointerleave func(event.Event)) {
	macro.Rewrite("$_.onpointerleave")
	return onpointerleave
}

func (*HTMLSelectElement) SetOnpointerleave(onpointerleave func(event.Event)) {
	macro.Rewrite("$_.onpointerleave = $1", onpointerleave)
}

func (*HTMLSelectElement) Onpointermove() (onpointermove func(event.Event)) {
	macro.Rewrite("$_.onpointermove")
	return onpointermove
}

func (*HTMLSelectElement) SetOnpointermove(onpointermove func(event.Event)) {
	macro.Rewrite("$_.onpointermove = $1", onpointermove)
}

func (*HTMLSelectElement) Onpointerout() (onpointerout func(event.Event)) {
	macro.Rewrite("$_.onpointerout")
	return onpointerout
}

func (*HTMLSelectElement) SetOnpointerout(onpointerout func(event.Event)) {
	macro.Rewrite("$_.onpointerout = $1", onpointerout)
}

func (*HTMLSelectElement) Onpointerover() (onpointerover func(event.Event)) {
	macro.Rewrite("$_.onpointerover")
	return onpointerover
}

func (*HTMLSelectElement) SetOnpointerover(onpointerover func(event.Event)) {
	macro.Rewrite("$_.onpointerover = $1", onpointerover)
}

func (*HTMLSelectElement) Onpointerup() (onpointerup func(event.Event)) {
	macro.Rewrite("$_.onpointerup")
	return onpointerup
}

func (*HTMLSelectElement) SetOnpointerup(onpointerup func(event.Event)) {
	macro.Rewrite("$_.onpointerup = $1", onpointerup)
}

func (*HTMLSelectElement) Onwheel() (onwheel func(event.Event)) {
	macro.Rewrite("$_.onwheel")
	return onwheel
}

func (*HTMLSelectElement) SetOnwheel(onwheel func(event.Event)) {
	macro.Rewrite("$_.onwheel = $1", onwheel)
}

func (*HTMLSelectElement) ChildElementCount() (childElementCount uint) {
	macro.Rewrite("$_.childElementCount")
	return childElementCount
}

func (*HTMLSelectElement) FirstElementChild() (firstElementChild element.Element) {
	macro.Rewrite("$_.firstElementChild")
	return firstElementChild
}

func (*HTMLSelectElement) LastElementChild() (lastElementChild element.Element) {
	macro.Rewrite("$_.lastElementChild")
	return lastElementChild
}

func (*HTMLSelectElement) NextElementSibling() (nextElementSibling element.Element) {
	macro.Rewrite("$_.nextElementSibling")
	return nextElementSibling
}

func (*HTMLSelectElement) PreviousElementSibling() (previousElementSibling element.Element) {
	macro.Rewrite("$_.previousElementSibling")
	return previousElementSibling
}

func (*HTMLSelectElement) Attributes() (attributes *dom.NamedNodeMap) {
	macro.Rewrite("$_.attributes")
	return attributes
}

func (*HTMLSelectElement) BaseURI() (baseURI string) {
	macro.Rewrite("$_.baseURI")
	return baseURI
}

func (*HTMLSelectElement) ChildNodes() (childNodes *dom.NodeList) {
	macro.Rewrite("$_.childNodes")
	return childNodes
}

func (*HTMLSelectElement) FirstChild() (firstChild dom.Node) {
	macro.Rewrite("$_.firstChild")
	return firstChild
}

func (*HTMLSelectElement) LastChild() (lastChild dom.Node) {
	macro.Rewrite("$_.lastChild")
	return lastChild
}

func (*HTMLSelectElement) LocalName() (localName string) {
	macro.Rewrite("$_.localName")
	return localName
}

func (*HTMLSelectElement) NamespaceURI() (namespaceURI string) {
	macro.Rewrite("$_.namespaceURI")
	return namespaceURI
}

func (*HTMLSelectElement) NextSibling() (nextSibling dom.Node) {
	macro.Rewrite("$_.nextSibling")
	return nextSibling
}

func (*HTMLSelectElement) NodeName() (nodeName string) {
	macro.Rewrite("$_.nodeName")
	return nodeName
}

func (*HTMLSelectElement) NodeType() (nodeType uint8) {
	macro.Rewrite("$_.nodeType")
	return nodeType
}

func (*HTMLSelectElement) NodeValue() (nodeValue string) {
	macro.Rewrite("$_.nodeValue")
	return nodeValue
}

func (*HTMLSelectElement) SetNodeValue(nodeValue string) {
	macro.Rewrite("$_.nodeValue = $1", nodeValue)
}

func (*HTMLSelectElement) OwnerDocument() (ownerDocument document.Document) {
	macro.Rewrite("$_.ownerDocument")
	return ownerDocument
}

func (*HTMLSelectElement) ParentElement() (parentElement HTMLElement) {
	macro.Rewrite("$_.parentElement")
	return parentElement
}

func (*HTMLSelectElement) ParentNode() (parentNode dom.Node) {
	macro.Rewrite("$_.parentNode")
	return parentNode
}

func (*HTMLSelectElement) PreviousSibling() (previousSibling dom.Node) {
	macro.Rewrite("$_.previousSibling")
	return previousSibling
}

func (*HTMLSelectElement) TextContent() (textContent string) {
	macro.Rewrite("$_.textContent")
	return textContent
}

func (*HTMLSelectElement) SetTextContent(textContent string) {
	macro.Rewrite("$_.textContent = $1", textContent)
}
