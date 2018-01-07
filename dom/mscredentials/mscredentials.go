package mscredentials

import (
	"github.com/matthewmueller/joy/dom/ms"
	"github.com/matthewmueller/joy/dom/mscredentialfilter"
	"github.com/matthewmueller/joy/macro"
)

// MSCredentials struct
// js:"MSCredentials,omit"
type MSCredentials struct {
}

// GetAssertion fn
// js:"getAssertion"
func (*MSCredentials) GetAssertion(challenge string, filter *mscredentialfilter.MSCredentialFilter, params *ms.MSSignatureParameters) (m ms.MSAssertion) {
	macro.Rewrite("await $_.getAssertion($1, $2, $3)", challenge, filter, params)
	return m
}

// MakeCredential fn
// js:"makeCredential"
func (*MSCredentials) MakeCredential(accountInfo *ms.MSAccountInfo, params []*ms.MSCredentialParameters, challenge *string) (m ms.MSAssertion) {
	macro.Rewrite("await $_.makeCredential($1, $2, $3)", accountInfo, params, challenge)
	return m
}
