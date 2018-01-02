package webrtc

import (
	"github.com/matthewmueller/joy/macro"
)

func New(candidateinitdict *RTCIceCandidateInit) *RTCIceCandidate {
	macro.Rewrite("new RTCIceCandidate($1)", candidateinitdict)
	return &RTCIceCandidate{}
}

type RTCIceCandidate struct {
}

func (*RTCIceCandidate) ToJSON() (i interface{}) {
	macro.Rewrite("$_.toJSON()")
	return i
}

func (*RTCIceCandidate) Candidate() (candidate string) {
	macro.Rewrite("$_.candidate")
	return candidate
}

func (*RTCIceCandidate) SetCandidate(candidate string) {
	macro.Rewrite("$_.candidate = $1", candidate)
}

func (*RTCIceCandidate) SdpMid() (sdpMid string) {
	macro.Rewrite("$_.sdpMid")
	return sdpMid
}

func (*RTCIceCandidate) SetSdpMid(sdpMid string) {
	macro.Rewrite("$_.sdpMid = $1", sdpMid)
}

func (*RTCIceCandidate) SdpMLineIndex() (sdpMLineIndex uint8) {
	macro.Rewrite("$_.sdpMLineIndex")
	return sdpMLineIndex
}

func (*RTCIceCandidate) SetSdpMLineIndex(sdpMLineIndex uint8) {
	macro.Rewrite("$_.sdpMLineIndex = $1", sdpMLineIndex)
}
