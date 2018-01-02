package utils

import "github.com/matthewmueller/joy/macro"

type ValidityState struct {
}

func (*ValidityState) BadInput() (badInput bool) {
	macro.Rewrite("$_.badInput")
	return badInput
}

func (*ValidityState) CustomError() (customError bool) {
	macro.Rewrite("$_.customError")
	return customError
}

func (*ValidityState) PatternMismatch() (patternMismatch bool) {
	macro.Rewrite("$_.patternMismatch")
	return patternMismatch
}

func (*ValidityState) RangeOverflow() (rangeOverflow bool) {
	macro.Rewrite("$_.rangeOverflow")
	return rangeOverflow
}

func (*ValidityState) RangeUnderflow() (rangeUnderflow bool) {
	macro.Rewrite("$_.rangeUnderflow")
	return rangeUnderflow
}

func (*ValidityState) StepMismatch() (stepMismatch bool) {
	macro.Rewrite("$_.stepMismatch")
	return stepMismatch
}

func (*ValidityState) TooLong() (tooLong bool) {
	macro.Rewrite("$_.tooLong")
	return tooLong
}

func (*ValidityState) TypeMismatch() (typeMismatch bool) {
	macro.Rewrite("$_.typeMismatch")
	return typeMismatch
}

func (*ValidityState) Valid() (valid bool) {
	macro.Rewrite("$_.valid")
	return valid
}

func (*ValidityState) ValueMissing() (valueMissing bool) {
	macro.Rewrite("$_.valueMissing")
	return valueMissing
}
