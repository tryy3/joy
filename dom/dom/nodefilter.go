package dom

import "github.com/matthewmueller/joy/macro"

type NodeFilter struct {
}

func (*NodeFilter) AcceptNode(n Node) (i int8) {
	macro.Rewrite("$_.acceptNode($1)", n)
	return i
}
