package payment

type PaymentOptions struct {
	requestPayerEmail	*bool
	requestPayerName	*bool
	requestPayerPhone	*bool
	requestShipping		*bool
	shippingType		*string
}
