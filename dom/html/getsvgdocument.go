package html

import "github.com/matthewmueller/joy/dom/document"

type GetSVGDocument interface {
	GetSVGDocument() (w document.Document)
}
