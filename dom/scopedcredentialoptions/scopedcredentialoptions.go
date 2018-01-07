package scopedcredentialoptions

import "github.com/matthewmueller/joy/dom/authentication"

type ScopedCredentialOptions struct {
	excludeList    *[]*authentication.ScopedCredentialDescriptor
	extensions     *authentication.WebAuthnExtensions
	rpId           *string
	timeoutSeconds *uint
}
