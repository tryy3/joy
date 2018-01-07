package assertionoptions

import "github.com/matthewmueller/joy/dom/authentication"

type AssertionOptions struct {
	allowList      *[]*authentication.ScopedCredentialDescriptor
	extensions     *authentication.WebAuthnExtensions
	rpId           *string
	timeoutSeconds *uint
}
