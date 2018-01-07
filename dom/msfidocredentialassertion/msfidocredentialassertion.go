package msfidocredentialassertion

import (
	"github.com/matthewmueller/joy/dom/ms"
	"github.com/matthewmueller/joy/macro"
)

var _ ms.MSAssertion = (*MSFIDOCredentialAssertion)(nil)

// MSFIDOCredentialAssertion struct
// js:"MSFIDOCredentialAssertion,omit"
type MSFIDOCredentialAssertion struct {
}

// Algorithm prop
// js:"algorithm"
func (*MSFIDOCredentialAssertion) Algorithm() (algorithm interface{}) {
	macro.Rewrite("$_.algorithm")
	return algorithm
}

// Attestation prop
// js:"attestation"
func (*MSFIDOCredentialAssertion) Attestation() (attestation bool) {
	macro.Rewrite("$_.attestation")
	return attestation
}

// PublicKey prop
// js:"publicKey"
func (*MSFIDOCredentialAssertion) PublicKey() (publicKey string) {
	macro.Rewrite("$_.publicKey")
	return publicKey
}

// TransportHints prop
// js:"transportHints"
func (*MSFIDOCredentialAssertion) TransportHints() (transportHints []*ms.MSTransportType) {
	macro.Rewrite("$_.transportHints")
	return transportHints
}

// ID prop
// js:"id"
func (*MSFIDOCredentialAssertion) ID() (id string) {
	macro.Rewrite("$_.id")
	return id
}

// Type prop
// js:"type"
func (*MSFIDOCredentialAssertion) Type() (kind *ms.MSCredentialType) {
	macro.Rewrite("$_.type")
	return kind
}
