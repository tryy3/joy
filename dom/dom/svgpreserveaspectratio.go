package dom

import "github.com/matthewmueller/joy/macro"

type SVGPreserveAspectRatio struct {
}

func (*SVGPreserveAspectRatio) Align() (align uint8) {
	macro.Rewrite("$_.align")
	return align
}

func (*SVGPreserveAspectRatio) SetAlign(align uint8) {
	macro.Rewrite("$_.align = $1", align)
}

func (*SVGPreserveAspectRatio) MeetOrSlice() (meetOrSlice uint8) {
	macro.Rewrite("$_.meetOrSlice")
	return meetOrSlice
}

func (*SVGPreserveAspectRatio) SetMeetOrSlice(meetOrSlice uint8) {
	macro.Rewrite("$_.meetOrSlice = $1", meetOrSlice)
}
