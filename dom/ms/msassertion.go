package ms


// MSAssertion interface
// js:"MSAssertion"
type MSAssertion interface {

	// ID prop
	// js:"id"
	// jsrewrite:"$_.id"
	ID() (id string)

	// Type prop
	// js:"type"
	// jsrewrite:"$_.type"
	Type() (kind *MSCredentialType)
}
