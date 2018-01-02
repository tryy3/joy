package dom

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/document"
)

type DOMImplementation struct {
}

func (*DOMImplementation) CreateDocument(namespaceURI string, qualifiedName string, doctype *document.DocumentType) (d document.Document) {
	macro.Rewrite("$_.createDocument($1, $2, $3)", namespaceURI, qualifiedName, doctype)
	return d
}

func (*DOMImplementation) CreateDocumentType(qualifiedName string, publicID string, systemID string) (d *document.DocumentType) {
	macro.Rewrite("$_.createDocumentType($1, $2, $3)", qualifiedName, publicID, systemID)
	return d
}

func (*DOMImplementation) CreateHTMLDocument(title string) (d document.Document) {
	macro.Rewrite("$_.createHTMLDocument($1)", title)
	return d
}

func (*DOMImplementation) HasFeature() (b bool) {
	macro.Rewrite("$_.hasFeature()")
	return b
}
