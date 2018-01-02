package document

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/dom"
	"html"
	"github.com/matthewmueller/joy/dom/event"
)

var _ dom.Node = (*DocumentType)(nil)
var _ event.EventTarget = (*DocumentType)(nil)

type DocumentType struct {
}

func (*DocumentType) AppendChild(newChild dom.Node) (n dom.Node) {
	macro.Rewrite("$_.appendChild($1)", newChild)
	return n
}

func (*DocumentType) CloneNode(deep *bool) (n dom.Node) {
	macro.Rewrite("$_.cloneNode($1)", deep)
	return n
}

func (*DocumentType) CompareDocumentPosition(other dom.Node) (u uint8) {
	macro.Rewrite("$_.compareDocumentPosition($1)", other)
	return u
}

func (*DocumentType) Contains(child dom.Node) (b bool) {
	macro.Rewrite("$_.contains($1)", child)
	return b
}

func (*DocumentType) HasAttributes() (b bool) {
	macro.Rewrite("$_.hasAttributes()")
	return b
}

func (*DocumentType) HasChildNodes() (b bool) {
	macro.Rewrite("$_.hasChildNodes()")
	return b
}

func (*DocumentType) InsertBefore(newChild dom.Node, refChild *dom.Node) (n dom.Node) {
	macro.Rewrite("$_.insertBefore($1, $2)", newChild, refChild)
	return n
}

func (*DocumentType) IsDefaultNamespace(namespaceURI string) (b bool) {
	macro.Rewrite("$_.isDefaultNamespace($1)", namespaceURI)
	return b
}

func (*DocumentType) IsEqualNode(arg dom.Node) (b bool) {
	macro.Rewrite("$_.isEqualNode($1)", arg)
	return b
}

func (*DocumentType) IsSameNode(other dom.Node) (b bool) {
	macro.Rewrite("$_.isSameNode($1)", other)
	return b
}

func (*DocumentType) LookupNamespaceURI(prefix string) (s string) {
	macro.Rewrite("$_.lookupNamespaceURI($1)", prefix)
	return s
}

func (*DocumentType) LookupPrefix(namespaceURI string) (s string) {
	macro.Rewrite("$_.lookupPrefix($1)", namespaceURI)
	return s
}

func (*DocumentType) Normalize() {
	macro.Rewrite("$_.normalize()")
}

func (*DocumentType) RemoveChild(oldChild dom.Node) (n dom.Node) {
	macro.Rewrite("$_.removeChild($1)", oldChild)
	return n
}

func (*DocumentType) ReplaceChild(newChild dom.Node, oldChild dom.Node) (n dom.Node) {
	macro.Rewrite("$_.replaceChild($1, $2)", newChild, oldChild)
	return n
}

func (*DocumentType) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*DocumentType) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*DocumentType) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*DocumentType) Entities() (entities *dom.NamedNodeMap) {
	macro.Rewrite("$_.entities")
	return entities
}

func (*DocumentType) InternalSubset() (internalSubset string) {
	macro.Rewrite("$_.internalSubset")
	return internalSubset
}

func (*DocumentType) Name() (name string) {
	macro.Rewrite("$_.name")
	return name
}

func (*DocumentType) Notations() (notations *dom.NamedNodeMap) {
	macro.Rewrite("$_.notations")
	return notations
}

func (*DocumentType) PublicID() (publicId string) {
	macro.Rewrite("$_.publicId")
	return publicId
}

func (*DocumentType) SystemID() (systemId string) {
	macro.Rewrite("$_.systemId")
	return systemId
}

func (*DocumentType) Attributes() (attributes *dom.NamedNodeMap) {
	macro.Rewrite("$_.attributes")
	return attributes
}

func (*DocumentType) BaseURI() (baseURI string) {
	macro.Rewrite("$_.baseURI")
	return baseURI
}

func (*DocumentType) ChildNodes() (childNodes *dom.NodeList) {
	macro.Rewrite("$_.childNodes")
	return childNodes
}

func (*DocumentType) FirstChild() (firstChild dom.Node) {
	macro.Rewrite("$_.firstChild")
	return firstChild
}

func (*DocumentType) LastChild() (lastChild dom.Node) {
	macro.Rewrite("$_.lastChild")
	return lastChild
}

func (*DocumentType) LocalName() (localName string) {
	macro.Rewrite("$_.localName")
	return localName
}

func (*DocumentType) NamespaceURI() (namespaceURI string) {
	macro.Rewrite("$_.namespaceURI")
	return namespaceURI
}

func (*DocumentType) NextSibling() (nextSibling dom.Node) {
	macro.Rewrite("$_.nextSibling")
	return nextSibling
}

func (*DocumentType) NodeName() (nodeName string) {
	macro.Rewrite("$_.nodeName")
	return nodeName
}

func (*DocumentType) NodeType() (nodeType uint8) {
	macro.Rewrite("$_.nodeType")
	return nodeType
}

func (*DocumentType) NodeValue() (nodeValue string) {
	macro.Rewrite("$_.nodeValue")
	return nodeValue
}

func (*DocumentType) SetNodeValue(nodeValue string) {
	macro.Rewrite("$_.nodeValue = $1", nodeValue)
}

func (*DocumentType) OwnerDocument() (ownerDocument Document) {
	macro.Rewrite("$_.ownerDocument")
	return ownerDocument
}

func (*DocumentType) ParentElement() (parentElement html.HTMLElement) {
	macro.Rewrite("$_.parentElement")
	return parentElement
}

func (*DocumentType) ParentNode() (parentNode dom.Node) {
	macro.Rewrite("$_.parentNode")
	return parentNode
}

func (*DocumentType) PreviousSibling() (previousSibling dom.Node) {
	macro.Rewrite("$_.previousSibling")
	return previousSibling
}

func (*DocumentType) TextContent() (textContent string) {
	macro.Rewrite("$_.textContent")
	return textContent
}

func (*DocumentType) SetTextContent(textContent string) {
	macro.Rewrite("$_.textContent = $1", textContent)
}
