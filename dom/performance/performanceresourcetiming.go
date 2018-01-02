package performance

import (
	"github.com/matthewmueller/joy/macro"
)

var _ PerformanceEntry = (*PerformanceResourceTiming)(nil)

type PerformanceResourceTiming struct {
}

func (*PerformanceResourceTiming) ConnectEnd() (connectEnd int) {
	macro.Rewrite("$_.connectEnd")
	return connectEnd
}

func (*PerformanceResourceTiming) ConnectStart() (connectStart int) {
	macro.Rewrite("$_.connectStart")
	return connectStart
}

func (*PerformanceResourceTiming) DomainLookupEnd() (domainLookupEnd int) {
	macro.Rewrite("$_.domainLookupEnd")
	return domainLookupEnd
}

func (*PerformanceResourceTiming) DomainLookupStart() (domainLookupStart int) {
	macro.Rewrite("$_.domainLookupStart")
	return domainLookupStart
}

func (*PerformanceResourceTiming) FetchStart() (fetchStart int) {
	macro.Rewrite("$_.fetchStart")
	return fetchStart
}

func (*PerformanceResourceTiming) InitiatorType() (initiatorType string) {
	macro.Rewrite("$_.initiatorType")
	return initiatorType
}

func (*PerformanceResourceTiming) RedirectEnd() (redirectEnd int) {
	macro.Rewrite("$_.redirectEnd")
	return redirectEnd
}

func (*PerformanceResourceTiming) RedirectStart() (redirectStart int) {
	macro.Rewrite("$_.redirectStart")
	return redirectStart
}

func (*PerformanceResourceTiming) RequestStart() (requestStart int) {
	macro.Rewrite("$_.requestStart")
	return requestStart
}

func (*PerformanceResourceTiming) ResponseEnd() (responseEnd int) {
	macro.Rewrite("$_.responseEnd")
	return responseEnd
}

func (*PerformanceResourceTiming) ResponseStart() (responseStart int) {
	macro.Rewrite("$_.responseStart")
	return responseStart
}

func (*PerformanceResourceTiming) Duration() (duration int) {
	macro.Rewrite("$_.duration")
	return duration
}

func (*PerformanceResourceTiming) EntryType() (entryType string) {
	macro.Rewrite("$_.entryType")
	return entryType
}

func (*PerformanceResourceTiming) Name() (name string) {
	macro.Rewrite("$_.name")
	return name
}

func (*PerformanceResourceTiming) StartTime() (startTime int) {
	macro.Rewrite("$_.startTime")
	return startTime
}
