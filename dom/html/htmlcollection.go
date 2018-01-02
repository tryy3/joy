package html

import "github.com/matthewmueller/joy/dom/element"

type HTMLCollection interface {
	Item(index uint) (e element.Element)

	NamedItem(name string) (e element.Element)

	Length() (length uint)
}
