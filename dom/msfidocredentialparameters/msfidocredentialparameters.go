package msfidocredentialparameters

import "github.com/matthewmueller/joy/dom/ms"

type MSFIDOCredentialParameters struct {
	*ms.MSCredentialParameters

	algorithm      *interface{}
	authenticators *[]string
}
