package payment

import (
	"github.com/matthewmueller/joy/macro"
)

type PaymentResponse struct {
}

func (*PaymentResponse) Complete(result *PaymentComplete) {
	macro.Rewrite("await $_.complete($1)", result)
}

func (*PaymentResponse) ToJSON() (i interface{}) {
	macro.Rewrite("$_.toJSON()")
	return i
}

func (*PaymentResponse) Details() (details interface{}) {
	macro.Rewrite("$_.details")
	return details
}

func (*PaymentResponse) MethodName() (methodName string) {
	macro.Rewrite("$_.methodName")
	return methodName
}

func (*PaymentResponse) PayerEmail() (payerEmail string) {
	macro.Rewrite("$_.payerEmail")
	return payerEmail
}

func (*PaymentResponse) PayerName() (payerName string) {
	macro.Rewrite("$_.payerName")
	return payerName
}

func (*PaymentResponse) PayerPhone() (payerPhone string) {
	macro.Rewrite("$_.payerPhone")
	return payerPhone
}

func (*PaymentResponse) ShippingAddress() (shippingAddress *PaymentAddress) {
	macro.Rewrite("$_.shippingAddress")
	return shippingAddress
}

func (*PaymentResponse) ShippingOption() (shippingOption string) {
	macro.Rewrite("$_.shippingOption")
	return shippingOption
}
