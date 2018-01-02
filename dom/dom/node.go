package dom

import (
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/document"
	"github.com/matthewmueller/joy/dom/html"
)

type Node interface {
	event.
		EventTarget

	AppendChild(newChild Node) (n Node)

	CloneNode(deep *bool) (n Node)

	CompareDocumentPosition(other Node) (u uint8)

	Contains(child Node) (b bool)

	HasAttributes() (b bool)

	HasChildNodes() (b bool)

	InsertBefore(newChild Node, refChild *Node) (n Node)

	IsDefaultNamespace(namespaceURI string) (b bool)

	IsEqualNode(arg Node) (b bool)

	IsSameNode(other Node) (b bool)

	LookupNamespaceURI(prefix string) (s string)

	LookupPrefix(namespaceURI string) (s string)

	Normalize()

	RemoveChild(oldChild Node) (n Node)

	ReplaceChild(newChild Node, oldChild Node) (n Node)

	Attributes() (attributes *NamedNodeMap)

	BaseURI() (baseURI string)

	ChildNodes() (childNodes *NodeList)

	FirstChild() (firstChild Node)

	LastChild() (lastChild Node)

	LocalName() (localName string)

	NamespaceURI() (namespaceURI string)

	NextSibling() (nextSibling Node)

	NodeName() (nodeName string)

	NodeType() (nodeType uint8)

	NodeValue() (nodeValue string)

	SetNodeValue(nodeValue string)

	OwnerDocument() (ownerDocument document.Document)

	ParentElement() (parentElement html.HTMLElement)

	ParentNode() (parentNode Node)

	PreviousSibling() (previousSibling Node)

	TextContent() (textContent string)

	SetTextContent(textContent string)
}
