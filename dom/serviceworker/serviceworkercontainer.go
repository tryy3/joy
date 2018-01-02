package serviceworker

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
)

var _ event.EventTarget = (*ServiceWorkerContainer)(nil)

type ServiceWorkerContainer struct {
}

func (*ServiceWorkerContainer) GetRegistration(clientURL *string) (i interface{}) {
	macro.Rewrite("await $_.getRegistration($1)", clientURL)
	return i
}

func (*ServiceWorkerContainer) GetRegistrations() (s []*ServiceWorkerRegistration) {
	macro.Rewrite("await $_.getRegistrations()")
	return s
}

func (*ServiceWorkerContainer) Register(scriptURL string, options *RegistrationOptions) (s *ServiceWorkerRegistration) {
	macro.Rewrite("await $_.register($1, $2)", scriptURL, options)
	return s
}

func (*ServiceWorkerContainer) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*ServiceWorkerContainer) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*ServiceWorkerContainer) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*ServiceWorkerContainer) Controller() (controller *ServiceWorker) {
	macro.Rewrite("$_.controller")
	return controller
}

func (*ServiceWorkerContainer) Oncontrollerchange() (oncontrollerchange func(event.Event)) {
	macro.Rewrite("$_.oncontrollerchange")
	return oncontrollerchange
}

func (*ServiceWorkerContainer) SetOncontrollerchange(oncontrollerchange func(event.Event)) {
	macro.Rewrite("$_.oncontrollerchange = $1", oncontrollerchange)
}

func (*ServiceWorkerContainer) Onmessage() (onmessage func(*ServiceWorkerMessageEvent)) {
	macro.Rewrite("$_.onmessage")
	return onmessage
}

func (*ServiceWorkerContainer) SetOnmessage(onmessage func(*ServiceWorkerMessageEvent)) {
	macro.Rewrite("$_.onmessage = $1", onmessage)
}

func (*ServiceWorkerContainer) Ready() (ready *ServiceWorkerRegistration) {
	macro.Rewrite("$_.ready")
	return ready
}
