package performance

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/navigation"
)

var _ PerformanceEntry = (*PerformanceNavigationTiming)(nil)

type PerformanceNavigationTiming struct {
}

func (*PerformanceNavigationTiming) ConnectEnd() (connectEnd int) {
	macro.Rewrite("$_.connectEnd")
	return connectEnd
}

func (*PerformanceNavigationTiming) ConnectStart() (connectStart int) {
	macro.Rewrite("$_.connectStart")
	return connectStart
}

func (*PerformanceNavigationTiming) DomainLookupEnd() (domainLookupEnd int) {
	macro.Rewrite("$_.domainLookupEnd")
	return domainLookupEnd
}

func (*PerformanceNavigationTiming) DomainLookupStart() (domainLookupStart int) {
	macro.Rewrite("$_.domainLookupStart")
	return domainLookupStart
}

func (*PerformanceNavigationTiming) DomComplete() (domComplete int) {
	macro.Rewrite("$_.domComplete")
	return domComplete
}

func (*PerformanceNavigationTiming) DomContentLoadedEventEnd() (domContentLoadedEventEnd int) {
	macro.Rewrite("$_.domContentLoadedEventEnd")
	return domContentLoadedEventEnd
}

func (*PerformanceNavigationTiming) DomContentLoadedEventStart() (domContentLoadedEventStart int) {
	macro.Rewrite("$_.domContentLoadedEventStart")
	return domContentLoadedEventStart
}

func (*PerformanceNavigationTiming) DomInteractive() (domInteractive int) {
	macro.Rewrite("$_.domInteractive")
	return domInteractive
}

func (*PerformanceNavigationTiming) DomLoading() (domLoading int) {
	macro.Rewrite("$_.domLoading")
	return domLoading
}

func (*PerformanceNavigationTiming) FetchStart() (fetchStart int) {
	macro.Rewrite("$_.fetchStart")
	return fetchStart
}

func (*PerformanceNavigationTiming) LoadEventEnd() (loadEventEnd int) {
	macro.Rewrite("$_.loadEventEnd")
	return loadEventEnd
}

func (*PerformanceNavigationTiming) LoadEventStart() (loadEventStart int) {
	macro.Rewrite("$_.loadEventStart")
	return loadEventStart
}

func (*PerformanceNavigationTiming) NavigationStart() (navigationStart int) {
	macro.Rewrite("$_.navigationStart")
	return navigationStart
}

func (*PerformanceNavigationTiming) RedirectCount() (redirectCount uint8) {
	macro.Rewrite("$_.redirectCount")
	return redirectCount
}

func (*PerformanceNavigationTiming) RedirectEnd() (redirectEnd int) {
	macro.Rewrite("$_.redirectEnd")
	return redirectEnd
}

func (*PerformanceNavigationTiming) RedirectStart() (redirectStart int) {
	macro.Rewrite("$_.redirectStart")
	return redirectStart
}

func (*PerformanceNavigationTiming) RequestStart() (requestStart int) {
	macro.Rewrite("$_.requestStart")
	return requestStart
}

func (*PerformanceNavigationTiming) ResponseEnd() (responseEnd int) {
	macro.Rewrite("$_.responseEnd")
	return responseEnd
}

func (*PerformanceNavigationTiming) ResponseStart() (responseStart int) {
	macro.Rewrite("$_.responseStart")
	return responseStart
}

func (*PerformanceNavigationTiming) Type() (kind *navigation.NavigationType) {
	macro.Rewrite("$_.type")
	return kind
}

func (*PerformanceNavigationTiming) UnloadEventEnd() (unloadEventEnd int) {
	macro.Rewrite("$_.unloadEventEnd")
	return unloadEventEnd
}

func (*PerformanceNavigationTiming) UnloadEventStart() (unloadEventStart int) {
	macro.Rewrite("$_.unloadEventStart")
	return unloadEventStart
}

func (*PerformanceNavigationTiming) Duration() (duration int) {
	macro.Rewrite("$_.duration")
	return duration
}

func (*PerformanceNavigationTiming) EntryType() (entryType string) {
	macro.Rewrite("$_.entryType")
	return entryType
}

func (*PerformanceNavigationTiming) Name() (name string) {
	macro.Rewrite("$_.name")
	return name
}

func (*PerformanceNavigationTiming) StartTime() (startTime int) {
	macro.Rewrite("$_.startTime")
	return startTime
}
