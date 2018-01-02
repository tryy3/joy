package xmldocument

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/document"
	"github.com/matthewmueller/joy/dom/utils"
	"github.com/matthewmueller/joy/dom/dom"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/element"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/dom/touch"
	"github.com/matthewmueller/joy/dom/treewalker"
	"github.com/matthewmueller/joy/dom/html"
	"github.com/matthewmueller/joy/dom/url"
	"github.com/matthewmueller/joy/dom/ui"
	"github.com/matthewmueller/joy/dom/ms"
	"github.com/matthewmueller/joy/dom/css"
)

var _ document.Document = (*XMLDocument)(nil)
var _ utils.GlobalEventHandlers = (*XMLDocument)(nil)
var _ dom.NodeSelector = (*XMLDocument)(nil)
var _ document.DocumentEvent = (*XMLDocument)(nil)
var _ dom.Node = (*XMLDocument)(nil)
var _ event.EventTarget = (*XMLDocument)(nil)

type XMLDocument struct {
}

func (*XMLDocument) AdoptNode(source dom.Node) (w dom.Node) {
	macro.Rewrite("$_.adoptNode($1)", source)
	return w
}

func (*XMLDocument) CaptureEvents() {
	macro.Rewrite("$_.captureEvents()")
}

func (*XMLDocument) CaretRangeFromPoint(x float32, y float32) (w *dom.Range) {
	macro.Rewrite("$_.caretRangeFromPoint($1, $2)", x, y)
	return w
}

func (*XMLDocument) Clear() {
	macro.Rewrite("$_.clear()")
}

func (*XMLDocument) Close() {
	macro.Rewrite("$_.close()")
}

func (*XMLDocument) CreateAttribute(name string) (w *dom.Attr) {
	macro.Rewrite("$_.createAttribute($1)", name)
	return w
}

func (*XMLDocument) CreateAttributeNS(namespaceURI string, qualifiedName string) (w *dom.Attr) {
	macro.Rewrite("$_.createAttributeNS($1, $2)", namespaceURI, qualifiedName)
	return w
}

func (*XMLDocument) CreateCDATASection(data string) (w *dom.CDATASection) {
	macro.Rewrite("$_.createCDATASection($1)", data)
	return w
}

func (*XMLDocument) CreateComment(data string) (w *dom.Comment) {
	macro.Rewrite("$_.createComment($1)", data)
	return w
}

func (*XMLDocument) CreateDocumentFragment() (w *document.DocumentFragment) {
	macro.Rewrite("$_.createDocumentFragment()")
	return w
}

func (*XMLDocument) CreateElement(tagName string) (w element.Element) {
	macro.Rewrite("$_.createElement($1)", tagName)
	return w
}

func (*XMLDocument) CreateElementNS(namespaceURI string, qualifiedName string) (w element.Element) {
	macro.Rewrite("$_.createElementNS($1, $2)", namespaceURI, qualifiedName)
	return w
}

func (*XMLDocument) CreateExpression(expression string, resolver *utils.XPathNSResolver) (w *utils.XPathExpression) {
	macro.Rewrite("$_.createExpression($1, $2)", expression, resolver)
	return w
}

func (*XMLDocument) CreateNodeIterator(root dom.Node, whatToShow *uint, filter *dom.NodeFilter, entityReferenceExpansion *bool) (w *dom.NodeIterator) {
	macro.Rewrite("$_.createNodeIterator($1, $2, $3, $4)", root, whatToShow, filter, entityReferenceExpansion)
	return w
}

func (*XMLDocument) CreateNSResolver(nodeResolver dom.Node) (x *utils.XPathNSResolver) {
	macro.Rewrite("$_.createNSResolver($1)", nodeResolver)
	return x
}

func (*XMLDocument) CreateProcessingInstruction(target string, data string) (w *dom.ProcessingInstruction) {
	macro.Rewrite("$_.createProcessingInstruction($1, $2)", target, data)
	return w
}

func (*XMLDocument) CreateRange() (w *dom.Range) {
	macro.Rewrite("$_.createRange()")
	return w
}

func (*XMLDocument) CreateTextNode(data string) (w dom.Text) {
	macro.Rewrite("$_.createTextNode($1)", data)
	return w
}

func (*XMLDocument) CreateTouch(view *window.Window, target event.EventTarget, identifier int, pageX int, pageY int, screenX int, screenY int) (w *touch.Touch) {
	macro.Rewrite("$_.createTouch($1, $2, $3, $4, $5, $6, $7)", view, target, identifier, pageX, pageY, screenX, screenY)
	return w
}

func (*XMLDocument) CreateTouchList(touches *touch.Touch) (w *touch.TouchList) {
	macro.Rewrite("$_.createTouchList($1)", touches)
	return w
}

func (*XMLDocument) CreateTreeWalker(root dom.Node, whatToShow *uint, filter *dom.NodeFilter, entityReferenceExpansion *bool) (w *treewalker.TreeWalker) {
	macro.Rewrite("$_.createTreeWalker($1, $2, $3, $4)", root, whatToShow, filter, entityReferenceExpansion)
	return w
}

func (*XMLDocument) ElementFromPoint(x int, y int) (w element.Element) {
	macro.Rewrite("$_.elementFromPoint($1, $2)", x, y)
	return w
}

func (*XMLDocument) Evaluate(expression string, contextNode dom.Node, resolver *utils.XPathNSResolver, kind uint8, result *utils.XPathResult) (w *utils.XPathResult) {
	macro.Rewrite("$_.evaluate($1, $2, $3, $4, $5)", expression, contextNode, resolver, kind, result)
	return w
}

func (*XMLDocument) ExecCommand(commandId string, showUI *bool, value *interface{}) (b bool) {
	macro.Rewrite("$_.execCommand($1, $2, $3)", commandId, showUI, value)
	return b
}

func (*XMLDocument) ExecCommandShowHelp(commandId string) (b bool) {
	macro.Rewrite("$_.execCommandShowHelp($1)", commandId)
	return b
}

func (*XMLDocument) ExitFullscreen() {
	macro.Rewrite("$_.exitFullscreen()")
}

func (*XMLDocument) ExitPointerLock() {
	macro.Rewrite("$_.exitPointerLock()")
}

func (*XMLDocument) Focus() {
	macro.Rewrite("$_.focus()")
}

func (*XMLDocument) GetElementByID(elementId string) (w element.Element) {
	macro.Rewrite("$_.getElementById($1)", elementId)
	return w
}

func (*XMLDocument) GetElementsByClassName(classNames string) (w *dom.NodeList) {
	macro.Rewrite("$_.getElementsByClassName($1)", classNames)
	return w
}

func (*XMLDocument) GetElementsByName(elementName string) (w *dom.NodeList) {
	macro.Rewrite("$_.getElementsByName($1)", elementName)
	return w
}

func (*XMLDocument) GetElementsByTagName(tagname string) (w *dom.NodeList) {
	macro.Rewrite("$_.getElementsByTagName($1)", tagname)
	return w
}

func (*XMLDocument) GetElementsByTagNameNS(namespaceURI string, localName string) (w *dom.NodeList) {
	macro.Rewrite("$_.getElementsByTagNameNS($1, $2)", namespaceURI, localName)
	return w
}

func (*XMLDocument) GetSelection() (w *dom.Selection) {
	macro.Rewrite("$_.getSelection()")
	return w
}

func (*XMLDocument) HasFocus() (b bool) {
	macro.Rewrite("$_.hasFocus()")
	return b
}

func (*XMLDocument) ImportNode(importedNode dom.Node, deep bool) (w dom.Node) {
	macro.Rewrite("$_.importNode($1, $2)", importedNode, deep)
	return w
}

func (*XMLDocument) MsElementsFromPoint(x float32, y float32) (w *dom.NodeList) {
	macro.Rewrite("$_.msElementsFromPoint($1, $2)", x, y)
	return w
}

func (*XMLDocument) MsElementsFromRect(left float32, top float32, width float32, height float32) (w *dom.NodeList) {
	macro.Rewrite("$_.msElementsFromRect($1, $2, $3, $4)", left, top, width, height)
	return w
}

func (*XMLDocument) Open(url *string, name *string, features *string, replace *bool) (i interface{}) {
	macro.Rewrite("$_.open($1, $2, $3, $4)", url, name, features, replace)
	return i
}

func (*XMLDocument) QueryCommandEnabled(commandId string) (b bool) {
	macro.Rewrite("$_.queryCommandEnabled($1)", commandId)
	return b
}

func (*XMLDocument) QueryCommandIndeterm(commandId string) (b bool) {
	macro.Rewrite("$_.queryCommandIndeterm($1)", commandId)
	return b
}

func (*XMLDocument) QueryCommandState(commandId string) (b bool) {
	macro.Rewrite("$_.queryCommandState($1)", commandId)
	return b
}

func (*XMLDocument) QueryCommandSupported(commandId string) (b bool) {
	macro.Rewrite("$_.queryCommandSupported($1)", commandId)
	return b
}

func (*XMLDocument) QueryCommandText(commandId string) (s string) {
	macro.Rewrite("$_.queryCommandText($1)", commandId)
	return s
}

func (*XMLDocument) QueryCommandValue(commandId string) (s string) {
	macro.Rewrite("$_.queryCommandValue($1)", commandId)
	return s
}

func (*XMLDocument) ReleaseEvents() {
	macro.Rewrite("$_.releaseEvents()")
}

func (*XMLDocument) UpdateSettings() {
	macro.Rewrite("$_.updateSettings()")
}

func (*XMLDocument) WebkitCancelFullScreen() {
	macro.Rewrite("$_.webkitCancelFullScreen()")
}

func (*XMLDocument) WebkitExitFullscreen() {
	macro.Rewrite("$_.webkitExitFullscreen()")
}

func (*XMLDocument) Write(content string) {
	macro.Rewrite("$_.write($1)", content)
}

func (*XMLDocument) Writeln(content string) {
	macro.Rewrite("$_.writeln($1)", content)
}

func (*XMLDocument) QuerySelector(selectors string) (w element.Element) {
	macro.Rewrite("$_.querySelector($1)", selectors)
	return w
}

func (*XMLDocument) QuerySelectorAll(selectors string) (w *dom.NodeList) {
	macro.Rewrite("$_.querySelectorAll($1)", selectors)
	return w
}

func (*XMLDocument) CreateEvent(eventInterface string) (w event.Event) {
	macro.Rewrite("$_.createEvent($1)", eventInterface)
	return w
}

func (*XMLDocument) AppendChild(newChild dom.Node) (w dom.Node) {
	macro.Rewrite("$_.appendChild($1)", newChild)
	return w
}

func (*XMLDocument) CloneNode(deep *bool) (w dom.Node) {
	macro.Rewrite("$_.cloneNode($1)", deep)
	return w
}

func (*XMLDocument) CompareDocumentPosition(other dom.Node) (u uint8) {
	macro.Rewrite("$_.compareDocumentPosition($1)", other)
	return u
}

func (*XMLDocument) Contains(child dom.Node) (b bool) {
	macro.Rewrite("$_.contains($1)", child)
	return b
}

func (*XMLDocument) HasAttributes() (b bool) {
	macro.Rewrite("$_.hasAttributes()")
	return b
}

func (*XMLDocument) HasChildNodes() (b bool) {
	macro.Rewrite("$_.hasChildNodes()")
	return b
}

func (*XMLDocument) InsertBefore(newChild dom.Node, refChild *dom.Node) (w dom.Node) {
	macro.Rewrite("$_.insertBefore($1, $2)", newChild, refChild)
	return w
}

func (*XMLDocument) IsDefaultNamespace(namespaceURI string) (b bool) {
	macro.Rewrite("$_.isDefaultNamespace($1)", namespaceURI)
	return b
}

func (*XMLDocument) IsEqualNode(arg dom.Node) (b bool) {
	macro.Rewrite("$_.isEqualNode($1)", arg)
	return b
}

func (*XMLDocument) IsSameNode(other dom.Node) (b bool) {
	macro.Rewrite("$_.isSameNode($1)", other)
	return b
}

func (*XMLDocument) LookupNamespaceURI(prefix string) (s string) {
	macro.Rewrite("$_.lookupNamespaceURI($1)", prefix)
	return s
}

func (*XMLDocument) LookupPrefix(namespaceURI string) (s string) {
	macro.Rewrite("$_.lookupPrefix($1)", namespaceURI)
	return s
}

func (*XMLDocument) Normalize() {
	macro.Rewrite("$_.normalize()")
}

func (*XMLDocument) RemoveChild(oldChild dom.Node) (w dom.Node) {
	macro.Rewrite("$_.removeChild($1)", oldChild)
	return w
}

func (*XMLDocument) ReplaceChild(newChild dom.Node, oldChild dom.Node) (w dom.Node) {
	macro.Rewrite("$_.replaceChild($1, $2)", newChild, oldChild)
	return w
}

func (*XMLDocument) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*XMLDocument) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*XMLDocument) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*XMLDocument) ActiveElement() (activeElement element.Element) {
	macro.Rewrite("$_.activeElement")
	return activeElement
}

func (*XMLDocument) AlinkColor() (alinkColor string) {
	macro.Rewrite("$_.alinkColor")
	return alinkColor
}

func (*XMLDocument) SetAlinkColor(alinkColor string) {
	macro.Rewrite("$_.alinkColor = $1", alinkColor)
}

func (*XMLDocument) All() (all *html.HTMLAllCollection) {
	macro.Rewrite("$_.all")
	return all
}

func (*XMLDocument) Anchors() (anchors html.HTMLCollection) {
	macro.Rewrite("$_.anchors")
	return anchors
}

func (*XMLDocument) Applets() (applets html.HTMLCollection) {
	macro.Rewrite("$_.applets")
	return applets
}

func (*XMLDocument) BgColor() (bgColor string) {
	macro.Rewrite("$_.bgColor")
	return bgColor
}

func (*XMLDocument) SetBgColor(bgColor string) {
	macro.Rewrite("$_.bgColor = $1", bgColor)
}

func (*XMLDocument) Body() (body html.HTMLElement) {
	macro.Rewrite("$_.body")
	return body
}

func (*XMLDocument) SetBody(body html.HTMLElement) {
	macro.Rewrite("$_.body = $1", body)
}

func (*XMLDocument) CharacterSet() (characterSet string) {
	macro.Rewrite("$_.characterSet")
	return characterSet
}

func (*XMLDocument) Charset() (charset string) {
	macro.Rewrite("$_.charset")
	return charset
}

func (*XMLDocument) SetCharset(charset string) {
	macro.Rewrite("$_.charset = $1", charset)
}

func (*XMLDocument) CompatMode() (compatMode string) {
	macro.Rewrite("$_.compatMode")
	return compatMode
}

func (*XMLDocument) Cookie() (cookie string) {
	macro.Rewrite("$_.cookie")
	return cookie
}

func (*XMLDocument) SetCookie(cookie string) {
	macro.Rewrite("$_.cookie = $1", cookie)
}

func (*XMLDocument) CurrentScript() (currentScript interface{}) {
	macro.Rewrite("$_.currentScript")
	return currentScript
}

func (*XMLDocument) DefaultView() (defaultView *window.Window) {
	macro.Rewrite("$_.defaultView")
	return defaultView
}

func (*XMLDocument) DesignMode() (designMode string) {
	macro.Rewrite("$_.designMode")
	return designMode
}

func (*XMLDocument) SetDesignMode(designMode string) {
	macro.Rewrite("$_.designMode = $1", designMode)
}

func (*XMLDocument) Dir() (dir string) {
	macro.Rewrite("$_.dir")
	return dir
}

func (*XMLDocument) SetDir(dir string) {
	macro.Rewrite("$_.dir = $1", dir)
}

func (*XMLDocument) Doctype() (doctype *document.DocumentType) {
	macro.Rewrite("$_.doctype")
	return doctype
}

func (*XMLDocument) DocumentElement() (documentElement element.Element) {
	macro.Rewrite("$_.documentElement")
	return documentElement
}

func (*XMLDocument) Domain() (domain string) {
	macro.Rewrite("$_.domain")
	return domain
}

func (*XMLDocument) SetDomain(domain string) {
	macro.Rewrite("$_.domain = $1", domain)
}

func (*XMLDocument) Embeds() (embeds html.HTMLCollection) {
	macro.Rewrite("$_.embeds")
	return embeds
}

func (*XMLDocument) FgColor() (fgColor string) {
	macro.Rewrite("$_.fgColor")
	return fgColor
}

func (*XMLDocument) SetFgColor(fgColor string) {
	macro.Rewrite("$_.fgColor = $1", fgColor)
}

func (*XMLDocument) Forms() (forms html.HTMLCollection) {
	macro.Rewrite("$_.forms")
	return forms
}

func (*XMLDocument) FullscreenElement() (fullscreenElement element.Element) {
	macro.Rewrite("$_.fullscreenElement")
	return fullscreenElement
}

func (*XMLDocument) FullscreenEnabled() (fullscreenEnabled bool) {
	macro.Rewrite("$_.fullscreenEnabled")
	return fullscreenEnabled
}

func (*XMLDocument) Head() (head *html.HTMLHeadElement) {
	macro.Rewrite("$_.head")
	return head
}

func (*XMLDocument) Hidden() (hidden bool) {
	macro.Rewrite("$_.hidden")
	return hidden
}

func (*XMLDocument) Images() (images html.HTMLCollection) {
	macro.Rewrite("$_.images")
	return images
}

func (*XMLDocument) Implementation() (implementation *dom.DOMImplementation) {
	macro.Rewrite("$_.implementation")
	return implementation
}

func (*XMLDocument) InputEncoding() (inputEncoding string) {
	macro.Rewrite("$_.inputEncoding")
	return inputEncoding
}

func (*XMLDocument) LastModified() (lastModified string) {
	macro.Rewrite("$_.lastModified")
	return lastModified
}

func (*XMLDocument) LinkColor() (linkColor string) {
	macro.Rewrite("$_.linkColor")
	return linkColor
}

func (*XMLDocument) SetLinkColor(linkColor string) {
	macro.Rewrite("$_.linkColor = $1", linkColor)
}

func (*XMLDocument) Links() (links html.HTMLCollection) {
	macro.Rewrite("$_.links")
	return links
}

func (*XMLDocument) Location() (location *url.Location) {
	macro.Rewrite("$_.location")
	return location
}

func (*XMLDocument) MsCapsLockWarningOff() (msCapsLockWarningOff bool) {
	macro.Rewrite("$_.msCapsLockWarningOff")
	return msCapsLockWarningOff
}

func (*XMLDocument) SetMsCapsLockWarningOff(msCapsLockWarningOff bool) {
	macro.Rewrite("$_.msCapsLockWarningOff = $1", msCapsLockWarningOff)
}

func (*XMLDocument) MsCSSOMElementFloatMetrics() (msCSSOMElementFloatMetrics bool) {
	macro.Rewrite("$_.msCSSOMElementFloatMetrics")
	return msCSSOMElementFloatMetrics
}

func (*XMLDocument) SetMsCSSOMElementFloatMetrics(msCSSOMElementFloatMetrics bool) {
	macro.Rewrite("$_.msCSSOMElementFloatMetrics = $1", msCSSOMElementFloatMetrics)
}

func (*XMLDocument) Onabort() (onabort func(event.Event)) {
	macro.Rewrite("$_.onabort")
	return onabort
}

func (*XMLDocument) SetOnabort(onabort func(event.Event)) {
	macro.Rewrite("$_.onabort = $1", onabort)
}

func (*XMLDocument) Onactivate() (onactivate func(event.Event)) {
	macro.Rewrite("$_.onactivate")
	return onactivate
}

func (*XMLDocument) SetOnactivate(onactivate func(event.Event)) {
	macro.Rewrite("$_.onactivate = $1", onactivate)
}

func (*XMLDocument) Onbeforeactivate() (onbeforeactivate func(event.Event)) {
	macro.Rewrite("$_.onbeforeactivate")
	return onbeforeactivate
}

func (*XMLDocument) SetOnbeforeactivate(onbeforeactivate func(event.Event)) {
	macro.Rewrite("$_.onbeforeactivate = $1", onbeforeactivate)
}

func (*XMLDocument) Onbeforedeactivate() (onbeforedeactivate func(event.Event)) {
	macro.Rewrite("$_.onbeforedeactivate")
	return onbeforedeactivate
}

func (*XMLDocument) SetOnbeforedeactivate(onbeforedeactivate func(event.Event)) {
	macro.Rewrite("$_.onbeforedeactivate = $1", onbeforedeactivate)
}

func (*XMLDocument) Onblur() (onblur func(event.Event)) {
	macro.Rewrite("$_.onblur")
	return onblur
}

func (*XMLDocument) SetOnblur(onblur func(event.Event)) {
	macro.Rewrite("$_.onblur = $1", onblur)
}

func (*XMLDocument) Oncanplay() (oncanplay func(event.Event)) {
	macro.Rewrite("$_.oncanplay")
	return oncanplay
}

func (*XMLDocument) SetOncanplay(oncanplay func(event.Event)) {
	macro.Rewrite("$_.oncanplay = $1", oncanplay)
}

func (*XMLDocument) Oncanplaythrough() (oncanplaythrough func(event.Event)) {
	macro.Rewrite("$_.oncanplaythrough")
	return oncanplaythrough
}

func (*XMLDocument) SetOncanplaythrough(oncanplaythrough func(event.Event)) {
	macro.Rewrite("$_.oncanplaythrough = $1", oncanplaythrough)
}

func (*XMLDocument) Onchange() (onchange func(event.Event)) {
	macro.Rewrite("$_.onchange")
	return onchange
}

func (*XMLDocument) SetOnchange(onchange func(event.Event)) {
	macro.Rewrite("$_.onchange = $1", onchange)
}

func (*XMLDocument) Onclick() (onclick func(event.Event)) {
	macro.Rewrite("$_.onclick")
	return onclick
}

func (*XMLDocument) SetOnclick(onclick func(event.Event)) {
	macro.Rewrite("$_.onclick = $1", onclick)
}

func (*XMLDocument) Oncontextmenu() (oncontextmenu func(event.Event)) {
	macro.Rewrite("$_.oncontextmenu")
	return oncontextmenu
}

func (*XMLDocument) SetOncontextmenu(oncontextmenu func(event.Event)) {
	macro.Rewrite("$_.oncontextmenu = $1", oncontextmenu)
}

func (*XMLDocument) Ondblclick() (ondblclick func(event.Event)) {
	macro.Rewrite("$_.ondblclick")
	return ondblclick
}

func (*XMLDocument) SetOndblclick(ondblclick func(event.Event)) {
	macro.Rewrite("$_.ondblclick = $1", ondblclick)
}

func (*XMLDocument) Ondeactivate() (ondeactivate func(event.Event)) {
	macro.Rewrite("$_.ondeactivate")
	return ondeactivate
}

func (*XMLDocument) SetOndeactivate(ondeactivate func(event.Event)) {
	macro.Rewrite("$_.ondeactivate = $1", ondeactivate)
}

func (*XMLDocument) Ondrag() (ondrag func(event.Event)) {
	macro.Rewrite("$_.ondrag")
	return ondrag
}

func (*XMLDocument) SetOndrag(ondrag func(event.Event)) {
	macro.Rewrite("$_.ondrag = $1", ondrag)
}

func (*XMLDocument) Ondragend() (ondragend func(event.Event)) {
	macro.Rewrite("$_.ondragend")
	return ondragend
}

func (*XMLDocument) SetOndragend(ondragend func(event.Event)) {
	macro.Rewrite("$_.ondragend = $1", ondragend)
}

func (*XMLDocument) Ondragenter() (ondragenter func(event.Event)) {
	macro.Rewrite("$_.ondragenter")
	return ondragenter
}

func (*XMLDocument) SetOndragenter(ondragenter func(event.Event)) {
	macro.Rewrite("$_.ondragenter = $1", ondragenter)
}

func (*XMLDocument) Ondragleave() (ondragleave func(event.Event)) {
	macro.Rewrite("$_.ondragleave")
	return ondragleave
}

func (*XMLDocument) SetOndragleave(ondragleave func(event.Event)) {
	macro.Rewrite("$_.ondragleave = $1", ondragleave)
}

func (*XMLDocument) Ondragover() (ondragover func(event.Event)) {
	macro.Rewrite("$_.ondragover")
	return ondragover
}

func (*XMLDocument) SetOndragover(ondragover func(event.Event)) {
	macro.Rewrite("$_.ondragover = $1", ondragover)
}

func (*XMLDocument) Ondragstart() (ondragstart func(event.Event)) {
	macro.Rewrite("$_.ondragstart")
	return ondragstart
}

func (*XMLDocument) SetOndragstart(ondragstart func(event.Event)) {
	macro.Rewrite("$_.ondragstart = $1", ondragstart)
}

func (*XMLDocument) Ondrop() (ondrop func(event.Event)) {
	macro.Rewrite("$_.ondrop")
	return ondrop
}

func (*XMLDocument) SetOndrop(ondrop func(event.Event)) {
	macro.Rewrite("$_.ondrop = $1", ondrop)
}

func (*XMLDocument) Ondurationchange() (ondurationchange func(event.Event)) {
	macro.Rewrite("$_.ondurationchange")
	return ondurationchange
}

func (*XMLDocument) SetOndurationchange(ondurationchange func(event.Event)) {
	macro.Rewrite("$_.ondurationchange = $1", ondurationchange)
}

func (*XMLDocument) Onemptied() (onemptied func(event.Event)) {
	macro.Rewrite("$_.onemptied")
	return onemptied
}

func (*XMLDocument) SetOnemptied(onemptied func(event.Event)) {
	macro.Rewrite("$_.onemptied = $1", onemptied)
}

func (*XMLDocument) Onended() (onended func(event.Event)) {
	macro.Rewrite("$_.onended")
	return onended
}

func (*XMLDocument) SetOnended(onended func(event.Event)) {
	macro.Rewrite("$_.onended = $1", onended)
}

func (*XMLDocument) Onerror() (onerror func(event.Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

func (*XMLDocument) SetOnerror(onerror func(event.Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

func (*XMLDocument) Onfocus() (onfocus func(event.Event)) {
	macro.Rewrite("$_.onfocus")
	return onfocus
}

func (*XMLDocument) SetOnfocus(onfocus func(event.Event)) {
	macro.Rewrite("$_.onfocus = $1", onfocus)
}

func (*XMLDocument) Onfullscreenchange() (onfullscreenchange func(event.Event)) {
	macro.Rewrite("$_.onfullscreenchange")
	return onfullscreenchange
}

func (*XMLDocument) SetOnfullscreenchange(onfullscreenchange func(event.Event)) {
	macro.Rewrite("$_.onfullscreenchange = $1", onfullscreenchange)
}

func (*XMLDocument) Onfullscreenerror() (onfullscreenerror func(event.Event)) {
	macro.Rewrite("$_.onfullscreenerror")
	return onfullscreenerror
}

func (*XMLDocument) SetOnfullscreenerror(onfullscreenerror func(event.Event)) {
	macro.Rewrite("$_.onfullscreenerror = $1", onfullscreenerror)
}

func (*XMLDocument) Oninput() (oninput func(event.Event)) {
	macro.Rewrite("$_.oninput")
	return oninput
}

func (*XMLDocument) SetOninput(oninput func(event.Event)) {
	macro.Rewrite("$_.oninput = $1", oninput)
}

func (*XMLDocument) Oninvalid() (oninvalid func(event.Event)) {
	macro.Rewrite("$_.oninvalid")
	return oninvalid
}

func (*XMLDocument) SetOninvalid(oninvalid func(event.Event)) {
	macro.Rewrite("$_.oninvalid = $1", oninvalid)
}

func (*XMLDocument) Onkeydown() (onkeydown func(event.Event)) {
	macro.Rewrite("$_.onkeydown")
	return onkeydown
}

func (*XMLDocument) SetOnkeydown(onkeydown func(event.Event)) {
	macro.Rewrite("$_.onkeydown = $1", onkeydown)
}

func (*XMLDocument) Onkeypress() (onkeypress func(event.Event)) {
	macro.Rewrite("$_.onkeypress")
	return onkeypress
}

func (*XMLDocument) SetOnkeypress(onkeypress func(event.Event)) {
	macro.Rewrite("$_.onkeypress = $1", onkeypress)
}

func (*XMLDocument) Onkeyup() (onkeyup func(event.Event)) {
	macro.Rewrite("$_.onkeyup")
	return onkeyup
}

func (*XMLDocument) SetOnkeyup(onkeyup func(event.Event)) {
	macro.Rewrite("$_.onkeyup = $1", onkeyup)
}

func (*XMLDocument) Onload() (onload func(event.Event)) {
	macro.Rewrite("$_.onload")
	return onload
}

func (*XMLDocument) SetOnload(onload func(event.Event)) {
	macro.Rewrite("$_.onload = $1", onload)
}

func (*XMLDocument) Onloadeddata() (onloadeddata func(event.Event)) {
	macro.Rewrite("$_.onloadeddata")
	return onloadeddata
}

func (*XMLDocument) SetOnloadeddata(onloadeddata func(event.Event)) {
	macro.Rewrite("$_.onloadeddata = $1", onloadeddata)
}

func (*XMLDocument) Onloadedmetadata() (onloadedmetadata func(event.Event)) {
	macro.Rewrite("$_.onloadedmetadata")
	return onloadedmetadata
}

func (*XMLDocument) SetOnloadedmetadata(onloadedmetadata func(event.Event)) {
	macro.Rewrite("$_.onloadedmetadata = $1", onloadedmetadata)
}

func (*XMLDocument) Onloadstart() (onloadstart func(event.Event)) {
	macro.Rewrite("$_.onloadstart")
	return onloadstart
}

func (*XMLDocument) SetOnloadstart(onloadstart func(event.Event)) {
	macro.Rewrite("$_.onloadstart = $1", onloadstart)
}

func (*XMLDocument) Onmousedown() (onmousedown func(event.Event)) {
	macro.Rewrite("$_.onmousedown")
	return onmousedown
}

func (*XMLDocument) SetOnmousedown(onmousedown func(event.Event)) {
	macro.Rewrite("$_.onmousedown = $1", onmousedown)
}

func (*XMLDocument) Onmousemove() (onmousemove func(event.Event)) {
	macro.Rewrite("$_.onmousemove")
	return onmousemove
}

func (*XMLDocument) SetOnmousemove(onmousemove func(event.Event)) {
	macro.Rewrite("$_.onmousemove = $1", onmousemove)
}

func (*XMLDocument) Onmouseout() (onmouseout func(event.Event)) {
	macro.Rewrite("$_.onmouseout")
	return onmouseout
}

func (*XMLDocument) SetOnmouseout(onmouseout func(event.Event)) {
	macro.Rewrite("$_.onmouseout = $1", onmouseout)
}

func (*XMLDocument) Onmouseover() (onmouseover func(event.Event)) {
	macro.Rewrite("$_.onmouseover")
	return onmouseover
}

func (*XMLDocument) SetOnmouseover(onmouseover func(event.Event)) {
	macro.Rewrite("$_.onmouseover = $1", onmouseover)
}

func (*XMLDocument) Onmouseup() (onmouseup func(event.Event)) {
	macro.Rewrite("$_.onmouseup")
	return onmouseup
}

func (*XMLDocument) SetOnmouseup(onmouseup func(event.Event)) {
	macro.Rewrite("$_.onmouseup = $1", onmouseup)
}

func (*XMLDocument) Onmousewheel() (onmousewheel func(event.Event)) {
	macro.Rewrite("$_.onmousewheel")
	return onmousewheel
}

func (*XMLDocument) SetOnmousewheel(onmousewheel func(event.Event)) {
	macro.Rewrite("$_.onmousewheel = $1", onmousewheel)
}

func (*XMLDocument) Onmscontentzoom() (onmscontentzoom func(ui.UIEvent)) {
	macro.Rewrite("$_.onmscontentzoom")
	return onmscontentzoom
}

func (*XMLDocument) SetOnmscontentzoom(onmscontentzoom func(ui.UIEvent)) {
	macro.Rewrite("$_.onmscontentzoom = $1", onmscontentzoom)
}

func (*XMLDocument) Onmsgesturechange() (onmsgesturechange func(event.Event)) {
	macro.Rewrite("$_.onmsgesturechange")
	return onmsgesturechange
}

func (*XMLDocument) SetOnmsgesturechange(onmsgesturechange func(event.Event)) {
	macro.Rewrite("$_.onmsgesturechange = $1", onmsgesturechange)
}

func (*XMLDocument) Onmsgesturedoubletap() (onmsgesturedoubletap func(event.Event)) {
	macro.Rewrite("$_.onmsgesturedoubletap")
	return onmsgesturedoubletap
}

func (*XMLDocument) SetOnmsgesturedoubletap(onmsgesturedoubletap func(event.Event)) {
	macro.Rewrite("$_.onmsgesturedoubletap = $1", onmsgesturedoubletap)
}

func (*XMLDocument) Onmsgestureend() (onmsgestureend func(event.Event)) {
	macro.Rewrite("$_.onmsgestureend")
	return onmsgestureend
}

func (*XMLDocument) SetOnmsgestureend(onmsgestureend func(event.Event)) {
	macro.Rewrite("$_.onmsgestureend = $1", onmsgestureend)
}

func (*XMLDocument) Onmsgesturehold() (onmsgesturehold func(event.Event)) {
	macro.Rewrite("$_.onmsgesturehold")
	return onmsgesturehold
}

func (*XMLDocument) SetOnmsgesturehold(onmsgesturehold func(event.Event)) {
	macro.Rewrite("$_.onmsgesturehold = $1", onmsgesturehold)
}

func (*XMLDocument) Onmsgesturestart() (onmsgesturestart func(event.Event)) {
	macro.Rewrite("$_.onmsgesturestart")
	return onmsgesturestart
}

func (*XMLDocument) SetOnmsgesturestart(onmsgesturestart func(event.Event)) {
	macro.Rewrite("$_.onmsgesturestart = $1", onmsgesturestart)
}

func (*XMLDocument) Onmsgesturetap() (onmsgesturetap func(event.Event)) {
	macro.Rewrite("$_.onmsgesturetap")
	return onmsgesturetap
}

func (*XMLDocument) SetOnmsgesturetap(onmsgesturetap func(event.Event)) {
	macro.Rewrite("$_.onmsgesturetap = $1", onmsgesturetap)
}

func (*XMLDocument) Onmsinertiastart() (onmsinertiastart func(event.Event)) {
	macro.Rewrite("$_.onmsinertiastart")
	return onmsinertiastart
}

func (*XMLDocument) SetOnmsinertiastart(onmsinertiastart func(event.Event)) {
	macro.Rewrite("$_.onmsinertiastart = $1", onmsinertiastart)
}

func (*XMLDocument) Onmsmanipulationstatechanged() (onmsmanipulationstatechanged func(*ms.MSManipulationEvent)) {
	macro.Rewrite("$_.onmsmanipulationstatechanged")
	return onmsmanipulationstatechanged
}

func (*XMLDocument) SetOnmsmanipulationstatechanged(onmsmanipulationstatechanged func(*ms.MSManipulationEvent)) {
	macro.Rewrite("$_.onmsmanipulationstatechanged = $1", onmsmanipulationstatechanged)
}

func (*XMLDocument) Onmspointercancel() (onmspointercancel func(event.Event)) {
	macro.Rewrite("$_.onmspointercancel")
	return onmspointercancel
}

func (*XMLDocument) SetOnmspointercancel(onmspointercancel func(event.Event)) {
	macro.Rewrite("$_.onmspointercancel = $1", onmspointercancel)
}

func (*XMLDocument) Onmspointerdown() (onmspointerdown func(event.Event)) {
	macro.Rewrite("$_.onmspointerdown")
	return onmspointerdown
}

func (*XMLDocument) SetOnmspointerdown(onmspointerdown func(event.Event)) {
	macro.Rewrite("$_.onmspointerdown = $1", onmspointerdown)
}

func (*XMLDocument) Onmspointerenter() (onmspointerenter func(event.Event)) {
	macro.Rewrite("$_.onmspointerenter")
	return onmspointerenter
}

func (*XMLDocument) SetOnmspointerenter(onmspointerenter func(event.Event)) {
	macro.Rewrite("$_.onmspointerenter = $1", onmspointerenter)
}

func (*XMLDocument) Onmspointerleave() (onmspointerleave func(event.Event)) {
	macro.Rewrite("$_.onmspointerleave")
	return onmspointerleave
}

func (*XMLDocument) SetOnmspointerleave(onmspointerleave func(event.Event)) {
	macro.Rewrite("$_.onmspointerleave = $1", onmspointerleave)
}

func (*XMLDocument) Onmspointermove() (onmspointermove func(event.Event)) {
	macro.Rewrite("$_.onmspointermove")
	return onmspointermove
}

func (*XMLDocument) SetOnmspointermove(onmspointermove func(event.Event)) {
	macro.Rewrite("$_.onmspointermove = $1", onmspointermove)
}

func (*XMLDocument) Onmspointerout() (onmspointerout func(event.Event)) {
	macro.Rewrite("$_.onmspointerout")
	return onmspointerout
}

func (*XMLDocument) SetOnmspointerout(onmspointerout func(event.Event)) {
	macro.Rewrite("$_.onmspointerout = $1", onmspointerout)
}

func (*XMLDocument) Onmspointerover() (onmspointerover func(event.Event)) {
	macro.Rewrite("$_.onmspointerover")
	return onmspointerover
}

func (*XMLDocument) SetOnmspointerover(onmspointerover func(event.Event)) {
	macro.Rewrite("$_.onmspointerover = $1", onmspointerover)
}

func (*XMLDocument) Onmspointerup() (onmspointerup func(event.Event)) {
	macro.Rewrite("$_.onmspointerup")
	return onmspointerup
}

func (*XMLDocument) SetOnmspointerup(onmspointerup func(event.Event)) {
	macro.Rewrite("$_.onmspointerup = $1", onmspointerup)
}

func (*XMLDocument) Onmssitemodejumplistitemremoved() (onmssitemodejumplistitemremoved func(*ms.MSSiteModeEvent)) {
	macro.Rewrite("$_.onmssitemodejumplistitemremoved")
	return onmssitemodejumplistitemremoved
}

func (*XMLDocument) SetOnmssitemodejumplistitemremoved(onmssitemodejumplistitemremoved func(*ms.MSSiteModeEvent)) {
	macro.Rewrite("$_.onmssitemodejumplistitemremoved = $1", onmssitemodejumplistitemremoved)
}

func (*XMLDocument) Onmsthumbnailclick() (onmsthumbnailclick func(*ms.MSSiteModeEvent)) {
	macro.Rewrite("$_.onmsthumbnailclick")
	return onmsthumbnailclick
}

func (*XMLDocument) SetOnmsthumbnailclick(onmsthumbnailclick func(*ms.MSSiteModeEvent)) {
	macro.Rewrite("$_.onmsthumbnailclick = $1", onmsthumbnailclick)
}

func (*XMLDocument) Onpause() (onpause func(event.Event)) {
	macro.Rewrite("$_.onpause")
	return onpause
}

func (*XMLDocument) SetOnpause(onpause func(event.Event)) {
	macro.Rewrite("$_.onpause = $1", onpause)
}

func (*XMLDocument) Onplay() (onplay func(event.Event)) {
	macro.Rewrite("$_.onplay")
	return onplay
}

func (*XMLDocument) SetOnplay(onplay func(event.Event)) {
	macro.Rewrite("$_.onplay = $1", onplay)
}

func (*XMLDocument) Onplaying() (onplaying func(event.Event)) {
	macro.Rewrite("$_.onplaying")
	return onplaying
}

func (*XMLDocument) SetOnplaying(onplaying func(event.Event)) {
	macro.Rewrite("$_.onplaying = $1", onplaying)
}

func (*XMLDocument) Onpointerlockchange() (onpointerlockchange func(event.Event)) {
	macro.Rewrite("$_.onpointerlockchange")
	return onpointerlockchange
}

func (*XMLDocument) SetOnpointerlockchange(onpointerlockchange func(event.Event)) {
	macro.Rewrite("$_.onpointerlockchange = $1", onpointerlockchange)
}

func (*XMLDocument) Onpointerlockerror() (onpointerlockerror func(event.Event)) {
	macro.Rewrite("$_.onpointerlockerror")
	return onpointerlockerror
}

func (*XMLDocument) SetOnpointerlockerror(onpointerlockerror func(event.Event)) {
	macro.Rewrite("$_.onpointerlockerror = $1", onpointerlockerror)
}

func (*XMLDocument) Onprogress() (onprogress func(event.Event)) {
	macro.Rewrite("$_.onprogress")
	return onprogress
}

func (*XMLDocument) SetOnprogress(onprogress func(event.Event)) {
	macro.Rewrite("$_.onprogress = $1", onprogress)
}

func (*XMLDocument) Onratechange() (onratechange func(event.Event)) {
	macro.Rewrite("$_.onratechange")
	return onratechange
}

func (*XMLDocument) SetOnratechange(onratechange func(event.Event)) {
	macro.Rewrite("$_.onratechange = $1", onratechange)
}

func (*XMLDocument) Onreadystatechange() (onreadystatechange func(event.Event)) {
	macro.Rewrite("$_.onreadystatechange")
	return onreadystatechange
}

func (*XMLDocument) SetOnreadystatechange(onreadystatechange func(event.Event)) {
	macro.Rewrite("$_.onreadystatechange = $1", onreadystatechange)
}

func (*XMLDocument) Onreset() (onreset func(event.Event)) {
	macro.Rewrite("$_.onreset")
	return onreset
}

func (*XMLDocument) SetOnreset(onreset func(event.Event)) {
	macro.Rewrite("$_.onreset = $1", onreset)
}

func (*XMLDocument) Onscroll() (onscroll func(event.Event)) {
	macro.Rewrite("$_.onscroll")
	return onscroll
}

func (*XMLDocument) SetOnscroll(onscroll func(event.Event)) {
	macro.Rewrite("$_.onscroll = $1", onscroll)
}

func (*XMLDocument) Onseeked() (onseeked func(event.Event)) {
	macro.Rewrite("$_.onseeked")
	return onseeked
}

func (*XMLDocument) SetOnseeked(onseeked func(event.Event)) {
	macro.Rewrite("$_.onseeked = $1", onseeked)
}

func (*XMLDocument) Onseeking() (onseeking func(event.Event)) {
	macro.Rewrite("$_.onseeking")
	return onseeking
}

func (*XMLDocument) SetOnseeking(onseeking func(event.Event)) {
	macro.Rewrite("$_.onseeking = $1", onseeking)
}

func (*XMLDocument) Onselect() (onselect func(event.Event)) {
	macro.Rewrite("$_.onselect")
	return onselect
}

func (*XMLDocument) SetOnselect(onselect func(event.Event)) {
	macro.Rewrite("$_.onselect = $1", onselect)
}

func (*XMLDocument) Onselectionchange() (onselectionchange func(event.Event)) {
	macro.Rewrite("$_.onselectionchange")
	return onselectionchange
}

func (*XMLDocument) SetOnselectionchange(onselectionchange func(event.Event)) {
	macro.Rewrite("$_.onselectionchange = $1", onselectionchange)
}

func (*XMLDocument) Onselectstart() (onselectstart func(event.Event)) {
	macro.Rewrite("$_.onselectstart")
	return onselectstart
}

func (*XMLDocument) SetOnselectstart(onselectstart func(event.Event)) {
	macro.Rewrite("$_.onselectstart = $1", onselectstart)
}

func (*XMLDocument) Onstalled() (onstalled func(event.Event)) {
	macro.Rewrite("$_.onstalled")
	return onstalled
}

func (*XMLDocument) SetOnstalled(onstalled func(event.Event)) {
	macro.Rewrite("$_.onstalled = $1", onstalled)
}

func (*XMLDocument) Onstop() (onstop func(event.Event)) {
	macro.Rewrite("$_.onstop")
	return onstop
}

func (*XMLDocument) SetOnstop(onstop func(event.Event)) {
	macro.Rewrite("$_.onstop = $1", onstop)
}

func (*XMLDocument) Onsubmit() (onsubmit func(event.Event)) {
	macro.Rewrite("$_.onsubmit")
	return onsubmit
}

func (*XMLDocument) SetOnsubmit(onsubmit func(event.Event)) {
	macro.Rewrite("$_.onsubmit = $1", onsubmit)
}

func (*XMLDocument) Onsuspend() (onsuspend func(event.Event)) {
	macro.Rewrite("$_.onsuspend")
	return onsuspend
}

func (*XMLDocument) SetOnsuspend(onsuspend func(event.Event)) {
	macro.Rewrite("$_.onsuspend = $1", onsuspend)
}

func (*XMLDocument) Ontimeupdate() (ontimeupdate func(event.Event)) {
	macro.Rewrite("$_.ontimeupdate")
	return ontimeupdate
}

func (*XMLDocument) SetOntimeupdate(ontimeupdate func(event.Event)) {
	macro.Rewrite("$_.ontimeupdate = $1", ontimeupdate)
}

func (*XMLDocument) Ontouchcancel() (ontouchcancel func(event.Event)) {
	macro.Rewrite("$_.ontouchcancel")
	return ontouchcancel
}

func (*XMLDocument) SetOntouchcancel(ontouchcancel func(event.Event)) {
	macro.Rewrite("$_.ontouchcancel = $1", ontouchcancel)
}

func (*XMLDocument) Ontouchend() (ontouchend func(event.Event)) {
	macro.Rewrite("$_.ontouchend")
	return ontouchend
}

func (*XMLDocument) SetOntouchend(ontouchend func(event.Event)) {
	macro.Rewrite("$_.ontouchend = $1", ontouchend)
}

func (*XMLDocument) Ontouchmove() (ontouchmove func(event.Event)) {
	macro.Rewrite("$_.ontouchmove")
	return ontouchmove
}

func (*XMLDocument) SetOntouchmove(ontouchmove func(event.Event)) {
	macro.Rewrite("$_.ontouchmove = $1", ontouchmove)
}

func (*XMLDocument) Ontouchstart() (ontouchstart func(event.Event)) {
	macro.Rewrite("$_.ontouchstart")
	return ontouchstart
}

func (*XMLDocument) SetOntouchstart(ontouchstart func(event.Event)) {
	macro.Rewrite("$_.ontouchstart = $1", ontouchstart)
}

func (*XMLDocument) Onvolumechange() (onvolumechange func(event.Event)) {
	macro.Rewrite("$_.onvolumechange")
	return onvolumechange
}

func (*XMLDocument) SetOnvolumechange(onvolumechange func(event.Event)) {
	macro.Rewrite("$_.onvolumechange = $1", onvolumechange)
}

func (*XMLDocument) Onwaiting() (onwaiting func(event.Event)) {
	macro.Rewrite("$_.onwaiting")
	return onwaiting
}

func (*XMLDocument) SetOnwaiting(onwaiting func(event.Event)) {
	macro.Rewrite("$_.onwaiting = $1", onwaiting)
}

func (*XMLDocument) Onwebkitfullscreenchange() (onwebkitfullscreenchange func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenchange")
	return onwebkitfullscreenchange
}

func (*XMLDocument) SetOnwebkitfullscreenchange(onwebkitfullscreenchange func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenchange = $1", onwebkitfullscreenchange)
}

func (*XMLDocument) Onwebkitfullscreenerror() (onwebkitfullscreenerror func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenerror")
	return onwebkitfullscreenerror
}

func (*XMLDocument) SetOnwebkitfullscreenerror(onwebkitfullscreenerror func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenerror = $1", onwebkitfullscreenerror)
}

func (*XMLDocument) Plugins() (plugins html.HTMLCollection) {
	macro.Rewrite("$_.plugins")
	return plugins
}

func (*XMLDocument) PointerLockElement() (pointerLockElement element.Element) {
	macro.Rewrite("$_.pointerLockElement")
	return pointerLockElement
}

func (*XMLDocument) ReadyState() (readyState string) {
	macro.Rewrite("$_.readyState")
	return readyState
}

func (*XMLDocument) Referrer() (referrer string) {
	macro.Rewrite("$_.referrer")
	return referrer
}

func (*XMLDocument) RootElement() (rootElement *dom.SVGSVGElement) {
	macro.Rewrite("$_.rootElement")
	return rootElement
}

func (*XMLDocument) Scripts() (scripts html.HTMLCollection) {
	macro.Rewrite("$_.scripts")
	return scripts
}

func (*XMLDocument) ScrollingElement() (scrollingElement element.Element) {
	macro.Rewrite("$_.scrollingElement")
	return scrollingElement
}

func (*XMLDocument) StyleSheets() (styleSheets *css.StyleSheetList) {
	macro.Rewrite("$_.styleSheets")
	return styleSheets
}

func (*XMLDocument) Title() (title string) {
	macro.Rewrite("$_.title")
	return title
}

func (*XMLDocument) SetTitle(title string) {
	macro.Rewrite("$_.title = $1", title)
}

func (*XMLDocument) URL() (URL string) {
	macro.Rewrite("$_.URL")
	return URL
}

func (*XMLDocument) URLUnencoded() (URLUnencoded string) {
	macro.Rewrite("$_.URLUnencoded")
	return URLUnencoded
}

func (*XMLDocument) VisibilityState() (visibilityState *document.VisibilityState) {
	macro.Rewrite("$_.visibilityState")
	return visibilityState
}

func (*XMLDocument) VlinkColor() (vlinkColor string) {
	macro.Rewrite("$_.vlinkColor")
	return vlinkColor
}

func (*XMLDocument) SetVlinkColor(vlinkColor string) {
	macro.Rewrite("$_.vlinkColor = $1", vlinkColor)
}

func (*XMLDocument) WebkitCurrentFullScreenElement() (webkitCurrentFullScreenElement element.Element) {
	macro.Rewrite("$_.webkitCurrentFullScreenElement")
	return webkitCurrentFullScreenElement
}

func (*XMLDocument) WebkitFullscreenElement() (webkitFullscreenElement element.Element) {
	macro.Rewrite("$_.webkitFullscreenElement")
	return webkitFullscreenElement
}

func (*XMLDocument) WebkitFullscreenEnabled() (webkitFullscreenEnabled bool) {
	macro.Rewrite("$_.webkitFullscreenEnabled")
	return webkitFullscreenEnabled
}

func (*XMLDocument) WebkitIsFullScreen() (webkitIsFullScreen bool) {
	macro.Rewrite("$_.webkitIsFullScreen")
	return webkitIsFullScreen
}

func (*XMLDocument) XMLEncoding() (xmlEncoding string) {
	macro.Rewrite("$_.xmlEncoding")
	return xmlEncoding
}

func (*XMLDocument) XMLStandalone() (xmlStandalone bool) {
	macro.Rewrite("$_.xmlStandalone")
	return xmlStandalone
}

func (*XMLDocument) SetXMLStandalone(xmlStandalone bool) {
	macro.Rewrite("$_.xmlStandalone = $1", xmlStandalone)
}

func (*XMLDocument) XMLVersion() (xmlVersion string) {
	macro.Rewrite("$_.xmlVersion")
	return xmlVersion
}

func (*XMLDocument) SetXMLVersion(xmlVersion string) {
	macro.Rewrite("$_.xmlVersion = $1", xmlVersion)
}

func (*XMLDocument) Onpointercancel() (onpointercancel func(event.Event)) {
	macro.Rewrite("$_.onpointercancel")
	return onpointercancel
}

func (*XMLDocument) SetOnpointercancel(onpointercancel func(event.Event)) {
	macro.Rewrite("$_.onpointercancel = $1", onpointercancel)
}

func (*XMLDocument) Onpointerdown() (onpointerdown func(event.Event)) {
	macro.Rewrite("$_.onpointerdown")
	return onpointerdown
}

func (*XMLDocument) SetOnpointerdown(onpointerdown func(event.Event)) {
	macro.Rewrite("$_.onpointerdown = $1", onpointerdown)
}

func (*XMLDocument) Onpointerenter() (onpointerenter func(event.Event)) {
	macro.Rewrite("$_.onpointerenter")
	return onpointerenter
}

func (*XMLDocument) SetOnpointerenter(onpointerenter func(event.Event)) {
	macro.Rewrite("$_.onpointerenter = $1", onpointerenter)
}

func (*XMLDocument) Onpointerleave() (onpointerleave func(event.Event)) {
	macro.Rewrite("$_.onpointerleave")
	return onpointerleave
}

func (*XMLDocument) SetOnpointerleave(onpointerleave func(event.Event)) {
	macro.Rewrite("$_.onpointerleave = $1", onpointerleave)
}

func (*XMLDocument) Onpointermove() (onpointermove func(event.Event)) {
	macro.Rewrite("$_.onpointermove")
	return onpointermove
}

func (*XMLDocument) SetOnpointermove(onpointermove func(event.Event)) {
	macro.Rewrite("$_.onpointermove = $1", onpointermove)
}

func (*XMLDocument) Onpointerout() (onpointerout func(event.Event)) {
	macro.Rewrite("$_.onpointerout")
	return onpointerout
}

func (*XMLDocument) SetOnpointerout(onpointerout func(event.Event)) {
	macro.Rewrite("$_.onpointerout = $1", onpointerout)
}

func (*XMLDocument) Onpointerover() (onpointerover func(event.Event)) {
	macro.Rewrite("$_.onpointerover")
	return onpointerover
}

func (*XMLDocument) SetOnpointerover(onpointerover func(event.Event)) {
	macro.Rewrite("$_.onpointerover = $1", onpointerover)
}

func (*XMLDocument) Onpointerup() (onpointerup func(event.Event)) {
	macro.Rewrite("$_.onpointerup")
	return onpointerup
}

func (*XMLDocument) SetOnpointerup(onpointerup func(event.Event)) {
	macro.Rewrite("$_.onpointerup = $1", onpointerup)
}

func (*XMLDocument) Onwheel() (onwheel func(event.Event)) {
	macro.Rewrite("$_.onwheel")
	return onwheel
}

func (*XMLDocument) SetOnwheel(onwheel func(event.Event)) {
	macro.Rewrite("$_.onwheel = $1", onwheel)
}

func (*XMLDocument) Attributes() (attributes *dom.NamedNodeMap) {
	macro.Rewrite("$_.attributes")
	return attributes
}

func (*XMLDocument) BaseURI() (baseURI string) {
	macro.Rewrite("$_.baseURI")
	return baseURI
}

func (*XMLDocument) ChildNodes() (childNodes *dom.NodeList) {
	macro.Rewrite("$_.childNodes")
	return childNodes
}

func (*XMLDocument) FirstChild() (firstChild dom.Node) {
	macro.Rewrite("$_.firstChild")
	return firstChild
}

func (*XMLDocument) LastChild() (lastChild dom.Node) {
	macro.Rewrite("$_.lastChild")
	return lastChild
}

func (*XMLDocument) LocalName() (localName string) {
	macro.Rewrite("$_.localName")
	return localName
}

func (*XMLDocument) NamespaceURI() (namespaceURI string) {
	macro.Rewrite("$_.namespaceURI")
	return namespaceURI
}

func (*XMLDocument) NextSibling() (nextSibling dom.Node) {
	macro.Rewrite("$_.nextSibling")
	return nextSibling
}

func (*XMLDocument) NodeName() (nodeName string) {
	macro.Rewrite("$_.nodeName")
	return nodeName
}

func (*XMLDocument) NodeType() (nodeType uint8) {
	macro.Rewrite("$_.nodeType")
	return nodeType
}

func (*XMLDocument) NodeValue() (nodeValue string) {
	macro.Rewrite("$_.nodeValue")
	return nodeValue
}

func (*XMLDocument) SetNodeValue(nodeValue string) {
	macro.Rewrite("$_.nodeValue = $1", nodeValue)
}

func (*XMLDocument) OwnerDocument() (ownerDocument document.Document) {
	macro.Rewrite("$_.ownerDocument")
	return ownerDocument
}

func (*XMLDocument) ParentElement() (parentElement html.HTMLElement) {
	macro.Rewrite("$_.parentElement")
	return parentElement
}

func (*XMLDocument) ParentNode() (parentNode dom.Node) {
	macro.Rewrite("$_.parentNode")
	return parentNode
}

func (*XMLDocument) PreviousSibling() (previousSibling dom.Node) {
	macro.Rewrite("$_.previousSibling")
	return previousSibling
}

func (*XMLDocument) TextContent() (textContent string) {
	macro.Rewrite("$_.textContent")
	return textContent
}

func (*XMLDocument) SetTextContent(textContent string) {
	macro.Rewrite("$_.textContent = $1", textContent)
}
