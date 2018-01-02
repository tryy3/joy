package webrtc

import (
	"github.com/matthewmueller/joy/macro"
)

func New(descriptioninitdict *RTCSessionDescriptionInit) *RTCSessionDescription {
	macro.Rewrite("new RTCSessionDescription($1)", descriptioninitdict)
	return &RTCSessionDescription{}
}

type RTCSessionDescription struct {
}

func (*RTCSessionDescription) ToJSON() (i interface{}) {
	macro.Rewrite("$_.toJSON()")
	return i
}

func (*RTCSessionDescription) Sdp() (sdp string) {
	macro.Rewrite("$_.sdp")
	return sdp
}

func (*RTCSessionDescription) SetSdp(sdp string) {
	macro.Rewrite("$_.sdp = $1", sdp)
}

func (*RTCSessionDescription) Type() (kind *RTCSdpType) {
	macro.Rewrite("$_.type")
	return kind
}

func (*RTCSessionDescription) SetType(kind *RTCSdpType) {
	macro.Rewrite("$_.type = $1", kind)
}
