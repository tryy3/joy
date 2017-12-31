package ms


import (
	"github.com/matthewmueller/joy/macro"
)

// MSCredentials struct
// js:"MSCredentials,omit"
type MSCredentials struct {
}

// GetAssertion fn
// js:"getAssertion"
func (*MSCredentials) GetAssertion(challenge string, filter *MSCredentialFilter, params *MSSignatureParameters) (m MSAssertion) {
	macro.Rewrite("await $_.getAssertion($1, $2, $3)", challenge, filter, params)
	return m
}

// MakeCredential fn
// js:"makeCredential"
func (*MSCredentials) MakeCredential(accountInfo *MSAccountInfo, params []*MSCredentialParameters, challenge *string) (m MSAssertion) {
	macro.Rewrite("await $_.makeCredential($1, $2, $3)", accountInfo, params, challenge)
	return m
}
