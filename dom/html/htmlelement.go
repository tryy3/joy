package html

import (
	"github.com/matthewmueller/joy/dom/element"
	"github.com/matthewmueller/joy/dom/dom"
	"github.com/matthewmueller/joy/dom/ms"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/ui"
	"github.com/matthewmueller/joy/dom/utils"
	"github.com/matthewmueller/joy/dom/mouse"
	"github.com/matthewmueller/joy/dom/draganddrop"
	"github.com/matthewmueller/joy/dom/css"
)

type HTMLElement interface {
	element.
		Element

	Blur()

	Click()

	DragDrop() (b bool)

	Focus()

	GetElementsByClassName(classNames string) (n *dom.NodeList)

	InsertAdjacentElement(position string, insertedElement element.Element) (e element.Element)

	InsertAdjacentHTML(where string, html string)

	InsertAdjacentText(where string, text string)

	MsGetInputContext() (m *ms.MSInputMethodContext)

	ScrollIntoView(top *bool)

	AccessKey() (accessKey string)

	SetAccessKey(accessKey string)

	Children() (children HTMLCollection)

	ContentEditable() (contentEditable string)

	SetContentEditable(contentEditable string)

	Dataset() (dataset *dom.DOMStringMap)

	Dir() (dir string)

	SetDir(dir string)

	Draggable() (draggable bool)

	SetDraggable(draggable bool)

	Hidden() (hidden bool)

	SetHidden(hidden bool)

	HideFocus() (hideFocus bool)

	SetHideFocus(hideFocus bool)

	InnerText() (innerText string)

	SetInnerText(innerText string)

	IsContentEditable() (isContentEditable bool)

	Lang() (lang string)

	SetLang(lang string)

	OffsetHeight() (offsetHeight int)

	OffsetLeft() (offsetLeft int)

	OffsetParent() (offsetParent element.Element)

	OffsetTop() (offsetTop int)

	OffsetWidth() (offsetWidth int)

	Onabort() (onabort func(event.Event))

	SetOnabort(onabort func(event.Event))

	Onactivate() (onactivate func(ui.UIEvent))

	SetOnactivate(onactivate func(ui.UIEvent))

	Onbeforeactivate() (onbeforeactivate func(ui.UIEvent))

	SetOnbeforeactivate(onbeforeactivate func(ui.UIEvent))

	Onbeforecopy() (onbeforecopy func(*utils.ClipboardEvent))

	SetOnbeforecopy(onbeforecopy func(*utils.ClipboardEvent))

	Onbeforecut() (onbeforecut func(*utils.ClipboardEvent))

	SetOnbeforecut(onbeforecut func(*utils.ClipboardEvent))

	Onbeforedeactivate() (onbeforedeactivate func(ui.UIEvent))

	SetOnbeforedeactivate(onbeforedeactivate func(ui.UIEvent))

	Onbeforepaste() (onbeforepaste func(*utils.ClipboardEvent))

	SetOnbeforepaste(onbeforepaste func(*utils.ClipboardEvent))

	Onblur() (onblur func(*utils.FocusEvent))

	SetOnblur(onblur func(*utils.FocusEvent))

	Oncanplay() (oncanplay func(event.Event))

	SetOncanplay(oncanplay func(event.Event))

	Oncanplaythrough() (oncanplaythrough func(event.Event))

	SetOncanplaythrough(oncanplaythrough func(event.Event))

	Onchange() (onchange func(event.Event))

	SetOnchange(onchange func(event.Event))

	Onclick() (onclick func(mouse.MouseEvent))

	SetOnclick(onclick func(mouse.MouseEvent))

	Oncontextmenu() (oncontextmenu func(*utils.PointerEvent))

	SetOncontextmenu(oncontextmenu func(*utils.PointerEvent))

	Oncopy() (oncopy func(*utils.ClipboardEvent))

	SetOncopy(oncopy func(*utils.ClipboardEvent))

	Oncuechange() (oncuechange func(event.Event))

	SetOncuechange(oncuechange func(event.Event))

	Oncut() (oncut func(*utils.ClipboardEvent))

	SetOncut(oncut func(*utils.ClipboardEvent))

	Ondblclick() (ondblclick func(mouse.MouseEvent))

	SetOndblclick(ondblclick func(mouse.MouseEvent))

	Ondeactivate() (ondeactivate func(ui.UIEvent))

	SetOndeactivate(ondeactivate func(ui.UIEvent))

	Ondrag() (ondrag func(*draganddrop.DragEvent))

	SetOndrag(ondrag func(*draganddrop.DragEvent))

	Ondragend() (ondragend func(*draganddrop.DragEvent))

	SetOndragend(ondragend func(*draganddrop.DragEvent))

	Ondragenter() (ondragenter func(*draganddrop.DragEvent))

	SetOndragenter(ondragenter func(*draganddrop.DragEvent))

	Ondragleave() (ondragleave func(*draganddrop.DragEvent))

	SetOndragleave(ondragleave func(*draganddrop.DragEvent))

	Ondragover() (ondragover func(*draganddrop.DragEvent))

	SetOndragover(ondragover func(*draganddrop.DragEvent))

	Ondragstart() (ondragstart func(*draganddrop.DragEvent))

	SetOndragstart(ondragstart func(*draganddrop.DragEvent))

	Ondrop() (ondrop func(*draganddrop.DragEvent))

	SetOndrop(ondrop func(*draganddrop.DragEvent))

	Ondurationchange() (ondurationchange func(event.Event))

	SetOndurationchange(ondurationchange func(event.Event))

	Onemptied() (onemptied func(event.Event))

	SetOnemptied(onemptied func(event.Event))

	Onended() (onended func(event.Event))

	SetOnended(onended func(event.Event))

	Onerror() (onerror func(event.Event))

	SetOnerror(onerror func(event.Event))

	Onfocus() (onfocus func(*utils.FocusEvent))

	SetOnfocus(onfocus func(*utils.FocusEvent))

	Oninput() (oninput func(event.Event))

	SetOninput(oninput func(event.Event))

	Oninvalid() (oninvalid func(event.Event))

	SetOninvalid(oninvalid func(event.Event))

	Onkeydown() (onkeydown func(*utils.KeyboardEvent))

	SetOnkeydown(onkeydown func(*utils.KeyboardEvent))

	Onkeypress() (onkeypress func(*utils.KeyboardEvent))

	SetOnkeypress(onkeypress func(*utils.KeyboardEvent))

	Onkeyup() (onkeyup func(*utils.KeyboardEvent))

	SetOnkeyup(onkeyup func(*utils.KeyboardEvent))

	Onload() (onload func(event.Event))

	SetOnload(onload func(event.Event))

	Onloadeddata() (onloadeddata func(event.Event))

	SetOnloadeddata(onloadeddata func(event.Event))

	Onloadedmetadata() (onloadedmetadata func(event.Event))

	SetOnloadedmetadata(onloadedmetadata func(event.Event))

	Onloadstart() (onloadstart func(event.Event))

	SetOnloadstart(onloadstart func(event.Event))

	Onmousedown() (onmousedown func(mouse.MouseEvent))

	SetOnmousedown(onmousedown func(mouse.MouseEvent))

	Onmouseenter() (onmouseenter func(mouse.MouseEvent))

	SetOnmouseenter(onmouseenter func(mouse.MouseEvent))

	Onmouseleave() (onmouseleave func(mouse.MouseEvent))

	SetOnmouseleave(onmouseleave func(mouse.MouseEvent))

	Onmousemove() (onmousemove func(mouse.MouseEvent))

	SetOnmousemove(onmousemove func(mouse.MouseEvent))

	Onmouseout() (onmouseout func(mouse.MouseEvent))

	SetOnmouseout(onmouseout func(mouse.MouseEvent))

	Onmouseover() (onmouseover func(mouse.MouseEvent))

	SetOnmouseover(onmouseover func(mouse.MouseEvent))

	Onmouseup() (onmouseup func(mouse.MouseEvent))

	SetOnmouseup(onmouseup func(mouse.MouseEvent))

	Onmousewheel() (onmousewheel func(*utils.WheelEvent))

	SetOnmousewheel(onmousewheel func(*utils.WheelEvent))

	Onmscontentzoom() (onmscontentzoom func(ui.UIEvent))

	SetOnmscontentzoom(onmscontentzoom func(ui.UIEvent))

	Onmsmanipulationstatechanged() (onmsmanipulationstatechanged func(*ms.MSManipulationEvent))

	SetOnmsmanipulationstatechanged(onmsmanipulationstatechanged func(*ms.MSManipulationEvent))

	Onpaste() (onpaste func(*utils.ClipboardEvent))

	SetOnpaste(onpaste func(*utils.ClipboardEvent))

	Onpause() (onpause func(event.Event))

	SetOnpause(onpause func(event.Event))

	Onplay() (onplay func(event.Event))

	SetOnplay(onplay func(event.Event))

	Onplaying() (onplaying func(event.Event))

	SetOnplaying(onplaying func(event.Event))

	Onprogress() (onprogress func(event.Event))

	SetOnprogress(onprogress func(event.Event))

	Onratechange() (onratechange func(event.Event))

	SetOnratechange(onratechange func(event.Event))

	Onreset() (onreset func(event.Event))

	SetOnreset(onreset func(event.Event))

	Onscroll() (onscroll func(ui.UIEvent))

	SetOnscroll(onscroll func(ui.UIEvent))

	Onseeked() (onseeked func(event.Event))

	SetOnseeked(onseeked func(event.Event))

	Onseeking() (onseeking func(event.Event))

	SetOnseeking(onseeking func(event.Event))

	Onselect() (onselect func(ui.UIEvent))

	SetOnselect(onselect func(ui.UIEvent))

	Onselectstart() (onselectstart func(event.Event))

	SetOnselectstart(onselectstart func(event.Event))

	Onstalled() (onstalled func(event.Event))

	SetOnstalled(onstalled func(event.Event))

	Onsubmit() (onsubmit func(event.Event))

	SetOnsubmit(onsubmit func(event.Event))

	Onsuspend() (onsuspend func(event.Event))

	SetOnsuspend(onsuspend func(event.Event))

	Ontimeupdate() (ontimeupdate func(event.Event))

	SetOntimeupdate(ontimeupdate func(event.Event))

	Onvolumechange() (onvolumechange func(event.Event))

	SetOnvolumechange(onvolumechange func(event.Event))

	Onwaiting() (onwaiting func(event.Event))

	SetOnwaiting(onwaiting func(event.Event))

	OuterText() (outerText string)

	SetOuterText(outerText string)

	Spellcheck() (spellcheck bool)

	SetSpellcheck(spellcheck bool)

	Style() (style *css.CSSStyleDeclaration)

	TabIndex() (tabIndex int8)

	SetTabIndex(tabIndex int8)

	Title() (title string)

	SetTitle(title string)
}
