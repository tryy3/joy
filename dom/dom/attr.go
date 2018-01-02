package dom

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/element"
	"github.com/matthewmueller/joy/dom/document"
	"github.com/matthewmueller/joy/dom/html"
)

var _ Node = (*Attr)(nil)
var _ event.EventTarget = (*Attr)(nil)

type Attr struct {
}

func (*Attr) AppendChild(newChild Node) (n Node) {
	macro.Rewrite("$_.appendChild($1)", newChild)
	return n
}

func (*Attr) CloneNode(deep *bool) (n Node) {
	macro.Rewrite("$_.cloneNode($1)", deep)
	return n
}

func (*Attr) CompareDocumentPosition(other Node) (u uint8) {
	macro.Rewrite("$_.compareDocumentPosition($1)", other)
	return u
}

func (*Attr) Contains(child Node) (b bool) {
	macro.Rewrite("$_.contains($1)", child)
	return b
}

func (*Attr) HasAttributes() (b bool) {
	macro.Rewrite("$_.hasAttributes()")
	return b
}

func (*Attr) HasChildNodes() (b bool) {
	macro.Rewrite("$_.hasChildNodes()")
	return b
}

func (*Attr) InsertBefore(newChild Node, refChild *Node) (n Node) {
	macro.Rewrite("$_.insertBefore($1, $2)", newChild, refChild)
	return n
}

func (*Attr) IsDefaultNamespace(namespaceURI string) (b bool) {
	macro.Rewrite("$_.isDefaultNamespace($1)", namespaceURI)
	return b
}

func (*Attr) IsEqualNode(arg Node) (b bool) {
	macro.Rewrite("$_.isEqualNode($1)", arg)
	return b
}

func (*Attr) IsSameNode(other Node) (b bool) {
	macro.Rewrite("$_.isSameNode($1)", other)
	return b
}

func (*Attr) LookupNamespaceURI(prefix string) (s string) {
	macro.Rewrite("$_.lookupNamespaceURI($1)", prefix)
	return s
}

func (*Attr) LookupPrefix(namespaceURI string) (s string) {
	macro.Rewrite("$_.lookupPrefix($1)", namespaceURI)
	return s
}

func (*Attr) Normalize() {
	macro.Rewrite("$_.normalize()")
}

func (*Attr) RemoveChild(oldChild Node) (n Node) {
	macro.Rewrite("$_.removeChild($1)", oldChild)
	return n
}

func (*Attr) ReplaceChild(newChild Node, oldChild Node) (n Node) {
	macro.Rewrite("$_.replaceChild($1, $2)", newChild, oldChild)
	return n
}

func (*Attr) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*Attr) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*Attr) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*Attr) Name() (name string) {
	macro.Rewrite("$_.name")
	return name
}

func (*Attr) OwnerElement() (ownerElement element.Element) {
	macro.Rewrite("$_.ownerElement")
	return ownerElement
}

func (*Attr) Prefix() (prefix string) {
	macro.Rewrite("$_.prefix")
	return prefix
}

func (*Attr) Specified() (specified bool) {
	macro.Rewrite("$_.specified")
	return specified
}

func (*Attr) Value() (value string) {
	macro.Rewrite("$_.value")
	return value
}

func (*Attr) SetValue(value string) {
	macro.Rewrite("$_.value = $1", value)
}

func (*Attr) Attributes() (attributes *NamedNodeMap) {
	macro.Rewrite("$_.attributes")
	return attributes
}

func (*Attr) BaseURI() (baseURI string) {
	macro.Rewrite("$_.baseURI")
	return baseURI
}

func (*Attr) ChildNodes() (childNodes *NodeList) {
	macro.Rewrite("$_.childNodes")
	return childNodes
}

func (*Attr) FirstChild() (firstChild Node) {
	macro.Rewrite("$_.firstChild")
	return firstChild
}

func (*Attr) LastChild() (lastChild Node) {
	macro.Rewrite("$_.lastChild")
	return lastChild
}

func (*Attr) LocalName() (localName string) {
	macro.Rewrite("$_.localName")
	return localName
}

func (*Attr) NamespaceURI() (namespaceURI string) {
	macro.Rewrite("$_.namespaceURI")
	return namespaceURI
}

func (*Attr) NextSibling() (nextSibling Node) {
	macro.Rewrite("$_.nextSibling")
	return nextSibling
}

func (*Attr) NodeName() (nodeName string) {
	macro.Rewrite("$_.nodeName")
	return nodeName
}

func (*Attr) NodeType() (nodeType uint8) {
	macro.Rewrite("$_.nodeType")
	return nodeType
}

func (*Attr) NodeValue() (nodeValue string) {
	macro.Rewrite("$_.nodeValue")
	return nodeValue
}

func (*Attr) SetNodeValue(nodeValue string) {
	macro.Rewrite("$_.nodeValue = $1", nodeValue)
}

func (*Attr) OwnerDocument() (ownerDocument document.Document) {
	macro.Rewrite("$_.ownerDocument")
	return ownerDocument
}

func (*Attr) ParentElement() (parentElement html.HTMLElement) {
	macro.Rewrite("$_.parentElement")
	return parentElement
}

func (*Attr) ParentNode() (parentNode Node) {
	macro.Rewrite("$_.parentNode")
	return parentNode
}

func (*Attr) PreviousSibling() (previousSibling Node) {
	macro.Rewrite("$_.previousSibling")
	return previousSibling
}

func (*Attr) TextContent() (textContent string) {
	macro.Rewrite("$_.textContent")
	return textContent
}

func (*Attr) SetTextContent(textContent string) {
	macro.Rewrite("$_.textContent = $1", textContent)
}
