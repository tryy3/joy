package authentication

import (
	"github.com/matthewmueller/joy/macro"
)

type ScopedCredential struct {
}

func (*ScopedCredential) ID() (id []byte) {
	macro.Rewrite("$_.id")
	return id
}

func (*ScopedCredential) Type() (kind *ScopedCredentialType) {
	macro.Rewrite("$_.type")
	return kind
}
