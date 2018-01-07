package ms

import "github.com/matthewmueller/joy/macro"

// MSFIDOSignature struct
// js:"MSFIDOSignature,omit"
type MSFIDOSignature struct {
}

// AuthnrData prop
// js:"authnrData"
func (*MSFIDOSignature) AuthnrData() (authnrData string) {
	macro.Rewrite("$_.authnrData")
	return authnrData
}

// ClientData prop
// js:"clientData"
func (*MSFIDOSignature) ClientData() (clientData string) {
	macro.Rewrite("$_.clientData")
	return clientData
}

// Signature prop
// js:"signature"
func (*MSFIDOSignature) Signature() (signature string) {
	macro.Rewrite("$_.signature")
	return signature
}
