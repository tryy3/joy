package dom

import "github.com/matthewmueller/joy/dom/element"

type NodeSelector interface {
	QuerySelector(selectors string) (e element.Element)

	QuerySelectorAll(selectors string) (n *NodeList)
}
