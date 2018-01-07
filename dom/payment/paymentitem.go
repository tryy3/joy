package payment

type PaymentItem struct {
	amount  *PaymentCurrencyAmount
	label   string
	pending *bool
}
