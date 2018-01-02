package ms

import "github.com/matthewmueller/joy/macro"

type MSFIDOSignature struct {
}

func (*MSFIDOSignature) AuthnrData() (authnrData string) {
	macro.Rewrite("$_.authnrData")
	return authnrData
}

func (*MSFIDOSignature) ClientData() (clientData string) {
	macro.Rewrite("$_.clientData")
	return clientData
}

func (*MSFIDOSignature) Signature() (signature string) {
	macro.Rewrite("$_.signature")
	return signature
}
