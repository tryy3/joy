package webrtc

type RTCSrtpSdesParameters struct {
	cryptoSuite	*string
	keyParams	*[]*RTCSrtpKeyParam
	sessionParams	*[]string
	tag		*uint8
}
