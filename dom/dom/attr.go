package dom


// and EventTarget (https://developer.mozilla.org/en-US/docs/Web/API/EventTarget)
package attr

import (
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/node"
	"github.com/matthewmueller/joy/macro"
)

var _ Node = (*Attr)(nil)
var _ EventTarget = (*Attr)(nil)

// Attr type represents a DOM element's attribute as an object.
// In most DOM methods, you will probably directly retrieve the attribute as a string
// (e.g., Element.getAttribute(), but certain functions (e.g., Element.getAttributeNode())
// or means of iterating give Attr types.
// js:"Attr,omit"
type Attr struct {
}

// AppendChild use Attr.SetValue() instead
// js:"appendChild"
func (*Attr) AppendChild(newChild node.Node) (n node.Node) {
	macro.Rewrite("$_.appendChild($1)", newChild)
	return n
}

// CloneNode is depecrated and you probably shouldn't have used it before
// js:"cloneNode"
func (*Attr) CloneNode(deep *bool) (n node.Node) {
	macro.Rewrite("$_.cloneNode($1)", deep)
	return n
}

// CompareDocumentPosition compares the position of the current node against another node in any other document.
// js:"compareDocumentPosition"
func (*Attr) CompareDocumentPosition(other node.Node) (u uint8) {
	macro.Rewrite("$_.compareDocumentPosition($1)", other)
	return u
}

// Contains returns a Boolean value indicating whether a node is a descendant of a given node or not.
// js:"contains"
func (*Attr) Contains(child node.Node) (b bool) {
	macro.Rewrite("$_.contains($1)", child)
	return b
}

// HasAttributes this method now always returns false.
// js:"hasAttributes"
func (*Attr) HasAttributes() (b bool) {
	macro.Rewrite("$_.hasAttributes()")
	return b
}

// HasChildNodes this method now always returns false.
// js:"hasChildNodes"
func (*Attr) HasChildNodes() (b bool) {
	macro.Rewrite("$_.hasChildNodes()")
	return b
}

// InsertBefore use Attr.SetValue() instead
// js:"insertBefore"
func (*Attr) InsertBefore(newChild Node, refChild *node.Node) (n node.Node) {
	macro.Rewrite("$_.insertBefore($1, $2)", newChild, refChild)
	return n
}

// IsDefaultNamespace accepts a namespace URI as an argument
// and returns a Boolean with a value of true if the namespace
// is the default namespace on the given node or false if not.
// js:"isDefaultNamespace"
func (*Attr) IsDefaultNamespace(namespaceURI string) (b bool) {
	macro.Rewrite("$_.isDefaultNamespace($1)", namespaceURI)
	return b
}

// IsEqualNode returns a Boolean which indicates whether or not two
// nodes are of the same type and all their defining data points match.
// js:"isEqualNode"
func (*Attr) IsEqualNode(arg node.Node) (b bool) {
	macro.Rewrite("$_.isEqualNode($1)", arg)
	return b
}

// IsSameNode returns a Boolean value indicating whether or not the two
// nodes are the same (that is, they reference the same object).
// js:"isSameNode"
func (*Attr) IsSameNode(other node.Node) (b bool) {
	macro.Rewrite("$_.isSameNode($1)", other)
	return b
}

// LookupNamespaceURI accepts a prefix and returns the namespace URI associated
// with it on the given node if found (and null if not).
// Supplying null for the prefix will return the default namespace.
// js:"lookupNamespaceURI"
func (*Attr) LookupNamespaceURI(prefix string) (s string) {
	macro.Rewrite("$_.lookupNamespaceURI($1)", prefix)
	return s
}

// LookupPrefix returns a string containing the prefix for a given namespace URI,
// if present, and null if not. When multiple prefixes are possible,
// the result is implementation-dependent.
// js:"lookupPrefix"
func (*Attr) LookupPrefix(namespaceURI string) (s string) {
	macro.Rewrite("$_.lookupPrefix($1)", namespaceURI)
	return s
}

// Normalize is depecrated and you probably shouldn't have used it before
// js:"normalize"
func (*Attr) Normalize() {
	macro.Rewrite("$_.normalize()")
}

// RemoveChild use Attr.SetValue() instead
// js:"removeChild"
func (*Attr) RemoveChild(oldChild node.Node) (n node.Node) {
	macro.Rewrite("$_.removeChild($1)", oldChild)
	return n
}

// ReplaceChild use Attr.SetValue() instead
// js:"replaceChild"
func (*Attr) ReplaceChild(newChild node.Node, oldChild node.Node) (n node.Node) {
	macro.Rewrite("$_.replaceChild($1, $2)", newChild, oldChild)
	return n
}

// AddEventListener register an event handler of a specific event type on the event.EventTarget.
// js:"addEventListener"
func (*Attr) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent dispatch an event to this event.EventTarget.
// js:"dispatchEvent"
func (*Attr) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener removes an event listener from the event.EventTarget.
// js:"removeEventListener"
func (*Attr) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Name the attribute's name.
// js:"name"
func (*Attr) Name() (name string) {
	macro.Rewrite("$_.name")
	return name
}

// OwnerElement prop
// js:"ownerElement"
func (*Attr) OwnerElement() (ownerElement Element) {
	macro.Rewrite("$_.ownerElement")
	return ownerElement
}

// Prefix prop
// js:"prefix"
func (*Attr) Prefix() (prefix string) {
	macro.Rewrite("$_.prefix")
	return prefix
}

// Specified prop
// js:"specified"
func (*Attr) Specified() (specified bool) {
	macro.Rewrite("$_.specified")
	return specified
}

// Value prop
// js:"value"
func (*Attr) Value() (value string) {
	macro.Rewrite("$_.value")
	return value
}

// SetValue prop
// js:"value"
func (*Attr) SetValue(value string) {
	macro.Rewrite("$_.value = $1", value)
}

// Attributes prop
// js:"attributes"
func (*Attr) Attributes() (attributes *NamedNodeMap) {
	macro.Rewrite("$_.attributes")
	return attributes
}

// BaseURI prop
// js:"baseURI"
func (*Attr) BaseURI() (baseURI string) {
	macro.Rewrite("$_.baseURI")
	return baseURI
}

// ChildNodes prop
// js:"childNodes"
func (*Attr) ChildNodes() (childNodes *NodeList) {
	macro.Rewrite("$_.childNodes")
	return childNodes
}

// FirstChild prop
// js:"firstChild"
func (*Attr) FirstChild() (firstChild Node) {
	macro.Rewrite("$_.firstChild")
	return firstChild
}

// LastChild prop
// js:"lastChild"
func (*Attr) LastChild() (lastChild Node) {
	macro.Rewrite("$_.lastChild")
	return lastChild
}

// LocalName prop
// js:"localName"
func (*Attr) LocalName() (localName string) {
	macro.Rewrite("$_.localName")
	return localName
}

// NamespaceURI prop
// js:"namespaceURI"
func (*Attr) NamespaceURI() (namespaceURI string) {
	macro.Rewrite("$_.namespaceURI")
	return namespaceURI
}

// NextSibling prop
// js:"nextSibling"
func (*Attr) NextSibling() (nextSibling Node) {
	macro.Rewrite("$_.nextSibling")
	return nextSibling
}

// NodeName prop
// js:"nodeName"
func (*Attr) NodeName() (nodeName string) {
	macro.Rewrite("$_.nodeName")
	return nodeName
}

// NodeType prop
// js:"nodeType"
func (*Attr) NodeType() (nodeType uint8) {
	macro.Rewrite("$_.nodeType")
	return nodeType
}

// NodeValue prop
// js:"nodeValue"
func (*Attr) NodeValue() (nodeValue string) {
	macro.Rewrite("$_.nodeValue")
	return nodeValue
}

// SetNodeValue prop
// js:"nodeValue"
func (*Attr) SetNodeValue(nodeValue string) {
	macro.Rewrite("$_.nodeValue = $1", nodeValue)
}

// OwnerDocument prop
// js:"ownerDocument"
func (*Attr) OwnerDocument() (ownerDocument Document) {
	macro.Rewrite("$_.ownerDocument")
	return ownerDocument
}

// ParentElement prop
// js:"parentElement"
func (*Attr) ParentElement() (parentElement HTMLElement) {
	macro.Rewrite("$_.parentElement")
	return parentElement
}

// ParentNode prop
// js:"parentNode"
func (*Attr) ParentNode() (parentNode Node) {
	macro.Rewrite("$_.parentNode")
	return parentNode
}

// PreviousSibling prop
// js:"previousSibling"
func (*Attr) PreviousSibling() (previousSibling Node) {
	macro.Rewrite("$_.previousSibling")
	return previousSibling
}

// TextContent prop
// js:"textContent"
func (*Attr) TextContent() (textContent string) {
	macro.Rewrite("$_.textContent")
	return textContent
}

// SetTextContent prop
// js:"textContent"
func (*Attr) SetTextContent(textContent string) {
	macro.Rewrite("$_.textContent = $1", textContent)
}
