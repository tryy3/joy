package storesitespecificexceptionsinformation

import "github.com/matthewmueller/joy/dom/utils"

type StoreSiteSpecificExceptionsInformation struct {
	*utils.StoreExceptionsInformation

	arrayOfDomainStrings *[]string
}
