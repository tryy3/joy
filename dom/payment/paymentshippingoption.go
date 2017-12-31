package payment


import "github.com/matthewmueller/joy/dom/paymentcurrencyamount"

type PaymentShippingOption struct {
	amount   *paymentcurrencyamount.PaymentCurrencyAmount
	id       string
	label    string
	selected *bool
}
