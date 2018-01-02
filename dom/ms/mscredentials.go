package ms

import (
	"github.com/matthewmueller/joy/macro"
)

type MSCredentials struct {
}

func (*MSCredentials) GetAssertion(challenge string, filter *MSCredentialFilter, params *MSSignatureParameters) (m MSAssertion) {
	macro.Rewrite("await $_.getAssertion($1, $2, $3)", challenge, filter, params)
	return m
}

func (*MSCredentials) MakeCredential(accountInfo *MSAccountInfo, params []*MSCredentialParameters, challenge *string) (m MSAssertion) {
	macro.Rewrite("await $_.makeCredential($1, $2, $3)", accountInfo, params, challenge)
	return m
}
