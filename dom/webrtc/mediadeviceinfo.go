package webrtc

import (
	"github.com/matthewmueller/joy/macro"
)

type MediaDeviceInfo struct {
}

func (*MediaDeviceInfo) DeviceID() (deviceId string) {
	macro.Rewrite("$_.deviceId")
	return deviceId
}

func (*MediaDeviceInfo) GroupID() (groupId string) {
	macro.Rewrite("$_.groupId")
	return groupId
}

func (*MediaDeviceInfo) Kind() (kind *MediaDeviceKind) {
	macro.Rewrite("$_.kind")
	return kind
}

func (*MediaDeviceInfo) Label() (label string) {
	macro.Rewrite("$_.label")
	return label
}
