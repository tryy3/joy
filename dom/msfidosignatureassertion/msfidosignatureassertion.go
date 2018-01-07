package msfidosignatureassertion

import (
	"github.com/matthewmueller/joy/dom/ms"
	"github.com/matthewmueller/joy/macro"
)

var _ ms.MSAssertion = (*MSFIDOSignatureAssertion)(nil)

// MSFIDOSignatureAssertion struct
// js:"MSFIDOSignatureAssertion,omit"
type MSFIDOSignatureAssertion struct {
}

// Signature prop
// js:"signature"
func (*MSFIDOSignatureAssertion) Signature() (signature *ms.MSFIDOSignature) {
	macro.Rewrite("$_.signature")
	return signature
}

// ID prop
// js:"id"
func (*MSFIDOSignatureAssertion) ID() (id string) {
	macro.Rewrite("$_.id")
	return id
}

// Type prop
// js:"type"
func (*MSFIDOSignatureAssertion) Type() (kind *ms.MSCredentialType) {
	macro.Rewrite("$_.type")
	return kind
}
