package ms

import (
	"github.com/matthewmueller/joy/macro"
)

func New(keysystem string) *MSMediaKeys {
	macro.Rewrite("new MSMediaKeys($1)", keysystem)
	return &MSMediaKeys{}
}

type MSMediaKeys struct {
}

func (*MSMediaKeys) CreateSession(kind string, initData []uint8, cdmData *[]uint8) (m *MSMediaKeySession) {
	macro.Rewrite("$_.createSession($1, $2, $3)", kind, initData, cdmData)
	return m
}

func (*MSMediaKeys) IsTypeSupported(keySystem string, kind *string) (b bool) {
	macro.Rewrite("$_.isTypeSupported($1, $2)", keySystem, kind)
	return b
}

func (*MSMediaKeys) IsTypeSupportedWithFeatures(keySystem string, kind *string) (s string) {
	macro.Rewrite("$_.isTypeSupportedWithFeatures($1, $2)", keySystem, kind)
	return s
}

func (*MSMediaKeys) KeySystem() (keySystem string) {
	macro.Rewrite("$_.keySystem")
	return keySystem
}
