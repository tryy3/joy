package mediastreams

import "github.com/matthewmueller/joy/macro"

type MediaStreamError struct {
}

func (*MediaStreamError) ConstraintName() (constraintName string) {
	macro.Rewrite("$_.constraintName")
	return constraintName
}

func (*MediaStreamError) Message() (message string) {
	macro.Rewrite("$_.message")
	return message
}

func (*MediaStreamError) Name() (name string) {
	macro.Rewrite("$_.name")
	return name
}
