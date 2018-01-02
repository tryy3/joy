package document

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/dom"
	"html"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/element"
)

var _ dom.Node = (*DocumentFragment)(nil)
var _ event.EventTarget = (*DocumentFragment)(nil)

type DocumentFragment struct {
}

func (*DocumentFragment) QuerySelector(selectors string) (e element.Element) {
	macro.Rewrite("$_.querySelector($1)", selectors)
	return e
}

func (*DocumentFragment) QuerySelectorAll(selectors string) (n *dom.NodeList) {
	macro.Rewrite("$_.querySelectorAll($1)", selectors)
	return n
}

func (*DocumentFragment) AppendChild(newChild dom.Node) (n dom.Node) {
	macro.Rewrite("$_.appendChild($1)", newChild)
	return n
}

func (*DocumentFragment) CloneNode(deep *bool) (n dom.Node) {
	macro.Rewrite("$_.cloneNode($1)", deep)
	return n
}

func (*DocumentFragment) CompareDocumentPosition(other dom.Node) (u uint8) {
	macro.Rewrite("$_.compareDocumentPosition($1)", other)
	return u
}

func (*DocumentFragment) Contains(child dom.Node) (b bool) {
	macro.Rewrite("$_.contains($1)", child)
	return b
}

func (*DocumentFragment) HasAttributes() (b bool) {
	macro.Rewrite("$_.hasAttributes()")
	return b
}

func (*DocumentFragment) HasChildNodes() (b bool) {
	macro.Rewrite("$_.hasChildNodes()")
	return b
}

func (*DocumentFragment) InsertBefore(newChild dom.Node, refChild *dom.Node) (n dom.Node) {
	macro.Rewrite("$_.insertBefore($1, $2)", newChild, refChild)
	return n
}

func (*DocumentFragment) IsDefaultNamespace(namespaceURI string) (b bool) {
	macro.Rewrite("$_.isDefaultNamespace($1)", namespaceURI)
	return b
}

func (*DocumentFragment) IsEqualNode(arg dom.Node) (b bool) {
	macro.Rewrite("$_.isEqualNode($1)", arg)
	return b
}

func (*DocumentFragment) IsSameNode(other dom.Node) (b bool) {
	macro.Rewrite("$_.isSameNode($1)", other)
	return b
}

func (*DocumentFragment) LookupNamespaceURI(prefix string) (s string) {
	macro.Rewrite("$_.lookupNamespaceURI($1)", prefix)
	return s
}

func (*DocumentFragment) LookupPrefix(namespaceURI string) (s string) {
	macro.Rewrite("$_.lookupPrefix($1)", namespaceURI)
	return s
}

func (*DocumentFragment) Normalize() {
	macro.Rewrite("$_.normalize()")
}

func (*DocumentFragment) RemoveChild(oldChild dom.Node) (n dom.Node) {
	macro.Rewrite("$_.removeChild($1)", oldChild)
	return n
}

func (*DocumentFragment) ReplaceChild(newChild dom.Node, oldChild dom.Node) (n dom.Node) {
	macro.Rewrite("$_.replaceChild($1, $2)", newChild, oldChild)
	return n
}

func (*DocumentFragment) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*DocumentFragment) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*DocumentFragment) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*DocumentFragment) Attributes() (attributes *dom.NamedNodeMap) {
	macro.Rewrite("$_.attributes")
	return attributes
}

func (*DocumentFragment) BaseURI() (baseURI string) {
	macro.Rewrite("$_.baseURI")
	return baseURI
}

func (*DocumentFragment) ChildNodes() (childNodes *dom.NodeList) {
	macro.Rewrite("$_.childNodes")
	return childNodes
}

func (*DocumentFragment) FirstChild() (firstChild dom.Node) {
	macro.Rewrite("$_.firstChild")
	return firstChild
}

func (*DocumentFragment) LastChild() (lastChild dom.Node) {
	macro.Rewrite("$_.lastChild")
	return lastChild
}

func (*DocumentFragment) LocalName() (localName string) {
	macro.Rewrite("$_.localName")
	return localName
}

func (*DocumentFragment) NamespaceURI() (namespaceURI string) {
	macro.Rewrite("$_.namespaceURI")
	return namespaceURI
}

func (*DocumentFragment) NextSibling() (nextSibling dom.Node) {
	macro.Rewrite("$_.nextSibling")
	return nextSibling
}

func (*DocumentFragment) NodeName() (nodeName string) {
	macro.Rewrite("$_.nodeName")
	return nodeName
}

func (*DocumentFragment) NodeType() (nodeType uint8) {
	macro.Rewrite("$_.nodeType")
	return nodeType
}

func (*DocumentFragment) NodeValue() (nodeValue string) {
	macro.Rewrite("$_.nodeValue")
	return nodeValue
}

func (*DocumentFragment) SetNodeValue(nodeValue string) {
	macro.Rewrite("$_.nodeValue = $1", nodeValue)
}

func (*DocumentFragment) OwnerDocument() (ownerDocument Document) {
	macro.Rewrite("$_.ownerDocument")
	return ownerDocument
}

func (*DocumentFragment) ParentElement() (parentElement html.HTMLElement) {
	macro.Rewrite("$_.parentElement")
	return parentElement
}

func (*DocumentFragment) ParentNode() (parentNode dom.Node) {
	macro.Rewrite("$_.parentNode")
	return parentNode
}

func (*DocumentFragment) PreviousSibling() (previousSibling dom.Node) {
	macro.Rewrite("$_.previousSibling")
	return previousSibling
}

func (*DocumentFragment) TextContent() (textContent string) {
	macro.Rewrite("$_.textContent")
	return textContent
}

func (*DocumentFragment) SetTextContent(textContent string) {
	macro.Rewrite("$_.textContent = $1", textContent)
}
