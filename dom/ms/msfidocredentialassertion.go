package ms

import (
	"github.com/matthewmueller/joy/macro"
)

var _ MSAssertion = (*MSFIDOCredentialAssertion)(nil)

type MSFIDOCredentialAssertion struct {
}

func (*MSFIDOCredentialAssertion) Algorithm() (algorithm interface{}) {
	macro.Rewrite("$_.algorithm")
	return algorithm
}

func (*MSFIDOCredentialAssertion) Attestation() (attestation bool) {
	macro.Rewrite("$_.attestation")
	return attestation
}

func (*MSFIDOCredentialAssertion) PublicKey() (publicKey string) {
	macro.Rewrite("$_.publicKey")
	return publicKey
}

func (*MSFIDOCredentialAssertion) TransportHints() (transportHints []*MSTransportType) {
	macro.Rewrite("$_.transportHints")
	return transportHints
}

func (*MSFIDOCredentialAssertion) ID() (id string) {
	macro.Rewrite("$_.id")
	return id
}

func (*MSFIDOCredentialAssertion) Type() (kind *MSCredentialType) {
	macro.Rewrite("$_.type")
	return kind
}
