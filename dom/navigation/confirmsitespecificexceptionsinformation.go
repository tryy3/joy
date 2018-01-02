package navigation

import "github.com/matthewmueller/joy/dom/utils"

type ConfirmSiteSpecificExceptionsInformation struct {
	*utils.ExceptionInformation

	arrayOfDomainStrings	*[]string
}
