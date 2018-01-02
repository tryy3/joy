package dom

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/document"
	"github.com/matthewmueller/joy/dom/html"
)

var _ CharacterData = (*Comment)(nil)
var _ ChildNode = (*Comment)(nil)
var _ Node = (*Comment)(nil)
var _ event.EventTarget = (*Comment)(nil)

type Comment struct {
}

func (*Comment) AppendData(arg string) {
	macro.Rewrite("$_.appendData($1)", arg)
}

func (*Comment) DeleteData(offset uint, count uint) {
	macro.Rewrite("$_.deleteData($1, $2)", offset, count)
}

func (*Comment) InsertData(offset uint, arg string) {
	macro.Rewrite("$_.insertData($1, $2)", offset, arg)
}

func (*Comment) ReplaceData(offset uint, count uint, arg string) {
	macro.Rewrite("$_.replaceData($1, $2, $3)", offset, count, arg)
}

func (*Comment) SubstringData(offset uint, count uint) (s string) {
	macro.Rewrite("$_.substringData($1, $2)", offset, count)
	return s
}

func (*Comment) AppendChild(newChild Node) (n Node) {
	macro.Rewrite("$_.appendChild($1)", newChild)
	return n
}

func (*Comment) CloneNode(deep *bool) (n Node) {
	macro.Rewrite("$_.cloneNode($1)", deep)
	return n
}

func (*Comment) CompareDocumentPosition(other Node) (u uint8) {
	macro.Rewrite("$_.compareDocumentPosition($1)", other)
	return u
}

func (*Comment) Contains(child Node) (b bool) {
	macro.Rewrite("$_.contains($1)", child)
	return b
}

func (*Comment) HasAttributes() (b bool) {
	macro.Rewrite("$_.hasAttributes()")
	return b
}

func (*Comment) HasChildNodes() (b bool) {
	macro.Rewrite("$_.hasChildNodes()")
	return b
}

func (*Comment) InsertBefore(newChild Node, refChild *Node) (n Node) {
	macro.Rewrite("$_.insertBefore($1, $2)", newChild, refChild)
	return n
}

func (*Comment) IsDefaultNamespace(namespaceURI string) (b bool) {
	macro.Rewrite("$_.isDefaultNamespace($1)", namespaceURI)
	return b
}

func (*Comment) IsEqualNode(arg Node) (b bool) {
	macro.Rewrite("$_.isEqualNode($1)", arg)
	return b
}

func (*Comment) IsSameNode(other Node) (b bool) {
	macro.Rewrite("$_.isSameNode($1)", other)
	return b
}

func (*Comment) LookupNamespaceURI(prefix string) (s string) {
	macro.Rewrite("$_.lookupNamespaceURI($1)", prefix)
	return s
}

func (*Comment) LookupPrefix(namespaceURI string) (s string) {
	macro.Rewrite("$_.lookupPrefix($1)", namespaceURI)
	return s
}

func (*Comment) Normalize() {
	macro.Rewrite("$_.normalize()")
}

func (*Comment) RemoveChild(oldChild Node) (n Node) {
	macro.Rewrite("$_.removeChild($1)", oldChild)
	return n
}

func (*Comment) ReplaceChild(newChild Node, oldChild Node) (n Node) {
	macro.Rewrite("$_.replaceChild($1, $2)", newChild, oldChild)
	return n
}

func (*Comment) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*Comment) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*Comment) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*Comment) Text() (text string) {
	macro.Rewrite("$_.text")
	return text
}

func (*Comment) SetText(text string) {
	macro.Rewrite("$_.text = $1", text)
}

func (*Comment) Data() (data string) {
	macro.Rewrite("$_.data")
	return data
}

func (*Comment) SetData(data string) {
	macro.Rewrite("$_.data = $1", data)
}

func (*Comment) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}

func (*Comment) Attributes() (attributes *NamedNodeMap) {
	macro.Rewrite("$_.attributes")
	return attributes
}

func (*Comment) BaseURI() (baseURI string) {
	macro.Rewrite("$_.baseURI")
	return baseURI
}

func (*Comment) ChildNodes() (childNodes *NodeList) {
	macro.Rewrite("$_.childNodes")
	return childNodes
}

func (*Comment) FirstChild() (firstChild Node) {
	macro.Rewrite("$_.firstChild")
	return firstChild
}

func (*Comment) LastChild() (lastChild Node) {
	macro.Rewrite("$_.lastChild")
	return lastChild
}

func (*Comment) LocalName() (localName string) {
	macro.Rewrite("$_.localName")
	return localName
}

func (*Comment) NamespaceURI() (namespaceURI string) {
	macro.Rewrite("$_.namespaceURI")
	return namespaceURI
}

func (*Comment) NextSibling() (nextSibling Node) {
	macro.Rewrite("$_.nextSibling")
	return nextSibling
}

func (*Comment) NodeName() (nodeName string) {
	macro.Rewrite("$_.nodeName")
	return nodeName
}

func (*Comment) NodeType() (nodeType uint8) {
	macro.Rewrite("$_.nodeType")
	return nodeType
}

func (*Comment) NodeValue() (nodeValue string) {
	macro.Rewrite("$_.nodeValue")
	return nodeValue
}

func (*Comment) SetNodeValue(nodeValue string) {
	macro.Rewrite("$_.nodeValue = $1", nodeValue)
}

func (*Comment) OwnerDocument() (ownerDocument document.Document) {
	macro.Rewrite("$_.ownerDocument")
	return ownerDocument
}

func (*Comment) ParentElement() (parentElement html.HTMLElement) {
	macro.Rewrite("$_.parentElement")
	return parentElement
}

func (*Comment) ParentNode() (parentNode Node) {
	macro.Rewrite("$_.parentNode")
	return parentNode
}

func (*Comment) PreviousSibling() (previousSibling Node) {
	macro.Rewrite("$_.previousSibling")
	return previousSibling
}

func (*Comment) TextContent() (textContent string) {
	macro.Rewrite("$_.textContent")
	return textContent
}

func (*Comment) SetTextContent(textContent string) {
	macro.Rewrite("$_.textContent = $1", textContent)
}
