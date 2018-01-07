package authentication

import "github.com/matthewmueller/joy/dom/utils"

type ScopedCredentialDescriptor struct {
	id         []byte
	transports *[]*utils.Transport
	kind       *ScopedCredentialType
}
