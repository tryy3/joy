package html

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
	"github.com/matthewmueller/joy/dom/url"
	"github.com/matthewmueller/joy/dom/ui"
	"github.com/matthewmueller/joy/dom/ms"
	"github.com/matthewmueller/joy/dom/css"
)

var _ document.Document = (*HTMLDocument)(nil)
var _ utils.GlobalEventHandlers = (*HTMLDocument)(nil)
var _ dom.NodeSelector = (*HTMLDocument)(nil)
var _ document.DocumentEvent = (*HTMLDocument)(nil)
var _ dom.Node = (*HTMLDocument)(nil)
var _ event.EventTarget = (*HTMLDocument)(nil)

type HTMLDocument struct {
}

func (*HTMLDocument) AdoptNode(source dom.Node) (w dom.Node) {
	macro.Rewrite("$_.adoptNode($1)", source)
	return w
}

func (*HTMLDocument) CaptureEvents() {
	macro.Rewrite("$_.captureEvents()")
}

func (*HTMLDocument) CaretRangeFromPoint(x float32, y float32) (w *dom.Range) {
	macro.Rewrite("$_.caretRangeFromPoint($1, $2)", x, y)
	return w
}

func (*HTMLDocument) Clear() {
	macro.Rewrite("$_.clear()")
}

func (*HTMLDocument) Close() {
	macro.Rewrite("$_.close()")
}

func (*HTMLDocument) CreateAttribute(name string) (w *dom.Attr) {
	macro.Rewrite("$_.createAttribute($1)", name)
	return w
}

func (*HTMLDocument) CreateAttributeNS(namespaceURI string, qualifiedName string) (w *dom.Attr) {
	macro.Rewrite("$_.createAttributeNS($1, $2)", namespaceURI, qualifiedName)
	return w
}

func (*HTMLDocument) CreateCDATASection(data string) (w *dom.CDATASection) {
	macro.Rewrite("$_.createCDATASection($1)", data)
	return w
}

func (*HTMLDocument) CreateComment(data string) (w *dom.Comment) {
	macro.Rewrite("$_.createComment($1)", data)
	return w
}

func (*HTMLDocument) CreateDocumentFragment() (w *document.DocumentFragment) {
	macro.Rewrite("$_.createDocumentFragment()")
	return w
}

func (*HTMLDocument) CreateElement(tagName string) (w element.Element) {
	macro.Rewrite("$_.createElement($1)", tagName)
	return w
}

func (*HTMLDocument) CreateElementNS(namespaceURI string, qualifiedName string) (w element.Element) {
	macro.Rewrite("$_.createElementNS($1, $2)", namespaceURI, qualifiedName)
	return w
}

func (*HTMLDocument) CreateExpression(expression string, resolver *utils.XPathNSResolver) (w *utils.XPathExpression) {
	macro.Rewrite("$_.createExpression($1, $2)", expression, resolver)
	return w
}

func (*HTMLDocument) CreateNodeIterator(root dom.Node, whatToShow *uint, filter *dom.NodeFilter, entityReferenceExpansion *bool) (w *dom.NodeIterator) {
	macro.Rewrite("$_.createNodeIterator($1, $2, $3, $4)", root, whatToShow, filter, entityReferenceExpansion)
	return w
}

func (*HTMLDocument) CreateNSResolver(nodeResolver dom.Node) (x *utils.XPathNSResolver) {
	macro.Rewrite("$_.createNSResolver($1)", nodeResolver)
	return x
}

func (*HTMLDocument) CreateProcessingInstruction(target string, data string) (w *dom.ProcessingInstruction) {
	macro.Rewrite("$_.createProcessingInstruction($1, $2)", target, data)
	return w
}

func (*HTMLDocument) CreateRange() (w *dom.Range) {
	macro.Rewrite("$_.createRange()")
	return w
}

func (*HTMLDocument) CreateTextNode(data string) (w dom.Text) {
	macro.Rewrite("$_.createTextNode($1)", data)
	return w
}

func (*HTMLDocument) CreateTouch(view *window.Window, target event.EventTarget, identifier int, pageX int, pageY int, screenX int, screenY int) (w *touch.Touch) {
	macro.Rewrite("$_.createTouch($1, $2, $3, $4, $5, $6, $7)", view, target, identifier, pageX, pageY, screenX, screenY)
	return w
}

func (*HTMLDocument) CreateTouchList(touches *touch.Touch) (w *touch.TouchList) {
	macro.Rewrite("$_.createTouchList($1)", touches)
	return w
}

func (*HTMLDocument) CreateTreeWalker(root dom.Node, whatToShow *uint, filter *dom.NodeFilter, entityReferenceExpansion *bool) (w *treewalker.TreeWalker) {
	macro.Rewrite("$_.createTreeWalker($1, $2, $3, $4)", root, whatToShow, filter, entityReferenceExpansion)
	return w
}

func (*HTMLDocument) ElementFromPoint(x int, y int) (w element.Element) {
	macro.Rewrite("$_.elementFromPoint($1, $2)", x, y)
	return w
}

func (*HTMLDocument) Evaluate(expression string, contextNode dom.Node, resolver *utils.XPathNSResolver, kind uint8, result *utils.XPathResult) (w *utils.XPathResult) {
	macro.Rewrite("$_.evaluate($1, $2, $3, $4, $5)", expression, contextNode, resolver, kind, result)
	return w
}

func (*HTMLDocument) ExecCommand(commandId string, showUI *bool, value *interface{}) (b bool) {
	macro.Rewrite("$_.execCommand($1, $2, $3)", commandId, showUI, value)
	return b
}

func (*HTMLDocument) ExecCommandShowHelp(commandId string) (b bool) {
	macro.Rewrite("$_.execCommandShowHelp($1)", commandId)
	return b
}

func (*HTMLDocument) ExitFullscreen() {
	macro.Rewrite("$_.exitFullscreen()")
}

func (*HTMLDocument) ExitPointerLock() {
	macro.Rewrite("$_.exitPointerLock()")
}

func (*HTMLDocument) Focus() {
	macro.Rewrite("$_.focus()")
}

func (*HTMLDocument) GetElementByID(elementId string) (w element.Element) {
	macro.Rewrite("$_.getElementById($1)", elementId)
	return w
}

func (*HTMLDocument) GetElementsByClassName(classNames string) (w *dom.NodeList) {
	macro.Rewrite("$_.getElementsByClassName($1)", classNames)
	return w
}

func (*HTMLDocument) GetElementsByName(elementName string) (w *dom.NodeList) {
	macro.Rewrite("$_.getElementsByName($1)", elementName)
	return w
}

func (*HTMLDocument) GetElementsByTagName(tagname string) (w *dom.NodeList) {
	macro.Rewrite("$_.getElementsByTagName($1)", tagname)
	return w
}

func (*HTMLDocument) GetElementsByTagNameNS(namespaceURI string, localName string) (w *dom.NodeList) {
	macro.Rewrite("$_.getElementsByTagNameNS($1, $2)", namespaceURI, localName)
	return w
}

func (*HTMLDocument) GetSelection() (w *dom.Selection) {
	macro.Rewrite("$_.getSelection()")
	return w
}

func (*HTMLDocument) HasFocus() (b bool) {
	macro.Rewrite("$_.hasFocus()")
	return b
}

func (*HTMLDocument) ImportNode(importedNode dom.Node, deep bool) (w dom.Node) {
	macro.Rewrite("$_.importNode($1, $2)", importedNode, deep)
	return w
}

func (*HTMLDocument) MsElementsFromPoint(x float32, y float32) (w *dom.NodeList) {
	macro.Rewrite("$_.msElementsFromPoint($1, $2)", x, y)
	return w
}

func (*HTMLDocument) MsElementsFromRect(left float32, top float32, width float32, height float32) (w *dom.NodeList) {
	macro.Rewrite("$_.msElementsFromRect($1, $2, $3, $4)", left, top, width, height)
	return w
}

func (*HTMLDocument) Open(url *string, name *string, features *string, replace *bool) (i interface{}) {
	macro.Rewrite("$_.open($1, $2, $3, $4)", url, name, features, replace)
	return i
}

func (*HTMLDocument) QueryCommandEnabled(commandId string) (b bool) {
	macro.Rewrite("$_.queryCommandEnabled($1)", commandId)
	return b
}

func (*HTMLDocument) QueryCommandIndeterm(commandId string) (b bool) {
	macro.Rewrite("$_.queryCommandIndeterm($1)", commandId)
	return b
}

func (*HTMLDocument) QueryCommandState(commandId string) (b bool) {
	macro.Rewrite("$_.queryCommandState($1)", commandId)
	return b
}

func (*HTMLDocument) QueryCommandSupported(commandId string) (b bool) {
	macro.Rewrite("$_.queryCommandSupported($1)", commandId)
	return b
}

func (*HTMLDocument) QueryCommandText(commandId string) (s string) {
	macro.Rewrite("$_.queryCommandText($1)", commandId)
	return s
}

func (*HTMLDocument) QueryCommandValue(commandId string) (s string) {
	macro.Rewrite("$_.queryCommandValue($1)", commandId)
	return s
}

func (*HTMLDocument) ReleaseEvents() {
	macro.Rewrite("$_.releaseEvents()")
}

func (*HTMLDocument) UpdateSettings() {
	macro.Rewrite("$_.updateSettings()")
}

func (*HTMLDocument) WebkitCancelFullScreen() {
	macro.Rewrite("$_.webkitCancelFullScreen()")
}

func (*HTMLDocument) WebkitExitFullscreen() {
	macro.Rewrite("$_.webkitExitFullscreen()")
}

func (*HTMLDocument) Write(content string) {
	macro.Rewrite("$_.write($1)", content)
}

func (*HTMLDocument) Writeln(content string) {
	macro.Rewrite("$_.writeln($1)", content)
}

func (*HTMLDocument) QuerySelector(selectors string) (w element.Element) {
	macro.Rewrite("$_.querySelector($1)", selectors)
	return w
}

func (*HTMLDocument) QuerySelectorAll(selectors string) (w *dom.NodeList) {
	macro.Rewrite("$_.querySelectorAll($1)", selectors)
	return w
}

func (*HTMLDocument) CreateEvent(eventInterface string) (w event.Event) {
	macro.Rewrite("$_.createEvent($1)", eventInterface)
	return w
}

func (*HTMLDocument) AppendChild(newChild dom.Node) (w dom.Node) {
	macro.Rewrite("$_.appendChild($1)", newChild)
	return w
}

func (*HTMLDocument) CloneNode(deep *bool) (w dom.Node) {
	macro.Rewrite("$_.cloneNode($1)", deep)
	return w
}

func (*HTMLDocument) CompareDocumentPosition(other dom.Node) (u uint8) {
	macro.Rewrite("$_.compareDocumentPosition($1)", other)
	return u
}

func (*HTMLDocument) Contains(child dom.Node) (b bool) {
	macro.Rewrite("$_.contains($1)", child)
	return b
}

func (*HTMLDocument) HasAttributes() (b bool) {
	macro.Rewrite("$_.hasAttributes()")
	return b
}

func (*HTMLDocument) HasChildNodes() (b bool) {
	macro.Rewrite("$_.hasChildNodes()")
	return b
}

func (*HTMLDocument) InsertBefore(newChild dom.Node, refChild *dom.Node) (w dom.Node) {
	macro.Rewrite("$_.insertBefore($1, $2)", newChild, refChild)
	return w
}

func (*HTMLDocument) IsDefaultNamespace(namespaceURI string) (b bool) {
	macro.Rewrite("$_.isDefaultNamespace($1)", namespaceURI)
	return b
}

func (*HTMLDocument) IsEqualNode(arg dom.Node) (b bool) {
	macro.Rewrite("$_.isEqualNode($1)", arg)
	return b
}

func (*HTMLDocument) IsSameNode(other dom.Node) (b bool) {
	macro.Rewrite("$_.isSameNode($1)", other)
	return b
}

func (*HTMLDocument) LookupNamespaceURI(prefix string) (s string) {
	macro.Rewrite("$_.lookupNamespaceURI($1)", prefix)
	return s
}

func (*HTMLDocument) LookupPrefix(namespaceURI string) (s string) {
	macro.Rewrite("$_.lookupPrefix($1)", namespaceURI)
	return s
}

func (*HTMLDocument) Normalize() {
	macro.Rewrite("$_.normalize()")
}

func (*HTMLDocument) RemoveChild(oldChild dom.Node) (w dom.Node) {
	macro.Rewrite("$_.removeChild($1)", oldChild)
	return w
}

func (*HTMLDocument) ReplaceChild(newChild dom.Node, oldChild dom.Node) (w dom.Node) {
	macro.Rewrite("$_.replaceChild($1, $2)", newChild, oldChild)
	return w
}

func (*HTMLDocument) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*HTMLDocument) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*HTMLDocument) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*HTMLDocument) ActiveElement() (activeElement element.Element) {
	macro.Rewrite("$_.activeElement")
	return activeElement
}

func (*HTMLDocument) AlinkColor() (alinkColor string) {
	macro.Rewrite("$_.alinkColor")
	return alinkColor
}

func (*HTMLDocument) SetAlinkColor(alinkColor string) {
	macro.Rewrite("$_.alinkColor = $1", alinkColor)
}

func (*HTMLDocument) All() (all *HTMLAllCollection) {
	macro.Rewrite("$_.all")
	return all
}

func (*HTMLDocument) Anchors() (anchors HTMLCollection) {
	macro.Rewrite("$_.anchors")
	return anchors
}

func (*HTMLDocument) Applets() (applets HTMLCollection) {
	macro.Rewrite("$_.applets")
	return applets
}

func (*HTMLDocument) BgColor() (bgColor string) {
	macro.Rewrite("$_.bgColor")
	return bgColor
}

func (*HTMLDocument) SetBgColor(bgColor string) {
	macro.Rewrite("$_.bgColor = $1", bgColor)
}

func (*HTMLDocument) Body() (body HTMLElement) {
	macro.Rewrite("$_.body")
	return body
}

func (*HTMLDocument) SetBody(body HTMLElement) {
	macro.Rewrite("$_.body = $1", body)
}

func (*HTMLDocument) CharacterSet() (characterSet string) {
	macro.Rewrite("$_.characterSet")
	return characterSet
}

func (*HTMLDocument) Charset() (charset string) {
	macro.Rewrite("$_.charset")
	return charset
}

func (*HTMLDocument) SetCharset(charset string) {
	macro.Rewrite("$_.charset = $1", charset)
}

func (*HTMLDocument) CompatMode() (compatMode string) {
	macro.Rewrite("$_.compatMode")
	return compatMode
}

func (*HTMLDocument) Cookie() (cookie string) {
	macro.Rewrite("$_.cookie")
	return cookie
}

func (*HTMLDocument) SetCookie(cookie string) {
	macro.Rewrite("$_.cookie = $1", cookie)
}

func (*HTMLDocument) CurrentScript() (currentScript interface{}) {
	macro.Rewrite("$_.currentScript")
	return currentScript
}

func (*HTMLDocument) DefaultView() (defaultView *window.Window) {
	macro.Rewrite("$_.defaultView")
	return defaultView
}

func (*HTMLDocument) DesignMode() (designMode string) {
	macro.Rewrite("$_.designMode")
	return designMode
}

func (*HTMLDocument) SetDesignMode(designMode string) {
	macro.Rewrite("$_.designMode = $1", designMode)
}

func (*HTMLDocument) Dir() (dir string) {
	macro.Rewrite("$_.dir")
	return dir
}

func (*HTMLDocument) SetDir(dir string) {
	macro.Rewrite("$_.dir = $1", dir)
}

func (*HTMLDocument) Doctype() (doctype *document.DocumentType) {
	macro.Rewrite("$_.doctype")
	return doctype
}

func (*HTMLDocument) DocumentElement() (documentElement element.Element) {
	macro.Rewrite("$_.documentElement")
	return documentElement
}

func (*HTMLDocument) Domain() (domain string) {
	macro.Rewrite("$_.domain")
	return domain
}

func (*HTMLDocument) SetDomain(domain string) {
	macro.Rewrite("$_.domain = $1", domain)
}

func (*HTMLDocument) Embeds() (embeds HTMLCollection) {
	macro.Rewrite("$_.embeds")
	return embeds
}

func (*HTMLDocument) FgColor() (fgColor string) {
	macro.Rewrite("$_.fgColor")
	return fgColor
}

func (*HTMLDocument) SetFgColor(fgColor string) {
	macro.Rewrite("$_.fgColor = $1", fgColor)
}

func (*HTMLDocument) Forms() (forms HTMLCollection) {
	macro.Rewrite("$_.forms")
	return forms
}

func (*HTMLDocument) FullscreenElement() (fullscreenElement element.Element) {
	macro.Rewrite("$_.fullscreenElement")
	return fullscreenElement
}

func (*HTMLDocument) FullscreenEnabled() (fullscreenEnabled bool) {
	macro.Rewrite("$_.fullscreenEnabled")
	return fullscreenEnabled
}

func (*HTMLDocument) Head() (head *HTMLHeadElement) {
	macro.Rewrite("$_.head")
	return head
}

func (*HTMLDocument) Hidden() (hidden bool) {
	macro.Rewrite("$_.hidden")
	return hidden
}

func (*HTMLDocument) Images() (images HTMLCollection) {
	macro.Rewrite("$_.images")
	return images
}

func (*HTMLDocument) Implementation() (implementation *dom.DOMImplementation) {
	macro.Rewrite("$_.implementation")
	return implementation
}

func (*HTMLDocument) InputEncoding() (inputEncoding string) {
	macro.Rewrite("$_.inputEncoding")
	return inputEncoding
}

func (*HTMLDocument) LastModified() (lastModified string) {
	macro.Rewrite("$_.lastModified")
	return lastModified
}

func (*HTMLDocument) LinkColor() (linkColor string) {
	macro.Rewrite("$_.linkColor")
	return linkColor
}

func (*HTMLDocument) SetLinkColor(linkColor string) {
	macro.Rewrite("$_.linkColor = $1", linkColor)
}

func (*HTMLDocument) Links() (links HTMLCollection) {
	macro.Rewrite("$_.links")
	return links
}

func (*HTMLDocument) Location() (location *url.Location) {
	macro.Rewrite("$_.location")
	return location
}

func (*HTMLDocument) MsCapsLockWarningOff() (msCapsLockWarningOff bool) {
	macro.Rewrite("$_.msCapsLockWarningOff")
	return msCapsLockWarningOff
}

func (*HTMLDocument) SetMsCapsLockWarningOff(msCapsLockWarningOff bool) {
	macro.Rewrite("$_.msCapsLockWarningOff = $1", msCapsLockWarningOff)
}

func (*HTMLDocument) MsCSSOMElementFloatMetrics() (msCSSOMElementFloatMetrics bool) {
	macro.Rewrite("$_.msCSSOMElementFloatMetrics")
	return msCSSOMElementFloatMetrics
}

func (*HTMLDocument) SetMsCSSOMElementFloatMetrics(msCSSOMElementFloatMetrics bool) {
	macro.Rewrite("$_.msCSSOMElementFloatMetrics = $1", msCSSOMElementFloatMetrics)
}

func (*HTMLDocument) Onabort() (onabort func(event.Event)) {
	macro.Rewrite("$_.onabort")
	return onabort
}

func (*HTMLDocument) SetOnabort(onabort func(event.Event)) {
	macro.Rewrite("$_.onabort = $1", onabort)
}

func (*HTMLDocument) Onactivate() (onactivate func(event.Event)) {
	macro.Rewrite("$_.onactivate")
	return onactivate
}

func (*HTMLDocument) SetOnactivate(onactivate func(event.Event)) {
	macro.Rewrite("$_.onactivate = $1", onactivate)
}

func (*HTMLDocument) Onbeforeactivate() (onbeforeactivate func(event.Event)) {
	macro.Rewrite("$_.onbeforeactivate")
	return onbeforeactivate
}

func (*HTMLDocument) SetOnbeforeactivate(onbeforeactivate func(event.Event)) {
	macro.Rewrite("$_.onbeforeactivate = $1", onbeforeactivate)
}

func (*HTMLDocument) Onbeforedeactivate() (onbeforedeactivate func(event.Event)) {
	macro.Rewrite("$_.onbeforedeactivate")
	return onbeforedeactivate
}

func (*HTMLDocument) SetOnbeforedeactivate(onbeforedeactivate func(event.Event)) {
	macro.Rewrite("$_.onbeforedeactivate = $1", onbeforedeactivate)
}

func (*HTMLDocument) Onblur() (onblur func(event.Event)) {
	macro.Rewrite("$_.onblur")
	return onblur
}

func (*HTMLDocument) SetOnblur(onblur func(event.Event)) {
	macro.Rewrite("$_.onblur = $1", onblur)
}

func (*HTMLDocument) Oncanplay() (oncanplay func(event.Event)) {
	macro.Rewrite("$_.oncanplay")
	return oncanplay
}

func (*HTMLDocument) SetOncanplay(oncanplay func(event.Event)) {
	macro.Rewrite("$_.oncanplay = $1", oncanplay)
}

func (*HTMLDocument) Oncanplaythrough() (oncanplaythrough func(event.Event)) {
	macro.Rewrite("$_.oncanplaythrough")
	return oncanplaythrough
}

func (*HTMLDocument) SetOncanplaythrough(oncanplaythrough func(event.Event)) {
	macro.Rewrite("$_.oncanplaythrough = $1", oncanplaythrough)
}

func (*HTMLDocument) Onchange() (onchange func(event.Event)) {
	macro.Rewrite("$_.onchange")
	return onchange
}

func (*HTMLDocument) SetOnchange(onchange func(event.Event)) {
	macro.Rewrite("$_.onchange = $1", onchange)
}

func (*HTMLDocument) Onclick() (onclick func(event.Event)) {
	macro.Rewrite("$_.onclick")
	return onclick
}

func (*HTMLDocument) SetOnclick(onclick func(event.Event)) {
	macro.Rewrite("$_.onclick = $1", onclick)
}

func (*HTMLDocument) Oncontextmenu() (oncontextmenu func(event.Event)) {
	macro.Rewrite("$_.oncontextmenu")
	return oncontextmenu
}

func (*HTMLDocument) SetOncontextmenu(oncontextmenu func(event.Event)) {
	macro.Rewrite("$_.oncontextmenu = $1", oncontextmenu)
}

func (*HTMLDocument) Ondblclick() (ondblclick func(event.Event)) {
	macro.Rewrite("$_.ondblclick")
	return ondblclick
}

func (*HTMLDocument) SetOndblclick(ondblclick func(event.Event)) {
	macro.Rewrite("$_.ondblclick = $1", ondblclick)
}

func (*HTMLDocument) Ondeactivate() (ondeactivate func(event.Event)) {
	macro.Rewrite("$_.ondeactivate")
	return ondeactivate
}

func (*HTMLDocument) SetOndeactivate(ondeactivate func(event.Event)) {
	macro.Rewrite("$_.ondeactivate = $1", ondeactivate)
}

func (*HTMLDocument) Ondrag() (ondrag func(event.Event)) {
	macro.Rewrite("$_.ondrag")
	return ondrag
}

func (*HTMLDocument) SetOndrag(ondrag func(event.Event)) {
	macro.Rewrite("$_.ondrag = $1", ondrag)
}

func (*HTMLDocument) Ondragend() (ondragend func(event.Event)) {
	macro.Rewrite("$_.ondragend")
	return ondragend
}

func (*HTMLDocument) SetOndragend(ondragend func(event.Event)) {
	macro.Rewrite("$_.ondragend = $1", ondragend)
}

func (*HTMLDocument) Ondragenter() (ondragenter func(event.Event)) {
	macro.Rewrite("$_.ondragenter")
	return ondragenter
}

func (*HTMLDocument) SetOndragenter(ondragenter func(event.Event)) {
	macro.Rewrite("$_.ondragenter = $1", ondragenter)
}

func (*HTMLDocument) Ondragleave() (ondragleave func(event.Event)) {
	macro.Rewrite("$_.ondragleave")
	return ondragleave
}

func (*HTMLDocument) SetOndragleave(ondragleave func(event.Event)) {
	macro.Rewrite("$_.ondragleave = $1", ondragleave)
}

func (*HTMLDocument) Ondragover() (ondragover func(event.Event)) {
	macro.Rewrite("$_.ondragover")
	return ondragover
}

func (*HTMLDocument) SetOndragover(ondragover func(event.Event)) {
	macro.Rewrite("$_.ondragover = $1", ondragover)
}

func (*HTMLDocument) Ondragstart() (ondragstart func(event.Event)) {
	macro.Rewrite("$_.ondragstart")
	return ondragstart
}

func (*HTMLDocument) SetOndragstart(ondragstart func(event.Event)) {
	macro.Rewrite("$_.ondragstart = $1", ondragstart)
}

func (*HTMLDocument) Ondrop() (ondrop func(event.Event)) {
	macro.Rewrite("$_.ondrop")
	return ondrop
}

func (*HTMLDocument) SetOndrop(ondrop func(event.Event)) {
	macro.Rewrite("$_.ondrop = $1", ondrop)
}

func (*HTMLDocument) Ondurationchange() (ondurationchange func(event.Event)) {
	macro.Rewrite("$_.ondurationchange")
	return ondurationchange
}

func (*HTMLDocument) SetOndurationchange(ondurationchange func(event.Event)) {
	macro.Rewrite("$_.ondurationchange = $1", ondurationchange)
}

func (*HTMLDocument) Onemptied() (onemptied func(event.Event)) {
	macro.Rewrite("$_.onemptied")
	return onemptied
}

func (*HTMLDocument) SetOnemptied(onemptied func(event.Event)) {
	macro.Rewrite("$_.onemptied = $1", onemptied)
}

func (*HTMLDocument) Onended() (onended func(event.Event)) {
	macro.Rewrite("$_.onended")
	return onended
}

func (*HTMLDocument) SetOnended(onended func(event.Event)) {
	macro.Rewrite("$_.onended = $1", onended)
}

func (*HTMLDocument) Onerror() (onerror func(event.Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

func (*HTMLDocument) SetOnerror(onerror func(event.Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

func (*HTMLDocument) Onfocus() (onfocus func(event.Event)) {
	macro.Rewrite("$_.onfocus")
	return onfocus
}

func (*HTMLDocument) SetOnfocus(onfocus func(event.Event)) {
	macro.Rewrite("$_.onfocus = $1", onfocus)
}

func (*HTMLDocument) Onfullscreenchange() (onfullscreenchange func(event.Event)) {
	macro.Rewrite("$_.onfullscreenchange")
	return onfullscreenchange
}

func (*HTMLDocument) SetOnfullscreenchange(onfullscreenchange func(event.Event)) {
	macro.Rewrite("$_.onfullscreenchange = $1", onfullscreenchange)
}

func (*HTMLDocument) Onfullscreenerror() (onfullscreenerror func(event.Event)) {
	macro.Rewrite("$_.onfullscreenerror")
	return onfullscreenerror
}

func (*HTMLDocument) SetOnfullscreenerror(onfullscreenerror func(event.Event)) {
	macro.Rewrite("$_.onfullscreenerror = $1", onfullscreenerror)
}

func (*HTMLDocument) Oninput() (oninput func(event.Event)) {
	macro.Rewrite("$_.oninput")
	return oninput
}

func (*HTMLDocument) SetOninput(oninput func(event.Event)) {
	macro.Rewrite("$_.oninput = $1", oninput)
}

func (*HTMLDocument) Oninvalid() (oninvalid func(event.Event)) {
	macro.Rewrite("$_.oninvalid")
	return oninvalid
}

func (*HTMLDocument) SetOninvalid(oninvalid func(event.Event)) {
	macro.Rewrite("$_.oninvalid = $1", oninvalid)
}

func (*HTMLDocument) Onkeydown() (onkeydown func(event.Event)) {
	macro.Rewrite("$_.onkeydown")
	return onkeydown
}

func (*HTMLDocument) SetOnkeydown(onkeydown func(event.Event)) {
	macro.Rewrite("$_.onkeydown = $1", onkeydown)
}

func (*HTMLDocument) Onkeypress() (onkeypress func(event.Event)) {
	macro.Rewrite("$_.onkeypress")
	return onkeypress
}

func (*HTMLDocument) SetOnkeypress(onkeypress func(event.Event)) {
	macro.Rewrite("$_.onkeypress = $1", onkeypress)
}

func (*HTMLDocument) Onkeyup() (onkeyup func(event.Event)) {
	macro.Rewrite("$_.onkeyup")
	return onkeyup
}

func (*HTMLDocument) SetOnkeyup(onkeyup func(event.Event)) {
	macro.Rewrite("$_.onkeyup = $1", onkeyup)
}

func (*HTMLDocument) Onload() (onload func(event.Event)) {
	macro.Rewrite("$_.onload")
	return onload
}

func (*HTMLDocument) SetOnload(onload func(event.Event)) {
	macro.Rewrite("$_.onload = $1", onload)
}

func (*HTMLDocument) Onloadeddata() (onloadeddata func(event.Event)) {
	macro.Rewrite("$_.onloadeddata")
	return onloadeddata
}

func (*HTMLDocument) SetOnloadeddata(onloadeddata func(event.Event)) {
	macro.Rewrite("$_.onloadeddata = $1", onloadeddata)
}

func (*HTMLDocument) Onloadedmetadata() (onloadedmetadata func(event.Event)) {
	macro.Rewrite("$_.onloadedmetadata")
	return onloadedmetadata
}

func (*HTMLDocument) SetOnloadedmetadata(onloadedmetadata func(event.Event)) {
	macro.Rewrite("$_.onloadedmetadata = $1", onloadedmetadata)
}

func (*HTMLDocument) Onloadstart() (onloadstart func(event.Event)) {
	macro.Rewrite("$_.onloadstart")
	return onloadstart
}

func (*HTMLDocument) SetOnloadstart(onloadstart func(event.Event)) {
	macro.Rewrite("$_.onloadstart = $1", onloadstart)
}

func (*HTMLDocument) Onmousedown() (onmousedown func(event.Event)) {
	macro.Rewrite("$_.onmousedown")
	return onmousedown
}

func (*HTMLDocument) SetOnmousedown(onmousedown func(event.Event)) {
	macro.Rewrite("$_.onmousedown = $1", onmousedown)
}

func (*HTMLDocument) Onmousemove() (onmousemove func(event.Event)) {
	macro.Rewrite("$_.onmousemove")
	return onmousemove
}

func (*HTMLDocument) SetOnmousemove(onmousemove func(event.Event)) {
	macro.Rewrite("$_.onmousemove = $1", onmousemove)
}

func (*HTMLDocument) Onmouseout() (onmouseout func(event.Event)) {
	macro.Rewrite("$_.onmouseout")
	return onmouseout
}

func (*HTMLDocument) SetOnmouseout(onmouseout func(event.Event)) {
	macro.Rewrite("$_.onmouseout = $1", onmouseout)
}

func (*HTMLDocument) Onmouseover() (onmouseover func(event.Event)) {
	macro.Rewrite("$_.onmouseover")
	return onmouseover
}

func (*HTMLDocument) SetOnmouseover(onmouseover func(event.Event)) {
	macro.Rewrite("$_.onmouseover = $1", onmouseover)
}

func (*HTMLDocument) Onmouseup() (onmouseup func(event.Event)) {
	macro.Rewrite("$_.onmouseup")
	return onmouseup
}

func (*HTMLDocument) SetOnmouseup(onmouseup func(event.Event)) {
	macro.Rewrite("$_.onmouseup = $1", onmouseup)
}

func (*HTMLDocument) Onmousewheel() (onmousewheel func(event.Event)) {
	macro.Rewrite("$_.onmousewheel")
	return onmousewheel
}

func (*HTMLDocument) SetOnmousewheel(onmousewheel func(event.Event)) {
	macro.Rewrite("$_.onmousewheel = $1", onmousewheel)
}

func (*HTMLDocument) Onmscontentzoom() (onmscontentzoom func(ui.UIEvent)) {
	macro.Rewrite("$_.onmscontentzoom")
	return onmscontentzoom
}

func (*HTMLDocument) SetOnmscontentzoom(onmscontentzoom func(ui.UIEvent)) {
	macro.Rewrite("$_.onmscontentzoom = $1", onmscontentzoom)
}

func (*HTMLDocument) Onmsgesturechange() (onmsgesturechange func(event.Event)) {
	macro.Rewrite("$_.onmsgesturechange")
	return onmsgesturechange
}

func (*HTMLDocument) SetOnmsgesturechange(onmsgesturechange func(event.Event)) {
	macro.Rewrite("$_.onmsgesturechange = $1", onmsgesturechange)
}

func (*HTMLDocument) Onmsgesturedoubletap() (onmsgesturedoubletap func(event.Event)) {
	macro.Rewrite("$_.onmsgesturedoubletap")
	return onmsgesturedoubletap
}

func (*HTMLDocument) SetOnmsgesturedoubletap(onmsgesturedoubletap func(event.Event)) {
	macro.Rewrite("$_.onmsgesturedoubletap = $1", onmsgesturedoubletap)
}

func (*HTMLDocument) Onmsgestureend() (onmsgestureend func(event.Event)) {
	macro.Rewrite("$_.onmsgestureend")
	return onmsgestureend
}

func (*HTMLDocument) SetOnmsgestureend(onmsgestureend func(event.Event)) {
	macro.Rewrite("$_.onmsgestureend = $1", onmsgestureend)
}

func (*HTMLDocument) Onmsgesturehold() (onmsgesturehold func(event.Event)) {
	macro.Rewrite("$_.onmsgesturehold")
	return onmsgesturehold
}

func (*HTMLDocument) SetOnmsgesturehold(onmsgesturehold func(event.Event)) {
	macro.Rewrite("$_.onmsgesturehold = $1", onmsgesturehold)
}

func (*HTMLDocument) Onmsgesturestart() (onmsgesturestart func(event.Event)) {
	macro.Rewrite("$_.onmsgesturestart")
	return onmsgesturestart
}

func (*HTMLDocument) SetOnmsgesturestart(onmsgesturestart func(event.Event)) {
	macro.Rewrite("$_.onmsgesturestart = $1", onmsgesturestart)
}

func (*HTMLDocument) Onmsgesturetap() (onmsgesturetap func(event.Event)) {
	macro.Rewrite("$_.onmsgesturetap")
	return onmsgesturetap
}

func (*HTMLDocument) SetOnmsgesturetap(onmsgesturetap func(event.Event)) {
	macro.Rewrite("$_.onmsgesturetap = $1", onmsgesturetap)
}

func (*HTMLDocument) Onmsinertiastart() (onmsinertiastart func(event.Event)) {
	macro.Rewrite("$_.onmsinertiastart")
	return onmsinertiastart
}

func (*HTMLDocument) SetOnmsinertiastart(onmsinertiastart func(event.Event)) {
	macro.Rewrite("$_.onmsinertiastart = $1", onmsinertiastart)
}

func (*HTMLDocument) Onmsmanipulationstatechanged() (onmsmanipulationstatechanged func(*ms.MSManipulationEvent)) {
	macro.Rewrite("$_.onmsmanipulationstatechanged")
	return onmsmanipulationstatechanged
}

func (*HTMLDocument) SetOnmsmanipulationstatechanged(onmsmanipulationstatechanged func(*ms.MSManipulationEvent)) {
	macro.Rewrite("$_.onmsmanipulationstatechanged = $1", onmsmanipulationstatechanged)
}

func (*HTMLDocument) Onmspointercancel() (onmspointercancel func(event.Event)) {
	macro.Rewrite("$_.onmspointercancel")
	return onmspointercancel
}

func (*HTMLDocument) SetOnmspointercancel(onmspointercancel func(event.Event)) {
	macro.Rewrite("$_.onmspointercancel = $1", onmspointercancel)
}

func (*HTMLDocument) Onmspointerdown() (onmspointerdown func(event.Event)) {
	macro.Rewrite("$_.onmspointerdown")
	return onmspointerdown
}

func (*HTMLDocument) SetOnmspointerdown(onmspointerdown func(event.Event)) {
	macro.Rewrite("$_.onmspointerdown = $1", onmspointerdown)
}

func (*HTMLDocument) Onmspointerenter() (onmspointerenter func(event.Event)) {
	macro.Rewrite("$_.onmspointerenter")
	return onmspointerenter
}

func (*HTMLDocument) SetOnmspointerenter(onmspointerenter func(event.Event)) {
	macro.Rewrite("$_.onmspointerenter = $1", onmspointerenter)
}

func (*HTMLDocument) Onmspointerleave() (onmspointerleave func(event.Event)) {
	macro.Rewrite("$_.onmspointerleave")
	return onmspointerleave
}

func (*HTMLDocument) SetOnmspointerleave(onmspointerleave func(event.Event)) {
	macro.Rewrite("$_.onmspointerleave = $1", onmspointerleave)
}

func (*HTMLDocument) Onmspointermove() (onmspointermove func(event.Event)) {
	macro.Rewrite("$_.onmspointermove")
	return onmspointermove
}

func (*HTMLDocument) SetOnmspointermove(onmspointermove func(event.Event)) {
	macro.Rewrite("$_.onmspointermove = $1", onmspointermove)
}

func (*HTMLDocument) Onmspointerout() (onmspointerout func(event.Event)) {
	macro.Rewrite("$_.onmspointerout")
	return onmspointerout
}

func (*HTMLDocument) SetOnmspointerout(onmspointerout func(event.Event)) {
	macro.Rewrite("$_.onmspointerout = $1", onmspointerout)
}

func (*HTMLDocument) Onmspointerover() (onmspointerover func(event.Event)) {
	macro.Rewrite("$_.onmspointerover")
	return onmspointerover
}

func (*HTMLDocument) SetOnmspointerover(onmspointerover func(event.Event)) {
	macro.Rewrite("$_.onmspointerover = $1", onmspointerover)
}

func (*HTMLDocument) Onmspointerup() (onmspointerup func(event.Event)) {
	macro.Rewrite("$_.onmspointerup")
	return onmspointerup
}

func (*HTMLDocument) SetOnmspointerup(onmspointerup func(event.Event)) {
	macro.Rewrite("$_.onmspointerup = $1", onmspointerup)
}

func (*HTMLDocument) Onmssitemodejumplistitemremoved() (onmssitemodejumplistitemremoved func(*ms.MSSiteModeEvent)) {
	macro.Rewrite("$_.onmssitemodejumplistitemremoved")
	return onmssitemodejumplistitemremoved
}

func (*HTMLDocument) SetOnmssitemodejumplistitemremoved(onmssitemodejumplistitemremoved func(*ms.MSSiteModeEvent)) {
	macro.Rewrite("$_.onmssitemodejumplistitemremoved = $1", onmssitemodejumplistitemremoved)
}

func (*HTMLDocument) Onmsthumbnailclick() (onmsthumbnailclick func(*ms.MSSiteModeEvent)) {
	macro.Rewrite("$_.onmsthumbnailclick")
	return onmsthumbnailclick
}

func (*HTMLDocument) SetOnmsthumbnailclick(onmsthumbnailclick func(*ms.MSSiteModeEvent)) {
	macro.Rewrite("$_.onmsthumbnailclick = $1", onmsthumbnailclick)
}

func (*HTMLDocument) Onpause() (onpause func(event.Event)) {
	macro.Rewrite("$_.onpause")
	return onpause
}

func (*HTMLDocument) SetOnpause(onpause func(event.Event)) {
	macro.Rewrite("$_.onpause = $1", onpause)
}

func (*HTMLDocument) Onplay() (onplay func(event.Event)) {
	macro.Rewrite("$_.onplay")
	return onplay
}

func (*HTMLDocument) SetOnplay(onplay func(event.Event)) {
	macro.Rewrite("$_.onplay = $1", onplay)
}

func (*HTMLDocument) Onplaying() (onplaying func(event.Event)) {
	macro.Rewrite("$_.onplaying")
	return onplaying
}

func (*HTMLDocument) SetOnplaying(onplaying func(event.Event)) {
	macro.Rewrite("$_.onplaying = $1", onplaying)
}

func (*HTMLDocument) Onpointerlockchange() (onpointerlockchange func(event.Event)) {
	macro.Rewrite("$_.onpointerlockchange")
	return onpointerlockchange
}

func (*HTMLDocument) SetOnpointerlockchange(onpointerlockchange func(event.Event)) {
	macro.Rewrite("$_.onpointerlockchange = $1", onpointerlockchange)
}

func (*HTMLDocument) Onpointerlockerror() (onpointerlockerror func(event.Event)) {
	macro.Rewrite("$_.onpointerlockerror")
	return onpointerlockerror
}

func (*HTMLDocument) SetOnpointerlockerror(onpointerlockerror func(event.Event)) {
	macro.Rewrite("$_.onpointerlockerror = $1", onpointerlockerror)
}

func (*HTMLDocument) Onprogress() (onprogress func(event.Event)) {
	macro.Rewrite("$_.onprogress")
	return onprogress
}

func (*HTMLDocument) SetOnprogress(onprogress func(event.Event)) {
	macro.Rewrite("$_.onprogress = $1", onprogress)
}

func (*HTMLDocument) Onratechange() (onratechange func(event.Event)) {
	macro.Rewrite("$_.onratechange")
	return onratechange
}

func (*HTMLDocument) SetOnratechange(onratechange func(event.Event)) {
	macro.Rewrite("$_.onratechange = $1", onratechange)
}

func (*HTMLDocument) Onreadystatechange() (onreadystatechange func(event.Event)) {
	macro.Rewrite("$_.onreadystatechange")
	return onreadystatechange
}

func (*HTMLDocument) SetOnreadystatechange(onreadystatechange func(event.Event)) {
	macro.Rewrite("$_.onreadystatechange = $1", onreadystatechange)
}

func (*HTMLDocument) Onreset() (onreset func(event.Event)) {
	macro.Rewrite("$_.onreset")
	return onreset
}

func (*HTMLDocument) SetOnreset(onreset func(event.Event)) {
	macro.Rewrite("$_.onreset = $1", onreset)
}

func (*HTMLDocument) Onscroll() (onscroll func(event.Event)) {
	macro.Rewrite("$_.onscroll")
	return onscroll
}

func (*HTMLDocument) SetOnscroll(onscroll func(event.Event)) {
	macro.Rewrite("$_.onscroll = $1", onscroll)
}

func (*HTMLDocument) Onseeked() (onseeked func(event.Event)) {
	macro.Rewrite("$_.onseeked")
	return onseeked
}

func (*HTMLDocument) SetOnseeked(onseeked func(event.Event)) {
	macro.Rewrite("$_.onseeked = $1", onseeked)
}

func (*HTMLDocument) Onseeking() (onseeking func(event.Event)) {
	macro.Rewrite("$_.onseeking")
	return onseeking
}

func (*HTMLDocument) SetOnseeking(onseeking func(event.Event)) {
	macro.Rewrite("$_.onseeking = $1", onseeking)
}

func (*HTMLDocument) Onselect() (onselect func(event.Event)) {
	macro.Rewrite("$_.onselect")
	return onselect
}

func (*HTMLDocument) SetOnselect(onselect func(event.Event)) {
	macro.Rewrite("$_.onselect = $1", onselect)
}

func (*HTMLDocument) Onselectionchange() (onselectionchange func(event.Event)) {
	macro.Rewrite("$_.onselectionchange")
	return onselectionchange
}

func (*HTMLDocument) SetOnselectionchange(onselectionchange func(event.Event)) {
	macro.Rewrite("$_.onselectionchange = $1", onselectionchange)
}

func (*HTMLDocument) Onselectstart() (onselectstart func(event.Event)) {
	macro.Rewrite("$_.onselectstart")
	return onselectstart
}

func (*HTMLDocument) SetOnselectstart(onselectstart func(event.Event)) {
	macro.Rewrite("$_.onselectstart = $1", onselectstart)
}

func (*HTMLDocument) Onstalled() (onstalled func(event.Event)) {
	macro.Rewrite("$_.onstalled")
	return onstalled
}

func (*HTMLDocument) SetOnstalled(onstalled func(event.Event)) {
	macro.Rewrite("$_.onstalled = $1", onstalled)
}

func (*HTMLDocument) Onstop() (onstop func(event.Event)) {
	macro.Rewrite("$_.onstop")
	return onstop
}

func (*HTMLDocument) SetOnstop(onstop func(event.Event)) {
	macro.Rewrite("$_.onstop = $1", onstop)
}

func (*HTMLDocument) Onsubmit() (onsubmit func(event.Event)) {
	macro.Rewrite("$_.onsubmit")
	return onsubmit
}

func (*HTMLDocument) SetOnsubmit(onsubmit func(event.Event)) {
	macro.Rewrite("$_.onsubmit = $1", onsubmit)
}

func (*HTMLDocument) Onsuspend() (onsuspend func(event.Event)) {
	macro.Rewrite("$_.onsuspend")
	return onsuspend
}

func (*HTMLDocument) SetOnsuspend(onsuspend func(event.Event)) {
	macro.Rewrite("$_.onsuspend = $1", onsuspend)
}

func (*HTMLDocument) Ontimeupdate() (ontimeupdate func(event.Event)) {
	macro.Rewrite("$_.ontimeupdate")
	return ontimeupdate
}

func (*HTMLDocument) SetOntimeupdate(ontimeupdate func(event.Event)) {
	macro.Rewrite("$_.ontimeupdate = $1", ontimeupdate)
}

func (*HTMLDocument) Ontouchcancel() (ontouchcancel func(event.Event)) {
	macro.Rewrite("$_.ontouchcancel")
	return ontouchcancel
}

func (*HTMLDocument) SetOntouchcancel(ontouchcancel func(event.Event)) {
	macro.Rewrite("$_.ontouchcancel = $1", ontouchcancel)
}

func (*HTMLDocument) Ontouchend() (ontouchend func(event.Event)) {
	macro.Rewrite("$_.ontouchend")
	return ontouchend
}

func (*HTMLDocument) SetOntouchend(ontouchend func(event.Event)) {
	macro.Rewrite("$_.ontouchend = $1", ontouchend)
}

func (*HTMLDocument) Ontouchmove() (ontouchmove func(event.Event)) {
	macro.Rewrite("$_.ontouchmove")
	return ontouchmove
}

func (*HTMLDocument) SetOntouchmove(ontouchmove func(event.Event)) {
	macro.Rewrite("$_.ontouchmove = $1", ontouchmove)
}

func (*HTMLDocument) Ontouchstart() (ontouchstart func(event.Event)) {
	macro.Rewrite("$_.ontouchstart")
	return ontouchstart
}

func (*HTMLDocument) SetOntouchstart(ontouchstart func(event.Event)) {
	macro.Rewrite("$_.ontouchstart = $1", ontouchstart)
}

func (*HTMLDocument) Onvolumechange() (onvolumechange func(event.Event)) {
	macro.Rewrite("$_.onvolumechange")
	return onvolumechange
}

func (*HTMLDocument) SetOnvolumechange(onvolumechange func(event.Event)) {
	macro.Rewrite("$_.onvolumechange = $1", onvolumechange)
}

func (*HTMLDocument) Onwaiting() (onwaiting func(event.Event)) {
	macro.Rewrite("$_.onwaiting")
	return onwaiting
}

func (*HTMLDocument) SetOnwaiting(onwaiting func(event.Event)) {
	macro.Rewrite("$_.onwaiting = $1", onwaiting)
}

func (*HTMLDocument) Onwebkitfullscreenchange() (onwebkitfullscreenchange func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenchange")
	return onwebkitfullscreenchange
}

func (*HTMLDocument) SetOnwebkitfullscreenchange(onwebkitfullscreenchange func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenchange = $1", onwebkitfullscreenchange)
}

func (*HTMLDocument) Onwebkitfullscreenerror() (onwebkitfullscreenerror func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenerror")
	return onwebkitfullscreenerror
}

func (*HTMLDocument) SetOnwebkitfullscreenerror(onwebkitfullscreenerror func(event.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenerror = $1", onwebkitfullscreenerror)
}

func (*HTMLDocument) Plugins() (plugins HTMLCollection) {
	macro.Rewrite("$_.plugins")
	return plugins
}

func (*HTMLDocument) PointerLockElement() (pointerLockElement element.Element) {
	macro.Rewrite("$_.pointerLockElement")
	return pointerLockElement
}

func (*HTMLDocument) ReadyState() (readyState string) {
	macro.Rewrite("$_.readyState")
	return readyState
}

func (*HTMLDocument) Referrer() (referrer string) {
	macro.Rewrite("$_.referrer")
	return referrer
}

func (*HTMLDocument) RootElement() (rootElement *dom.SVGSVGElement) {
	macro.Rewrite("$_.rootElement")
	return rootElement
}

func (*HTMLDocument) Scripts() (scripts HTMLCollection) {
	macro.Rewrite("$_.scripts")
	return scripts
}

func (*HTMLDocument) ScrollingElement() (scrollingElement element.Element) {
	macro.Rewrite("$_.scrollingElement")
	return scrollingElement
}

func (*HTMLDocument) StyleSheets() (styleSheets *css.StyleSheetList) {
	macro.Rewrite("$_.styleSheets")
	return styleSheets
}

func (*HTMLDocument) Title() (title string) {
	macro.Rewrite("$_.title")
	return title
}

func (*HTMLDocument) SetTitle(title string) {
	macro.Rewrite("$_.title = $1", title)
}

func (*HTMLDocument) URL() (URL string) {
	macro.Rewrite("$_.URL")
	return URL
}

func (*HTMLDocument) URLUnencoded() (URLUnencoded string) {
	macro.Rewrite("$_.URLUnencoded")
	return URLUnencoded
}

func (*HTMLDocument) VisibilityState() (visibilityState *document.VisibilityState) {
	macro.Rewrite("$_.visibilityState")
	return visibilityState
}

func (*HTMLDocument) VlinkColor() (vlinkColor string) {
	macro.Rewrite("$_.vlinkColor")
	return vlinkColor
}

func (*HTMLDocument) SetVlinkColor(vlinkColor string) {
	macro.Rewrite("$_.vlinkColor = $1", vlinkColor)
}

func (*HTMLDocument) WebkitCurrentFullScreenElement() (webkitCurrentFullScreenElement element.Element) {
	macro.Rewrite("$_.webkitCurrentFullScreenElement")
	return webkitCurrentFullScreenElement
}

func (*HTMLDocument) WebkitFullscreenElement() (webkitFullscreenElement element.Element) {
	macro.Rewrite("$_.webkitFullscreenElement")
	return webkitFullscreenElement
}

func (*HTMLDocument) WebkitFullscreenEnabled() (webkitFullscreenEnabled bool) {
	macro.Rewrite("$_.webkitFullscreenEnabled")
	return webkitFullscreenEnabled
}

func (*HTMLDocument) WebkitIsFullScreen() (webkitIsFullScreen bool) {
	macro.Rewrite("$_.webkitIsFullScreen")
	return webkitIsFullScreen
}

func (*HTMLDocument) XMLEncoding() (xmlEncoding string) {
	macro.Rewrite("$_.xmlEncoding")
	return xmlEncoding
}

func (*HTMLDocument) XMLStandalone() (xmlStandalone bool) {
	macro.Rewrite("$_.xmlStandalone")
	return xmlStandalone
}

func (*HTMLDocument) SetXMLStandalone(xmlStandalone bool) {
	macro.Rewrite("$_.xmlStandalone = $1", xmlStandalone)
}

func (*HTMLDocument) XMLVersion() (xmlVersion string) {
	macro.Rewrite("$_.xmlVersion")
	return xmlVersion
}

func (*HTMLDocument) SetXMLVersion(xmlVersion string) {
	macro.Rewrite("$_.xmlVersion = $1", xmlVersion)
}

func (*HTMLDocument) Onpointercancel() (onpointercancel func(event.Event)) {
	macro.Rewrite("$_.onpointercancel")
	return onpointercancel
}

func (*HTMLDocument) SetOnpointercancel(onpointercancel func(event.Event)) {
	macro.Rewrite("$_.onpointercancel = $1", onpointercancel)
}

func (*HTMLDocument) Onpointerdown() (onpointerdown func(event.Event)) {
	macro.Rewrite("$_.onpointerdown")
	return onpointerdown
}

func (*HTMLDocument) SetOnpointerdown(onpointerdown func(event.Event)) {
	macro.Rewrite("$_.onpointerdown = $1", onpointerdown)
}

func (*HTMLDocument) Onpointerenter() (onpointerenter func(event.Event)) {
	macro.Rewrite("$_.onpointerenter")
	return onpointerenter
}

func (*HTMLDocument) SetOnpointerenter(onpointerenter func(event.Event)) {
	macro.Rewrite("$_.onpointerenter = $1", onpointerenter)
}

func (*HTMLDocument) Onpointerleave() (onpointerleave func(event.Event)) {
	macro.Rewrite("$_.onpointerleave")
	return onpointerleave
}

func (*HTMLDocument) SetOnpointerleave(onpointerleave func(event.Event)) {
	macro.Rewrite("$_.onpointerleave = $1", onpointerleave)
}

func (*HTMLDocument) Onpointermove() (onpointermove func(event.Event)) {
	macro.Rewrite("$_.onpointermove")
	return onpointermove
}

func (*HTMLDocument) SetOnpointermove(onpointermove func(event.Event)) {
	macro.Rewrite("$_.onpointermove = $1", onpointermove)
}

func (*HTMLDocument) Onpointerout() (onpointerout func(event.Event)) {
	macro.Rewrite("$_.onpointerout")
	return onpointerout
}

func (*HTMLDocument) SetOnpointerout(onpointerout func(event.Event)) {
	macro.Rewrite("$_.onpointerout = $1", onpointerout)
}

func (*HTMLDocument) Onpointerover() (onpointerover func(event.Event)) {
	macro.Rewrite("$_.onpointerover")
	return onpointerover
}

func (*HTMLDocument) SetOnpointerover(onpointerover func(event.Event)) {
	macro.Rewrite("$_.onpointerover = $1", onpointerover)
}

func (*HTMLDocument) Onpointerup() (onpointerup func(event.Event)) {
	macro.Rewrite("$_.onpointerup")
	return onpointerup
}

func (*HTMLDocument) SetOnpointerup(onpointerup func(event.Event)) {
	macro.Rewrite("$_.onpointerup = $1", onpointerup)
}

func (*HTMLDocument) Onwheel() (onwheel func(event.Event)) {
	macro.Rewrite("$_.onwheel")
	return onwheel
}

func (*HTMLDocument) SetOnwheel(onwheel func(event.Event)) {
	macro.Rewrite("$_.onwheel = $1", onwheel)
}

func (*HTMLDocument) Attributes() (attributes *dom.NamedNodeMap) {
	macro.Rewrite("$_.attributes")
	return attributes
}

func (*HTMLDocument) BaseURI() (baseURI string) {
	macro.Rewrite("$_.baseURI")
	return baseURI
}

func (*HTMLDocument) ChildNodes() (childNodes *dom.NodeList) {
	macro.Rewrite("$_.childNodes")
	return childNodes
}

func (*HTMLDocument) FirstChild() (firstChild dom.Node) {
	macro.Rewrite("$_.firstChild")
	return firstChild
}

func (*HTMLDocument) LastChild() (lastChild dom.Node) {
	macro.Rewrite("$_.lastChild")
	return lastChild
}

func (*HTMLDocument) LocalName() (localName string) {
	macro.Rewrite("$_.localName")
	return localName
}

func (*HTMLDocument) NamespaceURI() (namespaceURI string) {
	macro.Rewrite("$_.namespaceURI")
	return namespaceURI
}

func (*HTMLDocument) NextSibling() (nextSibling dom.Node) {
	macro.Rewrite("$_.nextSibling")
	return nextSibling
}

func (*HTMLDocument) NodeName() (nodeName string) {
	macro.Rewrite("$_.nodeName")
	return nodeName
}

func (*HTMLDocument) NodeType() (nodeType uint8) {
	macro.Rewrite("$_.nodeType")
	return nodeType
}

func (*HTMLDocument) NodeValue() (nodeValue string) {
	macro.Rewrite("$_.nodeValue")
	return nodeValue
}

func (*HTMLDocument) SetNodeValue(nodeValue string) {
	macro.Rewrite("$_.nodeValue = $1", nodeValue)
}

func (*HTMLDocument) OwnerDocument() (ownerDocument document.Document) {
	macro.Rewrite("$_.ownerDocument")
	return ownerDocument
}

func (*HTMLDocument) ParentElement() (parentElement HTMLElement) {
	macro.Rewrite("$_.parentElement")
	return parentElement
}

func (*HTMLDocument) ParentNode() (parentNode dom.Node) {
	macro.Rewrite("$_.parentNode")
	return parentNode
}

func (*HTMLDocument) PreviousSibling() (previousSibling dom.Node) {
	macro.Rewrite("$_.previousSibling")
	return previousSibling
}

func (*HTMLDocument) TextContent() (textContent string) {
	macro.Rewrite("$_.textContent")
	return textContent
}

func (*HTMLDocument) SetTextContent(textContent string) {
	macro.Rewrite("$_.textContent = $1", textContent)
}
