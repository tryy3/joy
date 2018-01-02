package window

import "github.com/matthewmueller/joy/macro"

type StyleMedia struct {
}

func (*StyleMedia) MatchMedium(mediaquery string) (b bool) {
	macro.Rewrite("$_.matchMedium($1)", mediaquery)
	return b
}

func (*StyleMedia) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}
