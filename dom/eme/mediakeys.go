package eme

import (
	"github.com/matthewmueller/joy/macro"
)

type MediaKeys struct {
}

func (*MediaKeys) CreateSession(sessionType *MediaKeySessionType) (m *MediaKeySession) {
	macro.Rewrite("$_.createSession($1)", sessionType)
	return m
}

func (*MediaKeys) SetServerCertificate(serverCertificate []byte) {
	macro.Rewrite("await $_.setServerCertificate($1)", serverCertificate)
}
