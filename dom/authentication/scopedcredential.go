package authentication

import "github.com/matthewmueller/joy/macro"

// ScopedCredential struct
// js:"ScopedCredential,omit"
type ScopedCredential struct {
}

// ID prop
// js:"id"
func (*ScopedCredential) ID() (id []byte) {
	macro.Rewrite("$_.id")
	return id
}

// Type prop
// js:"type"
func (*ScopedCredential) Type() (kind *ScopedCredentialType) {
	macro.Rewrite("$_.type")
	return kind
}
