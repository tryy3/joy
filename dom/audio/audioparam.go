package audio

import "github.com/matthewmueller/joy/macro"

type AudioParam struct {
}

func (*AudioParam) CancelScheduledValues(startTime float32) (a *AudioParam) {
	macro.Rewrite("$_.cancelScheduledValues($1)", startTime)
	return a
}

func (*AudioParam) ExponentialRampToValueAtTime(value float32, endTime float32) (a *AudioParam) {
	macro.Rewrite("$_.exponentialRampToValueAtTime($1, $2)", value, endTime)
	return a
}

func (*AudioParam) LinearRampToValueAtTime(value float32, endTime float32) (a *AudioParam) {
	macro.Rewrite("$_.linearRampToValueAtTime($1, $2)", value, endTime)
	return a
}

func (*AudioParam) SetTargetAtTime(target float32, startTime float32, timeConstant float32) (a *AudioParam) {
	macro.Rewrite("$_.setTargetAtTime($1, $2, $3)", target, startTime, timeConstant)
	return a
}

func (*AudioParam) SetValueAtTime(value float32, startTime float32) (a *AudioParam) {
	macro.Rewrite("$_.setValueAtTime($1, $2)", value, startTime)
	return a
}

func (*AudioParam) SetValueCurveAtTime(values []float32, startTime float32, duration float32) (a *AudioParam) {
	macro.Rewrite("$_.setValueCurveAtTime($1, $2, $3)", values, startTime, duration)
	return a
}

func (*AudioParam) DefaultValue() (defaultValue float32) {
	macro.Rewrite("$_.defaultValue")
	return defaultValue
}

func (*AudioParam) Value() (value float32) {
	macro.Rewrite("$_.value")
	return value
}

func (*AudioParam) SetValue(value float32) {
	macro.Rewrite("$_.value = $1", value)
}
