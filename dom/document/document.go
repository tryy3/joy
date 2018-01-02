package document

import (
	"github.com/matthewmueller/joy/dom/dom"
	"github.com/matthewmueller/joy/dom/element"
	"github.com/matthewmueller/joy/dom/css"
	"html"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/ms"
	"github.com/matthewmueller/joy/dom/ui"
	"github.com/matthewmueller/joy/dom/url"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/dom/utils"
	"github.com/matthewmueller/joy/dom/treewalker"
	"github.com/matthewmueller/joy/dom/touch"
)

type Document interface {
	dom.
		Node

	QuerySelector(selectors string) (e element.Element)

	QuerySelectorAll(selectors string) (n *dom.NodeList)

	CreateEvent(eventInterface string) (e event.Event)

	AdoptNode(source dom.Node) (n dom.Node)

	CaptureEvents()

	CaretRangeFromPoint(x float32, y float32) (r *dom.Range)

	Clear()

	Close()

	CreateAttribute(name string) (a *dom.Attr)

	CreateAttributeNS(namespaceURI string, qualifiedName string) (a *dom.Attr)

	CreateCDATASection(data string) (c *dom.CDATASection)

	CreateComment(data string) (c *dom.Comment)

	CreateDocumentFragment() (d *DocumentFragment)

	CreateElement(tagName string) (e element.Element)

	CreateElementNS(namespaceURI string, qualifiedName string) (e element.Element)

	CreateExpression(expression string, resolver *utils.XPathNSResolver) (x *utils.XPathExpression)

	CreateNodeIterator(root dom.Node, whatToShow *uint, filter *dom.NodeFilter, entityReferenceExpansion *bool) (n *dom.NodeIterator)

	CreateNSResolver(nodeResolver dom.Node) (x *utils.XPathNSResolver)

	CreateProcessingInstruction(target string, data string) (p *dom.ProcessingInstruction)

	CreateRange() (r *dom.Range)

	CreateTextNode(data string) (t dom.Text)

	CreateTouch(view *window.Window, target event.EventTarget, identifier int, pageX int, pageY int, screenX int, screenY int) (t *touch.Touch)

	CreateTouchList(touches *touch.Touch) (t *touch.TouchList)

	CreateTreeWalker(root dom.Node, whatToShow *uint, filter *dom.NodeFilter, entityReferenceExpansion *bool) (t *treewalker.TreeWalker)

	ElementFromPoint(x int, y int) (e element.Element)

	Evaluate(expression string, contextNode dom.Node, resolver *utils.XPathNSResolver, kind uint8, result *utils.XPathResult) (x *utils.XPathResult)

	ExecCommand(commandId string, showUI *bool, value *interface{}) (b bool)

	ExecCommandShowHelp(commandId string) (b bool)

	ExitFullscreen()

	ExitPointerLock()

	Focus()

	GetElementByID(elementId string) (e element.Element)

	GetElementsByClassName(classNames string) (n *dom.NodeList)

	GetElementsByName(elementName string) (n *dom.NodeList)

	GetElementsByTagName(tagname string) (n *dom.NodeList)

	GetElementsByTagNameNS(namespaceURI string, localName string) (n *dom.NodeList)

	GetSelection() (s *dom.Selection)

	HasFocus() (b bool)

	ImportNode(importedNode dom.Node, deep bool) (n dom.Node)

	MsElementsFromPoint(x float32, y float32) (n *dom.NodeList)

	MsElementsFromRect(left float32, top float32, width float32, height float32) (n *dom.NodeList)

	Open(url *string, name *string, features *string, replace *bool) (i interface{})

	QueryCommandEnabled(commandId string) (b bool)

	QueryCommandIndeterm(commandId string) (b bool)

	QueryCommandState(commandId string) (b bool)

	QueryCommandSupported(commandId string) (b bool)

	QueryCommandText(commandId string) (s string)

	QueryCommandValue(commandId string) (s string)

	ReleaseEvents()

	UpdateSettings()

	WebkitCancelFullScreen()

	WebkitExitFullscreen()

	Write(content string)

	Writeln(content string)

	Onpointercancel() (onpointercancel func(event.Event))

	SetOnpointercancel(onpointercancel func(event.Event))

	Onpointerdown() (onpointerdown func(event.Event))

	SetOnpointerdown(onpointerdown func(event.Event))

	Onpointerenter() (onpointerenter func(event.Event))

	SetOnpointerenter(onpointerenter func(event.Event))

	Onpointerleave() (onpointerleave func(event.Event))

	SetOnpointerleave(onpointerleave func(event.Event))

	Onpointermove() (onpointermove func(event.Event))

	SetOnpointermove(onpointermove func(event.Event))

	Onpointerout() (onpointerout func(event.Event))

	SetOnpointerout(onpointerout func(event.Event))

	Onpointerover() (onpointerover func(event.Event))

	SetOnpointerover(onpointerover func(event.Event))

	Onpointerup() (onpointerup func(event.Event))

	SetOnpointerup(onpointerup func(event.Event))

	Onwheel() (onwheel func(event.Event))

	SetOnwheel(onwheel func(event.Event))

	ActiveElement() (activeElement element.Element)

	AlinkColor() (alinkColor string)

	SetAlinkColor(alinkColor string)

	All() (all *html.HTMLAllCollection)

	Anchors() (anchors html.HTMLCollection)

	Applets() (applets html.HTMLCollection)

	BgColor() (bgColor string)

	SetBgColor(bgColor string)

	Body() (body html.HTMLElement)

	SetBody(body html.HTMLElement)

	CharacterSet() (characterSet string)

	Charset() (charset string)

	SetCharset(charset string)

	CompatMode() (compatMode string)

	Cookie() (cookie string)

	SetCookie(cookie string)

	CurrentScript() (currentScript interface{})

	DefaultView() (defaultView *window.Window)

	DesignMode() (designMode string)

	SetDesignMode(designMode string)

	Dir() (dir string)

	SetDir(dir string)

	Doctype() (doctype *DocumentType)

	DocumentElement() (documentElement element.Element)

	Domain() (domain string)

	SetDomain(domain string)

	Embeds() (embeds html.HTMLCollection)

	FgColor() (fgColor string)

	SetFgColor(fgColor string)

	Forms() (forms html.HTMLCollection)

	FullscreenElement() (fullscreenElement element.Element)

	FullscreenEnabled() (fullscreenEnabled bool)

	Head() (head *html.HTMLHeadElement)

	Hidden() (hidden bool)

	Images() (images html.HTMLCollection)

	Implementation() (implementation *dom.DOMImplementation)

	InputEncoding() (inputEncoding string)

	LastModified() (lastModified string)

	LinkColor() (linkColor string)

	SetLinkColor(linkColor string)

	Links() (links html.HTMLCollection)

	Location() (location *url.Location)

	MsCapsLockWarningOff() (msCapsLockWarningOff bool)

	SetMsCapsLockWarningOff(msCapsLockWarningOff bool)

	MsCSSOMElementFloatMetrics() (msCSSOMElementFloatMetrics bool)

	SetMsCSSOMElementFloatMetrics(msCSSOMElementFloatMetrics bool)

	Onabort() (onabort func(event.Event))

	SetOnabort(onabort func(event.Event))

	Onactivate() (onactivate func(event.Event))

	SetOnactivate(onactivate func(event.Event))

	Onbeforeactivate() (onbeforeactivate func(event.Event))

	SetOnbeforeactivate(onbeforeactivate func(event.Event))

	Onbeforedeactivate() (onbeforedeactivate func(event.Event))

	SetOnbeforedeactivate(onbeforedeactivate func(event.Event))

	Onblur() (onblur func(event.Event))

	SetOnblur(onblur func(event.Event))

	Oncanplay() (oncanplay func(event.Event))

	SetOncanplay(oncanplay func(event.Event))

	Oncanplaythrough() (oncanplaythrough func(event.Event))

	SetOncanplaythrough(oncanplaythrough func(event.Event))

	Onchange() (onchange func(event.Event))

	SetOnchange(onchange func(event.Event))

	Onclick() (onclick func(event.Event))

	SetOnclick(onclick func(event.Event))

	Oncontextmenu() (oncontextmenu func(event.Event))

	SetOncontextmenu(oncontextmenu func(event.Event))

	Ondblclick() (ondblclick func(event.Event))

	SetOndblclick(ondblclick func(event.Event))

	Ondeactivate() (ondeactivate func(event.Event))

	SetOndeactivate(ondeactivate func(event.Event))

	Ondrag() (ondrag func(event.Event))

	SetOndrag(ondrag func(event.Event))

	Ondragend() (ondragend func(event.Event))

	SetOndragend(ondragend func(event.Event))

	Ondragenter() (ondragenter func(event.Event))

	SetOndragenter(ondragenter func(event.Event))

	Ondragleave() (ondragleave func(event.Event))

	SetOndragleave(ondragleave func(event.Event))

	Ondragover() (ondragover func(event.Event))

	SetOndragover(ondragover func(event.Event))

	Ondragstart() (ondragstart func(event.Event))

	SetOndragstart(ondragstart func(event.Event))

	Ondrop() (ondrop func(event.Event))

	SetOndrop(ondrop func(event.Event))

	Ondurationchange() (ondurationchange func(event.Event))

	SetOndurationchange(ondurationchange func(event.Event))

	Onemptied() (onemptied func(event.Event))

	SetOnemptied(onemptied func(event.Event))

	Onended() (onended func(event.Event))

	SetOnended(onended func(event.Event))

	Onerror() (onerror func(event.Event))

	SetOnerror(onerror func(event.Event))

	Onfocus() (onfocus func(event.Event))

	SetOnfocus(onfocus func(event.Event))

	Onfullscreenchange() (onfullscreenchange func(event.Event))

	SetOnfullscreenchange(onfullscreenchange func(event.Event))

	Onfullscreenerror() (onfullscreenerror func(event.Event))

	SetOnfullscreenerror(onfullscreenerror func(event.Event))

	Oninput() (oninput func(event.Event))

	SetOninput(oninput func(event.Event))

	Oninvalid() (oninvalid func(event.Event))

	SetOninvalid(oninvalid func(event.Event))

	Onkeydown() (onkeydown func(event.Event))

	SetOnkeydown(onkeydown func(event.Event))

	Onkeypress() (onkeypress func(event.Event))

	SetOnkeypress(onkeypress func(event.Event))

	Onkeyup() (onkeyup func(event.Event))

	SetOnkeyup(onkeyup func(event.Event))

	Onload() (onload func(event.Event))

	SetOnload(onload func(event.Event))

	Onloadeddata() (onloadeddata func(event.Event))

	SetOnloadeddata(onloadeddata func(event.Event))

	Onloadedmetadata() (onloadedmetadata func(event.Event))

	SetOnloadedmetadata(onloadedmetadata func(event.Event))

	Onloadstart() (onloadstart func(event.Event))

	SetOnloadstart(onloadstart func(event.Event))

	Onmousedown() (onmousedown func(event.Event))

	SetOnmousedown(onmousedown func(event.Event))

	Onmousemove() (onmousemove func(event.Event))

	SetOnmousemove(onmousemove func(event.Event))

	Onmouseout() (onmouseout func(event.Event))

	SetOnmouseout(onmouseout func(event.Event))

	Onmouseover() (onmouseover func(event.Event))

	SetOnmouseover(onmouseover func(event.Event))

	Onmouseup() (onmouseup func(event.Event))

	SetOnmouseup(onmouseup func(event.Event))

	Onmousewheel() (onmousewheel func(event.Event))

	SetOnmousewheel(onmousewheel func(event.Event))

	Onmscontentzoom() (onmscontentzoom func(ui.UIEvent))

	SetOnmscontentzoom(onmscontentzoom func(ui.UIEvent))

	Onmsgesturechange() (onmsgesturechange func(event.Event))

	SetOnmsgesturechange(onmsgesturechange func(event.Event))

	Onmsgesturedoubletap() (onmsgesturedoubletap func(event.Event))

	SetOnmsgesturedoubletap(onmsgesturedoubletap func(event.Event))

	Onmsgestureend() (onmsgestureend func(event.Event))

	SetOnmsgestureend(onmsgestureend func(event.Event))

	Onmsgesturehold() (onmsgesturehold func(event.Event))

	SetOnmsgesturehold(onmsgesturehold func(event.Event))

	Onmsgesturestart() (onmsgesturestart func(event.Event))

	SetOnmsgesturestart(onmsgesturestart func(event.Event))

	Onmsgesturetap() (onmsgesturetap func(event.Event))

	SetOnmsgesturetap(onmsgesturetap func(event.Event))

	Onmsinertiastart() (onmsinertiastart func(event.Event))

	SetOnmsinertiastart(onmsinertiastart func(event.Event))

	Onmsmanipulationstatechanged() (onmsmanipulationstatechanged func(*ms.MSManipulationEvent))

	SetOnmsmanipulationstatechanged(onmsmanipulationstatechanged func(*ms.MSManipulationEvent))

	Onmspointercancel() (onmspointercancel func(event.Event))

	SetOnmspointercancel(onmspointercancel func(event.Event))

	Onmspointerdown() (onmspointerdown func(event.Event))

	SetOnmspointerdown(onmspointerdown func(event.Event))

	Onmspointerenter() (onmspointerenter func(event.Event))

	SetOnmspointerenter(onmspointerenter func(event.Event))

	Onmspointerleave() (onmspointerleave func(event.Event))

	SetOnmspointerleave(onmspointerleave func(event.Event))

	Onmspointermove() (onmspointermove func(event.Event))

	SetOnmspointermove(onmspointermove func(event.Event))

	Onmspointerout() (onmspointerout func(event.Event))

	SetOnmspointerout(onmspointerout func(event.Event))

	Onmspointerover() (onmspointerover func(event.Event))

	SetOnmspointerover(onmspointerover func(event.Event))

	Onmspointerup() (onmspointerup func(event.Event))

	SetOnmspointerup(onmspointerup func(event.Event))

	Onmssitemodejumplistitemremoved() (onmssitemodejumplistitemremoved func(*ms.MSSiteModeEvent))

	SetOnmssitemodejumplistitemremoved(onmssitemodejumplistitemremoved func(*ms.MSSiteModeEvent))

	Onmsthumbnailclick() (onmsthumbnailclick func(*ms.MSSiteModeEvent))

	SetOnmsthumbnailclick(onmsthumbnailclick func(*ms.MSSiteModeEvent))

	Onpause() (onpause func(event.Event))

	SetOnpause(onpause func(event.Event))

	Onplay() (onplay func(event.Event))

	SetOnplay(onplay func(event.Event))

	Onplaying() (onplaying func(event.Event))

	SetOnplaying(onplaying func(event.Event))

	Onpointerlockchange() (onpointerlockchange func(event.Event))

	SetOnpointerlockchange(onpointerlockchange func(event.Event))

	Onpointerlockerror() (onpointerlockerror func(event.Event))

	SetOnpointerlockerror(onpointerlockerror func(event.Event))

	Onprogress() (onprogress func(event.Event))

	SetOnprogress(onprogress func(event.Event))

	Onratechange() (onratechange func(event.Event))

	SetOnratechange(onratechange func(event.Event))

	Onreadystatechange() (onreadystatechange func(event.Event))

	SetOnreadystatechange(onreadystatechange func(event.Event))

	Onreset() (onreset func(event.Event))

	SetOnreset(onreset func(event.Event))

	Onscroll() (onscroll func(event.Event))

	SetOnscroll(onscroll func(event.Event))

	Onseeked() (onseeked func(event.Event))

	SetOnseeked(onseeked func(event.Event))

	Onseeking() (onseeking func(event.Event))

	SetOnseeking(onseeking func(event.Event))

	Onselect() (onselect func(event.Event))

	SetOnselect(onselect func(event.Event))

	Onselectionchange() (onselectionchange func(event.Event))

	SetOnselectionchange(onselectionchange func(event.Event))

	Onselectstart() (onselectstart func(event.Event))

	SetOnselectstart(onselectstart func(event.Event))

	Onstalled() (onstalled func(event.Event))

	SetOnstalled(onstalled func(event.Event))

	Onstop() (onstop func(event.Event))

	SetOnstop(onstop func(event.Event))

	Onsubmit() (onsubmit func(event.Event))

	SetOnsubmit(onsubmit func(event.Event))

	Onsuspend() (onsuspend func(event.Event))

	SetOnsuspend(onsuspend func(event.Event))

	Ontimeupdate() (ontimeupdate func(event.Event))

	SetOntimeupdate(ontimeupdate func(event.Event))

	Ontouchcancel() (ontouchcancel func(event.Event))

	SetOntouchcancel(ontouchcancel func(event.Event))

	Ontouchend() (ontouchend func(event.Event))

	SetOntouchend(ontouchend func(event.Event))

	Ontouchmove() (ontouchmove func(event.Event))

	SetOntouchmove(ontouchmove func(event.Event))

	Ontouchstart() (ontouchstart func(event.Event))

	SetOntouchstart(ontouchstart func(event.Event))

	Onvolumechange() (onvolumechange func(event.Event))

	SetOnvolumechange(onvolumechange func(event.Event))

	Onwaiting() (onwaiting func(event.Event))

	SetOnwaiting(onwaiting func(event.Event))

	Onwebkitfullscreenchange() (onwebkitfullscreenchange func(event.Event))

	SetOnwebkitfullscreenchange(onwebkitfullscreenchange func(event.Event))

	Onwebkitfullscreenerror() (onwebkitfullscreenerror func(event.Event))

	SetOnwebkitfullscreenerror(onwebkitfullscreenerror func(event.Event))

	Plugins() (plugins html.HTMLCollection)

	PointerLockElement() (pointerLockElement element.Element)

	ReadyState() (readyState string)

	Referrer() (referrer string)

	RootElement() (rootElement *dom.SVGSVGElement)

	Scripts() (scripts html.HTMLCollection)

	ScrollingElement() (scrollingElement element.Element)

	StyleSheets() (styleSheets *css.StyleSheetList)

	Title() (title string)

	SetTitle(title string)

	URL() (URL string)

	URLUnencoded() (URLUnencoded string)

	VisibilityState() (visibilityState *VisibilityState)

	VlinkColor() (vlinkColor string)

	SetVlinkColor(vlinkColor string)

	WebkitCurrentFullScreenElement() (webkitCurrentFullScreenElement element.Element)

	WebkitFullscreenElement() (webkitFullscreenElement element.Element)

	WebkitFullscreenEnabled() (webkitFullscreenEnabled bool)

	WebkitIsFullScreen() (webkitIsFullScreen bool)

	XMLEncoding() (xmlEncoding string)

	XMLStandalone() (xmlStandalone bool)

	SetXMLStandalone(xmlStandalone bool)

	XMLVersion() (xmlVersion string)

	SetXMLVersion(xmlVersion string)
}
