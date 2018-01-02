package dom

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/document"
)

func New() *DOMParser {
	macro.Rewrite("new DOMParser()")
	return &DOMParser{}
}

type DOMParser struct {
}

func (*DOMParser) ParseFromString(source string, mimeType string) (w document.Document) {
	macro.Rewrite("$_.parseFromString($1, $2)", source, mimeType)
	return w
}
