package payment

type PaymentShippingOption struct {
	amount		*PaymentCurrencyAmount
	id		string
	label		string
	selected	*bool
}
