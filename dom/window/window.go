package window

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/navigation"
	"github.com/matthewmueller/joy/dom/element"
	"github.com/matthewmueller/joy/dom/css"
	"github.com/matthewmueller/joy/dom/dom"
	"github.com/matthewmueller/joy/dom/webkit"
	"github.com/matthewmueller/joy/dom/fetch"
	"github.com/matthewmueller/joy/dom/utils"
	"github.com/matthewmueller/joy/dom/cache"
	"github.com/matthewmueller/joy/dom/crypto"
	"github.com/matthewmueller/joy/dom/document"
	"github.com/matthewmueller/joy/dom/url"
	"github.com/matthewmueller/joy/dom/browser"
	"github.com/matthewmueller/joy/dom/ms"
	"github.com/matthewmueller/joy/dom/ui"
	"github.com/matthewmueller/joy/dom/storage"
	"github.com/matthewmueller/joy/dom/performance"
	"github.com/matthewmueller/joy/dom/speech"
	"github.com/matthewmueller/joy/dom/indexdb"
)

type External struct{}

var _ event.EventTarget = (*Window)(nil)

func New() *Window {
	macro.Rewrite("window")
	return &Window{}
}

type Window struct {
}

func (*Window) Alert(message *string) {
	macro.Rewrite("$_.alert($1)", message)
}

func (*Window) Blur() {
	macro.Rewrite("$_.blur()")
}

func (*Window) CancelAnimationFrame(handle int) {
	macro.Rewrite("$_.cancelAnimationFrame($1)", handle)
}

func (*Window) CaptureEvents() {
	macro.Rewrite("$_.captureEvents()")
}

func (*Window) Close() {
	macro.Rewrite("$_.close()")
}

func (*Window) Confirm(message *string) (b bool) {
	macro.Rewrite("$_.confirm($1)", message)
	return b
}

func (*Window) DepartFocus(navigationReason *navigation.NavigationReason, origin *navigation.FocusNavigationOrigin) {
	macro.Rewrite("$_.departFocus($1, $2)", navigationReason, origin)
}

func (*Window) Focus() {
	macro.Rewrite("$_.focus()")
}

func (*Window) GetComputedStyle(elt element.Element, pseudoElt *string) (c *css.CSSStyleDeclaration) {
	macro.Rewrite("$_.getComputedStyle($1, $2)", elt, pseudoElt)
	return c
}

func (*Window) GetMatchedCSSRules(elt element.Element, pseudoElt *string) (c *css.CSSRuleList) {
	macro.Rewrite("$_.getMatchedCSSRules($1, $2)", elt, pseudoElt)
	return c
}

func (*Window) GetSelection() (s *dom.Selection) {
	macro.Rewrite("$_.getSelection()")
	return s
}

func (*Window) MatchMedia(mediaQuery string) (m *css.MediaQueryList) {
	macro.Rewrite("$_.matchMedia($1)", mediaQuery)
	return m
}

func (*Window) MoveBy(x *int, y *int) {
	macro.Rewrite("$_.moveBy($1, $2)", x, y)
}

func (*Window) MoveTo(x *int, y *int) {
	macro.Rewrite("$_.moveTo($1, $2)", x, y)
}

func (*Window) MsWriteProfilerMark(profilerMarkName string) {
	macro.Rewrite("$_.msWriteProfilerMark($1)", profilerMarkName)
}

func (*Window) Open(url *string, target *string, features *string, replace *bool) (w *Window) {
	macro.Rewrite("$_.open($1, $2, $3, $4)", url, target, features, replace)
	return w
}

func (*Window) PostMessage(message interface{}, targetOrigin string, transfer *[]interface{}) {
	macro.Rewrite("$_.postMessage($1, $2, $3)", message, targetOrigin, transfer)
}

func (*Window) Print() {
	macro.Rewrite("$_.print()")
}

func (*Window) Prompt(message *string, def *string) (s string) {
	macro.Rewrite("$_.prompt($1, $2)", message, def)
	return s
}

func (*Window) ReleaseEvents() {
	macro.Rewrite("$_.releaseEvents()")
}

func (*Window) RequestAnimationFrame(callback func(time int)) (i int) {
	macro.Rewrite("$_.requestAnimationFrame($1)", callback)
	return i
}

func (*Window) ResizeBy(x *int, y *int) {
	macro.Rewrite("$_.resizeBy($1, $2)", x, y)
}

func (*Window) ResizeTo(x *int, y *int) {
	macro.Rewrite("$_.resizeTo($1, $2)", x, y)
}

func (*Window) Scroll(x *int, y *int) {
	macro.Rewrite("$_.scroll($1, $2)", x, y)
}

func (*Window) ScrollBy(x *int, y *int) {
	macro.Rewrite("$_.scrollBy($1, $2)", x, y)
}

func (*Window) ScrollTo(x *int, y *int) {
	macro.Rewrite("$_.scrollTo($1, $2)", x, y)
}

func (*Window) Stop() {
	macro.Rewrite("$_.stop()")
}

func (*Window) WebkitCancelAnimationFrame(handle int) {
	macro.Rewrite("$_.webkitCancelAnimationFrame($1)", handle)
}

func (*Window) WebkitConvertPointFromNodeToPage(node dom.Node, pt *webkit.WebKitPoint) (w *webkit.WebKitPoint) {
	macro.Rewrite("$_.webkitConvertPointFromNodeToPage($1, $2)", node, pt)
	return w
}

func (*Window) WebkitConvertPointFromPageToNode(node dom.Node, pt *webkit.WebKitPoint) (w *webkit.WebKitPoint) {
	macro.Rewrite("$_.webkitConvertPointFromPageToNode($1, $2)", node, pt)
	return w
}

func (*Window) WebkitRequestAnimationFrame(callback func(time int)) (i int) {
	macro.Rewrite("$_.webkitRequestAnimationFrame($1)", callback)
	return i
}

func (*Window) ClearInterval(handle int) {
	macro.Rewrite("$_.clearInterval($1)", handle)
}

func (*Window) ClearTimeout(handle int) {
	macro.Rewrite("$_.clearTimeout($1)", handle)
}

func (*Window) SetInterval(handler interface{}, timeout *interface{}, args interface{}) (i int) {
	macro.Rewrite("$_.setInterval($1, $2, $3)", handler, timeout, args)
	return i
}

func (*Window) SetTimeout(handler interface{}, timeout *interface{}, args interface{}) (i int) {
	macro.Rewrite("$_.setTimeout($1, $2, $3)", handler, timeout, args)
	return i
}

func (*Window) Atob(encodedString string) (s string) {
	macro.Rewrite("$_.atob($1)", encodedString)
	return s
}

func (*Window) Btoa(rawString string) (s string) {
	macro.Rewrite("$_.btoa($1)", rawString)
	return s
}

func (*Window) Fetch(input *fetch.Request, init *fetch.RequestInit) (r *fetch.Response) {
	macro.Rewrite("await $_.fetch($1, $2)", input, init)
	return r
}

func (*Window) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*Window) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*Window) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*Window) ApplicationCache() (applicationCache *utils.ApplicationCache) {
	macro.Rewrite("$_.applicationCache")
	return applicationCache
}

func (*Window) Caches() (caches *cache.CacheStorage) {
	macro.Rewrite("$_.caches")
	return caches
}

func (*Window) ClientInformation() (clientInformation *Navigator) {
	macro.Rewrite("$_.clientInformation")
	return clientInformation
}

func (*Window) Closed() (closed bool) {
	macro.Rewrite("$_.closed")
	return closed
}

func (*Window) Crypto() (crypto *crypto.Crypto) {
	macro.Rewrite("$_.crypto")
	return crypto
}

func (*Window) DefaultStatus() (defaultStatus string) {
	macro.Rewrite("$_.defaultStatus")
	return defaultStatus
}

func (*Window) SetDefaultStatus(defaultStatus string) {
	macro.Rewrite("$_.defaultStatus = $1", defaultStatus)
}

func (*Window) DevicePixelRatio() (devicePixelRatio float32) {
	macro.Rewrite("$_.devicePixelRatio")
	return devicePixelRatio
}

func (*Window) Document() (document document.Document) {
	macro.Rewrite("$_.document")
	return document
}

func (*Window) DoNotTrack() (doNotTrack string) {
	macro.Rewrite("$_.doNotTrack")
	return doNotTrack
}

func (*Window) Event() (event event.Event) {
	macro.Rewrite("$_.event")
	return event
}

func (*Window) SetEvent(event event.Event) {
	macro.Rewrite("$_.event = $1", event)
}

func (*Window) External() (external *External) {
	macro.Rewrite("$_.external")
	return external
}

func (*Window) FrameElement() (frameElement element.Element) {
	macro.Rewrite("$_.frameElement")
	return frameElement
}

func (*Window) Frames() (frames *Window) {
	macro.Rewrite("$_.frames")
	return frames
}

func (*Window) History() (history *History) {
	macro.Rewrite("$_.history")
	return history
}

func (*Window) InnerHeight() (innerHeight int) {
	macro.Rewrite("$_.innerHeight")
	return innerHeight
}

func (*Window) InnerWidth() (innerWidth int) {
	macro.Rewrite("$_.innerWidth")
	return innerWidth
}

func (*Window) IsSecureContext() (isSecureContext bool) {
	macro.Rewrite("$_.isSecureContext")
	return isSecureContext
}

func (*Window) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}

func (*Window) Location() (location *url.Location) {
	macro.Rewrite("$_.location")
	return location
}

func (*Window) Locationbar() (locationbar *browser.BarProp) {
	macro.Rewrite("$_.locationbar")
	return locationbar
}

func (*Window) Menubar() (menubar *browser.BarProp) {
	macro.Rewrite("$_.menubar")
	return menubar
}

func (*Window) MsContentScript() (msContentScript *ms.ExtensionScriptApis) {
	macro.Rewrite("$_.msContentScript")
	return msContentScript
}

func (*Window) MsCredentials() (msCredentials *ms.MSCredentials) {
	macro.Rewrite("$_.msCredentials")
	return msCredentials
}

func (*Window) Name() (name string) {
	macro.Rewrite("$_.name")
	return name
}

func (*Window) SetName(name string) {
	macro.Rewrite("$_.name = $1", name)
}

func (*Window) Navigator() (navigator *Navigator) {
	macro.Rewrite("$_.navigator")
	return navigator
}

func (*Window) OffscreenBuffering() (offscreenBuffering interface{}) {
	macro.Rewrite("$_.offscreenBuffering")
	return offscreenBuffering
}

func (*Window) SetOffscreenBuffering(offscreenBuffering interface{}) {
	macro.Rewrite("$_.offscreenBuffering = $1", offscreenBuffering)
}

func (*Window) Onabort() (onabort func(event.Event)) {
	macro.Rewrite("$_.onabort")
	return onabort
}

func (*Window) SetOnabort(onabort func(event.Event)) {
	macro.Rewrite("$_.onabort = $1", onabort)
}

func (*Window) Onafterprint() (onafterprint func(event.Event)) {
	macro.Rewrite("$_.onafterprint")
	return onafterprint
}

func (*Window) SetOnafterprint(onafterprint func(event.Event)) {
	macro.Rewrite("$_.onafterprint = $1", onafterprint)
}

func (*Window) Onbeforeprint() (onbeforeprint func(event.Event)) {
	macro.Rewrite("$_.onbeforeprint")
	return onbeforeprint
}

func (*Window) SetOnbeforeprint(onbeforeprint func(event.Event)) {
	macro.Rewrite("$_.onbeforeprint = $1", onbeforeprint)
}

func (*Window) Onbeforeunload() (onbeforeunload func(*utils.BeforeUnloadEvent)) {
	macro.Rewrite("$_.onbeforeunload")
	return onbeforeunload
}

func (*Window) SetOnbeforeunload(onbeforeunload func(*utils.BeforeUnloadEvent)) {
	macro.Rewrite("$_.onbeforeunload = $1", onbeforeunload)
}

func (*Window) Onblur() (onblur func(*utils.FocusEvent)) {
	macro.Rewrite("$_.onblur")
	return onblur
}

func (*Window) SetOnblur(onblur func(*utils.FocusEvent)) {
	macro.Rewrite("$_.onblur = $1", onblur)
}

func (*Window) Oncanplay() (oncanplay func(event.Event)) {
	macro.Rewrite("$_.oncanplay")
	return oncanplay
}

func (*Window) SetOncanplay(oncanplay func(event.Event)) {
	macro.Rewrite("$_.oncanplay = $1", oncanplay)
}

func (*Window) Oncanplaythrough() (oncanplaythrough func(event.Event)) {
	macro.Rewrite("$_.oncanplaythrough")
	return oncanplaythrough
}

func (*Window) SetOncanplaythrough(oncanplaythrough func(event.Event)) {
	macro.Rewrite("$_.oncanplaythrough = $1", oncanplaythrough)
}

func (*Window) Onchange() (onchange func(event.Event)) {
	macro.Rewrite("$_.onchange")
	return onchange
}

func (*Window) SetOnchange(onchange func(event.Event)) {
	macro.Rewrite("$_.onchange = $1", onchange)
}

func (*Window) Onclick() (onclick func(event.Event)) {
	macro.Rewrite("$_.onclick")
	return onclick
}

func (*Window) SetOnclick(onclick func(event.Event)) {
	macro.Rewrite("$_.onclick = $1", onclick)
}

func (*Window) Oncompassneedscalibration() (oncompassneedscalibration func(event.Event)) {
	macro.Rewrite("$_.oncompassneedscalibration")
	return oncompassneedscalibration
}

func (*Window) SetOncompassneedscalibration(oncompassneedscalibration func(event.Event)) {
	macro.Rewrite("$_.oncompassneedscalibration = $1", oncompassneedscalibration)
}

func (*Window) Oncontextmenu() (oncontextmenu func(event.Event)) {
	macro.Rewrite("$_.oncontextmenu")
	return oncontextmenu
}

func (*Window) SetOncontextmenu(oncontextmenu func(event.Event)) {
	macro.Rewrite("$_.oncontextmenu = $1", oncontextmenu)
}

func (*Window) Ondblclick() (ondblclick func(event.Event)) {
	macro.Rewrite("$_.ondblclick")
	return ondblclick
}

func (*Window) SetOndblclick(ondblclick func(event.Event)) {
	macro.Rewrite("$_.ondblclick = $1", ondblclick)
}

func (*Window) Ondevicelight() (ondevicelight func(*utils.DeviceLightEvent)) {
	macro.Rewrite("$_.ondevicelight")
	return ondevicelight
}

func (*Window) SetOndevicelight(ondevicelight func(*utils.DeviceLightEvent)) {
	macro.Rewrite("$_.ondevicelight = $1", ondevicelight)
}

func (*Window) Ondevicemotion() (ondevicemotion func(*utils.DeviceMotionEvent)) {
	macro.Rewrite("$_.ondevicemotion")
	return ondevicemotion
}

func (*Window) SetOndevicemotion(ondevicemotion func(*utils.DeviceMotionEvent)) {
	macro.Rewrite("$_.ondevicemotion = $1", ondevicemotion)
}

func (*Window) Ondeviceorientation() (ondeviceorientation func(*utils.DeviceOrientationEvent)) {
	macro.Rewrite("$_.ondeviceorientation")
	return ondeviceorientation
}

func (*Window) SetOndeviceorientation(ondeviceorientation func(*utils.DeviceOrientationEvent)) {
	macro.Rewrite("$_.ondeviceorientation = $1", ondeviceorientation)
}

func (*Window) Ondrag() (ondrag func(event.Event)) {
	macro.Rewrite("$_.ondrag")
	return ondrag
}

func (*Window) SetOndrag(ondrag func(event.Event)) {
	macro.Rewrite("$_.ondrag = $1", ondrag)
}

func (*Window) Ondragend() (ondragend func(event.Event)) {
	macro.Rewrite("$_.ondragend")
	return ondragend
}

func (*Window) SetOndragend(ondragend func(event.Event)) {
	macro.Rewrite("$_.ondragend = $1", ondragend)
}

func (*Window) Ondragenter() (ondragenter func(event.Event)) {
	macro.Rewrite("$_.ondragenter")
	return ondragenter
}

func (*Window) SetOndragenter(ondragenter func(event.Event)) {
	macro.Rewrite("$_.ondragenter = $1", ondragenter)
}

func (*Window) Ondragleave() (ondragleave func(event.Event)) {
	macro.Rewrite("$_.ondragleave")
	return ondragleave
}

func (*Window) SetOndragleave(ondragleave func(event.Event)) {
	macro.Rewrite("$_.ondragleave = $1", ondragleave)
}

func (*Window) Ondragover() (ondragover func(event.Event)) {
	macro.Rewrite("$_.ondragover")
	return ondragover
}

func (*Window) SetOndragover(ondragover func(event.Event)) {
	macro.Rewrite("$_.ondragover = $1", ondragover)
}

func (*Window) Ondragstart() (ondragstart func(event.Event)) {
	macro.Rewrite("$_.ondragstart")
	return ondragstart
}

func (*Window) SetOndragstart(ondragstart func(event.Event)) {
	macro.Rewrite("$_.ondragstart = $1", ondragstart)
}

func (*Window) Ondrop() (ondrop func(event.Event)) {
	macro.Rewrite("$_.ondrop")
	return ondrop
}

func (*Window) SetOndrop(ondrop func(event.Event)) {
	macro.Rewrite("$_.ondrop = $1", ondrop)
}

func (*Window) Ondurationchange() (ondurationchange func(event.Event)) {
	macro.Rewrite("$_.ondurationchange")
	return ondurationchange
}

func (*Window) SetOndurationchange(ondurationchange func(event.Event)) {
	macro.Rewrite("$_.ondurationchange = $1", ondurationchange)
}

func (*Window) Onemptied() (onemptied func(event.Event)) {
	macro.Rewrite("$_.onemptied")
	return onemptied
}

func (*Window) SetOnemptied(onemptied func(event.Event)) {
	macro.Rewrite("$_.onemptied = $1", onemptied)
}

func (*Window) Onended() (onended func(event.Event)) {
	macro.Rewrite("$_.onended")
	return onended
}

func (*Window) SetOnended(onended func(event.Event)) {
	macro.Rewrite("$_.onended = $1", onended)
}

func (*Window) Onerror() (onerror func(columnNumber *uint, event interface{}, fileno *uint, source *string)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

func (*Window) SetOnerror(onerror func(columnNumber *uint, event interface{}, fileno *uint, source *string)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

func (*Window) Onfocus() (onfocus func(*utils.FocusEvent)) {
	macro.Rewrite("$_.onfocus")
	return onfocus
}

func (*Window) SetOnfocus(onfocus func(*utils.FocusEvent)) {
	macro.Rewrite("$_.onfocus = $1", onfocus)
}

func (*Window) Onhashchange() (onhashchange func(*event.HashChangeEvent)) {
	macro.Rewrite("$_.onhashchange")
	return onhashchange
}

func (*Window) SetOnhashchange(onhashchange func(*event.HashChangeEvent)) {
	macro.Rewrite("$_.onhashchange = $1", onhashchange)
}

func (*Window) Oninput() (oninput func(event.Event)) {
	macro.Rewrite("$_.oninput")
	return oninput
}

func (*Window) SetOninput(oninput func(event.Event)) {
	macro.Rewrite("$_.oninput = $1", oninput)
}

func (*Window) Oninvalid() (oninvalid func(event.Event)) {
	macro.Rewrite("$_.oninvalid")
	return oninvalid
}

func (*Window) SetOninvalid(oninvalid func(event.Event)) {
	macro.Rewrite("$_.oninvalid = $1", oninvalid)
}

func (*Window) Onkeydown() (onkeydown func(event.Event)) {
	macro.Rewrite("$_.onkeydown")
	return onkeydown
}

func (*Window) SetOnkeydown(onkeydown func(event.Event)) {
	macro.Rewrite("$_.onkeydown = $1", onkeydown)
}

func (*Window) Onkeypress() (onkeypress func(event.Event)) {
	macro.Rewrite("$_.onkeypress")
	return onkeypress
}

func (*Window) SetOnkeypress(onkeypress func(event.Event)) {
	macro.Rewrite("$_.onkeypress = $1", onkeypress)
}

func (*Window) Onkeyup() (onkeyup func(event.Event)) {
	macro.Rewrite("$_.onkeyup")
	return onkeyup
}

func (*Window) SetOnkeyup(onkeyup func(event.Event)) {
	macro.Rewrite("$_.onkeyup = $1", onkeyup)
}

func (*Window) Onload() (onload func(event.Event)) {
	macro.Rewrite("$_.onload")
	return onload
}

func (*Window) SetOnload(onload func(event.Event)) {
	macro.Rewrite("$_.onload = $1", onload)
}

func (*Window) Onloadeddata() (onloadeddata func(event.Event)) {
	macro.Rewrite("$_.onloadeddata")
	return onloadeddata
}

func (*Window) SetOnloadeddata(onloadeddata func(event.Event)) {
	macro.Rewrite("$_.onloadeddata = $1", onloadeddata)
}

func (*Window) Onloadedmetadata() (onloadedmetadata func(event.Event)) {
	macro.Rewrite("$_.onloadedmetadata")
	return onloadedmetadata
}

func (*Window) SetOnloadedmetadata(onloadedmetadata func(event.Event)) {
	macro.Rewrite("$_.onloadedmetadata = $1", onloadedmetadata)
}

func (*Window) Onloadstart() (onloadstart func(event.Event)) {
	macro.Rewrite("$_.onloadstart")
	return onloadstart
}

func (*Window) SetOnloadstart(onloadstart func(event.Event)) {
	macro.Rewrite("$_.onloadstart = $1", onloadstart)
}

func (*Window) Onmessage() (onmessage func(*utils.MessageEvent)) {
	macro.Rewrite("$_.onmessage")
	return onmessage
}

func (*Window) SetOnmessage(onmessage func(*utils.MessageEvent)) {
	macro.Rewrite("$_.onmessage = $1", onmessage)
}

func (*Window) Onmousedown() (onmousedown func(event.Event)) {
	macro.Rewrite("$_.onmousedown")
	return onmousedown
}

func (*Window) SetOnmousedown(onmousedown func(event.Event)) {
	macro.Rewrite("$_.onmousedown = $1", onmousedown)
}

func (*Window) Onmouseenter() (onmouseenter func(event.Event)) {
	macro.Rewrite("$_.onmouseenter")
	return onmouseenter
}

func (*Window) SetOnmouseenter(onmouseenter func(event.Event)) {
	macro.Rewrite("$_.onmouseenter = $1", onmouseenter)
}

func (*Window) Onmouseleave() (onmouseleave func(event.Event)) {
	macro.Rewrite("$_.onmouseleave")
	return onmouseleave
}

func (*Window) SetOnmouseleave(onmouseleave func(event.Event)) {
	macro.Rewrite("$_.onmouseleave = $1", onmouseleave)
}

func (*Window) Onmousemove() (onmousemove func(event.Event)) {
	macro.Rewrite("$_.onmousemove")
	return onmousemove
}

func (*Window) SetOnmousemove(onmousemove func(event.Event)) {
	macro.Rewrite("$_.onmousemove = $1", onmousemove)
}

func (*Window) Onmouseout() (onmouseout func(event.Event)) {
	macro.Rewrite("$_.onmouseout")
	return onmouseout
}

func (*Window) SetOnmouseout(onmouseout func(event.Event)) {
	macro.Rewrite("$_.onmouseout = $1", onmouseout)
}

func (*Window) Onmouseover() (onmouseover func(event.Event)) {
	macro.Rewrite("$_.onmouseover")
	return onmouseover
}

func (*Window) SetOnmouseover(onmouseover func(event.Event)) {
	macro.Rewrite("$_.onmouseover = $1", onmouseover)
}

func (*Window) Onmouseup() (onmouseup func(event.Event)) {
	macro.Rewrite("$_.onmouseup")
	return onmouseup
}

func (*Window) SetOnmouseup(onmouseup func(event.Event)) {
	macro.Rewrite("$_.onmouseup = $1", onmouseup)
}

func (*Window) Onmousewheel() (onmousewheel func(event.Event)) {
	macro.Rewrite("$_.onmousewheel")
	return onmousewheel
}

func (*Window) SetOnmousewheel(onmousewheel func(event.Event)) {
	macro.Rewrite("$_.onmousewheel = $1", onmousewheel)
}

func (*Window) Onmsgesturechange() (onmsgesturechange func(event.Event)) {
	macro.Rewrite("$_.onmsgesturechange")
	return onmsgesturechange
}

func (*Window) SetOnmsgesturechange(onmsgesturechange func(event.Event)) {
	macro.Rewrite("$_.onmsgesturechange = $1", onmsgesturechange)
}

func (*Window) Onmsgesturedoubletap() (onmsgesturedoubletap func(event.Event)) {
	macro.Rewrite("$_.onmsgesturedoubletap")
	return onmsgesturedoubletap
}

func (*Window) SetOnmsgesturedoubletap(onmsgesturedoubletap func(event.Event)) {
	macro.Rewrite("$_.onmsgesturedoubletap = $1", onmsgesturedoubletap)
}

func (*Window) Onmsgestureend() (onmsgestureend func(event.Event)) {
	macro.Rewrite("$_.onmsgestureend")
	return onmsgestureend
}

func (*Window) SetOnmsgestureend(onmsgestureend func(event.Event)) {
	macro.Rewrite("$_.onmsgestureend = $1", onmsgestureend)
}

func (*Window) Onmsgesturehold() (onmsgesturehold func(event.Event)) {
	macro.Rewrite("$_.onmsgesturehold")
	return onmsgesturehold
}

func (*Window) SetOnmsgesturehold(onmsgesturehold func(event.Event)) {
	macro.Rewrite("$_.onmsgesturehold = $1", onmsgesturehold)
}

func (*Window) Onmsgesturestart() (onmsgesturestart func(event.Event)) {
	macro.Rewrite("$_.onmsgesturestart")
	return onmsgesturestart
}

func (*Window) SetOnmsgesturestart(onmsgesturestart func(event.Event)) {
	macro.Rewrite("$_.onmsgesturestart = $1", onmsgesturestart)
}

func (*Window) Onmsgesturetap() (onmsgesturetap func(event.Event)) {
	macro.Rewrite("$_.onmsgesturetap")
	return onmsgesturetap
}

func (*Window) SetOnmsgesturetap(onmsgesturetap func(event.Event)) {
	macro.Rewrite("$_.onmsgesturetap = $1", onmsgesturetap)
}

func (*Window) Onmsinertiastart() (onmsinertiastart func(event.Event)) {
	macro.Rewrite("$_.onmsinertiastart")
	return onmsinertiastart
}

func (*Window) SetOnmsinertiastart(onmsinertiastart func(event.Event)) {
	macro.Rewrite("$_.onmsinertiastart = $1", onmsinertiastart)
}

func (*Window) Onmspointercancel() (onmspointercancel func(event.Event)) {
	macro.Rewrite("$_.onmspointercancel")
	return onmspointercancel
}

func (*Window) SetOnmspointercancel(onmspointercancel func(event.Event)) {
	macro.Rewrite("$_.onmspointercancel = $1", onmspointercancel)
}

func (*Window) Onmspointerdown() (onmspointerdown func(event.Event)) {
	macro.Rewrite("$_.onmspointerdown")
	return onmspointerdown
}

func (*Window) SetOnmspointerdown(onmspointerdown func(event.Event)) {
	macro.Rewrite("$_.onmspointerdown = $1", onmspointerdown)
}

func (*Window) Onmspointerenter() (onmspointerenter func(event.Event)) {
	macro.Rewrite("$_.onmspointerenter")
	return onmspointerenter
}

func (*Window) SetOnmspointerenter(onmspointerenter func(event.Event)) {
	macro.Rewrite("$_.onmspointerenter = $1", onmspointerenter)
}

func (*Window) Onmspointerleave() (onmspointerleave func(event.Event)) {
	macro.Rewrite("$_.onmspointerleave")
	return onmspointerleave
}

func (*Window) SetOnmspointerleave(onmspointerleave func(event.Event)) {
	macro.Rewrite("$_.onmspointerleave = $1", onmspointerleave)
}

func (*Window) Onmspointermove() (onmspointermove func(event.Event)) {
	macro.Rewrite("$_.onmspointermove")
	return onmspointermove
}

func (*Window) SetOnmspointermove(onmspointermove func(event.Event)) {
	macro.Rewrite("$_.onmspointermove = $1", onmspointermove)
}

func (*Window) Onmspointerout() (onmspointerout func(event.Event)) {
	macro.Rewrite("$_.onmspointerout")
	return onmspointerout
}

func (*Window) SetOnmspointerout(onmspointerout func(event.Event)) {
	macro.Rewrite("$_.onmspointerout = $1", onmspointerout)
}

func (*Window) Onmspointerover() (onmspointerover func(event.Event)) {
	macro.Rewrite("$_.onmspointerover")
	return onmspointerover
}

func (*Window) SetOnmspointerover(onmspointerover func(event.Event)) {
	macro.Rewrite("$_.onmspointerover = $1", onmspointerover)
}

func (*Window) Onmspointerup() (onmspointerup func(event.Event)) {
	macro.Rewrite("$_.onmspointerup")
	return onmspointerup
}

func (*Window) SetOnmspointerup(onmspointerup func(event.Event)) {
	macro.Rewrite("$_.onmspointerup = $1", onmspointerup)
}

func (*Window) Onoffline() (onoffline func(event.Event)) {
	macro.Rewrite("$_.onoffline")
	return onoffline
}

func (*Window) SetOnoffline(onoffline func(event.Event)) {
	macro.Rewrite("$_.onoffline = $1", onoffline)
}

func (*Window) Ononline() (ononline func(event.Event)) {
	macro.Rewrite("$_.ononline")
	return ononline
}

func (*Window) SetOnonline(ononline func(event.Event)) {
	macro.Rewrite("$_.ononline = $1", ononline)
}

func (*Window) Onorientationchange() (onorientationchange func(event.Event)) {
	macro.Rewrite("$_.onorientationchange")
	return onorientationchange
}

func (*Window) SetOnorientationchange(onorientationchange func(event.Event)) {
	macro.Rewrite("$_.onorientationchange = $1", onorientationchange)
}

func (*Window) Onpagehide() (onpagehide func(*utils.PageTransitionEvent)) {
	macro.Rewrite("$_.onpagehide")
	return onpagehide
}

func (*Window) SetOnpagehide(onpagehide func(*utils.PageTransitionEvent)) {
	macro.Rewrite("$_.onpagehide = $1", onpagehide)
}

func (*Window) Onpageshow() (onpageshow func(*utils.PageTransitionEvent)) {
	macro.Rewrite("$_.onpageshow")
	return onpageshow
}

func (*Window) SetOnpageshow(onpageshow func(*utils.PageTransitionEvent)) {
	macro.Rewrite("$_.onpageshow = $1", onpageshow)
}

func (*Window) Onpause() (onpause func(event.Event)) {
	macro.Rewrite("$_.onpause")
	return onpause
}

func (*Window) SetOnpause(onpause func(event.Event)) {
	macro.Rewrite("$_.onpause = $1", onpause)
}

func (*Window) Onplay() (onplay func(event.Event)) {
	macro.Rewrite("$_.onplay")
	return onplay
}

func (*Window) SetOnplay(onplay func(event.Event)) {
	macro.Rewrite("$_.onplay = $1", onplay)
}

func (*Window) Onplaying() (onplaying func(event.Event)) {
	macro.Rewrite("$_.onplaying")
	return onplaying
}

func (*Window) SetOnplaying(onplaying func(event.Event)) {
	macro.Rewrite("$_.onplaying = $1", onplaying)
}

func (*Window) Onpopstate() (onpopstate func(*utils.PopStateEvent)) {
	macro.Rewrite("$_.onpopstate")
	return onpopstate
}

func (*Window) SetOnpopstate(onpopstate func(*utils.PopStateEvent)) {
	macro.Rewrite("$_.onpopstate = $1", onpopstate)
}

func (*Window) Onprogress() (onprogress func(event.Event)) {
	macro.Rewrite("$_.onprogress")
	return onprogress
}

func (*Window) SetOnprogress(onprogress func(event.Event)) {
	macro.Rewrite("$_.onprogress = $1", onprogress)
}

func (*Window) Onratechange() (onratechange func(event.Event)) {
	macro.Rewrite("$_.onratechange")
	return onratechange
}

func (*Window) SetOnratechange(onratechange func(event.Event)) {
	macro.Rewrite("$_.onratechange = $1", onratechange)
}

func (*Window) Onreadystatechange() (onreadystatechange func(event.Event)) {
	macro.Rewrite("$_.onreadystatechange")
	return onreadystatechange
}

func (*Window) SetOnreadystatechange(onreadystatechange func(event.Event)) {
	macro.Rewrite("$_.onreadystatechange = $1", onreadystatechange)
}

func (*Window) Onreset() (onreset func(event.Event)) {
	macro.Rewrite("$_.onreset")
	return onreset
}

func (*Window) SetOnreset(onreset func(event.Event)) {
	macro.Rewrite("$_.onreset = $1", onreset)
}

func (*Window) Onresize() (onresize func(ui.UIEvent)) {
	macro.Rewrite("$_.onresize")
	return onresize
}

func (*Window) SetOnresize(onresize func(ui.UIEvent)) {
	macro.Rewrite("$_.onresize = $1", onresize)
}

func (*Window) Onscroll() (onscroll func(event.Event)) {
	macro.Rewrite("$_.onscroll")
	return onscroll
}

func (*Window) SetOnscroll(onscroll func(event.Event)) {
	macro.Rewrite("$_.onscroll = $1", onscroll)
}

func (*Window) Onseeked() (onseeked func(event.Event)) {
	macro.Rewrite("$_.onseeked")
	return onseeked
}

func (*Window) SetOnseeked(onseeked func(event.Event)) {
	macro.Rewrite("$_.onseeked = $1", onseeked)
}

func (*Window) Onseeking() (onseeking func(event.Event)) {
	macro.Rewrite("$_.onseeking")
	return onseeking
}

func (*Window) SetOnseeking(onseeking func(event.Event)) {
	macro.Rewrite("$_.onseeking = $1", onseeking)
}

func (*Window) Onselect() (onselect func(event.Event)) {
	macro.Rewrite("$_.onselect")
	return onselect
}

func (*Window) SetOnselect(onselect func(event.Event)) {
	macro.Rewrite("$_.onselect = $1", onselect)
}

func (*Window) Onstalled() (onstalled func(event.Event)) {
	macro.Rewrite("$_.onstalled")
	return onstalled
}

func (*Window) SetOnstalled(onstalled func(event.Event)) {
	macro.Rewrite("$_.onstalled = $1", onstalled)
}

func (*Window) Onstorage() (onstorage func(*storage.StorageEvent)) {
	macro.Rewrite("$_.onstorage")
	return onstorage
}

func (*Window) SetOnstorage(onstorage func(*storage.StorageEvent)) {
	macro.Rewrite("$_.onstorage = $1", onstorage)
}

func (*Window) Onsubmit() (onsubmit func(event.Event)) {
	macro.Rewrite("$_.onsubmit")
	return onsubmit
}

func (*Window) SetOnsubmit(onsubmit func(event.Event)) {
	macro.Rewrite("$_.onsubmit = $1", onsubmit)
}

func (*Window) Onsuspend() (onsuspend func(event.Event)) {
	macro.Rewrite("$_.onsuspend")
	return onsuspend
}

func (*Window) SetOnsuspend(onsuspend func(event.Event)) {
	macro.Rewrite("$_.onsuspend = $1", onsuspend)
}

func (*Window) Ontimeupdate() (ontimeupdate func(event.Event)) {
	macro.Rewrite("$_.ontimeupdate")
	return ontimeupdate
}

func (*Window) SetOntimeupdate(ontimeupdate func(event.Event)) {
	macro.Rewrite("$_.ontimeupdate = $1", ontimeupdate)
}

func (*Window) Ontouchcancel() (ontouchcancel interface{}) {
	macro.Rewrite("$_.ontouchcancel")
	return ontouchcancel
}

func (*Window) SetOntouchcancel(ontouchcancel interface{}) {
	macro.Rewrite("$_.ontouchcancel = $1", ontouchcancel)
}

func (*Window) Ontouchend() (ontouchend interface{}) {
	macro.Rewrite("$_.ontouchend")
	return ontouchend
}

func (*Window) SetOntouchend(ontouchend interface{}) {
	macro.Rewrite("$_.ontouchend = $1", ontouchend)
}

func (*Window) Ontouchmove() (ontouchmove interface{}) {
	macro.Rewrite("$_.ontouchmove")
	return ontouchmove
}

func (*Window) SetOntouchmove(ontouchmove interface{}) {
	macro.Rewrite("$_.ontouchmove = $1", ontouchmove)
}

func (*Window) Ontouchstart() (ontouchstart interface{}) {
	macro.Rewrite("$_.ontouchstart")
	return ontouchstart
}

func (*Window) SetOntouchstart(ontouchstart interface{}) {
	macro.Rewrite("$_.ontouchstart = $1", ontouchstart)
}

func (*Window) Onunload() (onunload func(event.Event)) {
	macro.Rewrite("$_.onunload")
	return onunload
}

func (*Window) SetOnunload(onunload func(event.Event)) {
	macro.Rewrite("$_.onunload = $1", onunload)
}

func (*Window) Onvolumechange() (onvolumechange func(event.Event)) {
	macro.Rewrite("$_.onvolumechange")
	return onvolumechange
}

func (*Window) SetOnvolumechange(onvolumechange func(event.Event)) {
	macro.Rewrite("$_.onvolumechange = $1", onvolumechange)
}

func (*Window) Onwaiting() (onwaiting func(event.Event)) {
	macro.Rewrite("$_.onwaiting")
	return onwaiting
}

func (*Window) SetOnwaiting(onwaiting func(event.Event)) {
	macro.Rewrite("$_.onwaiting = $1", onwaiting)
}

func (*Window) Opener() (opener *Window) {
	macro.Rewrite("$_.opener")
	return opener
}

func (*Window) Orientation() (orientation string) {
	macro.Rewrite("$_.orientation")
	return orientation
}

func (*Window) OuterHeight() (outerHeight int) {
	macro.Rewrite("$_.outerHeight")
	return outerHeight
}

func (*Window) OuterWidth() (outerWidth int) {
	macro.Rewrite("$_.outerWidth")
	return outerWidth
}

func (*Window) PageXOffset() (pageXOffset int) {
	macro.Rewrite("$_.pageXOffset")
	return pageXOffset
}

func (*Window) PageYOffset() (pageYOffset int) {
	macro.Rewrite("$_.pageYOffset")
	return pageYOffset
}

func (*Window) Parent() (parent *Window) {
	macro.Rewrite("$_.parent")
	return parent
}

func (*Window) Performance() (performance *performance.Performance) {
	macro.Rewrite("$_.performance")
	return performance
}

func (*Window) Personalbar() (personalbar *browser.BarProp) {
	macro.Rewrite("$_.personalbar")
	return personalbar
}

func (*Window) Screen() (screen *utils.Screen) {
	macro.Rewrite("$_.screen")
	return screen
}

func (*Window) ScreenLeft() (screenLeft int) {
	macro.Rewrite("$_.screenLeft")
	return screenLeft
}

func (*Window) ScreenTop() (screenTop int) {
	macro.Rewrite("$_.screenTop")
	return screenTop
}

func (*Window) ScreenX() (screenX int) {
	macro.Rewrite("$_.screenX")
	return screenX
}

func (*Window) ScreenY() (screenY int) {
	macro.Rewrite("$_.screenY")
	return screenY
}

func (*Window) Scrollbars() (scrollbars *browser.BarProp) {
	macro.Rewrite("$_.scrollbars")
	return scrollbars
}

func (*Window) ScrollX() (scrollX int) {
	macro.Rewrite("$_.scrollX")
	return scrollX
}

func (*Window) ScrollY() (scrollY int) {
	macro.Rewrite("$_.scrollY")
	return scrollY
}

func (*Window) Self() (self *Window) {
	macro.Rewrite("$_.self")
	return self
}

func (*Window) SpeechSynthesis() (speechSynthesis *speech.SpeechSynthesis) {
	macro.Rewrite("$_.speechSynthesis")
	return speechSynthesis
}

func (*Window) Status() (status string) {
	macro.Rewrite("$_.status")
	return status
}

func (*Window) SetStatus(status string) {
	macro.Rewrite("$_.status = $1", status)
}

func (*Window) Statusbar() (statusbar *browser.BarProp) {
	macro.Rewrite("$_.statusbar")
	return statusbar
}

func (*Window) StyleMedia() (styleMedia *StyleMedia) {
	macro.Rewrite("$_.styleMedia")
	return styleMedia
}

func (*Window) Toolbar() (toolbar *browser.BarProp) {
	macro.Rewrite("$_.toolbar")
	return toolbar
}

func (*Window) Top() (top *Window) {
	macro.Rewrite("$_.top")
	return top
}

func (*Window) Window() (window *Window) {
	macro.Rewrite("$_.window")
	return window
}

func (*Window) SessionStorage() (sessionStorage *storage.Storage) {
	macro.Rewrite("$_.sessionStorage")
	return sessionStorage
}

func (*Window) LocalStorage() (localStorage *storage.Storage) {
	macro.Rewrite("$_.localStorage")
	return localStorage
}

func (*Window) Console() (console *utils.Console) {
	macro.Rewrite("$_.console")
	return console
}

func (*Window) Onpointercancel() (onpointercancel func(event.Event)) {
	macro.Rewrite("$_.onpointercancel")
	return onpointercancel
}

func (*Window) SetOnpointercancel(onpointercancel func(event.Event)) {
	macro.Rewrite("$_.onpointercancel = $1", onpointercancel)
}

func (*Window) Onpointerdown() (onpointerdown func(event.Event)) {
	macro.Rewrite("$_.onpointerdown")
	return onpointerdown
}

func (*Window) SetOnpointerdown(onpointerdown func(event.Event)) {
	macro.Rewrite("$_.onpointerdown = $1", onpointerdown)
}

func (*Window) Onpointerenter() (onpointerenter func(event.Event)) {
	macro.Rewrite("$_.onpointerenter")
	return onpointerenter
}

func (*Window) SetOnpointerenter(onpointerenter func(event.Event)) {
	macro.Rewrite("$_.onpointerenter = $1", onpointerenter)
}

func (*Window) Onpointerleave() (onpointerleave func(event.Event)) {
	macro.Rewrite("$_.onpointerleave")
	return onpointerleave
}

func (*Window) SetOnpointerleave(onpointerleave func(event.Event)) {
	macro.Rewrite("$_.onpointerleave = $1", onpointerleave)
}

func (*Window) Onpointermove() (onpointermove func(event.Event)) {
	macro.Rewrite("$_.onpointermove")
	return onpointermove
}

func (*Window) SetOnpointermove(onpointermove func(event.Event)) {
	macro.Rewrite("$_.onpointermove = $1", onpointermove)
}

func (*Window) Onpointerout() (onpointerout func(event.Event)) {
	macro.Rewrite("$_.onpointerout")
	return onpointerout
}

func (*Window) SetOnpointerout(onpointerout func(event.Event)) {
	macro.Rewrite("$_.onpointerout = $1", onpointerout)
}

func (*Window) Onpointerover() (onpointerover func(event.Event)) {
	macro.Rewrite("$_.onpointerover")
	return onpointerover
}

func (*Window) SetOnpointerover(onpointerover func(event.Event)) {
	macro.Rewrite("$_.onpointerover = $1", onpointerover)
}

func (*Window) Onpointerup() (onpointerup func(event.Event)) {
	macro.Rewrite("$_.onpointerup")
	return onpointerup
}

func (*Window) SetOnpointerup(onpointerup func(event.Event)) {
	macro.Rewrite("$_.onpointerup = $1", onpointerup)
}

func (*Window) Onwheel() (onwheel func(event.Event)) {
	macro.Rewrite("$_.onwheel")
	return onwheel
}

func (*Window) SetOnwheel(onwheel func(event.Event)) {
	macro.Rewrite("$_.onwheel = $1", onwheel)
}

func (*Window) IndexedDB() (indexedDB *indexdb.IDBFactory) {
	macro.Rewrite("$_.indexedDB")
	return indexedDB
}
