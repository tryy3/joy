package authentication

type ClientData struct {
	challenge    string
	extensions   *WebAuthnExtensions
	hashAlg      interface{}
	origin       string
	rpId         string
	tokenBinding *string
}
