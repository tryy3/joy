package utils

import "github.com/matthewmueller/joy/macro"

type XPathNSResolver struct {
}

func (*XPathNSResolver) LookupNamespaceURI(prefix string) (s string) {
	macro.Rewrite("$_.lookupNamespaceURI($1)", prefix)
	return s
}
