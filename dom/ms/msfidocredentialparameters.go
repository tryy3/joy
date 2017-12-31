package ms


import "github.com/matthewmueller/joy/dom/mscredentialparameters"

type MSFIDOCredentialParameters struct {
	*mscredentialparameters.MSCredentialParameters

	algorithm      *interface{}
	authenticators *[]string
}
