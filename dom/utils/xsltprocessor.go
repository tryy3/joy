package utils

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/dom"
	"github.com/matthewmueller/joy/dom/document"
)

func New() *XSLTProcessor {
	macro.Rewrite("new XSLTProcessor()")
	return &XSLTProcessor{}
}

type XSLTProcessor struct {
}

func (*XSLTProcessor) ClearParameters() {
	macro.Rewrite("$_.clearParameters()")
}

func (*XSLTProcessor) GetParameter(namespaceURI string, localName string) (i interface{}) {
	macro.Rewrite("$_.getParameter($1, $2)", namespaceURI, localName)
	return i
}

func (*XSLTProcessor) ImportStylesheet(style dom.Node) {
	macro.Rewrite("$_.importStylesheet($1)", style)
}

func (*XSLTProcessor) RemoveParameter(namespaceURI string, localName string) {
	macro.Rewrite("$_.removeParameter($1, $2)", namespaceURI, localName)
}

func (*XSLTProcessor) Reset() {
	macro.Rewrite("$_.reset()")
}

func (*XSLTProcessor) SetParameter(namespaceURI string, localName string, value interface{}) {
	macro.Rewrite("$_.setParameter($1, $2, $3)", namespaceURI, localName, value)
}

func (*XSLTProcessor) TransformToDocument(source dom.Node) (w document.Document) {
	macro.Rewrite("$_.transformToDocument($1)", source)
	return w
}

func (*XSLTProcessor) TransformToFragment(source dom.Node, document document.Document) (w *document.DocumentFragment) {
	macro.Rewrite("$_.transformToFragment($1, $2)", source, document)
	return w
}
