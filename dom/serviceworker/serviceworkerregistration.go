package serviceworker

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/notification"
	"github.com/matthewmueller/joy/dom/push"
)

var _ event.EventTarget = (*ServiceWorkerRegistration)(nil)

type ServiceWorkerRegistration struct {
}

func (*ServiceWorkerRegistration) GetNotifications(filter *GetNotificationOptions) (n []*notification.Notification) {
	macro.Rewrite("await $_.getNotifications($1)", filter)
	return n
}

func (*ServiceWorkerRegistration) ShowNotification(title string, options *notification.NotificationOptions) {
	macro.Rewrite("await $_.showNotification($1, $2)", title, options)
}

func (*ServiceWorkerRegistration) Unregister() (b bool) {
	macro.Rewrite("await $_.unregister()")
	return b
}

func (*ServiceWorkerRegistration) Update() {
	macro.Rewrite("await $_.update()")
}

func (*ServiceWorkerRegistration) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*ServiceWorkerRegistration) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*ServiceWorkerRegistration) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*ServiceWorkerRegistration) Active() (active *ServiceWorker) {
	macro.Rewrite("$_.active")
	return active
}

func (*ServiceWorkerRegistration) Installing() (installing *ServiceWorker) {
	macro.Rewrite("$_.installing")
	return installing
}

func (*ServiceWorkerRegistration) Onupdatefound() (onupdatefound func(event.Event)) {
	macro.Rewrite("$_.onupdatefound")
	return onupdatefound
}

func (*ServiceWorkerRegistration) SetOnupdatefound(onupdatefound func(event.Event)) {
	macro.Rewrite("$_.onupdatefound = $1", onupdatefound)
}

func (*ServiceWorkerRegistration) PushManager() (pushManager *push.PushManager) {
	macro.Rewrite("$_.pushManager")
	return pushManager
}

func (*ServiceWorkerRegistration) Scope() (scope string) {
	macro.Rewrite("$_.scope")
	return scope
}

func (*ServiceWorkerRegistration) Sync() (sync *SyncManager) {
	macro.Rewrite("$_.sync")
	return sync
}

func (*ServiceWorkerRegistration) Waiting() (waiting *ServiceWorker) {
	macro.Rewrite("$_.waiting")
	return waiting
}
