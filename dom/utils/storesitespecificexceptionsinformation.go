package utils


import "github.com/matthewmueller/joy/dom/storeexceptionsinformation"

type StoreSiteSpecificExceptionsInformation struct {
	*storeexceptionsinformation.StoreExceptionsInformation

	arrayOfDomainStrings *[]string
}
