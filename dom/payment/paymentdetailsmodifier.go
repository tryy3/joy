package payment

type PaymentDetailsModifier struct {
	additionalDisplayItems	*[]*PaymentItem
	data			*interface{}
	supportedMethods	[]string
	total			*PaymentItem
}
