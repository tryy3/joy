package css

import "github.com/matthewmueller/joy/macro"

type MediaQueryList struct {
}

func (*MediaQueryList) AddListener(listener func(mql *MediaQueryList)) {
	macro.Rewrite("$_.addListener($1)", listener)
}

func (*MediaQueryList) RemoveListener(listener func(mql *MediaQueryList)) {
	macro.Rewrite("$_.removeListener($1)", listener)
}

func (*MediaQueryList) Matches() (matches bool) {
	macro.Rewrite("$_.matches")
	return matches
}

func (*MediaQueryList) Media() (media string) {
	macro.Rewrite("$_.media")
	return media
}
