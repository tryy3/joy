package push

import "github.com/matthewmueller/joy/macro"

type PushSubscriptionOptions struct {
}

func (*PushSubscriptionOptions) ApplicationServerKey() (applicationServerKey []byte) {
	macro.Rewrite("$_.applicationServerKey")
	return applicationServerKey
}

func (*PushSubscriptionOptions) UserVisibleOnly() (userVisibleOnly bool) {
	macro.Rewrite("$_.userVisibleOnly")
	return userVisibleOnly
}
