package ms

type MSFIDOCredentialParameters struct {
	*MSCredentialParameters

	algorithm	*interface{}
	authenticators	*[]string
}
