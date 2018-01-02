package ms

type MSAssertion interface {
	ID() (id string)

	Type() (kind *MSCredentialType)
}
