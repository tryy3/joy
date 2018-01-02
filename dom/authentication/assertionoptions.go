package authentication

type AssertionOptions struct {
	allowList	*[]*ScopedCredentialDescriptor
	extensions	*WebAuthnExtensions
	rpId		*string
	timeoutSeconds	*uint
}
