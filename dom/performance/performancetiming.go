package performance

import "github.com/matthewmueller/joy/macro"

type PerformanceTiming struct {
}

func (*PerformanceTiming) ToJSON() (i interface{}) {
	macro.Rewrite("$_.toJSON()")
	return i
}

func (*PerformanceTiming) ConnectEnd() (connectEnd uint64) {
	macro.Rewrite("$_.connectEnd")
	return connectEnd
}

func (*PerformanceTiming) ConnectStart() (connectStart uint64) {
	macro.Rewrite("$_.connectStart")
	return connectStart
}

func (*PerformanceTiming) DomainLookupEnd() (domainLookupEnd uint64) {
	macro.Rewrite("$_.domainLookupEnd")
	return domainLookupEnd
}

func (*PerformanceTiming) DomainLookupStart() (domainLookupStart uint64) {
	macro.Rewrite("$_.domainLookupStart")
	return domainLookupStart
}

func (*PerformanceTiming) DomComplete() (domComplete uint64) {
	macro.Rewrite("$_.domComplete")
	return domComplete
}

func (*PerformanceTiming) DomContentLoadedEventEnd() (domContentLoadedEventEnd uint64) {
	macro.Rewrite("$_.domContentLoadedEventEnd")
	return domContentLoadedEventEnd
}

func (*PerformanceTiming) DomContentLoadedEventStart() (domContentLoadedEventStart uint64) {
	macro.Rewrite("$_.domContentLoadedEventStart")
	return domContentLoadedEventStart
}

func (*PerformanceTiming) DomInteractive() (domInteractive uint64) {
	macro.Rewrite("$_.domInteractive")
	return domInteractive
}

func (*PerformanceTiming) DomLoading() (domLoading uint64) {
	macro.Rewrite("$_.domLoading")
	return domLoading
}

func (*PerformanceTiming) FetchStart() (fetchStart uint64) {
	macro.Rewrite("$_.fetchStart")
	return fetchStart
}

func (*PerformanceTiming) LoadEventEnd() (loadEventEnd uint64) {
	macro.Rewrite("$_.loadEventEnd")
	return loadEventEnd
}

func (*PerformanceTiming) LoadEventStart() (loadEventStart uint64) {
	macro.Rewrite("$_.loadEventStart")
	return loadEventStart
}

func (*PerformanceTiming) MsFirstPaint() (msFirstPaint uint64) {
	macro.Rewrite("$_.msFirstPaint")
	return msFirstPaint
}

func (*PerformanceTiming) NavigationStart() (navigationStart uint64) {
	macro.Rewrite("$_.navigationStart")
	return navigationStart
}

func (*PerformanceTiming) RedirectEnd() (redirectEnd uint64) {
	macro.Rewrite("$_.redirectEnd")
	return redirectEnd
}

func (*PerformanceTiming) RedirectStart() (redirectStart uint64) {
	macro.Rewrite("$_.redirectStart")
	return redirectStart
}

func (*PerformanceTiming) RequestStart() (requestStart uint64) {
	macro.Rewrite("$_.requestStart")
	return requestStart
}

func (*PerformanceTiming) ResponseEnd() (responseEnd uint64) {
	macro.Rewrite("$_.responseEnd")
	return responseEnd
}

func (*PerformanceTiming) ResponseStart() (responseStart uint64) {
	macro.Rewrite("$_.responseStart")
	return responseStart
}

func (*PerformanceTiming) UnloadEventEnd() (unloadEventEnd uint64) {
	macro.Rewrite("$_.unloadEventEnd")
	return unloadEventEnd
}

func (*PerformanceTiming) UnloadEventStart() (unloadEventStart uint64) {
	macro.Rewrite("$_.unloadEventStart")
	return unloadEventStart
}
