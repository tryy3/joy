package dom

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/document"
	"github.com/matthewmueller/joy/dom/html"
)

var _ Text = (*CDATASection)(nil)
var _ CharacterData = (*CDATASection)(nil)
var _ ChildNode = (*CDATASection)(nil)
var _ Node = (*CDATASection)(nil)
var _ event.EventTarget = (*CDATASection)(nil)

type CDATASection struct {
}

func (*CDATASection) SplitText(offset uint) (t Text) {
	macro.Rewrite("$_.splitText($1)", offset)
	return t
}

func (*CDATASection) AppendData(arg string) {
	macro.Rewrite("$_.appendData($1)", arg)
}

func (*CDATASection) DeleteData(offset uint, count uint) {
	macro.Rewrite("$_.deleteData($1, $2)", offset, count)
}

func (*CDATASection) InsertData(offset uint, arg string) {
	macro.Rewrite("$_.insertData($1, $2)", offset, arg)
}

func (*CDATASection) ReplaceData(offset uint, count uint, arg string) {
	macro.Rewrite("$_.replaceData($1, $2, $3)", offset, count, arg)
}

func (*CDATASection) SubstringData(offset uint, count uint) (s string) {
	macro.Rewrite("$_.substringData($1, $2)", offset, count)
	return s
}

func (*CDATASection) AppendChild(newChild Node) (n Node) {
	macro.Rewrite("$_.appendChild($1)", newChild)
	return n
}

func (*CDATASection) CloneNode(deep *bool) (n Node) {
	macro.Rewrite("$_.cloneNode($1)", deep)
	return n
}

func (*CDATASection) CompareDocumentPosition(other Node) (u uint8) {
	macro.Rewrite("$_.compareDocumentPosition($1)", other)
	return u
}

func (*CDATASection) Contains(child Node) (b bool) {
	macro.Rewrite("$_.contains($1)", child)
	return b
}

func (*CDATASection) HasAttributes() (b bool) {
	macro.Rewrite("$_.hasAttributes()")
	return b
}

func (*CDATASection) HasChildNodes() (b bool) {
	macro.Rewrite("$_.hasChildNodes()")
	return b
}

func (*CDATASection) InsertBefore(newChild Node, refChild *Node) (n Node) {
	macro.Rewrite("$_.insertBefore($1, $2)", newChild, refChild)
	return n
}

func (*CDATASection) IsDefaultNamespace(namespaceURI string) (b bool) {
	macro.Rewrite("$_.isDefaultNamespace($1)", namespaceURI)
	return b
}

func (*CDATASection) IsEqualNode(arg Node) (b bool) {
	macro.Rewrite("$_.isEqualNode($1)", arg)
	return b
}

func (*CDATASection) IsSameNode(other Node) (b bool) {
	macro.Rewrite("$_.isSameNode($1)", other)
	return b
}

func (*CDATASection) LookupNamespaceURI(prefix string) (s string) {
	macro.Rewrite("$_.lookupNamespaceURI($1)", prefix)
	return s
}

func (*CDATASection) LookupPrefix(namespaceURI string) (s string) {
	macro.Rewrite("$_.lookupPrefix($1)", namespaceURI)
	return s
}

func (*CDATASection) Normalize() {
	macro.Rewrite("$_.normalize()")
}

func (*CDATASection) RemoveChild(oldChild Node) (n Node) {
	macro.Rewrite("$_.removeChild($1)", oldChild)
	return n
}

func (*CDATASection) ReplaceChild(newChild Node, oldChild Node) (n Node) {
	macro.Rewrite("$_.replaceChild($1, $2)", newChild, oldChild)
	return n
}

func (*CDATASection) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*CDATASection) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*CDATASection) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*CDATASection) WholeText() (wholeText string) {
	macro.Rewrite("$_.wholeText")
	return wholeText
}

func (*CDATASection) Data() (data string) {
	macro.Rewrite("$_.data")
	return data
}

func (*CDATASection) SetData(data string) {
	macro.Rewrite("$_.data = $1", data)
}

func (*CDATASection) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}

func (*CDATASection) Attributes() (attributes *NamedNodeMap) {
	macro.Rewrite("$_.attributes")
	return attributes
}

func (*CDATASection) BaseURI() (baseURI string) {
	macro.Rewrite("$_.baseURI")
	return baseURI
}

func (*CDATASection) ChildNodes() (childNodes *NodeList) {
	macro.Rewrite("$_.childNodes")
	return childNodes
}

func (*CDATASection) FirstChild() (firstChild Node) {
	macro.Rewrite("$_.firstChild")
	return firstChild
}

func (*CDATASection) LastChild() (lastChild Node) {
	macro.Rewrite("$_.lastChild")
	return lastChild
}

func (*CDATASection) LocalName() (localName string) {
	macro.Rewrite("$_.localName")
	return localName
}

func (*CDATASection) NamespaceURI() (namespaceURI string) {
	macro.Rewrite("$_.namespaceURI")
	return namespaceURI
}

func (*CDATASection) NextSibling() (nextSibling Node) {
	macro.Rewrite("$_.nextSibling")
	return nextSibling
}

func (*CDATASection) NodeName() (nodeName string) {
	macro.Rewrite("$_.nodeName")
	return nodeName
}

func (*CDATASection) NodeType() (nodeType uint8) {
	macro.Rewrite("$_.nodeType")
	return nodeType
}

func (*CDATASection) NodeValue() (nodeValue string) {
	macro.Rewrite("$_.nodeValue")
	return nodeValue
}

func (*CDATASection) SetNodeValue(nodeValue string) {
	macro.Rewrite("$_.nodeValue = $1", nodeValue)
}

func (*CDATASection) OwnerDocument() (ownerDocument document.Document) {
	macro.Rewrite("$_.ownerDocument")
	return ownerDocument
}

func (*CDATASection) ParentElement() (parentElement html.HTMLElement) {
	macro.Rewrite("$_.parentElement")
	return parentElement
}

func (*CDATASection) ParentNode() (parentNode Node) {
	macro.Rewrite("$_.parentNode")
	return parentNode
}

func (*CDATASection) PreviousSibling() (previousSibling Node) {
	macro.Rewrite("$_.previousSibling")
	return previousSibling
}

func (*CDATASection) TextContent() (textContent string) {
	macro.Rewrite("$_.textContent")
	return textContent
}

func (*CDATASection) SetTextContent(textContent string) {
	macro.Rewrite("$_.textContent = $1", textContent)
}
