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

var _ HTMLElement = (*HTMLButtonElement)(nil)
var _ element.Element = (*HTMLButtonElement)(nil)
var _ utils.GlobalEventHandlers = (*HTMLButtonElement)(nil)
var _ element.ElementTraversal = (*HTMLButtonElement)(nil)
var _ dom.NodeSelector = (*HTMLButtonElement)(nil)
var _ dom.ChildNode = (*HTMLButtonElement)(nil)
var _ dom.Node = (*HTMLButtonElement)(nil)
var _ event.EventTarget = (*HTMLButtonElement)(nil)

type HTMLButtonElement struct {
}

func (*HTMLButtonElement) CheckValidity() (b bool) {
	macro.Rewrite("$_.checkValidity()")
	return b
}

func (*HTMLButtonElement) SetCustomValidity(err string) {
	macro.Rewrite("$_.setCustomValidity($1)", err)
}

func (*HTMLButtonElement) Blur() {
	macro.Rewrite("$_.blur()")
}

func (*HTMLButtonElement) Click() {
	macro.Rewrite("$_.click()")
}

func (*HTMLButtonElement) DragDrop() (b bool) {
	macro.Rewrite("$_.dragDrop()")
	return b
}

func (*HTMLButtonElement) Focus() {
	macro.Rewrite("$_.focus()")
}

func (*HTMLButtonElement) GetElementsByClassName(classNames string) (w *dom.NodeList) {
	macro.Rewrite("$_.getElementsByClassName($1)", classNames)
	return w
}

func (*HTMLButtonElement) InsertAdjacentElement(position string, insertedElement element.Element) (w element.Element) {
	macro.Rewrite("$_.insertAdjacentElement($1, $2)", position, insertedElement)
	return w
}

func (*HTMLButtonElement) InsertAdjacentHTML(where string, html string) {
	macro.Rewrite("$_.insertAdjacentHTML($1, $2)", where, html)
}

func (*HTMLButtonElement) InsertAdjacentText(where string, text string) {
	macro.Rewrite("$_.insertAdjacentText($1, $2)", where, text)
}

func (*HTMLButtonElement) MsGetInputContext() (w *ms.MSInputMethodContext) {
	macro.Rewrite("$_.msGetInputContext()")
	return w
}

func (*HTMLButtonElement) ScrollIntoView(top *bool) {
	macro.Rewrite("$_.scrollIntoView($1)", top)
}

func (*HTMLButtonElement) GetAttribute(qualifiedName string) (s string) {
	macro.Rewrite("$_.getAttribute($1)", qualifiedName)
	return s
}

func (*HTMLButtonElement) GetAttributeNode(name string) (w *dom.Attr) {
	macro.Rewrite("$_.getAttributeNode($1)", name)
	return w
}

func (*HTMLButtonElement) GetAttributeNodeNS(namespaceURI string, localName string) (w *dom.Attr) {
	macro.Rewrite("$_.getAttributeNodeNS($1, $2)", namespaceURI, localName)
	return w
}

func (*HTMLButtonElement) GetAttributeNS(namespaceURI string, localName string) (s string) {
	macro.Rewrite("$_.getAttributeNS($1, $2)", namespaceURI, localName)
	return s
}

func (*HTMLButtonElement) GetBoundingClientRect() (c *dom.ClientRect) {
	macro.Rewrite("$_.getBoundingClientRect()")
	return c
}

func (*HTMLButtonElement) GetClientRects() (c *dom.ClientRectList) {
	macro.Rewrite("$_.getClientRects()")
	return c
}

func (*HTMLButtonElement) GetElementsByTagName(name string) (w *dom.NodeList) {
	macro.Rewrite("$_.getElementsByTagName($1)", name)
	return w
}

func (*HTMLButtonElement) GetElementsByTagNameNS(namespaceURI string, localName string) (w *dom.NodeList) {
	macro.Rewrite("$_.getElementsByTagNameNS($1, $2)", namespaceURI, localName)
	return w
}

func (*HTMLButtonElement) HasAttribute(name string) (b bool) {
	macro.Rewrite("$_.hasAttribute($1)", name)
	return b
}

func (*HTMLButtonElement) HasAttributeNS(namespaceURI string, localName string) (b bool) {
	macro.Rewrite("$_.hasAttributeNS($1, $2)", namespaceURI, localName)
	return b
}

func (*HTMLButtonElement) MsGetRegionContent() (w *ms.MSRangeCollection) {
	macro.Rewrite("$_.msGetRegionContent()")
	return w
}

func (*HTMLButtonElement) MsGetUntransformedBounds() (c *dom.ClientRect) {
	macro.Rewrite("$_.msGetUntransformedBounds()")
	return c
}

func (*HTMLButtonElement) MsMatchesSelector(selectors string) (b bool) {
	macro.Rewrite("$_.msMatchesSelector($1)", selectors)
	return b
}

func (*HTMLButtonElement) MsReleasePointerCapture(pointerId int) {
	macro.Rewrite("$_.msReleasePointerCapture($1)", pointerId)
}

func (*HTMLButtonElement) MsSetPointerCapture(pointerId int) {
	macro.Rewrite("$_.msSetPointerCapture($1)", pointerId)
}

func (*HTMLButtonElement) MsZoomTo(args *ms.MsZoomToOptions) {
	macro.Rewrite("$_.msZoomTo($1)", args)
}

func (*HTMLButtonElement) ReleasePointerCapture(pointerId int) {
	macro.Rewrite("$_.releasePointerCapture($1)", pointerId)
}

func (*HTMLButtonElement) RemoveAttribute(qualifiedName string) {
	macro.Rewrite("$_.removeAttribute($1)", qualifiedName)
}

func (*HTMLButtonElement) RemoveAttributeNode(oldAttr *dom.Attr) (w *dom.Attr) {
	macro.Rewrite("$_.removeAttributeNode($1)", oldAttr)
	return w
}

func (*HTMLButtonElement) RemoveAttributeNS(namespaceURI string, localName string) {
	macro.Rewrite("$_.removeAttributeNS($1, $2)", namespaceURI, localName)
}

func (*HTMLButtonElement) RequestFullscreen() {
	macro.Rewrite("$_.requestFullscreen()")
}

func (*HTMLButtonElement) RequestPointerLock() {
	macro.Rewrite("$_.requestPointerLock()")
}

func (*HTMLButtonElement) SetAttribute(qualifiedName string, value string) {
	macro.Rewrite("$_.setAttribute($1, $2)", qualifiedName, value)
}

func (*HTMLButtonElement) SetAttributeNode(newAttr *dom.Attr) (w *dom.Attr) {
	macro.Rewrite("$_.setAttributeNode($1)", newAttr)
	return w
}

func (*HTMLButtonElement) SetAttributeNodeNS(newAttr *dom.Attr) (w *dom.Attr) {
	macro.Rewrite("$_.setAttributeNodeNS($1)", newAttr)
	return w
}

func (*HTMLButtonElement) SetAttributeNS(namespaceURI string, qualifiedName string, value string) {
	macro.Rewrite("$_.setAttributeNS($1, $2, $3)", namespaceURI, qualifiedName, value)
}

func (*HTMLButtonElement) SetPointerCapture(pointerId int) {
	macro.Rewrite("$_.setPointerCapture($1)", pointerId)
}

func (*HTMLButtonElement) WebkitMatchesSelector(selectors string) (b bool) {
	macro.Rewrite("$_.webkitMatchesSelector($1)", selectors)
	return b
}

func (*HTMLButtonElement) WebkitRequestFullscreen() {
	macro.Rewrite("$_.webkitRequestFullscreen()")
}

func (*HTMLButtonElement) WebkitRequestFullScreen() {
	macro.Rewrite("$_.webkitRequestFullScreen()")
}

func (*HTMLButtonElement) QuerySelector(selectors string) (w element.Element) {
	macro.Rewrite("$_.querySelector($1)", selectors)
	return w
}

func (*HTMLButtonElement) QuerySelectorAll(selectors string) (w *dom.NodeList) {
	macro.Rewrite("$_.querySelectorAll($1)", selectors)
	return w
}

func (*HTMLButtonElement) AppendChild(newChild dom.Node) (w dom.Node) {
	macro.Rewrite("$_.appendChild($1)", newChild)
	return w
}

func (*HTMLButtonElement) CloneNode(deep *bool) (w dom.Node) {
	macro.Rewrite("$_.cloneNode($1)", deep)
	return w
}

func (*HTMLButtonElement) CompareDocumentPosition(other dom.Node) (u uint8) {
	macro.Rewrite("$_.compareDocumentPosition($1)", other)
	return u
}

func (*HTMLButtonElement) Contains(child dom.Node) (b bool) {
	macro.Rewrite("$_.contains($1)", child)
	return b
}

func (*HTMLButtonElement) HasAttributes() (b bool) {
	macro.Rewrite("$_.hasAttributes()")
	return b
}

func (*HTMLButtonElement) HasChildNodes() (b bool) {
	macro.Rewrite("$_.hasChildNodes()")
	return b
}

func (*HTMLButtonElement) InsertBefore(newChild dom.Node, refChild *dom.Node) (w dom.Node) {
	macro.Rewrite("$_.insertBefore($1, $2)", newChild, refChild)
	return w
}

func (*HTMLButtonElement) IsDefaultNamespace(namespaceURI string) (b bool) {
	macro.Rewrite("$_.isDefaultNamespace($1)", namespaceURI)
	return b
}

func (*HTMLButtonElement) IsEqualNode(arg dom.Node) (b bool) {
	macro.Rewrite("$_.isEqualNode($1)", arg)
	return b
}

func (*HTMLButtonElement) IsSameNode(other dom.Node) (b bool) {
	macro.Rewrite("$_.isSameNode($1)", other)
	return b
}

func (*HTMLButtonElement) LookupNamespaceURI(prefix string) (s string) {
	macro.Rewrite("$_.lookupNamespaceURI($1)", prefix)
	return s
}

func (*HTMLButtonElement) LookupPrefix(namespaceURI string) (s string) {
	macro.Rewrite("$_.lookupPrefix($1)", namespaceURI)
	return s
}

func (*HTMLButtonElement) Normalize() {
	macro.Rewrite("$_.normalize()")
}

func (*HTMLButtonElement) RemoveChild(oldChild dom.Node) (w dom.Node) {
	macro.Rewrite("$_.removeChild($1)", oldChild)
	return w
}

func (*HTMLButtonElement) ReplaceChild(newChild dom.Node, oldChild dom.Node) (w dom.Node) {
	macro.Rewrite("$_.replaceChild($1, $2)", newChild, oldChild)
	return w
}

func (*HTMLButtonElement) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*HTMLButtonElement) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*HTMLButtonElement) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*HTMLButtonElement) Autofocus() (autofocus bool) {
	macro.Rewrite("$_.autofocus")
	return autofocus
}

func (*HTMLButtonElement) SetAutofocus(autofocus bool) {
	macro.Rewrite("$_.autofocus = $1", autofocus)
}

func (*HTMLButtonElement) Disabled() (disabled bool) {
	macro.Rewrite("$_.disabled")
	return disabled
}

func (*HTMLButtonElement) SetDisabled(disabled bool) {
	macro.Rewrite("$_.disabled = $1", disabled)
}

func (*HTMLButtonElement) Form() (form *HTMLFormElement) {
	macro.Rewrite("$_.form")
	return form
}

func (*HTMLButtonElement) FormAction() (formAction string) {
	macro.Rewrite("$_.formAction")
	return formAction
}

func (*HTMLButtonElement) SetFormAction(formAction string) {
	macro.Rewrite("$_.formAction = $1", formAction)
}

func (*HTMLButtonElement) FormEnctype() (formEnctype string) {
	macro.Rewrite("$_.formEnctype")
	return formEnctype
}

func (*HTMLButtonElement) SetFormEnctype(formEnctype string) {
	macro.Rewrite("$_.formEnctype = $1", formEnctype)
}

func (*HTMLButtonElement) FormMethod() (formMethod string) {
	macro.Rewrite("$_.formMethod")
	return formMethod
}

func (*HTMLButtonElement) SetFormMethod(formMethod string) {
	macro.Rewrite("$_.formMethod = $1", formMethod)
}

func (*HTMLButtonElement) FormNoValidate() (formNoValidate string) {
	macro.Rewrite("$_.formNoValidate")
	return formNoValidate
}

func (*HTMLButtonElement) SetFormNoValidate(formNoValidate string) {
	macro.Rewrite("$_.formNoValidate = $1", formNoValidate)
}

func (*HTMLButtonElement) FormTarget() (formTarget string) {
	macro.Rewrite("$_.formTarget")
	return formTarget
}

func (*HTMLButtonElement) SetFormTarget(formTarget string) {
	macro.Rewrite("$_.formTarget = $1", formTarget)
}

func (*HTMLButtonElement) Name() (name string) {
	macro.Rewrite("$_.name")
	return name
}

func (*HTMLButtonElement) SetName(name string) {
	macro.Rewrite("$_.name = $1", name)
}

func (*HTMLButtonElement) Status() (status interface{}) {
	macro.Rewrite("$_.status")
	return status
}

func (*HTMLButtonElement) SetStatus(status interface{}) {
	macro.Rewrite("$_.status = $1", status)
}

func (*HTMLButtonElement) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

func (*HTMLButtonElement) SetType(kind string) {
	macro.Rewrite("$_.type = $1", kind)
}

func (*HTMLButtonElement) ValidationMessage() (validationMessage string) {
	macro.Rewrite("$_.validationMessage")
	return validationMessage
}

func (*HTMLButtonElement) Validity() (validity *utils.ValidityState) {
	macro.Rewrite("$_.validity")
	return validity
}

func (*HTMLButtonElement) Value() (value string) {
	macro.Rewrite("$_.value")
	return value
}

func (*HTMLButtonElement) SetValue(value string) {
	macro.Rewrite("$_.value = $1", value)
}

func (*HTMLButtonElement) WillValidate() (willValidate bool) {
	macro.Rewrite("$_.willValidate")
	return willValidate
}

func (*HTMLButtonElement) AccessKey() (accessKey string) {
	macro.Rewrite("$_.accessKey")
	return accessKey
}

func (*HTMLButtonElement) SetAccessKey(accessKey string) {
	macro.Rewrite("$_.accessKey = $1", accessKey)
}

func (*HTMLButtonElement) Children() (children HTMLCollection) {
	macro.Rewrite("$_.children")
	return children
}

func (*HTMLButtonElement) ContentEditable() (contentEditable string) {
	macro.Rewrite("$_.contentEditable")
	return contentEditable
}

func (*HTMLButtonElement) SetContentEditable(contentEditable string) {
	macro.Rewrite("$_.contentEditable = $1", contentEditable)
}

func (*HTMLButtonElement) Dataset() (dataset *dom.DOMStringMap) {
	macro.Rewrite("$_.dataset")
	return dataset
}

func (*HTMLButtonElement) Dir() (dir string) {
	macro.Rewrite("$_.dir")
	return dir
}

func (*HTMLButtonElement) SetDir(dir string) {
	macro.Rewrite("$_.dir = $1", dir)
}

func (*HTMLButtonElement) Draggable() (draggable bool) {
	macro.Rewrite("$_.draggable")
	return draggable
}

func (*HTMLButtonElement) SetDraggable(draggable bool) {
	macro.Rewrite("$_.draggable = $1", draggable)
}

func (*HTMLButtonElement) Hidden() (hidden bool) {
	macro.Rewrite("$_.hidden")
	return hidden
}

func (*HTMLButtonElement) SetHidden(hidden bool) {
	macro.Rewrite("$_.hidden = $1", hidden)
}

func (*HTMLButtonElement) HideFocus() (hideFocus bool) {
	macro.Rewrite("$_.hideFocus")
	return hideFocus
}

func (*HTMLButtonElement) SetHideFocus(hideFocus bool) {
	macro.Rewrite("$_.hideFocus = $1", hideFocus)
}

func (*HTMLButtonElement) InnerText() (innerText string) {
	macro.Rewrite("$_.innerText")
	return innerText
}

func (*HTMLButtonElement) SetInnerText(innerText string) {
	macro.Rewrite("$_.innerText = $1", innerText)
}

func (*HTMLButtonElement) IsContentEditable() (isContentEditable bool) {
	macro.Rewrite("$_.isContentEditable")
	return isContentEditable
}

func (*HTMLButtonElement) Lang() (lang string) {
	macro.Rewrite("$_.lang")
	return lang
}

func (*HTMLButtonElement) SetLang(lang string) {
	macro.Rewrite("$_.lang = $1", lang)
}

func (*HTMLButtonElement) OffsetHeight() (offsetHeight int) {
	macro.Rewrite("$_.offsetHeight")
	return offsetHeight
}

func (*HTMLButtonElement) OffsetLeft() (offsetLeft int) {
	macro.Rewrite("$_.offsetLeft")
	return offsetLeft
}

func (*HTMLButtonElement) OffsetParent() (offsetParent element.Element) {
	macro.Rewrite("$_.offsetParent")
	return offsetParent
}

func (*HTMLButtonElement) OffsetTop() (offsetTop int) {
	macro.Rewrite("$_.offsetTop")
	return offsetTop
}

func (*HTMLButtonElement) OffsetWidth() (offsetWidth int) {
	macro.Rewrite("$_.offsetWidth")
	return offsetWidth
}

func (*HTMLButtonElement) Onabort() (onabort func(event.Event)) {
	macro.Rewrite("$_.onabort")
	return onabort
}

func (*HTMLButtonElement) SetOnabort(onabort func(event.Event)) {
	macro.Rewrite("$_.onabort = $1", onabort)
}

func (*HTMLButtonElement) Onactivate() (onactivate func(ui.UIEvent)) {
	macro.Rewrite("$_.onactivate")
	return onactivate
}

func (*HTMLButtonElement) SetOnactivate(onactivate func(ui.UIEvent)) {
	macro.Rewrite("$_.onactivate = $1", onactivate)
}

func (*HTMLButtonElement) Onbeforeactivate() (onbeforeactivate func(ui.UIEvent)) {
	macro.Rewrite("$_.onbeforeactivate")
	return onbeforeactivate
}

func (*HTMLButtonElement) SetOnbeforeactivate(onbeforeactivate func(ui.UIEvent)) {
	macro.Rewrite("$_.onbeforeactivate = $1", onbeforeactivate)
}

func (*HTMLButtonElement) Onbeforecopy() (onbeforecopy func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.onbeforecopy")
	return onbeforecopy
}

func (*HTMLButtonElement) SetOnbeforecopy(onbeforecopy func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.onbeforecopy = $1", onbeforecopy)
}

func (*HTMLButtonElement) Onbeforecut() (onbeforecut func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.onbeforecut")
	return onbeforecut
}

func (*HTMLButtonElement) SetOnbeforecut(onbeforecut func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.onbeforecut = $1", onbeforecut)
}

func (*HTMLButtonElement) Onbeforedeactivate() (onbeforedeactivate func(ui.UIEvent)) {
	macro.Rewrite("$_.onbeforedeactivate")
	return onbeforedeactivate
}

func (*HTMLButtonElement) SetOnbeforedeactivate(onbeforedeactivate func(ui.UIEvent)) {
	macro.Rewrite("$_.onbeforedeactivate = $1", onbeforedeactivate)
}

func (*HTMLButtonElement) Onbeforepaste() (onbeforepaste func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.onbeforepaste")
	return onbeforepaste
}

func (*HTMLButtonElement) SetOnbeforepaste(onbeforepaste func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.onbeforepaste = $1", onbeforepaste)
}

func (*HTMLButtonElement) Onblur() (onblur func(*utils.FocusEvent)) {
	macro.Rewrite("$_.onblur")
	return onblur
}

func (*HTMLButtonElement) SetOnblur(onblur func(*utils.FocusEvent)) {
	macro.Rewrite("$_.onblur = $1", onblur)
}

func (*HTMLButtonElement) Oncanplay() (oncanplay func(event.Event)) {
	macro.Rewrite("$_.oncanplay")
	return oncanplay
}

func (*HTMLButtonElement) SetOncanplay(oncanplay func(event.Event)) {
	macro.Rewrite("$_.oncanplay = $1", oncanplay)
}

func (*HTMLButtonElement) Oncanplaythrough() (oncanplaythrough func(event.Event)) {
	macro.Rewrite("$_.oncanplaythrough")
	return oncanplaythrough
}

func (*HTMLButtonElement) SetOncanplaythrough(oncanplaythrough func(event.Event)) {
	macro.Rewrite("$_.oncanplaythrough = $1", oncanplaythrough)
}

func (*HTMLButtonElement) Onchange() (onchange func(event.Event)) {
	macro.Rewrite("$_.onchange")
	return onchange
}

func (*HTMLButtonElement) SetOnchange(onchange func(event.Event)) {
	macro.Rewrite("$_.onchange = $1", onchange)
}

func (*HTMLButtonElement) Onclick() (onclick func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onclick")
	return onclick
}

func (*HTMLButtonElement) SetOnclick(onclick func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onclick = $1", onclick)
}

func (*HTMLButtonElement) Oncontextmenu() (oncontextmenu func(*utils.PointerEvent)) {
	macro.Rewrite("$_.oncontextmenu")
	return oncontextmenu
}

func (*HTMLButtonElement) SetOncontextmenu(oncontextmenu func(*utils.PointerEvent)) {
	macro.Rewrite("$_.oncontextmenu = $1", oncontextmenu)
}

func (*HTMLButtonElement) Oncopy() (oncopy func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.oncopy")
	return oncopy
}

func (*HTMLButtonElement) SetOncopy(oncopy func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.oncopy = $1", oncopy)
}

func (*HTMLButtonElement) Oncuechange() (oncuechange func(event.Event)) {
	macro.Rewrite("$_.oncuechange")
	return oncuechange
}

func (*HTMLButtonElement) SetOncuechange(oncuechange func(event.Event)) {
	macro.Rewrite("$_.oncuechange = $1", oncuechange)
}

func (*HTMLButtonElement) Oncut() (oncut func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.oncut")
	return oncut
}

func (*HTMLButtonElement) SetOncut(oncut func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.oncut = $1", oncut)
}

func (*HTMLButtonElement) Ondblclick() (ondblclick func(mouse.MouseEvent)) {
	macro.Rewrite("$_.ondblclick")
	return ondblclick
}

func (*HTMLButtonElement) SetOndblclick(ondblclick func(mouse.MouseEvent)) {
	macro.Rewrite("$_.ondblclick = $1", ondblclick)
}

func (*HTMLButtonElement) Ondeactivate() (ondeactivate func(ui.UIEvent)) {
	macro.Rewrite("$_.ondeactivate")
	return ondeactivate
}

func (*HTMLButtonElement) SetOndeactivate(ondeactivate func(ui.UIEvent)) {
	macro.Rewrite("$_.ondeactivate = $1", ondeactivate)
}

func (*HTMLButtonElement) Ondrag() (ondrag func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondrag")
	return ondrag
}

func (*HTMLButtonElement) SetOndrag(ondrag func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondrag = $1", ondrag)
}

func (*HTMLButtonElement) Ondragend() (ondragend func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragend")
	return ondragend
}

func (*HTMLButtonElement) SetOndragend(ondragend func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragend = $1", ondragend)
}

func (*HTMLButtonElement) Ondragenter() (ondragenter func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragenter")
	return ondragenter
}

func (*HTMLButtonElement) SetOndragenter(ondragenter func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragenter = $1", ondragenter)
}

func (*HTMLButtonElement) Ondragleave() (ondragleave func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragleave")
	return ondragleave
}

func (*HTMLButtonElement) SetOndragleave(ondragleave func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragleave = $1", ondragleave)
}

func (*HTMLButtonElement) Ondragover() (ondragover func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragover")
	return ondragover
}

func (*HTMLButtonElement) SetOndragover(ondragover func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragover = $1", ondragover)
}

func (*HTMLButtonElement) Ondragstart() (ondragstart func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragstart")
	return ondragstart
}

func (*HTMLButtonElement) SetOndragstart(ondragstart func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondragstart = $1", ondragstart)
}

func (*HTMLButtonElement) Ondrop() (ondrop func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondrop")
	return ondrop
}

func (*HTMLButtonElement) SetOndrop(ondrop func(*draganddrop.DragEvent)) {
	macro.Rewrite("$_.ondrop = $1", ondrop)
}

func (*HTMLButtonElement) Ondurationchange() (ondurationchange func(event.Event)) {
	macro.Rewrite("$_.ondurationchange")
	return ondurationchange
}

func (*HTMLButtonElement) SetOndurationchange(ondurationchange func(event.Event)) {
	macro.Rewrite("$_.ondurationchange = $1", ondurationchange)
}

func (*HTMLButtonElement) Onemptied() (onemptied func(event.Event)) {
	macro.Rewrite("$_.onemptied")
	return onemptied
}

func (*HTMLButtonElement) SetOnemptied(onemptied func(event.Event)) {
	macro.Rewrite("$_.onemptied = $1", onemptied)
}

func (*HTMLButtonElement) Onended() (onended func(event.Event)) {
	macro.Rewrite("$_.onended")
	return onended
}

func (*HTMLButtonElement) SetOnended(onended func(event.Event)) {
	macro.Rewrite("$_.onended = $1", onended)
}

func (*HTMLButtonElement) Onerror() (onerror func(event.Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

func (*HTMLButtonElement) SetOnerror(onerror func(event.Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

func (*HTMLButtonElement) Onfocus() (onfocus func(*utils.FocusEvent)) {
	macro.Rewrite("$_.onfocus")
	return onfocus
}

func (*HTMLButtonElement) SetOnfocus(onfocus func(*utils.FocusEvent)) {
	macro.Rewrite("$_.onfocus = $1", onfocus)
}

func (*HTMLButtonElement) Oninput() (oninput func(event.Event)) {
	macro.Rewrite("$_.oninput")
	return oninput
}

func (*HTMLButtonElement) SetOninput(oninput func(event.Event)) {
	macro.Rewrite("$_.oninput = $1", oninput)
}

func (*HTMLButtonElement) Oninvalid() (oninvalid func(event.Event)) {
	macro.Rewrite("$_.oninvalid")
	return oninvalid
}

func (*HTMLButtonElement) SetOninvalid(oninvalid func(event.Event)) {
	macro.Rewrite("$_.oninvalid = $1", oninvalid)
}

func (*HTMLButtonElement) Onkeydown() (onkeydown func(*utils.KeyboardEvent)) {
	macro.Rewrite("$_.onkeydown")
	return onkeydown
}

func (*HTMLButtonElement) SetOnkeydown(onkeydown func(*utils.KeyboardEvent)) {
	macro.Rewrite("$_.onkeydown = $1", onkeydown)
}

func (*HTMLButtonElement) Onkeypress() (onkeypress func(*utils.KeyboardEvent)) {
	macro.Rewrite("$_.onkeypress")
	return onkeypress
}

func (*HTMLButtonElement) SetOnkeypress(onkeypress func(*utils.KeyboardEvent)) {
	macro.Rewrite("$_.onkeypress = $1", onkeypress)
}

func (*HTMLButtonElement) Onkeyup() (onkeyup func(*utils.KeyboardEvent)) {
	macro.Rewrite("$_.onkeyup")
	return onkeyup
}

func (*HTMLButtonElement) SetOnkeyup(onkeyup func(*utils.KeyboardEvent)) {
	macro.Rewrite("$_.onkeyup = $1", onkeyup)
}

func (*HTMLButtonElement) Onload() (onload func(event.Event)) {
	macro.Rewrite("$_.onload")
	return onload
}

func (*HTMLButtonElement) SetOnload(onload func(event.Event)) {
	macro.Rewrite("$_.onload = $1", onload)
}

func (*HTMLButtonElement) Onloadeddata() (onloadeddata func(event.Event)) {
	macro.Rewrite("$_.onloadeddata")
	return onloadeddata
}

func (*HTMLButtonElement) SetOnloadeddata(onloadeddata func(event.Event)) {
	macro.Rewrite("$_.onloadeddata = $1", onloadeddata)
}

func (*HTMLButtonElement) Onloadedmetadata() (onloadedmetadata func(event.Event)) {
	macro.Rewrite("$_.onloadedmetadata")
	return onloadedmetadata
}

func (*HTMLButtonElement) SetOnloadedmetadata(onloadedmetadata func(event.Event)) {
	macro.Rewrite("$_.onloadedmetadata = $1", onloadedmetadata)
}

func (*HTMLButtonElement) Onloadstart() (onloadstart func(event.Event)) {
	macro.Rewrite("$_.onloadstart")
	return onloadstart
}

func (*HTMLButtonElement) SetOnloadstart(onloadstart func(event.Event)) {
	macro.Rewrite("$_.onloadstart = $1", onloadstart)
}

func (*HTMLButtonElement) Onmousedown() (onmousedown func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmousedown")
	return onmousedown
}

func (*HTMLButtonElement) SetOnmousedown(onmousedown func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmousedown = $1", onmousedown)
}

func (*HTMLButtonElement) Onmouseenter() (onmouseenter func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseenter")
	return onmouseenter
}

func (*HTMLButtonElement) SetOnmouseenter(onmouseenter func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseenter = $1", onmouseenter)
}

func (*HTMLButtonElement) Onmouseleave() (onmouseleave func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseleave")
	return onmouseleave
}

func (*HTMLButtonElement) SetOnmouseleave(onmouseleave func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseleave = $1", onmouseleave)
}

func (*HTMLButtonElement) Onmousemove() (onmousemove func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmousemove")
	return onmousemove
}

func (*HTMLButtonElement) SetOnmousemove(onmousemove func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmousemove = $1", onmousemove)
}

func (*HTMLButtonElement) Onmouseout() (onmouseout func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseout")
	return onmouseout
}

func (*HTMLButtonElement) SetOnmouseout(onmouseout func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseout = $1", onmouseout)
}

func (*HTMLButtonElement) Onmouseover() (onmouseover func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseover")
	return onmouseover
}

func (*HTMLButtonElement) SetOnmouseover(onmouseover func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseover = $1", onmouseover)
}

func (*HTMLButtonElement) Onmouseup() (onmouseup func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseup")
	return onmouseup
}

func (*HTMLButtonElement) SetOnmouseup(onmouseup func(mouse.MouseEvent)) {
	macro.Rewrite("$_.onmouseup = $1", onmouseup)
}

func (*HTMLButtonElement) Onmousewheel() (onmousewheel func(*utils.WheelEvent)) {
	macro.Rewrite("$_.onmousewheel")
	return onmousewheel
}

func (*HTMLButtonElement) SetOnmousewheel(onmousewheel func(*utils.WheelEvent)) {
	macro.Rewrite("$_.onmousewheel = $1", onmousewheel)
}

func (*HTMLButtonElement) Onmscontentzoom() (onmscontentzoom func(ui.UIEvent)) {
	macro.Rewrite("$_.onmscontentzoom")
	return onmscontentzoom
}

func (*HTMLButtonElement) SetOnmscontentzoom(onmscontentzoom func(ui.UIEvent)) {
	macro.Rewrite("$_.onmscontentzoom = $1", onmscontentzoom)
}

func (*HTMLButtonElement) Onmsmanipulationstatechanged() (onmsmanipulationstatechanged func(*ms.MSManipulationEvent)) {
	macro.Rewrite("$_.onmsmanipulationstatechanged")
	return onmsmanipulationstatechanged
}

func (*HTMLButtonElement) SetOnmsmanipulationstatechanged(onmsmanipulationstatechanged func(*ms.MSManipulationEvent)) {
	macro.Rewrite("$_.onmsmanipulationstatechanged = $1", onmsmanipulationstatechanged)
}

func (*HTMLButtonElement) Onpaste() (onpaste func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.onpaste")
	return onpaste
}

func (*HTMLButtonElement) SetOnpaste(onpaste func(*utils.ClipboardEvent)) {
	macro.Rewrite("$_.onpaste = $1", onpaste)
}

func (*HTMLButtonElement) Onpause() (onpause func(event.Event)) {
	macro.Rewrite("$_.onpause")
	return onpause
}

func (*HTMLButtonElement) SetOnpause(onpause func(event.Event)) {
	macro.Rewrite("$_.onpause = $1", onpause)
}

func (*HTMLButtonElement) Onplay() (onplay func(event.Event)) {
	macro.Rewrite("$_.onplay")
	return onplay
}

func (*HTMLButtonElement) SetOnplay(onplay func(event.Event)) {
	macro.Rewrite("$_.onplay = $1", onplay)
}

func (*HTMLButtonElement) Onplaying() (onplaying func(event.Event)) {
	macro.Rewrite("$_.onplaying")
	return onplaying
}

func (*HTMLButtonElement) SetOnplaying(onplaying func(event.Event)) {
	macro.Rewrite("$_.onplaying = $1", onplaying)
}

func (*HTMLButtonElement) Onprogress() (onprogress func(event.Event)) {
	macro.Rewrite("$_.onprogress")
	return onprogress
}

func (*HTMLButtonElement) SetOnprogress(onprogress func(event.Event)) {
	macro.Rewrite("$_.onprogress = $1", onprogress)
}

func (*HTMLButtonElement) Onratechange() (onratechange func(event.Event)) {
	macro.Rewrite("$_.onratechange")
	return onratechange
}

func (*HTMLButtonElement) SetOnratechange(onratechange func(event.Event)) {
	macro.Rewrite("$_.onratechange = $1", onratechange)
}

func (*HTMLButtonElement) Onreset() (onreset func(event.Event)) {
	macro.Rewrite("$_.onreset")
	return onreset
}

func (*HTMLButtonElement) SetOnreset(onreset func(event.Event)) {
	macro.Rewrite("$_.onreset = $1", onreset)
}

func (*HTMLButtonElement) Onscroll() (onscroll func(ui.UIEvent)) {
	macro.Rewrite("$_.onscroll")
	return onscroll
}

func (*HTMLButtonElement) SetOnscroll(onscroll func(ui.UIEvent)) {
	macro.Rewrite("$_.onscroll = $1", onscroll)
}

func (*HTMLButtonElement) Onseeked() (onseeked func(event.Event)) {
	macro.Rewrite("$_.onseeked")
	return onseeked
}

func (*HTMLButtonElement) SetOnseeked(onseeked func(event.Event)) {
	macro.Rewrite("$_.onseeked = $1", onseeked)
}

func (*HTMLButtonElement) Onseeking() (onseeking func(event.Event)) {
	macro.Rewrite("$_.onseeking")
	return onseeking
}

func (*HTMLButtonElement) SetOnseeking(onseeking func(event.Event)) {
	macro.Rewrite("$_.onseeking = $1", onseeking)
}

func (*HTMLButtonElement) Onselect() (onselect func(ui.UIEvent)) {
	macro.Rewrite("$_.onselect")
	return onselect
}

func (*HTMLButtonElement) SetOnselect(onselect func(ui.UIEvent)) {
	macro.Rewrite("$_.onselect = $1", onselect)
}

func (*HTMLButtonElement) Onselectstart() (onselectstart func(event.Event)) {
	macro.Rewrite("$_.onselectstart")
	return onselectstart
}

func (*HTMLButtonElement) SetOnselectstart(onselectstart func(event.Event)) {
	macro.Rewrite("$_.onselectstart = $1", onselectstart)
}

func (*HTMLButtonElement) Onstalled() (onstalled func(event.Event)) {
	macro.Rewrite("$_.onstalled")
	return onstalled
}

func (*HTMLButtonElement) SetOnstalled(onstalled func(event.Event)) {
	macro.Rewrite("$_.onstalled = $1", onstalled)
}

func (*HTMLButtonElement) Onsubmit() (onsubmit func(event.Event)) {
	macro.Rewrite("$_.onsubmit")
	return onsubmit
}

func (*HTMLButtonElement) SetOnsubmit(onsubmit func(event.Event)) {
	macro.Rewrite("$_.onsubmit = $1", onsubmit)
}

func (*HTMLButtonElement) Onsuspend() (onsuspend func(event.Event)) {
	macro.Rewrite("$_.onsuspend")
	return onsuspend
}

func (*HTMLButtonElement) SetOnsuspend(onsuspend func(event.Event)) {
	macro.Rewrite("$_.onsuspend = $1", onsuspend)
}

func (*HTMLButtonElement) Ontimeupdate() (ontimeupdate func(event.Event)) {
	macro.Rewrite("$_.ontimeupdate")
	return ontimeupdate
}

func (*HTMLButtonElement) SetOntimeupdate(ontimeupdate func(event.Event)) {
	macro.Rewrite("$_.ontimeupdate = $1", ontimeupdate)
}

func (*HTMLButtonElement) Onvolumechange() (onvolumechange func(event.Event)) {
	macro.Rewrite("$_.onvolumechange")
	return onvolumechange
}

func (*HTMLButtonElement) SetOnvolumechange(onvolumechange func(event.Event)) {
	macro.Rewrite("$_.onvolumechange = $1", onvolumechange)
}

func (*HTMLButtonElement) Onwaiting() (onwaiting func(event.Event)) {
	macro.Rewrite("$_.onwaiting")
	return onwaiting
}

func (*HTMLButtonElement) SetOnwaiting(onwaiting func(event.Event)) {
	macro.Rewrite("$_.onwaiting = $1", onwaiting)
}

func (*HTMLButtonElement) OuterText() (outerText string) {
	macro.Rewrite("$_.outerText")
	return outerText
}

func (*HTMLButtonElement) SetOuterText(outerText string) {
	macro.Rewrite("$_.outerText = $1", outerText)
}

func (*HTMLButtonElement) Spellcheck() (spellcheck bool) {
	macro.Rewrite("$_.spellcheck")
	return spellcheck
}

func (*HTMLButtonElement) SetSpellcheck(spellcheck bool) {
	macro.Rewrite("$_.spellcheck = $1", spellcheck)
}

func (*HTMLButtonElement) Style() (style *css.CSSStyleDeclaration) {
	macro.Rewrite("$_.style")
	return style
}

func (*HTMLButtonElement) TabIndex() (tabIndex int8) {
	macro.Rewrite("$_.tabIndex")
	return tabIndex
}

func (*HTMLButtonElement) SetTabIndex(tabIndex int8) {
	macro.Rewrite("$_.tabIndex = $1", tabIndex)
}

func (*HTMLButtonElement) Title() (title string) {
	macro.Rewrite("$_.title")
	return title
}

func (*HTMLButtonElement) SetTitle(title string) {
	macro.Rewrite("$_.title = $1", title)
}

func (*HTMLButtonElement) ClassList() (classList dom.DOMTokenList) {
	macro.Rewrite("$_.classList")
	return classList
}

func (*HTMLButtonElement) ClassName() (className string) {
	macro.Rewrite("$_.className")
	return className
}

func (*HTMLButtonElement) SetClassName(className string) {
	macro.Rewrite("$_.className = $1", className)
}

func (*HTMLButtonElement) ClientHeight() (clientHeight int) {
	macro.Rewrite("$_.clientHeight")
	return clientHeight
}

func (*HTMLButtonElement) ClientLeft() (clientLeft int) {
	macro.Rewrite("$_.clientLeft")
	return clientLeft
}

func (*HTMLButtonElement) ClientTop() (clientTop int) {
	macro.Rewrite("$_.clientTop")
	return clientTop
}

func (*HTMLButtonElement) ClientWidth() (clientWidth int) {
	macro.Rewrite("$_.clientWidth")
	return clientWidth
}

func (*HTMLButtonElement) ID() (id string) {
	macro.Rewrite("$_.id")
	return id
}

func (*HTMLButtonElement) SetID(id string) {
	macro.Rewrite("$_.id = $1", id)
}

func (*HTMLButtonElement) InnerHTML() (innerHTML string) {
	macro.Rewrite("$_.innerHTML")
	return innerHTML
}

func (*HTMLButtonElement) SetInnerHTML(innerHTML string) {
	macro.Rewrite("$_.innerHTML = $1", innerHTML)
}

func (*HTMLButtonElement) MsContentZoomFactor() (msContentZoomFactor float32) {
	macro.Rewrite("$_.msContentZoomFactor")
	return msContentZoomFactor
}

func (*HTMLButtonElement) SetMsContentZoomFactor(msContentZoomFactor float32) {
	macro.Rewrite("$_.msContentZoomFactor = $1", msContentZoomFactor)
}

func (*HTMLButtonElement) MsRegionOverflow() (msRegionOverflow string) {
	macro.Rewrite("$_.msRegionOverflow")
	return msRegionOverflow
}

func (*HTMLButtonElement) Onariarequest() (onariarequest func(event.Event)) {
	macro.Rewrite("$_.onariarequest")
	return onariarequest
}

func (*HTMLButtonElement) SetOnariarequest(onariarequest func(event.Event)) {
	macro.Rewrite("$_.onariarequest = $1", onariarequest)
}

func (*HTMLButtonElement) Oncommand() (oncommand func(event.Event)) {
	macro.Rewrite("$_.oncommand")
	return oncommand
}

func (*HTMLButtonElement) SetOncommand(oncommand func(event.Event)) {
	macro.Rewrite("$_.oncommand = $1", oncommand)
}

func (*HTMLButtonElement) Ongotpointercapture() (ongotpointercapture func(*utils.PointerEvent)) {
	macro.Rewrite("$_.ongotpointercapture")
	return ongotpointercapture
}

func (*HTMLButtonElement) SetOngotpointercapture(ongotpointercapture func(*utils.PointerEvent)) {
	macro.Rewrite("$_.ongotpointercapture = $1", ongotpointercapture)
}

func (*HTMLButtonElement) Onlostpointercapture() (onlostpointercapture func(*utils.PointerEvent)) {
	macro.Rewrite("$_.onlostpointercapture")
	return onlostpointercapture
}

func (*HTMLButtonElement) SetOnlostpointercapture(onlostpointercapture func(*utils.PointerEvent)) {
	macro.Rewrite("$_.onlostpointercapture = $1", onlostpointercapture)
}

func (*HTMLButtonElement) Onmsgesturechange() (onmsgesturechange func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturechange")
	return onmsgesturechange
}

func (*HTMLButtonElement) SetOnmsgesturechange(onmsgesturechange func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturechange = $1", onmsgesturechange)
}

func (*HTMLButtonElement) Onmsgesturedoubletap() (onmsgesturedoubletap func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturedoubletap")
	return onmsgesturedoubletap
}

func (*HTMLButtonElement) SetOnmsgesturedoubletap(onmsgesturedoubletap func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturedoubletap = $1", onmsgesturedoubletap)
}

func (*HTMLButtonElement) Onmsgestureend() (onmsgestureend func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgestureend")
	return onmsgestureend
}

func (*HTMLButtonElement) SetOnmsgestureend(onmsgestureend func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgestureend = $1", onmsgestureend)
}

func (*HTMLButtonElement) Onmsgesturehold() (onmsgesturehold func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturehold")
	return onmsgesturehold
}

func (*HTMLButtonElement) SetOnmsgesturehold(onmsgesturehold func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturehold = $1", onmsgesturehold)
}

func (*HTMLButtonElement) Onmsgesturestart() (onmsgesturestart func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturestart")
	return onmsgesturestart
}

func (*HTMLButtonElement) SetOnmsgesturestart(onmsgesturestart func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturestart = $1", onmsgesturestart)
}

func (*HTMLButtonElement) Onmsgesturetap() (onmsgesturetap func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturetap")
	return onmsgesturetap
}

func (*HTMLButtonElement) SetOnmsgesturetap(onmsgesturetap func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturetap = $1", onmsgesturetap)
}

func (*HTMLButtonElement) Onmsgotpointercapture() (onmsgotpointercapture func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmsgotpointercapture")
	return onmsgotpointercapture
}

func (*HTMLButtonElement) SetOnmsgotpointercapture(onmsgotpointercapture func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmsgotpointercapture = $1", onmsgotpointercapture)
}

func (*HTMLButtonElement) Onmsinertiastart() (onmsinertiastart func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsinertiastart")
	return onmsinertiastart
}

func (*HTMLButtonElement) SetOnmsinertiastart(onmsinertiastart func(*ms.MSGestureEvent)) {
	macro.Rewrite("$_.onmsinertiastart = $1", onmsinertiastart)
}

func (*HTMLButtonElement) Onmslostpointercapture() (onmslostpointercapture func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmslostpointercapture")
	return onmslostpointercapture
}

func (*HTMLButtonElement) SetOnmslostpointercapture(onmslostpointercapture func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmslostpointercapture = $1", onmslostpointercapture)
}

func (*HTMLButtonElement) Onmspointercancel() (onmspointercancel func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointercancel")
	return onmspointercancel
}

func (*HTMLButtonElement) SetOnmspointercancel(onmspointercancel func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointercancel = $1", onmspointercancel)
}

func (*HTMLButtonElement) Onmspointerdown() (onmspointerdown func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerdown")
	return onmspointerdown
}

func (*HTMLButtonElement) SetOnmspointerdown(onmspointerdown func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerdown = $1", onmspointerdown)
}

func (*HTMLButtonElement) Onmspointerenter() (onmspointerenter func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerenter")
	return onmspointerenter
}

func (*HTMLButtonElement) SetOnmspointerenter(onmspointerenter func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerenter = $1", onmspointerenter)
}

func (*HTMLButtonElement) Onmspointerleave() (onmspointerleave func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerleave")
	return onmspointerleave
}

func (*HTMLButtonElement) SetOnmspointerleave(onmspointerleave func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerleave = $1", onmspointerleave)
}

func (*HTMLButtonElement) Onmspointermove() (onmspointermove func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointermove")
	return onmspointermove
}

func (*HTMLButtonElement) SetOnmspointermove(onmspointermove func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointermove = $1", onmspointermove)
}

func (*HTMLButtonElement) Onmspointerout() (onmspointerout func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerout")
	return onmspointerout
}

func (*HTMLButtonElement) SetOnmspointerout(onmspointerout func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerout = $1", onmspointerout)
}

func (*HTMLButtonElement) Onmspointerover() (onmspointerover func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerover")
	return onmspointerover
}

func (*HTMLButtonElement) SetOnmspointerover(onmspointerover func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerover = $1", onmspointerover)
}

func (*HTMLButtonElement) Onmspointerup() (onmspointerup func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerup")
	return onmspointerup
}

func (*HTMLButtonElement) SetOnmspointerup(onmspointerup func(*ms.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerup = $1", onmspointerup)
}

func (*HTMLButtonElement) Ontouchcancel() (ontouchcancel func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchcancel")
	return ontouchcancel
}

func (*HTMLButtonElement) SetOntouchcancel(ontouchcancel func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchcancel = $1", ontouchcancel)
}

func (*HTMLButtonElement) Ontouchend() (ontouchend func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchend")
	return ontouchend
}

func (*HTMLButtonElement) SetOntouchend(ontouchend func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchend = $1", ontouchend)
}

func (*HTMLButtonElement) Ontouchmove() (ontouchmove func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchmove")
	return ontouchmove
}

func (*HTMLButtonElement) SetOntouchmove(ontouchmove func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchmove = $1", ontouchmove)
}

func (*HTMLButtonElement) Ontouchstart() (ontouchstart func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchstart")
	return ontouchstart
}

func (*HTMLButtonElement) SetOntouchstart(ontouchstart func(*touch.TouchEvent)) {
	macro.Rewrite("$_.ontouchstart = $1", ontouchstart)
}

func (*HTMLButtonElement) Onwebkitfullscreenchange() (onwebkitfullscreenchange func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenchange")
	return onwebkitfullscreenchange
}

func (*HTMLButtonElement) SetOnwebkitfullscreenchange(onwebkitfullscreenchange func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenchange = $1", onwebkitfullscreenchange)
}

func (*HTMLButtonElement) Onwebkitfullscreenerror() (onwebkitfullscreenerror func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenerror")
	return onwebkitfullscreenerror
}

func (*HTMLButtonElement) SetOnwebkitfullscreenerror(onwebkitfullscreenerror func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenerror = $1", onwebkitfullscreenerror)
}

func (*HTMLButtonElement) OuterHTML() (outerHTML string) {
	macro.Rewrite("$_.outerHTML")
	return outerHTML
}

func (*HTMLButtonElement) SetOuterHTML(outerHTML string) {
	macro.Rewrite("$_.outerHTML = $1", outerHTML)
}

func (*HTMLButtonElement) Prefix() (prefix string) {
	macro.Rewrite("$_.prefix")
	return prefix
}

func (*HTMLButtonElement) ScrollHeight() (scrollHeight int) {
	macro.Rewrite("$_.scrollHeight")
	return scrollHeight
}

func (*HTMLButtonElement) ScrollLeft() (scrollLeft int) {
	macro.Rewrite("$_.scrollLeft")
	return scrollLeft
}

func (*HTMLButtonElement) SetScrollLeft(scrollLeft int) {
	macro.Rewrite("$_.scrollLeft = $1", scrollLeft)
}

func (*HTMLButtonElement) ScrollTop() (scrollTop int) {
	macro.Rewrite("$_.scrollTop")
	return scrollTop
}

func (*HTMLButtonElement) SetScrollTop(scrollTop int) {
	macro.Rewrite("$_.scrollTop = $1", scrollTop)
}

func (*HTMLButtonElement) ScrollWidth() (scrollWidth int) {
	macro.Rewrite("$_.scrollWidth")
	return scrollWidth
}

func (*HTMLButtonElement) TagName() (tagName string) {
	macro.Rewrite("$_.tagName")
	return tagName
}

func (*HTMLButtonElement) Onpointercancel() (onpointercancel func(event.Event)) {
	macro.Rewrite("$_.onpointercancel")
	return onpointercancel
}

func (*HTMLButtonElement) SetOnpointercancel(onpointercancel func(event.Event)) {
	macro.Rewrite("$_.onpointercancel = $1", onpointercancel)
}

func (*HTMLButtonElement) Onpointerdown() (onpointerdown func(event.Event)) {
	macro.Rewrite("$_.onpointerdown")
	return onpointerdown
}

func (*HTMLButtonElement) SetOnpointerdown(onpointerdown func(event.Event)) {
	macro.Rewrite("$_.onpointerdown = $1", onpointerdown)
}

func (*HTMLButtonElement) Onpointerenter() (onpointerenter func(event.Event)) {
	macro.Rewrite("$_.onpointerenter")
	return onpointerenter
}

func (*HTMLButtonElement) SetOnpointerenter(onpointerenter func(event.Event)) {
	macro.Rewrite("$_.onpointerenter = $1", onpointerenter)
}

func (*HTMLButtonElement) Onpointerleave() (onpointerleave func(event.Event)) {
	macro.Rewrite("$_.onpointerleave")
	return onpointerleave
}

func (*HTMLButtonElement) SetOnpointerleave(onpointerleave func(event.Event)) {
	macro.Rewrite("$_.onpointerleave = $1", onpointerleave)
}

func (*HTMLButtonElement) Onpointermove() (onpointermove func(event.Event)) {
	macro.Rewrite("$_.onpointermove")
	return onpointermove
}

func (*HTMLButtonElement) SetOnpointermove(onpointermove func(event.Event)) {
	macro.Rewrite("$_.onpointermove = $1", onpointermove)
}

func (*HTMLButtonElement) Onpointerout() (onpointerout func(event.Event)) {
	macro.Rewrite("$_.onpointerout")
	return onpointerout
}

func (*HTMLButtonElement) SetOnpointerout(onpointerout func(event.Event)) {
	macro.Rewrite("$_.onpointerout = $1", onpointerout)
}

func (*HTMLButtonElement) Onpointerover() (onpointerover func(event.Event)) {
	macro.Rewrite("$_.onpointerover")
	return onpointerover
}

func (*HTMLButtonElement) SetOnpointerover(onpointerover func(event.Event)) {
	macro.Rewrite("$_.onpointerover = $1", onpointerover)
}

func (*HTMLButtonElement) Onpointerup() (onpointerup func(event.Event)) {
	macro.Rewrite("$_.onpointerup")
	return onpointerup
}

func (*HTMLButtonElement) SetOnpointerup(onpointerup func(event.Event)) {
	macro.Rewrite("$_.onpointerup = $1", onpointerup)
}

func (*HTMLButtonElement) Onwheel() (onwheel func(event.Event)) {
	macro.Rewrite("$_.onwheel")
	return onwheel
}

func (*HTMLButtonElement) SetOnwheel(onwheel func(event.Event)) {
	macro.Rewrite("$_.onwheel = $1", onwheel)
}

func (*HTMLButtonElement) ChildElementCount() (childElementCount uint) {
	macro.Rewrite("$_.childElementCount")
	return childElementCount
}

func (*HTMLButtonElement) FirstElementChild() (firstElementChild element.Element) {
	macro.Rewrite("$_.firstElementChild")
	return firstElementChild
}

func (*HTMLButtonElement) LastElementChild() (lastElementChild element.Element) {
	macro.Rewrite("$_.lastElementChild")
	return lastElementChild
}

func (*HTMLButtonElement) NextElementSibling() (nextElementSibling element.Element) {
	macro.Rewrite("$_.nextElementSibling")
	return nextElementSibling
}

func (*HTMLButtonElement) PreviousElementSibling() (previousElementSibling element.Element) {
	macro.Rewrite("$_.previousElementSibling")
	return previousElementSibling
}

func (*HTMLButtonElement) Attributes() (attributes *dom.NamedNodeMap) {
	macro.Rewrite("$_.attributes")
	return attributes
}

func (*HTMLButtonElement) BaseURI() (baseURI string) {
	macro.Rewrite("$_.baseURI")
	return baseURI
}

func (*HTMLButtonElement) ChildNodes() (childNodes *dom.NodeList) {
	macro.Rewrite("$_.childNodes")
	return childNodes
}

func (*HTMLButtonElement) FirstChild() (firstChild dom.Node) {
	macro.Rewrite("$_.firstChild")
	return firstChild
}

func (*HTMLButtonElement) LastChild() (lastChild dom.Node) {
	macro.Rewrite("$_.lastChild")
	return lastChild
}

func (*HTMLButtonElement) LocalName() (localName string) {
	macro.Rewrite("$_.localName")
	return localName
}

func (*HTMLButtonElement) NamespaceURI() (namespaceURI string) {
	macro.Rewrite("$_.namespaceURI")
	return namespaceURI
}

func (*HTMLButtonElement) NextSibling() (nextSibling dom.Node) {
	macro.Rewrite("$_.nextSibling")
	return nextSibling
}

func (*HTMLButtonElement) NodeName() (nodeName string) {
	macro.Rewrite("$_.nodeName")
	return nodeName
}

func (*HTMLButtonElement) NodeType() (nodeType uint8) {
	macro.Rewrite("$_.nodeType")
	return nodeType
}

func (*HTMLButtonElement) NodeValue() (nodeValue string) {
	macro.Rewrite("$_.nodeValue")
	return nodeValue
}

func (*HTMLButtonElement) SetNodeValue(nodeValue string) {
	macro.Rewrite("$_.nodeValue = $1", nodeValue)
}

func (*HTMLButtonElement) OwnerDocument() (ownerDocument document.Document) {
	macro.Rewrite("$_.ownerDocument")
	return ownerDocument
}

func (*HTMLButtonElement) ParentElement() (parentElement HTMLElement) {
	macro.Rewrite("$_.parentElement")
	return parentElement
}

func (*HTMLButtonElement) ParentNode() (parentNode dom.Node) {
	macro.Rewrite("$_.parentNode")
	return parentNode
}

func (*HTMLButtonElement) PreviousSibling() (previousSibling dom.Node) {
	macro.Rewrite("$_.previousSibling")
	return previousSibling
}

func (*HTMLButtonElement) TextContent() (textContent string) {
	macro.Rewrite("$_.textContent")
	return textContent
}

func (*HTMLButtonElement) SetTextContent(textContent string) {
	macro.Rewrite("$_.textContent = $1", textContent)
}
