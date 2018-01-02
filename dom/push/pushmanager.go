package push

import (
	"github.com/matthewmueller/joy/macro"
)

type PushManager struct {
}

func (*PushManager) GetSubscription() (p *PushSubscription) {
	macro.Rewrite("await $_.getSubscription()")
	return p
}

func (*PushManager) PermissionState(options *PushSubscriptionOptionsInit) (p *PushPermissionState) {
	macro.Rewrite("await $_.permissionState($1)", options)
	return p
}

func (*PushManager) Subscribe(options *PushSubscriptionOptionsInit) (p *PushSubscription) {
	macro.Rewrite("await $_.subscribe($1)", options)
	return p
}
