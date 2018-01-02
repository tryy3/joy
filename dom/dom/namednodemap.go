package dom

import "github.com/matthewmueller/joy/macro"

type NamedNodeMap struct {
}

func (*NamedNodeMap) GetNamedItem(name string) (a *Attr) {
	macro.Rewrite("$_.getNamedItem($1)", name)
	return a
}

func (*NamedNodeMap) GetNamedItemNS(namespaceURI string, localName string) (a *Attr) {
	macro.Rewrite("$_.getNamedItemNS($1, $2)", namespaceURI, localName)
	return a
}

func (*NamedNodeMap) Item(index uint) (a *Attr) {
	macro.Rewrite("$_.item($1)", index)
	return a
}

func (*NamedNodeMap) RemoveNamedItem(name string) (a *Attr) {
	macro.Rewrite("$_.removeNamedItem($1)", name)
	return a
}

func (*NamedNodeMap) RemoveNamedItemNS(namespaceURI string, localName string) (a *Attr) {
	macro.Rewrite("$_.removeNamedItemNS($1, $2)", namespaceURI, localName)
	return a
}

func (*NamedNodeMap) SetNamedItem(arg *Attr) (a *Attr) {
	macro.Rewrite("$_.setNamedItem($1)", arg)
	return a
}

func (*NamedNodeMap) SetNamedItemNS(arg *Attr) (a *Attr) {
	macro.Rewrite("$_.setNamedItemNS($1)", arg)
	return a
}

func (*NamedNodeMap) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}
