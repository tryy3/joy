package performance

import (
	"github.com/matthewmueller/joy/macro"
)

var _ PerformanceEntry = (*PerformanceMark)(nil)

type PerformanceMark struct {
}

func (*PerformanceMark) Duration() (duration int) {
	macro.Rewrite("$_.duration")
	return duration
}

func (*PerformanceMark) EntryType() (entryType string) {
	macro.Rewrite("$_.entryType")
	return entryType
}

func (*PerformanceMark) Name() (name string) {
	macro.Rewrite("$_.name")
	return name
}

func (*PerformanceMark) StartTime() (startTime int) {
	macro.Rewrite("$_.startTime")
	return startTime
}
