package paymentdetailsmodifier

import "github.com/matthewmueller/joy/dom/payment"

type PaymentDetailsModifier struct {
	additionalDisplayItems *[]*payment.PaymentItem
	data                   *interface{}
	supportedMethods       []string
	total                  *payment.PaymentItem
}
