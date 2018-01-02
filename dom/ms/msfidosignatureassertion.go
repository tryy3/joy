package ms

import (
	"github.com/matthewmueller/joy/macro"
)

var _ MSAssertion = (*MSFIDOSignatureAssertion)(nil)

type MSFIDOSignatureAssertion struct {
}

func (*MSFIDOSignatureAssertion) Signature() (signature *MSFIDOSignature) {
	macro.Rewrite("$_.signature")
	return signature
}

func (*MSFIDOSignatureAssertion) ID() (id string) {
	macro.Rewrite("$_.id")
	return id
}

func (*MSFIDOSignatureAssertion) Type() (kind *MSCredentialType) {
	macro.Rewrite("$_.type")
	return kind
}
