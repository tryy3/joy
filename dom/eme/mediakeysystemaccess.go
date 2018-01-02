package eme

import (
	"github.com/matthewmueller/joy/macro"
)

type MediaKeySystemAccess struct {
}

func (*MediaKeySystemAccess) CreateMediaKeys() (m *MediaKeys) {
	macro.Rewrite("await $_.createMediaKeys()")
	return m
}

func (*MediaKeySystemAccess) GetConfiguration() (m *MediaKeySystemConfiguration) {
	macro.Rewrite("$_.getConfiguration()")
	return m
}

func (*MediaKeySystemAccess) KeySystem() (keySystem string) {
	macro.Rewrite("$_.keySystem")
	return keySystem
}
