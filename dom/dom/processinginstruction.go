package dom

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/document"
	"github.com/matthewmueller/joy/dom/html"
)

var _ CharacterData = (*ProcessingInstruction)(nil)
var _ ChildNode = (*ProcessingInstruction)(nil)
var _ Node = (*ProcessingInstruction)(nil)
var _ event.EventTarget = (*ProcessingInstruction)(nil)

type ProcessingInstruction struct {
}

func (*ProcessingInstruction) AppendData(arg string) {
	macro.Rewrite("$_.appendData($1)", arg)
}

func (*ProcessingInstruction) DeleteData(offset uint, count uint) {
	macro.Rewrite("$_.deleteData($1, $2)", offset, count)
}

func (*ProcessingInstruction) InsertData(offset uint, arg string) {
	macro.Rewrite("$_.insertData($1, $2)", offset, arg)
}

func (*ProcessingInstruction) ReplaceData(offset uint, count uint, arg string) {
	macro.Rewrite("$_.replaceData($1, $2, $3)", offset, count, arg)
}

func (*ProcessingInstruction) SubstringData(offset uint, count uint) (s string) {
	macro.Rewrite("$_.substringData($1, $2)", offset, count)
	return s
}

func (*ProcessingInstruction) AppendChild(newChild Node) (n Node) {
	macro.Rewrite("$_.appendChild($1)", newChild)
	return n
}

func (*ProcessingInstruction) CloneNode(deep *bool) (n Node) {
	macro.Rewrite("$_.cloneNode($1)", deep)
	return n
}

func (*ProcessingInstruction) CompareDocumentPosition(other Node) (u uint8) {
	macro.Rewrite("$_.compareDocumentPosition($1)", other)
	return u
}

func (*ProcessingInstruction) Contains(child Node) (b bool) {
	macro.Rewrite("$_.contains($1)", child)
	return b
}

func (*ProcessingInstruction) HasAttributes() (b bool) {
	macro.Rewrite("$_.hasAttributes()")
	return b
}

func (*ProcessingInstruction) HasChildNodes() (b bool) {
	macro.Rewrite("$_.hasChildNodes()")
	return b
}

func (*ProcessingInstruction) InsertBefore(newChild Node, refChild *Node) (n Node) {
	macro.Rewrite("$_.insertBefore($1, $2)", newChild, refChild)
	return n
}

func (*ProcessingInstruction) IsDefaultNamespace(namespaceURI string) (b bool) {
	macro.Rewrite("$_.isDefaultNamespace($1)", namespaceURI)
	return b
}

func (*ProcessingInstruction) IsEqualNode(arg Node) (b bool) {
	macro.Rewrite("$_.isEqualNode($1)", arg)
	return b
}

func (*ProcessingInstruction) IsSameNode(other Node) (b bool) {
	macro.Rewrite("$_.isSameNode($1)", other)
	return b
}

func (*ProcessingInstruction) LookupNamespaceURI(prefix string) (s string) {
	macro.Rewrite("$_.lookupNamespaceURI($1)", prefix)
	return s
}

func (*ProcessingInstruction) LookupPrefix(namespaceURI string) (s string) {
	macro.Rewrite("$_.lookupPrefix($1)", namespaceURI)
	return s
}

func (*ProcessingInstruction) Normalize() {
	macro.Rewrite("$_.normalize()")
}

func (*ProcessingInstruction) RemoveChild(oldChild Node) (n Node) {
	macro.Rewrite("$_.removeChild($1)", oldChild)
	return n
}

func (*ProcessingInstruction) ReplaceChild(newChild Node, oldChild Node) (n Node) {
	macro.Rewrite("$_.replaceChild($1, $2)", newChild, oldChild)
	return n
}

func (*ProcessingInstruction) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*ProcessingInstruction) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*ProcessingInstruction) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*ProcessingInstruction) Target() (target string) {
	macro.Rewrite("$_.target")
	return target
}

func (*ProcessingInstruction) Data() (data string) {
	macro.Rewrite("$_.data")
	return data
}

func (*ProcessingInstruction) SetData(data string) {
	macro.Rewrite("$_.data = $1", data)
}

func (*ProcessingInstruction) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}

func (*ProcessingInstruction) Attributes() (attributes *NamedNodeMap) {
	macro.Rewrite("$_.attributes")
	return attributes
}

func (*ProcessingInstruction) BaseURI() (baseURI string) {
	macro.Rewrite("$_.baseURI")
	return baseURI
}

func (*ProcessingInstruction) ChildNodes() (childNodes *NodeList) {
	macro.Rewrite("$_.childNodes")
	return childNodes
}

func (*ProcessingInstruction) FirstChild() (firstChild Node) {
	macro.Rewrite("$_.firstChild")
	return firstChild
}

func (*ProcessingInstruction) LastChild() (lastChild Node) {
	macro.Rewrite("$_.lastChild")
	return lastChild
}

func (*ProcessingInstruction) LocalName() (localName string) {
	macro.Rewrite("$_.localName")
	return localName
}

func (*ProcessingInstruction) NamespaceURI() (namespaceURI string) {
	macro.Rewrite("$_.namespaceURI")
	return namespaceURI
}

func (*ProcessingInstruction) NextSibling() (nextSibling Node) {
	macro.Rewrite("$_.nextSibling")
	return nextSibling
}

func (*ProcessingInstruction) NodeName() (nodeName string) {
	macro.Rewrite("$_.nodeName")
	return nodeName
}

func (*ProcessingInstruction) NodeType() (nodeType uint8) {
	macro.Rewrite("$_.nodeType")
	return nodeType
}

func (*ProcessingInstruction) NodeValue() (nodeValue string) {
	macro.Rewrite("$_.nodeValue")
	return nodeValue
}

func (*ProcessingInstruction) SetNodeValue(nodeValue string) {
	macro.Rewrite("$_.nodeValue = $1", nodeValue)
}

func (*ProcessingInstruction) OwnerDocument() (ownerDocument document.Document) {
	macro.Rewrite("$_.ownerDocument")
	return ownerDocument
}

func (*ProcessingInstruction) ParentElement() (parentElement html.HTMLElement) {
	macro.Rewrite("$_.parentElement")
	return parentElement
}

func (*ProcessingInstruction) ParentNode() (parentNode Node) {
	macro.Rewrite("$_.parentNode")
	return parentNode
}

func (*ProcessingInstruction) PreviousSibling() (previousSibling Node) {
	macro.Rewrite("$_.previousSibling")
	return previousSibling
}

func (*ProcessingInstruction) TextContent() (textContent string) {
	macro.Rewrite("$_.textContent")
	return textContent
}

func (*ProcessingInstruction) SetTextContent(textContent string) {
	macro.Rewrite("$_.textContent = $1", textContent)
}
