package performance

import "github.com/matthewmueller/joy/macro"

type PerformanceNavigation struct {
}

func (*PerformanceNavigation) ToJSON() (i interface{}) {
	macro.Rewrite("$_.toJSON()")
	return i
}

func (*PerformanceNavigation) RedirectCount() (redirectCount uint) {
	macro.Rewrite("$_.redirectCount")
	return redirectCount
}

func (*PerformanceNavigation) Type() (kind uint) {
	macro.Rewrite("$_.type")
	return kind
}
