package html


import (
	"github.com/matthewmueller/joy/dom/childnode"
	"github.com/matthewmueller/joy/dom/clientrect"
	"github.com/matthewmueller/joy/dom/clientrectlist"
	"github.com/matthewmueller/joy/dom/domstringmap"
	"github.com/matthewmueller/joy/dom/domtokenlist"
	"github.com/matthewmueller/joy/dom/htmlformelement"
	"github.com/matthewmueller/joy/dom/mszoomtooptions"
	"github.com/matthewmueller/joy/dom/validitystate"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ window.HTMLElement = (*HTMLObjectElement)(nil)
var _ window.Element = (*HTMLObjectElement)(nil)
var _ window.GlobalEventHandlers = (*HTMLObjectElement)(nil)
var _ window.ElementTraversal = (*HTMLObjectElement)(nil)
var _ window.NodeSelector = (*HTMLObjectElement)(nil)
var _ childnode.ChildNode = (*HTMLObjectElement)(nil)
var _ window.Node = (*HTMLObjectElement)(nil)
var _ window.EventTarget = (*HTMLObjectElement)(nil)

// HTMLObjectElement struct
// js:"HTMLObjectElement,omit"
type HTMLObjectElement struct {
}

// CheckValidity fn Returns whether a form will validate when it is submitted, without having to submit it.
// js:"checkValidity"
func (*HTMLObjectElement) CheckValidity() (b bool) {
	macro.Rewrite("$_.checkValidity()")
	return b
}

// SetCustomValidity fn Sets a custom error message that is displayed when a form is submitted.
//     * @param error Sets a custom error message that is displayed when a form is submitted.
// js:"setCustomValidity"
func (*HTMLObjectElement) SetCustomValidity(err string) {
	macro.Rewrite("$_.setCustomValidity($1)", err)
}

// GetSVGDocument fn
// js:"getSVGDocument"
func (*HTMLObjectElement) GetSVGDocument() (w window.Document) {
	macro.Rewrite("$_.getSVGDocument()")
	return w
}

// Blur fn
// js:"blur"
func (*HTMLObjectElement) Blur() {
	macro.Rewrite("$_.blur()")
}

// Click fn
// js:"click"
func (*HTMLObjectElement) Click() {
	macro.Rewrite("$_.click()")
}

// DragDrop fn
// js:"dragDrop"
func (*HTMLObjectElement) DragDrop() (b bool) {
	macro.Rewrite("$_.dragDrop()")
	return b
}

// Focus fn
// js:"focus"
func (*HTMLObjectElement) Focus() {
	macro.Rewrite("$_.focus()")
}

// GetElementsByClassName fn
// js:"getElementsByClassName"
func (*HTMLObjectElement) GetElementsByClassName(classNames string) (w *window.NodeList) {
	macro.Rewrite("$_.getElementsByClassName($1)", classNames)
	return w
}

// InsertAdjacentElement fn
// js:"insertAdjacentElement"
func (*HTMLObjectElement) InsertAdjacentElement(position string, insertedElement window.Element) (w window.Element) {
	macro.Rewrite("$_.insertAdjacentElement($1, $2)", position, insertedElement)
	return w
}

// InsertAdjacentHTML fn
// js:"insertAdjacentHTML"
func (*HTMLObjectElement) InsertAdjacentHTML(where string, html string) {
	macro.Rewrite("$_.insertAdjacentHTML($1, $2)", where, html)
}

// InsertAdjacentText fn
// js:"insertAdjacentText"
func (*HTMLObjectElement) InsertAdjacentText(where string, text string) {
	macro.Rewrite("$_.insertAdjacentText($1, $2)", where, text)
}

// MsGetInputContext fn
// js:"msGetInputContext"
func (*HTMLObjectElement) MsGetInputContext() (w *window.MSInputMethodContext) {
	macro.Rewrite("$_.msGetInputContext()")
	return w
}

// ScrollIntoView fn
// js:"scrollIntoView"
func (*HTMLObjectElement) ScrollIntoView(top *bool) {
	macro.Rewrite("$_.scrollIntoView($1)", top)
}

// GetAttribute fn
// js:"getAttribute"
func (*HTMLObjectElement) GetAttribute(qualifiedName string) (s string) {
	macro.Rewrite("$_.getAttribute($1)", qualifiedName)
	return s
}

// GetAttributeNode fn
// js:"getAttributeNode"
func (*HTMLObjectElement) GetAttributeNode(name string) (w *window.Attr) {
	macro.Rewrite("$_.getAttributeNode($1)", name)
	return w
}

// GetAttributeNodeNS fn
// js:"getAttributeNodeNS"
func (*HTMLObjectElement) GetAttributeNodeNS(namespaceURI string, localName string) (w *window.Attr) {
	macro.Rewrite("$_.getAttributeNodeNS($1, $2)", namespaceURI, localName)
	return w
}

// GetAttributeNS fn
// js:"getAttributeNS"
func (*HTMLObjectElement) GetAttributeNS(namespaceURI string, localName string) (s string) {
	macro.Rewrite("$_.getAttributeNS($1, $2)", namespaceURI, localName)
	return s
}

// GetBoundingClientRect fn
// js:"getBoundingClientRect"
func (*HTMLObjectElement) GetBoundingClientRect() (c *clientrect.ClientRect) {
	macro.Rewrite("$_.getBoundingClientRect()")
	return c
}

// GetClientRects fn
// js:"getClientRects"
func (*HTMLObjectElement) GetClientRects() (c *clientrectlist.ClientRectList) {
	macro.Rewrite("$_.getClientRects()")
	return c
}

// GetElementsByTagName fn
// js:"getElementsByTagName"
func (*HTMLObjectElement) GetElementsByTagName(name string) (w *window.NodeList) {
	macro.Rewrite("$_.getElementsByTagName($1)", name)
	return w
}

// GetElementsByTagNameNS fn
// js:"getElementsByTagNameNS"
func (*HTMLObjectElement) GetElementsByTagNameNS(namespaceURI string, localName string) (w *window.NodeList) {
	macro.Rewrite("$_.getElementsByTagNameNS($1, $2)", namespaceURI, localName)
	return w
}

// HasAttribute fn
// js:"hasAttribute"
func (*HTMLObjectElement) HasAttribute(name string) (b bool) {
	macro.Rewrite("$_.hasAttribute($1)", name)
	return b
}

// HasAttributeNS fn
// js:"hasAttributeNS"
func (*HTMLObjectElement) HasAttributeNS(namespaceURI string, localName string) (b bool) {
	macro.Rewrite("$_.hasAttributeNS($1, $2)", namespaceURI, localName)
	return b
}

// MsGetRegionContent fn
// js:"msGetRegionContent"
func (*HTMLObjectElement) MsGetRegionContent() (w *window.MSRangeCollection) {
	macro.Rewrite("$_.msGetRegionContent()")
	return w
}

// MsGetUntransformedBounds fn
// js:"msGetUntransformedBounds"
func (*HTMLObjectElement) MsGetUntransformedBounds() (c *clientrect.ClientRect) {
	macro.Rewrite("$_.msGetUntransformedBounds()")
	return c
}

// MsMatchesSelector fn
// js:"msMatchesSelector"
func (*HTMLObjectElement) MsMatchesSelector(selectors string) (b bool) {
	macro.Rewrite("$_.msMatchesSelector($1)", selectors)
	return b
}

// MsReleasePointerCapture fn
// js:"msReleasePointerCapture"
func (*HTMLObjectElement) MsReleasePointerCapture(pointerId int) {
	macro.Rewrite("$_.msReleasePointerCapture($1)", pointerId)
}

// MsSetPointerCapture fn
// js:"msSetPointerCapture"
func (*HTMLObjectElement) MsSetPointerCapture(pointerId int) {
	macro.Rewrite("$_.msSetPointerCapture($1)", pointerId)
}

// MsZoomTo fn
// js:"msZoomTo"
func (*HTMLObjectElement) MsZoomTo(args *mszoomtooptions.MsZoomToOptions) {
	macro.Rewrite("$_.msZoomTo($1)", args)
}

// ReleasePointerCapture fn
// js:"releasePointerCapture"
func (*HTMLObjectElement) ReleasePointerCapture(pointerId int) {
	macro.Rewrite("$_.releasePointerCapture($1)", pointerId)
}

// RemoveAttribute fn
// js:"removeAttribute"
func (*HTMLObjectElement) RemoveAttribute(qualifiedName string) {
	macro.Rewrite("$_.removeAttribute($1)", qualifiedName)
}

// RemoveAttributeNode fn
// js:"removeAttributeNode"
func (*HTMLObjectElement) RemoveAttributeNode(oldAttr *window.Attr) (w *window.Attr) {
	macro.Rewrite("$_.removeAttributeNode($1)", oldAttr)
	return w
}

// RemoveAttributeNS fn
// js:"removeAttributeNS"
func (*HTMLObjectElement) RemoveAttributeNS(namespaceURI string, localName string) {
	macro.Rewrite("$_.removeAttributeNS($1, $2)", namespaceURI, localName)
}

// RequestFullscreen fn
// js:"requestFullscreen"
func (*HTMLObjectElement) RequestFullscreen() {
	macro.Rewrite("$_.requestFullscreen()")
}

// RequestPointerLock fn
// js:"requestPointerLock"
func (*HTMLObjectElement) RequestPointerLock() {
	macro.Rewrite("$_.requestPointerLock()")
}

// SetAttribute fn
// js:"setAttribute"
func (*HTMLObjectElement) SetAttribute(qualifiedName string, value string) {
	macro.Rewrite("$_.setAttribute($1, $2)", qualifiedName, value)
}

// SetAttributeNode fn
// js:"setAttributeNode"
func (*HTMLObjectElement) SetAttributeNode(newAttr *window.Attr) (w *window.Attr) {
	macro.Rewrite("$_.setAttributeNode($1)", newAttr)
	return w
}

// SetAttributeNodeNS fn
// js:"setAttributeNodeNS"
func (*HTMLObjectElement) SetAttributeNodeNS(newAttr *window.Attr) (w *window.Attr) {
	macro.Rewrite("$_.setAttributeNodeNS($1)", newAttr)
	return w
}

// SetAttributeNS fn
// js:"setAttributeNS"
func (*HTMLObjectElement) SetAttributeNS(namespaceURI string, qualifiedName string, value string) {
	macro.Rewrite("$_.setAttributeNS($1, $2, $3)", namespaceURI, qualifiedName, value)
}

// SetPointerCapture fn
// js:"setPointerCapture"
func (*HTMLObjectElement) SetPointerCapture(pointerId int) {
	macro.Rewrite("$_.setPointerCapture($1)", pointerId)
}

// WebkitMatchesSelector fn
// js:"webkitMatchesSelector"
func (*HTMLObjectElement) WebkitMatchesSelector(selectors string) (b bool) {
	macro.Rewrite("$_.webkitMatchesSelector($1)", selectors)
	return b
}

// WebkitRequestFullscreen fn
// js:"webkitRequestFullscreen"
func (*HTMLObjectElement) WebkitRequestFullscreen() {
	macro.Rewrite("$_.webkitRequestFullscreen()")
}

// WebkitRequestFullScreen fn
// js:"webkitRequestFullScreen"
func (*HTMLObjectElement) WebkitRequestFullScreen() {
	macro.Rewrite("$_.webkitRequestFullScreen()")
}

// QuerySelector fn
// js:"querySelector"
func (*HTMLObjectElement) QuerySelector(selectors string) (w window.Element) {
	macro.Rewrite("$_.querySelector($1)", selectors)
	return w
}

// QuerySelectorAll fn
// js:"querySelectorAll"
func (*HTMLObjectElement) QuerySelectorAll(selectors string) (w *window.NodeList) {
	macro.Rewrite("$_.querySelectorAll($1)", selectors)
	return w
}

// AppendChild fn
// js:"appendChild"
func (*HTMLObjectElement) AppendChild(newChild window.Node) (w window.Node) {
	macro.Rewrite("$_.appendChild($1)", newChild)
	return w
}

// CloneNode fn
// js:"cloneNode"
func (*HTMLObjectElement) CloneNode(deep *bool) (w window.Node) {
	macro.Rewrite("$_.cloneNode($1)", deep)
	return w
}

// CompareDocumentPosition fn
// js:"compareDocumentPosition"
func (*HTMLObjectElement) CompareDocumentPosition(other window.Node) (u uint8) {
	macro.Rewrite("$_.compareDocumentPosition($1)", other)
	return u
}

// Contains fn
// js:"contains"
func (*HTMLObjectElement) Contains(child window.Node) (b bool) {
	macro.Rewrite("$_.contains($1)", child)
	return b
}

// HasAttributes fn
// js:"hasAttributes"
func (*HTMLObjectElement) HasAttributes() (b bool) {
	macro.Rewrite("$_.hasAttributes()")
	return b
}

// HasChildNodes fn
// js:"hasChildNodes"
func (*HTMLObjectElement) HasChildNodes() (b bool) {
	macro.Rewrite("$_.hasChildNodes()")
	return b
}

// InsertBefore fn
// js:"insertBefore"
func (*HTMLObjectElement) InsertBefore(newChild window.Node, refChild *window.Node) (w window.Node) {
	macro.Rewrite("$_.insertBefore($1, $2)", newChild, refChild)
	return w
}

// IsDefaultNamespace fn
// js:"isDefaultNamespace"
func (*HTMLObjectElement) IsDefaultNamespace(namespaceURI string) (b bool) {
	macro.Rewrite("$_.isDefaultNamespace($1)", namespaceURI)
	return b
}

// IsEqualNode fn
// js:"isEqualNode"
func (*HTMLObjectElement) IsEqualNode(arg window.Node) (b bool) {
	macro.Rewrite("$_.isEqualNode($1)", arg)
	return b
}

// IsSameNode fn
// js:"isSameNode"
func (*HTMLObjectElement) IsSameNode(other window.Node) (b bool) {
	macro.Rewrite("$_.isSameNode($1)", other)
	return b
}

// LookupNamespaceURI fn
// js:"lookupNamespaceURI"
func (*HTMLObjectElement) LookupNamespaceURI(prefix string) (s string) {
	macro.Rewrite("$_.lookupNamespaceURI($1)", prefix)
	return s
}

// LookupPrefix fn
// js:"lookupPrefix"
func (*HTMLObjectElement) LookupPrefix(namespaceURI string) (s string) {
	macro.Rewrite("$_.lookupPrefix($1)", namespaceURI)
	return s
}

// Normalize fn
// js:"normalize"
func (*HTMLObjectElement) Normalize() {
	macro.Rewrite("$_.normalize()")
}

// RemoveChild fn
// js:"removeChild"
func (*HTMLObjectElement) RemoveChild(oldChild window.Node) (w window.Node) {
	macro.Rewrite("$_.removeChild($1)", oldChild)
	return w
}

// ReplaceChild fn
// js:"replaceChild"
func (*HTMLObjectElement) ReplaceChild(newChild window.Node, oldChild window.Node) (w window.Node) {
	macro.Rewrite("$_.replaceChild($1, $2)", newChild, oldChild)
	return w
}

// AddEventListener fn
// js:"addEventListener"
func (*HTMLObjectElement) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*HTMLObjectElement) DispatchEvent(evt window.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*HTMLObjectElement) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Align prop
// js:"align"
func (*HTMLObjectElement) Align() (align string) {
	macro.Rewrite("$_.align")
	return align
}

// SetAlign prop
// js:"align"
func (*HTMLObjectElement) SetAlign(align string) {
	macro.Rewrite("$_.align = $1", align)
}

// Alt prop Sets or retrieves a text alternative to the graphic.
// js:"alt"
func (*HTMLObjectElement) Alt() (alt string) {
	macro.Rewrite("$_.alt")
	return alt
}

// SetAlt prop Sets or retrieves a text alternative to the graphic.
// js:"alt"
func (*HTMLObjectElement) SetAlt(alt string) {
	macro.Rewrite("$_.alt = $1", alt)
}

// AltHTML prop Gets or sets the optional alternative HTML script to execute if the object fails to load.
// js:"altHtml"
func (*HTMLObjectElement) AltHTML() (altHtml string) {
	macro.Rewrite("$_.altHtml")
	return altHtml
}

// SetAltHTML prop Gets or sets the optional alternative HTML script to execute if the object fails to load.
// js:"altHtml"
func (*HTMLObjectElement) SetAltHTML(altHtml string) {
	macro.Rewrite("$_.altHtml = $1", altHtml)
}

// Archive prop Sets or retrieves a character string that can be used to implement your own archive functionality for the object.
// js:"archive"
func (*HTMLObjectElement) Archive() (archive string) {
	macro.Rewrite("$_.archive")
	return archive
}

// SetArchive prop Sets or retrieves a character string that can be used to implement your own archive functionality for the object.
// js:"archive"
func (*HTMLObjectElement) SetArchive(archive string) {
	macro.Rewrite("$_.archive = $1", archive)
}

// BaseHref prop Retrieves a string of the URL where the object tag can be found. This is often the href of the document that the object is in, or the value set by a base element.
// js:"BaseHref"
func (*HTMLObjectElement) BaseHref() (BaseHref string) {
	macro.Rewrite("$_.BaseHref")
	return BaseHref
}

// Border prop
// js:"border"
func (*HTMLObjectElement) Border() (border string) {
	macro.Rewrite("$_.border")
	return border
}

// SetBorder prop
// js:"border"
func (*HTMLObjectElement) SetBorder(border string) {
	macro.Rewrite("$_.border = $1", border)
}

// Code prop Sets or retrieves the URL of the file containing the compiled Java class.
// js:"code"
func (*HTMLObjectElement) Code() (code string) {
	macro.Rewrite("$_.code")
	return code
}

// SetCode prop Sets or retrieves the URL of the file containing the compiled Java class.
// js:"code"
func (*HTMLObjectElement) SetCode(code string) {
	macro.Rewrite("$_.code = $1", code)
}

// CodeBase prop Sets or retrieves the URL of the component.
// js:"codeBase"
func (*HTMLObjectElement) CodeBase() (codeBase string) {
	macro.Rewrite("$_.codeBase")
	return codeBase
}

// SetCodeBase prop Sets or retrieves the URL of the component.
// js:"codeBase"
func (*HTMLObjectElement) SetCodeBase(codeBase string) {
	macro.Rewrite("$_.codeBase = $1", codeBase)
}

// CodeType prop Sets or retrieves the Internet media type for the code associated with the object.
// js:"codeType"
func (*HTMLObjectElement) CodeType() (codeType string) {
	macro.Rewrite("$_.codeType")
	return codeType
}

// SetCodeType prop Sets or retrieves the Internet media type for the code associated with the object.
// js:"codeType"
func (*HTMLObjectElement) SetCodeType(codeType string) {
	macro.Rewrite("$_.codeType = $1", codeType)
}

// ContentDocument prop Retrieves the document object of the page or frame.
// js:"contentDocument"
func (*HTMLObjectElement) ContentDocument() (contentDocument window.Document) {
	macro.Rewrite("$_.contentDocument")
	return contentDocument
}

// Data prop Sets or retrieves the URL that references the data of the object.
// js:"data"
func (*HTMLObjectElement) Data() (data string) {
	macro.Rewrite("$_.data")
	return data
}

// SetData prop Sets or retrieves the URL that references the data of the object.
// js:"data"
func (*HTMLObjectElement) SetData(data string) {
	macro.Rewrite("$_.data = $1", data)
}

// Declare prop
// js:"declare"
func (*HTMLObjectElement) Declare() (declare bool) {
	macro.Rewrite("$_.declare")
	return declare
}

// SetDeclare prop
// js:"declare"
func (*HTMLObjectElement) SetDeclare(declare bool) {
	macro.Rewrite("$_.declare = $1", declare)
}

// Form prop Retrieves a reference to the form that the object is embedded in.
// js:"form"
func (*HTMLObjectElement) Form() (form *htmlformelement.HTMLFormElement) {
	macro.Rewrite("$_.form")
	return form
}

// Height prop Sets or retrieves the height of the object.
// js:"height"
func (*HTMLObjectElement) Height() (height string) {
	macro.Rewrite("$_.height")
	return height
}

// SetHeight prop Sets or retrieves the height of the object.
// js:"height"
func (*HTMLObjectElement) SetHeight(height string) {
	macro.Rewrite("$_.height = $1", height)
}

// Hspace prop
// js:"hspace"
func (*HTMLObjectElement) Hspace() (hspace int) {
	macro.Rewrite("$_.hspace")
	return hspace
}

// SetHspace prop
// js:"hspace"
func (*HTMLObjectElement) SetHspace(hspace int) {
	macro.Rewrite("$_.hspace = $1", hspace)
}

// MsPlayToDisabled prop Gets or sets whether the DLNA PlayTo device is available.
// js:"msPlayToDisabled"
func (*HTMLObjectElement) MsPlayToDisabled() (msPlayToDisabled bool) {
	macro.Rewrite("$_.msPlayToDisabled")
	return msPlayToDisabled
}

// SetMsPlayToDisabled prop Gets or sets whether the DLNA PlayTo device is available.
// js:"msPlayToDisabled"
func (*HTMLObjectElement) SetMsPlayToDisabled(msPlayToDisabled bool) {
	macro.Rewrite("$_.msPlayToDisabled = $1", msPlayToDisabled)
}

// MsPlayToPreferredSourceURI prop Gets or sets the path to the preferred media source. This enables the Play To target device to stream the media content, which can be DRM protected, from a different location, such as a cloud media server.
// js:"msPlayToPreferredSourceUri"
func (*HTMLObjectElement) MsPlayToPreferredSourceURI() (msPlayToPreferredSourceUri string) {
	macro.Rewrite("$_.msPlayToPreferredSourceUri")
	return msPlayToPreferredSourceUri
}

// SetMsPlayToPreferredSourceURI prop Gets or sets the path to the preferred media source. This enables the Play To target device to stream the media content, which can be DRM protected, from a different location, such as a cloud media server.
// js:"msPlayToPreferredSourceUri"
func (*HTMLObjectElement) SetMsPlayToPreferredSourceURI(msPlayToPreferredSourceUri string) {
	macro.Rewrite("$_.msPlayToPreferredSourceUri = $1", msPlayToPreferredSourceUri)
}

// MsPlayToPrimary prop Gets or sets the primary DLNA PlayTo device.
// js:"msPlayToPrimary"
func (*HTMLObjectElement) MsPlayToPrimary() (msPlayToPrimary bool) {
	macro.Rewrite("$_.msPlayToPrimary")
	return msPlayToPrimary
}

// SetMsPlayToPrimary prop Gets or sets the primary DLNA PlayTo device.
// js:"msPlayToPrimary"
func (*HTMLObjectElement) SetMsPlayToPrimary(msPlayToPrimary bool) {
	macro.Rewrite("$_.msPlayToPrimary = $1", msPlayToPrimary)
}

// MsPlayToSource prop Gets the source associated with the media element for use by the PlayToManager.
// js:"msPlayToSource"
func (*HTMLObjectElement) MsPlayToSource() (msPlayToSource interface{}) {
	macro.Rewrite("$_.msPlayToSource")
	return msPlayToSource
}

// Name prop Sets or retrieves the name of the object.
// js:"name"
func (*HTMLObjectElement) Name() (name string) {
	macro.Rewrite("$_.name")
	return name
}

// SetName prop Sets or retrieves the name of the object.
// js:"name"
func (*HTMLObjectElement) SetName(name string) {
	macro.Rewrite("$_.name = $1", name)
}

// ReadyState prop
// js:"readyState"
func (*HTMLObjectElement) ReadyState() (readyState uint8) {
	macro.Rewrite("$_.readyState")
	return readyState
}

// Standby prop Sets or retrieves a message to be displayed while an object is loading.
// js:"standby"
func (*HTMLObjectElement) Standby() (standby string) {
	macro.Rewrite("$_.standby")
	return standby
}

// SetStandby prop Sets or retrieves a message to be displayed while an object is loading.
// js:"standby"
func (*HTMLObjectElement) SetStandby(standby string) {
	macro.Rewrite("$_.standby = $1", standby)
}

// Type prop Sets or retrieves the MIME type of the object.
// js:"type"
func (*HTMLObjectElement) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

// SetType prop Sets or retrieves the MIME type of the object.
// js:"type"
func (*HTMLObjectElement) SetType(kind string) {
	macro.Rewrite("$_.type = $1", kind)
}

// UseMap prop Sets or retrieves the URL, often with a bookmark extension (#name), to use as a client-side image map.
// js:"useMap"
func (*HTMLObjectElement) UseMap() (useMap string) {
	macro.Rewrite("$_.useMap")
	return useMap
}

// SetUseMap prop Sets or retrieves the URL, often with a bookmark extension (#name), to use as a client-side image map.
// js:"useMap"
func (*HTMLObjectElement) SetUseMap(useMap string) {
	macro.Rewrite("$_.useMap = $1", useMap)
}

// ValidationMessage prop Returns the error message that would be displayed if the user submits the form, or an empty string if no error message. It also triggers the standard error message, such as "this is a required field". The result is that the user sees validation messages without actually submitting.
// js:"validationMessage"
func (*HTMLObjectElement) ValidationMessage() (validationMessage string) {
	macro.Rewrite("$_.validationMessage")
	return validationMessage
}

// Validity prop Returns a  ValidityState object that represents the validity states of an element.
// js:"validity"
func (*HTMLObjectElement) Validity() (validity *validitystate.ValidityState) {
	macro.Rewrite("$_.validity")
	return validity
}

// Vspace prop
// js:"vspace"
func (*HTMLObjectElement) Vspace() (vspace int) {
	macro.Rewrite("$_.vspace")
	return vspace
}

// SetVspace prop
// js:"vspace"
func (*HTMLObjectElement) SetVspace(vspace int) {
	macro.Rewrite("$_.vspace = $1", vspace)
}

// Width prop Sets or retrieves the width of the object.
// js:"width"
func (*HTMLObjectElement) Width() (width string) {
	macro.Rewrite("$_.width")
	return width
}

// SetWidth prop Sets or retrieves the width of the object.
// js:"width"
func (*HTMLObjectElement) SetWidth(width string) {
	macro.Rewrite("$_.width = $1", width)
}

// WillValidate prop Returns whether an element will successfully validate based on forms validation rules and constraints.
// js:"willValidate"
func (*HTMLObjectElement) WillValidate() (willValidate bool) {
	macro.Rewrite("$_.willValidate")
	return willValidate
}

// AccessKey prop
// js:"accessKey"
func (*HTMLObjectElement) AccessKey() (accessKey string) {
	macro.Rewrite("$_.accessKey")
	return accessKey
}

// SetAccessKey prop
// js:"accessKey"
func (*HTMLObjectElement) SetAccessKey(accessKey string) {
	macro.Rewrite("$_.accessKey = $1", accessKey)
}

// Children prop
// js:"children"
func (*HTMLObjectElement) Children() (children window.HTMLCollection) {
	macro.Rewrite("$_.children")
	return children
}

// ContentEditable prop
// js:"contentEditable"
func (*HTMLObjectElement) ContentEditable() (contentEditable string) {
	macro.Rewrite("$_.contentEditable")
	return contentEditable
}

// SetContentEditable prop
// js:"contentEditable"
func (*HTMLObjectElement) SetContentEditable(contentEditable string) {
	macro.Rewrite("$_.contentEditable = $1", contentEditable)
}

// Dataset prop
// js:"dataset"
func (*HTMLObjectElement) Dataset() (dataset *domstringmap.DOMStringMap) {
	macro.Rewrite("$_.dataset")
	return dataset
}

// Dir prop
// js:"dir"
func (*HTMLObjectElement) Dir() (dir string) {
	macro.Rewrite("$_.dir")
	return dir
}

// SetDir prop
// js:"dir"
func (*HTMLObjectElement) SetDir(dir string) {
	macro.Rewrite("$_.dir = $1", dir)
}

// Draggable prop
// js:"draggable"
func (*HTMLObjectElement) Draggable() (draggable bool) {
	macro.Rewrite("$_.draggable")
	return draggable
}

// SetDraggable prop
// js:"draggable"
func (*HTMLObjectElement) SetDraggable(draggable bool) {
	macro.Rewrite("$_.draggable = $1", draggable)
}

// Hidden prop
// js:"hidden"
func (*HTMLObjectElement) Hidden() (hidden bool) {
	macro.Rewrite("$_.hidden")
	return hidden
}

// SetHidden prop
// js:"hidden"
func (*HTMLObjectElement) SetHidden(hidden bool) {
	macro.Rewrite("$_.hidden = $1", hidden)
}

// HideFocus prop
// js:"hideFocus"
func (*HTMLObjectElement) HideFocus() (hideFocus bool) {
	macro.Rewrite("$_.hideFocus")
	return hideFocus
}

// SetHideFocus prop
// js:"hideFocus"
func (*HTMLObjectElement) SetHideFocus(hideFocus bool) {
	macro.Rewrite("$_.hideFocus = $1", hideFocus)
}

// InnerText prop
// js:"innerText"
func (*HTMLObjectElement) InnerText() (innerText string) {
	macro.Rewrite("$_.innerText")
	return innerText
}

// SetInnerText prop
// js:"innerText"
func (*HTMLObjectElement) SetInnerText(innerText string) {
	macro.Rewrite("$_.innerText = $1", innerText)
}

// IsContentEditable prop
// js:"isContentEditable"
func (*HTMLObjectElement) IsContentEditable() (isContentEditable bool) {
	macro.Rewrite("$_.isContentEditable")
	return isContentEditable
}

// Lang prop
// js:"lang"
func (*HTMLObjectElement) Lang() (lang string) {
	macro.Rewrite("$_.lang")
	return lang
}

// SetLang prop
// js:"lang"
func (*HTMLObjectElement) SetLang(lang string) {
	macro.Rewrite("$_.lang = $1", lang)
}

// OffsetHeight prop
// js:"offsetHeight"
func (*HTMLObjectElement) OffsetHeight() (offsetHeight int) {
	macro.Rewrite("$_.offsetHeight")
	return offsetHeight
}

// OffsetLeft prop
// js:"offsetLeft"
func (*HTMLObjectElement) OffsetLeft() (offsetLeft int) {
	macro.Rewrite("$_.offsetLeft")
	return offsetLeft
}

// OffsetParent prop
// js:"offsetParent"
func (*HTMLObjectElement) OffsetParent() (offsetParent window.Element) {
	macro.Rewrite("$_.offsetParent")
	return offsetParent
}

// OffsetTop prop
// js:"offsetTop"
func (*HTMLObjectElement) OffsetTop() (offsetTop int) {
	macro.Rewrite("$_.offsetTop")
	return offsetTop
}

// OffsetWidth prop
// js:"offsetWidth"
func (*HTMLObjectElement) OffsetWidth() (offsetWidth int) {
	macro.Rewrite("$_.offsetWidth")
	return offsetWidth
}

// Onabort prop
// js:"onabort"
func (*HTMLObjectElement) Onabort() (onabort func(window.Event)) {
	macro.Rewrite("$_.onabort")
	return onabort
}

// SetOnabort prop
// js:"onabort"
func (*HTMLObjectElement) SetOnabort(onabort func(window.Event)) {
	macro.Rewrite("$_.onabort = $1", onabort)
}

// Onactivate prop
// js:"onactivate"
func (*HTMLObjectElement) Onactivate() (onactivate func(window.UIEvent)) {
	macro.Rewrite("$_.onactivate")
	return onactivate
}

// SetOnactivate prop
// js:"onactivate"
func (*HTMLObjectElement) SetOnactivate(onactivate func(window.UIEvent)) {
	macro.Rewrite("$_.onactivate = $1", onactivate)
}

// Onbeforeactivate prop
// js:"onbeforeactivate"
func (*HTMLObjectElement) Onbeforeactivate() (onbeforeactivate func(window.UIEvent)) {
	macro.Rewrite("$_.onbeforeactivate")
	return onbeforeactivate
}

// SetOnbeforeactivate prop
// js:"onbeforeactivate"
func (*HTMLObjectElement) SetOnbeforeactivate(onbeforeactivate func(window.UIEvent)) {
	macro.Rewrite("$_.onbeforeactivate = $1", onbeforeactivate)
}

// Onbeforecopy prop
// js:"onbeforecopy"
func (*HTMLObjectElement) Onbeforecopy() (onbeforecopy func(*window.ClipboardEvent)) {
	macro.Rewrite("$_.onbeforecopy")
	return onbeforecopy
}

// SetOnbeforecopy prop
// js:"onbeforecopy"
func (*HTMLObjectElement) SetOnbeforecopy(onbeforecopy func(*window.ClipboardEvent)) {
	macro.Rewrite("$_.onbeforecopy = $1", onbeforecopy)
}

// Onbeforecut prop
// js:"onbeforecut"
func (*HTMLObjectElement) Onbeforecut() (onbeforecut func(*window.ClipboardEvent)) {
	macro.Rewrite("$_.onbeforecut")
	return onbeforecut
}

// SetOnbeforecut prop
// js:"onbeforecut"
func (*HTMLObjectElement) SetOnbeforecut(onbeforecut func(*window.ClipboardEvent)) {
	macro.Rewrite("$_.onbeforecut = $1", onbeforecut)
}

// Onbeforedeactivate prop
// js:"onbeforedeactivate"
func (*HTMLObjectElement) Onbeforedeactivate() (onbeforedeactivate func(window.UIEvent)) {
	macro.Rewrite("$_.onbeforedeactivate")
	return onbeforedeactivate
}

// SetOnbeforedeactivate prop
// js:"onbeforedeactivate"
func (*HTMLObjectElement) SetOnbeforedeactivate(onbeforedeactivate func(window.UIEvent)) {
	macro.Rewrite("$_.onbeforedeactivate = $1", onbeforedeactivate)
}

// Onbeforepaste prop
// js:"onbeforepaste"
func (*HTMLObjectElement) Onbeforepaste() (onbeforepaste func(*window.ClipboardEvent)) {
	macro.Rewrite("$_.onbeforepaste")
	return onbeforepaste
}

// SetOnbeforepaste prop
// js:"onbeforepaste"
func (*HTMLObjectElement) SetOnbeforepaste(onbeforepaste func(*window.ClipboardEvent)) {
	macro.Rewrite("$_.onbeforepaste = $1", onbeforepaste)
}

// Onblur prop
// js:"onblur"
func (*HTMLObjectElement) Onblur() (onblur func(*window.FocusEvent)) {
	macro.Rewrite("$_.onblur")
	return onblur
}

// SetOnblur prop
// js:"onblur"
func (*HTMLObjectElement) SetOnblur(onblur func(*window.FocusEvent)) {
	macro.Rewrite("$_.onblur = $1", onblur)
}

// Oncanplay prop
// js:"oncanplay"
func (*HTMLObjectElement) Oncanplay() (oncanplay func(window.Event)) {
	macro.Rewrite("$_.oncanplay")
	return oncanplay
}

// SetOncanplay prop
// js:"oncanplay"
func (*HTMLObjectElement) SetOncanplay(oncanplay func(window.Event)) {
	macro.Rewrite("$_.oncanplay = $1", oncanplay)
}

// Oncanplaythrough prop
// js:"oncanplaythrough"
func (*HTMLObjectElement) Oncanplaythrough() (oncanplaythrough func(window.Event)) {
	macro.Rewrite("$_.oncanplaythrough")
	return oncanplaythrough
}

// SetOncanplaythrough prop
// js:"oncanplaythrough"
func (*HTMLObjectElement) SetOncanplaythrough(oncanplaythrough func(window.Event)) {
	macro.Rewrite("$_.oncanplaythrough = $1", oncanplaythrough)
}

// Onchange prop
// js:"onchange"
func (*HTMLObjectElement) Onchange() (onchange func(window.Event)) {
	macro.Rewrite("$_.onchange")
	return onchange
}

// SetOnchange prop
// js:"onchange"
func (*HTMLObjectElement) SetOnchange(onchange func(window.Event)) {
	macro.Rewrite("$_.onchange = $1", onchange)
}

// Onclick prop
// js:"onclick"
func (*HTMLObjectElement) Onclick() (onclick func(window.MouseEvent)) {
	macro.Rewrite("$_.onclick")
	return onclick
}

// SetOnclick prop
// js:"onclick"
func (*HTMLObjectElement) SetOnclick(onclick func(window.MouseEvent)) {
	macro.Rewrite("$_.onclick = $1", onclick)
}

// Oncontextmenu prop
// js:"oncontextmenu"
func (*HTMLObjectElement) Oncontextmenu() (oncontextmenu func(*window.PointerEvent)) {
	macro.Rewrite("$_.oncontextmenu")
	return oncontextmenu
}

// SetOncontextmenu prop
// js:"oncontextmenu"
func (*HTMLObjectElement) SetOncontextmenu(oncontextmenu func(*window.PointerEvent)) {
	macro.Rewrite("$_.oncontextmenu = $1", oncontextmenu)
}

// Oncopy prop
// js:"oncopy"
func (*HTMLObjectElement) Oncopy() (oncopy func(*window.ClipboardEvent)) {
	macro.Rewrite("$_.oncopy")
	return oncopy
}

// SetOncopy prop
// js:"oncopy"
func (*HTMLObjectElement) SetOncopy(oncopy func(*window.ClipboardEvent)) {
	macro.Rewrite("$_.oncopy = $1", oncopy)
}

// Oncuechange prop
// js:"oncuechange"
func (*HTMLObjectElement) Oncuechange() (oncuechange func(window.Event)) {
	macro.Rewrite("$_.oncuechange")
	return oncuechange
}

// SetOncuechange prop
// js:"oncuechange"
func (*HTMLObjectElement) SetOncuechange(oncuechange func(window.Event)) {
	macro.Rewrite("$_.oncuechange = $1", oncuechange)
}

// Oncut prop
// js:"oncut"
func (*HTMLObjectElement) Oncut() (oncut func(*window.ClipboardEvent)) {
	macro.Rewrite("$_.oncut")
	return oncut
}

// SetOncut prop
// js:"oncut"
func (*HTMLObjectElement) SetOncut(oncut func(*window.ClipboardEvent)) {
	macro.Rewrite("$_.oncut = $1", oncut)
}

// Ondblclick prop
// js:"ondblclick"
func (*HTMLObjectElement) Ondblclick() (ondblclick func(window.MouseEvent)) {
	macro.Rewrite("$_.ondblclick")
	return ondblclick
}

// SetOndblclick prop
// js:"ondblclick"
func (*HTMLObjectElement) SetOndblclick(ondblclick func(window.MouseEvent)) {
	macro.Rewrite("$_.ondblclick = $1", ondblclick)
}

// Ondeactivate prop
// js:"ondeactivate"
func (*HTMLObjectElement) Ondeactivate() (ondeactivate func(window.UIEvent)) {
	macro.Rewrite("$_.ondeactivate")
	return ondeactivate
}

// SetOndeactivate prop
// js:"ondeactivate"
func (*HTMLObjectElement) SetOndeactivate(ondeactivate func(window.UIEvent)) {
	macro.Rewrite("$_.ondeactivate = $1", ondeactivate)
}

// Ondrag prop
// js:"ondrag"
func (*HTMLObjectElement) Ondrag() (ondrag func(*window.DragEvent)) {
	macro.Rewrite("$_.ondrag")
	return ondrag
}

// SetOndrag prop
// js:"ondrag"
func (*HTMLObjectElement) SetOndrag(ondrag func(*window.DragEvent)) {
	macro.Rewrite("$_.ondrag = $1", ondrag)
}

// Ondragend prop
// js:"ondragend"
func (*HTMLObjectElement) Ondragend() (ondragend func(*window.DragEvent)) {
	macro.Rewrite("$_.ondragend")
	return ondragend
}

// SetOndragend prop
// js:"ondragend"
func (*HTMLObjectElement) SetOndragend(ondragend func(*window.DragEvent)) {
	macro.Rewrite("$_.ondragend = $1", ondragend)
}

// Ondragenter prop
// js:"ondragenter"
func (*HTMLObjectElement) Ondragenter() (ondragenter func(*window.DragEvent)) {
	macro.Rewrite("$_.ondragenter")
	return ondragenter
}

// SetOndragenter prop
// js:"ondragenter"
func (*HTMLObjectElement) SetOndragenter(ondragenter func(*window.DragEvent)) {
	macro.Rewrite("$_.ondragenter = $1", ondragenter)
}

// Ondragleave prop
// js:"ondragleave"
func (*HTMLObjectElement) Ondragleave() (ondragleave func(*window.DragEvent)) {
	macro.Rewrite("$_.ondragleave")
	return ondragleave
}

// SetOndragleave prop
// js:"ondragleave"
func (*HTMLObjectElement) SetOndragleave(ondragleave func(*window.DragEvent)) {
	macro.Rewrite("$_.ondragleave = $1", ondragleave)
}

// Ondragover prop
// js:"ondragover"
func (*HTMLObjectElement) Ondragover() (ondragover func(*window.DragEvent)) {
	macro.Rewrite("$_.ondragover")
	return ondragover
}

// SetOndragover prop
// js:"ondragover"
func (*HTMLObjectElement) SetOndragover(ondragover func(*window.DragEvent)) {
	macro.Rewrite("$_.ondragover = $1", ondragover)
}

// Ondragstart prop
// js:"ondragstart"
func (*HTMLObjectElement) Ondragstart() (ondragstart func(*window.DragEvent)) {
	macro.Rewrite("$_.ondragstart")
	return ondragstart
}

// SetOndragstart prop
// js:"ondragstart"
func (*HTMLObjectElement) SetOndragstart(ondragstart func(*window.DragEvent)) {
	macro.Rewrite("$_.ondragstart = $1", ondragstart)
}

// Ondrop prop
// js:"ondrop"
func (*HTMLObjectElement) Ondrop() (ondrop func(*window.DragEvent)) {
	macro.Rewrite("$_.ondrop")
	return ondrop
}

// SetOndrop prop
// js:"ondrop"
func (*HTMLObjectElement) SetOndrop(ondrop func(*window.DragEvent)) {
	macro.Rewrite("$_.ondrop = $1", ondrop)
}

// Ondurationchange prop
// js:"ondurationchange"
func (*HTMLObjectElement) Ondurationchange() (ondurationchange func(window.Event)) {
	macro.Rewrite("$_.ondurationchange")
	return ondurationchange
}

// SetOndurationchange prop
// js:"ondurationchange"
func (*HTMLObjectElement) SetOndurationchange(ondurationchange func(window.Event)) {
	macro.Rewrite("$_.ondurationchange = $1", ondurationchange)
}

// Onemptied prop
// js:"onemptied"
func (*HTMLObjectElement) Onemptied() (onemptied func(window.Event)) {
	macro.Rewrite("$_.onemptied")
	return onemptied
}

// SetOnemptied prop
// js:"onemptied"
func (*HTMLObjectElement) SetOnemptied(onemptied func(window.Event)) {
	macro.Rewrite("$_.onemptied = $1", onemptied)
}

// Onended prop
// js:"onended"
func (*HTMLObjectElement) Onended() (onended func(window.Event)) {
	macro.Rewrite("$_.onended")
	return onended
}

// SetOnended prop
// js:"onended"
func (*HTMLObjectElement) SetOnended(onended func(window.Event)) {
	macro.Rewrite("$_.onended = $1", onended)
}

// Onerror prop
// js:"onerror"
func (*HTMLObjectElement) Onerror() (onerror func(window.Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

// SetOnerror prop
// js:"onerror"
func (*HTMLObjectElement) SetOnerror(onerror func(window.Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

// Onfocus prop
// js:"onfocus"
func (*HTMLObjectElement) Onfocus() (onfocus func(*window.FocusEvent)) {
	macro.Rewrite("$_.onfocus")
	return onfocus
}

// SetOnfocus prop
// js:"onfocus"
func (*HTMLObjectElement) SetOnfocus(onfocus func(*window.FocusEvent)) {
	macro.Rewrite("$_.onfocus = $1", onfocus)
}

// Oninput prop
// js:"oninput"
func (*HTMLObjectElement) Oninput() (oninput func(window.Event)) {
	macro.Rewrite("$_.oninput")
	return oninput
}

// SetOninput prop
// js:"oninput"
func (*HTMLObjectElement) SetOninput(oninput func(window.Event)) {
	macro.Rewrite("$_.oninput = $1", oninput)
}

// Oninvalid prop
// js:"oninvalid"
func (*HTMLObjectElement) Oninvalid() (oninvalid func(window.Event)) {
	macro.Rewrite("$_.oninvalid")
	return oninvalid
}

// SetOninvalid prop
// js:"oninvalid"
func (*HTMLObjectElement) SetOninvalid(oninvalid func(window.Event)) {
	macro.Rewrite("$_.oninvalid = $1", oninvalid)
}

// Onkeydown prop
// js:"onkeydown"
func (*HTMLObjectElement) Onkeydown() (onkeydown func(*window.KeyboardEvent)) {
	macro.Rewrite("$_.onkeydown")
	return onkeydown
}

// SetOnkeydown prop
// js:"onkeydown"
func (*HTMLObjectElement) SetOnkeydown(onkeydown func(*window.KeyboardEvent)) {
	macro.Rewrite("$_.onkeydown = $1", onkeydown)
}

// Onkeypress prop
// js:"onkeypress"
func (*HTMLObjectElement) Onkeypress() (onkeypress func(*window.KeyboardEvent)) {
	macro.Rewrite("$_.onkeypress")
	return onkeypress
}

// SetOnkeypress prop
// js:"onkeypress"
func (*HTMLObjectElement) SetOnkeypress(onkeypress func(*window.KeyboardEvent)) {
	macro.Rewrite("$_.onkeypress = $1", onkeypress)
}

// Onkeyup prop
// js:"onkeyup"
func (*HTMLObjectElement) Onkeyup() (onkeyup func(*window.KeyboardEvent)) {
	macro.Rewrite("$_.onkeyup")
	return onkeyup
}

// SetOnkeyup prop
// js:"onkeyup"
func (*HTMLObjectElement) SetOnkeyup(onkeyup func(*window.KeyboardEvent)) {
	macro.Rewrite("$_.onkeyup = $1", onkeyup)
}

// Onload prop
// js:"onload"
func (*HTMLObjectElement) Onload() (onload func(window.Event)) {
	macro.Rewrite("$_.onload")
	return onload
}

// SetOnload prop
// js:"onload"
func (*HTMLObjectElement) SetOnload(onload func(window.Event)) {
	macro.Rewrite("$_.onload = $1", onload)
}

// Onloadeddata prop
// js:"onloadeddata"
func (*HTMLObjectElement) Onloadeddata() (onloadeddata func(window.Event)) {
	macro.Rewrite("$_.onloadeddata")
	return onloadeddata
}

// SetOnloadeddata prop
// js:"onloadeddata"
func (*HTMLObjectElement) SetOnloadeddata(onloadeddata func(window.Event)) {
	macro.Rewrite("$_.onloadeddata = $1", onloadeddata)
}

// Onloadedmetadata prop
// js:"onloadedmetadata"
func (*HTMLObjectElement) Onloadedmetadata() (onloadedmetadata func(window.Event)) {
	macro.Rewrite("$_.onloadedmetadata")
	return onloadedmetadata
}

// SetOnloadedmetadata prop
// js:"onloadedmetadata"
func (*HTMLObjectElement) SetOnloadedmetadata(onloadedmetadata func(window.Event)) {
	macro.Rewrite("$_.onloadedmetadata = $1", onloadedmetadata)
}

// Onloadstart prop
// js:"onloadstart"
func (*HTMLObjectElement) Onloadstart() (onloadstart func(window.Event)) {
	macro.Rewrite("$_.onloadstart")
	return onloadstart
}

// SetOnloadstart prop
// js:"onloadstart"
func (*HTMLObjectElement) SetOnloadstart(onloadstart func(window.Event)) {
	macro.Rewrite("$_.onloadstart = $1", onloadstart)
}

// Onmousedown prop
// js:"onmousedown"
func (*HTMLObjectElement) Onmousedown() (onmousedown func(window.MouseEvent)) {
	macro.Rewrite("$_.onmousedown")
	return onmousedown
}

// SetOnmousedown prop
// js:"onmousedown"
func (*HTMLObjectElement) SetOnmousedown(onmousedown func(window.MouseEvent)) {
	macro.Rewrite("$_.onmousedown = $1", onmousedown)
}

// Onmouseenter prop
// js:"onmouseenter"
func (*HTMLObjectElement) Onmouseenter() (onmouseenter func(window.MouseEvent)) {
	macro.Rewrite("$_.onmouseenter")
	return onmouseenter
}

// SetOnmouseenter prop
// js:"onmouseenter"
func (*HTMLObjectElement) SetOnmouseenter(onmouseenter func(window.MouseEvent)) {
	macro.Rewrite("$_.onmouseenter = $1", onmouseenter)
}

// Onmouseleave prop
// js:"onmouseleave"
func (*HTMLObjectElement) Onmouseleave() (onmouseleave func(window.MouseEvent)) {
	macro.Rewrite("$_.onmouseleave")
	return onmouseleave
}

// SetOnmouseleave prop
// js:"onmouseleave"
func (*HTMLObjectElement) SetOnmouseleave(onmouseleave func(window.MouseEvent)) {
	macro.Rewrite("$_.onmouseleave = $1", onmouseleave)
}

// Onmousemove prop
// js:"onmousemove"
func (*HTMLObjectElement) Onmousemove() (onmousemove func(window.MouseEvent)) {
	macro.Rewrite("$_.onmousemove")
	return onmousemove
}

// SetOnmousemove prop
// js:"onmousemove"
func (*HTMLObjectElement) SetOnmousemove(onmousemove func(window.MouseEvent)) {
	macro.Rewrite("$_.onmousemove = $1", onmousemove)
}

// Onmouseout prop
// js:"onmouseout"
func (*HTMLObjectElement) Onmouseout() (onmouseout func(window.MouseEvent)) {
	macro.Rewrite("$_.onmouseout")
	return onmouseout
}

// SetOnmouseout prop
// js:"onmouseout"
func (*HTMLObjectElement) SetOnmouseout(onmouseout func(window.MouseEvent)) {
	macro.Rewrite("$_.onmouseout = $1", onmouseout)
}

// Onmouseover prop
// js:"onmouseover"
func (*HTMLObjectElement) Onmouseover() (onmouseover func(window.MouseEvent)) {
	macro.Rewrite("$_.onmouseover")
	return onmouseover
}

// SetOnmouseover prop
// js:"onmouseover"
func (*HTMLObjectElement) SetOnmouseover(onmouseover func(window.MouseEvent)) {
	macro.Rewrite("$_.onmouseover = $1", onmouseover)
}

// Onmouseup prop
// js:"onmouseup"
func (*HTMLObjectElement) Onmouseup() (onmouseup func(window.MouseEvent)) {
	macro.Rewrite("$_.onmouseup")
	return onmouseup
}

// SetOnmouseup prop
// js:"onmouseup"
func (*HTMLObjectElement) SetOnmouseup(onmouseup func(window.MouseEvent)) {
	macro.Rewrite("$_.onmouseup = $1", onmouseup)
}

// Onmousewheel prop
// js:"onmousewheel"
func (*HTMLObjectElement) Onmousewheel() (onmousewheel func(*window.WheelEvent)) {
	macro.Rewrite("$_.onmousewheel")
	return onmousewheel
}

// SetOnmousewheel prop
// js:"onmousewheel"
func (*HTMLObjectElement) SetOnmousewheel(onmousewheel func(*window.WheelEvent)) {
	macro.Rewrite("$_.onmousewheel = $1", onmousewheel)
}

// Onmscontentzoom prop
// js:"onmscontentzoom"
func (*HTMLObjectElement) Onmscontentzoom() (onmscontentzoom func(window.UIEvent)) {
	macro.Rewrite("$_.onmscontentzoom")
	return onmscontentzoom
}

// SetOnmscontentzoom prop
// js:"onmscontentzoom"
func (*HTMLObjectElement) SetOnmscontentzoom(onmscontentzoom func(window.UIEvent)) {
	macro.Rewrite("$_.onmscontentzoom = $1", onmscontentzoom)
}

// Onmsmanipulationstatechanged prop
// js:"onmsmanipulationstatechanged"
func (*HTMLObjectElement) Onmsmanipulationstatechanged() (onmsmanipulationstatechanged func(*window.MSManipulationEvent)) {
	macro.Rewrite("$_.onmsmanipulationstatechanged")
	return onmsmanipulationstatechanged
}

// SetOnmsmanipulationstatechanged prop
// js:"onmsmanipulationstatechanged"
func (*HTMLObjectElement) SetOnmsmanipulationstatechanged(onmsmanipulationstatechanged func(*window.MSManipulationEvent)) {
	macro.Rewrite("$_.onmsmanipulationstatechanged = $1", onmsmanipulationstatechanged)
}

// Onpaste prop
// js:"onpaste"
func (*HTMLObjectElement) Onpaste() (onpaste func(*window.ClipboardEvent)) {
	macro.Rewrite("$_.onpaste")
	return onpaste
}

// SetOnpaste prop
// js:"onpaste"
func (*HTMLObjectElement) SetOnpaste(onpaste func(*window.ClipboardEvent)) {
	macro.Rewrite("$_.onpaste = $1", onpaste)
}

// Onpause prop
// js:"onpause"
func (*HTMLObjectElement) Onpause() (onpause func(window.Event)) {
	macro.Rewrite("$_.onpause")
	return onpause
}

// SetOnpause prop
// js:"onpause"
func (*HTMLObjectElement) SetOnpause(onpause func(window.Event)) {
	macro.Rewrite("$_.onpause = $1", onpause)
}

// Onplay prop
// js:"onplay"
func (*HTMLObjectElement) Onplay() (onplay func(window.Event)) {
	macro.Rewrite("$_.onplay")
	return onplay
}

// SetOnplay prop
// js:"onplay"
func (*HTMLObjectElement) SetOnplay(onplay func(window.Event)) {
	macro.Rewrite("$_.onplay = $1", onplay)
}

// Onplaying prop
// js:"onplaying"
func (*HTMLObjectElement) Onplaying() (onplaying func(window.Event)) {
	macro.Rewrite("$_.onplaying")
	return onplaying
}

// SetOnplaying prop
// js:"onplaying"
func (*HTMLObjectElement) SetOnplaying(onplaying func(window.Event)) {
	macro.Rewrite("$_.onplaying = $1", onplaying)
}

// Onprogress prop
// js:"onprogress"
func (*HTMLObjectElement) Onprogress() (onprogress func(window.Event)) {
	macro.Rewrite("$_.onprogress")
	return onprogress
}

// SetOnprogress prop
// js:"onprogress"
func (*HTMLObjectElement) SetOnprogress(onprogress func(window.Event)) {
	macro.Rewrite("$_.onprogress = $1", onprogress)
}

// Onratechange prop
// js:"onratechange"
func (*HTMLObjectElement) Onratechange() (onratechange func(window.Event)) {
	macro.Rewrite("$_.onratechange")
	return onratechange
}

// SetOnratechange prop
// js:"onratechange"
func (*HTMLObjectElement) SetOnratechange(onratechange func(window.Event)) {
	macro.Rewrite("$_.onratechange = $1", onratechange)
}

// Onreset prop
// js:"onreset"
func (*HTMLObjectElement) Onreset() (onreset func(window.Event)) {
	macro.Rewrite("$_.onreset")
	return onreset
}

// SetOnreset prop
// js:"onreset"
func (*HTMLObjectElement) SetOnreset(onreset func(window.Event)) {
	macro.Rewrite("$_.onreset = $1", onreset)
}

// Onscroll prop
// js:"onscroll"
func (*HTMLObjectElement) Onscroll() (onscroll func(window.UIEvent)) {
	macro.Rewrite("$_.onscroll")
	return onscroll
}

// SetOnscroll prop
// js:"onscroll"
func (*HTMLObjectElement) SetOnscroll(onscroll func(window.UIEvent)) {
	macro.Rewrite("$_.onscroll = $1", onscroll)
}

// Onseeked prop
// js:"onseeked"
func (*HTMLObjectElement) Onseeked() (onseeked func(window.Event)) {
	macro.Rewrite("$_.onseeked")
	return onseeked
}

// SetOnseeked prop
// js:"onseeked"
func (*HTMLObjectElement) SetOnseeked(onseeked func(window.Event)) {
	macro.Rewrite("$_.onseeked = $1", onseeked)
}

// Onseeking prop
// js:"onseeking"
func (*HTMLObjectElement) Onseeking() (onseeking func(window.Event)) {
	macro.Rewrite("$_.onseeking")
	return onseeking
}

// SetOnseeking prop
// js:"onseeking"
func (*HTMLObjectElement) SetOnseeking(onseeking func(window.Event)) {
	macro.Rewrite("$_.onseeking = $1", onseeking)
}

// Onselect prop
// js:"onselect"
func (*HTMLObjectElement) Onselect() (onselect func(window.UIEvent)) {
	macro.Rewrite("$_.onselect")
	return onselect
}

// SetOnselect prop
// js:"onselect"
func (*HTMLObjectElement) SetOnselect(onselect func(window.UIEvent)) {
	macro.Rewrite("$_.onselect = $1", onselect)
}

// Onselectstart prop
// js:"onselectstart"
func (*HTMLObjectElement) Onselectstart() (onselectstart func(window.Event)) {
	macro.Rewrite("$_.onselectstart")
	return onselectstart
}

// SetOnselectstart prop
// js:"onselectstart"
func (*HTMLObjectElement) SetOnselectstart(onselectstart func(window.Event)) {
	macro.Rewrite("$_.onselectstart = $1", onselectstart)
}

// Onstalled prop
// js:"onstalled"
func (*HTMLObjectElement) Onstalled() (onstalled func(window.Event)) {
	macro.Rewrite("$_.onstalled")
	return onstalled
}

// SetOnstalled prop
// js:"onstalled"
func (*HTMLObjectElement) SetOnstalled(onstalled func(window.Event)) {
	macro.Rewrite("$_.onstalled = $1", onstalled)
}

// Onsubmit prop
// js:"onsubmit"
func (*HTMLObjectElement) Onsubmit() (onsubmit func(window.Event)) {
	macro.Rewrite("$_.onsubmit")
	return onsubmit
}

// SetOnsubmit prop
// js:"onsubmit"
func (*HTMLObjectElement) SetOnsubmit(onsubmit func(window.Event)) {
	macro.Rewrite("$_.onsubmit = $1", onsubmit)
}

// Onsuspend prop
// js:"onsuspend"
func (*HTMLObjectElement) Onsuspend() (onsuspend func(window.Event)) {
	macro.Rewrite("$_.onsuspend")
	return onsuspend
}

// SetOnsuspend prop
// js:"onsuspend"
func (*HTMLObjectElement) SetOnsuspend(onsuspend func(window.Event)) {
	macro.Rewrite("$_.onsuspend = $1", onsuspend)
}

// Ontimeupdate prop
// js:"ontimeupdate"
func (*HTMLObjectElement) Ontimeupdate() (ontimeupdate func(window.Event)) {
	macro.Rewrite("$_.ontimeupdate")
	return ontimeupdate
}

// SetOntimeupdate prop
// js:"ontimeupdate"
func (*HTMLObjectElement) SetOntimeupdate(ontimeupdate func(window.Event)) {
	macro.Rewrite("$_.ontimeupdate = $1", ontimeupdate)
}

// Onvolumechange prop
// js:"onvolumechange"
func (*HTMLObjectElement) Onvolumechange() (onvolumechange func(window.Event)) {
	macro.Rewrite("$_.onvolumechange")
	return onvolumechange
}

// SetOnvolumechange prop
// js:"onvolumechange"
func (*HTMLObjectElement) SetOnvolumechange(onvolumechange func(window.Event)) {
	macro.Rewrite("$_.onvolumechange = $1", onvolumechange)
}

// Onwaiting prop
// js:"onwaiting"
func (*HTMLObjectElement) Onwaiting() (onwaiting func(window.Event)) {
	macro.Rewrite("$_.onwaiting")
	return onwaiting
}

// SetOnwaiting prop
// js:"onwaiting"
func (*HTMLObjectElement) SetOnwaiting(onwaiting func(window.Event)) {
	macro.Rewrite("$_.onwaiting = $1", onwaiting)
}

// OuterText prop
// js:"outerText"
func (*HTMLObjectElement) OuterText() (outerText string) {
	macro.Rewrite("$_.outerText")
	return outerText
}

// SetOuterText prop
// js:"outerText"
func (*HTMLObjectElement) SetOuterText(outerText string) {
	macro.Rewrite("$_.outerText = $1", outerText)
}

// Spellcheck prop
// js:"spellcheck"
func (*HTMLObjectElement) Spellcheck() (spellcheck bool) {
	macro.Rewrite("$_.spellcheck")
	return spellcheck
}

// SetSpellcheck prop
// js:"spellcheck"
func (*HTMLObjectElement) SetSpellcheck(spellcheck bool) {
	macro.Rewrite("$_.spellcheck = $1", spellcheck)
}

// Style prop
// js:"style"
func (*HTMLObjectElement) Style() (style *window.CSSStyleDeclaration) {
	macro.Rewrite("$_.style")
	return style
}

// TabIndex prop
// js:"tabIndex"
func (*HTMLObjectElement) TabIndex() (tabIndex int8) {
	macro.Rewrite("$_.tabIndex")
	return tabIndex
}

// SetTabIndex prop
// js:"tabIndex"
func (*HTMLObjectElement) SetTabIndex(tabIndex int8) {
	macro.Rewrite("$_.tabIndex = $1", tabIndex)
}

// Title prop
// js:"title"
func (*HTMLObjectElement) Title() (title string) {
	macro.Rewrite("$_.title")
	return title
}

// SetTitle prop
// js:"title"
func (*HTMLObjectElement) SetTitle(title string) {
	macro.Rewrite("$_.title = $1", title)
}

// ClassList prop
// js:"classList"
func (*HTMLObjectElement) ClassList() (classList domtokenlist.DOMTokenList) {
	macro.Rewrite("$_.classList")
	return classList
}

// ClassName prop
// js:"className"
func (*HTMLObjectElement) ClassName() (className string) {
	macro.Rewrite("$_.className")
	return className
}

// SetClassName prop
// js:"className"
func (*HTMLObjectElement) SetClassName(className string) {
	macro.Rewrite("$_.className = $1", className)
}

// ClientHeight prop
// js:"clientHeight"
func (*HTMLObjectElement) ClientHeight() (clientHeight int) {
	macro.Rewrite("$_.clientHeight")
	return clientHeight
}

// ClientLeft prop
// js:"clientLeft"
func (*HTMLObjectElement) ClientLeft() (clientLeft int) {
	macro.Rewrite("$_.clientLeft")
	return clientLeft
}

// ClientTop prop
// js:"clientTop"
func (*HTMLObjectElement) ClientTop() (clientTop int) {
	macro.Rewrite("$_.clientTop")
	return clientTop
}

// ClientWidth prop
// js:"clientWidth"
func (*HTMLObjectElement) ClientWidth() (clientWidth int) {
	macro.Rewrite("$_.clientWidth")
	return clientWidth
}

// ID prop
// js:"id"
func (*HTMLObjectElement) ID() (id string) {
	macro.Rewrite("$_.id")
	return id
}

// SetID prop
// js:"id"
func (*HTMLObjectElement) SetID(id string) {
	macro.Rewrite("$_.id = $1", id)
}

// InnerHTML prop
// js:"innerHTML"
func (*HTMLObjectElement) InnerHTML() (innerHTML string) {
	macro.Rewrite("$_.innerHTML")
	return innerHTML
}

// SetInnerHTML prop
// js:"innerHTML"
func (*HTMLObjectElement) SetInnerHTML(innerHTML string) {
	macro.Rewrite("$_.innerHTML = $1", innerHTML)
}

// MsContentZoomFactor prop
// js:"msContentZoomFactor"
func (*HTMLObjectElement) MsContentZoomFactor() (msContentZoomFactor float32) {
	macro.Rewrite("$_.msContentZoomFactor")
	return msContentZoomFactor
}

// SetMsContentZoomFactor prop
// js:"msContentZoomFactor"
func (*HTMLObjectElement) SetMsContentZoomFactor(msContentZoomFactor float32) {
	macro.Rewrite("$_.msContentZoomFactor = $1", msContentZoomFactor)
}

// MsRegionOverflow prop
// js:"msRegionOverflow"
func (*HTMLObjectElement) MsRegionOverflow() (msRegionOverflow string) {
	macro.Rewrite("$_.msRegionOverflow")
	return msRegionOverflow
}

// Onariarequest prop
// js:"onariarequest"
func (*HTMLObjectElement) Onariarequest() (onariarequest func(window.Event)) {
	macro.Rewrite("$_.onariarequest")
	return onariarequest
}

// SetOnariarequest prop
// js:"onariarequest"
func (*HTMLObjectElement) SetOnariarequest(onariarequest func(window.Event)) {
	macro.Rewrite("$_.onariarequest = $1", onariarequest)
}

// Oncommand prop
// js:"oncommand"
func (*HTMLObjectElement) Oncommand() (oncommand func(window.Event)) {
	macro.Rewrite("$_.oncommand")
	return oncommand
}

// SetOncommand prop
// js:"oncommand"
func (*HTMLObjectElement) SetOncommand(oncommand func(window.Event)) {
	macro.Rewrite("$_.oncommand = $1", oncommand)
}

// Ongotpointercapture prop
// js:"ongotpointercapture"
func (*HTMLObjectElement) Ongotpointercapture() (ongotpointercapture func(*window.PointerEvent)) {
	macro.Rewrite("$_.ongotpointercapture")
	return ongotpointercapture
}

// SetOngotpointercapture prop
// js:"ongotpointercapture"
func (*HTMLObjectElement) SetOngotpointercapture(ongotpointercapture func(*window.PointerEvent)) {
	macro.Rewrite("$_.ongotpointercapture = $1", ongotpointercapture)
}

// Onlostpointercapture prop
// js:"onlostpointercapture"
func (*HTMLObjectElement) Onlostpointercapture() (onlostpointercapture func(*window.PointerEvent)) {
	macro.Rewrite("$_.onlostpointercapture")
	return onlostpointercapture
}

// SetOnlostpointercapture prop
// js:"onlostpointercapture"
func (*HTMLObjectElement) SetOnlostpointercapture(onlostpointercapture func(*window.PointerEvent)) {
	macro.Rewrite("$_.onlostpointercapture = $1", onlostpointercapture)
}

// Onmsgesturechange prop
// js:"onmsgesturechange"
func (*HTMLObjectElement) Onmsgesturechange() (onmsgesturechange func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturechange")
	return onmsgesturechange
}

// SetOnmsgesturechange prop
// js:"onmsgesturechange"
func (*HTMLObjectElement) SetOnmsgesturechange(onmsgesturechange func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturechange = $1", onmsgesturechange)
}

// Onmsgesturedoubletap prop
// js:"onmsgesturedoubletap"
func (*HTMLObjectElement) Onmsgesturedoubletap() (onmsgesturedoubletap func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturedoubletap")
	return onmsgesturedoubletap
}

// SetOnmsgesturedoubletap prop
// js:"onmsgesturedoubletap"
func (*HTMLObjectElement) SetOnmsgesturedoubletap(onmsgesturedoubletap func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturedoubletap = $1", onmsgesturedoubletap)
}

// Onmsgestureend prop
// js:"onmsgestureend"
func (*HTMLObjectElement) Onmsgestureend() (onmsgestureend func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgestureend")
	return onmsgestureend
}

// SetOnmsgestureend prop
// js:"onmsgestureend"
func (*HTMLObjectElement) SetOnmsgestureend(onmsgestureend func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgestureend = $1", onmsgestureend)
}

// Onmsgesturehold prop
// js:"onmsgesturehold"
func (*HTMLObjectElement) Onmsgesturehold() (onmsgesturehold func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturehold")
	return onmsgesturehold
}

// SetOnmsgesturehold prop
// js:"onmsgesturehold"
func (*HTMLObjectElement) SetOnmsgesturehold(onmsgesturehold func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturehold = $1", onmsgesturehold)
}

// Onmsgesturestart prop
// js:"onmsgesturestart"
func (*HTMLObjectElement) Onmsgesturestart() (onmsgesturestart func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturestart")
	return onmsgesturestart
}

// SetOnmsgesturestart prop
// js:"onmsgesturestart"
func (*HTMLObjectElement) SetOnmsgesturestart(onmsgesturestart func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturestart = $1", onmsgesturestart)
}

// Onmsgesturetap prop
// js:"onmsgesturetap"
func (*HTMLObjectElement) Onmsgesturetap() (onmsgesturetap func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturetap")
	return onmsgesturetap
}

// SetOnmsgesturetap prop
// js:"onmsgesturetap"
func (*HTMLObjectElement) SetOnmsgesturetap(onmsgesturetap func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturetap = $1", onmsgesturetap)
}

// Onmsgotpointercapture prop
// js:"onmsgotpointercapture"
func (*HTMLObjectElement) Onmsgotpointercapture() (onmsgotpointercapture func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmsgotpointercapture")
	return onmsgotpointercapture
}

// SetOnmsgotpointercapture prop
// js:"onmsgotpointercapture"
func (*HTMLObjectElement) SetOnmsgotpointercapture(onmsgotpointercapture func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmsgotpointercapture = $1", onmsgotpointercapture)
}

// Onmsinertiastart prop
// js:"onmsinertiastart"
func (*HTMLObjectElement) Onmsinertiastart() (onmsinertiastart func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsinertiastart")
	return onmsinertiastart
}

// SetOnmsinertiastart prop
// js:"onmsinertiastart"
func (*HTMLObjectElement) SetOnmsinertiastart(onmsinertiastart func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsinertiastart = $1", onmsinertiastart)
}

// Onmslostpointercapture prop
// js:"onmslostpointercapture"
func (*HTMLObjectElement) Onmslostpointercapture() (onmslostpointercapture func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmslostpointercapture")
	return onmslostpointercapture
}

// SetOnmslostpointercapture prop
// js:"onmslostpointercapture"
func (*HTMLObjectElement) SetOnmslostpointercapture(onmslostpointercapture func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmslostpointercapture = $1", onmslostpointercapture)
}

// Onmspointercancel prop
// js:"onmspointercancel"
func (*HTMLObjectElement) Onmspointercancel() (onmspointercancel func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointercancel")
	return onmspointercancel
}

// SetOnmspointercancel prop
// js:"onmspointercancel"
func (*HTMLObjectElement) SetOnmspointercancel(onmspointercancel func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointercancel = $1", onmspointercancel)
}

// Onmspointerdown prop
// js:"onmspointerdown"
func (*HTMLObjectElement) Onmspointerdown() (onmspointerdown func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerdown")
	return onmspointerdown
}

// SetOnmspointerdown prop
// js:"onmspointerdown"
func (*HTMLObjectElement) SetOnmspointerdown(onmspointerdown func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerdown = $1", onmspointerdown)
}

// Onmspointerenter prop
// js:"onmspointerenter"
func (*HTMLObjectElement) Onmspointerenter() (onmspointerenter func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerenter")
	return onmspointerenter
}

// SetOnmspointerenter prop
// js:"onmspointerenter"
func (*HTMLObjectElement) SetOnmspointerenter(onmspointerenter func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerenter = $1", onmspointerenter)
}

// Onmspointerleave prop
// js:"onmspointerleave"
func (*HTMLObjectElement) Onmspointerleave() (onmspointerleave func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerleave")
	return onmspointerleave
}

// SetOnmspointerleave prop
// js:"onmspointerleave"
func (*HTMLObjectElement) SetOnmspointerleave(onmspointerleave func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerleave = $1", onmspointerleave)
}

// Onmspointermove prop
// js:"onmspointermove"
func (*HTMLObjectElement) Onmspointermove() (onmspointermove func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointermove")
	return onmspointermove
}

// SetOnmspointermove prop
// js:"onmspointermove"
func (*HTMLObjectElement) SetOnmspointermove(onmspointermove func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointermove = $1", onmspointermove)
}

// Onmspointerout prop
// js:"onmspointerout"
func (*HTMLObjectElement) Onmspointerout() (onmspointerout func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerout")
	return onmspointerout
}

// SetOnmspointerout prop
// js:"onmspointerout"
func (*HTMLObjectElement) SetOnmspointerout(onmspointerout func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerout = $1", onmspointerout)
}

// Onmspointerover prop
// js:"onmspointerover"
func (*HTMLObjectElement) Onmspointerover() (onmspointerover func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerover")
	return onmspointerover
}

// SetOnmspointerover prop
// js:"onmspointerover"
func (*HTMLObjectElement) SetOnmspointerover(onmspointerover func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerover = $1", onmspointerover)
}

// Onmspointerup prop
// js:"onmspointerup"
func (*HTMLObjectElement) Onmspointerup() (onmspointerup func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerup")
	return onmspointerup
}

// SetOnmspointerup prop
// js:"onmspointerup"
func (*HTMLObjectElement) SetOnmspointerup(onmspointerup func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerup = $1", onmspointerup)
}

// Ontouchcancel prop
// js:"ontouchcancel"
func (*HTMLObjectElement) Ontouchcancel() (ontouchcancel func(*window.TouchEvent)) {
	macro.Rewrite("$_.ontouchcancel")
	return ontouchcancel
}

// SetOntouchcancel prop
// js:"ontouchcancel"
func (*HTMLObjectElement) SetOntouchcancel(ontouchcancel func(*window.TouchEvent)) {
	macro.Rewrite("$_.ontouchcancel = $1", ontouchcancel)
}

// Ontouchend prop
// js:"ontouchend"
func (*HTMLObjectElement) Ontouchend() (ontouchend func(*window.TouchEvent)) {
	macro.Rewrite("$_.ontouchend")
	return ontouchend
}

// SetOntouchend prop
// js:"ontouchend"
func (*HTMLObjectElement) SetOntouchend(ontouchend func(*window.TouchEvent)) {
	macro.Rewrite("$_.ontouchend = $1", ontouchend)
}

// Ontouchmove prop
// js:"ontouchmove"
func (*HTMLObjectElement) Ontouchmove() (ontouchmove func(*window.TouchEvent)) {
	macro.Rewrite("$_.ontouchmove")
	return ontouchmove
}

// SetOntouchmove prop
// js:"ontouchmove"
func (*HTMLObjectElement) SetOntouchmove(ontouchmove func(*window.TouchEvent)) {
	macro.Rewrite("$_.ontouchmove = $1", ontouchmove)
}

// Ontouchstart prop
// js:"ontouchstart"
func (*HTMLObjectElement) Ontouchstart() (ontouchstart func(*window.TouchEvent)) {
	macro.Rewrite("$_.ontouchstart")
	return ontouchstart
}

// SetOntouchstart prop
// js:"ontouchstart"
func (*HTMLObjectElement) SetOntouchstart(ontouchstart func(*window.TouchEvent)) {
	macro.Rewrite("$_.ontouchstart = $1", ontouchstart)
}

// Onwebkitfullscreenchange prop
// js:"onwebkitfullscreenchange"
func (*HTMLObjectElement) Onwebkitfullscreenchange() (onwebkitfullscreenchange func(window.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenchange")
	return onwebkitfullscreenchange
}

// SetOnwebkitfullscreenchange prop
// js:"onwebkitfullscreenchange"
func (*HTMLObjectElement) SetOnwebkitfullscreenchange(onwebkitfullscreenchange func(window.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenchange = $1", onwebkitfullscreenchange)
}

// Onwebkitfullscreenerror prop
// js:"onwebkitfullscreenerror"
func (*HTMLObjectElement) Onwebkitfullscreenerror() (onwebkitfullscreenerror func(window.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenerror")
	return onwebkitfullscreenerror
}

// SetOnwebkitfullscreenerror prop
// js:"onwebkitfullscreenerror"
func (*HTMLObjectElement) SetOnwebkitfullscreenerror(onwebkitfullscreenerror func(window.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenerror = $1", onwebkitfullscreenerror)
}

// OuterHTML prop
// js:"outerHTML"
func (*HTMLObjectElement) OuterHTML() (outerHTML string) {
	macro.Rewrite("$_.outerHTML")
	return outerHTML
}

// SetOuterHTML prop
// js:"outerHTML"
func (*HTMLObjectElement) SetOuterHTML(outerHTML string) {
	macro.Rewrite("$_.outerHTML = $1", outerHTML)
}

// Prefix prop
// js:"prefix"
func (*HTMLObjectElement) Prefix() (prefix string) {
	macro.Rewrite("$_.prefix")
	return prefix
}

// ScrollHeight prop
// js:"scrollHeight"
func (*HTMLObjectElement) ScrollHeight() (scrollHeight int) {
	macro.Rewrite("$_.scrollHeight")
	return scrollHeight
}

// ScrollLeft prop
// js:"scrollLeft"
func (*HTMLObjectElement) ScrollLeft() (scrollLeft int) {
	macro.Rewrite("$_.scrollLeft")
	return scrollLeft
}

// SetScrollLeft prop
// js:"scrollLeft"
func (*HTMLObjectElement) SetScrollLeft(scrollLeft int) {
	macro.Rewrite("$_.scrollLeft = $1", scrollLeft)
}

// ScrollTop prop
// js:"scrollTop"
func (*HTMLObjectElement) ScrollTop() (scrollTop int) {
	macro.Rewrite("$_.scrollTop")
	return scrollTop
}

// SetScrollTop prop
// js:"scrollTop"
func (*HTMLObjectElement) SetScrollTop(scrollTop int) {
	macro.Rewrite("$_.scrollTop = $1", scrollTop)
}

// ScrollWidth prop
// js:"scrollWidth"
func (*HTMLObjectElement) ScrollWidth() (scrollWidth int) {
	macro.Rewrite("$_.scrollWidth")
	return scrollWidth
}

// TagName prop
// js:"tagName"
func (*HTMLObjectElement) TagName() (tagName string) {
	macro.Rewrite("$_.tagName")
	return tagName
}

// Onpointercancel prop
// js:"onpointercancel"
func (*HTMLObjectElement) Onpointercancel() (onpointercancel func(window.Event)) {
	macro.Rewrite("$_.onpointercancel")
	return onpointercancel
}

// SetOnpointercancel prop
// js:"onpointercancel"
func (*HTMLObjectElement) SetOnpointercancel(onpointercancel func(window.Event)) {
	macro.Rewrite("$_.onpointercancel = $1", onpointercancel)
}

// Onpointerdown prop
// js:"onpointerdown"
func (*HTMLObjectElement) Onpointerdown() (onpointerdown func(window.Event)) {
	macro.Rewrite("$_.onpointerdown")
	return onpointerdown
}

// SetOnpointerdown prop
// js:"onpointerdown"
func (*HTMLObjectElement) SetOnpointerdown(onpointerdown func(window.Event)) {
	macro.Rewrite("$_.onpointerdown = $1", onpointerdown)
}

// Onpointerenter prop
// js:"onpointerenter"
func (*HTMLObjectElement) Onpointerenter() (onpointerenter func(window.Event)) {
	macro.Rewrite("$_.onpointerenter")
	return onpointerenter
}

// SetOnpointerenter prop
// js:"onpointerenter"
func (*HTMLObjectElement) SetOnpointerenter(onpointerenter func(window.Event)) {
	macro.Rewrite("$_.onpointerenter = $1", onpointerenter)
}

// Onpointerleave prop
// js:"onpointerleave"
func (*HTMLObjectElement) Onpointerleave() (onpointerleave func(window.Event)) {
	macro.Rewrite("$_.onpointerleave")
	return onpointerleave
}

// SetOnpointerleave prop
// js:"onpointerleave"
func (*HTMLObjectElement) SetOnpointerleave(onpointerleave func(window.Event)) {
	macro.Rewrite("$_.onpointerleave = $1", onpointerleave)
}

// Onpointermove prop
// js:"onpointermove"
func (*HTMLObjectElement) Onpointermove() (onpointermove func(window.Event)) {
	macro.Rewrite("$_.onpointermove")
	return onpointermove
}

// SetOnpointermove prop
// js:"onpointermove"
func (*HTMLObjectElement) SetOnpointermove(onpointermove func(window.Event)) {
	macro.Rewrite("$_.onpointermove = $1", onpointermove)
}

// Onpointerout prop
// js:"onpointerout"
func (*HTMLObjectElement) Onpointerout() (onpointerout func(window.Event)) {
	macro.Rewrite("$_.onpointerout")
	return onpointerout
}

// SetOnpointerout prop
// js:"onpointerout"
func (*HTMLObjectElement) SetOnpointerout(onpointerout func(window.Event)) {
	macro.Rewrite("$_.onpointerout = $1", onpointerout)
}

// Onpointerover prop
// js:"onpointerover"
func (*HTMLObjectElement) Onpointerover() (onpointerover func(window.Event)) {
	macro.Rewrite("$_.onpointerover")
	return onpointerover
}

// SetOnpointerover prop
// js:"onpointerover"
func (*HTMLObjectElement) SetOnpointerover(onpointerover func(window.Event)) {
	macro.Rewrite("$_.onpointerover = $1", onpointerover)
}

// Onpointerup prop
// js:"onpointerup"
func (*HTMLObjectElement) Onpointerup() (onpointerup func(window.Event)) {
	macro.Rewrite("$_.onpointerup")
	return onpointerup
}

// SetOnpointerup prop
// js:"onpointerup"
func (*HTMLObjectElement) SetOnpointerup(onpointerup func(window.Event)) {
	macro.Rewrite("$_.onpointerup = $1", onpointerup)
}

// Onwheel prop
// js:"onwheel"
func (*HTMLObjectElement) Onwheel() (onwheel func(window.Event)) {
	macro.Rewrite("$_.onwheel")
	return onwheel
}

// SetOnwheel prop
// js:"onwheel"
func (*HTMLObjectElement) SetOnwheel(onwheel func(window.Event)) {
	macro.Rewrite("$_.onwheel = $1", onwheel)
}

// ChildElementCount prop
// js:"childElementCount"
func (*HTMLObjectElement) ChildElementCount() (childElementCount uint) {
	macro.Rewrite("$_.childElementCount")
	return childElementCount
}

// FirstElementChild prop
// js:"firstElementChild"
func (*HTMLObjectElement) FirstElementChild() (firstElementChild window.Element) {
	macro.Rewrite("$_.firstElementChild")
	return firstElementChild
}

// LastElementChild prop
// js:"lastElementChild"
func (*HTMLObjectElement) LastElementChild() (lastElementChild window.Element) {
	macro.Rewrite("$_.lastElementChild")
	return lastElementChild
}

// NextElementSibling prop
// js:"nextElementSibling"
func (*HTMLObjectElement) NextElementSibling() (nextElementSibling window.Element) {
	macro.Rewrite("$_.nextElementSibling")
	return nextElementSibling
}

// PreviousElementSibling prop
// js:"previousElementSibling"
func (*HTMLObjectElement) PreviousElementSibling() (previousElementSibling window.Element) {
	macro.Rewrite("$_.previousElementSibling")
	return previousElementSibling
}

// Attributes prop
// js:"attributes"
func (*HTMLObjectElement) Attributes() (attributes *window.NamedNodeMap) {
	macro.Rewrite("$_.attributes")
	return attributes
}

// BaseURI prop
// js:"baseURI"
func (*HTMLObjectElement) BaseURI() (baseURI string) {
	macro.Rewrite("$_.baseURI")
	return baseURI
}

// ChildNodes prop
// js:"childNodes"
func (*HTMLObjectElement) ChildNodes() (childNodes *window.NodeList) {
	macro.Rewrite("$_.childNodes")
	return childNodes
}

// FirstChild prop
// js:"firstChild"
func (*HTMLObjectElement) FirstChild() (firstChild window.Node) {
	macro.Rewrite("$_.firstChild")
	return firstChild
}

// LastChild prop
// js:"lastChild"
func (*HTMLObjectElement) LastChild() (lastChild window.Node) {
	macro.Rewrite("$_.lastChild")
	return lastChild
}

// LocalName prop
// js:"localName"
func (*HTMLObjectElement) LocalName() (localName string) {
	macro.Rewrite("$_.localName")
	return localName
}

// NamespaceURI prop
// js:"namespaceURI"
func (*HTMLObjectElement) NamespaceURI() (namespaceURI string) {
	macro.Rewrite("$_.namespaceURI")
	return namespaceURI
}

// NextSibling prop
// js:"nextSibling"
func (*HTMLObjectElement) NextSibling() (nextSibling window.Node) {
	macro.Rewrite("$_.nextSibling")
	return nextSibling
}

// NodeName prop
// js:"nodeName"
func (*HTMLObjectElement) NodeName() (nodeName string) {
	macro.Rewrite("$_.nodeName")
	return nodeName
}

// NodeType prop
// js:"nodeType"
func (*HTMLObjectElement) NodeType() (nodeType uint8) {
	macro.Rewrite("$_.nodeType")
	return nodeType
}

// NodeValue prop
// js:"nodeValue"
func (*HTMLObjectElement) NodeValue() (nodeValue string) {
	macro.Rewrite("$_.nodeValue")
	return nodeValue
}

// SetNodeValue prop
// js:"nodeValue"
func (*HTMLObjectElement) SetNodeValue(nodeValue string) {
	macro.Rewrite("$_.nodeValue = $1", nodeValue)
}

// OwnerDocument prop
// js:"ownerDocument"
func (*HTMLObjectElement) OwnerDocument() (ownerDocument window.Document) {
	macro.Rewrite("$_.ownerDocument")
	return ownerDocument
}

// ParentElement prop
// js:"parentElement"
func (*HTMLObjectElement) ParentElement() (parentElement window.HTMLElement) {
	macro.Rewrite("$_.parentElement")
	return parentElement
}

// ParentNode prop
// js:"parentNode"
func (*HTMLObjectElement) ParentNode() (parentNode window.Node) {
	macro.Rewrite("$_.parentNode")
	return parentNode
}

// PreviousSibling prop
// js:"previousSibling"
func (*HTMLObjectElement) PreviousSibling() (previousSibling window.Node) {
	macro.Rewrite("$_.previousSibling")
	return previousSibling
}

// TextContent prop
// js:"textContent"
func (*HTMLObjectElement) TextContent() (textContent string) {
	macro.Rewrite("$_.textContent")
	return textContent
}

// SetTextContent prop
// js:"textContent"
func (*HTMLObjectElement) SetTextContent(textContent string) {
	macro.Rewrite("$_.textContent = $1", textContent)
}
