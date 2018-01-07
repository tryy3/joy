package pushmanager

import (
	"github.com/matthewmueller/joy/dom/push"
	"github.com/matthewmueller/joy/macro"
)

// PushManager struct
// js:"PushManager,omit"
type PushManager struct {
}

// GetSubscription fn
// js:"getSubscription"
func (*PushManager) GetSubscription() (p *push.PushSubscription) {
	macro.Rewrite("await $_.getSubscription()")
	return p
}

// PermissionState fn
// js:"permissionState"
func (*PushManager) PermissionState(options *push.PushSubscriptionOptionsInit) (p *push.PushPermissionState) {
	macro.Rewrite("await $_.permissionState($1)", options)
	return p
}

// Subscribe fn
// js:"subscribe"
func (*PushManager) Subscribe(options *push.PushSubscriptionOptionsInit) (p *push.PushSubscription) {
	macro.Rewrite("await $_.subscribe($1)", options)
	return p
}
