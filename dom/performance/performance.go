package performance

import (
	"github.com/matthewmueller/joy/macro"
)

type Performance struct {
}

func (*Performance) ClearMarks(markName *string) {
	macro.Rewrite("$_.clearMarks($1)", markName)
}

func (*Performance) ClearMeasures(measureName *string) {
	macro.Rewrite("$_.clearMeasures($1)", measureName)
}

func (*Performance) ClearResourceTimings() {
	macro.Rewrite("$_.clearResourceTimings()")
}

func (*Performance) GetEntries() (i interface{}) {
	macro.Rewrite("$_.getEntries()")
	return i
}

func (*Performance) GetEntriesByName(name string, entryType *string) (i interface{}) {
	macro.Rewrite("$_.getEntriesByName($1, $2)", name, entryType)
	return i
}

func (*Performance) GetEntriesByType(entryType string) (i interface{}) {
	macro.Rewrite("$_.getEntriesByType($1)", entryType)
	return i
}

func (*Performance) GetMarks(markName *string) (i interface{}) {
	macro.Rewrite("$_.getMarks($1)", markName)
	return i
}

func (*Performance) GetMeasures(measureName *string) (i interface{}) {
	macro.Rewrite("$_.getMeasures($1)", measureName)
	return i
}

func (*Performance) Mark(markName string) {
	macro.Rewrite("$_.mark($1)", markName)
}

func (*Performance) Measure(measureName string, startMarkName *string, endMarkName *string) {
	macro.Rewrite("$_.measure($1, $2, $3)", measureName, startMarkName, endMarkName)
}

func (*Performance) Now() (i int) {
	macro.Rewrite("$_.now()")
	return i
}

func (*Performance) SetResourceTimingBufferSize(maxSize uint) {
	macro.Rewrite("$_.setResourceTimingBufferSize($1)", maxSize)
}

func (*Performance) ToJSON() (i interface{}) {
	macro.Rewrite("$_.toJSON()")
	return i
}

func (*Performance) Navigation() (navigation *PerformanceNavigation) {
	macro.Rewrite("$_.navigation")
	return navigation
}

func (*Performance) Timing() (timing *PerformanceTiming) {
	macro.Rewrite("$_.timing")
	return timing
}
