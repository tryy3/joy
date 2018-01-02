package performance

import (
	"github.com/matthewmueller/joy/macro"
)

var _ PerformanceEntry = (*PerformanceMeasure)(nil)

type PerformanceMeasure struct {
}

func (*PerformanceMeasure) Duration() (duration int) {
	macro.Rewrite("$_.duration")
	return duration
}

func (*PerformanceMeasure) EntryType() (entryType string) {
	macro.Rewrite("$_.entryType")
	return entryType
}

func (*PerformanceMeasure) Name() (name string) {
	macro.Rewrite("$_.name")
	return name
}

func (*PerformanceMeasure) StartTime() (startTime int) {
	macro.Rewrite("$_.startTime")
	return startTime
}
