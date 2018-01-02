package payment

import "github.com/matthewmueller/joy/macro"

type PaymentAddress struct {
}

func (*PaymentAddress) ToJSON() (i interface{}) {
	macro.Rewrite("$_.toJSON()")
	return i
}

func (*PaymentAddress) AddressLine() (addressLine []string) {
	macro.Rewrite("$_.addressLine")
	return addressLine
}

func (*PaymentAddress) City() (city string) {
	macro.Rewrite("$_.city")
	return city
}

func (*PaymentAddress) Country() (country string) {
	macro.Rewrite("$_.country")
	return country
}

func (*PaymentAddress) DependentLocality() (dependentLocality string) {
	macro.Rewrite("$_.dependentLocality")
	return dependentLocality
}

func (*PaymentAddress) LanguageCode() (languageCode string) {
	macro.Rewrite("$_.languageCode")
	return languageCode
}

func (*PaymentAddress) Organization() (organization string) {
	macro.Rewrite("$_.organization")
	return organization
}

func (*PaymentAddress) Phone() (phone string) {
	macro.Rewrite("$_.phone")
	return phone
}

func (*PaymentAddress) PostalCode() (postalCode string) {
	macro.Rewrite("$_.postalCode")
	return postalCode
}

func (*PaymentAddress) Recipient() (recipient string) {
	macro.Rewrite("$_.recipient")
	return recipient
}

func (*PaymentAddress) Region() (region string) {
	macro.Rewrite("$_.region")
	return region
}

func (*PaymentAddress) SortingCode() (sortingCode string) {
	macro.Rewrite("$_.sortingCode")
	return sortingCode
}
