package window

import "github.com/matthewmueller/joy/macro"

type History struct {
}

func (*History) Back(distance *interface{}) {
	macro.Rewrite("$_.back($1)", distance)
}

func (*History) Forward(distance *interface{}) {
	macro.Rewrite("$_.forward($1)", distance)
}

func (*History) Go(delta *interface{}) {
	macro.Rewrite("$_.go($1)", delta)
}

func (*History) PushState(statedata interface{}, title *string, url *string) {
	macro.Rewrite("$_.pushState($1, $2, $3)", statedata, title, url)
}

func (*History) ReplaceState(statedata interface{}, title *string, url *string) {
	macro.Rewrite("$_.replaceState($1, $2, $3)", statedata, title, url)
}

func (*History) Length() (length int) {
	macro.Rewrite("$_.length")
	return length
}

func (*History) State() (state interface{}) {
	macro.Rewrite("$_.state")
	return state
}
