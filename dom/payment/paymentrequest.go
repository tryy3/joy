package payment

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
)

var _ event.EventTarget = (*PaymentRequest)(nil)

func New(methoddata []*PaymentMethodData, details *PaymentDetails, options *PaymentOptions) *PaymentRequest {
	macro.Rewrite("new PaymentRequest($1, $2, $3)", methoddata, details, options)
	return &PaymentRequest{}
}

type PaymentRequest struct {
}

func (*PaymentRequest) Abort() {
	macro.Rewrite("await $_.abort()")
}

func (*PaymentRequest) Show() (p *PaymentResponse) {
	macro.Rewrite("await $_.show()")
	return p
}

func (*PaymentRequest) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*PaymentRequest) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*PaymentRequest) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*PaymentRequest) Onshippingaddresschange() (onshippingaddresschange func(event.Event)) {
	macro.Rewrite("$_.onshippingaddresschange")
	return onshippingaddresschange
}

func (*PaymentRequest) SetOnshippingaddresschange(onshippingaddresschange func(event.Event)) {
	macro.Rewrite("$_.onshippingaddresschange = $1", onshippingaddresschange)
}

func (*PaymentRequest) Onshippingoptionchange() (onshippingoptionchange func(event.Event)) {
	macro.Rewrite("$_.onshippingoptionchange")
	return onshippingoptionchange
}

func (*PaymentRequest) SetOnshippingoptionchange(onshippingoptionchange func(event.Event)) {
	macro.Rewrite("$_.onshippingoptionchange = $1", onshippingoptionchange)
}

func (*PaymentRequest) ShippingAddress() (shippingAddress *PaymentAddress) {
	macro.Rewrite("$_.shippingAddress")
	return shippingAddress
}

func (*PaymentRequest) ShippingOption() (shippingOption string) {
	macro.Rewrite("$_.shippingOption")
	return shippingOption
}

func (*PaymentRequest) ShippingType() (shippingType *PaymentShippingType) {
	macro.Rewrite("$_.shippingType")
	return shippingType
}
