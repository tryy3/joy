package paymentdetails

import (
	"github.com/matthewmueller/joy/dom/payment"
	"github.com/matthewmueller/joy/dom/paymentdetailsmodifier"
)

type PaymentDetails struct {
	displayItems    *[]*payment.PaymentItem
	err             *string
	modifiers       *[]*paymentdetailsmodifier.PaymentDetailsModifier
	shippingOptions *[]*payment.PaymentShippingOption
	total           *payment.PaymentItem
}
