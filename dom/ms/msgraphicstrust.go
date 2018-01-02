package ms

import "github.com/matthewmueller/joy/macro"

type MSGraphicsTrust struct {
}

func (*MSGraphicsTrust) ConstrictionActive() (constrictionActive bool) {
	macro.Rewrite("$_.constrictionActive")
	return constrictionActive
}

func (*MSGraphicsTrust) Status() (status string) {
	macro.Rewrite("$_.status")
	return status
}
