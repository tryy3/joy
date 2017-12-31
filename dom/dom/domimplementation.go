package dom


import (
	"github.com/matthewmueller/joy/dom/document"
	"github.com/matthewmueller/joy/macro"
)

// DOMImplementation struct
// js:"DOMImplementation,omit"
type DOMImplementation struct {
}

// CreateDocument fn
// js:"createDocument"
func (*DOMImplementation) CreateDocument(namespaceURI string, qualifiedName string, doctype *document.DocumentType) (d document.Document) {
	macro.Rewrite("$_.createDocument($1, $2, $3)", namespaceURI, qualifiedName, doctype)
	return d
}

// CreateDocumentType fn
// js:"createDocumentType"
func (*DOMImplementation) CreateDocumentType(qualifiedName string, publicID string, systemID string) (d *document.DocumentType) {
	macro.Rewrite("$_.createDocumentType($1, $2, $3)", qualifiedName, publicID, systemID)
	return d
}

// CreateHTMLDocument fn
// js:"createHTMLDocument"
func (*DOMImplementation) CreateHTMLDocument(title string) (d document.Document) {
	macro.Rewrite("$_.createHTMLDocument($1)", title)
	return d
}

// HasFeature fn
// js:"hasFeature"
func (*DOMImplementation) HasFeature() (b bool) {
	macro.Rewrite("$_.hasFeature()")
	return b
}
