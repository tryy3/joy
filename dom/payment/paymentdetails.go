package payment

type PaymentDetails struct {
	displayItems	*[]*PaymentItem
	err		*string
	modifiers	*[]*PaymentDetailsModifier
	shippingOptions	*[]*PaymentShippingOption
	total		*PaymentItem
}
