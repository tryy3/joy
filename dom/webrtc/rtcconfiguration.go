package webrtc

type RTCConfiguration struct {
	bundlePolicy       *RTCBundlePolicy
	iceServers         *[]*RTCIceServer
	iceTransportPolicy *RTCIceTransportPolicy
	peerIdentity       *string
}
