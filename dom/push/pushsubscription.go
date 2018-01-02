package push

import (
	"github.com/matthewmueller/joy/macro"
)

type PushSubscription struct {
}

func (*PushSubscription) GetKey(name *PushEncryptionKeyName) (b []byte) {
	macro.Rewrite("$_.getKey($1)", name)
	return b
}

func (*PushSubscription) ToJSON() (i interface{}) {
	macro.Rewrite("$_.toJSON()")
	return i
}

func (*PushSubscription) Unsubscribe() (b bool) {
	macro.Rewrite("await $_.unsubscribe()")
	return b
}

func (*PushSubscription) Endpoint() (endpoint string) {
	macro.Rewrite("$_.endpoint")
	return endpoint
}

func (*PushSubscription) Options() (options *PushSubscriptionOptions) {
	macro.Rewrite("$_.options")
	return options
}
