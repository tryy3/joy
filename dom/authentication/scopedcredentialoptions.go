package authentication

type ScopedCredentialOptions struct {
	excludeList	*[]*ScopedCredentialDescriptor
	extensions	*WebAuthnExtensions
	rpId		*string
	timeoutSeconds	*uint
}
